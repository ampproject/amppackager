package signer_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ampproject/amppackager/packager/signer"
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
	assert.EqualError(t, errorFrom(signer.ParseURL("", "sign")), "sign URL is unspecified")
	if err := errorFrom(signer.ParseURL("abc-@#79!%^/", "sign")); assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "Error parsing sign URL")
	}
	assert.EqualError(t, errorFrom(signer.ParseURL("abc/def", "sign")), "sign URL is relative")

	assert.Equal(t, "http://foo.com/baz", urlFrom(signer.ParseURL("http://foo.com/bar/../baz", "sign")).String())
}

func TestFetchURLMatches(t *testing.T) {
	assert.NoError(t, signer.FetchURLMatches(nil, nil))
	assert.NoError(t, signer.FetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}))
	assert.NoError(t, signer.FetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, Domain: "example.com", PathRE: stringPtr("/"), QueryRE: stringPtr(""), MaxLength: 2000}))
	assert.NoError(t, signer.FetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, DomainRE: "example.*", PathRE: stringPtr("/"), QueryRE: stringPtr(""), MaxLength: 2000}))

	assert.EqualError(t, signer.FetchURLMatches(urlOrDie("http://example.com/"), nil),
		"If URLSet.Fetch is unspecified, then so should ?fetch= be.")
	assert.EqualError(t, signer.FetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"https"}, PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}),
		"Scheme doesn't match")
	assert.EqualError(t, signer.FetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, Domain: "wrongexample.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}),
		"Domain doesn't match")
	assert.EqualError(t, signer.FetchURLMatches(urlOrDie("http://example.com:1234/"),
		&util.URLPattern{Scheme: []string{"http"}, Domain: "example.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}),
		"Domain doesn't match")
	assert.EqualError(t, signer.FetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, DomainRE: "xample", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}),
		"DomainRE doesn't match")

	assert.EqualError(t, signer.FetchURLMatches(urlOrDie("http:example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}),
		"URL is opaque")
	assert.EqualError(t, signer.FetchURLMatches(urlOrDie("http://user@example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}),
		"URL contains user")
	assert.EqualError(t, signer.FetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, PathRE: stringPtr("/amp/.*"), QueryRE: stringPtr(".*"), MaxLength: 2000}),
		"PathRE doesn't match")
	assert.EqualError(t, signer.FetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, PathRE: stringPtr(".*"), PathExcludeRE: []string{"/"}, QueryRE: stringPtr(".*"), MaxLength: 2000}),
		"PathExcludeRE matches: /")
	assert.EqualError(t, signer.FetchURLMatches(urlOrDie("http://example.com/?sessid=foo"),
		&util.URLPattern{Scheme: []string{"http"}, PathRE: stringPtr(".*"), QueryRE: stringPtr(""), MaxLength: 2000}),
		"QueryRE doesn't match")
	assert.EqualError(t, signer.FetchURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Scheme: []string{"http"}, PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 10}),
		"URL too long")
}

