// Copyright 2020 Google LLC
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

package mux

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	pkgt "github.com/ampproject/amppackager/packager/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// expand propagates hardcoded values into template test url.
func expand(templateURL string) string {
	templateURL = strings.Replace(templateURL, "$HOST", "http://www.publisher_amp_server.com", 1)
	templateURL = strings.Replace(templateURL, "$FETCH", "http://www.publisher_main_server.com/some_page", 1)
	templateURL = strings.Replace(templateURL, "$SIGN", "https://www.publisher_main_server.com/some_page", 1)
	templateURL = strings.Replace(templateURL, "$CERT", pkgt.CertName, 1)
	return templateURL
}

// mockedHandler mocks mux' underlying http handlers - signer, cert etc.
type mockedHandler struct {
	mock.Mock
}

func (m *mockedHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	m.Called(Params(req))
}

func TestServeHTTPSuccess(t *testing.T) {
	templateTests := []struct {
		testName      string
		testURL       string
		expectHandler string
		expectParams  map[string]string
	}{
		{
			testName:      `Signer - empty`,
			testURL:       `$HOST/priv/doc`,
			expectHandler: `signer`,
			expectParams:  map[string]string{},
		}, {
			testName:      `Signer - with query, empty`,
			testURL:       `$HOST/priv/doc?`,
			expectHandler: `signer`,
			expectParams:  map[string]string{},
		}, {
			testName:      `Signer - with query, regular`,
			testURL:       `$HOST/priv/doc?fetch=$FETCH&sign=$SIGN`,
			expectHandler: `signer`,
			expectParams:  map[string]string{},
		},
		{
			testName:      `Signer - with query, escaping`,
			testURL:       `$HOST/priv/doc?fetch=$FETCH&sign=$SIGN%2A\`,
			expectHandler: `signer`,
			expectParams:  map[string]string{},
		}, {
			testName:      `Signer - with path, empty`,
			testURL:       `$HOST/priv/doc/`,
			expectHandler: `signer`,
			expectParams:  map[string]string{`signURL`: ``},
		}, {
			testName:      `Signer -  with path, regular`,
			testURL:       `$HOST/priv/doc/$FETCH`,
			expectHandler: `signer`,
			expectParams:  map[string]string{`signURL`: `$FETCH`},
		}, {
			testName:      `Signer - with path, escaping`,
			testURL:       `$HOST/priv/doc/$FETCH%2A\`,
			expectHandler: `signer`,
			expectParams:  map[string]string{`signURL`: `$FETCH%2A%5C`},
		}, {
			testName:      `Signer - with path and query, regular`,
			testURL:       `$HOST/priv/doc/$FETCH?amp=1`,
			expectHandler: `signer`,
			expectParams:  map[string]string{`signURL`: `$FETCH?amp=1`},
		}, {
			testName:      `Signer - with path and query, escaping`,
			testURL:       `$HOST/priv/doc/$FETCH%2A\?amp=1%2A\`,
			expectHandler: `signer`,
			expectParams:  map[string]string{`signURL`: `$FETCH%2A%5C?amp=1%2A\`},
		}, {
			testName:      `Cert - empty`,
			testURL:       `$HOST/amppkg/cert/`,
			expectHandler: `cert`,
			expectParams:  map[string]string{`certName`: ``},
		}, {
			testName:      `Cert - regular`,
			testURL:       `$HOST/amppkg/cert/$CERT`,
			expectHandler: `cert`,
			expectParams:  map[string]string{`certName`: `$CERT`},
		}, {
			testName:      `Cert - escaping`,
			testURL:       `$HOST/amppkg/cert/$CERT%2A\`,
			expectHandler: `cert`,
			expectParams:  map[string]string{`certName`: `$CERT*\`},
		}, {
			testName:      `ValidityMap - regular`,
			testURL:       `$HOST/amppkg/validity`,
			expectHandler: `validityMap`,
			expectParams:  map[string]string{},
		}, {
			testName:      `Healthz - regular`,
			testURL:       `$HOST/healthz`,
			expectHandler: `healthz`,
			expectParams:  map[string]string{},
		},
	}
	for _, tt := range templateTests {
		testName := tt.testName
		t.Run(testName, func(t *testing.T) {
			// Defer validation to ensure it does happen.
			mocks := map[string](*mockedHandler){"signer": &mockedHandler{}, "healthz": &mockedHandler{}, "cert": &mockedHandler{}, "validityMap": &mockedHandler{}}
			var actualResp *http.Response
			defer func() {
				// Expect no errors.
				assert.Equal(t, 200, actualResp.StatusCode, "No error expected: %#v", actualResp)

				// Expect the right call to the right handler, and no calls to the rest.
				for _, mockedHandler := range mocks {
					mockedHandler.AssertExpectations(t)
				}
			}()

			// expand template URL and expectParams.
			tt.testURL = expand(tt.testURL)
			for v, k := range tt.expectParams {
				tt.expectParams[v] = expand(k)
			}

			// Set expectation.
			expectMockedHandler := mocks[tt.expectHandler]
			expectMockedHandler.On("ServeHTTP", tt.expectParams)

			// Run.
			mux := New(mocks["cert"], mocks["signer"], mocks["validityMap"], mocks["healthz"])
			actualResp = pkgt.Get(t, mux, tt.testURL)
		})
	}
}

func expectError(t *testing.T, url string, expectErrorMessage string, expectErrorCode int, body io.Reader) {
	// Defer validation to ensure it does happen.
	mockedHandler := new(mockedHandler)
	var actualResp *http.Response
	var actualErrorMessage string
	defer func() {
		// Expect the right error.
		assert.Equal(t, expectErrorCode, actualResp.StatusCode)
		assert.Equal(t, expectErrorMessage, actualErrorMessage)

		// Expect no calls to mocks.
		mockedHandler.AssertExpectations(t)
	}()

	// Initialize mux with 4 identical mocked handlers, because no calls are expect to any of them.
	mux := New(mockedHandler, mockedHandler, mockedHandler, mockedHandler)

	// Run and extract error.
	actualResp = pkgt.GetBHH(t, mux, url, "", body, http.Header{})
	actualErrorMessageBuffer, _ := ioutil.ReadAll(actualResp.Body)
	actualErrorMessage = fmt.Sprintf("%s", actualErrorMessageBuffer)

}

func TestServeHTTPexpect404s(t *testing.T) {
	templateTests := []struct {
		testName string
		URL      string
	}{
		{"No such endpoint                      ", "$HOST/abc"},
		{"Signer - unexpected extra char        ", "$HOST/priv/doc1"},
		{"Cert - no closing slash               ", "$HOST/amppkg/cert"},
		{"ValidityMap - unexpected closing slash", "$HOST/amppkg/validity/"},
		{"Healthz - unexpected closing slash    ", "$HOST/healthz/"},
		{"Healthz - unexpected extra char       ", "$HOST/healthz1"},
	}
	for _, tt := range templateTests {
		t.Run(tt.testName, func(t *testing.T) {
			expectError(t, expand(tt.URL), "404 page not found\n", http.StatusNotFound, nil)
		})
	}
}

func TestServeHTTPexpect405(t *testing.T) {
	body := strings.NewReader("Non empty body so GetBHH sends a POST request")
	expectError(t, expand("$HOST/healthz"), "405 method not allowed\n", http.StatusMethodNotAllowed, body)
}

func TestParamsIncorrectValueType(t *testing.T) {
	req := httptest.NewRequest("", "http://abc.com", nil)

	// Pass string instead of expected map[string]string.
	req = req.WithContext(context.WithValue(req.Context(), paramsKey, "Some string"))

	// Expect Params to handle invalid input gracefully.
	assert.Equal(t, Params(req), map[string]string{})
}
