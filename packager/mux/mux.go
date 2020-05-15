// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Implements the HTTP routing strategy used by the amppkg server. The main
// purpose for rolling our own is to be able to route the following types of
// URLs:
//   /priv/doc/https://example.com/esc%61ped%2Furl.html
// and have the signer sign the URL exactly as it is encoded in the request, in
// order to meet docs/cache_requirements.md.
//
// The default http mux had some problems which led to me investigating 3rd
// party routers. (TODO(twifkak): Remember and document those problems.)
//
// I investigated two 3rd party routers capable of handling catch-all suffix
// parameters:
//   https://github.com/julienschmidt/httprouter
//   https://github.com/dimfeld/httptreemux
// Both libs unescape their catch-all parameters, making the above use-case
// impossible. The latter does so despite documenting support for unmodified
// URL escapings. This, plus a lack of feature needs, led me to believe that
// writing our own mux was the best approach.
package mux

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/ampproject/amppackager/packager/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// routingRule maps a URL path prefix to three entities:
// * suffixValidatorFunc - a function that validates the suffix of URL path,
// * handler - an http.Handler that should handle such prefix,
// * handlerPrometheusLabel - a label (dimension) to be used in
//       handler-agnostic Prometheus metrics like requests count.
//       Must adhere to the Prometheus data model:
// 		 https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels
type routingRule struct {
	urlPathPrefix          string
	suffixValidatorFunc    func(suffix string, req *http.Request, params *map[string]string, errorMsg *string, errorCode *int)
	handler                http.Handler
	handlerPrometheusLabel string
}

// mux stores a routingMatrix, an array of routing rules that define the mux'
// routing logic. Note that the order of rules in the matrix is important: the
// first matching rule will be applied by ServeHTTP, so the rule for a
// particular prefix should precede the rule for it's sub-prefix. E.g. the rule
// for routing requests prefixed with “/priv/doc/” (w/ trailing slash) must go
// before the rule for routing requests prefixed with “/priv/doc” (w/o trailing
// slash).
// The last rule in the matrix must have empty urlPathPrefix, so it would match
// any URL. This ensures there's at least one matching rule for any URL.
type mux struct {
	routingMatrix []routingRule
	defaultRule   routingRule
}

// return404 is a URL Path Suffix Validator that always returns 404.
func return404(suffix string, req *http.Request, params *map[string]string, errorMsg *string, errorCode *int) {
	*errorMsg, *errorCode = "404 page not found", http.StatusNotFound
}

// expectNoSuffix is a URL Path Suffix Validator that expects an empty suffix.
func expectNoSuffix(suffix string, req *http.Request, params *map[string]string, errorMsg *string, errorCode *int) {
	if suffix != "" {
		return404(suffix, req, params, errorMsg, errorCode)
	}
}

// expectSignerQuery is a URL Path Suffix Validator specific to signer requests.
func expectSignerQuery(suffix string, req *http.Request, params *map[string]string, errorMsg *string, errorCode *int) {
	(*params)["signURL"] = suffix
	if req.URL.RawQuery != "" {
		(*params)["signURL"] += "?" + req.URL.RawQuery
	}
}

// expectCertQuery is a URL Path Suffix Validator specific to cert requests.
func expectCertQuery(suffix string, req *http.Request, params *map[string]string, errorMsg *string, errorCode *int) {
	unescaped, err := url.PathUnescape(suffix)
	if err != nil {
		*errorMsg, *errorCode = "400 bad request - bad URL encoding", http.StatusBadRequest
	} else {
		(*params)["certName"] = unescaped
	}
}

