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
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/nyaxt/webpackage/go/signedexchange"
)

var fakeBody = []byte("They like to OPINE. Get it? (Is he fir real? Yew gotta be kidding me.)")
var lastRequestURL string

func newPackager(t *testing.T, urlSets []URLSet) *Packager {
	handler, err := NewPackager(cert, key, "https://example.com/", urlSets)
	if err != nil {
		t.Fatal(errors.WithStack(err))
	}
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
		Sign:  &URLPattern{[]string{"https"}, "", "example.com", stringPtr("/amp/.*"), []string{}, stringPtr(""), false, nil},
		Fetch: &URLPattern{[]string{"http"}, "", "example.com", stringPtr("/amp/.*"), []string{}, stringPtr(""), false, boolPtr(true)},
	}}
	resp := get(t, newPackager(t, urlSets), `/priv/doc?fetch=http%3A%2F%2Fexample.com%2Famp%2Fsecret-life-of-pine-trees.html&sign=https%3A%2F%2Fexample.com%2Famp%2Fsecret-life-of-pine-trees.html`)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("incorrect status: %#v", resp)
	}
	exchange, err := signedexchange.ReadExchangeFile(resp.Body)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "/example.com/amp/secret-life-of-pine-trees.html?usqp=mq331AQCSAE", lastRequestURL)
	assert.Equal(t, "https://example.com/amp/secret-life-of-pine-trees.html", exchange.RequestUri.String())
	assert.Equal(t, http.Header{}, exchange.RequestHeaders)
	assert.Equal(t, 200, exchange.ResponseStatus)
	assert.Equal(t, []string{"content-encoding", "content-length", "content-type", "date", "mi", "signature"}, headerNames(exchange.ResponseHeaders))
	// The response header values are untested here, as that is covered by signedexchange tests.
	assert.Equal(t, fakeBody, exchange.Payload)
}

func TestNoFetchParam(t *testing.T) {
	urlSets := []URLSet{URLSet{
		Sign: &URLPattern{[]string{"https"}, "", "example.com", stringPtr("/amp/.*"), []string{}, stringPtr(""), false, nil},
	}}
	resp := get(t, newPackager(t, urlSets), `/priv/doc?sign=https%3A%2F%2Fexample.com%2Famp%2Fsecret-life-of-pine-trees.html`)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("incorrect status: %#v", resp)
	}
	exchange, err := signedexchange.ReadExchangeFile(resp.Body)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "/s/example.com/amp/secret-life-of-pine-trees.html?usqp=mq331AQCSAE", lastRequestURL)
	assert.Equal(t, "https://example.com/amp/secret-life-of-pine-trees.html", exchange.RequestUri.String())
}

func TestErrorNoCache(t *testing.T) {
	urlSets := []URLSet{URLSet{
		Fetch: &URLPattern{[]string{"http"}, "", "example.com", stringPtr("/amp/.*"), []string{}, stringPtr(""), false, boolPtr(true)},
	}}
	// Missing sign param generates an error.
	resp := get(t, newPackager(t, urlSets), `/priv/doc?fetch=http%3A%2F%2Fexample.com%2Famp%2Fsecret-life-of-pine-trees.html`)
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("incorrect status: %#v", resp)
	}
	assert.Equal(t, "no-store", resp.Header.Get("Cache-Control"))
}

func TestMain(m *testing.M) {
	// Mock out AMP CDN endpoint.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		lastRequestURL = req.URL.String()
		w.Write(fakeBody)
	}))
	defer server.Close()
	url, _ := url.Parse(server.URL)
	AmpCDNBase = "http://" + url.Host + "/"

	os.Exit(m.Run())
}

// TODO(twifkak): Write lots more tests.
// TODO(twifkak): Fuzz-test.
