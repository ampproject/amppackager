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

package transformers_test

import (
	"net/url"
	"strings"
	"testing"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	tt "github.com/ampproject/amppackager/transformer/internal/testing"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

const (
	PublisherURL             = "https://publisher.com/amp-url.html"
	LinkGoogleFontPreconnect = "<link href=\"https://publisher.com\" rel=\"dns-prefetch preconnect\"/>"
)

func TestLinkTag(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc: "Adds link for Google Font Preconnect",
			Input: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.LinkGoogleFont, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡=\"\"><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.LinkGoogleFontPreconnect, tt.LinkGoogleFont,
				tt.LinkCanonical, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				LinkGoogleFontPreconnect,
				"</head><body></body></html>"),
		},
		{
			Desc: "Adds link for Google Font Preconnect only once",
			Input: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.LinkGoogleFont, tt.LinkGoogleFont,
				tt.LinkCanonical, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡=\"\"><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.LinkGoogleFontPreconnect, tt.LinkGoogleFont,
				tt.LinkGoogleFont, tt.LinkCanonical, tt.StyleAMPBoilerplate,
				tt.NoscriptAMPBoilerplate, LinkGoogleFontPreconnect,
				"</head><body></body></html>"),
		},
	}
	for _, tc := range tcs {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s: html.Parse on %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		inputDOM, err := amphtml.NewDOM(inputDoc)
		if err != nil {
			t.Errorf("%s\namphtml.NewDOM for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		context := transformers.Context{DOM: inputDOM}
		context.DocumentURL, err = url.Parse(PublisherURL)
		transformers.LinkTag(&context)

		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
			t.Errorf("%s: html.Render on %s failed %q", tc.Desc, tc.Input, err)
			continue
		}

		expectedDoc, err := html.Parse(strings.NewReader(tc.Expected))
		if err != nil {
			t.Errorf("%s: html.Parse for %s failed %q", tc.Desc, tc.Expected, err)
			continue
		}
		var expected strings.Builder
		if err := html.Render(&expected, expectedDoc); err != nil {
			t.Errorf("%s: html.Render for %s failed %q", tc.Desc, tc.Expected, err)
			continue
		}
		if input.String() != expected.String() {
			t.Errorf("%s: LinkTag=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
