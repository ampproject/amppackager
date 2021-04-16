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
	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

func TestAMPRuntimeCSS(t *testing.T) {
	tcs := []struct{ desc, input, expected, rtv, css string }{
		{
			desc:  "no css or rtv",
			input: "<html><head></head></html>",
			expected: tt.Concat("<html><head>",
				"<style amp-runtime=\"\" i-amphtml-version=\"latest\"></style>",
				"</head><body></body></html>"),
			rtv: "",
			css: "",
		},
		{
			desc:  "no css",
			input: "<html><head></head></html>",
			expected: tt.Concat("<html><head>",
				"<style amp-runtime=\"\" i-amphtml-version=\"0142\"></style>",
				"</head><body></body></html>"),
			rtv: "0142",
			css: "",
		},
		{
			desc:  "inline css",
			input: "<html><head></head></html>",
			expected: tt.Concat("<html><head>",
				"<style amp-runtime=\"\" i-amphtml-version=\"0142\">CSS contents to inline</style>",
				"</head><body></body></html>"),
			rtv: "0142",
			css: "CSS contents to inline",
		},
		{
			desc:  "inline trimmed css",
			input: "<html><head></head></html>",
			expected: tt.Concat("<html><head>",
				"<style amp-runtime=\"\" i-amphtml-version=\"0142\">CSS contents to inline</style>",
				"</head><body></body></html>"),
			rtv: "0142",
			css: " \t\n CSS contents to inline \n\t\r\n",
		},
		{
			desc: "keeps existing <style amp-runtime>s if no local CSS",
			input: tt.Concat("<html><head>",
				"<style amp-runtime i-amphtml-version=0141>CSS contents to inline</style>",
				"<style amp-runtime i-amphtml-version=0141></style>",
				"</head></html>"),
			expected: tt.Concat("<html><head>",
				"<style amp-runtime=\"\" i-amphtml-version=\"0141\">CSS contents to inline</style>",
				"<style amp-runtime=\"\" i-amphtml-version=\"0141\"></style>",
				"</head><body></body></html>"),
			rtv: "0142",
			css: "",
		},
		{
			desc: "keeps existing <style amp-runtime>s if any newer",
			input: tt.Concat("<html><head>",
				"<style amp-runtime i-amphtml-version=0141>CSS contents to inline</style>",
				"<style amp-runtime i-amphtml-version=0152></style>",
				"</head></html>"),
			expected: tt.Concat("<html><head>",
				"<style amp-runtime=\"\" i-amphtml-version=\"0141\">CSS contents to inline</style>",
				"<style amp-runtime=\"\" i-amphtml-version=\"0152\"></style>",
				"</head><body></body></html>"),
			rtv: "0142",
			css: "CSS contents to inline",
		},
		{
			desc: "removes existing <style amp-runtime>s if older",
			input: tt.Concat("<html><head>",
				"<style amp-runtime i-amphtml-version=0141>CSS contents to inline</style>",
				"<style amp-runtime i-amphtml-version=0141></style>",
				"</head></html>"),
			expected: tt.Concat("<html><head>",
				"<style amp-runtime=\"\" i-amphtml-version=\"0142\">CSS contents to inline</style>",
				"</head><body></body></html>"),
			rtv: "0142",
			css: "CSS contents to inline",
		},
		{
			desc: "removes existing <style amp-runtime>s if atypical",
			input: tt.Concat("<html><head>",
				"<style amp-runtime=0252>CSS contents to inline</style>",
				"<style amp-runtime=0252></style>",
				"</head></html>"),
			expected: tt.Concat("<html><head>",
				"<style amp-runtime=\"\" i-amphtml-version=\"0142\">CSS contents to inline</style>",
				"</head><body></body></html>"),
			rtv: "0142",
			css: "CSS contents to inline",
		},
	}

	for _, tc := range tcs {
		inputDoc, err := html.Parse(strings.NewReader(tc.input))
		if err != nil {
			t.Errorf("%s: html.Parse on %s failed %q", tc.desc, tc.input, err)
			continue
		}
		inputDOM, err := amphtml.NewDOM(inputDoc)
		if err != nil {
			t.Errorf("%s\namphtml.NewDOM for %s failed %q", tc.desc, tc.input, err)
			continue
		}
		transformers.AMPRuntimeCSS(&transformers.Context{DOM: inputDOM, Request: &rpb.Request{Rtv: tc.rtv, Css: tc.css}})
		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
			t.Errorf("%s: html.Render on %s failed %q", tc.desc, tc.input, err)
			continue
		}

		if input.String() != tc.expected {
			t.Errorf("%s: AMPRuntimeCSS=\n%q\nwant=\n%q", tc.desc, &input, tc.expected)
		}
	}
}
