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

func TestStrip(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc: "strips comments",
			Input: tt.Concat("<!-- comment -->",
				tt.BuildHTML("<foo><!-- comment --></foo>")),
			Expected: tt.BuildHTML("<foo></foo>"),
		},
		{
			Desc:     "strip duplicate attributes",
			Input:    tt.BuildHTML("<a class=foo class=foo></a>"),
			Expected: tt.BuildHTML("<a class=foo></a>"),
		},
		{
			Desc:     "verify first attr is kept",
			Input:    tt.BuildHTML("<a class=bar href='#' class=foo></a>"),
			Expected: tt.BuildHTML("<a class=bar href='#'></a>"),
		},
		{
			Desc:     "dedupe attr, case-insensitive",
			Input:    tt.BuildHTML("<a CLASS=foo class=foo></a>"),
			Expected: tt.BuildHTML("<a class=foo></a>"),
		},
		{
			Desc:     "dedupe attr, case-insensitive, order irrelevant",
			Input:    tt.BuildHTML("<a class=foo CLASS=bar></a>"),
			Expected: tt.BuildHTML("<a class=foo></a>"),
		},
		{
			Desc: "Strips child whitespace nodes from <html> and <head>",
			Input: tt.Concat(
				"<!doctype html><html ⚡>  <head>\n",
				"\t\t",
				tt.ScriptAMPRuntime,
				"  ",
				tt.LinkFavicon,
				"</head>\n<body>\n",
				"    foo<b> </b>bar\n\n",
				"</body></html>"),
			Expected: tt.BuildHTML("\n    foo<b> </b>bar\n\n"),
		},
		{
			// Stray text in head will automatically start a body tag, (and will
			// NOT be stripped because it's not all whitespace). Note also that
			// all subsequent tags after the stray text are moved to body too.
			Desc: "strip stray text in head",
			Input: `<!doctype html>
<html ⚡>
  <head>
    <meta charset="utf-8">
    <link ref=canonical href=http://www.example.com>stray text
    <script async src="https://cdn.ampproject.org/v0.js"></script>
  </head>
  <body class="foo">
</body>
</html>`,
			Expected: tt.Concat(`<!DOCTYPE html><html ⚡=""><head><meta charset="utf-8"/><link ref="canonical" href="http://www.example.com"/></head><body class="foo">stray text
`,
				`    <script async="" src="https://cdn.ampproject.org/v0.js"></script>`,
				"\n  ",
				"\n  ",
				"\n\n</body></html>"),
		},
		{
			Desc:     "Strip empty amp-custom style",
			Input:    "<style amp-custom></style>",
			Expected: "",
		},
		{
			Desc:     "No-op (not empty amp-custom style)",
			Input:    "<style amp-custom>amp-gist { color: red; }</style>",
			Expected: "<style amp-custom>amp-gist { color: red; }</style>",
		},
		{
			Desc:     "strip extra attrs from style amp-custom",
			Input:    "<style amp-custom=amp-custom type=text/css>amp-gist { color: red; }</style>",
			Expected: "<style amp-custom>amp-gist { color: red; }</style>",
		},
		{
			Desc: "Sanitize URIs in src",
			// (src has space, space, and tab)
			Input: `<img src="  	">`,
			Expected: `<img src="  "/>`,
		},
		{
			Desc: "Sanitize URIs in href",
			// (href has space, space, and tab)
			Input: `<a href="  	">`,
			Expected: `<a href="  "/>`,
		},
		{
			Desc: "untouched URI",
			Input: `<lemur uri="  	">`,
			Expected: `<lemur uri="  	">`,
		},
		{
			Desc:     "Strip extra <title> elements",
			Input:    `<!doctype html><html ⚡>  <head><title>a</title><title>b</title></head>`,
			Expected: `<!doctype html><html ⚡>  <head><title>a</title></head>`,
		},
		{
			Desc:     "Strip all <title> elements in body",
			Input:    `<!doctype html><html ⚡><body><title>a</title><title>b</title></body>`,
			Expected: `<!doctype html><html ⚡><body></body>`,
		},
		{
			Desc: "Preserve svg <title> elements",
			Input: tt.Concat("<!doctype html><html ⚡><body>",
				"<svg><title>a</title></svg>",
				"<svg><symbol><title>b</title></symbol></svg>",
				"</body>"),
			Expected: tt.Concat("<!doctype html><html ⚡><body>",
				"<svg><title>a</title></svg>",
				"<svg><symbol><title>b</title></symbol></svg>",
				"</body>"),
		},
	}
	runAllTestCases(t, tcs)
}

func TestDoctype(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc:     "doctype no-op",
			Input:    "<!doctype html>",
			Expected: "<!doctype html>",
		},
		{
			Desc:     "doctype add html",
			Input:    "<!doctype>",
			Expected: "<!doctype html>",
		},
		{
			Desc:     "doctype strip all",
			Input:    `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">`,
			Expected: "<!doctype html>",
		},
		{
			Desc:     "doctype strip bogus",
			Input:    `<!DOCTYPE HTML PUBLIC "bogus" "notreal">`,
			Expected: "<!doctype html>",
		},
		{
			Desc:     "doctype ignore non-html",
			Input:    `<!DOCTYPE document SYSTEM "subjects.dtd">`,
			Expected: "<!doctype html>",
		},
	}
	runAllTestCases(t, tcs)
}

func TestWellFormedHtml(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc: "wellformed",
			Input: tt.Concat("<!doctype html><html ⚡>",
				tt.ScriptAMPRuntime,
				tt.LinkFavicon,
				"<foo>"),
			Expected: `<!DOCTYPE html><html ⚡=""><head><script async="" src="https://cdn.ampproject.org/v0.js"></script><link href="https://example.com/favicon.ico" rel="icon"/></head><body><foo></foo></body></html>`,
		},
	}
	runAllTestCases(t, tcs)
}

func TestNonceRemoved(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc:     "remove nonce",
			Input:    "<script nonce async>",
			Expected: "<script async>",
		},
	}
	runAllTestCases(t, tcs)
}

func TestNoScriptParsed(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc:     "parse noscript",
			Input:    "<body><noscript><lemur z b y></noscript></body>",
			Expected: `<body><noscript><lemur z="" b="" y=""></lemur></noscript></body>`,
		},
	}
	runAllTestCases(t, tcs)
}

func runAllTestCases(t *testing.T, tcs []tt.TestCase) {
	for _, tc := range tcs {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s: html.Parse failed %q", tc.Input, err)
			continue
		}
		engine := transformers.Engine{Doc: inputDoc}
		engine.Transform()
		var input strings.Builder
		if err := html.Render(&input, engine.Doc); err != nil {
			t.Errorf("%s: html.Render failed %q", tc.Input, err)
			continue
		}

		expectedDoc, err := html.Parse(strings.NewReader(tc.Expected))
		if err != nil {
			t.Errorf("%s: html.Parse failed %q", tc.Expected, err)
			continue
		}
		var expected strings.Builder
		err = html.Render(&expected, expectedDoc)
		if err != nil {
			t.Errorf("%s: html.Render failed %q", tc.Expected, err)
			continue
		}
		if input.String() != expected.String() {
			t.Errorf("%s: Transform=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
