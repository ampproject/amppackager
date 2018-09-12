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

package printer_test

import (
	"strings"
	"testing"

	tt "github.com/ampproject/amppackager/transformer/internal/testing"

	"github.com/ampproject/amppackager/transformer/printer"
	"golang.org/x/net/html"
)

func TestQuotes(t *testing.T) {
	testCases := []tt.TestCase{
		{
			"Drops unnecessary quotes",
			`<div class="content" id="content">`,
			"<div class=content id=content></div>",
		},
		{
			"Doesn't strip quotes",
			`<p class="normal text" style="font-size: 12pt;">Hello World!</p>`,
			`<p class="normal text" style="font-size: 12pt;">Hello World!</p>`,
		},
		{
			"Some quotes stripped",
			`<IMG src="http://www.google.com/a.jpg" alt="This is a image.">`,
			`<img alt="This is a image." src=http://www.google.com/a.jpg>`,
		},
		{
			"space is double quoted",
			`<lemur x=" ">`,
			`<lemur x=" ">`,
		},
		{
			"double quote is single quoted",
			"<lemur x='\"'>",
			"<lemur x=&#34;>",
		},
		{
			"= is quoted",
			`<lemur x='a=b'>`,
			`<lemur x="a=b">`,
		},
		{
			"< is escaped and no quotes",
			`<lemur x='a<b'>`,
			`<lemur x=a&lt;b>`,
		},
		{
			"> is escaped and no quotes",
			`<lemur x='a>b'>`,
			`<lemur x=a&gt;b>`,
		},
		{
			"utf8 symbol is unquoted",
			`<lemur x="❄">`,
			`<lemur x=❄>`,
		},
		{
			"utf8 turkish is unquoted",
			`<lemur x="Beşiktaş">`,
			`<lemur x=Beşiktaş>`,
		},
		{
			"utf8 russian is unquoted",
			`<lemur x="Вконтакте">`,
			`<lemur x=Вконтакте>`,
		},
	}
	runAllTestCases(t, testCases)
}