func TestSignURLMatches(t *testing.T) {
	assert.NoError(t, signer.SignURLMatches(urlOrDie("https://example.com/"),
		&util.URLPattern{Domain: "example.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}))

	assert.EqualError(t, signer.SignURLMatches(urlOrDie("http://example.com/"),
		&util.URLPattern{Domain: "example.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}),
		"Scheme doesn't match")
	assert.EqualError(t, signer.SignURLMatches(urlOrDie("https://wrongexample.com/"),
		&util.URLPattern{Domain: "example.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}),
		"Domain doesn't match")
}

func TestURLsMatch(t *testing.T) {
	config := util.URLSet{
		Fetch: &util.URLPattern{
			Scheme: []string{"http"}, Domain: "fetch.com",
			PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000,
			SamePath: boolPtr(true)},
		Sign: &util.URLPattern{
			Domain: "sign.com",
			PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000},
	}

	assert.NoError(t, signer.URLsMatch(urlOrDie("http://fetch.com/"), urlOrDie("https://sign.com/"), config))

	assert.EqualError(t, signer.URLsMatch(urlOrDie("https://fetch.com/"), urlOrDie("https://sign.com/"), config),
		"fetch URL: Scheme doesn't match")
	assert.EqualError(t, signer.URLsMatch(urlOrDie("http://fetch.com/"), urlOrDie("http://sign.com/"), config),
		"sign URL: Scheme doesn't match")
	assert.EqualError(t, signer.URLsMatch(urlOrDie("http://fetch.com/"), urlOrDie("https://sign.com/other"), config),
		"fetch and sign paths don't match")

	*config.Fetch.SamePath = false
	assert.NoError(t, signer.URLsMatch(urlOrDie("http://fetch.com/"), urlOrDie("https://sign.com/other"), config))
}

func TestParseURLs(t *testing.T) {
	if _, _, _, err := signer.ParseURLs("a%-", "b", []util.URLSet{}); assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "fetch URL")
	}
	if _, _, _, err := signer.ParseURLs("http://a", "b%-", []util.URLSet{}); assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "sign URL")
	}

	fetch, sign, errorOnStatefulHeaders, err := signer.ParseURLs("", "https://example.com/", []util.URLSet{
		{Sign: &util.URLPattern{Domain: "wrongexample.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}},
		{Sign: &util.URLPattern{Domain: "example.com", PathRE: stringPtr("/amp/.*"), QueryRE: stringPtr(".*"), MaxLength: 2000}},
		{Sign: &util.URLPattern{Domain: "example.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000, ErrorOnStatefulHeaders: true}},
		{Sign: &util.URLPattern{Domain: "badexample.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, "https://example.com/", fetch.String())
		assert.Equal(t, "https://example.com/", sign.String())
		assert.True(t, errorOnStatefulHeaders)
	}

	_, _, _, err = signer.ParseURLs("", "https://example.com/", []util.URLSet{
		{Sign: &util.URLPattern{Domain: "wrongexample.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}},
		{Sign: &util.URLPattern{Domain: "example.com", PathRE: stringPtr("/amp/.*"), QueryRE: stringPtr(".*"), MaxLength: 2000}},
		{Sign: &util.URLPattern{Domain: "badexample.com", PathRE: stringPtr(".*"), QueryRE: stringPtr(".*"), MaxLength: 2000}},
	})
	if assert.NotNil(t, err) {
		assert.EqualError(t, err, "fetch/sign URLs do not match config")
	}
}

func TestValidateFetch(t *testing.T) {
	req := httptest.NewRequest("", "/", nil)
	resp := http.Response{Header: http.Header{}}
	resp.Header.Set("Cache-Control", "max-age=ph'nglui mglw'nafh Cthulhu R'lyeh wgah'nagl fhtagn")
	if err := signer.ValidateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Parsing cache headers")
	}

	resp.Header.Set("Cache-Control", "private")
	if err := signer.ValidateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Non-cacheable response")
	}

	resp.Header.Del("Cache-Control")
	if err := signer.ValidateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Non-cacheable response")
	}

	resp.Header.Set("Cache-Control", "public")

	resp.Header.Set("Content-Type", "text//html")
	if err := signer.ValidateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Parsing Content-Type")
	}

	resp.Header.Set("Content-Type", "text/html;charset=utf-8;charset=ebcdic")
	if err := signer.ValidateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Parsing Content-Type")
	}

	resp.Header.Set("Content-Type", "text/htmlol")
	if err := signer.ValidateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Wrong Content-Type")
	}

	resp.Header.Set("Content-Type", "text/html;charset=ebcdic")
	if err := signer.ValidateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Wrong charset")
	}

	resp.Header.Set("Content-Type", "text/html;CHARSET=ebcdic")
	if err := signer.ValidateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Wrong charset")
	}

	resp.Header.Set("Content-Type", `text/html; charset ="ebcdic"`)
	if err := signer.ValidateFetch(req, &resp); assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Wrong charset")
	}

	resp.Header.Set("Content-Type", "text/html")
	assert.NoError(t, signer.ValidateFetch(req, &resp))

	// Examples from https://tools.ietf.org/html/rfc7231#section-3.1.1.1:

	resp.Header.Set("Content-Type", "text/html;charset=utf-8")
	assert.NoError(t, signer.ValidateFetch(req, &resp))

	resp.Header.Set("Content-Type", "text/html;charset=UTF-8")
	assert.NoError(t, signer.ValidateFetch(req, &resp))

	resp.Header.Set("Content-Type", `Text/HTML;Charset="utf-8"`)
	assert.NoError(t, signer.ValidateFetch(req, &resp))

	resp.Header.Set("Content-Type", `text/html; charset="utf-8"`)
	assert.NoError(t, signer.ValidateFetch(req, &resp))
}
