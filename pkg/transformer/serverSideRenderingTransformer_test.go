package transformer_test

import (
	"strings"
	"testing"

	"github.com/ampproject/amppackager/pkg/transformer"
	"golang.org/x/net/html"

	tt "github.com/ampproject/amppackager/internal/pkg/testing"
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
			"Modifies document only once",
			// The expected output is actually not correctly server-side
			// rendered, but the presence of i-amphtml-layout attribute halts
			// processing, so this is effectively a no-op.
			"<html i-amphtml-layout><body><amp-img layout=container></amp-img></body>",
			"<html i-amphtml-layout><body><amp-img layout=container></amp-img></body>",
		},
		{
			"Preserves noscript in body",
			"<body><noscript><img src=lemur.png></noscript></body>",
			`<html i-amphtml-layout="" i-amphtml-no-boilerplate=""><head><style amp-runtime=""></style></head><body><noscript><img src=lemur.png></noscript></body></html>`,
		},
		{
			"No changes within template tag",
			"<body><template><amp-img height=42 layout=responsive width=42></amp-img></template></body>",
			`<html i-amphtml-layout="" i-amphtml-no-boilerplate=""><head><style amp-runtime=""></style></head><body><template><amp-img height="42" layout="responsive" width="42"></amp-img></template></body></html>`,
		},
		{
			"Boilerplate removed and layout applied",
			tt.Concat("<!doctype html><html ⚡><head>",
				tt.ScriptAMPRuntime, tt.LinkFavicon, tt.StyleAMPBoilerplate,
				tt.NoscriptAMPBoilerplate, "</head>",
				"<body><amp-img layout=container></amp-img></body></html>"),
			tt.Concat(`<!doctype html><html ⚡ i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				`<style amp-runtime=""></style>`,
				tt.ScriptAMPRuntime, tt.LinkFavicon,
				"</head><body>",
				`<amp-img layout="container" class="i-amphtml-layout-container" i-amphtml-layout="container"></amp-img>`,
				"</body></html>"),
		},
		{
			"Amp4Email Boilerplate removed and layout applied",
			tt.Concat("<!doctype html><html ⚡4email><head>",
				tt.ScriptAMPRuntime, tt.StyleAMP4EmailBoilerplate,
				tt.MetaCharset, "</head>",
				"<body><amp-img layout=container></amp-img></body></html>"),
			tt.Concat(`<!doctype html><html ⚡4email i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				`<style amp-runtime=""></style>`,
				tt.ScriptAMPRuntime, tt.MetaCharset,
				"</head><body>",
				`<amp-img layout="container" class="i-amphtml-layout-container" i-amphtml-layout="container"></amp-img>`,
				"</body></html>"),
		},
		{
			"Amp4Ads Boilerplate removed and layout applied",
			tt.Concat("<!doctype html><html ⚡><head>",
				tt.ScriptAMPRuntime, tt.LinkFavicon, tt.StyleAMP4AdsBoilerplate,
				tt.NoscriptAMPBoilerplate, "</head>",
				"<body><amp-img layout=container></amp-img></body></html>"),
			tt.Concat(`<!doctype html><html ⚡ i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
				`<style amp-runtime=""></style>`,
				tt.ScriptAMPRuntime, tt.LinkFavicon,
				"</head><body>",
				`<amp-img layout="container" class="i-amphtml-layout-container" i-amphtml-layout="container"></amp-img>`,
				"</body></html>"),
		},
		{
			"Boilerplate removed despite sizes (in head tho)",
			tt.Concat("<!doctype html><html ⚡><head>",
				`<link rel="shortcut icon" type="a" href="b" sizes="c">`,
				tt.ScriptAMPRuntime, tt.LinkFavicon, tt.StyleAMP4AdsBoilerplate,
				tt.NoscriptAMPBoilerplate, "</head>",
				"<body><amp-img layout=container></amp-img></body></html>"),
			tt.Concat(`<!doctype html><html ⚡ i-amphtml-layout="" i-amphtml-no-boilerplate=""><head>`,
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
			"amp-audio",
			input("", "<amp-audio></amp-audio>"),
			expected("", "<amp-audio></amp-audio>"),
		},
		{
			"amp-experiment",
			input("", "<amp-experiment></amp-experiment>"),
			expected("", `<amp-experiment class="i-amphtml-layout-container" i-amphtml-layout="container"></amp-experiment>`),
		},
		{
			"amp-story",
			input(tt.ScriptAMPStory, ""),
			expected(tt.ScriptAMPStory, ""),
		},
		{
			"amp-dynamic-css-classes",
			input(tt.ScriptAMPDynamicCSSClasses, ""),
			expected(tt.ScriptAMPDynamicCSSClasses, ""),
		},
		{
			"heights attr",
			input("", `<amp-img height=256 heights="(min-width:500px) 200px, 80%" layout=responsive width=320></amp-img>`),
			expected("", `<amp-img height=256 heights="(min-width:500px) 200px, 80%" layout="responsive" width="320" class="i-amphtml-layout-responsive i-amphtml-layout-size-defined" i-amphtml-layout="responsive"><i-amphtml-sizer style="display:block;padding-top:80.0000%;"></i-amphtml-sizer></amp-img>`),
		},
		{
			"media attr",
			input("", `<amp-img height=355 layout=fixed media="(min-width: 650px) and handheld" src=wide.jpg width=466></amp-img>`),
			expected("", `<amp-img height="355" layout="fixed" media="(min-width: 650px) and handheld" src="wide.jpg" width="466" class="i-amphtml-layout-fixed i-amphtml-layout-size-defined" style="width:466px;height:355px;" i-amphtml-layout="fixed"></amp-img>`),
		},
		{
			"sizes attr",
			input("", `<amp-img height=300 layout=responsive sizes="(min-width: 320px) 320px, 100vw" src=https://acme.org/image1.png width=400></amp-img>`),
			expected("", `<amp-img height=300 layout=responsive sizes="(min-width: 320px) 320px, 100vw" src=https://acme.org/image1.png width=400 class="i-amphtml-layout-responsive i-amphtml-layout-size-defined" i-amphtml-layout="responsive"><i-amphtml-sizer style="display:block;padding-top:75.0000%;"></i-amphtml-sizer></amp-img>`),
		},
		{
			"style attr",
			input("", `<amp-img height=300 layout=fixed src=https://acme.org/image1.png style=position:relative width=400></amp-img>`),
			expected("", `<amp-img height=300 layout=fixed src=https://acme.org/image1.png style=position:relative width=400></amp-img>`),
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
		transformer.ServerSideRenderingTransformer(&transformer.Engine{Doc: inputDoc})

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
		err = html.Render(&expected, expectedDoc)
		if err != nil {
			t.Errorf("%s: html.Render for %s failed %q", tc.Desc, tc.Expected, err)
			continue
		}
		if input.String() != expected.String() {
			t.Errorf("%s: ServerSideRenderingTransformer=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
