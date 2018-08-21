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
	"encoding/binary"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var httpURL, httpsURL string
var httpsClient *http.Client

var fakeBody = []byte("They like to OPINE. Get it? (Is he fir real? Yew gotta be kidding me.)")
var lastRequestURL string

// Don't override this manually; use replacingFakeHandler() instead.
var fakeHandler = func(w http.ResponseWriter, req *http.Request) {
	lastRequestURL = req.URL.String()
	w.Write(fakeBody)
}

func fetchURL(serverURL string) *url.URL {
	u, err := url.Parse(serverURL)
	if err != nil {
		log.Panic("parsing server url:", err)
	}
	u.Path = "/amp/secret-life-of-pine-trees.html"
	return u
}

func signURL(serverURL string) *url.URL {
	u := fetchURL(serverURL)
	u.Scheme = "https"
	return u

}

func replacingFakeHandler(newFake func(w http.ResponseWriter, req *http.Request), testCode func()) {
	oldFake := fakeHandler
	defer func() { fakeHandler = oldFake }()
	fakeHandler = newFake
	testCode()
}

func newPackager(t *testing.T, urlSets []URLSet) *Packager {
	handler, err := NewPackager(cert, key, "https://example.com/", urlSets)
	if err != nil {
		t.Fatal(errors.WithStack(err))
	}
	// Accept the self-signed certificate generated by the test server.
	handler.client = httpsClient
	return handler
}

func boolPtr(x bool) *bool       { return &x }
func stringPtr(x string) *string { return &x }

func headerNames(headers http.Header) []string {
	names := make([]string, len(headers))
	i := 0
	for name := range headers {
		names[i] = strings.ToLower(name)
		i++
	}
	sort.Strings(names)
	return names
}

func TestSimple(t *testing.T) {
	urlSets := []URLSet{URLSet{
		Sign:  &URLPattern{[]string{"https"}, "", signURL(httpURL).Host, stringPtr("/amp/.*"), []string{}, stringPtr(""), false, nil},
		Fetch: &URLPattern{[]string{"http"}, "", fetchURL(httpURL).Host, stringPtr("/amp/.*"), []string{}, stringPtr(""), false, boolPtr(true)},
	}}
	resp := get(t, newPackager(t, urlSets),
		"/priv/doc?fetch="+url.QueryEscape(fetchURL(httpURL).String())+"&sign="+url.QueryEscape(signURL(httpURL).String()))
	assert.Equal(t, http.StatusOK, resp.StatusCode, "incorrect status: %#v", resp)

	exchange, err := signedexchange.ReadExchange(resp.Body)
	if assert.NoError(t, err) {
		assert.Equal(t, "/amp/secret-life-of-pine-trees.html", lastRequestURL)
		assert.Equal(t, signURL(httpURL).String(), exchange.RequestURI.String())
		assert.Equal(t, http.Header{":method": []string{"GET"}}, exchange.RequestHeaders)
		assert.Equal(t, 200, exchange.ResponseStatus)
		assert.Equal(t, []string{"content-encoding", "content-length", "content-type", "date", "mi-draft2"}, headerNames(exchange.ResponseHeaders))
		// The response header values are untested here, as that is covered by signedexchange tests.

		// For small enough bodies, the only thing that MICE does is add a record size prefix.
		var payloadPrefix bytes.Buffer
		binary.Write(&payloadPrefix, binary.BigEndian, uint64(miRecordSize))
		assert.Equal(t, append(payloadPrefix.Bytes(), fakeBody...), exchange.Payload)
	}
}

func TestNoFetchParam(t *testing.T) {
	urlSets := []URLSet{URLSet{
		Sign: &URLPattern{[]string{"https"}, "", signURL(httpsURL).Host, stringPtr("/amp/.*"), []string{}, stringPtr(""), false, nil},
	}}
	resp := get(t, newPackager(t, urlSets), "/priv/doc?sign="+url.QueryEscape(signURL(httpsURL).String()))
	assert.Equal(t, http.StatusOK, resp.StatusCode, "incorrect status: %#v", resp)

	exchange, err := signedexchange.ReadExchange(resp.Body)
	if assert.NoError(t, err) {
		assert.Equal(t, "/amp/secret-life-of-pine-trees.html", lastRequestURL)
		assert.Equal(t, signURL(httpsURL).String(), exchange.RequestURI.String())
	}
}

