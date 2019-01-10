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

func TestStripJS(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc:     "strips script wrong src",
			Input:    "<script src=main.js/>",
			Expected: "<head></head><body></body>",
		},
		{
			Desc:     "keeps script correct src",
			Input:    "<script async src=\"https://cdn.ampproject.org/v0.js\"></script>",
			Expected: "<head><script async src=https://cdn.ampproject.org/v0.js></script></head>",
		},
		{
			Desc:     "keeps script correct src case-insensitive",
			Input:    "<script async custom-element='amp-analytics' src='https://cDn.AMPproject.org/v0/amp-analytics-0.1.js'></script>",
			Expected: "<head><script async custom-element=amp-analytics src=https://cDn.AMPproject.org/v0/amp-analytics-0.1.js></script></head>",
		},
		{
			Desc:     "keeps script correct src with type",
			Input:    "<script async src=\"https://cdn.ampproject.org/v0.js\" type=\"text/javascript\"></script>",
			Expected: "<head><script async src=https://cdn.ampproject.org/v0.js type=text/javascript></script></head>",
		},
		{
			Desc:     "strips script no src, no type",
			Input:    "<script>foo</script>",
			Expected: "<head></head><body></body>",
		},
		{
			Desc:     "strips script wrong type",
			Input:    "<script type=application/javascript>foo</script>",
			Expected: "<head></head><body></body>",
		},
		{
			Desc:     "keeps script corect type",
			Input:    "<script type=application/json>foo</script>",
			Expected: "<head><script type=application/json>foo</script></head><body></body>",
		},
		{
			Desc:     "strip tag attr ona",
			Input:    "<body><select ona=\"myFunction()\"></body>",
			Expected: "<body><select></select></body>",
		},
		{
			Desc:     "strips tag event attr",
			Input:    "<body><select onchange=\"myFunction()\"></body>",
			Expected: "<body><select></select></body>",
		},
		{
			Desc:     "strip tag attr onfoo",
			Input:    "<body><select onfoo=\"myFunction()\"></body>",
			Expected: "<body><select></select></body>",
		},
		{
			Desc:     "keep tag attr 'on'",
			Input:    "<body><select on=\"myFunction()\"></body>",
			Expected: "<body><select on=myFunction()></select></body>",
		},
		{
			Desc:     "keep tag attr on-foo",
			Input:    "<body><select on-foo=\"myFunction()\"></body>",
			Expected: "<body><select on-foo=myFunction()></select></body>",
		},
		{
			Desc:     "keep tag attr notonchange",
			Input:    "<body><select notonchange=\"myFunction()\"></body>",
			Expected: "<body><select notonchange=myFunction()></select></body>",
		},
	}

	for _, tc := range tcs {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s: html.Parse failed %q", tc.Input, err)
			continue
		}
		inputDOM, err := amphtml.NewDOM(inputDoc)
		if err != nil {
			t.Errorf("%s\namphtml.NewDOM for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		transformers.StripJS(&transformers.Context{DOM: inputDOM})
		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
			t.Errorf("%s: html.Render failed %q", tc.Input, err)
			continue
		}

		expectedDoc, err := html.Parse(strings.NewReader(tc.Expected))
		if err != nil {
			t.Errorf("%s: html.Parse failed %q", tc.Expected, err)
			continue
		}
		var expected strings.Builder
		err = html.Render(&expected, expectedDoc)
		if err != nil {
			t.Errorf("%s: html.Render failed %q", tc.Expected, err)
			continue
		}
		if input.String() != expected.String() {
			t.Errorf("%s: Transform=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}

}
