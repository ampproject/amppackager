// Auxiliary functions for use by Signer.
package signer

import (
	"mime"
	"net/http"
	"net/url"
	"regexp"

	"github.com/ampproject/amppackager/packager/util"
	"github.com/pkg/errors"
	"github.com/pquerna/cachecontrol"
)

func parseURL(rawURL string, name string) (*url.URL, *util.HTTPError) {
	if rawURL == "" {
		return nil, util.NewHTTPError(http.StatusBadRequest, name, " URL is unspecified")
	}
	ret, err := url.Parse(rawURL)
	if err != nil {
		return nil, util.NewHTTPError(http.StatusBadRequest, "Error parsing ", name, " URL: ", err)
	}
	if !ret.IsAbs() {
		return nil, util.NewHTTPError(http.StatusBadRequest, name, " URL is relative")
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

func urlMatches(url *url.URL, pattern util.URLPattern) error {
	if url.Opaque != "" {
		return errors.New("URL is opaque")
	}
	if url.User != nil {
		return errors.New("URL contains user")
	}
	if !regexpFullMatch(*pattern.PathRE, url.EscapedPath()) {
		return errors.New("PathRE doesn't match")
	}
	for _, re := range pattern.PathExcludeRE {
		if regexpFullMatch(re, url.EscapedPath()) {
			return errors.Errorf("PathExcludeRE matches: %s", re)
		}
	}
	if !regexpFullMatch(*pattern.QueryRE, url.RawQuery) {
		return errors.New("QueryRE doesn't match")
	}
	return nil
}

func schemeMatches(actual string, expecteds []string) bool {
	for _, expected := range expecteds {
		if actual == expected {
			return true
		}
	}
	return false
}

func fetchURLMatches(url *url.URL, pattern *util.URLPattern) error {
	if pattern == nil {
		if url == nil {
			return nil
		} else {
			return errors.New("If URLSet.Fetch is unspecified, then so should ?fetch= be.")
		}
	}
	if !schemeMatches(url.Scheme, pattern.Scheme) {
		return errors.New("Scheme doesn't match")
	}
	if pattern.Domain != "" && url.Host != pattern.Domain {
		return errors.New("Domain doesn't match")
	}
	if pattern.DomainRE != "" && !regexpFullMatch(pattern.DomainRE, url.Host) {
		return errors.New("DomainRE doesn't match")
	}
	return urlMatches(url, *pattern)
}

func signURLMatches(url *url.URL, pattern *util.URLPattern) error {
	if url.Scheme != "https" {
		return errors.New("Scheme doesn't match")
	}
	if url.Host != pattern.Domain {
		return errors.New("Domain doesn't match")
	}
	return urlMatches(url, *pattern)
}

func urlsMatch(fetchURL *url.URL, signURL *url.URL, set util.URLSet) error {
	if err := fetchURLMatches(fetchURL, set.Fetch); err != nil {
		return errors.Wrap(err, "fetch URL")
	}
	if err := signURLMatches(signURL, set.Sign); err != nil {
		return errors.Wrap(err, "sign URL")
	}
	theyMatch := set.Fetch == nil || !*set.Fetch.SamePath || fetchURL.RequestURI() == signURL.RequestURI()
	if !theyMatch {
		return errors.New("fetch and sign paths don't match")
	}
	return nil
}

// Returns parsed URLs and whether to fail on stateful headers.
func parseURLs(fetch string, sign string, urlSets []util.URLSet) (*url.URL, *url.URL, bool, *util.HTTPError) {
	var fetchURL *url.URL
	var err *util.HTTPError
	if fetch != "" {
		fetchURL, err = parseURL(fetch, "fetch")
		if err != nil {
			// TODO(twifkak): Use errors.Wrap() after changing return types to error.
			return nil, nil, false, err
		}
	}
	signURL, err := parseURL(sign, "sign")
	if err != nil {
		// TODO(twifkak): Use errors.Wrap() after changing return types to error.
		return nil, nil, false, err
	}
	for _, set := range urlSets {
		err := urlsMatch(fetchURL, signURL, set)
		if err == nil {
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
		return errors.Wrap(err, "Parsing cache headers")
	}
	if len(nonCachableReasons) > 0 {
		return errors.Errorf("Non-cacheable response: %s", nonCachableReasons)
	}

	// Validate that Content-Type seems right. This doesn't validate its
	// params (such as charset); we just want to verify we're not
	// misinterpreting the server's intent. We override the Content-Type
	// later for unambiguous interpretation by the browser.
	contentType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return errors.Wrap(err, "Parsing Content-Type")
	}
	if contentType != "text/html" {
		return errors.Errorf("Wrong Content-Type: %s", contentType)
	}
	return nil
}
