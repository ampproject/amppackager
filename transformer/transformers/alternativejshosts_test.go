// Copyright 2022 Google LLC
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

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	tt "github.com/ampproject/amppackager/transformer/internal/testing"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

func TestAlternativeJsHosts(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc: "Do Nothing To Cdn Js",
			Input: tt.Concat("<html><head>",
				"<script async src=https://cdn.ampproject.org/v0.js></script>",
				"<script async src=https://cdn.ampproject.org/amp4ads-v0.js></script>",
				"<script async src=https://cdn.ampproject.org/v0/amp-ad-exit-0.1.js>",
				"</script>",
				"</head><body></body></html>"),
			Expected: tt.Concat("<html><head>",
				"<script async src=https://cdn.ampproject.org/v0.js></script>",
				"<script async src=https://cdn.ampproject.org/amp4ads-v0.js></script>",
				"<script async src=https://cdn.ampproject.org/v0/amp-ad-exit-0.1.js>",
				"</script>",
				"</head><body></body></html>"),
		},
		{
			Desc: "Do Nothing To Cdn mjs",
			Input: tt.Concat("<html><head>",
				"<script async src=https://cdn.ampproject.org/v0.mjs></script>",
				"<script async src=https://cdn.ampproject.org/amp4ads-v0.mjs></script>",
				"<script async src=https://cdn.ampproject.org/v0/amp-ad-exit-0.1.mjs>",
				"</script>",
				"</head><body></body></html>"),
			Expected: tt.Concat("<html><head>",
				"<script async src=https://cdn.ampproject.org/v0.mjs></script>",
				"<script async src=https://cdn.ampproject.org/amp4ads-v0.mjs></script>",
				"<script async src=https://cdn.ampproject.org/v0/amp-ad-exit-0.1.mjs>",
				"</script>",
				"</head><body></body></html>"),
		},
		{
			Desc: "Do Nothing To Unexpected Js",
			Input: tt.Concat("<html><head>",
				"<script async src=https://site.example/unexpected/script.js></script>",
				"</head><body></body></html>"),
			Expected: tt.Concat("<html><head>",
				"<script async src=https://site.example/unexpected/script.js></script>",
				"</head><body></body></html>"),
		},
		{
			Desc: "Do Nothing To Upper Case Paths",
			Input: tt.Concat("<html><head>",
				"<script async src=https://site.example/V0.JS></script>",
				"</head><body></body></html>"),
			Expected: tt.Concat("<html><head>",
				"<script async src=https://site.example/V0.JS></script>",
				"</head><body></body></html>"),
		},
		{
			Desc: "Rewrites AmpJs To AmpProject",
			Input: tt.Concat("<html><head>",
				"<script async src=https://ampjs.org/v0.js></script>",
				"<script async src=https://ampjs.org/amp4ads-v0.js></script>",
				"<script async src=https://ampjs.org/v0/amp-ad-exit-0.1.js>",
				"</script>",
				"</head><body></body></html>"),
			Expected: tt.Concat("<html><head>",
				"<script async src=https://cdn.ampproject.org/v0.js></script>",
				"<script async src=https://cdn.ampproject.org/amp4ads-v0.js></script>",
				"<script async src=https://cdn.ampproject.org/v0/amp-ad-exit-0.1.js>",
				"</script>",
				"</head><body></body></html>"),
		},
		{
			Desc: "Rewrites Self Host To AmpProject",
			Input: tt.Concat("<html><head>",
				"<script async src=https://selfhost.com/v0.js></script>",
				"<script async src=https://selfhost.com/amp4ads-v0.js></script>",
				"<script async src=https://selfhost.com/v0/amp-ad-exit-0.1.js>",
				"</script>",
				"</head><body></body></html>"),
			Expected: tt.Concat("<html><head>",
				"<script async src=https://cdn.ampproject.org/v0.js></script>",
				"<script async src=https://cdn.ampproject.org/amp4ads-v0.js></script>",
				"<script async src=https://cdn.ampproject.org/v0/amp-ad-exit-0.1.js>",
				"</script>",
				"</head><body></body></html>"),
		},
	}

	runAlternativeJsHostsTestcases(t, tcs)
}

func runAlternativeJsHostsTestcases(t *testing.T, tcs []tt.TestCase) {
	for _, tc := range tcs {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s\nhtml.Parse for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		inputDOM, err := amphtml.NewDOM(inputDoc)
		if err != nil {
			t.Errorf("%s\namphtml.NewDOM for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		transformers.AlternativeJsHosts(&transformers.Context{DOM: inputDOM})

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
			t.Errorf("%s: AlternativeJsHosts=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
