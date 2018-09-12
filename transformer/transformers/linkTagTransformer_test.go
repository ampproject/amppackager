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
	"strings"
	"testing"

	tt "github.com/ampproject/amppackager/transformer/internal/testing"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

// These tests do NOT run through the custom transformations of the
// Engine, and instead rely exclusively on vanilla golang parser and
// renderer (otherwise the scope of these tests would creep past unit
// testing). Therefore, the test data must be made to match, and is not
// the expected normalized output from transformer.go, nor from any other
// transformers.

func TestLinkTagTransformer(t *testing.T) {
	testCases := []tt.TestCase{
		{
			Desc: "Adds link for Google Font Preconnect",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.ScriptAMPRuntime, tt.MetaCharset, tt.StyleAMPBoilerplate,
				tt.NoscriptAMPBoilerplate, tt.LinkGoogleFont,
				"</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡=\"\"><head>",
				tt.ScriptAMPRuntime, tt.MetaCharset, tt.StyleAMPBoilerplate,
				tt.NoscriptAMPBoilerplate, tt.LinkGoogleFontPreconnect, tt.LinkGoogleFont,
				"</head><body></body></html>"),
		},
		{
			Desc: "Adds link for Google Font Preconnect only once",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.ScriptAMPRuntime, tt.MetaCharset, tt.StyleAMPBoilerplate,
				tt.NoscriptAMPBoilerplate, tt.LinkGoogleFont, tt.LinkGoogleFont,
				"</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡=\"\"><head>",
				tt.ScriptAMPRuntime, tt.MetaCharset, tt.StyleAMPBoilerplate,
				tt.NoscriptAMPBoilerplate, tt.LinkGoogleFontPreconnect, tt.LinkGoogleFont,
				tt.LinkGoogleFont, "</head><body></body></html>"),
		},
		{
			Desc: "Renames author supplied resource hints",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.ScriptAMPRuntime, tt.MetaCharset, tt.StyleAMPBoilerplate,
				"<link href=https://cdn.ampproject.org/ rel=preconnect>",
				tt.NoscriptAMPBoilerplate, "</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡=\"\"><head>",
				tt.ScriptAMPRuntime, tt.MetaCharset, tt.StyleAMPBoilerplate,
				"<link href=https://cdn.ampproject.org/ disabled-rel=preconnect>",
				tt.NoscriptAMPBoilerplate, "</head><body></body></html>"),
		},
	}
	for _, tc := range testCases {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s: html.Parse on %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		transformers.LinkTagTransformer(&transformers.Engine{Doc: inputDoc})

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
			t.Errorf("%s: LinkTagTransformer=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
