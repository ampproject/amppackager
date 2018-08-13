package transformer_test

import (
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

// ampBoilerplateNoscriptWithAttr is the equivalent of
// tt.AMPBoilerplateNoscript except it includes empty attribute
// values. This is needed because the golang parser treats everything
// inside <noscript> as one giant text node, so the <noscript> tag in
// the expected HTML isn't parsed properly.
const ampBoilerplateNoscriptWithAttr = "<noscript><style amp-boilerplate=\"\">body{-webkit-animation:none;-moz-animation:none;-ms-animation:none;animation:none}</style></noscript>"

func TestAMPBoilerplate(t *testing.T) {
	canonicalExpected := tt.Concat("<!doctype html><html ⚡><head>",
		tt.ScriptAMPRuntime, tt.LinkFavicon, tt.StyleAMPBoilerplate,
		ampBoilerplateNoscriptWithAttr, "</head><body></body></html>")

	testCases := []tt.TestCase{
		{
			"Keeps boilerplate",
			tt.Concat("<!doctype html><html ⚡><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			canonicalExpected,
		},
		{
			"Adds boilerplate if missing",
			tt.Concat("<!doctype html><html ⚡><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, "</head><body></body></html>"),
			canonicalExpected,
		},
		{
			"Upgrades old boilerplate",
			tt.Concat("<!doctype html><html ⚡><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon,
				"<style>body {opacity: 0}</style>",
				"<noscript><style>body {opacity: 1}</style></noscript></head>",
				"<body></body></html>"),
			canonicalExpected,
		},
		{
			// The validator actually allows both old and new boilerplate to be present.
			// This test ensures we always strip multiple instances and end up
			// with just the new boilerplate.
			"Strips old and new if both present",
			tt.Concat("<!doctype html><html ⚡><head>", tt.ScriptAMPRuntime,
				tt.NoscriptAMPBoilerplate, tt.LinkFavicon,
				"<style>body {opacity: 0}</style>", tt.StyleAMPBoilerplate,
				"<noscript><style>body {opacity: 1}</style></noscript></head>",
				"</head><body></body></html>"),
			canonicalExpected,
		},
	}

	runAMPBoilerplateTransformerTestcases(t, testCases)
}

func TestAMP4Ads(t *testing.T) {
	expected := func(attr string) string {
		return tt.Concat("<!doctype html><html ", attr, "><head>",
			tt.ScriptAMPRuntime, tt.LinkFavicon,
			"<style amp4ads-boilerplate>body{visibility:hidden}</style></head>",
			"<body></body></html>")
	}

	testCases := []tt.TestCase{
		{
			"Keeps boilerplate",
			expected("amp4ads"),
			expected("amp4ads"),
		},
		{
			"Keeps boilerplate for ⚡4ads",
			expected("⚡4ads"),
			expected("⚡4ads"),
		},
		{
			"Adds boilerplate if missing",
			tt.Concat("<!doctype html><html amp4ads><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, "</head><body></body></html>"),
			expected("amp4ads"),
		},
		{
			"Adds boilerplate for ⚡4ads if missing",
			tt.Concat("<!doctype html><html ⚡4ads><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, "</head><body></body></html>"),
			expected("⚡4ads"),
		},
	}

	runAMPBoilerplateTransformerTestcases(t, testCases)
}

func TestAMP4Email(t *testing.T) {
	expected := func(attr string) string {
		return tt.Concat("<!doctype html><html ", attr, "><head>",
			tt.ScriptAMPRuntime, tt.LinkFavicon,
			"<style amp4email-boilerplate>body{visibility:hidden}</style></head>",
			"<body></body></html>")
	}

	testCases := []tt.TestCase{
		{
			"Keeps boilerplate",
			expected("amp4email"),
			expected("amp4email"),
		},
		{
			"Keeps boilerplate for ⚡4email",
			expected("⚡4email"),
			expected("⚡4email"),
		},
		{
			"Adds boilerplate if missing",
			tt.Concat("<!doctype html><html amp4email><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, "</head><body></body></html>"),
			expected("amp4email"),
		},
		{
			"Adds boilerplate for ⚡4email if missing",
			tt.Concat("<!doctype html><html ⚡4email><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, "</head><body></body></html>"),
			expected("⚡4email"),
		},
	}

	runAMPBoilerplateTransformerTestcases(t, testCases)
}

func runAMPBoilerplateTransformerTestcases(t *testing.T, testCases []tt.TestCase) {
	for _, tc := range testCases {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s\nhtml.Parse for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		transformer.AMPBoilerplateTransformer(&transformer.Engine{Doc: inputDoc})

		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
			t.Errorf("%s\nhtml.Render for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}

		expectedDoc, err := html.Parse(strings.NewReader(tc.Expected))
		if err != nil {
			t.Errorf("%s\nhtml.Parse for %s failed %q", tc.Desc, tc.Expected, err)
			continue
		}
		var expected strings.Builder
		err = html.Render(&expected, expectedDoc)
		if err != nil {
			t.Errorf("%s\nhtml.Render for %s failed %q", tc.Desc, tc.Expected, err)
			continue
		}
		if input.String() != expected.String() {
			t.Errorf("%s: AMPBoilerplateTransformer=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
