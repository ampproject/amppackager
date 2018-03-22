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

// TODO(twifkak): Test this.
// TODO(twifkak): Document code.
package main

import (
	"bytes"
	"crypto"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/WICG/webpackage/go/signedexchange/certurl"
	"github.com/nyaxt/webpackage/go/signedexchange"
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
	"github.com/pquerna/cachecontrol"

	amppkg "github.com/ampproject/amppackager"
)

var flagConfig = flag.String("config", "amppkg.toml", "Path to the config toml file.")

// Allowed schemes for the PackagerBase URL, from which certUrls are constructed.
var acceptablePackagerSchemes = map[string]bool{"http": true, "https": true}

// Must start without a slash, for PackagerBase's sake.
const certUrlPrefix = "amppkg/cert"

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

// The basename for the given cert, as served by this packager's cert cache.
// Should be stable and unique (e.g. content-addressing). Clients should
// url.PathEscape this, just in case its format changes to need escaping in the
// future.
func certName(cert *x509.Certificate) string {
	sum := sha256.Sum256(cert.Raw)
	return base64.URLEncoding.EncodeToString(sum[:])
}

type CertCache struct {
	// TODO(twifkak): Support multiple certs.
	certName    string
	certMessage []byte
}

func newCertCache(cert *x509.Certificate, pemContent []byte) (*CertCache, error) {
	this := new(CertCache)
	this.certName = certName(cert)
	// TODO(twifkak): Refactor CertificateMessageFromPEM to be based on the x509.Certificate instead.
	var err error
	this.certMessage, err = certurl.CertificateMessageFromPEM(pemContent)
	if err != nil {
		return nil, errors.Wrap(err, "extracting certificate from CertFile")
	}
	return this, nil
}

func (this CertCache) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	println("path", req.URL.Path)
	if req.URL.Path == path.Join("/", certUrlPrefix, this.certName) {
		// https://jyasskin.github.io/webpackage/implementation-draft/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cert-chain-format
		resp.Header().Set("Content-Type", "application/tls-cert-chain")
		resp.Header().Set("ETag", "\""+this.certName+"\"")
		http.ServeContent(resp, req, "", time.Time{}, bytes.NewReader(this.certMessage))
	} else {
		http.NotFound(resp, req)
	}
}

