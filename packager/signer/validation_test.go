package signer

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ampproject/amppackager/packager/util"
	"github.com/stretchr/testify/assert"
)

func urlFrom(url *url.URL, err *util.HTTPError) *url.URL { return url }

func errorFrom(url *url.URL, err *util.HTTPError) *util.HTTPError { return err }

func urlOrDie(spec string) *url.URL {
	url, err := url.Parse(spec)
	if err != nil {
		panic(err)
	}
	return url
}

func TestParseURL(t *testing.T) {
	assert.EqualError(t, errorFrom(parseURL("", "sign")), "sign URL is unspecified")
	if err := errorFrom(parseURL("abc-@#79!%^/", "sign")); assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "Error parsing sign URL")
	}
	assert.EqualError(t, errorFrom(parseURL("abc/def", "sign")), "sign URL is relative")

	assert.Equal(t, "http://foo.com/baz", urlFrom(parseURL("http://foo.com/bar/../baz", "sign")).String())
}

func TestFetchURLMatches(t *testing.T) {
	assert.NoError(t, fetchURLMatches(nil, nil))
	assert.NoError(t, fetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}))
	assert.NoError(t, fetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, Domain: "example.com", PathRE: stringPtr("/"), QueryRE: stringPtr("")}))
	assert.NoError(t, fetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, DomainRE: "example.*", PathRE: stringPtr("/"), QueryRE: stringPtr("")}))

	assert.EqualError(t, fetchURLMatches(urlOrDie("http://example.com/"), nil),
		"If URLSet.Fetch is unspecified, then so should ?fetch= be.")
	assert.EqualError(t, fetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"https"}, PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}),
		"Scheme doesn't match")
	assert.EqualError(t, fetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, Domain: "wrongexample.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}),
		"Domain doesn't match")
	assert.EqualError(t, fetchURLMatches(urlOrDie("http://example.com:1234/"),
		&util.URLPattern{Scheme: []string{"http"}, Domain: "example.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}),
		"Domain doesn't match")
	assert.EqualError(t, fetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, DomainRE: "xample", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}),
		"DomainRE doesn't match")

	assert.EqualError(t, fetchURLMatches(urlOrDie("http:example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}),
		"URL is opaque")
	assert.EqualError(t, fetchURLMatches(urlOrDie("http://user@example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}),
		"URL contains user")
	assert.EqualError(t, fetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, PathRE: stringPtr("/amp/.*"), QueryRE: stringPtr(".*")}),
		"PathRE doesn't match")
	assert.EqualError(t, fetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, PathRE: stringPtr(".*"), PathExcludeRE: []string{"/"}, QueryRE: stringPtr(".*")}),
		"PathExcludeRE matches: /")
	assert.EqualError(t, fetchURLMatches(urlOrDie("http://example.com/?sessid=foo"),
		&util.URLPattern{Scheme: []string{"http"}, PathRE: stringPtr(".*"), QueryRE: stringPtr("")}),
		"QueryRE doesn't match")
}

func TestSignURLMatches(t *testing.T) {
	assert.NoError(t, signURLMatches(urlOrDie("https://example.com/"),
		&util.URLPattern{Domain: "example.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}))

	assert.EqualError(t, signURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Domain: "example.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}),
		"Scheme doesn't match")
	assert.EqualError(t, signURLMatches(urlOrDie("https://wrongexample.com/"),
		&util.URLPattern{Domain: "example.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}),
		"Domain doesn't match")
}

func TestURLsMatch(t *testing.T) {
	config := util.URLSet{
		Fetch: &util.URLPattern{
			Scheme: []string{"http"}, Domain: "fetch.com",
			PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"),
			SamePath: boolPtr(true)},
		Sign: &util.URLPattern{
			Domain: "sign.com",
			PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")},
	}

	assert.NoError(t, urlsMatch(urlOrDie("http://fetch.com/"), urlOrDie("https://sign.com/"), config))

	assert.EqualError(t, urlsMatch(urlOrDie("https://fetch.com/"), urlOrDie("https://sign.com/"), config),
		"fetch URL: Scheme doesn't match")
	assert.EqualError(t, urlsMatch(urlOrDie("http://fetch.com/"), urlOrDie("http://sign.com/"), config),
		"sign URL: Scheme doesn't match")
	assert.EqualError(t, urlsMatch(urlOrDie("http://fetch.com/"), urlOrDie("https://sign.com/other"), config),
		"fetch and sign paths don't match")

	*config.Fetch.SamePath = false
	assert.NoError(t, urlsMatch(urlOrDie("http://fetch.com/"), urlOrDie("https://sign.com/other"), config))
}

func TestParseURLs(t *testing.T) {
	if _, _, _, err := parseURLs("a%-", "b", []util.URLSet{}); assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "fetch URL")
	}
	if _, _, _, err := parseURLs("http://a", "b%-", []util.URLSet{}); assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "sign URL")
	}

	fetch, sign, errorOnStatefulHeaders, err := parseURLs("", "https://example.com/", []util.URLSet{
		{Sign: &util.URLPattern{Domain: "wrongexample.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}},
		{Sign: &util.URLPattern{Domain: "example.com", PathRE: stringPtr("/amp/.*"), QueryRE: stringPtr(".*")}},
		{Sign: &util.URLPattern{Domain: "example.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), ErrorOnStatefulHeaders: true}},
		{Sign: &util.URLPattern{Domain: "badexample.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, "https://example.com/", fetch.String())
		assert.Equal(t, "https://example.com/", sign.String())
		assert.True(t, errorOnStatefulHeaders)
	}

	_, _, _, err = parseURLs("", "https://example.com/", []util.URLSet{
		{Sign: &util.URLPattern{Domain: "wrongexample.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}},
		{Sign: &util.URLPattern{Domain: "example.com", PathRE: stringPtr("/amp/.*"), QueryRE: stringPtr(".*")}},
		{Sign: &util.URLPattern{Domain: "badexample.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*")}},
	})
	if assert.NotNil(t, err) {
		assert.EqualError(t, err, "fetch/sign URLs do not match config")
	}
}

func TestValidateFetch(t *testing.T) {
	req := httptest.NewRequest("", "/", nil)
	resp := http.Response{Header: http.Header{}}
	resp.Header.Set("Cache-Control", "max-age=ph'nglui mglw'nafh Cthulhu R'lyeh wgah'nagl fhtagn")
	if err := validateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Parsing cache headers")
	}

	resp.Header.Set("Cache-Control", "private")
	if err := validateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Non-cacheable response")
	}

	resp.Header.Del("Cache-Control")
	if err := validateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Non-cacheable response")
	}

	resp.Header.Set("Cache-Control", "public")

	resp.Header.Set("Content-Type", "text//html")
	if err := validateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Parsing Content-Type")
	}

	resp.Header.Set("Content-Type", "text/htmlol")
	if err := validateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Wrong Content-Type")
	}

	resp.Header.Set("Content-Type", "TEXT/HTML")
	assert.NoError(t, validateFetch(req, &resp))

	resp.Header.Set("Content-Type", "text/html;charset=ebcdic") // Close enough.
	assert.NoError(t, validateFetch(req, &resp))
}
