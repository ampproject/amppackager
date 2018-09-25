// Copyright 2018 Google LLC
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

package signer

import (
	"bytes"
	"crypto"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/WICG/webpackage/go/signedexchange/version"
	"github.com/ampproject/amppackager/packager/accept"
	"github.com/ampproject/amppackager/packager/amp_cache_transform"
	"github.com/ampproject/amppackager/packager/rtv"
	"github.com/ampproject/amppackager/packager/util"
	"github.com/ampproject/amppackager/transformer"
	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/pquerna/cachecontrol"
)

// The Content-Security-Policy in use by the AMP Cache today. Specifying here
// provides protection for the publisher against bugs in the transformers, as
// these pages will now run on the publisher's origin. In the future, this
// value will likely be versioned along with the transforms.
var contentSecurityPolicy = "default-src * blob: data:; script-src blob: https://cdn.ampproject.org/rtv/ https://cdn.ampproject.org/v0.js https://cdn.ampproject.org/v0/ https://cdn.ampproject.org/viewer/; object-src 'none'; style-src 'unsafe-inline' https://cdn.ampproject.org/rtv/ https://cdn.materialdesignicons.com https://cloud.typography.com https://fast.fonts.net https://fonts.googleapis.com https://maxcdn.bootstrapcdn.com https://p.typekit.net https://pro.fontawesome.com https://use.fontawesome.com https://use.typekit.net; report-uri https://csp-collector.appspot.com/csp/amp"

// The user agent to send when issuing fetches. Should look like a mobile device.
const userAgent = "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) " +
	"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile " +
	"Safari/537.36 (compatible; amppackager/0.0.0; +https://github.com/ampproject/amppackager)"

// Conditional request headers that ServeHTTP may receive and need to be sent with fetchURL.
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Conditional_requests#Conditional_headers
var conditionalRequestHeaders = map[string]bool{
	"If-Match":            true,
	"If-None-Match":       true,
	"If-Modified-Since":   true,
	"If-Unmodified-Since": true,
	"If-Range":            true,
}

// Advised against, per
// https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00#section-4.1
// and blocked in http://crrev.com/c/958945.
var statefulResponseHeaders = map[string]bool{
	"Authentication-Control":    true,
	"Authentication-Info":       true,
	"Optional-WWW-Authenticate": true,
	"Proxy-Authenticate":        true,
	"Proxy-Authentication-Info": true,
	"Sec-WebSocket-Accept":      true,
	"Set-Cookie":                true,
	"Set-Cookie2":               true,
	"SetProfile":                true,
	"WWW-Authenticate":          true,
}

// The server generating a 304 response MUST generate any of the
// following header fields that would have been sent in a 200 (OK) response
// to the same request.
// https://tools.ietf.org/html/rfc7232#section-4.1
var statusNotModifiedHeaders = map[string]bool{
	"Cache-Control":    true,
	"Content-Location": true,
	"Date":             true,
	"ETag":             true,
	"Expires":          true,
	"Vary":             true,
}

// TODO(twifkak): Remove this restriction by allowing streamed responses from the signedexchange library.
const maxBodyLength = 4 * 1 << 20

// TODO(twifkak): What value should this be?
const miRecordSize = 4096

// Overrideable for testing.
var getTransformerRequest = func(r *rtv.RTVCache, s, u string) *rpb.Request {
	return &rpb.Request{Html: string(s), DocumentUrl: u, Rtv: r.GetRTV(), Css: r.GetCSS()}
}

func parseURL(rawURL string, name string) (*url.URL, *util.HTTPError) {
	if rawURL == "" {
		return nil, util.NewHTTPError(http.StatusBadRequest, name, " URL is unspecified")
	}
	ret, err := url.Parse(rawURL)
	if err != nil {
		return nil, util.NewHTTPError(http.StatusBadRequest, "Error parsing ", name, " url: ", err)
	}
	if !ret.IsAbs() {
		return nil, util.NewHTTPError(http.StatusBadRequest, name, " url is relative")
	}
	// Evaluate "/..", by resolving the URL as a reference from itself.
	// This prevents malformed URLs from eluding the PathRE protections.
	ret = ret.ResolveReference(ret)
	return ret, nil
}

