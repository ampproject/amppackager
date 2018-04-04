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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func newPackager(t *testing.T, urlSets []URLSet) *Packager {
	handler, err := NewPackager(cert, key, "https://example.com/", urlSets)
	if err != nil {
		t.Fatal(errors.WithStack(err))
	}
	return handler
}

func stringPtr(s string) *string { return &s }

func TestSimple(t *testing.T) {
	urlSets := []URLSet{URLSet{
		SamePath: true,
		Fetch:    URLPattern{[]string{"http"}, "example.com", stringPtr("/amp/.*"), []string{}, stringPtr(""), false},
		Sign:     URLPattern{[]string{"https"}, "example.com", stringPtr("/amp/.*"), []string{}, stringPtr(""), false}}}
	resp := get(t, newPackager(t, urlSets), `/priv/doc?fetch=http%3A%2F%2Fexample.com%2Famp%2Fsecret-life-of-pine-trees.html&sign=https%3A%2F%2Fexample.com%2Famp%2Fsecret-life-of-pine-trees.html`)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("incorrect status: %#v", resp)
	}
	_, _ = ioutil.ReadAll(resp.Body)
	// TODO(twifkak): Test the body somehow.
}

func TestMain(m *testing.M) {
	// Mock out AMP CDN endpoint.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("yum yum yum"))
	}))
	defer server.Close()
	url, _ := url.Parse(server.URL)
	ampCDNBase = url.Host

	os.Exit(m.Run())
}

// TODO(twifkak): Write lots more tests.
// TODO(twifkak): Fuzz-test.
