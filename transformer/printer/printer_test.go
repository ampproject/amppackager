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
			Desc: "Drops unnecessary quotes",
			Input: `<div class="content" id="content">`,
			Expected: "<div class=content id=content></div>",
		},
		{
			Desc: "Doesn't strip quotes",
			Input: `<p class="normal text" style="font-size: 12pt;">Hello World!</p>`,
			Expected: `<p class="normal text" style="font-size: 12pt;">Hello World!</p>`,
		},
		{
			Desc: "Some quotes stripped",
			Input: `<IMG src="http://www.google.com/a.jpg" alt="This is a image.">`,
			Expected: `<img alt="This is a image." src=http://www.google.com/a.jpg>`,
		},
		{
			Desc: "space is double quoted",
			Input: `<lemur x=" ">`,
			Expected: `<lemur x=" ">`,
		},
		{
			Desc: "double quote is single quoted",
			Input: "<lemur x='\"'>",
			Expected: "<lemur x=&#34;>",
		},
		{
			Desc: "= is quoted",
			Input: `<lemur x='a=b'>`,
			Expected: `<lemur x="a=b">`,
		},
		{
			Desc: "< is escaped and no quotes",
			Input: `<lemur x='a<b'>`,
			Expected: `<lemur x=a&lt;b>`,
		},
		{
			Desc: "> is escaped and no quotes",
			Input: `<lemur x='a>b'>`,
			Expected: `<lemur x=a&gt;b>`,
		},
		{
			Desc: "utf8 symbol is unquoted",
			Input: `<lemur x="❄">`,
			Expected: `<lemur x=❄>`,
		},
		{
			Desc: "utf8 turkish is unquoted",
			Input: `<lemur x="Beşiktaş">`,
			Expected: `<lemur x=Beşiktaş>`,
		},
		{
			Desc: "utf8 russian is unquoted",
			Input: `<lemur x="Вконтакте">`,
			Expected: `<lemur x=Вконтакте>`,
		},
	}
	runAllTestCases(t, testCases)
}

