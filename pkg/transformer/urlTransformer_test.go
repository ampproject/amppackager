package transformer_test

import (
	"net/url"
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
// the expected normalized output from transformer.go.

const (
	baseURL     = "https://www.example.com/foo"
	barBaseURL  = "https://www.example.com/bar"
	relativeURL = "/foo"
)

// testCase stores the input HTML, expected output HTML, and an optional
// baseURL.
type urlTransformerTestCase struct {
	desc     string
	input    string
	expected string
	docURL   string
}

func TestURLTansformer(t *testing.T) {
	tcs := []urlTransformerTestCase{
		{
			"AmpImgSrcUrlNotChanged",
			"<amp-img src=" + relativeURL + "></amp-img>",
			"<amp-img src=" + relativeURL + "></amp-img>",
			baseURL,
		},
		{
			"PortableUrlHasHash",
			"<div src=" + relativeURL + "></div>",
			"<div src=#></div>",
			baseURL,
		},
		{
			"AbsoluteUrlHasNoHash",
			"<form action=" + relativeURL + "></form>",
			"<form action=" + baseURL + "></form>",
			barBaseURL,
		},
		{
			"AttributeUrlsOnAnyTagBecomePortable",
			"<div src=" + relativeURL + "></div>",
			"<div src=" + baseURL + "></div>",
			barBaseURL,
		},
		{
			"AttributeUrlsOnAmpInstallServiceworkerTagBecomePortable",
			tt.Concat("<amp-install-serviceworker data-iframe-src=", relativeURL, " data-no-service-worker-fallback-shell-url=",
				relativeURL, "></amp-install-serviceworker>"),
			tt.Concat("<amp-install-serviceworker data-iframe-src=", baseURL, " data-no-service-worker-fallback-shell-url=",
				baseURL, "></amp-install-serviceworker>"),
			barBaseURL,
		},
		{
			"AttributeUrlsOnAmpStoryTagBecomePortable",
			tt.Concat("<amp-story background-audio=", relativeURL, " bookend-config-src=", relativeURL,
				" poster-landscape-src=", relativeURL, " poster-square-src=", relativeURL,
				" publisher-logo-src=", relativeURL, "></amp-story>"),
			tt.Concat("<amp-story background-audio=", baseURL, " bookend-config-src=", baseURL,
				" poster-landscape-src=", baseURL, " poster-square-src=", baseURL,
				" publisher-logo-src=", baseURL, "></amp-story>"),
			barBaseURL,
		},
		{
			"AttributeUrlsOnAmpStoryPageTagBecomePortable",
			"<amp-story-page background-audio=" + relativeURL + "></amp-story-page>",
			"<amp-story-page background-audio=" + baseURL + "></amp-story-page>",
			barBaseURL,
		},
		{
			"AttributeUrlsOnFormTagBecomeAbsolute",
			tt.Concat("<form action=", relativeURL, " action-xhr=", relativeURL, "></form>"),
			tt.Concat("<form action=", baseURL, " action-xhr=", baseURL, "></form>"),
			barBaseURL,
		},
		{
			"AttributeUrlsOnImageTagBecomePortable",
			"<img longdesc=" + relativeURL + "/>",
			"<img longdesc=" + baseURL + "/>",
			barBaseURL,
		},
		{
			"BaseHrefRemoved",
			"<base href=" + baseURL + "/>",
			"",
			baseURL,
		},
		{
			"LinkCanonicalHrefBecomeAbsolute",
			"<link href=" + relativeURL + "/ rel=canonical>",
			"<link href=" + baseURL + "/ rel=canonical>",
			baseURL,
		},
		{
			"AnchorTagHrefBecomesFragmentAndNoTargetAdded",
			"<a href=" + relativeURL + ">anchor</a>",
			"<a href=#>anchor</a>",
			baseURL,
		},
		{
			"AnchorTagTargetDefaultsToTop",
			"<a href=" + baseURL + "/>anchor</a>",
			"<a href=" + baseURL + "/ target=_top>anchor</a>",
			baseURL,
		},
		{
			"AnchorTagTargetStaysBlank",
			"<a href=" + baseURL + "/ target=_blank>anchor</a>",
			"<a href=" + baseURL + "/ target=_blank>anchor</a>",
			baseURL,
		},
		{
			"AnchorTagTargetOverridesToDefault",
			"<a href=" + baseURL + "/ target=popup>anchor</a>",
			"<a href=" + baseURL + "/ target=_top>anchor</a>",
			baseURL,
		},
		{
			"NonAnchorHrefUrlBecomePortable",
			"<link href=" + relativeURL + "/ itemprop=sameas/>",
			"<link href=" + baseURL + "/ itemprop=sameas/>",
			barBaseURL,
		},
	}
	runURLTransformerTestCases(t, tcs)
}

func runURLTransformerTestCases(t *testing.T, tcs []urlTransformerTestCase) {

	for _, tc := range tcs {
		rawInput := tt.Concat("<html><head>", tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
			"</head><body>", tc.input, "</body></html>")
		inputDoc, err := html.Parse(strings.NewReader(rawInput))
		if err != nil {
			t.Errorf("%s\nhtml.Parse for %s failed %q", tc.desc, rawInput, err)
			continue
		}
		engine := transformer.Engine{Doc: inputDoc}
		engine.DocumentURL, err = url.Parse(tc.docURL)
		if err != nil {
			t.Errorf("%s\nurl.Parse for %s failed %q", tc.desc, tc.docURL, err)
			continue
		}
		transformer.URLTransformer(&engine)

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
		err = html.Render(&expected, expectedDoc)
		if err != nil {
			t.Errorf("%s\nhtml.Render for %s failed %q", tc.desc, rawExpected, err)
			continue
		}
		if input.String() != expected.String() {
			t.Errorf("%s: URLTransformer=\n%q\nwant=\n%q", tc.desc, &input, &expected)
		}
	}
}
