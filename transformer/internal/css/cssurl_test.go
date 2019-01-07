package css

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseURLs(t *testing.T) {
	tcs := []struct {
		desc, input string
		expected    Segments
	}{
		{
			desc:  "image url single-quote",
			input: "body{background-image: url('http://foo.com/blah.png');}",
			expected: Segments{
				Segment{ByteType, "body{background-image: "},
				Segment{ImageURLType, "http://foo.com/blah.png"},
				Segment{ByteType, ";}"},
			},
		},
		{
			desc:  "image url double-quote",
			input: "body{background-image: url(\"http://foo.com/blah.png\");}",
			expected: Segments{
				Segment{ByteType, "body{background-image: "},
				Segment{ImageURLType, "http://foo.com/blah.png"},
				Segment{ByteType, ";}"},
			},
		},
		{
			desc:  "image url function",
			input: "body{background-image: url(http://foo.com/blah.png);}",
			expected: Segments{
				Segment{ByteType, "body{background-image: "},
				Segment{ImageURLType, "http://foo.com/blah.png"},
				Segment{ByteType, ";}"},
			},
		},
		{
			desc:  "image with unicode",
			input: "body{background-image: url('http://a.com/b/c=d\u0026e=f_g*h');}",
			expected: Segments{
				Segment{ByteType, "body{background-image: "},
				Segment{ImageURLType, "http://a.com/b/c=d&e=f_g*h"},
				Segment{ByteType, ";}"},
			},
		},
		{
			desc:  "image with character escape",
			input: "body{background-image: url('http://a.com/b/c=d\\000026e=f_g*h');}",
			expected: Segments{
				Segment{ByteType, "body{background-image: "},
				Segment{ImageURLType, "http://a.com/b/c=d&e=f_g*h"},
				Segment{ByteType, ";}"},
			},
		},
		{
			desc:  "ident with unicode",
			input: "body{background-image: \\75 \\72 \\6c('http://a.com/b/c=d\u0026e=f_g*h');}",
			expected: Segments{
				Segment{ByteType, "body{background-image: "},
				Segment{ImageURLType, "http://a.com/b/c=d&e=f_g*h"},
				Segment{ByteType, ";}"},
			},
		},
		{
			desc:  "font",
			input: "@font-face {font-family: 'Foo'; src: url('http://foo.com/bar.ttf');}",
			expected: Segments{
				Segment{ByteType, "@font-face {font-family: 'Foo'; src: "},
				Segment{FontURLType, "http://foo.com/bar.ttf"},
				Segment{ByteType, ";}"},
			},
		},
		{
			desc: "font with nested parens",
			input: "@font-face {foo: {} font-family: 'Foo'; src: url('http://foo.com/bar.ttf');}\n" +
				".a { background-image: url(http://foo.com/baz.png); }",
			expected: Segments{
				Segment{ByteType, "@font-face {foo: {} font-family: 'Foo'; src: "},
				Segment{FontURLType, "http://foo.com/bar.ttf"},
				Segment{ByteType, ";}\n.a { background-image: "},
				Segment{ImageURLType, "http://foo.com/baz.png"},
				Segment{ByteType, "; }"},
			},
		},
		{
			// NOTE this will fail AMP Cache validation.
			desc: "font with imbalanced open paren",
			input: "@font-face {foo: {; font-family: 'Foo'; src: url('http://foo.com/bar.ttf');}\n" +
				".a { background-image: url(http://foo.com/baz.png); }",
			expected: Segments{
				Segment{ByteType, "@font-face {foo: {; font-family: 'Foo'; src: "},
				Segment{FontURLType, "http://foo.com/bar.ttf"},
				Segment{ByteType, ";}\n.a { background-image: "},
				// NOTE the incorrect type assignation is WAI. The parser thinks the at-rule is still open.
				Segment{FontURLType, "http://foo.com/baz.png"},
				Segment{ByteType, "; }"},
			},
		},
		{
			// NOTE this will fail AMP Cache validation.
			desc: "font with imbalanced close paren",
			input: "@font-face {foo: }; font-family: 'Foo'; src: url('http://foo.com/bar.ttf');}\n" +
				".a { background-image: url(http://foo.com/baz.png); }",
			expected: Segments{
				Segment{ByteType, "@font-face {foo: }; font-family: 'Foo'; src: "},
				// NOTE the incorrect type assignation is WAI. The parser thinks the at-rule has been closed.
				Segment{ImageURLType, "http://foo.com/bar.ttf"},
				Segment{ByteType, ";}\n.a { background-image: "},
				Segment{ImageURLType, "http://foo.com/baz.png"},
				Segment{ByteType, "; }"},
			},
		},
		{
			desc: "longer example",
			input: ".a { color:red; background-image:url(4.png) }\n" +
				".b { color:black; background-image:url('http://a.com/b.png') }\n" +
				"@font-face {font-family: 'Medium';src: url('http://a.com/1.woff')\n" +
				"format('woff'),url('http://b.com/1.ttf') format('truetype'),\n" +
				"src:url('') format('embedded-opentype');}\n" +
				".c { color:blue; background-image:url(5.png) }\n",
			expected: Segments{
				Segment{ByteType, ".a { color:red; background-image:"},
				Segment{ImageURLType, "4.png"},
				Segment{ByteType, " }\n.b { color:black; background-image:"},
				Segment{ImageURLType, "http://a.com/b.png"},
				Segment{ByteType, " }\n@font-face {font-family: 'Medium';src: "},
				Segment{FontURLType, "http://a.com/1.woff"},
				Segment{ByteType, "\nformat('woff'),"},
				Segment{FontURLType, "http://b.com/1.ttf"},
				Segment{ByteType, " format('truetype'),\nsrc:"},
				Segment{FontURLType, ""},
				Segment{ByteType, " format('embedded-opentype');}\n.c { color:blue; background-image:"},
				Segment{ImageURLType, "5.png"},
				Segment{ByteType, " }\n"},
			},
		},
		{
			desc: "newlines",
			input: ".a \r\n{ color:red; background-image:url(4.png) }\r\n" +
				".b { color:black; \r\nbackground-image:url('http://a.com/b.png') }",
			expected: Segments{
				Segment{ByteType, ".a \n{ color:red; background-image:"},
				Segment{ImageURLType, "4.png"},
				Segment{ByteType, " }\n.b { color:black; \nbackground-image:"},
				Segment{ImageURLType, "http://a.com/b.png"},
				Segment{ByteType, " }"},
			},
		},
		{
			desc: "html characters",
			input: ".x{background: url(&#39;&#39;) url(&#39;&#39;) " +
				"url(&#39;&#39;) " +
				"url(&#39;https://leak.com&#39;)};",
			expected: Segments{
				Segment{ByteType, ".x{background: "},
				Segment{ImageURLType, "&#39;&#39;"},
				Segment{ByteType, " "},
				Segment{ImageURLType, "&#39;&#39;"},
				Segment{ByteType, " "},
				Segment{ImageURLType, "&#39;&#39;"},
				Segment{ByteType, " "},
				Segment{ImageURLType, "&#39;https://leak.com&#39;"},
				Segment{ByteType, "};"},
			},
		},
		{
			desc:  "function pass through",
			input: "amp-lightbox {  background-color: rgba(0, 0, 0, 0.9); }",
			expected: Segments{
				Segment{ByteType, "amp-lightbox {  background-color: rgba(0, 0, 0, 0.9); }"},
			},
		},
	}

	for _, tc := range tcs {
		if actual, err := ParseURLs(tc.input); err == nil {
			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("%s: Segment() returned diff (-want +got):\n%s", tc.desc, diff)
			}
		} else {
			t.Errorf("%s: Segment() generated unexpected error %v", tc.desc, err)
		}
	}
}
