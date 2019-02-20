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
	"net/url"
	"strings"
	"testing"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	tt "github.com/ampproject/amppackager/transformer/internal/testing"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

const (
	fooURL      = "https://www.example.com/foo"
	barURL      = "https://www.example.com/bar"
	httpURL     = "http://www.example.com/"
	relativeURL = "/foo"
)

func TestAbsoluteURLTansformer(t *testing.T) {
	tcs := []struct {
		desc        string
		input       string
		expected    string
		baseURL     string
		documentURL string
	}{
		{
			desc: "Self URL not changed to fragment",
			// In this case the URL is the same as the document and base URL, but we
			// don't want it to be changed to be a fragment ("#").
			input:       "<div src=" + fooURL + "></div>",
			expected:    "<div src=" + fooURL + "></div>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "Non-http protocol preserved",
			input:       "<a href=mailto:devnull@example.com>mail</a>",
			expected:    "<a href=mailto:devnull@example.com target=_top>mail</a>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc: "Empty fragment preserved",
			// An URL of the current document with an empty fragment should not lose
			// the fragment, as this means that the link does nothing, rather than
			// reload the page.
			input:       "<a href=#>link</a>",
			expected:    "<a href=#>link</a>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "Simple fragment preserved",
			input:       "<a href=#top>link</a>",
			expected:    "<a href=#top>link</a>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "No fragment preserved",
			input:       "<form action=" + fooURL + "></form>",
			expected:    "<form action=" + fooURL + "></form>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "Empty URL preserved",
			input:       "<foo action></form>",
			expected:    "<foo action></form>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "Matching fragment simplified",
			input:       "<div src=/foo#dogs></div>",
			expected:    "<div src=#dogs></a>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "Non-matching fragment made absolute",
			input:       "<div src=/bar#dogs></div>",
			expected:    "<div src=" + barURL + "#dogs></a>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "Full relative preserves http",
			input:       "<div src=/></div>",
			expected:    "<div src=" + httpURL + "></a>",
			baseURL:     httpURL,
			documentURL: httpURL,
		},
		{
			desc:        "Protocol relative forces https",
			input:       "<div src=//www.example.com/foo></div>",
			expected:    "<div src=" + fooURL + "></a>",
			baseURL:     httpURL,
			documentURL: httpURL,
		},
		{
			desc:        "Base href removed",
			input:       "<base href='/'>",
			expected:    "",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "Link canonical href become absolute",
			input:       "<link href=" + relativeURL + " rel=canonical>",
			expected:    "<link href=" + fooURL + " rel=canonical>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "Link canonical ignores fragment",
			input:       "<link href=# rel=canonical>",
			expected:    "<link href=" + fooURL + " rel=canonical>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "Anchor tag target defaults to top",
			input:       "<a href=" + fooURL + "/>anchor</a>",
			expected:    "<a href=" + fooURL + "/ target=_top>anchor</a>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "Anchor tag target stays blank",
			input:       "<a href=" + fooURL + "/ target=_blank>anchor</a>",
			expected:    "<a href=" + fooURL + "/ target=_blank>anchor</a>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "Anchor tag target overrides to default",
			input:       "<a href=" + fooURL + "/ target=popup>anchor</a>",
			expected:    "<a href=" + fooURL + "/ target=_top>anchor</a>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "Anchor tag target in template no-op",
			input:       "<template><a href=" + fooURL + "/ target=popup>anchor</a></template>",
			expected:    "<template><a href=" + fooURL + "/ target=popup>anchor</a></template>",
			baseURL:     fooURL,
			documentURL: fooURL,
		},
		{
			desc:        "Non-anchor href URL becomes absolute",
			input:       "<link href=" + relativeURL + " itemprop=sameas/>",
			expected:    "<link href=" + fooURL + " itemprop=sameas/>",
			baseURL:     barURL,
			documentURL: barURL,
		},
		{
			// https://github.com/ampproject/amphtml/issues/19688
			desc:        "Base URL matches anchor",
			input:       "<a href='/'>foo</a>",
			expected:    "<a href=https://example.com/ target=_top>foo</a>",
			baseURL:     "https://example.com/",
			documentURL: barURL,
		},
		{
			desc:  "srcset rewritten",
			input: "<amp-img srcset=\"200.png 200w, 400.png 400w\"></amp-img>",
			expected: "<amp-img srcset=\"https://example.com/200.png 200w, " +
				"https://example.com/400.png 400w\"></amp-img>",
			baseURL:     "https://example.com/",
			documentURL: barURL,
		},
		{
			desc:        "missing authority",
			input:       "<a href=https:/foo.com/baz.html>Foo</a>",
			expected:    "<a href=https://www.example.com/foo.com/baz.html target=_top>Foo</a>",
			baseURL:     "http://www.example.com",
			documentURL: fooURL,
		},
	}
	for _, tc := range tcs {
		rawInput := tt.Concat(
			"<html><head>", tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
			"</head><body>", tc.input, "</body></html>")
		inputDoc, err := html.Parse(strings.NewReader(rawInput))
		if err != nil {
			t.Errorf("%s\nhtml.Parse for %s failed %q", tc.desc, rawInput, err)
			continue
		}
		inputDOM, err := amphtml.NewDOM(inputDoc)
		if err != nil {
			t.Errorf("%s\namphtml.NewDOM for %s failed %q", tc.desc, tc.input, err)
			continue
		}
		context := transformers.Context{DOM: inputDOM}
		context.BaseURL, err = url.Parse(tc.baseURL)
		if err != nil {
			t.Errorf("%s\nurl.Parse for %s failed %q", tc.desc, tc.baseURL, err)
			continue
		}
		context.DocumentURL, err = url.Parse(tc.documentURL)
		if err != nil {
			t.Errorf("%s\nurl.Parse for %s failed %q", tc.desc, tc.baseURL, err)
			continue
		}
		transformers.AbsoluteURL(&context)

		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
			t.Errorf("%s\nhtml.Render for %s failed %q", tc.desc, rawInput, err)
			continue
		}

		rawExpected := tt.Concat("<html><head>", tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
			"</head><body>", tc.expected, "</body></html>")
		expectedDoc, err := html.Parse(strings.NewReader(rawExpected))
		if err != nil {
			t.Errorf("%s\nhtml.Parse for %s failed %q", tc.desc, rawExpected, err)
			continue
		}
		var expected strings.Builder
		if err := html.Render(&expected, expectedDoc); err != nil {
			t.Errorf("%s\nhtml.Render for %s failed %q", tc.desc, rawExpected, err)
			continue
		}
		if input.String() != expected.String() {
			t.Errorf("%s: URL=\n%q\nwant=\n%q", tc.desc, &input, &expected)
		}
	}
}

