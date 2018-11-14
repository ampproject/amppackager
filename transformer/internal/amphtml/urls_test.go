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

package amphtml

import (
	"net/url"
	"testing"
)

const (
	fooBaseURL  = "https://www.example.com/foo"
	barBaseURL  = "https://www.example.com/bar"
	relativeURL = "/foo"
)

func TestToURLs(t *testing.T) {
	tcs := []struct {
		desc, input, expectedPortable, expectedAbsolute, baseURL string
	}{
		{
			desc:             "Empty",
			input:            "",
			expectedPortable: "",
			expectedAbsolute: "",
			baseURL:          barBaseURL,
		},
		{
			desc:             "protocol relative path",
			input:            "//domain.com",
			expectedPortable: "https://domain.com",
			expectedAbsolute: "https://domain.com",
			baseURL:          barBaseURL,
		},
		{
			desc:             "invalid",
			input:            "file://foo.txt",
			expectedPortable: "file://foo.txt",
			expectedAbsolute: "file://foo.txt",
			baseURL:          barBaseURL,
		},
		{
			desc:             "valid absolute",
			input:            fooBaseURL,
			expectedPortable: fooBaseURL,
			expectedAbsolute: fooBaseURL,
			baseURL:          barBaseURL,
		},
		{
			desc:             "valid relative",
			input:            relativeURL,
			expectedPortable: fooBaseURL,
			expectedAbsolute: fooBaseURL,
			baseURL:          barBaseURL,
		},
		{
			desc:             "same replaced with fragment",
			input:            barBaseURL,
			expectedPortable: "#",
			expectedAbsolute: barBaseURL,
			baseURL:          barBaseURL,
		},
		{
			desc:             "fragment same base",
			input:            barBaseURL + "#dogs",
			expectedPortable: "#dogs",
			expectedAbsolute: barBaseURL + "#dogs",
			baseURL:          barBaseURL,
		},
		{
			desc:             "fragment different base",
			input:            barBaseURL + "#dogs",
			expectedPortable: barBaseURL + "#dogs",
			expectedAbsolute: barBaseURL + "#dogs",
			baseURL:          "http://otherdomain.com",
		},
	}
	for _, tc := range tcs {
		baseParsed, _ := url.Parse(tc.baseURL)
		actual := ToAbsoluteURL(baseParsed, tc.input)
		if actual != tc.expectedAbsolute {
			t.Errorf("%s: ToAbsoluteURL=%s want=%s", tc.desc, actual, tc.expectedAbsolute)
		}

		actual = ToPortableURL(baseParsed, tc.input)
		if actual != tc.expectedPortable {
			t.Errorf("%s: ToPortableURL=%s want=%s", tc.desc, actual, tc.expectedPortable)
		}
	}
}
