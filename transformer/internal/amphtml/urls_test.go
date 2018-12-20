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
			desc:             "absolute with different base",
			input:            fooBaseURL.String(),
			expectedPortable: fooBaseURL.String(),
			expectedAbsolute: fooBaseURL.String(),
			baseURL:          otherURL,
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

func TestGetCacheURL(t *testing.T) {
	tcs := []struct {
		desc, input, expectedImage, expectedOther string
		width                                     int
		expectError                               bool
	}{
		{
			desc:        "empty string",
			input:       "",
			expectError: true,
		},
		{
			desc:          "image",
			input:         "http://www.example.com/blah.jpg",
			expectedImage: "https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg",
			expectedOther: "https://www-example-com.cdn.ampproject.org/r/www.example.com/blah.jpg",
		},
		{
			desc:          "secure",
			input:         "https://www.example.com/blah.jpg",
			expectedImage: "https://www-example-com.cdn.ampproject.org/i/s/www.example.com/blah.jpg",
			expectedOther: "https://www-example-com.cdn.ampproject.org/r/s/www.example.com/blah.jpg",
		},
		{
			desc:          "image with requested width",
			input:         "http://www.example.com/blah.jpg",
			width:         50,
			expectedImage: "https://www-example-com.cdn.ampproject.org/ii/w50/www.example.com/blah.jpg",
			expectedOther: "https://www-example-com.cdn.ampproject.org/r/www.example.com/blah.jpg",
		},
		{
			desc:          "image negative width",
			input:         "http://www.example.com/blah.jpg",
			width:         -50,
			expectedImage: "https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg",
			expectedOther: "https://www-example-com.cdn.ampproject.org/r/www.example.com/blah.jpg",
		},
		{
			desc:          "fragment",
			input:         "https://localhost.test/icons/below.svg#icon-whatsapp",
			expectedImage: "https://localhost-test.cdn.ampproject.org/i/s/localhost.test/icons/below.svg#icon-whatsapp",
			expectedOther: "https://localhost-test.cdn.ampproject.org/r/s/localhost.test/icons/below.svg#icon-whatsapp",
		},
		{
			desc:          "port is dropped",
			input:         "http://www.example.com:8080/blah.jpg",
			expectedImage: "https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg",
			expectedOther: "https://www-example-com.cdn.ampproject.org/r/www.example.com/blah.jpg",
		},
		{
			desc:        "unsupported scheme noop",
			input:       "data:image/png.foo",
			expectError: true,
		},
		{
			desc:        "unsupported scheme with width",
			input:       "itshappening.gif",
			width:       100,
			expectError: true,
		},
	}
	base, _ := url.Parse("")
	for _, tc := range tcs {
		for _, subtype := range []SubresourceType{OtherType, ImageType} {
			expected := tc.expectedOther
			if subtype == ImageType {
				expected = tc.expectedImage
			}
			so := SubresourceOffset{SubType: subtype, Start: 0, End: len(tc.input), DesiredImageWidth: tc.width}
			cu, err := so.GetCacheURL(base, tc.input)
			if tc.expectError {
				if err == nil {
					t.Errorf("%s: ToCacheImageURL(%s, %d) expected error. Got none", tc.desc, tc.input, tc.width)
				}
			} else if cu.String() != expected {
				t.Errorf("%s: ToCacheImageURL(%s, %d)=%s, want=%s", tc.desc, tc.input, tc.width, cu.String(), expected)
			}
		}
	}
}

func TestToCacheURLDomain(t *testing.T) {
	tcs := []struct {
		desc, input, expected string
	}{
		{
			desc:     "simple case",
			input:    "example.com",
			expected: "example-com",
		},
		{
			desc:     "simple case, case-insensitive",
			input:    "ExAMpLE.Com",
			expected: "example-com",
		},
		{
			desc:     "origin has no dots or hyphes, use hash",
			input:    "toplevelnohyphens",
			expected: "qsgpfjzulvuaxb66z77vlhb5gu2irvcnyp6t67cz6tqo5ae6fysa",
		},
		{
			desc:     "Human-readable form too long; use hash",
			input:    "itwasadarkandstormynight.therainfellintorrents.exceptatoccasionalintervalswhenitwascheckedby.aviolentgustofwindwhichsweptupthestreets.com",
			expected: "dgz4cnrxufaulnwku4ow5biptyqnenjievjht56hd7wqinbdbteq",
		},
		{
			desc:     "IDN",
			input:    "xn--bcher-kva.ch",
			expected: "xn--bcher-ch-65a",
		},
		{
			desc:     "RTL",
			input:    "xn--4gbrim.xn----rmckbbajlc6dj7bxne2c.xn--wgbh1c",
			expected: "xn-------i5fvcbaopc6fkc0de0d9jybegt6cd",
		},
		{
			desc:     "Mixed Bidi, use hash",
			input:    "hello.xn--4gbrim.xn----rmckbbajlc6dj7bxne2c.xn--wgbh1c",
			expected: "a6h5moukddengbsjm77rvbosevwuduec2blkjva4223o4bgafgla",
		},
		{
			desc:     "Punify(ز۰.ز٠) = xn--xgb49a.xn--xgb6g. Cannot mix two alternative Arabic ranges. Use hash",
			input:    "xn--xgb49a.xn--xgb6g",
			expected: "asdk26k2mfqxgc6cdx3oh3vlnx42rqwn6uvsuqrufnx622tguq6q",
		},
		{
			desc:     "Same Arabic range is ok",
			input:    "xn--xgb49a.xn--xgb49a",
			expected: "xn----lncb27eca",
		},
		{
			desc:     "R-LDH: cannot contain double hyphen in 3 and 4th char positions",
			input:    "in--trouble.com",
			expected: "r5s7rxu53tjelpr7ngbxkxpirbrylvbwcuueckh7gmn5mim5cjna",
		},
		{
			desc:     "R-LDH #2",
			input:    "in-trouble.com",
			expected: "j7pweznglei73fva3bo6oidjt74j3hx4tfyncjsdwud7r7cci4va",
		},
		{
			desc:     "R-LDH #3",
			input:    "a--problem.com",
			expected: "a47psvede4jpgjom2kzmuhop74zzmdpjzasoctyoqqaxbkdbsyiq",
		},
		{
			desc:     "Transition mapping per UTS #46",
			input:    "faß.de",
			expected: "fass-de",
		},
	}
	for _, tc := range tcs {
		actual := ToCacheURLSubdomain(tc.input)
		if actual != tc.expected {
			t.Errorf("%s: ToCacheURLDomain(%s)=%s, want=%s", tc.desc, tc.input, actual, tc.expected)
		}
	}
}
