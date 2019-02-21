// Auxiliary functions for use by Signer.
package signer

import (
	"mime"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/ampproject/amppackager/packager/util"
	"github.com/pkg/errors"
	"github.com/pquerna/cachecontrol"
)

// Converts an URL string into an URL object with an unambiguous interpretation.
func parseURL(rawURL string, name string) (*url.URL, *util.HTTPError) {
	if rawURL == "" {
		return nil, util.NewHTTPError(http.StatusBadRequest, name, " URL is unspecified")
	}
	ret, err := url.Parse(rawURL)
	if err != nil {
		return nil, util.NewHTTPError(http.StatusBadRequest, "Error parsing ", name, " URL: ", err)
	}
	if !ret.IsAbs() {
		// Relative URLs don't make sense. For a fetch URL, it's not
		// clear what from what base URL the signer should resolve it
		// before fetching. For a sign URL, relative URLs are
		// disallowed by the SXG spec:
		// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#application-signed-exchange
		return nil, util.NewHTTPError(http.StatusBadRequest, name, " URL is relative")
	}
	// Evaluate "/..", by resolving the URL as a reference from itself.
	// This prevents malformed URLs from eluding the PathRE protections.
	ret = ret.ResolveReference(ret)
	// Escape special characters in the query component such as "<" or "|"
	// (but not "&" or "=").
	ret.RawQuery = url.PathEscape(ret.RawQuery)
	return ret, nil
}

// Returns true iff the given pattern matches the entire test string.
func regexpFullMatch(pattern string, test string) bool {
	// This is how regexp/exec_test.go turns a partial pattern into a full pattern.
	fullRe := `\A(?:` + pattern + `)\z`
	matches, _ := regexp.MatchString(fullRe, test)
	return matches
}

// Implements the URL-matching common to both fetchURLMatches and signURLMatches.
func urlMatches(url *url.URL, pattern util.URLPattern) error {
	if url.Opaque != "" {
		// Opaque URLs are unfetchable, and also disallowed by the spec
		// as sign URLs.
		return errors.New("URL is opaque")
	}
	if url.User != nil {
		// The `user:pass@` portion of a URL is not technically
		// disallowed by the spec, but is a weird enough request that
		// it seems wise to disable this capability by default (i.e.
		// more likely a sign of attack than a legitimate request).
		// Please open an issue if you have a legitimate need for this
		// in a fetch/sign URL.
		return errors.New("URL contains user")
	}
	// PathRE matches the path component of the URL, including the
	// beginning slash.
	if !regexpFullMatch(*pattern.PathRE, url.EscapedPath()) {
		return errors.New("PathRE doesn't match")
	}
	// If any of PathExcludeRE matches, the URL does not match.
	for _, re := range pattern.PathExcludeRE {
		if regexpFullMatch(re, url.EscapedPath()) {
			return errors.Errorf("PathExcludeRE matches: %s", re)
		}
	}
	// QueryRE matches the query component of the URL, *not* including the
	// beginning question mark.
	if !regexpFullMatch(*pattern.QueryRE, url.RawQuery) {
		return errors.New("QueryRE doesn't match")
	}
	if len(url.String()) > pattern.MaxLength {
		return errors.New("URL too long")
	}
	return nil
}

// True iff actualScheme is an element of expectedSchemes.
func schemeMatches(actualScheme string, expectedSchemes []string) bool {
	for _, expectedScheme := range expectedSchemes {
		if actualScheme == expectedScheme {
			return true
		}
	}
	return false
}

