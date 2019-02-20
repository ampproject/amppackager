package transformers_test

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/transformers"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

type urlRewriteTestCase struct {
	desc, input, expected, base string
}

func TestURLRewrite_images(t *testing.T) {
	baseTcs := []urlRewriteTestCase{
		{
			desc:     "%s in template noop",
			input:    `<template><%s src="foo" width="92" height="10" srcset="bar 50w"></template>`,
			expected: `<template><%s src="foo" width="92" height="10" srcset="bar 50w"></%s></template>`,
		},
		{
			desc:     "%s src and srcset rewritten",
			input:    `<%s src=http://www.example.com/blah.jpg width=92 height=10 srcset="http://www.example.com/blah.jpg 50w">`,
			expected: `<%s src="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg" width="92" height="10" srcset="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 50w"></%s>`,
		},
		{
			desc:     "%s does not add srcset with no width",
			input:    `<%s src=http://www.example.com/blah.jpg height=10>`,
			expected: `<%s src="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg" height="10"></%s>`,
		},
		{
			desc:     "%s does not add srcset with 0 width",
			input:    `<%s src=http://www.example.com/blah.jpg height=10 width=0>`,
			expected: `<%s src="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg" height="10" width="0"></%s>`,
		},
		{
			desc:     "%s adds srcset",
			input:    `<%s src=http://www.example.com/blah.jpg width=92 height=10>`,
			expected: `<%s src="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg" width="92" height="10" srcset="https://www-example-com.cdn.ampproject.org/ii/w100/www.example.com/blah.jpg 100w, https://www-example-com.cdn.ampproject.org/ii/w220/www.example.com/blah.jpg 220w, https://www-example-com.cdn.ampproject.org/ii/w330/www.example.com/blah.jpg 330w"></%s>`,
		},
		{
			desc:     "%s adds srcset if empty",
			input:    `<%s src=http://www.example.com/blah.jpg width=92 height=10 srcset="">`,
			expected: `<%s src="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg" width="92" height="10" srcset="https://www-example-com.cdn.ampproject.org/ii/w100/www.example.com/blah.jpg 100w, https://www-example-com.cdn.ampproject.org/ii/w220/www.example.com/blah.jpg 220w, https://www-example-com.cdn.ampproject.org/ii/w330/www.example.com/blah.jpg 330w"></%s>`,
		},
		{
			desc:     "%s adds srcset no base",
			input:    `<%s src=blah.jpg width=42 height=42>`,
			expected: `<%s src="blah.jpg" width="42" height="42" srcset="blah.jpg 47w, blah.jpg 100w, blah.jpg 150w"></%s>`,
		},
		{
			desc:     "%s preserve srcset",
			input:    `<%s srcset="image-1x.png 1x, image-2x.png 2x" layout="fill">`,
			expected: `<%s srcset="image-1x.png 1x, image-2x.png 2x" layout="fill"></%s>`,
		},
		{
			desc:     "%s src and srcset rewritten with baseURL",
			input:    `<%s src=blah.jpg width=92 height=10 srcset="blah.jpg 50w">`,
			expected: `<%s src="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg" width="92" height="10" srcset="https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 50w"></%s>`,
			base:     "http://www.example.com",
		},
		{
			desc:     "%s data:image noop",
			input:    `<%s src="data:image/png,foo">`,
			expected: `<%s src="data:image/png,foo"></%s>`,
		},
		{
			desc:     "%s empty src noop",
			input:    `<%s src="">`,
			expected: `<%s src=""></%s>`,
		},
	}
	tcs := []urlRewriteTestCase{}
	for _, tag := range []string{"amp-img", "amp-anim"} {
		for _, baseTc := range baseTcs {
			tc := urlRewriteTestCase{
				desc:     fmt.Sprintf(baseTc.desc, tag),
				input:    fmt.Sprintf(baseTc.input, tag),
				expected: fmt.Sprintf(baseTc.expected, tag, tag),
				base:     baseTc.base,
			}
			tcs = append(tcs, tc)
		}
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_link(t *testing.T) {
	tcs := []urlRewriteTestCase{
		{
			desc:     "link rel=icon in template noop",
			input:    `<template><link rel=icon href=foo></template>`,
			expected: `<template><link rel="icon" href="foo"/></template>`,
		},
		{
			desc:     "link rel=shortcuticon noop",
			input:    `<link rel=shortcuticon href=foo>`,
			expected: `<link rel="shortcuticon" href="foo"/>`,
		},
		{
			desc:     "link rel=notsupported noop",
			input:    `<link rel=notsupported href=foo>`,
			expected: `<link rel="notsupported" href="foo"/>`,
		},
		{
			desc:     "link rel=icon",
			input:    `<link rel=icon href=foo>`,
			expected: `<link rel="icon" href="https://www-example-com.cdn.ampproject.org/i/www.example.com/foo"/>`,
			base:     "http://www.example.com",
		},
		{
			desc:     "fragment",
			input:    `<link rel=icon href=foo#bar>`,
			expected: `<link rel="icon" href="https://www-example-com.cdn.ampproject.org/i/www.example.com/foo#bar"/>`,
			base:     "http://www.example.com",
		},
		{
			desc:     `link rel="shortcut icon"`,
			input:    `<link rel="shortcut icon" href=foo>`,
			expected: `<link rel="shortcut icon" href="https://www-example-com.cdn.ampproject.org/i/www.example.com/foo"/>`,
			base:     "http://www.example.com",
		},
		{
			desc:     `link rel="icon shortcut"`,
			input:    `<link rel="icon shortcut" href=foo>`,
			expected: `<link rel="icon shortcut" href="https://www-example-com.cdn.ampproject.org/i/www.example.com/foo"/>`,
			base:     "http://www.example.com",
		},
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_background(t *testing.T) {
	tcs := []urlRewriteTestCase{
		{
			desc:     "insecure",
			input:    `<body background=http://leak.com></body>`,
			expected: `<body background="https://leak-com.cdn.ampproject.org/i/leak.com">`,
		},
		{
			desc:     "secure ",
			input:    `<body background=https://leak.com></body>`,
			expected: `<body background="https://leak-com.cdn.ampproject.org/i/s/leak.com">`,
		},
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_svgImage(t *testing.T) {
	tcs := []urlRewriteTestCase{
		{
			desc:     "in template noop",
			input:    `<template><svg><image xlink:href=https://leak.com></image></svg></template>`,
			expected: `<template><svg><image xlink:href="https://leak.com"></image></svg></template>`,
		},
		{
			desc:     "xlink:href",
			input:    `<svg><image xlink:href=https://leak.com></image></svg>`,
			expected: `<svg><image xlink:href="https://leak-com.cdn.ampproject.org/i/s/leak.com"></image></svg>`,
		},
		{
			desc:     "href",
			input:    `<svg><image href=https://leak.com></image></svg>`,
			expected: `<svg><image href="https://leak-com.cdn.ampproject.org/i/s/leak.com"></image></svg>`,
		},
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_svgUse(t *testing.T) {
	tcs := []urlRewriteTestCase{
		{
			desc:     "in template noop",
			input:    `<template><use xlink:href=https://leak.com></use></template>`,
			expected: `<template><use xlink:href="https://leak.com"></use></template>`,
		},
		{
			desc:     "rewrite",
			input:    `<svg><use xlink:href=https://leak.com></use></svg>`,
			expected: `<svg><use xlink:href="https://leak-com.cdn.ampproject.org/i/s/leak.com"></use></svg>`,
		},
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_poster(t *testing.T) {
	baseTcs := []urlRewriteTestCase{
		{
			desc:     "%s in template noop",
			input:    `<template><%s poster="foo"></template>`,
			expected: `<template><%s poster="foo"></%s></template>`,
		},
		{
			desc:     "%s rewrite",
			input:    `<%s poster="foo">`,
			expected: `<%s poster="https://www-example-com.cdn.ampproject.org/i/www.example.com/foo"></%s>`,
			base:     "http://www.example.com",
		},
	}
	tcs := []urlRewriteTestCase{}
	for _, tag := range []string{"amp-video", "video"} {
		for _, baseTc := range baseTcs {
			tc := urlRewriteTestCase{
				desc:     fmt.Sprintf(baseTc.desc, tag),
				input:    fmt.Sprintf(baseTc.input, tag),
				expected: fmt.Sprintf(baseTc.expected, tag, tag),
				base:     baseTc.base,
			}
			tcs = append(tcs, tc)
		}
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_preconnect(t *testing.T) {
	tcs := []urlRewriteTestCase{
		{
			desc:     "preconnects added and sorted",
			input:    `<amp-img src=http://notexample.com/blah.jpg width=92 height=10 srcset="http://alsonotexample.com/blah.jpg 50w">`,
			expected: `<head><link href="https://alsonotexample-com.cdn.ampproject.org" rel="dns-prefetch preconnect"/><link href="https://notexample-com.cdn.ampproject.org" rel="dns-prefetch preconnect"/></head>`,
			base:     "http://www.example.com",
		},
		{
			desc:     "does not add duplicate preconnects",
			input:    `<amp-img src=http://notexample.com/blah.jpg width=92 height=10 srcset="http://notexample.com/another.jpg 50w">`,
			expected: `<head><link href="https://notexample-com.cdn.ampproject.org" rel="dns-prefetch preconnect"/></head>`,
			base:     "http://www.example.com",
		},
		{
			desc:     "keeps preconnect intact if in head already",
			input:    `<head><link crossorigin href="https://notexample-com.cdn.ampproject.org" rel="dns-prefetch preconnect"/></head><body><amp-img src=http://notexample.com/blah.jpg width=92 height=10 srcset="http://notexample.com/another.jpg 50w">`,
			expected: `<head><link crossorigin="" href="https://notexample-com.cdn.ampproject.org" rel="dns-prefetch preconnect"/></head>`,
			base:     "http://www.example.com",
		},
		{
			desc:     "does not add unnecessary preconnect",
			input:    `<amp-img src=http://www.example.com/blah.jpg width=92 height=10 srcset="http://www.example.com/blah.jpg 50w">`,
			expected: `<head></head>`,
			base:     "http://www.example.com",
		},
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_style(t *testing.T) {
	baseTcs := []struct{ desc, input, replacement string }{
		{
			desc: "stylesheet with images",
			input: "<head><style amp-custom=\"\">" +
				"a:after {content: url('https://leak.com')} " +
				"a::after {content: url('https://leak.com')} " +
				"a:before {content: url('https://leak.com')} " +
				"a::before {content: url('https://leak.com')} " +
				"big {" +
				"  list-style: url('https://leak.com'); " +
				"  list-style-image: url('https://leak.com'); " +
				"  background: url('https://leak.com'); " +
				"  background-image: url('https://leak.com'); " +
				"  border-image: url('https://leak.com'); " +
				"  -moz-border-image: url('https://leak.com'); " +
				"  -webkit-border-image: url('https://leak.com'); " +
				"  border-image-source: url('https://leak.com'); " +
				"  shape-outside: url('https://leak.com'); " +
				"  cursor: url('https://leak.com'), auto; " +
				"}" +
				"</style>",
			replacement: "https://leak-com.cdn.ampproject.org/i/s/leak.com",
		},
		{
			desc: "stylesheet with fonts",
			input: "<head><style amp-custom=\"\">" +
				"@font-face { " +
				" font-family: 'leak'; " +
				" src: url('https://leak.com') format('eot'), url('https://leak.com') " +
				"format('woff'), url('https://leak.com') format('truetype'); " +
				"} " +
				"@font-face { " +
				"  font-family: 'leak'; " +
				"  src: url('https://leak.com') format('eot'), url('https://leak.com') " +
				"format('woff'), url('https://leak.com') format('truetype'); " +
				"} " +
				"</style>",
			replacement: "https://leak-com.cdn.ampproject.org/r/s/leak.com",
		},
		{
			desc: "inline div",
			input: "<div style=\"background: url(&#39;&#39;) url(&#39;&#39;) " +
				"url(&#39;&#39;) " +
				"url(&#39;https://leak.com&#39;);\"></div>\n" +
				"<div style=\"behavior: url(&#39;https://leak.com&#39;);\"></div>\n" +
				"<div style=\"-ms-behavior: url(&#39;https://leak.com&#39;);\"></div>\n",
			replacement: "https://leak-com.cdn.ampproject.org/i/s/leak.com",
		},
		{
			desc: "another inline",
			input: "<b style=\"\n" +
				"  list-style: url(&#39;https://leak.com&#39;); \n" +
				"  list-style-image: url(&#39;https://leak.com&#39;); \n" +
				"  background: url(&#39;https://leak.com&#39;); \n" +
				"  background-image: url(&#39;https://leak.com&#39;); \n" +
				"  border-image: url(&#39;https://leak.com&#39;); \n" +
				"  -moz-border-image: url(&#39;https://leak.com&#39;); \n" +
				"  -webkit-border-image: url(&#39;https://leak.com&#39;); \n" +
				"  border-image-source: url(&#39;https://leak.com&#39;); \n" +
				"  shape-outside: url(&#39;https://leak.com&#39;); \n" +
				"  cursor: url(&#39;https://leak.com&#39;), auto; \n" +
				"\">MNO</b>",
			replacement: "https://leak-com.cdn.ampproject.org/i/s/leak.com",
		},
		{
			desc: "URLs reused as variables",
			input: "<style amp-custom=\"\">s {\n  --leak: url('https://leak.com');\n" +
				"}\ns{\n  background: var(--leak);\n}\n</style>",
			replacement: "https://leak-com.cdn.ampproject.org/i/s/leak.com",
		},
	}
	tcs := []urlRewriteTestCase{}
	for _, baseTc := range baseTcs {
		tc := urlRewriteTestCase{
			desc:     baseTc.desc,
			input:    baseTc.input,
			expected: strings.Replace(baseTc.input, "https://leak.com", baseTc.replacement, -1)}
		tcs = append(tcs, tc)
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_styleEdgeCases(t *testing.T) {
	tcs := []urlRewriteTestCase{
		{
			desc: "escaped code points",
			input: "<head><style amp-custom>" +
				"p#p1 {\n" +
				"  background-image: \\75 \\72 \\6C (https://leak.com);\n" +
				"}\n" +
				"p#p2 {\n" +
				"  background-image: \\000075\\000072\\00006C (https://leak.com);\n" +
				"}\n" +
				"</style>",
			expected: "<style amp-custom=\"\">" +
				"p#p1 {\n" +
				"  background-image: " +
				"url('https://leak-com.cdn.ampproject.org/i/s/leak.com');\n" +
				"}\n" +
				"p#p2 {\n" +
				"  background-image: " +
				"url('https://leak-com.cdn.ampproject.org/i/s/leak.com');\n" +
				"}\n" +
				"</style>",
		},
		{
			desc: "re-escape closing style in data and newline",
			input: "<head><style amp-custom>" +
				".foo { background-image:url('data:image/svg+xml;utf8," +
				"<svg xmlns=\"http://www.w3.org/2000/svg\">\\3c /style>" +
				"<path d=\"M108.044,42.407\\a c-22.58,65.505z\"/></svg>');}" +
				"</style></head>",
			expected: "<style amp-custom=\"\">" +
				".foo { background-image:url('data:image/svg+xml;utf8," +
				"<svg xmlns=\"http://www.w3.org/2000/svg\">\\3C /style>" +
				"<path d=\"M108.044,42.407\\A c-22.58,65.505z\"/>" +
				"</svg>');}</style>",
		},
	}
	runURLRewriteTestcases(t, tcs)
}

func TestURLRewrite_idempotent(t *testing.T) {
	html := `
<html amp>
<head>
<style amp-custom="">
	a:after {content: url('https://foo.com')}
	a::after {content: url('https://foo.com')}
	a:before {content: url('https://foo.com')}
	a::before {content: url('https://foo.com')}
	big {
	  list-style: url('https://foo.com');
	  list-style-image: url('https://foo.com');
	  background: url('https://foo.com');
	  background-image: url('https://foo.com');
	  border-image: url('https://foo.com');
	  -moz-border-image: url('https://foo.com');
	  -webkit-border-image: url('https://foo.com');
	  border-image-source: url('https://foo.com');
	  shape-outside: url('https://foo.com');
	  cursor: url('https://foo.com'), auto;
	}
</style>
</head>
<body>
<amp-img src=http://www.example.com/blah.jpg width=92 height=10 srcset="http://www.example.com/blah.jpg 50w"></amp-img>
<amp-anim src=http://www.example.com/blah.jpg width=92 height=10></amp-anim>
</body>
</html>
`

	first, err := transformAndPrint(html, "")
	if err != nil {
		t.Fatalf("transformAndPrint failed %q", err)
	}
	second, err := transformAndPrint(first, "")
	if err != nil {
		t.Fatalf("transformAndPrint failed %q", err)
	}
	if diff := cmp.Diff(first, second); diff != "" {
		t.Errorf("URLRewrite not idempotent (-want, +got):\n%s", diff)
	}
}

func transformAndPrint(input, base string) (string, error) {
	documentURL, _ := url.Parse("https://example.com/")
	baseURL, _ := url.Parse(base)
	inputDoc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return "", err
	}
	inputDOM, err := amphtml.NewDOM(inputDoc)
	if err != nil {
		return "", err
	}
	transformers.URLRewrite(&transformers.Context{
		DOM:         inputDOM,
		BaseURL:     baseURL,
		DocumentURL: documentURL,
	})
	var output strings.Builder
	if err := html.Render(&output, inputDoc); err != nil {
		return "", err
	}
	return output.String(), nil
}

func runURLRewriteTestcases(t *testing.T, tcs []urlRewriteTestCase) {
	for _, tc := range tcs {
		output, err := transformAndPrint(tc.input, tc.base)
		if err != nil {
			t.Errorf("%s: unexpected error %q", tc.desc, err)
		}
		if !strings.Contains(output, tc.expected) {
			t.Errorf("%s: URLRewrite=\n%q\ndoes not contain=\n%q", tc.desc, output, tc.expected)
		}
	}
}
