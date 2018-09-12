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
	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

func TestAMPRuntimeCSSTransformer(t *testing.T) {
	tcs := []struct{ desc, input, expected, css string }{
		{
			desc:     "Empty doc",
			input:    "",
			expected: "<html><head></head><body></body></html>",
			css:      "",
		},
		{
			desc:     "no ssr",
			input:    "<html></html>",
			expected: "<html><head></head><body></body></html>",
			css:      "",
		},
		{
			desc:  "link to css",
			input: "<html><head><style amp-runtime></style></head></html>",
			expected: tt.Concat("<html><head>",
				"<style amp-runtime=\"\" i-amphtml-version=\"42\"></style>",
				"<link rel=\"stylesheet\" href=\"https://cdn.ampproject.org/rtv/42/v0.css\"/>",
				"</head><body></body></html>"),
			css: "",
		},
		{
			desc:  "inline css",
			input: "<html><head><style amp-runtime></style></head></html>",
			expected: tt.Concat("<html><head>",
				"<style amp-runtime=\"\" i-amphtml-version=\"42\">",
				"CSS contents to inline</style></head>",
				"<body></body></html>"),
			css: "CSS contents to inline",
		},
		{
			desc:  "trim css",
			input: "<html><head><style amp-runtime></style></head></html>",
			expected: tt.Concat("<html><head>",
				"<style amp-runtime=\"\" i-amphtml-version=\"42\">",
				"CSS contents to inline</style></head>",
				"<body></body></html>"),
			css: " \t\n CSS contents to inline \n\t\r\n",
		},
	}

	for _, tc := range tcs {
		inputDoc, err := html.Parse(strings.NewReader(tc.input))
		if err != nil {
			t.Errorf("%s: html.Parse on %s failed %q", tc.desc, tc.input, err)
			continue
		}
		transformers.AMPRuntimeCSSTransformer(&transformers.Engine{Doc: inputDoc, Request: &rpb.Request{Rtv: "42", Css: tc.css}})
		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
			t.Errorf("%s: html.Render on %s failed %q", tc.desc, tc.input, err)
			continue
		}

		if input.String() != tc.expected {
			t.Errorf("%s: AMPRuntimeCSSTransformer=\n%q\nwant=\n%q", tc.desc, &input, tc.expected)
		}
	}
}
