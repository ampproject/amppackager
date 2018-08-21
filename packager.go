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

package amppackager

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
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/pkg/errors"
	"github.com/pquerna/cachecontrol"
)

// The base URL for transformed fetch URLs.
var AmpCDNBase = "https://cdn.ampproject.org/c/"

// Allowed schemes for the PackagerBase URL, from which certUrls are constructed.
var acceptablePackagerSchemes = map[string]bool{"http": true, "https": true}

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

func parseURL(rawURL string, name string) (*url.URL, *HTTPError) {
	if rawURL == "" {
		return nil, NewHTTPError(http.StatusBadRequest, name, " URL is unspecified")
	}
	ret, err := url.Parse(rawURL)
	if err != nil {
		return nil, NewHTTPError(http.StatusBadRequest, "Error parsing ", name, " url: ", err)
	}
	if !ret.IsAbs() {
		return nil, NewHTTPError(http.StatusBadRequest, name, " url is relative")
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

func urlMatches(url *url.URL, pattern URLPattern) bool {
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

func fetchUrlMatches(url *url.URL, pattern *URLPattern) bool {
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

func signUrlMatches(url *url.URL, pattern *URLPattern) bool {
	if url.Scheme != "https" {
		return false
	}
	if url.Host != pattern.Domain {
		return false
	}
	return urlMatches(url, *pattern)
}

func urlsMatch(fetchURL *url.URL, signURL *url.URL, set URLSet) bool {
	return fetchUrlMatches(fetchURL, set.Fetch) && signUrlMatches(signURL, set.Sign) &&
		(set.Fetch == nil || !*set.Fetch.SamePath || fetchURL.RequestURI() == signURL.RequestURI())
}

// Returns parsed URLs and whether to fail on stateful headers.
func parseURLs(fetch string, sign string, urlSets []URLSet) (*url.URL, *url.URL, bool, *HTTPError) {
	var fetchURL *url.URL
	var err *HTTPError
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
	return nil, nil, false, NewHTTPError(http.StatusBadRequest, "fetch/sign URLs do not match config")
}

func validateFetch(req *http.Request, resp *http.Response) *HTTPError {
	if resp.StatusCode != http.StatusOK {
		return NewHTTPError(http.StatusBadGateway, "Non-OK fetch: ", resp.StatusCode)
	}
	// Validate response is publicly-cacheable, per
	// https://tools.ietf.org/html/draft-yasskin-http-origin-signed-responses-03#section-6.1, as referenced by
	// https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00#section-6.
	// TODO(twifkak): Set {PrivateCache: false} after we switch from
	// fetching through the AMP CDN to fetching directly and using the
	// transformer API. For now, the AMP CDN validates that the origin
	// response is publicly-cacheable.
	nonCachableReasons, _, err := cachecontrol.CachableResponse(req, resp, cachecontrol.Options{PrivateCache: true})
	if err != nil {
		return NewHTTPError(http.StatusBadGateway, "Error parsing cache headers: ", err)
	}
	if len(nonCachableReasons) > 0 {
		return NewHTTPError(http.StatusBadGateway, "Non-cacheable response: ", nonCachableReasons)
	}
	return nil
}

type Packager struct {
	// TODO(twifkak): Support multiple certs. This will require generating
	// a signature for each one. Note that Chrome only supports 1 signature
	// at the moment.
	cert *x509.Certificate
	// TODO(twifkak): Do we want to allow multiple keys?
	key         crypto.PrivateKey
	validityURL *url.URL
	client      *http.Client
	baseURL     *url.URL
	urlSets     []URLSet
}

func noRedirects(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func NewPackager(cert *x509.Certificate, key crypto.PrivateKey, packagerBase string, urlSets []URLSet) (*Packager, error) {
	baseURL, err := url.Parse(packagerBase)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing PackagerBase %q", packagerBase)
	}
	if !baseURL.IsAbs() {
		return nil, errors.Errorf("PackagerBase %q must be an absolute URL.", packagerBase)
	}
	if !acceptablePackagerSchemes[baseURL.Scheme] {
		return nil, errors.Errorf("PackagerBase %q must be over http or https.", packagerBase)
	}
	// packagerBase is always guaranteed to have a trailing slash due to config.go
	validityURL, err := url.Parse(packagerBase + ValidityMapURL)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing PackagerBase %q with ValidityMapURL %q", packagerBase, ValidityMapURL)
	}
	client := http.Client{
		CheckRedirect: noRedirects,
		// TODO(twifkak): Load-test and see if default transport settings are okay.
		Timeout: 60 * time.Second,
	}
	return &Packager{cert, key, validityURL, &client, baseURL, urlSets}, nil
}

func (this Packager) fetchURL(orig *url.URL, serveHTTPReq http.Header) (*http.Request, *http.Response, *HTTPError) {
	// Make a copy so destructive changes don't persist.
	fetch := *orig
	// Add the query parameter to enable web package transforms.
	query := fetch.Query()
	query.Add("usqp", "mq331AQCSAE")
	fetch.RawQuery = query.Encode()

	ampURL := AmpCDNBase
	if fetch.Scheme == "https" {
		ampURL += "s/"
	}
	ampURL += fetch.Host + fetch.RequestURI()

	log.Printf("Fetching URL: %q\n", ampURL)
	// TODO(twifkak): Translate into AMP CDN URL, until transform API is available.
	req, err := http.NewRequest(http.MethodGet, ampURL, nil)
	// TODO(twifkak): Should we add 'Accept-Charset: utf-8'? The AMP Transformer API requires utf-8.
	if err != nil {
		return nil, nil, NewHTTPError(http.StatusInternalServerError, "Error building request: ", err)
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
		return nil, nil, NewHTTPError(http.StatusBadGateway, "Error fetching: ", err)
	}
	return req, resp, nil
}

func (this Packager) genCertURL(cert *x509.Certificate) (*url.URL, error) {
	urlPath := path.Join(CertURLPrefix, url.PathEscape(CertName(cert)))
	certURL, err := url.Parse(urlPath)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing cert URL %q", urlPath)
	}
	ret := this.baseURL.ResolveReference(certURL)
	return ret, nil
}

func (this Packager) ServeHTTP(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// TODO(twifkak): See if there are any other validations or sanitizations that need adding.
	if err := req.ParseForm(); err != nil {
		NewHTTPError(http.StatusBadRequest, "Form input parsing failed: ", err).LogAndRespond(resp)
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
			NewHTTPError(http.StatusBadRequest, "More than 1 fetch param").LogAndRespond(resp)
			return
		}
		if len(req.Form["sign"]) != 1 {
			NewHTTPError(http.StatusBadRequest, "Not exactly 1 sign param").LogAndRespond(resp)
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

	// If fetchURL returns a redirect, then forward that along; do not sign it and do not error out.
	if fetchResp.StatusCode == 301 || fetchResp.StatusCode == 302 || fetchResp.StatusCode == 303 {
		resp.Header().Set("location", fetchResp.Header.Get("location"))
		resp.WriteHeader(fetchResp.StatusCode)
		_, err := io.Copy(resp, fetchResp.Body)
		if err != nil {
			log.Println("Error writing redirect body:", err)
		}
		return
	}

	// If fetchURL returns a 304, then also return a 304 with appropriate headers.
	if fetchResp.StatusCode == 304 {
		for header := range statusNotModifiedHeaders {
			if fetchResp.Header.Get(header) != "" {
				resp.Header().Set(header, fetchResp.Header.Get(header))
			}
		}
		resp.WriteHeader(http.StatusNotModified)
		return
	}

	if httpErr := validateFetch(fetchReq, fetchResp); httpErr != nil {
		httpErr.LogAndRespond(resp)
		return
	}

	// TODO(twifkak): Add config: either ensure Expires is + 5 days, or reject. (Or at least do one and document it in the README.)
	// TODO(twifkak): Should I be more restrictive and just whitelist some response headers?
	for header := range statefulResponseHeaders {
		if errorOnStatefulHeaders && fetchResp.Header.Get(header) != "" {
			NewHTTPError(http.StatusBadGateway, "Fetch response contains stateful header: ", header).LogAndRespond(resp)
			return
		}
		fetchResp.Header.Del(header)
	}
	// TODO(twifkak): Are there any headers that AMP CDNs sets that publishers wouldn't want
	// running on their origin? Are there any (such as CSP) that we absolutely need to run?
	// TODO(twifkak): After the Transformer API, just add whatever headers are provided by the
	// transformer plus a few extra (e.g. Content-Type).

	fetchBody, err := ioutil.ReadAll(io.LimitReader(fetchResp.Body, maxBodyLength))
	if err != nil {
		NewHTTPError(http.StatusBadGateway, "Error reading body: ", err).LogAndRespond(resp)
		return
	}

	exchange, err := signedexchange.NewExchange(signURL, http.Header{}, fetchResp.StatusCode, fetchResp.Header, fetchBody)
	if err != nil {
		NewHTTPError(http.StatusInternalServerError, "Error building exchange: ", err).LogAndRespond(resp)
		return
	}
	if err := exchange.MiEncodePayload(miRecordSize); err != nil {
		NewHTTPError(http.StatusInternalServerError, "Error MI-encoding: ", err).LogAndRespond(resp)
		return
	}
	certURL, err := this.genCertURL(this.cert)
	if err != nil {
		NewHTTPError(http.StatusInternalServerError, "Error building cert URL: ", err).LogAndRespond(resp)
		return
	}
	now := time.Now()
	signer := signedexchange.Signer{
		// Expires - Date must be <= 604800 seconds, per
		// https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00#section-3.5.
		Date:        now.Add(-24 * time.Hour),
		Expires:     now.Add(6 * 24 * time.Hour),
		Certs:       []*x509.Certificate{this.cert},
		CertUrl:     certURL,
		ValidityUrl: this.validityURL,
		PrivKey:     this.key,
		// TODO(twifkak): Should we make Rand user-configurable? The
		// default is to use getrandom(2) if available, else
		// /dev/urandom.
	}
	if err := exchange.AddSignatureHeader(&signer); err != nil {
		NewHTTPError(http.StatusInternalServerError, "Error signing exchange: ", err).LogAndRespond(resp)
		return
	}
	// TODO(twifkak): Make this a streaming response. How will we handle errors after part of the response has already been sent?
	var body bytes.Buffer
	if err := exchange.Write(&body); err != nil {
		NewHTTPError(http.StatusInternalServerError, "Error serializing exchange: ", err).LogAndRespond(resp)
	}

	// TODO(twifkak): Add Cache-Control: public with expiry to match when we think the AMP Cache
	// should fetch an update (half-way between signature date & expires).
	// TODO(twifkak): Add `X-Amppkg-Version: 0.0.0`.
	resp.Header().Set("Content-Type", "application/signed-exchange;v=b0")
	resp.Header().Set("Cache-Control", "no-transform")
	if _, err := resp.Write(body.Bytes()); err != nil {
		log.Println("Error writing response:", err)
		return
	}
}
