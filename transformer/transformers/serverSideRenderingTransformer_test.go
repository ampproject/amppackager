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

func TestServerSideRenderingTransformer(t *testing.T) {
	testCases := []tt.TestCase{
		{
			Desc: "Modifies document only once",
			// The expected output is actually not correctly server-side
			// rendered, but the presence of i-amphtml-layout attribute halts
			// processing, so this is effectively a no-op.
			Input:    "<html i-amphtml-layout><body><amp-img layout=container></amp-img></body>",
			Expected: "<html i-amphtml-layout><body><amp-img layout=container></amp-img></body>",
		},
		{
			Desc:     "Preserves noscript in body",
			Input:    "<body><noscript><img src=lemur.png></noscript></body>",
			Expected: `<html i-amphtml-layout="" i-amphtml-no-boilerplate=""><head><style amp-runtime=""></style></head><body><noscript><img src=lemur.png></noscript></body></html>`,
		},
		{
			Desc:     "No changes within template tag",
			Input:    "<body><template><amp-img height=42 layout=responsive width=42></amp-img></template></body>",
			Expected: `<html i-amphtml-layout="" i-amphtml-no-boilerplate=""><head><style amp-runtime=""></style></head><body><template><amp-img height="42" layout="responsive" width="42"></amp-img></template></body></html>`,
		},
		{
			Desc: "Boilerplate removed and layout applied",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.ScriptAMPRuntime, tt.LinkFavicon, tt.StyleAMPBoilerplate,
				tt.NoscriptAMPBoilerplate, "</head>",
				"<body><amp-img layout=container></amp-img></body></html>"),
			Expected: tt.Concat(`<!doctype html><html ⚡ i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				`<style amp-runtime=""></style>`,
				tt.ScriptAMPRuntime, tt.LinkFavicon,
				"</head><body>",
				`<amp-img layout="container" class="i-amphtml-layout-container" i-amphtml-layout="container"></amp-img>`,
				"</body></html>"),
		},
		{
			Desc: "Amp4Email Boilerplate removed and layout applied",
			Input: tt.Concat("<!doctype html><html ⚡4email><head>",
				tt.ScriptAMPRuntime, tt.StyleAMP4EmailBoilerplate,
				tt.MetaCharset, "</head>",
				"<body><amp-img layout=container></amp-img></body></html>"),
			Expected: tt.Concat(`<!doctype html><html ⚡4email i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				`<style amp-runtime=""></style>`,
				tt.ScriptAMPRuntime, tt.MetaCharset,
				"</head><body>",
				`<amp-img layout="container" class="i-amphtml-layout-container" i-amphtml-layout="container"></amp-img>`,
				"</body></html>"),
		},
		{
			Desc: "Amp4Ads Boilerplate removed and layout applied",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.ScriptAMPRuntime, tt.LinkFavicon, tt.StyleAMP4AdsBoilerplate,
				tt.NoscriptAMPBoilerplate, "</head>",
				"<body><amp-img layout=container></amp-img></body></html>"),
			Expected: tt.Concat(`<!doctype html><html ⚡ i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				`<style amp-runtime=""></style>`,
				tt.ScriptAMPRuntime, tt.LinkFavicon,
				"</head><body>",
				`<amp-img layout="container" class="i-amphtml-layout-container" i-amphtml-layout="container"></amp-img>`,
				"</body></html>"),
		},
		{
			Desc: "Boilerplate removed despite sizes (in head tho)",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				`<link rel="shortcut icon" type="a" href="b" sizes="c">`,
				tt.ScriptAMPRuntime, tt.LinkFavicon, tt.StyleAMP4AdsBoilerplate,
				tt.NoscriptAMPBoilerplate, "</head>",
				"<body><amp-img layout=container></amp-img></body></html>"),
			Expected: tt.Concat(`<!doctype html><html ⚡ i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				`<style amp-runtime=""></style>`,
				`<link rel="shortcut icon" type="a" href="b" sizes="c">`,
				tt.ScriptAMPRuntime, tt.LinkFavicon,
				"</head><body>",
				`<amp-img layout="container" class="i-amphtml-layout-container" i-amphtml-layout="container"></amp-img>`,
				"</body></html>"),
		},
	}
	runServerSideRenderingTransformerTestcases(t, testCases)
}

func TestBoilerplatePreserved(t *testing.T) {
	input := func(extrahead, body string) string {
		return tt.Concat("<!doctype html><html ⚡><head>",
			tt.ScriptAMPRuntime, tt.LinkFavicon, tt.StyleAMPBoilerplate,
			tt.NoscriptAMPBoilerplate, extrahead, "</head><body>",
			body, "</body></html>")
	}
	expected := func(extrahead, body string) string {
		return tt.Concat(`<!doctype html><html ⚡ i-amphtml-layout=""><head>`,
			`<style amp-runtime=""></style>`,
			tt.ScriptAMPRuntime, tt.LinkFavicon, tt.StyleAMPBoilerplate,
			tt.NoscriptAMPBoilerplate, extrahead, "</head><body>",
			body, "</body></html>")
	}

	testCases := []tt.TestCase{
		{
			Desc:     "amp-audio",
			Input:    input("", "<amp-audio></amp-audio>"),
			Expected: expected("", "<amp-audio></amp-audio>"),
		},
		{
			Desc:     "amp-experiment",
			Input:    input("", "<amp-experiment></amp-experiment>"),
			Expected: expected("", `<amp-experiment class="i-amphtml-layout-container" i-amphtml-layout="container"></amp-experiment>`),
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
		{
			Desc:     "style attr",
			Input:    input("", `<amp-img height=300 layout=fixed src=https://acme.org/image1.png style=position:relative width=400></amp-img>`),
			Expected: expected("", `<amp-img height=300 layout=fixed src=https://acme.org/image1.png style=position:relative width=400></amp-img>`),
		},
	}
	runServerSideRenderingTransformerTestcases(t, testCases)
}

func runServerSideRenderingTransformerTestcases(t *testing.T, testCases []tt.TestCase) {
	for _, tc := range testCases {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s: html.Parse for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		transformers.ServerSideRenderingTransformer(&transformers.Engine{Doc: inputDoc})

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
			t.Errorf("%s: ServerSideRenderingTransformer=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