// New is the main entry point. Use the return value for http.Server.Handler.
func New(certCache http.Handler, signer http.Handler, validityMap http.Handler, healthz http.Handler, metrics http.Handler) http.Handler {
	return &mux{
		// Note that the order of rules in the matrix matters: the first
		// matching rule will be applied, so the rule for “/priv/doc/” precedes
		// the rule for “/priv/doc” (note that SignerURLPrefix is "/priv/doc").
		// Also note that the last rule matches any URL.
		[]routingRule{
			{util.SignerURLPrefix + "/", expectSignerQuery, signer, "signer"},
			{util.SignerURLPrefix, expectNoSuffix, signer, "signer"},
			{util.CertURLPrefix + "/", expectCertQuery, certCache, "certCache"},
			{util.ValidityMapPath, expectNoSuffix, validityMap, "validityMap"},
			{util.HealthzPath, expectNoSuffix, healthz, "healthz"},
			{util.MetricsPath, expectNoSuffix, metrics, "metrics"},
		},
		/* defaultRule= */ routingRule{"", return404, nil, "handler_not_assigned"},
	}
}

var promTotalRequests = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "total_requests_by_code_and_url",
		Help: "Total number of requests by HTTP code and URL.",
	},
	[]string{"code", "handler"},
)

// tryTrimPrefix trims prefix off s if s starts with prefix, and keeps s as is
// otherwise. It returns the remaining suffix and a boolean success indicator.
// If prefix is an empty string, returns s and "true".
func tryTrimPrefix(s, prefix string) (string, bool) {
	sLen := len(s)
	trimmed := strings.TrimPrefix(s, prefix)
	return trimmed, len(prefix)+len(trimmed) == sLen
}

var allowedMethods = map[string]bool{http.MethodGet: true, http.MethodHead: true}

func (this *mux) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// Use EscapedPath rather than RequestURI because the latter can take
	// absolute-form, per https://tools.ietf.org/html/rfc7230#section-5.3.
	//
	// Use EscapedPath rather than RawPath to verify that the request URI
	// is a valid encoding, per
	// https://wicg.github.io/webpackage/draft-yasskin-http-origin-signed-responses.html#application-signed-exchange
	// item 3.
	path := req.URL.EscapedPath()

	// Find the first matching routing rule.
	var matchingRule *routingRule
	var suffix string
	for _, currentRule := range this.routingMatrix {
		if currentSuffix, isMatch := tryTrimPrefix(path, currentRule.urlPathPrefix); isMatch {
			matchingRule = &currentRule
			suffix = currentSuffix
			break
		}
	}

	if matchingRule == nil {
		matchingRule = &this.defaultRule
	}

	errorMsg := ""
	errorCode := 0
	// Validate HTTP method and params, parse params and attach them to req.
	if !allowedMethods[req.Method] {
		errorMsg, errorCode = "405 method not allowed", http.StatusMethodNotAllowed
	} else {
		params := map[string]string{}
		req = WithParams(req, params)
		matchingRule.suffixValidatorFunc(suffix, req, &params, &errorMsg, &errorCode)
	}

	// Prepare the handler.
	var handlerFunc http.Handler
	if errorCode == 0 {
		handlerFunc = matchingRule.handler
	} else {
		handlerFunc = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.Error(w, errorMsg, errorCode) })
	}

	// Decorate the call to handlerFunc with a Prometheus requests counter
	// pre-labelled (curried) with the right handler label.
	promhttp.InstrumentHandlerCounter(promTotalRequests.MustCurryWith(
		prometheus.Labels{"handler": matchingRule.handlerPrometheusLabel}),
		handlerFunc).ServeHTTP(resp, req)
}

type paramsKeyType struct{}

var paramsKey = paramsKeyType{}

// Params gets the params from the request context, injected by the mux.
// Guaranteed to be non-nil. Call from the handlers.
func Params(req *http.Request) map[string]string {
	params := req.Context().Value(paramsKey)
	switch v := params.(type) {
	case map[string]string:
		return v
	default:
		// This should never happen, but just in case, let's not panic.
		return map[string]string{}
	}
}

// WithParams returns a copy of req annotated with the given params. (Params is
// stored by reference, and so may be mutated afterwards.) To be called only by
// this library itself or by tests.
func WithParams(req *http.Request, params map[string]string) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), paramsKey, params))
}
