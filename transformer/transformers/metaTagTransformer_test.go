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
// the expected normalized output from transformer.go, nor from any other
// transformers.

func TestMetaTagTransformer(t *testing.T) {
	testCases := []tt.TestCase{
		{
			Desc: "Strips some meta tags",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.ScriptAMPRuntime, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				tt.MetaCharset, tt.MetaViewport,
				"<meta http-equiv=x-dns-prefetch-control>", // gets stripped
				"<meta content=com.nytimes.com:basic itemprop=productID>",
				"<meta itemprop=productID name=nytimes>", // gets stripped
				"<meta name=Author content=lorem>",       // gets stripped
				"<meta content=experiment-a name=amp-experiments-opt-in>",
				"<meta name=robots content=index>", // gets stripped
				"<meta property=rendition:spread>", // gets stripped
				"<meta as=script href=v0.js rel=preload>",
				"</head><body></body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡><head>",
				tt.ScriptAMPRuntime, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				tt.MetaCharset, tt.MetaViewport,
				"<meta content=com.nytimes.com:basic itemprop=productID>",
				"<meta content=experiment-a name=amp-experiments-opt-in>",
				"<meta as=script href=v0.js rel=preload>",
				"</head><body></body></html>"),
		},
		{
			Desc: "Strips and moves some meta tags",
			Input: tt.Concat("<!doctype html><html ⚡><head>",
				tt.ScriptAMPRuntime, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				tt.MetaCharset, tt.MetaViewport,
				"<meta name=author content=lorem>", // gets stripped
				"<meta content=experiment-a name=amp-experiments-opt-in>",
				"<meta content=experiment-b name=amp-experiments-opt-in>",
				"<meta property=rendition:spread>", // gets stripped
				"</head><body>",
				"<meta name=author content=ipsum>",                        // gets stripped
				"<meta content=experiment-c name=amp-experiments-opt-in>", // moves to head
				"</body></html>"),
			Expected: tt.Concat("<!doctype html><html ⚡><head>",
				tt.ScriptAMPRuntime, tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				tt.MetaCharset, tt.MetaViewport,
				"<meta content=experiment-a name=amp-experiments-opt-in>",
				"<meta content=experiment-b name=amp-experiments-opt-in>",
				"<meta content=experiment-c name=amp-experiments-opt-in>",
				"</head><body></body></html>"),
		},
	}
	for _, tc := range testCases {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s: html.Parse on %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		transformers.MetaTagTransformer(&transformers.Engine{Doc: inputDoc})

		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
			t.Errorf("%s: html.Render on %s failed %q", tc.Desc, tc.Input, err)
			continue
		}

		expectedDoc, err := html.Parse(strings.NewReader(tc.Expected))
		if err != nil {
			t.Errorf("%s: html.Parse for %s failed %q", tc.Desc, tc.Expected, err)
			continue
		}
		var expected strings.Builder
		if err := html.Render(&expected, expectedDoc); err != nil {
			t.Errorf("%s: html.Render for %s failed %q", tc.Desc, tc.Expected, err)
			continue
		}
		if input.String() != expected.String() {
			t.Errorf("%s: MetaTagTransformer=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