func regexpFullMatch(pattern string, test string) bool {
	// This is how regexp/exec_test.go turns a partial pattern into a full pattern.
	fullRe := `\A(?:` + pattern + `)\z`
	matches, _ := regexp.MatchString(fullRe, test)
	return matches
}

func urlMatches(url *url.URL, pattern util.URLPattern) bool {
	if url.Opaque != "" {
		return false
	}
	if url.User != nil {
		return false
	}
	if !regexpFullMatch(*pattern.PathRE, url.EscapedPath()) {
		return false
	}
	for _, re := range pattern.PathExcludeRE {
		if regexpFullMatch(re, url.EscapedPath()) {
			return false
		}
	}
	if !regexpFullMatch(*pattern.QueryRE, url.RawQuery) {
		return false
	}
	return true
}

func fetchUrlMatches(url *url.URL, pattern *util.URLPattern) bool {
	if pattern == nil {
		// If URLSet.Fetch was unspecified, then so should ?fetch= be.
		return url == nil
	}
	schemeMatches := false
	for _, scheme := range pattern.Scheme {
		if url.Scheme == scheme {
			schemeMatches = true
		}
	}
	if !schemeMatches {
		return false
	}
	if pattern.Domain != "" && url.Host != pattern.Domain {
		return false
	}
	if pattern.DomainRE != "" && !regexpFullMatch(pattern.DomainRE, url.Host) {
		return false
	}
	return urlMatches(url, *pattern)
}

func signUrlMatches(url *url.URL, pattern *util.URLPattern) bool {
	if url.Scheme != "https" {
		return false
	}
	if url.Host != pattern.Domain {
		return false
	}
	return urlMatches(url, *pattern)
}

func urlsMatch(fetchURL *url.URL, signURL *url.URL, set util.URLSet) bool {
	fetchOK := fetchUrlMatches(fetchURL, set.Fetch)
	signOK := signUrlMatches(signURL, set.Sign)
	theyMatch := set.Fetch == nil || !*set.Fetch.SamePath || fetchURL.RequestURI() == signURL.RequestURI()
	return fetchOK && signOK && theyMatch
}

// Returns parsed URLs and whether to fail on stateful headers.
func parseURLs(fetch string, sign string, urlSets []util.URLSet) (*url.URL, *url.URL, bool, *util.HTTPError) {
	var fetchURL *url.URL
	var err *util.HTTPError
	if fetch != "" {
		fetchURL, err = parseURL(fetch, "fetch")
		if err != nil {
			return nil, nil, false, err
		}
	}
	signURL, err := parseURL(sign, "sign")
	if err != nil {
		return nil, nil, false, err
	}
	for _, set := range urlSets {
		if urlsMatch(fetchURL, signURL, set) {
			if fetchURL == nil {
				fetchURL = signURL
			}
			return fetchURL, signURL, set.Sign.ErrorOnStatefulHeaders, nil
		}
	}
	return nil, nil, false, util.NewHTTPError(http.StatusBadRequest, "fetch/sign URLs do not match config")
}

func validateFetch(req *http.Request, resp *http.Response) error {
	// Validate response is publicly-cacheable, per
	// https://tools.ietf.org/html/draft-yasskin-http-origin-signed-responses-03#section-6.1, as referenced by
	// https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00#section-6.
	nonCachableReasons, _, err := cachecontrol.CachableResponse(req, resp, cachecontrol.Options{PrivateCache: false})
	if err != nil {
		return errors.Wrap(err, "Error parsing cache headers.")
	}
	if len(nonCachableReasons) > 0 {
		return errors.Errorf("Non-cacheable response: %s", nonCachableReasons)
	}

	// Validate that Content-Type seems right. This is an approximation,
	// because parsing media types is hard and we just want to verify we're
	// not misinterpreting the server's intent. We override the
	// Content-Type below for unambiguous interpretation.
	content_type := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(strings.ToLower(content_type), "text/html") {
		return errors.Errorf("Wrong content-type: %s", content_type)
	}
	return nil
}

