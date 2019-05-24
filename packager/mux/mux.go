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
)

type mux struct {
	certCache   http.Handler
	signer      http.Handler
	validityMap http.Handler
}

// The main entry point. Use the return value for http.Server.Handler.
func New(certCache http.Handler, signer http.Handler, validityMap http.Handler) http.Handler {
	return &mux{certCache, signer, validityMap}
}

func tryTrimPrefix(s, prefix string) (string, bool) {
	sLen := len(s)
	trimmed := strings.TrimPrefix(s, prefix)
	return trimmed, len(trimmed) != sLen
}

var allowedMethods = map[string]bool{http.MethodGet: true, http.MethodHead: true}

func (this *mux) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if !allowedMethods[req.Method] {
		http.Error(resp, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	params := map[string]string{}
	req = WithParams(req, params)

	// Use EscapedPath rather than RequestURI because the latter can take
	// absolute-form, per https://tools.ietf.org/html/rfc7230#section-5.3.
	//
	// Use EscapedPath rather than RawPath to verify that the request URI
	// is a valid encoding, per
	// https://wicg.github.io/webpackage/draft-yasskin-http-origin-signed-responses.html#application-signed-exchange
	// item 3.
	path := req.URL.EscapedPath()
	if suffix, ok := tryTrimPrefix(path, "/priv/doc"); ok {
		if suffix == "" {
			this.signer.ServeHTTP(resp, req)
		} else if suffix[0] == '/' {
			params["signURL"] = suffix[1:]
			if req.URL.RawQuery != "" {
				params["signURL"] += "?" + req.URL.RawQuery
			}
			this.signer.ServeHTTP(resp, req)
		} else {
			http.NotFound(resp, req)
		}
	} else if suffix, ok := tryTrimPrefix(path, util.CertURLPrefix+"/"); ok {
		unescaped, err := url.PathUnescape(suffix)
		if err != nil {
			http.Error(resp, "400 bad request - bad URL encoding", http.StatusBadRequest)
		} else {
			params["certName"] = unescaped
			this.certCache.ServeHTTP(resp, req)
		}
	} else if path == util.ValidityMapPath {
		this.validityMap.ServeHTTP(resp, req)
	} else {
		http.NotFound(resp, req)
	}
}

type paramsKeyType struct{}

var paramsKey = paramsKeyType{}

// Gets the params from the request context, injected by the mux. Guaranteed to
// be non-nil. Call from the handlers.
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

// Returns a copy of req annotated with the given params. (Params is stored by
// reference, and so may be mutated afterwards.) To be called only by this
// library itself or by tests.
func WithParams(req *http.Request, params map[string]string) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), paramsKey, params))
}
