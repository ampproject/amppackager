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

const relativeURL = "/foo"

func TestToURLs(t *testing.T) {
	fooBaseURL, _ := url.Parse("https://www.example.com/foo")
	barBaseURL, _ := url.Parse("https://www.example.com/bar")
	otherURL, _ := url.Parse("http://otherdomain.com")

	tcs := []struct {
		desc, input, expectedPortable, expectedAbsolute string
		baseURL                                         *url.URL
	}{
		{
			desc:             "Empty",
			input:            "",
			expectedPortable: "",
			expectedAbsolute: "",
			baseURL:          barBaseURL,
		},
		{
			desc:             "Null base",
			input:            fooBaseURL.String(),
			expectedPortable: fooBaseURL.String(),
			expectedAbsolute: fooBaseURL.String(),
			baseURL:          nil,
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
			input:            fooBaseURL.String(),
			expectedPortable: fooBaseURL.String(),
			expectedAbsolute: fooBaseURL.String(),
			baseURL:          barBaseURL,
		},
		{
			desc:             "valid relative",
			input:            relativeURL,
			expectedPortable: fooBaseURL.String(),
			expectedAbsolute: fooBaseURL.String(),
			baseURL:          barBaseURL,
		},
		{
			desc:             "same replaced with fragment",
			input:            barBaseURL.String(),
			expectedPortable: "#",
			expectedAbsolute: barBaseURL.String(),
			baseURL:          barBaseURL,
		},
		{
			desc:             "fragment same base",
			input:            barBaseURL.String() + "#dogs",
			expectedPortable: "#dogs",
			expectedAbsolute: barBaseURL.String() + "#dogs",
			baseURL:          barBaseURL,
		},
		{
			desc:             "fragment different base",
			input:            barBaseURL.String() + "#dogs",
			expectedPortable: barBaseURL.String() + "#dogs",
			expectedAbsolute: barBaseURL.String() + "#dogs",
			baseURL:          otherURL,
		},
	}
	for _, tc := range tcs {
		actual := ToAbsoluteURL(tc.baseURL, tc.input)
		if actual != tc.expectedAbsolute {
			t.Errorf("%s: ToAbsoluteURL=%s want=%s", tc.desc, actual, tc.expectedAbsolute)
		}

		actual = ToPortableURL(tc.baseURL, tc.input)
		if actual != tc.expectedPortable {
			t.Errorf("%s: ToPortableURL=%s want=%s", tc.desc, actual, tc.expectedPortable)
		}
	}
}
