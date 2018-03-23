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

	"github.com/nyaxt/webpackage/go/signedexchange"
	"github.com/pkg/errors"
	"github.com/pquerna/cachecontrol"
)

// Allowed schemes for the PackagerBase URL, from which certUrls are constructed.
var acceptablePackagerSchemes = map[string]bool{"http": true, "https": true}

// The user agent to send when issuing fetches. Should look like a mobile device.
const userAgent = "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) " +
	"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile " +
	"Safari/537.36 (compatible; amppackager/0.0.0; +https://github.com/ampproject/amppackager)"

// Advised against, per
// https://jyasskin.github.io/webpackage/implementation-draft/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#stateful-headers,
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
	return ret, nil
}

func regexpFullMatch(pattern string, test string) bool {
	// This is how regexp/exec_test.go turns a partial pattern into a full pattern.
	fullRe := `\A(?:` + pattern + `)\z`
	matches, _ := regexp.MatchString(fullRe, test)
	return matches
}

func urlMatches(url *url.URL, pattern URLPattern) bool {
	schemeMatches := false
	for _, scheme := range pattern.Scheme {
		if url.Scheme == scheme {
			schemeMatches = true
		}
	}
	if !schemeMatches {
		return false
	}
	if url.Opaque != "" {
		return false
	}
	if url.User != nil {
		return false
	}
	if url.Host != pattern.Domain {
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

func urlsMatch(fetchURL *url.URL, signURL *url.URL, set URLSet) bool {
	return urlMatches(fetchURL, set.Fetch) && urlMatches(signURL, set.Sign) &&
		(!set.SamePath || fetchURL.RequestURI() == signURL.RequestURI())
}

// Returns parsed sign URL and whether to fail on stateful headers.
func parseURLs(fetch string, sign string, urlSets []URLSet) (*url.URL, bool, *HTTPError) {
	fetchURL, err := parseURL(fetch, "fetch")
	if err != nil {
		return nil, false, err
	}
	signURL, err := parseURL(sign, "sign")
	if err != nil {
		return nil, false, err
	}
	for _, set := range urlSets {
		if urlsMatch(fetchURL, signURL, set) {
			return signURL, set.Fetch.ErrorOnStatefulHeaders, nil
		}
	}
	return nil, false, NewHTTPError(http.StatusBadRequest, "fetch/sign URLs do not match config")
}

func validateFetch(req *http.Request, resp *http.Response) *HTTPError {
	if resp.StatusCode != http.StatusOK {
		return NewHTTPError(http.StatusBadGateway, "Non-OK fetch: ", resp.StatusCode)
	}
	// Validate response is publicly-cacheable, per
	// https://tools.ietf.org/html/draft-yasskin-http-origin-signed-responses-03#section-6.1.
	nonCachableReasons, _, err := cachecontrol.CachableResponse(req, resp, cachecontrol.Options{PrivateCache: false})
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
	validityURL, err := url.Parse("https://cdn.ampproject.org/null-validity")
	if err != nil {
		return nil, errors.Wrap(err, "parsing null-validity URL")
	}
	client := http.Client{
		// TODO(twifkak): Load-test and see if default transport settings are okay.
		Timeout: 60 * time.Second,
	}
	return &Packager{cert, key, validityURL, &client, baseURL, urlSets}, nil
}

func (this Packager) fetchURL(fetch string) (*http.Request, *http.Response, *HTTPError) {
	log.Printf("Fetching URL: %q\n", fetch)
	// TODO(twifkak): Translate into AMP CDN URL, until transform API is available.
	req, err := http.NewRequest(http.MethodGet, fetch, nil)
	req.Header.Set("User-Agent", userAgent)
	// TODO(twifkak): Should we add 'Accept-Charset: utf-8'? The AMP Transformer API requires utf-8.
	if err != nil {
		return nil, nil, NewHTTPError(http.StatusInternalServerError, "Error building request: ", err)
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

func (this Packager) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// TODO(twifkak): See if there are any other validations or sanitizations that need adding.
	fetch := req.FormValue("fetch")
	sign := req.FormValue("sign")
	signURL, errorOnStatefulHeaders, httpErr := parseURLs(fetch, sign, this.urlSets)
	if httpErr != nil {
		httpErr.LogAndRespond(resp)
		return
	}

	fetchReq, fetchResp, httpErr := this.fetchURL(fetch)
	if httpErr != nil {
		httpErr.LogAndRespond(resp)
		return
	}
	defer func() {
		if err := fetchResp.Body.Close(); err != nil {
			log.Println("Error closing fetchResp body:", err)
		}
	}()

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
	fetchBody, err := ioutil.ReadAll(io.LimitReader(fetchResp.Body, maxBodyLength))
	if err != nil {
		log.Println("Error reading body:", err)
		http.Error(resp, "502 bad gateway", http.StatusBadGateway)
		return
	}
	exchange, err := signedexchange.NewExchange(signURL, http.Header{}, fetchResp.StatusCode, fetchResp.Header, fetchBody, miRecordSize)
	if err != nil {
		NewHTTPError(http.StatusInternalServerError, "Error building exchange: ", err).LogAndRespond(resp)
		return
	}
	certURL, err := this.genCertURL(this.cert)
	if err != nil {
		NewHTTPError(http.StatusInternalServerError, "Error building cert URL: ", err).LogAndRespond(resp)
		return
	}
	signer := signedexchange.Signer{
		// Expires - Date must be <= 604800 seconds, per
		// https://jyasskin.github.io/webpackage/implementation-draft/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#signature-validity.
		Date:        time.Now().Add(-24 * time.Hour),
		Expires:     time.Now().Add(6 * 24 * time.Hour),
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
	if err := signedexchange.WriteExchangeFile(&body, exchange); err != nil {
		NewHTTPError(http.StatusInternalServerError, "Error serializing exchange: ", err).LogAndRespond(resp)
	}

	// TODO(twifkak): Add Cache-Control: public with expiry to match the signature.
	// TODO(twifkak): Set some other headers?
	resp.Header().Set("Content-Type", "application/signed-exchange;v=b0")
	if _, err := resp.Write(body.Bytes()); err != nil {
		log.Println("Error writing response:", err)
		return
	}
}
