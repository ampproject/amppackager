package transformers_test

import (
	"net/url"
	"strings"
	"testing"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

func TestURLRewrite_images(t *testing.T) {
	baseURL, _ := url.Parse("http://www.example.com")

	tcs := []struct {
		desc, input, expected, baseURL string
	}{
		{
			desc:     "in template noop",
			input:    `<template><amp-img src="foo" width="92" height="10" srcset="bar 50w"></template>`,
			expected: `<template><amp-img src="foo" width="92" height="10" srcset="bar 50w"></amp-img></template>`,
		},

		{
			desc:     "amp-img src and srcset rewritten",
			input:    "<amp-img src=http://www.example.com/blah.jpg width=92 height=10 srcset=\"http://www.example.com/blah.jpg 50w\">",
			expected: "<amp-img src=\"https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg\" width=\"92\" height=\"10\" srcset=\"https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 50w\">",
		},
		{
			desc:     "amp-img src and srcset rewritten with baseURL",
			input:    "<amp-img src=blah.jpg width=92 height=10 srcset=\"blah.jpg 50w\">",
			expected: "<amp-img src=\"https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg\" width=\"92\" height=\"10\" srcset=\"https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 50w\">",
		},
		{
			desc:     "amp-img srcset default density",
			input:    "<amp-img srcset=\"http://www.example.com/blah.jpg\">",
			expected: "<amp-img srcset=\"https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 1x\">",
		},
		{
			desc:     "amp-img srcset multiple",
			input:    "<amp-img srcset=\"image1 2x, image2, image3 3x, image4 4x \">",
			expected: "<amp-img srcset=\"https://www-example-com.cdn.ampproject.org/i/www.example.com/image1 2x, https://www-example-com.cdn.ampproject.org/i/www.example.com/image2 1x, https://www-example-com.cdn.ampproject.org/i/www.example.com/image3 3x, https://www-example-com.cdn.ampproject.org/i/www.example.com/image4 4x\">",
		},
		{
			desc:     "amp-img srcset whitespace",
			input:    "<amp-img srcset=\"  \t\n http://www.example.com/blah.jpg  \n\t\t \">",
			expected: "<amp-img srcset=\"https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 1x\">",
		},
		{
			desc:     "amp-img srcset leading comma",
			input:    "<amp-img srcset=\" , http://www.example.com/blah.jpg \n\t\t \">",
			expected: "<amp-img srcset=\"https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 1x\">",
		},
		{
			desc:     "amp-img srcset comma in url",
			input:    "<amp-img srcset=\" , http://www.example.com/image,1 \n\t\t \">",
			expected: "<amp-img srcset=\"https://www-example-com.cdn.ampproject.org/i/www.example.com/image,1 1x\">",
		},
		{
			desc:     "amp-img srcset commas everywhehre",
			input:    "<amp-img srcset=\",http://www.example.com/,/,/,/,50w,\">",
			expected: "<amp-img srcset=\"https://www-example-com.cdn.ampproject.org/i/www.example.com/,/,/,/,50w 1x\">",
		},
		{
			desc:     "amp-img srcset missing delimiter noop",
			input:    "<amp-img srcset=\"image1 100w image2 50w\">",
			expected: "<amp-img srcset=\"image1 100w image2 50w\">",
		},
		{
			desc:     "amp-img srcset negative value noop",
			input:    "<amp-img srcset=\"image1 100w, image2 -50w\">",
			expected: "<amp-img srcset=\"image1 100w, image2 -50w\">",
		},
	}

	for _, tc := range tcs {
		inputDoc, err := html.Parse(strings.NewReader(tc.input))
		if err != nil {
			t.Errorf("%s: html.Parse failed %q", tc.input, err)
			continue
		}
		inputDOM, err := amphtml.NewDOM(inputDoc)
		if err != nil {
			t.Errorf("%s\namphtml.NewDOM for %s failed %q", tc.desc, tc.input, err)
			continue
		}
		transformers.URLRewrite(&transformers.Context{DOM: inputDOM, BaseURL: baseURL})
		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
			t.Errorf("%s: html.Render failed %q", tc.input, err)
			continue
		}

		if !strings.Contains(input.String(), tc.expected) {
			t.Errorf("%s: URLRewrite=\n%q\ndoes not contain=\n%q", tc.desc, &input, tc.expected)
		}
	}
}
