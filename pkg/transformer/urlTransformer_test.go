package transformer_test

import (
	"net/url"
	"strings"
	"testing"

	tt "github.com/ampproject/amppackager/internal/pkg/testing"
	"github.com/ampproject/amppackager/pkg/transformer"
	"golang.org/x/net/html"
)

// These tests do NOT run through the custom transformations of the
// Engine, and instead rely exclusively on vanilla golang parser and
// renderer (otherwise the scope of these tests would creep past unit
// testing). Therefore, the test data must be made to match, and is not
// the expected normalized output from transformer.go.

const (
	baseURL     = "https://www.example.com/foo"
	barBaseURL  = "https://www.example.com/bar"
	relativeURL = "/foo"
)

// testCase stores the input HTML, expected output HTML, and an optional
// baseURL.
type urlTransformerTestCase struct {
	Desc     string
	Input    string
	Expected string
	DocURL   string
}

func TestURLTansformer(t *testing.T) {
	tcs := []urlTransformerTestCase{
		{
			Desc:     "AmpImgSrcUrlNotChanged",
			Input:    "<amp-img src=" + relativeURL + "></amp-img>",
			Expected: "<amp-img src=" + relativeURL + "></amp-img>",
			DocURL:   baseURL,
		},
		{
			Desc:     "PortableUrlHasHash",
			Input:    "<div src=" + relativeURL + "></div>",
			Expected: "<div src=#></div>",
			DocURL:   baseURL,
		},
		{
			Desc:     "AbsoluteUrlHasNoHash",
			Input:    "<form action=" + relativeURL + "></form>",
			Expected: "<form action=" + baseURL + "></form>",
			DocURL:   barBaseURL,
		},
		{
			Desc:     "AttributeUrlsOnAnyTagBecomePortable",
			Input:    "<div src=" + relativeURL + "></div>",
			Expected: "<div src=" + baseURL + "></div>",
			DocURL:   barBaseURL,
		},
		{
			Desc: "AttributeUrlsOnAmpInstallServiceworkerTagBecomePortable",
			Input: tt.Concat("<amp-install-serviceworker data-iframe-src=", relativeURL, " data-no-service-worker-fallback-shell-url=",
				relativeURL, "></amp-install-serviceworker>"),
			Expected: tt.Concat("<amp-install-serviceworker data-iframe-src=", baseURL, " data-no-service-worker-fallback-shell-url=",
				baseURL, "></amp-install-serviceworker>"),
			DocURL: barBaseURL,
		},
		{
			Desc: "AttributeUrlsOnAmpStoryTagBecomePortable",
			Input: tt.Concat("<amp-story background-audio=", relativeURL, " bookend-config-src=", relativeURL,
				" poster-landscape-src=", relativeURL, " poster-square-src=", relativeURL,
				" publisher-logo-src=", relativeURL, "></amp-story>"),
			Expected: tt.Concat("<amp-story background-audio=", baseURL, " bookend-config-src=", baseURL,
				" poster-landscape-src=", baseURL, " poster-square-src=", baseURL,
				" publisher-logo-src=", baseURL, "></amp-story>"),
			DocURL: barBaseURL,
		},
		{
			Desc:     "AttributeUrlsOnAmpStoryPageTagBecomePortable",
			Input:    "<amp-story-page background-audio=" + relativeURL + "></amp-story-page>",
			Expected: "<amp-story-page background-audio=" + baseURL + "></amp-story-page>",
			DocURL:   barBaseURL,
		},
		{
			Desc:     "AttributeUrlsOnFormTagBecomeAbsolute",
			Input:    tt.Concat("<form action=", relativeURL, " action-xhr=", relativeURL, "></form>"),
			Expected: tt.Concat("<form action=", baseURL, " action-xhr=", baseURL, "></form>"),
			DocURL:   barBaseURL,
		},
		{
			Desc:     "AttributeUrlsOnImageTagBecomePortable",
			Input:    "<img longdesc=" + relativeURL + "/>",
			Expected: "<img longdesc=" + baseURL + "/>",
			DocURL:   barBaseURL,
		},
		{
			Desc:     "BaseHrefRemoved",
			Input:    "<base href=" + baseURL + "/>",
			Expected: "",
			DocURL:   baseURL,
		},
		{
			Desc:     "LinkCanonicalHrefBecomeAbsolute",
			Input:    "<link href=" + relativeURL + "/ rel=canonical>",
			Expected: "<link href=" + baseURL + "/ rel=canonical>",
			DocURL:   baseURL,
		},
		{
			Desc:     "AnchorTagHrefBecomesFragmentAndNoTargetAdded",
			Input:    "<a href=" + relativeURL + ">anchor</a>",
			Expected: "<a href=#>anchor</a>",
			DocURL:   baseURL,
		},
		{
			Desc:     "AnchorTagTargetDefaultsToTop",
			Input:    "<a href=" + baseURL + "/>anchor</a>",
			Expected: "<a href=" + baseURL + "/ target=_top>anchor</a>",
			DocURL:   baseURL,
		},
		{
			Desc:     "AnchorTagTargetStaysBlank",
			Input:    "<a href=" + baseURL + "/ target=_blank>anchor</a>",
			Expected: "<a href=" + baseURL + "/ target=_blank>anchor</a>",
			DocURL:   baseURL,
		},
		{
			Desc:     "AnchorTagTargetOverridesToDefault",
			Input:    "<a href=" + baseURL + "/ target=popup>anchor</a>",
			Expected: "<a href=" + baseURL + "/ target=_top>anchor</a>",
			DocURL:   baseURL,
		},
		{
			Desc:     "NonAnchorHrefUrlBecomePortable",
			Input:    "<link href=" + relativeURL + "/ itemprop=sameas/>",
			Expected: "<link href=" + baseURL + "/ itemprop=sameas/>",
			DocURL:   barBaseURL,
		},
	}
	runURLTransformerTestCases(t, tcs)
}

func runURLTransformerTestCases(t *testing.T, tcs []urlTransformerTestCase) {

	for _, tc := range tcs {
		rawInput := tt.Concat("<html><head>", tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
			"</head><body>", tc.Input, "</body></html>")
		inputDoc, err := html.Parse(strings.NewReader(rawInput))
		if err != nil {
			t.Errorf("%s\nhtml.Parse for %s failed %q", tc.Desc, rawInput, err)
			continue
		}
		engine := transformer.Engine{Doc: inputDoc}
		engine.DocumentURL, err = url.Parse(tc.DocURL)
		if err != nil {
			t.Errorf("%s\nurl.Parse for %s failed %q", tc.Desc, tc.DocURL, err)
			continue
		}
		transformer.URLTransformer(&engine)

		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
			t.Errorf("%s\nhtml.Render for %s failed %q", tc.Desc, rawInput, err)
			continue
		}

		rawExpected := tt.Concat("<html><head>", tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
			"</head><body>", tc.Expected, "</body></html>")
		expectedDoc, err := html.Parse(strings.NewReader(rawExpected))
		if err != nil {
			t.Errorf("%s\nhtml.Parse for %s failed %q", tc.Desc, rawExpected, err)
			continue
		}
		var expected strings.Builder
		err = html.Render(&expected, expectedDoc)
		if err != nil {
			t.Errorf("%s\nhtml.Render for %s failed %q", tc.Desc, rawExpected, err)
			continue
		}
		if input.String() != expected.String() {
			t.Errorf("%s: URLTransformer=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