func TestSignAsPathParam(t *testing.T) {
	urlSets := []URLSet{URLSet{
		Sign: &URLPattern{[]string{"https"}, "", signURL(httpsURL).Host, stringPtr("/amp/.*"), []string{}, stringPtr(""), false, nil},
	}}
	resp := getP(t, newPackager(t, urlSets), `/priv/doc/`, httprouter.Params{httprouter.Param{"signURL", "/" + signURL(httpsURL).String()}})
	assert.Equal(t, http.StatusOK, resp.StatusCode, "incorrect status: %#v", resp)

	exchange, err := signedexchange.ReadExchange(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, "/amp/secret-life-of-pine-trees.html", lastRequestURL)
	assert.Equal(t, signURL(httpsURL).String(), exchange.RequestURI.String())
}

func TestErrorNoCache(t *testing.T) {
	urlSets := []URLSet{URLSet{
		Fetch: &URLPattern{[]string{"http"}, "", fetchURL(httpURL).Host, stringPtr("/amp/.*"), []string{}, stringPtr(""), false, boolPtr(true)},
	}}
	// Missing sign param generates an error.
	resp := get(t, newPackager(t, urlSets), "/priv/doc?fetch="+url.QueryEscape(fetchURL(httpURL).String()))
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "incorrect status: %#v", resp)
	assert.Equal(t, "no-store", resp.Header.Get("Cache-Control"))
}

func TestRedirectIsProxiedUnsigned(t *testing.T) {
	urlSets := []URLSet{URLSet{
		Sign: &URLPattern{[]string{"https"}, "", signURL(httpsURL).Host, stringPtr("/amp/.*"), []string{}, stringPtr(""), false, nil},
	}}
	replacingFakeHandler(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("cookie", "yum yum yum")
		w.Header().Set("location", "/login")
		w.WriteHeader(301)
	}, func() {
		resp := get(t, newPackager(t, urlSets), "/priv/doc?sign="+url.QueryEscape(signURL(httpsURL).String()))
		assert.Equal(t, 301, resp.StatusCode)
		assert.Equal(t, "", resp.Header.Get("cookie"))
		assert.Equal(t, "/login", resp.Header.Get("location"))
	})
}

func TestNotModifiedIsProxiedUnsigned(t *testing.T) {
	urlSets := []URLSet{URLSet{
		Sign: &URLPattern{[]string{"https"}, "", signURL(httpsURL).Host, stringPtr("/amp/.*"), []string{}, stringPtr(""), false, nil},
	}}
	replacingFakeHandler(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("cache-control", "private")
		w.Header().Set("cookie", "yum yum yum")
		w.Header().Set("etag", "superrad")
		w.WriteHeader(304)
	}, func() {
		resp := get(t, newPackager(t, urlSets), "/priv/doc?sign="+signURL(httpsURL).String())
		assert.Equal(t, 304, resp.StatusCode)
		assert.Equal(t, "private", resp.Header.Get("cache-control"))
		assert.Equal(t, "", resp.Header.Get("cookie"))
		assert.Equal(t, "superrad", resp.Header.Get("etag"))
	})
}

func TestMain(m *testing.M) {
	// Mock out example.com endpoint.
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fakeHandler(w, req)
	}))
	defer httpServer.Close()
	httpURL = httpServer.URL

	tlsServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fakeHandler(w, req)
	}))
	defer tlsServer.Close()
	httpsURL = tlsServer.URL
	httpsClient = tlsServer.Client()
	// Configure the test httpsClient to have the same redirect policy as production.
	httpsClient.CheckRedirect = noRedirects

	os.Exit(m.Run())
}

// TODO(twifkak): Write lots more tests.
// TODO(twifkak): Fuzz-test.