func TestSpecificTagsAreAbsoluted(t *testing.T) {
	rawInput := tt.Concat(
		"<html><head>", tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
		`</head><body>
		<amp-install-serviceworker
       data-iframe-src=/foo
       data-no-service-worker-fallback-shell-url=/foo>
    </amp-install-serviceworker>
    <amp-story
       background-audio=/foo
       bookend-config-src=/foo
       poster-landscape-src=/foo
       poster-square-src=/foo
       publisher-logo-src=/foo>
    </amp-story>
    <amp-story-page background-audio=/foo></amp-story-page>
    <form action=/foo></form>
    <form action-xhr=/foo></form>
		</body></html>`)
	// TODO(gregable): Another good test would be:
	//   <noscript><img longdesc=/foo></form></noscript>
	// but net/url currently doesn't parse noscript contents.
	inputDoc, err := html.Parse(strings.NewReader(rawInput))
	if err != nil {
		t.Errorf("html.Parse for %s failed %q", rawInput, err)
		return
	}
	inputDOM, err := amphtml.NewDOM(inputDoc)
	if err != nil {
		t.Errorf("amphtml.NewDOM failed %q", err)
		return
	}
	context := transformers.Context{DOM: inputDOM}
	context.BaseURL, _ = url.Parse(fooURL)
	context.DocumentURL, _ = url.Parse(fooURL)
	transformers.AbsoluteURL(&context)
	var seen strings.Builder
	if err := html.Render(&seen, inputDoc); err != nil {
		t.Errorf("html.Render for %s failed %q", rawInput, err)
		return
	}

	if strings.Contains(seen.String(), "=/foo") ||
		strings.Contains(seen.String(), "\"/foo") ||
		strings.Contains(seen.String(), "'/foo") {
		t.Errorf("Relative URL found: %s", seen.String())
	}
}
