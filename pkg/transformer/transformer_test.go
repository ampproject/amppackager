package transformer_test

import (
	"strings"
	"testing"

	"github.com/ampproject/amppackager/pkg/transformer"
	"golang.org/x/net/html"

	tt "github.com/ampproject/amppackager/internal/pkg/testing"
)

// testCase stores the input HTML, expected output HTML, and an optional
// transformer to execute.
type transformerTestCase struct {
	desc        string
	input       string
	expected    string
	transformer func(*transformer.Engine)
}

func TestStrip(t *testing.T) {
	tcs := []transformerTestCase{
		{
			"strips comments",
			tt.Concat("<!-- comment -->",
				tt.BuildHTML("<foo><!-- comment --></foo>")),
			tt.BuildHTML("<foo></foo>"),
			nil,
		},
		{
			"strip duplicate attributes",
			tt.BuildHTML("<a class=foo class=foo></a>"),
			tt.BuildHTML("<a class=foo></a>"),
			nil,
		},
		{
			"verify first attr is kept",
			tt.BuildHTML("<a class=bar href='#' class=foo></a>"),
			tt.BuildHTML("<a class=bar href='#'></a>"),
			nil,
		},
		{
			"dedupe attr, case-insensitive",
			tt.BuildHTML("<a CLASS=foo class=foo></a>"),
			tt.BuildHTML("<a class=foo></a>"),
			nil,
		},
		{
			"dedupe attr, case-insensitive, order irrelevant",
			tt.BuildHTML("<a class=foo CLASS=bar></a>"),
			tt.BuildHTML("<a class=foo></a>"),
			nil,
		},
		{
			"Strips child whitespace nodes from <html> and <head>",
			tt.Concat(
				"<!doctype html><html ⚡>  <head>\n",
				"\t\t",
				tt.ScriptAMPRuntime,
				"  ",
				tt.LinkFavicon,
				"</head>\n<body>\n",
				"    foo<b> </b>bar\n\n",
				"</body></html>"),
			tt.BuildHTML("\n    foo<b> </b>bar\n\n"),
			nil,
		},
		{
			// Stray text in head will automatically start a body tag, (and will
			// NOT be stripped because it's not all whitespace). Note also that
			// all subsequent tags after the stray text are moved to body too.
			"strip stray text in head",
			`<!doctype html>
<html ⚡>
  <head>
    <meta charset="utf-8">
    <link ref=canonical href=http://www.example.com>stray text
    <script async src="https://cdn.ampproject.org/v0.js"></script>
  </head>
  <body class="foo">
</body>
</html>`,
			tt.Concat(`<!DOCTYPE html><html ⚡=""><head><meta charset="utf-8"/><link ref="canonical" href="http://www.example.com"/></head><body class="foo">stray text
`,
				`    <script async="" src="https://cdn.ampproject.org/v0.js"></script>`,
				"\n  ",
				"\n  ",
				"\n\n</body></html>"),
			nil,
		},
		{
			"Strip empty amp-custom style",
			"<style amp-custom></style>",
			"",
			nil,
		},
		{
			"No-op (not empty amp-custom style)",
			"<style amp-custom>amp-gist { color: red; }</style>",
			"<style amp-custom>amp-gist { color: red; }</style>",
			nil,
		},
		{
			"strip extra attrs from style amp-custom",
			"<style amp-custom=amp-custom type=text/css>amp-gist { color: red; }</style>",
			"<style amp-custom>amp-gist { color: red; }</style>",
			nil,
		},
		{
			"Sanitize URIs in src",
			// (src has space, space, and tab)
			`<img src="  	">`,
			`<img src="  "/>`,
			nil,
		},
		{
			"Sanitize URIs in href",
			// (href has space, space, and tab)
			`<a href="  	">`,
			`<a href="  "/>`,
			nil,
		},
		{
			"untouched URI",
			`<lemur uri="  	">`,
			`<lemur uri="  	">`,
			nil,
		},
		{
			"Strip extra <title> elements",
			`<!doctype html><html ⚡>  <head><title>a</title><title>b</title></head>`,
			`<!doctype html><html ⚡>  <head><title>a</title></head>`,
			nil,
		},
		{
			"Strip all <title> elements in body",
			`<!doctype html><html ⚡><body><title>a</title><title>b</title></body>`,
			`<!doctype html><html ⚡><body></body>`,
			nil,
		},
		{
			"Preserve svg <title> elements",
			tt.Concat("<!doctype html><html ⚡><body>",
				"<svg><title>a</title></svg>",
				"<svg><symbol><title>b</title></symbol></svg>",
				"</body>"),
			tt.Concat("<!doctype html><html ⚡><body>",
				"<svg><title>a</title></svg>",
				"<svg><symbol><title>b</title></symbol></svg>",
				"</body>"),
			nil,
		},
	}
	runAllTestCases(t, tcs)
}

