package transformer_test

import (
	"strings"
	"testing"

	tt "github.com/ampproject/amppackager/internal/pkg/testing"
	rpb "github.com/ampproject/amppackager/pkg/transform
	"github.com/ampproject/amppackager/pkg/transformer"
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
	}

	for _, tc := range tcs {
		inputDoc, err := html.Parse(strings.NewReader(tc.input))
		if err != nil {
			t.Errorf("%s: html.Parse on %s failed %q", tc.desc, tc.input, err)
			continue
		}
		transformer.AMPRuntimeCSSTransformer(&transformer.Engine{Doc: inputDoc, Request: &rpb.Request{Rtv: "42", Css: tc.css}})
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
