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

func TestUnusedExtensions(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc: "removes unused extension",
			Input: tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPAd, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"</body></html>"),
		},
		{
			Desc: "keeps used extension",
			Input: tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPAd, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<amp-ad></amp-ad>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPAd, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<amp-ad></amp-ad>",
				"</body></html>"),
		},
		{
			Desc: "keeps used extension (lowercase)",
			Input: tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPAd, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<AMP-AD></AMP-AD>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPAd, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<AMP-AD></AMP-AD>",
				"</body></html>"),
		},
		{
			Desc: "keeps used extension (entity)",
			Input: tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				"<script async custom-element=amp&#45;ad src=https://cdn.ampproject.org/v0/amp-ad-0.1.js></script>",
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<amp-ad></amp-ad>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPAd, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<amp-ad></amp-ad>",
				"</body></html>"),
		},
		{
			Desc: "keeps used extension (uppercase)",
			Input: tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				"<SCRIPT ASYNC CUSTOM-ELEMENT=amp&#45;ad SRC=https://cdn.ampproject.org/v0/amp-ad-0.1.js></script>",
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<amp-ad></amp-ad>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPAd, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<amp-ad></amp-ad>",
				"</body></html>"),
		},
		{
			Desc: "keeps used extension (amp-access)",
			Input: tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPAccess,
				tt.LinkFavicon,
				"<script id=amp-access></script>",
				tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPAccess,
				tt.LinkFavicon,
				"<script id=amp-access></script>",
				tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"</body></html>"),
		},
		{
			Desc: "keeps used extension (amp-embed)",
			Input: tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPAd, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<amp-embed></amp-embed>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPAd, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<amp-embed></amp-embed>",
				"</body></html>"),
		},
		{
			Desc: "keeps used extension (amp-form)",
			Input: tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPForm, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<form></form>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPForm, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<form></form>",
				"</body></html>"),
		},
		{
			Desc: "keeps used extension (amp-mustache)",
			Input: tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPMustache, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<template></template>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPMustache, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<template></template>",
				"</body></html>"),
		},
		{
			Desc: "keeps non-grandfathered extension",
			Input: tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPAnalytics, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.ScriptAMPAnalytics, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"</body></html>"),
		},
	}
	for _, tc := range tcs {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s: html.Parse for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		inputDOM, err := amphtml.NewDOM(inputDoc)
		if err != nil {
			t.Errorf("%s\namphtml.NewDOM for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		transformers.UnusedExtensions(&transformers.Context{DOM: inputDOM})

		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
			t.Errorf("%s: html.Render for %s failed %q", tc.Desc, tc.Input, err)
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
			t.Errorf("%s: UnusedExtensions=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