func TestStripsDocTypeAttrs(t *testing.T) {
	testCases := []tt.TestCase{
		{
			"No-op",
			"<!doctype html>",
			"<!doctype html>",
		},
		{
			"Strips all attrs",
			`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">`,
			"<!doctype html>",
		},
		{
			"Strips bogus attrs",
			`<!DOCTYPE HTML PUBLIC "bogus" "notreal">`,
			"<!doctype html>",
		},
		{
			"Bogus doctype",
			`<!DOCTYPE document SYSTEM "subjects.dtd">`,
			"<!doctype document>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestAddsRequiredTags(t *testing.T) {
	testCases := []tt.TestCase{
		{
			"Adds html, head, and body tags.",
			tt.Concat("<!doctype html>", tt.ScriptAMPRuntime, tt.LinkFavicon, "hello world"),
			tt.Concat("<!doctype html><html><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, "</head><body>hello world</body></html>"),
		},
	}
	runAllTestCases(t, testCases)
}

func TestStripsWhitespace(t *testing.T) {
	testCases := []tt.TestCase{
		{
			"Strip whitespace pre doctype",
			"\n\n\t\t    <!doctype html>",
			"<!doctype html><html><head></head><body></body></html>",
		},
		{
			"Strips intra tag whitespace",
			`<lemur   lemur = "lemur" ></lemur>`,
			"<lemur lemur=lemur></lemur>",
		},
		{
			"Should not affect pre tags",
			"<pre>   foo   </pre>",
			"<pre>   foo   </pre>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestStripComments(t *testing.T) {
	testCases := []tt.TestCase{
		{
			"Strip comments",
			tt.Concat("<!-- comment --><!doctype html><html ⚡>",
				"<foo><!-- comment --></foo>"),
			"<!doctype html><html ⚡><head></head><body><foo></foo></body></html>",
		},
		{
			"Strip comments embedded within text",
			tt.Concat(
				"<!-- All comments --><!doctype html><!-- are --><html ⚡><head>",
				"</head><body>are <!-- belong --><p><!-- to --> us!</p></body>",
				"</html>"),
			"<!doctype html><html ⚡><head></head><body>are <p> us!</p></body></html>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestClosesTags(t *testing.T) {
	testCases := []tt.TestCase{
		{
			"Closes tags.",
			"<lemur>",
			"<lemur></lemur>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestVoidTags(t *testing.T) {
	testCases := []tt.TestCase{
		{
			"Noop",
			"<br>",
			"<br>",
		},
		{
			"Strip end tag",
			"<br/>",
			"<br>",
		},
		{
			"End tag without start tag",
			"</br>",
			"<br>",
		},
		{
			"Strip end tag with crazy spacing",
			"<img src  = 'lemur.png' />",
			"<img src=lemur.png>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestSelfClosedTags(t *testing.T) {
	testCases := []tt.TestCase{
		{
			// Self-closed tags are invalid HTML5. They should either be void tags or
			// proper start followed by end tag.
			"Self-closed tags",
			"<lemur />foo",
			"<lemur>foo</lemur>",
		},
		{
			"Add end tag (when closing body tag is encountered)",
			"<lemur/>foo",
			"<lemur>foo</lemur>",
		},
		{
			"Drop redundant end tag.",
			"<lemur/>foo</lemur>bar",
			"<lemur>foo</lemur>bar",
		},
		{
			// <style> and <script> tags are handled differently
			// One would expect to produce "<style>foo</style> but we don't. Instead the
			// rest of the document gets read in and then the parser closes the tags
			// already seen to that point.
			// Given:
			// <html><head></head><body><style/>foo</body></html>
			// Reserializes:
			// <html><head></head><body><style>foo</body></html></style></body></html>
			// This is mostly parser behavior and maybe unnecessary for this test.
			"Style tag",
			"<html><head></head><body><style />foo</body></html>",
			"<style>foo</body></html></style>",
		},
		{
			// However, within SVG tags, self-closed tags are interpreted in
			// HTML5. This is handled by the parser, so might be
			// unnecessary here.
			"svg",
			"<svg><lemur />foo</svg>",
			"<svg><lemur></lemur>foo</svg>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestEscaping(t *testing.T) {
	testCases := []tt.TestCase{
		{
			"Attr value",
			`<lemur lemur="<>'">`,
			`<lemur lemur=&lt;&gt;&#39;></lemur>`,
		},
		{
			"Quote",
			`<lemur koala='"'>`,
			// Note that &#34; is used over &quot;
			// https://github.com/golang/net/blob/master/html/escape.go#L215
			`<lemur koala=&#34;></lemur>`,
		},
		{
			// Make sure that we don't unescape and then fail to re-escape.
			"No-op",
			"&lt;script&gt;",
			"&lt;script&gt;",
		},
		{
			"Script attrs escaped",
			`<script type="application/json">{ "AmpBind": true }</script>`,
			`<script type=application/json>{ "AmpBind": true }</script>`,
		},
	}
	runAllTestCases(t, testCases)
}

func TestOrderedAttrs(t *testing.T) {
	testCases := []tt.TestCase{
		{
			"Order",
			"<lemur x=3 y=4 b=5 />",
			"<lemur b=5 x=3 y=4></lemur>",
		},
		{
			"No secondary sort, instead relies on order.",
			"<lemur x=4 x=3 b=5 />",
			"<lemur b=5 x=4 x=3></lemur>",
		},
		{
			"svg",
			`<svg version="1.0" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="640" height="480"></svg>`,
			"<svg height=480 version=1.0 width=640 xmlns=http://www.w3.org/2000/svg xmlns:xlink=http://www.w3.org/1999/xlink></svg>",
		},
		{
			"Namespace",
			"<lemur x:foo y a>",
			"<lemur a x:foo y></lemur>",
		},
		{
			"More namespace.",
			"<lemur x foob:az foo:bar a>",
			"<lemur a foo:bar foob:az x></lemur>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestLowerCaseTagsAndAttrs(t *testing.T) {
	testCases := []tt.TestCase{
		{
			"svg attribute names",
			"<svg height=296 viewBox=\"1400 500 3000 2500\" width=400></svg>",
			"<svg height=296 viewbox=\"1400 500 3000 2500\" width=400></svg>",
		},
		{
			// This test looks like a no-op, but in fact, the golang parser
			// only camelCases svg tags if they are lowercase to being with.
			// So the test is verifying we are re-lowercasing it.
			"svg child tag",
			"<svg><lineargradient>",
			"<svg><lineargradient>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestEOFNewline(t *testing.T) {
	testCases := []tt.TestCase{
		{
			"newline added",
			"hello world",
			"<html><head></head><body>hello world</body></html>\n",
		},
	}
	runAllTestCases(t, testCases)
}

func runAllTestCases(t *testing.T, testCases []tt.TestCase) {
	for _, tc := range testCases {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s: htmlParse on %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		var input strings.Builder
		err = printer.Print(&input, inputDoc)
		if err != nil {
			t.Errorf("%s: printer.Print on %s failed %q", tc.Desc, tc.Input, err)
		}

		if !strings.Contains(input.String(), tc.Expected) {
			t.Errorf("%s: Print=\n%q\ndoes not contain expected=\n%q", tc.Desc, &input, tc.Expected)
		}
	}
}