type Packager struct {
	// TODO(twifkak): Support multiple certs. This will require generating
	// a signature for each one. Note that Chrome only supports 1 signature
	// at the moment.
	cert *x509.Certificate
	// TODO(twifkak): Do we want to allow multiple keys?
	key             crypto.PrivateKey
	client          *http.Client
	urlSets         []util.URLSet
	rtvCache        *rtv.RTVCache
	shouldPackage   func() bool
	overrideBaseURL *url.URL
	requireHeaders  bool
}

func noRedirects(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func New(cert *x509.Certificate, key crypto.PrivateKey, urlSets []util.URLSet,
	rtvCache *rtv.RTVCache, shouldPackage func() bool, overrideBaseURL *url.URL,
	requireHeaders bool) (*Packager, error) {
	client := http.Client{
		CheckRedirect: noRedirects,
		// TODO(twifkak): Load-test and see if default transport settings are okay.
		Timeout: 60 * time.Second,
	}

	return &Packager{cert, key, &client, urlSets, rtvCache, shouldPackage, overrideBaseURL, requireHeaders}, nil
}

func (this *Packager) fetchURL(fetch *url.URL, serveHTTPReq http.Header) (*http.Request, *http.Response, *util.HTTPError) {
	ampURL := fetch.String()

	log.Printf("Fetching URL: %q\n", ampURL)
	req, err := http.NewRequest(http.MethodGet, ampURL, nil)
	// TODO(twifkak): Should we add 'Accept-Charset: utf-8'? Do AMP Caches require it? Will it break more servers than it fixes?
	if err != nil {
		return nil, nil, util.NewHTTPError(http.StatusInternalServerError, "Error building request: ", err)
	}
	req.Header.Set("User-Agent", userAgent)
	// Set conditional headers that were included in ServeHTTP's Request.
	for header := range conditionalRequestHeaders {
		if serveHTTPReq.Get(header) != "" {
			req.Header.Set(header, serveHTTPReq.Get(header))
		}
	}
	resp, err := this.client.Do(req)
	if err != nil {
		return nil, nil, util.NewHTTPError(http.StatusBadGateway, "Error fetching: ", err)
	}
	return req, resp, nil
}

func (this *Packager) genCertURL(cert *x509.Certificate, signURL *url.URL) (*url.URL, error) {
	var baseURL *url.URL
	if this.overrideBaseURL != nil {
		baseURL = this.overrideBaseURL
	} else {
		baseURL = signURL
	}
	urlPath := path.Join(util.CertURLPrefix, url.PathEscape(util.CertName(cert)))
	certHRef, err := url.Parse(urlPath)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing cert URL %q", urlPath)
	}
	ret := baseURL.ResolveReference(certHRef)
	return ret, nil
}

