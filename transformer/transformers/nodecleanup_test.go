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
	"fmt"
	"strings"
	"testing"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	tt "github.com/ampproject/amppackager/transformer/internal/testing"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

// BuildHTML returns AMPHTML with the given body string. Everything
// in the head tag is the minimum required for a valid AMP document
// except it is missing NoscriptAMPBoilerplate as noscript tags are
// stripped by this transformer (see b/79415817).
func BuildHTML(body string) string {
	return tt.Concat(
		tt.Doctype,
		"<html ⚡><head>",
		tt.MetaCharset,
		tt.MetaViewport,
		tt.ScriptAMPRuntime,
		tt.LinkFavicon,
		tt.LinkCanonical,
		tt.StyleAMPBoilerplate,
		"</head><body>",
		body,
		"</body></html>")
}

func TestNodeCleanup_Strip(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc: "strips comments",
			Input: tt.Concat("<!-- comment -->",
				BuildHTML("<foo><!-- comment --></foo>")),
			Expected: BuildHTML("<foo></foo>"),
		},
		{
			Desc:     "strip duplicate attributes",
			Input:    BuildHTML("<a class=foo class=foo></a>"),
			Expected: BuildHTML("<a class=foo></a>"),
		},
		{
			Desc:     "verify first attr is kept",
			Input:    BuildHTML("<a class=bar href='#' class=foo></a>"),
			Expected: BuildHTML("<a class=bar href='#'></a>"),
		},
		{
			Desc:     "dedupe attr, case-insensitive",
			Input:    BuildHTML("<a CLASS=foo class=foo></a>"),
			Expected: BuildHTML("<a class=foo></a>"),
		},
		{
			Desc:     "dedupe attr, case-insensitive, order irrelevant",
			Input:    BuildHTML("<a class=foo CLASS=bar></a>"),
			Expected: BuildHTML("<a class=foo></a>"),
		},
		{
			Desc: "strips child whitespace nodes from <html> and <head>",
			Input: tt.Concat(
				tt.Doctype, "<html ⚡>  <head>\n",
				"\t\t",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				"  ",
				tt.LinkFavicon, tt.LinkCanonical, tt.StyleAMPBoilerplate,
				"</head>\n<body>\n",
				"    foo<b> </b>bar\n\n",
				"</body></html>"),
			Expected: BuildHTML("\n    foo<b> </b>bar\n\n"),
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
			// Whitespace should be stripped, then empty style stripped too.
			Desc:     "Strip amp-custom style with pure whitespace",
			Input:    "<style amp-custom>  </style>",
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
			// The extra whitespace node after second title element is to
			// prove that siblings after the removed node are still processed.
			// In this case, the whitespace is removed.
			Desc:     "Strip extra <title> elements",
			Input:    `<!doctype html><html ⚡><head><title>a</title><title>b</title> <script></script></head>`,
			Expected: `<!doctype html><html ⚡><head><title>a</title><script></script></head>`,
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
	runNodeCleanupTestCases(t, tcs)
}

func TestNodeCleanup_Doctype(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc:     "doctype no-op",
			Input:    tt.Doctype,
			Expected: tt.Doctype,
		},
		{
			Desc:     "doctype add html",
			Input:    "<!doctype>",
			Expected: tt.Doctype,
		},
		{
			Desc:     "doctype strip all",
			Input:    `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">`,
			Expected: tt.Doctype,
		},
		{
			Desc:     "doctype strip bogus",
			Input:    `<!DOCTYPE HTML PUBLIC "bogus" "notreal">`,
			Expected: tt.Doctype,
		},
		{
			Desc:     "doctype ignore non-html",
			Input:    `<!DOCTYPE document SYSTEM "subjects.dtd">`,
			Expected: tt.Doctype,
		},
	}
	runNodeCleanupTestCases(t, tcs)
}

func TestNodeCleanup_WellFormedHtml(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc: "wellformed",
			Input: tt.Concat(tt.Doctype, "<html ⚡>",
				tt.ScriptAMPRuntime,
				tt.LinkFavicon,
				"<foo>"),
			Expected: `<!DOCTYPE html><html ⚡=""><head><script async="" src="https://cdn.ampproject.org/v0.js"></script><link href="https://example.com/favicon.ico" rel="icon"/></head><body><foo></foo></body></html>`,
		},
	}
	runNodeCleanupTestCases(t, tcs)
}

