package transformer_test

import (
	"strings"
	"testing"

	tt "github.com/ampproject/amppackager/internal/pkg/testing"
	"github.com/ampproject/amppackager/pkg/transformer"
	"golang.org/x/net/html"
)

// These tests do NOT run through the custom transformations of the
// Engine, and instead rely exclusively on vanilla golang parser and
// renderer (otherwise the scope of these tests would creep past unit
// testing). Therefore, the test data must be made to match, and is not
// the expected normalized output from transformer.go, nor from any other
// transformers.

func TestTransformedIdentifierTransformer(t *testing.T) {
	testCases := []tt.TestCase{
		{
			"Adds identifier to html tag",
			tt.Concat("<!doctype html><html ⚡><head>",
				tt.ScriptAMPRuntime, tt.MetaCharset, tt.StyleAMPBoilerplate,
				tt.NoscriptAMPBoilerplate, "</head><body></body></html>"),
			tt.Concat("<!doctype html><html ⚡=\"\" transformed=google><head>",
				tt.ScriptAMPRuntime, tt.MetaCharset, tt.StyleAMPBoilerplate,
				tt.NoscriptAMPBoilerplate, "</head><body></body></html>"),
		},
	}
	runTransformedIdentifierTransformerTestcases(t, testCases)
}

func runTransformedIdentifierTransformerTestcases(t *testing.T, testCases []tt.TestCase) {
	for _, tc := range testCases {
		inputDoc, err := html.Parse(strings.NewReader(tc.Input))
		if err != nil {
			t.Errorf("%s: html.Parse for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}
		transformer.TransformedIdentifierTransformer(&transformer.Engine{Doc: inputDoc})

		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
			t.Errorf("%s: html.Render for %s failed %q", tc.Desc, tc.Input, err)
			continue
		}

		expectedDoc, err := html.Parse(strings.NewReader(tc.Expected))
		if err != nil {
			t.Errorf("%s: html.Parse for %s failed %q", tc.Desc, tc.Expected, err)
			continue
		}
		var expected strings.Builder
		err = html.Render(&expected, expectedDoc)
		if err != nil {
			t.Errorf("%s: html.Render for %s failed %q", tc.Desc, tc.Expected, err)
			continue
		}
		if input.String() != expected.String() {
			t.Errorf("%s: TransformedIdentifierTransformer=\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
