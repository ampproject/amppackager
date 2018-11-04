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
	baseURL     = "https://www.example.com/foo"
	barBaseURL  = "https://www.example.com/bar"
	relativeURL = "/foo"
)

func TestURLTansformer(t *testing.T) {
	tcs := []struct {
		desc, input, expected, docURL string
	}{
		{
			desc:     "AmpImgSrcUrlNotChanged",
			input:    "<amp-img src=" + relativeURL + "></amp-img>",
			expected: "<amp-img src=" + relativeURL + "></amp-img>",
			docURL:   baseURL,
		},
		{
			desc:     "PortableUrlHasHash",
			input:    "<div src=" + relativeURL + "></div>",
			expected: "<div src=#></div>",
			docURL:   baseURL,
		},
		{
			desc:     "AbsoluteUrlHasNoHash",
			input:    "<form action=" + relativeURL + "></form>",
			expected: "<form action=" + baseURL + "></form>",
			docURL:   barBaseURL,
		},
		{
			desc:     "AttributeUrlsOnAnyTagBecomePortable",
			input:    "<div src=" + relativeURL + "></div>",
			expected: "<div src=" + baseURL + "></div>",
			docURL:   barBaseURL,
		},
		{
			desc: "AttributeUrlsOnAmpInstallServiceworkerTagBecomePortable",
			input: tt.Concat("<amp-install-serviceworker data-iframe-src=", relativeURL, " data-no-service-worker-fallback-shell-url=",
				relativeURL, "></amp-install-serviceworker>"),
			expected: tt.Concat("<amp-install-serviceworker data-iframe-src=", baseURL, " data-no-service-worker-fallback-shell-url=",
				baseURL, "></amp-install-serviceworker>"),
			docURL: barBaseURL,
		},
		{
			desc: "AttributeUrlsOnAmpStoryTagBecomePortable",
			input: tt.Concat("<amp-story background-audio=", relativeURL, " bookend-config-src=", relativeURL,
				" poster-landscape-src=", relativeURL, " poster-square-src=", relativeURL,
				" publisher-logo-src=", relativeURL, "></amp-story>"),
			expected: tt.Concat("<amp-story background-audio=", baseURL, " bookend-config-src=", baseURL,
				" poster-landscape-src=", baseURL, " poster-square-src=", baseURL,
				" publisher-logo-src=", baseURL, "></amp-story>"),
			docURL: barBaseURL,
		},
		{
			desc:     "AttributeUrlsOnAmpStoryPageTagBecomePortable",
			input:    "<amp-story-page background-audio=" + relativeURL + "></amp-story-page>",
			expected: "<amp-story-page background-audio=" + baseURL + "></amp-story-page>",
			docURL:   barBaseURL,
		},
		{
			desc:     "AttributeUrlsOnFormTagBecomeAbsolute",
			input:    tt.Concat("<form action=", relativeURL, " action-xhr=", relativeURL, "></form>"),
			expected: tt.Concat("<form action=", baseURL, " action-xhr=", baseURL, "></form>"),
			docURL:   barBaseURL,
		},
		{
			desc:     "AttributeUrlsOnImageTagBecomePortable",
			input:    "<img longdesc=" + relativeURL + "/>",
			expected: "<img longdesc=" + baseURL + "/>",
			docURL:   barBaseURL,
		},
		{
			desc:     "BaseHrefRemoved",
			input:    "<base href=" + baseURL + "/>",
			expected: "",
			docURL:   baseURL,
		},
		{
			desc:     "Both tags parsed.",
			input:    "<base href=" + baseURL + "/><link href=" + relativeURL + "/ rel=canonical>",
			expected: "<link href=" + baseURL + "/ rel=canonical>",
			docURL:   baseURL,
		},
		{
			desc:     "LinkCanonicalHrefBecomeAbsolute",
			input:    "<link href=" + relativeURL + "/ rel=canonical>",
			expected: "<link href=" + baseURL + "/ rel=canonical>",
			docURL:   baseURL,
		},
		{
			desc:     "AnchorTagHrefBecomesFragmentAndNoTargetAdded",
			input:    "<a href=" + relativeURL + ">anchor</a>",
			expected: "<a href=#>anchor</a>",
			docURL:   baseURL,
		},
		{
			desc:     "AnchorTagTargetDefaultsToTop",
			input:    "<a href=" + baseURL + "/>anchor</a>",
			expected: "<a href=" + baseURL + "/ target=_top>anchor</a>",
			docURL:   baseURL,
		},
		{
			desc:     "AnchorTagTargetStaysBlank",
			input:    "<a href=" + baseURL + "/ target=_blank>anchor</a>",
			expected: "<a href=" + baseURL + "/ target=_blank>anchor</a>",
			docURL:   baseURL,
		},
		{
			desc:     "AnchorTagTargetOverridesToDefault",
			input:    "<a href=" + baseURL + "/ target=popup>anchor</a>",
			expected: "<a href=" + baseURL + "/ target=_top>anchor</a>",
			docURL:   baseURL,
		},
		{
			desc:     "NonAnchorHrefUrlBecomePortable",
			input:    "<link href=" + relativeURL + "/ itemprop=sameas/>",
			expected: "<link href=" + baseURL + "/ itemprop=sameas/>",
			docURL:   barBaseURL,
		},
	}
	for _, tc := range tcs {
		rawInput := tt.Concat("<html><head>", tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
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
		context.DocumentURL, err = url.Parse(tc.docURL)
		if err != nil {
			t.Errorf("%s\nurl.Parse for %s failed %q", tc.desc, tc.docURL, err)
			continue
		}
		transformers.URL(&context)

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
