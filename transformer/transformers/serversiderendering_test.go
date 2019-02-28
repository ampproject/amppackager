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

func TestServerSideRendering(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc: "Modifies document only once",
			// The expected output is actually not correctly server-side
			// rendered, but the presence of i-amphtml-layout attribute halts
			// processing, so this is effectively a no-op.
			Input:    tt.Concat(tt.Doctype,
				"<html ⚡ i-amphtml-layout><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime, tt.LinkFavicon,
				tt.LinkCanonical, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<amp-img layout=container></amp-img>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype,
				"<html ⚡ i-amphtml-layout><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime, tt.LinkFavicon,
				tt.LinkCanonical, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<amp-img layout=container></amp-img>",
				"</body></html>"),
		},
		{
			Desc:     "Boilerplate removed and preserves noscript in body",
			Input:    tt.Concat(tt.Doctype,
				"<html  ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime, tt.LinkFavicon,
				tt.LinkCanonical, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body><noscript><img src=lemur.png></noscript></body></html>"),
			Expected: tt.Concat(tt.Doctype,
				`<html ⚡ i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.LinkCanonical,
				"</head><body>",
				"<noscript><img src=lemur.png></noscript>",
				"</body></html>"),
		},
		{
			Desc:     "Boilerplate removed and no changes within template tag",
			Input:    tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime, tt.LinkFavicon,
				tt.LinkCanonical, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<template><amp-img height=42 layout=responsive width=42></amp-img></template>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype,
				`<html ⚡ i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.LinkCanonical,
				"</head><body>",
				`<template><amp-img height="42" layout="responsive" width="42"></amp-img></template>`,
				"</body></html>"),
		},
		{
			Desc:     "Boilerplate removed and layout applied",
			Input:    tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime, tt.LinkFavicon,
        tt.LinkCanonical, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				`<amp-img class="" layout=container></amp-img>`,
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype,
				`<html ⚡ i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.LinkCanonical,
				"</head><body>",
				`<amp-img class="i-amphtml-layout-container" layout="container" i-amphtml-layout="container"></amp-img>`,
				"</body></html>"),
		},
		{
			Desc:     "Amp4Email Boilerplate removed and layout applied",
			Input:    tt.Concat(tt.Doctype,
				"<html ⚡4email><head>",
				tt.MetaCharset, tt.ScriptAMPRuntime, tt.StyleAMP4EmailBoilerplate,
				"</head><body>",
				"<amp-img layout=container></amp-img>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype,
				`<html ⚡4email i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				tt.MetaCharset, tt.ScriptAMPRuntime,
				"</head><body>",
				`<amp-img layout="container" class="i-amphtml-layout-container" i-amphtml-layout="container"></amp-img>`,
				"</body></html>"),
		},
		{
			Desc:     "Amp4Ads Boilerplate removed and layout applied",
			Input:    tt.Concat(tt.Doctype,
				"<html ⚡4ads><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.StyleAMP4AdsBoilerplate,
				"</head><body>",
				"<amp-img layout=container></amp-img>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype,
				`<html ⚡4ads i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				"</head><body>",
				`<amp-img layout="container" class="i-amphtml-layout-container" i-amphtml-layout="container"></amp-img>`,
				"</body></html>"),
		},
		{
			Desc:     "Boilerplate removed despite sizes (in head tho)",
			Input:    tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime, tt.LinkFavicon,
				`<link rel="shortcut icon" type="a" href="b" sizes="c">`,
        tt.LinkCanonical, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<amp-img layout=container></amp-img>",
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype,
				`<html ⚡ i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime, tt.LinkFavicon,
				`<link rel="shortcut icon" type="a" href="b" sizes="c">`,
				tt.LinkCanonical,
				"</head><body>",
				`<amp-img layout="container" class="i-amphtml-layout-container" i-amphtml-layout="container"></amp-img>`,
				"</body></html>"),
		},
		{
      Desc:     "Boilerplate removed when amp-experiment is present but empty",
      Input:    tt.Concat(tt.Doctype,
				"<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime, tt.LinkFavicon,
				tt.LinkCanonical, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
        "</head><body>",
				`<amp-experiment><script type="application/json">{ }</script></amp-experiment>`,
				"</body></html>"),
      Expected: tt.Concat(tt.Doctype,
				`<html ⚡ i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.LinkCanonical,
        "</head><body>",
				`<amp-experiment class=i-amphtml-layout-container i-amphtml-layout=container><script type="application/json">{ }</script></amp-experiment>`,
        "</body></html>"),
    },

	}
	runServerSideRenderingTestcases(t, tcs)
}

func TestBoilerplatePreserved(t *testing.T) {
	input := func(extrahead, body string) string {
		return tt.Concat(tt.Doctype,
			"<html ⚡><head>",
			tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
			extrahead,
			tt.LinkFavicon, tt.LinkCanonical,
			tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
			"</head><body>", body, "</body></html>")
	}
	expected := func(extrahead, body string) string {
		return tt.Concat(tt.Doctype,
			`<html ⚡ i-amphtml-layout=""><head>`,
			tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
			extrahead,
			tt.LinkFavicon, tt.LinkCanonical,
			tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
			"</head><body>", body, "</body></html>")
	}

	tcs := []tt.TestCase{
		{
			Desc:     "amp-audio",
			Input:    input("", "<amp-audio></amp-audio>"),
			Expected: expected("", "<amp-audio></amp-audio>"),
		},
		{
			Desc:     "amp-experiment is non-empty",
			Input:    input("", `<amp-experiment><script type="application/json">{ "exp": { "variants": { "a": 25, "b": 25 } } }</script></amp-experiment>`),
			Expected: expected("", `<amp-experiment class="i-amphtml-layout-container" i-amphtml-layout="container"><script type="application/json">{ "exp": { "variants": { "a": 25, "b": 25 } } }</script></amp-experiment>`),
		},
		{
			Desc:     "amp-story",
			Input:    input(tt.ScriptAMPStory, ""),
			Expected: expected(tt.ScriptAMPStory, ""),
		},
		{
			Desc:     "amp-dynamic-css-classes",
			Input:    input(tt.ScriptAMPDynamicCSSClasses, ""),
			Expected: expected(tt.ScriptAMPDynamicCSSClasses, ""),
		},
		{
			Desc:     "heights attr",
			Input:    input("", `<amp-img height=256 heights="(min-width:500px) 200px, 80%" layout=responsive width=320></amp-img>`),
			Expected: expected("", `<amp-img height=256 heights="(min-width:500px) 200px, 80%" layout="responsive" width="320" class="i-amphtml-layout-responsive i-amphtml-layout-size-defined" i-amphtml-layout="responsive"><i-amphtml-sizer style="display:block;padding-top:80.0000%;"></i-amphtml-sizer></amp-img>`),
		},
		{
			Desc:     "media attr",
			Input:    input("", `<amp-img height=355 layout=fixed media="(min-width: 650px) and handheld" src=wide.jpg width=466></amp-img>`),
			Expected: expected("", `<amp-img height="355" layout="fixed" media="(min-width: 650px) and handheld" src="wide.jpg" width="466" class="i-amphtml-layout-fixed i-amphtml-layout-size-defined" style="width:466px;height:355px;" i-amphtml-layout="fixed"></amp-img>`),
		},
		{
			Desc:     "sizes attr",
			Input:    input("", `<amp-img height=300 layout=responsive sizes="(min-width: 320px) 320px, 100vw" src=https://acme.org/image1.png width=400></amp-img>`),
			Expected: expected("", `<amp-img height=300 layout=responsive sizes="(min-width: 320px) 320px, 100vw" src=https://acme.org/image1.png width=400 class="i-amphtml-layout-responsive i-amphtml-layout-size-defined" i-amphtml-layout="responsive"><i-amphtml-sizer style="display:block;padding-top:75.0000%;"></i-amphtml-sizer></amp-img>`),
		},
	}
	runServerSideRenderingTestcases(t, tcs)
}

func runServerSideRenderingTestcases(t *testing.T, tcs []tt.TestCase) {
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
		transformers.ServerSideRendering(&transformers.Context{DOM: inputDOM})

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
			t.Errorf("%s: ServerSideRendering=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
