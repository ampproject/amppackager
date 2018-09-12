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
// renderer (otherwise the scope of these tests would creep beyond unit
// testing). Therefore, the test data must be made to match, and is not
// the expected normalized output from transformer.go, nor from any other
// transformers.

func TestReorderHeadTransformer(t *testing.T) {
	testCases := []tt.TestCase{
		{
			Desc: "Reorders head children for AMP document",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.Title, tt.StyleAMPBoilerplate,
				tt.ScriptAMPExperiment, tt.ScriptAMPAudio,
				tt.NoscriptAMPBoilerplate, tt.StyleAMPRuntime,
				tt.ScriptAMPRuntime, tt.LinkStylesheetGoogleFont,
				tt.LinkGoogleFontPreconnect, tt.MetaCharset,
				tt.MetaViewport, tt.StyleAMPCustom,
				tt.LinkFavicon, tt.ScriptAMPViewerRuntime,
				"</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡><head>",
				// (0) <meta charset> tag
				tt.MetaCharset,
				// (1) <style amp-runtime> (inserted by ServerSideRenderingTransformer)
				tt.StyleAMPRuntime,
				// (2) remaining <meta> tags
				tt.MetaViewport,
				// (3) AMP runtime .js <script> tag
				tt.ScriptAMPRuntime,
				// (4) AMP viewer runtime .js <script> tag
				// (inserted by AmpViewerScriptTransformer)
				tt.ScriptAMPViewerRuntime,
				// (5) <script> tags that are render delaying
				tt.ScriptAMPExperiment,
				// (6) remaining <script> tags
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
				// (12) amp boilerplate (first style amp-boilerplate, then noscript)
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Reorders head children for AMP4Ads document",
			Input: tt.Concat("<!doctype html><html ⚡4ads><head>",
				tt.Title, tt.StyleAMP4AdsBoilerplate, tt.ScriptAMPAudio,
				tt.ScriptAMP4AdsRuntime, tt.LinkStylesheetGoogleFont,
				tt.LinkGoogleFontPreconnect, tt.MetaCharset, tt.MetaViewport, tt.StyleAMPCustom,
				"</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡4ads><head>",
				// (0) <meta charset> tag
				tt.MetaCharset,
				// (1) <style amp-runtime> (inserted by ServerSideRenderingTransformer)
				// n/a for AMP4Ads, no ServerSideRenderingTransformer
				// (2) remaining <meta> tags
				tt.MetaViewport,
				// (3) AMP runtime .js <script> tag
				tt.ScriptAMP4AdsRuntime,
				// (4) <script> tags that are render delaying
				// n/a for AMP4Ads, no render delaying <script> tags allowed
				// (5) remaining <script> tags
				tt.ScriptAMPAudio,
				// (6) <link> tag for favicons
				// n/a for AMP4Ads, no favicons allowed
				// (7) <link> tag for resource hints
				tt.LinkGoogleFontPreconnect,
				// (8) <link rel=stylesheet> tags before <style amp-custom>
				tt.LinkStylesheetGoogleFont,
				// (9) <style amp-custom>
				tt.StyleAMPCustom,
				// (10) any other tags allowed in <head>
				tt.Title,
				// (11) amp boilerplate (first style amp-boilerplate, then noscript)
				tt.StyleAMP4AdsBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Preserves style sheet ordering",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.LinkFavicon,
				tt.LinkStylesheetGoogleFont, tt.StyleAMPCustom,
				"<link href=another-font rel=stylesheet>",
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.LinkFavicon,
				tt.LinkStylesheetGoogleFont, tt.StyleAMPCustom,
				"<link href=another-font rel=stylesheet>",
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "AMP Runtime script is reordered as first script",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPAudio, tt.ScriptAMPRuntime, tt.LinkFavicon,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.ScriptAMPAudio, tt.LinkFavicon,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Render delaying scripts before non-render delaying scripts",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.ScriptAMPAudio,
				tt.ScriptAMPExperiment, tt.LinkFavicon,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.ScriptAMPExperiment,
				tt.ScriptAMPAudio, tt.LinkFavicon,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Removes duplicate custom element script",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.ScriptAMPAudio,
				tt.ScriptAMPAudio, tt.LinkFavicon,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.ScriptAMPAudio, tt.LinkFavicon,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Sorts custom element scripts",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.ScriptAMPExperiment,
				tt.ScriptAMPDynamicCSSClasses, tt.LinkFavicon,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.ScriptAMPDynamicCSSClasses,
				tt.ScriptAMPExperiment, tt.LinkFavicon,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Removes duplicate custom template script",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.ScriptAMPMustache,
				tt.ScriptAMPMustache, tt.LinkFavicon,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.ScriptAMPMustache,
				tt.LinkFavicon, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Preserves multiple favicons",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.ScriptAMPAudio,
				tt.LinkFavicon, tt.LinkFavicon,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.ScriptAMPAudio,
				tt.LinkFavicon, tt.LinkFavicon,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
		{
			Desc: "Case insensitive rel value",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.ScriptAMPAudio,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				`<link href=https://example.com/favicon.ico rel="Shortcut Icon">`,
				"</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.ScriptAMPAudio,
				`<link href=https://example.com/favicon.ico rel="Shortcut Icon">`,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
		},
	}
	runReorderHeadTransformerTestcases(t, testCases)
}

func runReorderHeadTransformerTestcases(t *testing.T, testCases []tt.TestCase) {
	for _, tc := range testCases {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s: html.Parse for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		transformers.ReorderHeadTransformer(&transformers.Engine{Doc: inputDoc})

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
			t.Errorf("%s: ReorderHeadTransformer=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
