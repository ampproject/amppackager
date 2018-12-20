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
	fooBaseURL  = "https://www.example.com/foo"
	barBaseURL  = "https://www.example.com/bar"
	relativeURL = "/foo"
)

func TestAbsoluteURLTansformer(t *testing.T) {
	tcs := []struct {
		desc, input, expected, baseURL string
	}{
		{
			desc:     "AmpImgSrcUrlNotChanged",
			input:    "<amp-img src=" + relativeURL + "></amp-img>",
			expected: "<amp-img src=" + relativeURL + "></amp-img>",
			baseURL:  fooBaseURL,
		},
		{
			desc:     "PortableUrlHasHash",
			input:    "<div src=" + relativeURL + "></div>",
			expected: "<div src=#></div>",
			baseURL:  fooBaseURL,
		},
		{
			desc:     "AbsoluteUrlHasNoHash",
			input:    "<form action=" + relativeURL + "></form>",
			expected: "<form action=" + fooBaseURL + "></form>",
			baseURL:  barBaseURL,
		},
		{
			desc:     "AttributeUrlsOnAnyTagBecomePortable",
			input:    "<div src=" + relativeURL + "></div>",
			expected: "<div src=" + fooBaseURL + "></div>",
			baseURL:  barBaseURL,
		},
		{
			desc: "AttributeUrlsOnAmpInstallServiceworkerTagBecomePortable",
			input: tt.Concat("<amp-install-serviceworker data-iframe-src=", relativeURL, " data-no-service-worker-fallback-shell-url=",
				relativeURL, "></amp-install-serviceworker>"),
			expected: tt.Concat("<amp-install-serviceworker data-iframe-src=", fooBaseURL, " data-no-service-worker-fallback-shell-url=",
				fooBaseURL, "></amp-install-serviceworker>"),
			baseURL: barBaseURL,
		},
		{
			desc: "AttributeUrlsOnAmpStoryTagBecomePortable",
			input: tt.Concat("<amp-story background-audio=", relativeURL, " bookend-config-src=", relativeURL,
				" poster-landscape-src=", relativeURL, " poster-square-src=", relativeURL,
				" publisher-logo-src=", relativeURL, "></amp-story>"),
			expected: tt.Concat("<amp-story background-audio=", fooBaseURL, " bookend-config-src=", fooBaseURL,
				" poster-landscape-src=", fooBaseURL, " poster-square-src=", fooBaseURL,
				" publisher-logo-src=", fooBaseURL, "></amp-story>"),
			baseURL: barBaseURL,
		},
		{
			desc:     "AttributeUrlsOnAmpStoryPageTagBecomePortable",
			input:    "<amp-story-page background-audio=" + relativeURL + "></amp-story-page>",
			expected: "<amp-story-page background-audio=" + fooBaseURL + "></amp-story-page>",
			baseURL:  barBaseURL,
		},
		{
			desc:     "AttributeUrlsOnFormTagBecomeAbsolute",
			input:    tt.Concat("<form action=", relativeURL, " action-xhr=", relativeURL, "></form>"),
			expected: tt.Concat("<form action=", fooBaseURL, " action-xhr=", fooBaseURL, "></form>"),
			baseURL:  barBaseURL,
		},
		{
			desc:     "AttributeUrlsOnImageTagBecomePortable",
			input:    "<img longdesc=" + relativeURL + "/>",
			expected: "<img longdesc=" + fooBaseURL + "/>",
			baseURL:  barBaseURL,
		},
		{
			desc:     "BaseHrefRemoved",
			input:    "<base href=" + fooBaseURL + "/>",
			expected: "",
			baseURL:  fooBaseURL,
		},
		{
			desc:     "Both tags parsed.",
			input:    "<base href=" + fooBaseURL + "/><link href=" + relativeURL + "/ rel=canonical>",
			expected: "<link href=" + fooBaseURL + "/ rel=canonical>",
			baseURL:  fooBaseURL,
		},
		{
			desc:     "LinkCanonicalHrefBecomeAbsolute",
			input:    "<link href=" + relativeURL + "/ rel=canonical>",
			expected: "<link href=" + fooBaseURL + "/ rel=canonical>",
			baseURL:  fooBaseURL,
		},
		{
			desc:     "AnchorTagHrefBecomesFragmentAndNoTargetAdded",
			input:    "<a href=" + relativeURL + ">anchor</a>",
			expected: "<a href=#>anchor</a>",
			baseURL:  fooBaseURL,
		},
		{
			desc:     "AnchorTagTargetDefaultsToTop",
			input:    "<a href=" + fooBaseURL + "/>anchor</a>",
			expected: "<a href=" + fooBaseURL + "/ target=_top>anchor</a>",
			baseURL:  fooBaseURL,
		},
		{
			desc:     "AnchorTagTargetStaysBlank",
			input:    "<a href=" + fooBaseURL + "/ target=_blank>anchor</a>",
			expected: "<a href=" + fooBaseURL + "/ target=_blank>anchor</a>",
			baseURL:  fooBaseURL,
		},
		{
			desc:     "AnchorTagTargetOverridesToDefault",
			input:    "<a href=" + fooBaseURL + "/ target=popup>anchor</a>",
			expected: "<a href=" + fooBaseURL + "/ target=_top>anchor</a>",
			baseURL:  fooBaseURL,
		},
		{
			desc:     "AnchorTagTargetInTemplateNoop",
			input:    "<template><a href=" + fooBaseURL + "/ target=popup>anchor</a></template>",
			expected: "<template><a href=" + fooBaseURL + "/ target=popup>anchor</a></template>",
			baseURL:  fooBaseURL,
		},
		{
			desc:     "NonAnchorHrefUrlBecomePortable",
			input:    "<link href=" + relativeURL + "/ itemprop=sameas/>",
			expected: "<link href=" + fooBaseURL + "/ itemprop=sameas/>",
			baseURL:  barBaseURL,
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
		context.BaseURL, err = url.Parse(tc.baseURL)
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
