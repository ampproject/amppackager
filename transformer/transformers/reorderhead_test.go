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

func TestReorderHead(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc: "Reorders head children for AMP document",
			Input: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.Title, tt.StyleAMPBoilerplate,
				tt.ScriptAMPExperiment, tt.ScriptAMPAudio,
				tt.NoscriptAMPBoilerplate, tt.StyleAMPRuntime,
				tt.ScriptAMPRuntime, tt.LinkStylesheetGoogleFont,
				tt.LinkGoogleFontPreconnect, tt.MetaCharset,
				tt.MetaViewport, tt.StyleAMPCustom, tt.LinkCanonical,
				tt.LinkFavicon, tt.ScriptAMPViewerRuntime,
				"</head><body></body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				// (0) <meta charset> tag
				tt.MetaCharset,
				// (1) <style amp-runtime> (inserted by ampruntimecss.go)
				tt.StyleAMPRuntime,
				// (2) remaining <meta> tags (those other than <meta charset>)
				tt.MetaViewport,
				// (3) AMP runtime .js <script> tag
				tt.ScriptAMPRuntime,
				// (4) AMP viewer runtime .js <script> tag (inserted by AmpViewerScript)
				tt.ScriptAMPViewerRuntime,
				// (5) <script> tags that are render delaying
				tt.ScriptAMPExperiment,
				// (6) <script> tags for remaining extensions
				tt.ScriptAMPAudio,
				// (7) <link> tag for favicons
				tt.LinkFavicon,
				// (8) <link> tag for resource hints
				tt.LinkGoogleFontPreconnect,
				// (9) <link rel=stylesheet> tags before <style amp-custom>
				tt.LinkStylesheetGoogleFont,
				// (10) <style amp-custom>
				tt.StyleAMPCustom,
				// (11) any other tags allowed in <head>
				tt.Title,
				tt.LinkCanonical,
				// (12) amp boilerplate (first style amp-boilerplate, then noscript)
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Reorders head children for AMP4ADS document",
			Input: tt.Concat(tt.Doctype, "<html ⚡4ads><head>",
				tt.Title, tt.StyleAMP4AdsBoilerplate, tt.ScriptAMPAudio,
				tt.ScriptAMP4AdsRuntime, tt.LinkStylesheetGoogleFont,
				tt.LinkGoogleFontPreconnect, tt.MetaCharset, tt.MetaViewport, tt.StyleAMPCustom,
				"</head><body></body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡4ads><head>",
				// (0) <meta charset> tag
				tt.MetaCharset,
				// (1) <style amp-runtime> (inserted by ampruntimecss.go)
				// n/a for AMP4ADS
				// (2) remaining <meta> tags (those other than <meta charset>)
				tt.MetaViewport,
				// (3) AMP runtime .js <script> tag
				tt.ScriptAMP4AdsRuntime,
				// (4) AMP viewer runtime .js <script> tag (inserted by AmpViewerScript)
				// n/a for AMP4ADS, no viewer
				// (5) <script> tags that are render delaying
				// n/a for AMP4ADS, no render delaying <script> tags allowed
				// (6) <script tags> for remaining extensions
				tt.ScriptAMPAudio,
				// (7) <link> tag for favicons
				// n/a for AMP4ADS, no favicons allowed
				// (8) <link> tag for resource hints
				tt.LinkGoogleFontPreconnect,
				// (9) <link rel=stylesheet> tags before <style amp-custom>
				tt.LinkStylesheetGoogleFont,
				// (10) <style amp-custom>
				tt.StyleAMPCustom,
				// (11) any other tags allowed in <head>
				tt.Title,
				// (12) amp boilerplate (first style amp-boilerplate, then noscript)
				tt.StyleAMP4AdsBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Preserves style sheet ordering",
			Input: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.LinkStylesheetGoogleFont, tt.StyleAMPCustom,
				"<link href=another-font rel=stylesheet>",
				tt.LinkCanonical, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.LinkStylesheetGoogleFont, tt.StyleAMPCustom,
				"<link href=another-font rel=stylesheet>",
				tt.LinkCanonical, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "AMP Runtime script is reordered as first script",
			Input: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPAudio, tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPRuntime, tt.ScriptAMPAudio,
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Render delaying scripts before non-render delaying scripts",
			Input: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPRuntime, tt.ScriptAMPAudio, tt.ScriptAMPExperiment,
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPRuntime, tt.ScriptAMPExperiment, tt.ScriptAMPAudio,
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Removes duplicate custom element script",
			Input: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPRuntime, tt.ScriptAMPAudio, tt.ScriptAMPAudio,
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPRuntime, tt.ScriptAMPAudio,
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Sorts custom element scripts",
			Input: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPRuntime, tt.ScriptAMPExperiment, tt.ScriptAMPDynamicCSSClasses,
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPRuntime, tt.ScriptAMPDynamicCSSClasses, tt.ScriptAMPExperiment,
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Removes duplicate custom template script",
			Input: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPRuntime, tt.ScriptAMPMustache, tt.ScriptAMPMustache,
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPRuntime, tt.ScriptAMPMustache,
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Preserves multiple favicons",
			Input: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPRuntime, tt.ScriptAMPAudio,
				tt.LinkFavicon, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPRuntime, tt.ScriptAMPAudio,
				tt.LinkFavicon, tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Case insensitive rel value",
			Input: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPRuntime, tt.ScriptAMPAudio, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				`<link href=https://example.com/favicon.ico rel="Shortcut Icon">`,
				"</head><body></body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport,
				tt.ScriptAMPRuntime, tt.ScriptAMPAudio,
				`<link href=https://example.com/favicon.ico rel="Shortcut Icon">`,
				tt.LinkCanonical, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
	}
	runReorderHeadTestcases(t, tcs)
}

func runReorderHeadTestcases(t *testing.T, tcs []tt.TestCase) {
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
		transformers.ReorderHead(&transformers.Context{DOM: inputDOM})

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
			t.Errorf("%s: ReorderHead=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