func TestStripsDocTypeAttrs(t *testing.T) {
	testCases := []tt.TestCase{
		{
			Desc: "No-op",
			Input: "<!doctype html>",
			Expected: "<!doctype html>",
		},
		{
			Desc: "Strips all attrs",
			Input: `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">`,
			Expected: "<!doctype html>",
		},
		{
			Desc: "Strips bogus attrs",
			Input: `<!DOCTYPE HTML PUBLIC "bogus" "notreal">`,
			Expected: "<!doctype html>",
		},
		{
			Desc: "Bogus doctype",
			Input: `<!DOCTYPE document SYSTEM "subjects.dtd">`,
			Expected: "<!doctype document>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestAddsRequiredTags(t *testing.T) {
	testCases := []tt.TestCase{
		{
			Desc: "Adds html, head, and body tags.",
			Input: tt.Concat("<!doctype html>", tt.ScriptAMPRuntime, tt.LinkFavicon, "hello world"),
			Expected: tt.Concat("<!doctype html><html><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, "</head><body>hello world</body></html>"),
		},
	}
	runAllTestCases(t, testCases)
}

func TestStripsWhitespace(t *testing.T) {
	testCases := []tt.TestCase{
		{
			Desc: "Strip whitespace pre doctype",
			Input: "\n\n\t\t    <!doctype html>",
			Expected: "<!doctype html><html><head></head><body></body></html>",
		},
		{
			Desc: "Strips intra tag whitespace",
			Input: `<lemur   lemur = "lemur" ></lemur>`,
			Expected: "<lemur lemur=lemur></lemur>",
		},
		{
			Desc: "Should not affect pre tags",
			Input: "<pre>   foo   </pre>",
			Expected: "<pre>   foo   </pre>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestStripComments(t *testing.T) {
	testCases := []tt.TestCase{
		{
			Desc: "Strip comments",
			Input: tt.Concat("<!-- comment --><!doctype html><html ⚡>",
				"<foo><!-- comment --></foo>"),
			Expected: "<!doctype html><html ⚡><head></head><body><foo></foo></body></html>",
		},
		{
			Desc: "Strip comments embedded within text",
			Input: tt.Concat(
				"<!-- All comments --><!doctype html><!-- are --><html ⚡><head>",
				"</head><body>are <!-- belong --><p><!-- to --> us!</p></body>",
				"</html>"),
			Expected: "<!doctype html><html ⚡><head></head><body>are <p> us!</p></body></html>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestClosesTags(t *testing.T) {
	testCases := []tt.TestCase{
		{
			Desc: "Closes tags.",
			Input: "<lemur>",
			Expected: "<lemur></lemur>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestVoidTags(t *testing.T) {
	testCases := []tt.TestCase{
		{
			Desc: "Noop",
			Input: "<br>",
			Expected: "<br>",
		},
		{
			Desc: "Strip ending slash",
			Input: "<br/>",
			Expected: "<br>",
		},
		{
			Desc: "End tag without start tag",
			Input: "</br>",
			Expected: "<br>",
		},
		{
			Desc: "Strip end tag with crazy spacing",
			Input: "<img src  = 'lemur.png' />",
			Expected: "<img src=lemur.png>",
		},
		{
			Desc: "Keep ending slash for void element in foreign content (SVG)",
			Input: "<svg><link rel=alternate /></svg>",
			Expected: "<svg><link rel=alternate /></svg>",
		},
		{
			Desc: "Keep ending slash for void element in foreign content (MathML)",
			Input: "<math><link rel=alternate /></math>",
			Expected: "<math><link rel=alternate /></math>",
		},
		{
			Desc: "Strip ending slash for void element in HTML integration point (SVG)",
			Input: "<svg><foreignobject><link rel=alternate /></foreignobject></svg>",
			Expected: "<svg><foreignobject><link rel=alternate></foreignobject></svg>",
		},
		{
			Desc: "Strip ending slash for void element in HTML integration point (MathML)",
			Input: "<math><annotation-xml encoding=text/html><link rel=alternate /></annotation-xml></math>",
			Expected: "<math><annotation-xml encoding=text/html><link rel=alternate></annotation-xml></math>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestSelfClosedTags(t *testing.T) {
	testCases := []tt.TestCase{
		{
			// Self-closed tags are invalid HTML5. They should either be void tags or
			// proper start followed by end tag.
			Desc: "Self-closed tags",
			Input: "<lemur />foo",
			Expected: "<lemur>foo</lemur>",
		},
		{
			Desc: "Add end tag (when closing body tag is encountered)",
			Input: "<lemur/>foo",
			Expected: "<lemur>foo</lemur>",
		},
		{
			Desc: "Drop redundant end tag.",
			Input: "<lemur/>foo</lemur>bar",
			Expected: "<lemur>foo</lemur>bar",
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
			Desc: "Style tag",
			Input: "<html><head></head><body><style />foo</body></html>",
			Expected: "<style>foo</body></html></style>",
		},
		{
			// However, within SVG tags, self-closed tags are interpreted in
			// HTML5. This is handled by the parser, so might be
			// unnecessary here.
			Desc: "svg",
			Input: "<svg><lemur />foo</svg>",
			Expected: "<svg><lemur></lemur>foo</svg>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestEscaping(t *testing.T) {
	testCases := []tt.TestCase{
		{
			Desc: "Attr value",
			Input: `<lemur lemur="<>'">`,
			Expected: `<lemur lemur=&lt;&gt;&#39;></lemur>`,
		},
		{
			Desc: "Quote",
			Input: `<lemur koala='"'>`,
			// Note that &#34; is used over &quot;
			// https://github.com/golang/net/blob/master/html/escape.go#L215
			Expected: `<lemur koala=&#34;></lemur>`,
		},
		{
			// Make sure that we don't unescape and then fail to re-escape.
			Desc: "No-op",
			Input: "&lt;script&gt;",
			Expected: "&lt;script&gt;",
		},
		{
			Desc: "Script attrs escaped",
			Input: `<script type="application/json">{ "AmpBind": true }</script>`,
			Expected: `<script type=application/json>{ "AmpBind": true }</script>`,
		},
	}
	runAllTestCases(t, testCases)
}

func TestOrderedAttrs(t *testing.T) {
	testCases := []tt.TestCase{
		{
			Desc: "Order",
			Input: "<lemur x=3 y=4 b=5 />",
			Expected: "<lemur b=5 x=3 y=4></lemur>",
		},
		{
			Desc: "No secondary sort, instead relies on order.",
			Input: "<lemur x=4 x=3 b=5 />",
			Expected: "<lemur b=5 x=4 x=3></lemur>",
		},
		{
			Desc: "svg",
			Input: `<svg version="1.0" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="640" height="480"></svg>`,
			Expected: "<svg height=480 version=1.0 width=640 xmlns=http://www.w3.org/2000/svg xmlns:xlink=http://www.w3.org/1999/xlink></svg>",
		},
		{
			Desc: "Namespace",
			Input: "<lemur x:foo y a>",
			Expected: "<lemur a x:foo y></lemur>",
		},
		{
			Desc: "More namespace.",
			Input: "<lemur x foob:az foo:bar a>",
			Expected: "<lemur a foo:bar foob:az x></lemur>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestLowerCaseTagsAndAttrs(t *testing.T) {
	testCases := []tt.TestCase{
		{
			Desc: "svg attribute names",
			Input: "<svg height=296 viewBox=\"1400 500 3000 2500\" width=400></svg>",
			Expected: "<svg height=296 viewbox=\"1400 500 3000 2500\" width=400></svg>",
		},
		{
			// This test looks like a no-op, but in fact, the golang parser
			// only camelCases svg tags if they are lowercase to being with.
			// So the test is verifying we are re-lowercasing it.
			Desc: "svg child tag",
			Input: "<svg><lineargradient>",
			Expected: "<svg><lineargradient>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestPreLeadingNewline(t *testing.T) {
	testCases := []tt.TestCase{
		{
			// This is actually done in the parser, as per spec. Verifying
			// behavior here.
			Desc: "Newline at start of <pre> is dropped.",
			Input: "<pre>&#13;</pre>",
			Expected: "<pre></pre>",
		},
		{
			// This is also done in the parser, as per spec, where the
			// combo of CR LF is treated as a single newline. Verifying
			// behavior here.
			Desc: "CR LF",
			Input: "<pre>&#13;\n</pre>",
			Expected: "<pre></pre>",
		},
		{
			Desc: "Add LF to <pre> when start with CR.",
			Input: "<pre>&#13;&#13;</pre>",
			Expected: "<pre>\n&#13;</pre>",
		},
		{
			Desc: "Add LF to <pre> when start with LF.",
			Input: "<pre>&#10;&#10;</pre>",
			Expected: "<pre>\n\n</pre>",
		},
		{
			Desc: "Preserve LF LF when comment in the middle.",
			Input: "<pre>&#10;<!-- comment -->&#10;</pre>",
			Expected: "<pre>\n\n</pre>",
		},
		{
			Desc: "Add LF to <pre> when comment followed by LF.",
			Input: "<pre><!-- comment -->&#10;</pre>",
			Expected: "<pre>\n\n</pre>",
		},
		{
			Desc: "Add LF to LF LF preceded by comment.",
			Input: "<pre><!-- comment -->&#10;&#10;</pre>",
			// HTML parsers will strip the first LF, thus
			// preserving the meaning of the originally non-leading
			// LF LF:
			Expected: "<pre>\n\n\n</pre>",
		},
		{
			Desc: "LF LF with more text",
			Input: `<pre>&#10;&#10;lemur</pre>`,
			Expected: "<pre>\n\nlemur</pre>",
		},
		{
			Desc: "Don't add extra LF to <pre> with leading text.",
			Input: "<pre>blah&#10;&#13;</pre>",
			Expected: "<pre>blah\n&#13;</pre>",
		},
		{
			Desc: "Don't add extra LF to <pre> with leading element.",
			Input: "<pre><strong>boo</strong>&#10;&#13;</pre>",
			Expected: "<pre><strong>boo</strong>\n&#13;</pre>",
		},
	}
	runAllTestCases(t, testCases)
}

func TestTextareaLeadingNewline(t *testing.T) {
	testCases := []tt.TestCase{
		{
			// The same logic for <pre> also applies for <textarea>.
			Desc: "Add LF to <textarea> when start with CR.",
			Input: "<pre>&#13;&#13;</pre>",
			Expected: "<pre>\n&#13;</pre>",
		},
	}
	runAllTestCases(t, testCases)
}

func runAllTestCases(t *testing.T, testCases []tt.TestCase) {
	for _, tc := range testCases {
		InputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s: htmlParse on %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		var Input strings.Builder
		err = printer.Print(&Input, InputDoc)
		if err != nil {
			t.Errorf("%s: printer.Print on %s failed %q", tc.Desc, tc.Input, err)
		}

		if !strings.Contains(Input.String(), tc.Expected) {
			t.Errorf("%s: Print=\n%q\ndoes not contain Expected=\n%q", tc.Desc, &Input, tc.Expected)
		}
	}
}