// True iff url matches pattern, as defined by an [URLSet.Fetch] block in the
// config file. The format of this URLPattern is validated by
// validateFetchURLPattern in config.go.
func fetchURLMatches(url *url.URL, pattern *util.URLPattern) error {
	// If the fetch block is not specified, then this particular URLSet is
	// a "sign-only" config. That is: only the sign URL should be passed to
	// the Signer; this will be used as the fetch URL as well.
	if pattern == nil {
		if url == nil {
			return nil
		} else {
			return errors.New("If URLSet.Fetch is unspecified, then so should ?fetch= be.")
		}
	}
	// The fetch block may specify which schemes are allowed.
	if !schemeMatches(url.Scheme, pattern.Scheme) {
		return errors.New("Scheme doesn't match")
	}
	// The fetch block may specify either Domain or DomainRE.
	if pattern.Domain != "" && url.Host != pattern.Domain {
		return errors.New("Domain doesn't match")
	}
	if pattern.DomainRE != "" && !regexpFullMatch(pattern.DomainRE, url.Host) {
		return errors.New("DomainRE doesn't match")
	}
	return urlMatches(url, *pattern)
}

// True iff url matches pattern, as defined by an [URLSet.Sign] block in the
// config file. The format of this URLPattern is validated by
// validateSignURLPattern in config.go.
func signURLMatches(url *url.URL, pattern *util.URLPattern) error {
	// The sign block may not specify which schemes are allowed. Only HTTPS
	// is allowed:
	// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#rfc.section.5.3
	if url.Scheme != "https" {
		return errors.New("Scheme doesn't match")
	}
	// The sign block may only specify Domain. DomainRE would only be
	// useful for wildcard SXG certificates. Please open an issue if you
	// have a valid wildcard SXG certificate and a legitimate need for
	// this. This should be implemented with some thought into how to
	// ensure that the sign URL matches the fetch URL.
	if url.Host != pattern.Domain {
		return errors.New("Domain doesn't match")
	}
	return urlMatches(url, *pattern)
}

// True iff the given fetchURL and signURL match the given set (as specified by
// an [[URLSet]] block in the config file), and, if SamePath is true (default),
// fetchURL and signURL match each other.
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

// If the given fetch and sign URLs are valid, and match at least one of the
// urlSets (as specified by the [[URLSet]] blocks in the config file), then
// this returns the parsed URLs as well as a bool containing the value of
// ErrorOnStatefulHeaders for the first matching URLSet. Otherwise, returns an
// error.
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

// Given a request/response pair for the fetch from the packager to the backend
// content server, validates that the response is fit for including in an AMP
// SXG.
func validateFetch(req *http.Request, resp *http.Response) error {
	// Validate response is publicly-cacheable, per
	// https://tools.ietf.org/html/draft-yasskin-http-origin-signed-responses-03#section-6.1, as referenced by
	// https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00#section-6.
	//
	// Note: If the cachecontrol library ever adds support for no-cache
	// with field name arguments, then instruct the signer to remove these
	// headers, per https://github.com/WICG/webpackage/pull/339.
	nonCachableReasons, _, err := cachecontrol.CachableResponse(req, resp, cachecontrol.Options{PrivateCache: false})
	if err != nil {
		return errors.Wrap(err, "Parsing cache headers")
	}
	if len(nonCachableReasons) > 0 {
		return errors.Errorf("Non-cacheable response: %s", nonCachableReasons)
	}

	// Validate that no Content-Encoding is specified. Otherwise, it was
	// encoded as something that http.Client was unable to decode (e.g. br).
	if encoding := resp.Header.Get("Content-Encoding"); encoding != "" {
		return errors.Errorf("Invalid Content-Encoding: %s", encoding)
	}

	// Validate that Content-Type seems right. This doesn't validate its
	// params (such as charset); we just want to verify we're not
	// misinterpreting the server's intent. We override the Content-Type
	// later for unambiguous interpretation by the browser.
	contentType, params, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return errors.Wrap(err, "Parsing Content-Type")
	}
	if contentType != "text/html" {
		return errors.Errorf("Wrong Content-Type: %s", contentType)
	}

	// Don't allow charset other than utf-8, as this overrides <meta charset>.
	charset := strings.ToLower(params["charset"])
	if charset != "" && charset != "utf-8" {
		return errors.Errorf("Wrong charset: %s", charset)
	}
	return nil
}