func (this *Packager) ServeHTTP(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// TODO(twifkak): See if there are any other validations or sanitizations that need adding.
	if err := req.ParseForm(); err != nil {
		util.NewHTTPError(http.StatusBadRequest, "Form input parsing failed: ", err).LogAndRespond(resp)
		return
	}
	var fetch, sign string
	if inPathSignURL := params.ByName("signURL"); inPathSignURL != "" {
		sign = inPathSignURL[1:] // Strip leading "/" produced by httprouter.
		if req.URL.RawQuery != "" {
			sign += "?" + req.URL.RawQuery
		}
	} else {
		if len(req.Form["fetch"]) > 1 {
			util.NewHTTPError(http.StatusBadRequest, "More than 1 fetch param").LogAndRespond(resp)
			return
		}
		if len(req.Form["sign"]) != 1 {
			util.NewHTTPError(http.StatusBadRequest, "Not exactly 1 sign param").LogAndRespond(resp)
			return
		}
		fetch = req.FormValue("fetch")
		sign = req.FormValue("sign")
	}
	fetchURL, signURL, errorOnStatefulHeaders, httpErr := parseURLs(fetch, sign, this.urlSets)
	if httpErr != nil {
		httpErr.LogAndRespond(resp)
		return
	}

	fetchReq, fetchResp, httpErr := this.fetchURL(fetchURL, req.Header)
	if httpErr != nil {
		httpErr.LogAndRespond(resp)
		return
	}

	defer func() {
		if err := fetchResp.Body.Close(); err != nil {
			log.Println("Error closing fetchResp body:", err)
		}
	}()

	if !this.shouldPackage() {
		log.Println("Not packaging because server is unhealthy; see above log statements.")
		proxy(resp, fetchResp)
		return
	}
	if this.requireHeaders {
		act := amp_cache_transform.ShouldSendSXG(req.Header.Get("AMP-Cache-Transform"))
		if act == "" {
			log.Println("Not packaging because AMP-Cache-Transform request header is missing.")
			proxy(resp, fetchResp)
			return
		}
		resp.Header().Set("AMP-Cache-Transform", act)
	}
	if this.requireHeaders && !accept.CanSatisfy(req.Header.Get("Accept")) {
		log.Printf("Not packaging because Accept request header lacks application/signed-exchange;v=%s.\n", accept.AcceptedSxgVersion)
		proxy(resp, fetchResp)
		return
	}

	switch fetchResp.StatusCode {
	case 200:
		// If fetchURL returns an OK status, then validate, munge, and package.
		if err := validateFetch(fetchReq, fetchResp); err != nil {
			log.Println("Not packaging because of invalid fetch: ", err)
			proxy(resp, fetchResp)
			return
		}
		// TODO(twifkak): Add config: either ensure Expires is + 5 days, or reject. (Or at least do one and document it in the README.)
		// TODO(twifkak): Should I be more restrictive and just whitelist some response headers?
		for header := range statefulResponseHeaders {
			if errorOnStatefulHeaders && fetchResp.Header.Get(header) != "" {
				log.Println("Not packaging because ErrorOnStatefulHeaders = True and fetch response contains stateful header: ", header)
				proxy(resp, fetchResp)
				return
			}
			fetchResp.Header.Del(header)
		}

		// charset=utf-8 would be redundant, as it is specified in the <meta> of a valid AMPHTML document:
		fetchResp.Header.Set("Content-Type", "text/html")
		fetchResp.Header.Set("Content-Security-Policy", contentSecurityPolicy)
		fetchResp.Header.Del("Link") // Ensure there are no privacy-violating Link:rel=preload headers.

		this.serveSignedExchange(resp, fetchResp, signURL)

	case 301, 302, 303:
		// If fetchURL returns a redirect, then forward that along; do not sign it and do not error out.
		resp.Header().Set("location", fetchResp.Header.Get("location"))
		resp.WriteHeader(fetchResp.StatusCode)
		if _, err := io.Copy(resp, fetchResp.Body); err != nil {
			log.Println("Error writing redirect body:", err)
		}

	case 304:
		// If fetchURL returns a 304, then also return a 304 with appropriate headers.
		for header := range statusNotModifiedHeaders {
			if fetchResp.Header.Get(header) != "" {
				resp.Header().Set(header, fetchResp.Header.Get(header))
			}
		}
		resp.WriteHeader(http.StatusNotModified)

	default:
		util.NewHTTPError(http.StatusBadGateway, "Non-OK fetch: ", fetchResp.StatusCode).LogAndRespond(resp)
	}
}

