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

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	tt "github.com/ampproject/amppackager/transformer/internal/testing"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

// ampBoilerplateNoscriptWithAttr is the equivalent of
// tt.AMPBoilerplateNoscript except it includes empty attribute
// values. This is needed because the golang parser treats everything
// inside <noscript> as one giant text node, so the <noscript> tag in
// the expected HTML isn't parsed properly.
const ampBoilerplateNoscriptWithAttr = "<noscript><style amp-boilerplate=\"\">body{-webkit-animation:none;-moz-animation:none;-ms-animation:none;animation:none}</style></noscript>"

func TestAMPBoilerplate(t *testing.T) {
	canonicalExpected := tt.Concat(tt.Doctype, "<html ⚡><head>",
		tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
		tt.LinkFavicon, tt.LinkCanonical, tt.StyleAMPBoilerplate,
		ampBoilerplateNoscriptWithAttr, "</head><body></body></html>")

	tcs := []tt.TestCase{
		{
			Desc:     "Keeps boilerplate",
			Input:    tt.Concat(tt.Doctype, "<html ⚡><head>",
				  tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime, tt.LinkFavicon,
				  tt.LinkCanonical, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				  "</head><body></body></html>"),
			Expected: canonicalExpected,
		},
		{
			Desc:     "Adds boilerplate if missing",
			Input:    tt.Concat(tt.Doctype, "<html ⚡><head>",
				  tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				  tt.LinkFavicon, tt.LinkCanonical,
				  "</head><body></body></html>"),
			Expected: canonicalExpected,
		},
	}

	runAMPBoilerplateTestcases(t, tcs)
}

func TestAMP4Ads(t *testing.T) {
	expected := func(attr string) string {
		return tt.Concat(tt.Doctype, "<html ", attr, "><head>",
			tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
			"<style amp4ads-boilerplate>body{visibility:hidden}</style></head>",
			"<body></body></html>")
	}

	tcs := []tt.TestCase{
		{
			Desc:     "Keeps boilerplate",
			Input:    expected("amp4ads"),
			Expected: expected("amp4ads"),
		},
		{
			Desc:     "Keeps boilerplate for ⚡4ads",
			Input:    expected("⚡4ads"),
			Expected: expected("⚡4ads"),
		},
		{
			Desc:     "Adds boilerplate if missing",
			Input:    tt.Concat(tt.Doctype, "<html amp4ads><head>",
				  tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				  "</head><body></body></html>"),
			Expected: expected("amp4ads"),
		},
		{
			Desc:     "Adds boilerplate for ⚡4ads if missing",
			Input:    tt.Concat(tt.Doctype, "<html ⚡4ads><head>",
				  tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				  "</head><body></body></html>"),
			Expected: expected("⚡4ads"),
		},
	}

	runAMPBoilerplateTestcases(t, tcs)
}

func TestAMP4Email(t *testing.T) {
	expected := func(attr string) string {
		return tt.Concat(tt.Doctype, "<html ", attr, "><head>",
			tt.MetaCharset, tt.ScriptAMPRuntime,
			"<style amp4email-boilerplate>body{visibility:hidden}</style></head>",
			"<body></body></html>")
	}

	tcs := []tt.TestCase{
		{
			Desc:     "Keeps boilerplate",
			Input:    expected("amp4email"),
			Expected: expected("amp4email"),
		},
		{
			Desc:     "Keeps boilerplate for ⚡4email",
			Input:    expected("⚡4email"),
			Expected: expected("⚡4email"),
		},
		{
			Desc:     "Adds boilerplate if missing",
			Input:    tt.Concat(tt.Doctype, "<html amp4email><head>",
				  tt.MetaCharset, tt.ScriptAMPRuntime,
				  "</head><body></body></html>"),
			Expected: expected("amp4email"),
		},
		{
			Desc:     "Adds boilerplate for ⚡4email if missing",
			Input:    tt.Concat(tt.Doctype, "<html ⚡4email><head>",
				  tt.MetaCharset, tt.ScriptAMPRuntime,
				  "</head><body></body></html>"),
			Expected: expected("⚡4email"),
		},
	}

	runAMPBoilerplateTestcases(t, tcs)
}

func runAMPBoilerplateTestcases(t *testing.T, tcs []tt.TestCase) {
	for _, tc := range tcs {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s\nhtml.Parse for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		inputDOM, err := amphtml.NewDOM(inputDoc)
		if err != nil {
			t.Errorf("%s\namphtml.NewDOM for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		transformers.AMPBoilerplate(&transformers.Context{DOM: inputDOM})

		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
			t.Errorf("%s\nhtml.Render for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}

		expectedDoc, err := html.Parse(strings.NewReader(tc.Expected))
		if err != nil {
			t.Errorf("%s\nhtml.Parse for %s failed %q", tc.Desc, tc.Expected, err)
			continue
		}
		var expected strings.Builder
		if err := html.Render(&expected, expectedDoc); err != nil {
			t.Errorf("%s\nhtml.Render for %s failed %q", tc.Desc, tc.Expected, err)
			continue
		}
		if input.String() != expected.String() {
			t.Errorf("%s: AMPBoilerplate=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
