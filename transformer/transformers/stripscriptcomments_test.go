package transformers_test

import (
	"strings"
	"testing"
)

import (
	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
	tt "github.com/ampproject/amppackager/transformer/internal/testing"
)

func TestStripScriptComments(t *testing.T) {
	tcs := []tt.TestCase{
		{
			Desc: "No comments in script",
		        Input: "<html><head><script type=\"application/json\">alert('hi');</script></head><body></body></html>",
		        Expected: "<html><head><script type=\"application/json\">alert('hi');</script></head><body></body></html>",
		},
		{
			Desc: "Basic html comment in script",
		        Input: "<html><head><script type=\"application/json\"><!--alert('hi');--></script></head><body></body></html>",
		        Expected: "<html><head><script type=\"application/json\"></script></head><body></body></html>",
		},
		{
			Desc: "Basic html comment with condition",
		        Input: "<html><head><script type=\"application/json\"><!--[if IE]>alert('hi');<![endif]--></script></head><body></body></html>",
		        Expected: "<html><head><script type=\"application/json\"></script></head><body></body></html>",
		},
		{
			Desc: "Basic html comment with IE 10 condition",
		        Input: "<html><head><script type=\"application/json\"><!--[if gte IE 10]>alert('hi');<!--<![endif]--></script></head><body></body></html>",
		        Expected: "<html><head><script type=\"application/json\"></script></head><body></body></html>",
		},
		{
			Desc: "Basic html comment with different syntax",
		        Input: "<html><head><script type=\"application/json\"><!--[if gte IE 10]><!-->alert('hi');<!--<![endif]--></script></head><body></body></html>",
		        Expected: "<html><head><script type=\"application/json\"></script></head><body></body></html>",
		},
		{
			Desc: "Script with comment inside string variable",
		        Input: "<html><head><script type=\"application/json\">var k = \"<!--[if !IE]> -->alert('hi');<!-- <![endif]-->\";</script></head><body></body></html>",
		        Expected: "<html><head><script type=\"application/json\">var k = \"<!--[if !IE]> -->alert('hi');<!-- <![endif]-->\";</script></head><body></body></html>",
		},
		{
			Desc: "Script with comment inside string variable escaped slashes",
		        Input: "<html><head><script type=\"application/json\">var k = \"<!--[if !IE]> -->alert(\\\"hi\\\");<!-- <![endif]-->\";</script></head><body></body></html>",
		        Expected: "<html><head><script type=\"application/json\">var k = \"<!--[if !IE]> -->alert(\\\"hi\\\");<!-- <![endif]-->\";</script></head><body></body></html>",
		},
		{
			Desc: "Script with comment containing quotes",
		        Input: "<html><head><script type=\"application/json\">var k = \"<!--[if !IE]> -->alert(\"hi\");<!-- <![endif]-->\";</script></head><body></body></html>",
		        Expected: "<html><head><script type=\"application/json\">var k = \"<!--[if !IE]> -->alert(\"hi\");<!-- <![endif]-->\";</script></head><body></body></html>",
		},
		{
			Desc: "Script with comment and some valid content",
		        Input: "<html><head><script type=\"application/json\"><!--[if !IE]>alert('hi');<![endif]-->var a = 1; var k = \"hi\";</script></head><body></body></html>",
		        Expected: "<html><head><script type=\"application/json\">var a = 1; var k = \"hi\";</script></head><body></body></html>",
		},
		{
			Desc: "Script with empty comment",
		        Input: "<html><head><script type=\"application/json\"><!-- --></script></head><body></body></html>",
		        Expected: "<html><head><script type=\"application/json\"></script></head><body></body></html>",
		},
		{
			Desc: "Script with empty comment and variables",
		        Input: "<html><head><script type=\"application/json\"><!-- -->var a = 1;</script></head><body></body></html>",
		        Expected: "<html><head><script type=\"application/json\">var a = 1;</script></head><body></body></html>",
		},
		{
			Desc: "Script with too many comments",
		        Input: "<html><head><script type=\"application/json\"><!-- comment -->var a = 1;<!--[if IE]>It's an IE<![endif]-->var b = \"hello\";var c = false;</script></head><body></body></html>",
		        Expected: "<html><head><script type=\"application/json\">var a = 1;var b = \"hello\";var c = false;</script></head><body></body></html>",
		},
		{
			Desc: "Too many comments V2",
			Input: "<html><head><script type=\"application/json\"><!-- comment -->var a = 1;<!--[if IE]>It's an IE<![endif]-->var b = \"hello\";var c = false;<!-- hello -->var d = \"foo\";<!--[if !IE]>One more conditional comment<![endif]-->var f = \"<!-- comment variable -->\";</script>",
			Expected: "<html><head><script type=application/json>var a = 1;var b = \"hello\";var c = false;var d = \"foo\";var f = \"<!-- comment variable -->\";</script></head><body></body></html>",
		},
		{
			Desc: "Simple JSON",
			Input: "<html><head><script type=\"application/json\"><!-- comment -->{'a': 'b', 'c': true, 'd': 123}</script>",
			Expected: "<html><head><script type=application/json>{'a': 'b', 'c': true, 'd': 123}</script></head><body></body></html>",
		},
		{
			Desc: "JSON with comment as string variable",
			Input: "<html><head><script type=\"application/ld+json\"><!-- comment -->{'a': 'b', 'c': true, 'd': '<!-- should preserve -->'}</script>",
			Expected: "<html><head><script type=application/ld+json>{'a': 'b', 'c': true, 'd': '<!-- should preserve -->'}</script></head><body></body></html>",
		},
		{
			Desc: "JSON with double quoted values",
			Input: "<script type=\"application/ld+json\"><!-- comment -->{\"a\": \"b\", \"c\": true, \"d\": \"<!-- should preserve -->\"}</script>",
			Expected: "<html><head><script type=application/ld+json>{\"a\": \"b\", \"c\": true, \"d\": \"<!-- should preserve -->\"}</script></head><body></body></html>",
		},
		{
			Desc: "JSON with too many escape chars",
			Input: "<html><head><script type=\"application/ld+json\"><!-- comment -->{\"a\": \"\\\\\"b\", \"c\": true, \"d\": \"<!-- should preserve -->\"}</script>",
			Expected: "<html><head><script type=application/ld+json>{\"a\": \"\\\\\"b\", \"c\": true, \"d\": \"<!-- should preserve -->\"}</script></head><body></body></html>",
		},
		{
			Desc: "JSON with multiple comments",
			Input: "<html><head><script type=\"application/json\"><!-- comment -->{\"a\": \"b\", \"c\": true, \"d\": \"<!-- should preserve -->\", <!-- more comments --> e: 123}</script>",
			Expected: "<html><head><script type=application/json>{\"a\": \"b\", \"c\": true, \"d\": \"<!-- should preserve -->\",  e: 123}</script></head><body></body></html>",
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
		transformers.StripScriptComments(&transformers.Context{DOM: inputDOM})
		var input strings.Builder
		if err = html.Render(&input, inputDoc); err != nil {
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
			t.Errorf("%s: Transform\n%q\nwant=\n%q", tc.Desc, &input, &expected)
		}
	}
}
