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
			Desc: "Keeps boilerplate",
			Input: tt.Concat("<!doctype html><html ⚡><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body></body></html>"),
			Expected: canonicalExpected,
		},
		{
			Desc: "Adds boilerplate if missing",
			Input: tt.Concat("<!doctype html><html ⚡><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, "</head><body></body></html>"),
			Expected: canonicalExpected,
		},
		{
			Desc: "Upgrades old boilerplate",
			Input: tt.Concat("<!doctype html><html ⚡><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon,
				"<style>body {opacity: 0}</style>",
				"<noscript><style>body {opacity: 1}</style></noscript></head>",
				"<body></body></html>"),
			Expected: canonicalExpected,
		},
		{
			// The validator actually allows both old and new boilerplate to be present.
			// This test ensures we always strip multiple instances and end up
			// with just the new boilerplate.
			Desc: "Strips old and new if both present",
			Input: tt.Concat("<!doctype html><html ⚡><head>", tt.ScriptAMPRuntime,
				tt.NoscriptAMPBoilerplate, tt.LinkFavicon,
				"<style>body {opacity: 0}</style>", tt.StyleAMPBoilerplate,
				"<noscript><style>body {opacity: 1}</style></noscript></head>",
				"</head><body></body></html>"),
			Expected: canonicalExpected,
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
			Desc:     "Keeps boilerplate",
			Input:    expected("amp4ads"),
			Expected: expected("amp4ads"),
		},
		{
			Desc:     "Keeps boilerplate for ⚡4ads",
			Input:    expected("⚡4ads"),
			Expected: expected("⚡4ads"),
		},
		{
			Desc: "Adds boilerplate if missing",
			Input: tt.Concat("<!doctype html><html amp4ads><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, "</head><body></body></html>"),
			Expected: expected("amp4ads"),
		},
		{
			Desc: "Adds boilerplate for ⚡4ads if missing",
			Input: tt.Concat("<!doctype html><html ⚡4ads><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, "</head><body></body></html>"),
			Expected: expected("⚡4ads"),
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
			Desc:     "Keeps boilerplate",
			Input:    expected("amp4email"),
			Expected: expected("amp4email"),
		},
		{
			Desc:     "Keeps boilerplate for ⚡4email",
			Input:    expected("⚡4email"),
			Expected: expected("⚡4email"),
		},
		{
			Desc: "Adds boilerplate if missing",
			Input: tt.Concat("<!doctype html><html amp4email><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, "</head><body></body></html>"),
			Expected: expected("amp4email"),
		},
		{
			Desc: "Adds boilerplate for ⚡4email if missing",
			Input: tt.Concat("<!doctype html><html ⚡4email><head>", tt.ScriptAMPRuntime,
				tt.LinkFavicon, "</head><body></body></html>"),
			Expected: expected("⚡4email"),
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
		transformers.AMPBoilerplateTransformer(&transformers.Engine{Doc: inputDoc})

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
		if err := html.Render(&expected, expectedDoc); err != nil {
			t.Errorf("%s\nhtml.Render for %s failed %q", tc.Desc, tc.Expected, err)
			continue
		}
		if input.String() != expected.String() {
			t.Errorf("%s: AMPBoilerplateTransformer=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