func TestDoctype(t *testing.T) {
	tcs := []transformerTestCase{
		{
			"doctype no-op",
			"<!doctype html>",
			"<!doctype html>",
			nil,
		},
		{
			"doctype add html",
			"<!doctype>",
			"<!doctype html>",
			nil,
		},
		{
			"doctype strip all",
			`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">`,
			"<!doctype html>",
			nil,
		},
		{
			"doctype strip bogus",
			`<!DOCTYPE HTML PUBLIC "bogus" "notreal">`,
			"<!doctype html>",
			nil,
		},
		{
			"doctype ignore non-html",
			`<!DOCTYPE document SYSTEM "subjects.dtd">`,
			"<!doctype html>",
			nil,
		},
	}
	runAllTestCases(t, tcs)
}

func TestWellFormedHtml(t *testing.T) {
	tcs := []transformerTestCase{
		{
			"wellformed",
			tt.Concat("<!doctype html><html ⚡>",
				tt.ScriptAMPRuntime,
				tt.LinkFavicon,
				"<foo>"),
			`<!DOCTYPE html><html ⚡=""><head><script async="" src="https://cdn.ampproject.org/v0.js"></script><link href="https://example.com/favicon.ico" rel="icon"/></head><body><foo></foo></body></html>`,
			nil,
		},
	}
	runAllTestCases(t, tcs)
}

func TestNonceRemoved(t *testing.T) {
	tcs := []transformerTestCase{
		{
			"remove nonce",
			"<script nonce async>",
			"<script async>",
			nil,
		},
	}
	runAllTestCases(t, tcs)
}

func TestNoScriptParsed(t *testing.T) {
	tcs := []transformerTestCase{
		{
			"parse noscript",
			"<body><noscript><lemur z b y></noscript></body>",
			`<body><noscript><lemur z="" b="" y=""></lemur></noscript></body>`,
			nil,
		},
	}
	runAllTestCases(t, tcs)
}

func runAllTestCases(t *testing.T, tcs []transformerTestCase) {
	for _, tc := range tcs {
		inputDoc, err := html.Parse(strings.NewReader(tc.input))
		if err != nil {
			t.Errorf("%s: html.Parse failed %q", tc.input, err)
			continue
		}
		engine := transformer.Engine{Doc: inputDoc}
		if tc.transformer != nil {
			engine.Transformers = append(engine.Transformers, tc.transformer)
		}
		engine.Transform()
		var input strings.Builder
		if err := html.Render(&input, engine.Doc); err != nil {
			t.Errorf("%s: html.Render failed %q", tc.input, err)
			continue
		}

		expectedDoc, err := html.Parse(strings.NewReader(tc.expected))
		if err != nil {
			t.Errorf("%s: html.Parse failed %q", tc.expected, err)
			continue
		}
		var expected strings.Builder
		err = html.Render(&expected, expectedDoc)
		if err != nil {
			t.Errorf("%s: html.Render failed %q", tc.expected, err)
			continue
		}
		if input.String() != expected.String() {
			t.Errorf("%s: Transform=\n%q\nwant=\n%q", tc.desc, &input, &expected)
		}
	}
}