func parseUrl(rawUrl string, name string) (*url.URL, *amppkg.HttpError) {
	if rawUrl == "" {
		return nil, amppkg.NewHttpError(http.StatusBadRequest, name, " URL is unspecified")
	}
	ret, err := url.Parse(rawUrl)
	if err != nil {
		return nil, amppkg.NewHttpError(http.StatusBadRequest, "Error parsing ", name, " url: ", err)
	}
	if !ret.IsAbs() {
		return nil, amppkg.NewHttpError(http.StatusBadRequest, name, " url is relative")
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

// Returns parsed sign URL and whether to fail on stateful headers.
func parseUrls(fetch string, sign string, urlSets []URLSet) (*url.URL, bool, *amppkg.HttpError) {
	fetchUrl, err := parseUrl(fetch, "fetch")
	if err != nil {
		return nil, false, err
	}
	signUrl, err := parseUrl(sign, "sign")
	if err != nil {
		return nil, false, err
	}
	for _, pattern := range urlSets {
		if urlMatches(fetchUrl, pattern.Fetch) && urlMatches(signUrl, pattern.Sign) {
			return signUrl, pattern.Fetch.ErrorOnStatefulHeaders, nil
		}
	}
	return nil, false, amppkg.NewHttpError(http.StatusBadRequest, "fetch/sign URLs do not match config")
}

func validateFetch(req *http.Request, resp *http.Response) *amppkg.HttpError {
	if resp.StatusCode != http.StatusOK {
		return amppkg.NewHttpError(http.StatusBadGateway, "Non-OK fetch: ", resp.StatusCode)
	}
	// Validate response is publicly-cacheable, per
	// https://tools.ietf.org/html/draft-yasskin-http-origin-signed-responses-03#section-6.1.
	nonCachableReasons, _, err := cachecontrol.CachableResponse(req, resp, cachecontrol.Options{PrivateCache: false})
	if err != nil {
		return amppkg.NewHttpError(http.StatusBadGateway, "Error parsing cache headers: ", err)
	}
	if len(nonCachableReasons) > 0 {
		return amppkg.NewHttpError(http.StatusBadGateway, "Non-cacheable response: ", nonCachableReasons)
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
	validityUrl *url.URL
	client      *http.Client
	baseUrl     *url.URL
	urlSets     []URLSet
}

func newPackager(cert *x509.Certificate, key crypto.PrivateKey, packagerBase string, urlSets []URLSet) (*Packager, error) {
	baseUrl, err := url.Parse(packagerBase)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing PackagerBase %q", packagerBase)
	}
	if !baseUrl.IsAbs() {
		return nil, errors.Errorf("PackagerBase %q must be an absolute URL.", packagerBase)
	}
	if !acceptablePackagerSchemes[baseUrl.Scheme] {
		return nil, errors.Errorf("PackagerBase %q must be over http or https.", packagerBase)
	}
	validityUrl, err := url.Parse("https://cdn.ampproject.org/null-validity")
	if err != nil {
		return nil, errors.Wrap(err, "parsing null-validity URL")
	}
	client := http.Client{
		// TODO(twifkak): Load-test and see if default transport settings are okay.
		Timeout: 60 * time.Second,
	}
	return &Packager{cert, key, validityUrl, &client, baseUrl, urlSets}, nil
}

func (this Packager) fetchUrl(fetch string) (*http.Request, *http.Response, *amppkg.HttpError) {
	log.Printf("Fetching URL: %q\n", fetch)
	// TODO(twifkak): Translate into AMP CDN URL, until transform API is available.
	req, err := http.NewRequest(http.MethodGet, fetch, nil)
	req.Header.Set("User-Agent", "amppackager-0.0.0")
	// TODO(twifkak): Should we add 'Accept-Charset: utf-8'? The AMP Transformer API requires utf-8.
	if err != nil {
		return nil, nil, amppkg.NewHttpError(http.StatusInternalServerError, "Error building request: ", err)
	}
	resp, err := this.client.Do(req)
	if err != nil {
		return nil, nil, amppkg.NewHttpError(http.StatusBadGateway, "Error fetching: ", err)
	}
	return req, resp, nil
}

func (this Packager) genCertUrl(cert *x509.Certificate) (*url.URL, error) {
	urlPath := path.Join(certUrlPrefix, url.PathEscape(certName(cert)))
	certUrl, err := url.Parse(urlPath)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing certUrl %q", urlPath)
	}
	ret := this.baseUrl.ResolveReference(certUrl)
	return ret, nil
}

func (this Packager) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// TODO(twifkak): See if there are any other validations or sanitizations that need adding.
	fetch := req.FormValue("fetch")
	sign := req.FormValue("sign")
	signUrl, errorOnStatefulHeaders, httpErr := parseUrls(fetch, sign, this.urlSets)
	if httpErr != nil {
		httpErr.LogAndRespond(resp)
		return
	}

	fetchReq, fetchResp, httpErr := this.fetchUrl(fetch)
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
	for header, _ := range statefulResponseHeaders {
		if errorOnStatefulHeaders && fetchResp.Header.Get(header) != "" {
			amppkg.NewHttpError(http.StatusBadGateway, "Fetch response contains stateful header: ", header).LogAndRespond(resp)
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
	exchange, err := signedexchange.NewExchange(signUrl, http.Header{}, fetchResp.StatusCode, fetchResp.Header, fetchBody, miRecordSize)
	if err != nil {
		amppkg.NewHttpError(http.StatusInternalServerError, "Error building exchange: ", err).LogAndRespond(resp)
		return
	}
	certUrl, err := this.genCertUrl(this.cert)
	if err != nil {
		amppkg.NewHttpError(http.StatusInternalServerError, "Error building cert URL: ", err).LogAndRespond(resp)
		return
	}
	signer := signedexchange.Signer{
		// Expires - Date must be <= 604800 seconds, per
		// https://jyasskin.github.io/webpackage/implementation-draft/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#signature-validity.
		Date:        time.Now().Add(-24 * time.Hour),
		Expires:     time.Now().Add(6 * 24 * time.Hour),
		Certs:       []*x509.Certificate{this.cert},
		CertUrl:     certUrl,
		ValidityUrl: this.validityUrl,
		PrivKey:     this.key,
		// TODO(twifkak): Should we make Rand user-configurable? The
		// default is to use getrandom(2) if available, else
		// /dev/urandom.
	}
	if err := exchange.AddSignatureHeader(&signer); err != nil {
		amppkg.NewHttpError(http.StatusInternalServerError, "Error signing exchange: ", err).LogAndRespond(resp)
		return
	}
	// TODO(twifkak): Make this a streaming response. How will we handle errors after part of the response has already been sent?
	var body bytes.Buffer
	if err := signedexchange.WriteExchangeFile(&body, exchange); err != nil {
		amppkg.NewHttpError(http.StatusInternalServerError, "Error serializing exchange: ", err).LogAndRespond(resp)
	}

	// TODO(twifkak): Add Cache-Control: public with expiry to match the signature.
	// TODO(twifkak): Set some other headers?
	resp.Header().Set("Content-Type", "application/signed-exchange;v=b0")
	if _, err := resp.Write(body.Bytes()); err != nil {
		log.Println("Error writing response:", err)
		return
	}
}

type Config struct {
	LocalOnly    bool
	Port         int
	PackagerBase string // The base URL under which /amppkg/ URLs will be served on the internet.
	CertFile     string // This must be the full certificate chain.
	KeyFile      string // Just for the first cert, obviously.
	GoogleAPIKey string
	URLSet       []URLSet
}

type URLSet struct {
	Fetch URLPattern
	Sign  URLPattern
}

type URLPattern struct {
	Scheme                 []string
	Domain                 string
	PathRE                 *string
	PathExcludeRE          []string
	QueryRE                *string
	ErrorOnStatefulHeaders bool
}

var dotStarRegexp = ".*"

// Also sets defaults.
func validateURLPattern(pattern *URLPattern, allowedSchemes map[string]bool) error {
	if len(pattern.Scheme) == 0 {
		// Default Scheme to the list of keys in allowedSchemes.
		pattern.Scheme = make([]string, len(allowedSchemes))
		i := 0
		for scheme := range allowedSchemes {
			pattern.Scheme[i] = scheme
			i++
		}
	} else {
		for _, scheme := range pattern.Scheme {
			if !allowedSchemes[scheme] {
				return errors.Errorf("Scheme contains invalid value %q", scheme)
			}
		}
	}
	if pattern.Domain == "" {
		return errors.New("Domain must be specified")
	}
	if pattern.PathRE == nil {
		pattern.PathRE = &dotStarRegexp
	} else if _, err := regexp.Compile(*pattern.PathRE); err != nil {
		return errors.New("PathRE must be a valid regexp")
	}
	for _, exclude := range pattern.PathExcludeRE {
		if _, err := regexp.Compile(exclude); err != nil {
			return errors.Errorf("PathExcludeRE contains be invalid regexp %q", exclude)
		}
	}
	if pattern.QueryRE == nil {
		pattern.QueryRE = &dotStarRegexp
	} else if _, err := regexp.Compile(*pattern.QueryRE); err != nil {
		return errors.New("QueryRE must be a valid regexp")
	}
	return nil
}

var allowedFetchSchemes = map[string]bool{"http": true, "https": true}
var allowedSignSchemes = map[string]bool{"https": true}

// Reads the config file specified at --config and validates it.
func readConfig() (*Config, error) {
	if *flagConfig == "" {
		return nil, errors.New("must specify --config")
	}
	tree, err := toml.LoadFile(*flagConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse config at: %s", *flagConfig)
	}
	config := Config{}
	if err = tree.Unmarshal(&config); err != nil {
		return nil, errors.Wrapf(err, "failed to parse config at: %s", *flagConfig)
	}
	// TODO(twifkak): Return an error if the TOML includes any fields that aren't part of the Config struct.

	if config.Port == 0 {
		config.Port = 8080
	}
	if !strings.HasSuffix(config.PackagerBase, "/") {
		// This ensures that the ResolveReference call doesn't replace the last path component.
		config.PackagerBase += "/"
	}
	if config.CertFile == "" {
		return nil, errors.New("must specify CertFile")
	}
	if config.KeyFile == "" {
		return nil, errors.New("must specify KeyFile")
	}
	if config.GoogleAPIKey == "" {
		return nil, errors.New("must specify GoogleAPIKey")
	}
	if len(config.URLSet) == 0 {
		return nil, errors.New("must specify one or more [[URLSet]]")
	}
	for i := range config.URLSet {
		if err := validateURLPattern(&config.URLSet[i].Fetch, allowedFetchSchemes); err != nil {
			return nil, errors.Wrapf(err, "parsing URLSet.%d.Fetch", i)
		}
		if err := validateURLPattern(&config.URLSet[i].Sign, allowedSignSchemes); err != nil {
			return nil, errors.Wrapf(err, "parsing URLSet.%d.Sign", i)
		}
		if config.URLSet[i].Sign.ErrorOnStatefulHeaders {
			return nil, errors.Errorf("URLSet.%d.Sign.ErrorOnStatefulHeaders is not allowed; perhaps you meant to put this in the Fetch section?", i)
		}
	}
	return &config, nil
}

type LogIntercept struct {
	handler http.Handler
}

func (this LogIntercept) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// TODO(twifkak): Adopt whatever the standard format is nowadays.
	log.Println("Serving", req.URL, "to", req.RemoteAddr)
	this.handler.ServeHTTP(resp, req)
	// TODO(twifkak): Get status code from resp. This requires making a ResponseWriter wrapper.
	// TODO(twifkak): Separate the typical weblog from the detailed error log.
}

// Exposes an HTTP server. Don't run this on the open internet, for at least two reasons:
//  - It exposes an API that allows people to sign any URL as any other URL.
//  - It is in cleartext.
func main() {
	flag.Parse()
	config, err := readConfig()
	if err != nil {
		panic(errors.Wrap(err, "reading config"))
	}

	// TODO(twifkak): Document what cert/key storage formats this accepts.
	certPem, err := ioutil.ReadFile(config.CertFile)
	if err != nil {
		panic(errors.Wrapf(err, "reading %s", config.CertFile))
	}
	keyPem, err := ioutil.ReadFile(config.KeyFile)
	if err != nil {
		panic(errors.Wrapf(err, "reading %s", config.KeyFile))
	}

	certs, err := signedexchange.ParseCertificates(certPem)
	if err != nil {
		panic(errors.Wrapf(err, "parsing %s", config.CertFile))
	}
	if certs == nil || len(certs) == 0 {
		panic(fmt.Sprintf("no cert found in %s", config.CertFile))
	}
	cert := certs[0]
	// TODO(twifkak): Verify that cert covers all the signing domains in the config.

	keyBlock, _ := pem.Decode(keyPem)
	if keyBlock == nil {
		panic(fmt.Sprintf("no key found in %s", config.KeyFile))
	}

	key, err := signedexchange.ParsePrivateKey(keyBlock.Bytes)
	if err != nil {
		panic(errors.Wrapf(err, "parsing %s", config.KeyFile))
	}
	// TODO(twifkak): Verify that key matches cert.

	packager, err := newPackager(cert, key, config.PackagerBase, config.URLSet)
	if err != nil {
		panic(errors.Wrap(err, "building packager"))
	}
	certCache, err := newCertCache(cert, certPem)
	if err != nil {
		panic(errors.Wrap(err, "building cert cache"))
	}

	// TODO(twifkak): Make log output configurable.
	mux := http.NewServeMux()
	mux.Handle("/priv-amppkg/doc", packager)
	mux.Handle(path.Join("/", certUrlPrefix)+"/", certCache)
	addr := ""
	if config.LocalOnly {
		addr = "localhost"
	}
	addr += fmt.Sprint(":", config.Port)
	server := http.Server{
		Addr: addr,
		// Don't use DefaultServeMux, per
		// https://blog.cloudflare.com/exposing-go-on-the-internet/.
		Handler:           LogIntercept{mux},
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		// If needing to stream the response, disable WriteTimeout and
		// use TimeoutHandler instead, per
		// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/.
		WriteTimeout: 60 * time.Second,
		// Needs Go 1.8.
		IdleTimeout: 120 * time.Second,
		// TODO(twifkak): Specify ErrorLog?
	}

	// TODO(twifkak): Add monitoring (e.g. per the above Cloudflare blog).

	log.Println("Serving on port", config.Port)

	// TCP keep-alive timeout on ListenAndServe is 3 minutes. To shorten,
	// follow the above Cloudflare blog.
	log.Fatal(server.ListenAndServe())
}
