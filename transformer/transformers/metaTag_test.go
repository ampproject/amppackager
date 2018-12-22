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

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	tt "github.com/ampproject/amppackager/transformer/internal/testing"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

func TestMetaTag(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc: "Moves some meta tags",
			Input: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				"<meta content=experiment-a name=amp-experiments-opt-in>",
				"<meta content=experiment-b name=amp-experiments-opt-in>",
				tt.LinkFavicon, tt.LinkCanonical,
				tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"</head><body>",
				"<meta content=experiment-c name=amp-experiments-opt-in>", // moves to head
				"</body></html>"),
			Expected: tt.Concat(tt.Doctype, "<html ⚡><head>",
				tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
				"<meta content=experiment-a name=amp-experiments-opt-in>",
				"<meta content=experiment-b name=amp-experiments-opt-in>",
				tt.LinkFavicon, tt.LinkCanonical,
                                tt.StyleAMPBoilerplate, tt.NoscriptAMPBoilerplate,
				"<meta content=experiment-c name=amp-experiments-opt-in>",
				"</head><body></body></html>"),
		},
	}
	for _, tc := range tcs {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s: html.Parse on %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		inputDOM, err := amphtml.NewDOM(inputDoc)
		if err != nil {
			t.Errorf("%s\namphtml.NewDOM for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		transformers.MetaTag(&transformers.Context{DOM: inputDOM})

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
			t.Errorf("%s: MetaTag=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
