package transformers_test

import (
	"strings"
	"testing"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	tt "github.com/ampproject/amppackager/transformer/internal/testing"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

func TestAmpRuntimeJS(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc:     "no script node",
			Input:    "<head></head><body></body>",
			Expected: "<head></head><body></body>",
		},
		{
			Desc:     "no prefix",
			Input:    "<head><script src=main.js/></head>",
			Expected: "<head><script src=main.js/></head>",
		},
		{
			Desc:     "no suffix",
			Input:    `<head><script src="https://cdn.ampproject.org"/></head>`,
			Expected: `<head><script src="https://cdn.ampproject.org"/></head>`,
		},
		{
			Desc:     "transformation",
			Input:    `<head><script async src="https://cdn.ampproject.org/v0.js"></script></head>`,
			Expected: `<head><script async src="https://cdn.ampproject.org/v0.js?f=sxg"></script></head>`,
		},
		{
			Desc:     "transform on two scripts",
			Input:    `<head><script async src="https://cdn.ampproject.org/foo.js"></script><script async src="https://cdn.ampproject.org/bar.js"></script></head>`,
			Expected: `<head><script async src="https://cdn.ampproject.org/foo.js?f=sxg"></script><script async src="https://cdn.ampproject.org/bar.js?f=sxg"></script></head>`,
		},
		{
			Desc:     "skip one, transform one",
			Input:    `<head><script async src="https://cdn.ampproject.org/"></script><script async src="https://cdn.ampproject.org/bar.js"></script></head>`,
			Expected: `<head><script async src="https://cdn.ampproject.org/"></script><script async src="https://cdn.ampproject.org/bar.js?f=sxg"></script></head>`,
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
		transformers.AMPRuntimeJS(&transformers.Context{DOM: inputDOM})
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
