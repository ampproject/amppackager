package transformers_test

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	tt "github.com/ampproject/amppackager/transformer/internal/testing"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

func TestURLRewrite_images(t *testing.T) {
	baseTcs := []tt.TestCase{
		{
			Desc:     "%s in template noop",
			Input:    `<template><%s src="foo" width="92" height="10" srcset="bar 50w"></template>`,
			Expected: `<template><%s src="foo" width="92" height="10" srcset="bar 50w"></%s></template>`,
		},
		{
			Desc:     "%s src and srcset rewritten",
			Input:    `<%s src=http://www.example.com/blah.jpg width=92 height=10 srcset="http://www.example.com/blah.jpg 50w">`,
			Expected: `<%s src="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg" width="92" height="10" srcset="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 50w"></%s>`,
		},
		{
			Desc:     "%s does not add srcset with no width",
			Input:    `<%s src=http://www.example.com/blah.jpg height=10>`,
			Expected: `<%s src="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg" height="10"></%s>`,
		},
		{
			Desc:     "%s does not add srcset with 0 width",
			Input:    `<%s src=http://www.example.com/blah.jpg height=10 width=0>`,
			Expected: `<%s src="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg" height="10" width="0"></%s>`,
		},
		{
			Desc:     "%s adds srcset",
			Input:    `<%s src=http://www.example.com/blah.jpg width=92 height=10>`,
			Expected: `<%s src="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg" width="92" height="10" srcset="https://www-example-com.cdn.ampproject.org/ii/w100/www.example.com/blah.jpg 100w, https://www-example-com.cdn.ampproject.org/ii/w220/www.example.com/blah.jpg 220w, https://www-example-com.cdn.ampproject.org/ii/w330/www.example.com/blah.jpg 330w"></%s>`,
		},
		{
			Desc:     "%s src and srcset rewritten with baseURL",
			Input:    `<%s src=blah.jpg width=92 height=10 srcset="blah.jpg 50w">`,
			Expected: `<%s src="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg" width="92" height="10" srcset="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 50w"></%s>`,
		},
		{
			Desc:     "%s data:image noop",
			Input:    `<%s src="data:image/png,foo">`,
			Expected: `<%s src="data:image/png,foo"></%s>`,
		},
	}
	tcs := []tt.TestCase{}
	for _, tag := range []string{"amp-img", "amp-anim"} {
		for _, baseTc := range baseTcs {
			tc := tt.TestCase{
				Desc:     fmt.Sprintf(baseTc.Desc, tag),
				Input:    fmt.Sprintf(baseTc.Input, tag),
				Expected: fmt.Sprintf(baseTc.Expected, tag, tag),
			}
			tcs = append(tcs, tc)
		}
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_link(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc:     "link rel=icon in template noop",
			Input:    `<template><link rel=icon href=foo></template>`,
			Expected: `<template><link rel="icon" href="foo"/></template>`,
		},
		{
			Desc:     "link rel=shortcuticon noop",
			Input:    `<link rel=shortcuticon href=foo>`,
			Expected: `<link rel="shortcuticon" href="foo"/>`,
		},
		{
			Desc:     "link rel=notsupported noop",
			Input:    `<link rel=notsupported href=foo>`,
			Expected: `<link rel="notsupported" href="foo"/>`,
		},
		{
			Desc:     "link rel=icon",
			Input:    `<link rel=icon href=foo>`,
			Expected: `<link rel="icon" href="https://www-example-com.cdn.ampproject.org/i/www.example.com/foo"/>`,
		},
		{
			Desc:     `link rel="shortcut icon"`,
			Input:    `<link rel="shortcut icon" href=foo>`,
			Expected: `<link rel="shortcut icon" href="https://www-example-com.cdn.ampproject.org/i/www.example.com/foo"/>`,
		},
		{
			Desc:     `link rel="icon shortcut"`,
			Input:    `<link rel="icon shortcut" href=foo>`,
			Expected: `<link rel="icon shortcut" href="https://www-example-com.cdn.ampproject.org/i/www.example.com/foo"/>`,
		},
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_background(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc:     "background",
			Input:    `<body background=https://leak.com></body>`,
			Expected: `<body background="https://leak-com.cdn.ampproject.org/i/s/leak.com">`,
		},
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_svgImage(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc:     "in template noop",
			Input:    `<template><svg><image xlink:href=https://leak.com></image></svg></template>`,
			Expected: `<template><svg><image xlink:href="https://leak.com"></image></svg></template>`,
		},
		{
			Desc:     "xlink:href",
			Input:    `<svg><image xlink:href=https://leak.com></image></svg>`,
			Expected: `<svg><image xlink:href="https://leak-com.cdn.ampproject.org/i/s/leak.com"></image></svg>`,
		},
		{
			Desc:     "href",
			Input:    `<svg><image href=https://leak.com></image></svg>`,
			Expected: `<svg><image href="https://leak-com.cdn.ampproject.org/i/s/leak.com"></image></svg>`,
		},
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_svgUse(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc:     "in template noop",
			Input:    `<template><use xlink:href=https://leak.com></use></template>`,
			Expected: `<template><use xlink:href="https://leak.com"></use></template>`,
		},
		{
			Desc:     "rewrite",
			Input:    `<svg><use xlink:href=https://leak.com></use></svg>`,
			Expected: `<svg><use xlink:href="https://leak-com.cdn.ampproject.org/i/s/leak.com"></use></svg>`,
		},
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_poster(t *testing.T) {
	baseTcs := []tt.TestCase{
		{
			Desc:     "%s in template noop",
			Input:    `<template><%s poster="foo"></template>`,
			Expected: `<template><%s poster="foo"></%s></template>`,
		},
		{
			Desc:     "%s rewrite",
			Input:    `<%s poster="foo">`,
			Expected: `<%s poster="https://www-example-com.cdn.ampproject.org/i/www.example.com/foo"></%s>`,
		},
	}
	tcs := []tt.TestCase{}
	for _, tag := range []string{"amp-video", "video"} {
		for _, baseTc := range baseTcs {
			tc := tt.TestCase{
				Desc:     fmt.Sprintf(baseTc.Desc, tag),
				Input:    fmt.Sprintf(baseTc.Input, tag),
				Expected: fmt.Sprintf(baseTc.Expected, tag, tag),
			}
			tcs = append(tcs, tc)
		}
	}
	runURLRewriteTestcases(t, tcs)
}

func runURLRewriteTestcases(t *testing.T, tcs []tt.TestCase) {
	baseURL, _ := url.Parse("http://www.example.com")

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
		transformers.URLRewrite(&transformers.Context{DOM: inputDOM, BaseURL: baseURL})
		var input strings.Builder
		if err := html.Render(&input, inputDoc); err != nil {
			t.Errorf("%s: html.Render failed %q", tc.Input, err)
			continue
		}

		if !strings.Contains(input.String(), tc.Expected) {
			t.Errorf("%s: URLRewrite=\n%q\ndoes not contain=\n%q", tc.Desc, &input, tc.Expected)
		}
	}
}