func TestNodeCleanup_NonceRemoved(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc:     "remove nonce",
			Input:    "<script nonce async>",
			Expected: "<script async>",
		},
	}
	runNodeCleanupTestCases(t, tcs)
}

func TestNodeCleanup_NoScriptRemoved(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc:     "remove noscript in head",
			Input:    "<head><noscript><style></style></noscript></head>",
			Expected: `<head></head>`,
		},
		{
			Desc:     "remove noscript in body",
			Input:    "<body><noscript><lemur z b y></noscript></body>",
			Expected: `<body></body>`,
		},
		{
			Desc:     "imbalanced comment",
			Input:    "<body><noscript><!-- comment </noscript>--></body>",
			Expected: `<body>--&gt;</body>`,
		},
		{
			Desc:     "imbalanced",
			Input:    "<body><noscript><noscript></noscript></body>",
			Expected: `<body></body>`,
		},
		{
			Desc:     "imbalanced 2",
			Input:    "<body><noscript/></noscript></body>",
			Expected: `<body></body>`,
		},
	}
	runNodeCleanupTestCases(t, tcs)
}

func TestNodeCleanup_ReescapeText(t *testing.T) {
	basetcs := []tt.TestCase{
		{
			Desc:     "escape <",
			Input:    "<body>3 < 4</body>",
			Expected: "<body>3 &lt; 4</body>",
		},
		{
			Desc:     "&lt; noop",
			Input:    "<body>3 &lt; 4</body>",
			Expected: "<body>3 &lt; 4</body>",
		},
		{
			Desc:     "escape double quote",
			Input:    "<body>foo \" bar</body>",
			Expected: "<body>foo &#34; bar</body>",
		},
		{
			Desc:     "crazy",
			Input:    "<body><<!---->script>alert(42)<<!---->/script></body>",
			Expected: "<body>&lt;script&gt;alert(42)&lt;/script&gt;</body>",
		},
		{
			Desc:     "utf-8",
			Input:    "<body>\u00A0<body>",
			Expected: "<body>&nbsp;</body>",
		},
		{
			Desc:     "hex encoded",
			Input:    "<body>&#x61;</body>",
			Expected: "<body>a</body>",
		},
	}
	tcs := []tt.TestCase{}
	for _, tag := range []string{"body, title"} {
		for _, basetc := range basetcs {
			tc := tt.TestCase{
				Desc:     basetc.Desc,
				Input:    fmt.Sprintf("<%s>%s</%s>", tag, basetc.Input, tag),
				Expected: fmt.Sprintf("<%s>%s</%s>", tag, basetc.Expected, tag),
			}
			tcs = append(tcs, tc)
		}
	}
	runNodeCleanupTestCases(t, tcs)
}

func TestNodeCleanup_EscapeJspCharactersInScriptAndStyle(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc: "escape jsp",
			Input: `<head><style amp-custom>
           a { color: "<% %>" }
           b { color: "<% %>" }
         </style><script type=application/json>
           {
             "foo": "<% %>",
             "bar": "<% %>"
           }
         </script></head>`,
			Expected: `<head><style amp-custom>
           a { color: "\3c % %\3e " }
           b { color: "\3c % %\3e " }
         </style><script type=application/json>
           {
             "foo": "\u003c% %\u003e",
             "bar": "\u003c% %\u003e"
           }
         </script></head>`,
		},
	}
	runNodeCleanupTestCases(t, tcs)
}

func runNodeCleanupTestCases(t *testing.T, tcs []tt.TestCase) {
	for _, tc := range tcs {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s: html.Parse failed %q", tc.Input, err)
			continue
		}
		inputDOM, err := amphtml.NewDOM(inputDoc)
		if err != nil {
			t.Errorf("%s\namphtml.NewDOM for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		transformers.NodeCleanup(&transformers.Context{DOM: inputDOM})
		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
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
