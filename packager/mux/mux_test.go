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

// Propagate hardcoded values into template test url.
func Expand(templateURL string) string {
	templateURL = strings.Replace(templateURL, "$HOST",
		"http://www.publisher_amp_server.com", 1)
	templateURL = strings.Replace(templateURL, "$FETCH",
		"http://www.publisher_main_server.com/some_page", 1)
	templateURL = strings.Replace(templateURL, "$SIGN",
		"https://www.publisher_main_server.com/some_page", 1)
	templateURL = strings.Replace(templateURL, "$CERT", pkgt.CertName, 1)
	return templateURL
}

// Dedicated type for annotated params - for tests readability.
type params map[string]string

// Mock for underlying http handlers - signer, cert etc.
type MockedHandler struct {
	mock.Mock
}

// Mock ServeHTTP: record annotated params, don't forward the call.
func (m *MockedHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// Convert annotated parameters to dedicated "params" type for tests readability.
	m.Called(params(Params(req)))
}

func TestServeHTTPSuccess(t *testing.T) {
	tests := []struct {
		expectedHandlerName  string
		testScenario         string
		url                  string
		expectedParsedParams params
	}{
		{"signer",
			" empty",
			Expand("$HOST/priv/doc"),
			params{}},
		{"signer",
			" with query, empty",
			Expand("$HOST/priv/doc?"),
			params{}},
		{"signer",
			" with query, regular",
			Expand("$HOST/priv/doc?fetch=$FETCH&sign=$SIGN"),
			params{}},
		{"signer",
			" with query, escaping",
			Expand("$HOST/priv/doc?fetch=$FETCH&sign=$SIGN%2A\\"),
			params{}},
		{"signer",
			" with path, empty",
			Expand("$HOST/priv/doc/"),
			params{"signURL": ""}},
		{"signer",
			"  with path, regular",
			Expand("$HOST/priv/doc/$FETCH"),
			params{"signURL": Expand("$FETCH")}},
		{"signer",
			" with path, escaping",
			Expand("$HOST/priv/doc/$FETCH%2A\\"),
			params{"signURL": Expand("$FETCH%2A%5C")}},
		{"signer",
			" with path and query, regular",
			Expand("$HOST/priv/doc/$FETCH?amp=1"),
			params{"signURL": Expand("$FETCH?amp=1")}},
		{"signer",
			" with path and query, escaping",
			Expand("$HOST/priv/doc/$FETCH%2A\\?amp=1%2A\\"),
			params{"signURL": Expand("$FETCH%2A%5C?amp=1%2A\\")}},
		{"cert",
			" empty",
			Expand("$HOST/amppkg/cert/"),
			params{"certName": ""}},
		{"cert",
			" regular",
			Expand("$HOST/amppkg/cert/$CERT"),
			params{"certName": Expand("$CERT")}},
		{"cert",
			" escaping",
			Expand("$HOST/amppkg/cert/$CERT%2A\\"),
			params{"certName": Expand("$CERT*\\")}},
		{"validityMap",
			" regular",
			Expand("$HOST/amppkg/validity"),
			params{}},
		{"healthz",
			" regular",
			Expand("$HOST/healthz"),
			params{}},
	}
	for _, tt := range tests {
		testName := tt.expectedHandlerName + tt.testScenario
		t.Run(testName, func(t *testing.T) {
			mocks := map[string](*MockedHandler){"signer": &MockedHandler{}, "healthz": &MockedHandler{}, "cert": &MockedHandler{}, "validityMap": &MockedHandler{}}
			mux := New(mocks["cert"], mocks["signer"], mocks["validityMap"], mocks["healthz"])

			// Set expectation.
			expectedMockedHandler := mocks[tt.expectedHandlerName]
			expectedMockedHandler.On("ServeHTTP", tt.expectedParsedParams)

			// Run.
			actualResp := pkgt.Get(t, mux, tt.url)

			// Expect no errors.
			assert.Equal(t, 200, actualResp.StatusCode, "No error expected: %#v", actualResp)

			// Expect the right call to the right handler, and no calls to the rest.
			for _, mockedHandler := range mocks {
				mockedHandler.AssertExpectations(t)
			}
		})
	}
}

func ExpectError(t *testing.T, url string, expectedErrorMessage string, expectedErrorCode int, body io.Reader) {
	// Initialize mux with 4 identical mocked handlers, because no calls are expected to any of them.
	mockedHandler := new(MockedHandler)
	mux := New(mockedHandler, mockedHandler, mockedHandler, mockedHandler)

	// Run and extract error.
	actualResp := pkgt.GetBHH(t, mux, url, "", body, http.Header{})
	actualErrorMessageBuffer, _ := ioutil.ReadAll(actualResp.Body)
	actualErrorMessage := fmt.Sprintf("%s", actualErrorMessageBuffer)

	// Expect the right error.
	assert.Equal(t, expectedErrorCode, actualResp.StatusCode,
		"incorrect error code: %#v", actualResp)
	assert.Equal(t, expectedErrorMessage, actualErrorMessage,
		"incorrect error message: %#v", actualErrorMessage)

	// Expect no calls to mocks.
	mockedHandler.AssertExpectations(t)
}

func TestServeHTTPExpected404s(t *testing.T) {
	tests := []struct {
		testName string
		url      string
	}{
		{"No such endpoint                      ", Expand("$HOST/abc")},
		{"Signer - unexpected extra char        ", Expand("$HOST/priv/doc1")},
		{"Cert - no closing slash               ", Expand("$HOST/amppkg/cert")},
		{"ValidityMap - unexpected closing slash", Expand("$HOST/amppkg/validity/")},
		{"Healthz - unexpected closing slash    ", Expand("$HOST/healthz/")},
		{"Healthz - unexpected extra char       ", Expand("$HOST/healthz1")},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ExpectError(t, tt.url, "404 page not found\n", http.StatusNotFound, nil)
		})
	}
}

func TestServeHTTPExpected405(t *testing.T) {
	body := strings.NewReader("Non empty body so GetBHH sends a POST request")
	ExpectError(t, Expand("$HOST/healthz"),
		"405 method not allowed\n", http.StatusMethodNotAllowed, body)
}

func TestParamsIncorrectValueType(t *testing.T) {
	req := httptest.NewRequest("", "http://abc.com", nil)

	// Pass string instead of expected map[string]string.
	req = req.WithContext(context.WithValue(req.Context(), paramsKey, "Some string"))

	// Expect Params to handle invalid input gracefully.
	assert.Equal(t, Params(req), map[string]string{})
}