// serveSignedExchange does the actual work of transforming, packaging and signed and writing to the response.
func (this *Packager) serveSignedExchange(resp http.ResponseWriter, fetchResp *http.Response, signURL *url.URL) {
	// Override the content-type of the fetch response to ensure browsers
	// interpret the contents as HTML. The AMP Cache will validate that the
	// payload is valid AMPHTML. Alternatively, we could reject responses
	// with the wrong content-type, but:
	//  1. Some existing AMP servers may not be setting the proper content
	//     type, and may be relying on the AMP cache to rewrite it.
	//  2. This would require a media-type parser plus some logic for
	//     determining equivalence.
	fetchResp.Header.Set("Content-Type", "text/html")

	// TODO(twifkak): Are there any headers that AMP CDNs sets that publishers wouldn't want
	// running on their origin? Are there any (such as CSP) that we absolutely need to run?
	// TODO(twifkak): After the Transformer API, just add whatever headers are provided by the
	// transformer plus a few extra (e.g. Content-Type).

	fetchBody, err := ioutil.ReadAll(io.LimitReader(fetchResp.Body, maxBodyLength))
	if err != nil {
		util.NewHTTPError(http.StatusBadGateway, "Error reading body: ", err).LogAndRespond(resp)
		return
	}

	// Perform local transformations.
	r := getTransformerRequest(this.rtvCache, string(fetchBody), signURL.String())
	transformed, err := transformer.Process(r)
	if err != nil {
		util.NewHTTPError(http.StatusInternalServerError, "Error transforming: ", err).LogAndRespond(resp)
		return
	}
	fetchResp.Header.Set("Content-Length", strconv.Itoa(len(transformed)))

	exchange, err := signedexchange.NewExchange(signURL, http.Header{}, fetchResp.StatusCode, fetchResp.Header, []byte(transformed))
	if err != nil {
		util.NewHTTPError(http.StatusInternalServerError, "Error building exchange: ", err).LogAndRespond(resp)
		return
	}
	if err := exchange.MiEncodePayload(miRecordSize, version.Version1b2); err != nil {
		util.NewHTTPError(http.StatusInternalServerError, "Error MI-encoding: ", err).LogAndRespond(resp)
		return
	}
	certURL, err := this.genCertURL(this.cert, signURL)
	if err != nil {
		util.NewHTTPError(http.StatusInternalServerError, "Error building cert URL: ", err).LogAndRespond(resp)
		return
	}
	now := time.Now()
	validityHRef, err := url.Parse(util.ValidityMapPath)
	if err != nil {
		util.NewHTTPError(http.StatusInternalServerError, "Error building validity href: ", err).LogAndRespond(resp)
	}
	signer := signedexchange.Signer{
		// Expires - Date must be <= 604800 seconds, per
		// https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00#section-3.5.
		Date:        now.Add(-24 * time.Hour),
		Expires:     now.Add(6 * 24 * time.Hour),
		Certs:       []*x509.Certificate{this.cert},
		CertUrl:     certURL,
		ValidityUrl: signURL.ResolveReference(validityHRef),
		PrivKey:     this.key,
		// TODO(twifkak): Should we make Rand user-configurable? The
		// default is to use getrandom(2) if available, else
		// /dev/urandom.
	}
	if err := exchange.AddSignatureHeader(&signer, version.Version1b2); err != nil {
		util.NewHTTPError(http.StatusInternalServerError, "Error signing exchange: ", err).LogAndRespond(resp)
		return
	}
	// TODO(twifkak): Make this a streaming response. How will we handle errors after part of the response has already been sent?
	var body bytes.Buffer
	if err := exchange.Write(&body, version.Version1b2); err != nil {
		util.NewHTTPError(http.StatusInternalServerError, "Error serializing exchange: ", err).LogAndRespond(resp)
	}

	// TODO(twifkak): Add Cache-Control: public with expiry to match when we think the AMP Cache
	// should fetch an update (half-way between signature date & expires).
	// TODO(twifkak): Add `X-Amppkg-Version: 0.0.0`.
	resp.Header().Set("Content-Type", "application/signed-exchange;v=b2")
	resp.Header().Set("Cache-Control", "no-transform")
	if _, err := resp.Write(body.Bytes()); err != nil {
		log.Println("Error writing response:", err)
		return
	}
}

// Proxy the content unsigned.
func proxy(resp http.ResponseWriter, fetchResp *http.Response) {
	for k, v := range fetchResp.Header {
		resp.Header()[k] = v
	}
	bytesCopied, err := io.Copy(resp, fetchResp.Body)
	if err != nil {
		if bytesCopied == 0 {
			util.NewHTTPError(http.StatusInternalServerError, "Error copying response body").LogAndRespond(resp)
		} else {
			log.Printf("Error copying response body, %d bytes into stream\n", bytesCopied)
		}
	}
}
