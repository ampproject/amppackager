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

package css

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewTokenizer(t *testing.T) {
	tcs := []struct{ desc, input, expected string }{
		{
			desc:     "no preprocessing",
			input:    "hi there",
			expected: "hi there",
		},
		{
			desc:     "CR",
			input:    "hi\rthere",
			expected: "hi\nthere",
		},
		{
			desc:     "FF",
			input:    "hi\fthere",
			expected: "hi\nthere",
		},
		{
			desc:     "CR LF",
			input:    "hi\r\nthere",
			expected: "hi\nthere",
		},
		{
			desc:     "CR CR LF",
			input:    "hi\r\r\nthere",
			expected: "hi\n\nthere",
		},
		{
			desc:     "crazy combo",
			input:    "hi\n\r\n\f\r\nthere",
			expected: "hi\n\n\n\nthere",
		},
	}
	for _, tc := range tcs {
		z := NewTokenizer(tc.input)
		if (string)(z.input) != tc.expected {
			t.Errorf("%s: NewTokenizer(%q) = %q, want = %q", tc.desc, tc.input, z.input, tc.expected)
		}
	}
}

func TestSingleToken(t *testing.T) {
	tcs := []struct {
		desc, input    string
		expected       Token
		expectedString string
	}{
		{
			desc:           "whitespace",
			input:          " \n \r\t \r\n \f   ",
			expected:       Token{Type: WhitespaceToken, Value: " \n \n\t \n \n   "},
			expectedString: " \n \n\t \n \n   ",
		},
		{
			desc:           "string",
			input:          "\"foo'bar\"",
			expected:       Token{Type: StringToken, Value: "foo'bar", Extra: "\""},
			expectedString: "\"foo'bar\"",
		},
		{
			desc:           "double quoted string with embedded double quote",
			input:          `"this is a \"string\""`,
			expected:       Token{Type: StringToken, Value: "this is a \"string\"", Extra: "\""},
			expectedString: `"this is a \"string\""`,
		},
		{
			desc:           "string with escape",
			input:          "'\\00066 \\0006f \\0006f'",
			expected:       Token{Type: StringToken, Value: "foo", Extra: "'"},
			expectedString: "'\\00066 \\0006f \\0006f'",
		},
		{
			desc:           "raw string with newline escape",
			input:          "'" + `English \a Version` + "'",
			expected:       Token{Type: StringToken, Value: "English \nVersion", Extra: "'"},
			expectedString: "'English \\a Version'",
		},
		{
			desc:           "full newline escape",
			input:          "'" + `hi\00000abye` + "'",
			expected:       Token{Type: StringToken, Value: "hi\nbye", Extra: "'"},
			expectedString: "'hi\\00000abye'",
		},
		{
			desc:           "hash",
			input:          "#foo",
			expected:       Token{Type: HashToken, Value: "foo"},
			expectedString: "#foo",
		},
		{
			desc:           "hash with escape",
			input:          "#\\0066 \\006f \\006f",
			expected:       Token{Type: HashToken, Value: "foo"},
			expectedString: "#\\0066 \\006f \\006f",
		},
		{
			desc:           "hash delim",
			input:          "# ",
			expected:       Token{Type: DelimToken, Value: "#"},
			expectedString: "#",
		},
		{
			desc:           "suffix match",
			input:          "$=",
			expected:       Token{Type: SuffixMatchToken, Value: "$="},
			expectedString: "$=",
		},
		{
			desc:           "suffix delim",
			input:          "$ ",
			expected:       Token{Type: DelimToken, Value: "$"},
			expectedString: "$",
		},
		{
			desc:           "string single quote",
			input:          "'foo\"bar'",
			expected:       Token{Type: StringToken, Value: "foo\"bar", Extra: "'"},
			expectedString: "'foo\"bar'",
		},
		{
			desc:           "bad string",
			input:          "'foo\n'",
			expected:       Token{Type: BadStringToken, Value: "foo", Extra: "'"},
			expectedString: "'foo",
		},
		{
			desc:           "imbalanced quote (still valid string)",
			input:          "'foo",
			expected:       Token{Type: StringToken, Value: "foo", Extra: "'"},
			expectedString: "'foo",
		},
		{
			desc:           "left paren",
			input:          "(",
			expected:       Token{Type: OpenParenToken, Value: "("},
			expectedString: "(",
		},
		{
			desc:           "right paren",
			input:          ")",
			expected:       Token{Type: CloseParenToken, Value: ")"},
			expectedString: ")",
		},
		{
			desc:           "substring match",
			input:          "*=",
			expected:       Token{Type: SubstringMatchToken, Value: "*="},
			expectedString: "*=",
		},
		{
			desc:           "asterisk delim",
			input:          "* ",
			expected:       Token{Type: DelimToken, Value: "*"},
			expectedString: "*",
		},
		{
			desc:           "positive number",
			input:          "+123.45",
			expected:       Token{Type: NumberToken, Value: "+123.45"},
			expectedString: "+123.45",
		},
		{
			desc:           "dimension",
			input:          "+12px",
			expected:       Token{Type: DimensionToken, Value: "+12", Extra: "px"},
			expectedString: "+12px",
		},
		{
			desc:           "comma",
			input:          ",",
			expected:       Token{Type: CommaToken, Value: ","},
			expectedString: ",",
		},
		{
			desc:           "negative number",
			input:          "-123.45",
			expected:       Token{Type: NumberToken, Value: "-123.45"},
			expectedString: "-123.45",
		},
		{
			desc:           "hyphen ident",
			input:          "-abc",
			expected:       Token{Type: IdentToken, Value: "-abc"},
			expectedString: "-abc",
		},
		{
			desc:           "cdc",
			input:          "-->",
			expected:       Token{Type: CDCToken, Value: "-->"},
			expectedString: "-->",
		},
		{
			desc:           "hypen delim",
			input:          "- ",
			expected:       Token{Type: DelimToken, Value: "-"},
			expectedString: "-",
		},
		{
			desc:           "decimal number",
			input:          ".42",
			expected:       Token{Type: NumberToken, Value: ".42"},
			expectedString: ".42",
		},
		{
			desc:           "decimal delim",
			input:          ". ",
			expected:       Token{Type: DelimToken, Value: "."},
			expectedString: ".",
		},
		{
			desc:           "comment ignored as token value, but preserved in StringValue()",
			input:          "/* this is throw away */",
			expected:       Token{Type: EOFToken, Value: ""},
			expectedString: "/* this is throw away */",
		},
		{
			desc:           "slash delim",
			input:          "/ ",
			expected:       Token{Type: DelimToken, Value: "/"},
			expectedString: "/",
		},
		{
			desc:           "colon",
			input:          ":",
			expected:       Token{Type: ColonToken, Value: ":"},
			expectedString: ":",
		},
		{
			desc:           "semicolon",
			input:          ";",
			expected:       Token{Type: SemicolonToken, Value: ";"},
			expectedString: ";",
		},
		{
			desc:           "cdo",
			input:          "<!--",
			expected:       Token{Type: CDOToken, Value: "<!--"},
			expectedString: "<!--",
		},
		{
			desc:           "< delim",
			input:          "< ",
			expected:       Token{Type: DelimToken, Value: "<"},
			expectedString: "<",
		},
		{
			desc:           "at keyword",
			input:          "@font-face",
			expected:       Token{Type: AtKeywordToken, Value: "font-face"},
			expectedString: "@font-face",
		},
		{
			desc:           "at delim",
			input:          "@ ",
			expected:       Token{Type: DelimToken, Value: "@"},
			expectedString: "@",
		},
		{
			desc:           "left bracket",
			input:          "[",
			expected:       Token{Type: OpenSquareToken, Value: "["},
			expectedString: "[",
		},
		{
			desc:           "right bracket",
			input:          "]",
			expected:       Token{Type: CloseSquareToken, Value: "]"},
			expectedString: "]",
		},
		{
			desc:           "prefix match",
			input:          "^=",
			expected:       Token{Type: PrefixMatchToken, Value: "^="},
			expectedString: "^=",
		},
		{
			desc:           "carat delim",
			input:          "^",
			expected:       Token{Type: DelimToken, Value: "^"},
			expectedString: "^",
		},
		{
			desc:           "left curly",
			input:          "{}",
			expected:       Token{Type: OpenCurlyToken, Value: "{"},
			expectedString: "{",
		},
		{
			desc:           "right curly",
			input:          "}",
			expected:       Token{Type: CloseCurlyToken, Value: "}"},
			expectedString: "}",
		},
		{
			desc:           "digits",
			input:          "123.45",
			expected:       Token{Type: NumberToken, Value: "123.45"},
			expectedString: "123.45",
		},
		{
			desc:           "ident",
			input:          "foo",
			expected:       Token{Type: IdentToken, Value: "foo"},
			expectedString: "foo",
		},
		{
			desc:           "hash",
			input:          "#p1",
			expected:       Token{Type: HashToken, Value: "p1"},
			expectedString: "#p1",
		},
		{
			desc:           "dash match",
			input:          "|=",
			expected:       Token{Type: DashMatchToken, Value: "|="},
			expectedString: "|=",
		},
		{
			desc:           "column token",
			input:          "||",
			expected:       Token{Type: ColumnToken, Value: "||"},
			expectedString: "||",
		},
		{
			desc:           "column delim",
			input:          "| ",
			expected:       Token{Type: DelimToken, Value: "|"},
			expectedString: "|",
		},
		{
			desc:           "eof",
			input:          "",
			expected:       Token{Type: EOFToken},
			expectedString: "",
		},
		{
			desc:           "url function",
			input:          "url(foo.gif)",
			expected:       Token{Type: URLToken, Value: "foo.gif"},
			expectedString: "url(foo.gif)",
		},
		{
			desc:           "data url with escape newline",
			input:          "url('data:image/svg+xml;\\a <svg></svg>')",
			expected:       Token{Type: URLToken, Value: "data:image/svg+xml;\n<svg></svg>"},
			expectedString: "url('data:image/svg+xml;\\a <svg></svg>')",
		},
		{
			desc:           "url single quote",
			input:          "url('foo.gif')",
			expected:       Token{Type: URLToken, Value: "foo.gif"},
			expectedString: "url('foo.gif')",
		},
		{
			desc:           "url double quote",
			input:          "url(\"foo.gif\")",
			expected:       Token{Type: URLToken, Value: "foo.gif"},
			expectedString: "url(\"foo.gif\")",
		},
		{
			desc:           "url with whitespace, drops whitespace in token value, StringValue is preprocessed",
			input:          "url(  \n\f  'foo.gif'  \t\t )",
			expected:       Token{Type: URLToken, Value: "foo.gif"},
			expectedString: "url(  \n\n  'foo.gif'  \t\t )",
		},
		{
			desc:           "url with html chars",
			input:          "url(  &#39;&#39;)",
			expected:       Token{Type: URLToken, Value: "&#39;&#39;"},
			expectedString: "url(  &#39;&#39;)",
		},
		{
			desc:           "url as escaped code points",
			input:          "\\75 \\72 \\6c('foo.gif')",
			expected:       Token{Type: URLToken, Value: "foo.gif"},
			expectedString: "\\75 \\72 \\6c('foo.gif')",
		},
		{
			desc:           "url as escaped code points #2",
			input:          "\\000075 \\000072 \\00006c('foo.gif')",
			expected:       Token{Type: URLToken, Value: "foo.gif"},
			expectedString: "\\000075 \\000072 \\00006c('foo.gif')",
		},
		{
			desc:           "url with embedded single quote",
			input:          `url("fo'o.gif")`,
			expected:       Token{Type: URLToken, Value: "fo'o.gif"},
			expectedString: `url("fo'o.gif")`,
		},
		{
			desc:           "raw url with embedded single quote",
			input:          `url('fo\'o.gif')`,
			expected:       Token{Type: URLToken, Value: "fo'o.gif"},
			expectedString: "url('fo\\'o.gif')",
		},
		{
			desc:           "url with escape in value",
			input:          "url(http://a.com/b/c=d\\000026e=f_g*h)",
			expected:       Token{Type: URLToken, Value: "http://a.com/b/c=d&e=f_g*h"},
			expectedString: "url(http://a.com/b/c=d\\000026e=f_g*h)",
		},
		{
			desc:           "unicode range start only",
			input:          "u+26",
			expected:       Token{Type: UnicodeRangeToken, Value: "U+0026"},
			expectedString: "u+26",
		},
		{
			desc:           "unicode range",
			input:          "u+0-7F",
			expected:       Token{Type: UnicodeRangeToken, Value: "U+0000-007F"},
			expectedString: "u+0-7F",
		},
		{
			desc:           "unicode range",
			input:          "u+0025-00FF",
			expected:       Token{Type: UnicodeRangeToken, Value: "U+0025-00FF"},
			expectedString: "u+0025-00FF",
		},
		{
			desc:           "unicode range wild",
			input:          "u+4??",
			expected:       Token{Type: UnicodeRangeToken, Value: "U+0400-04FF"},
			expectedString: "u+4??",
		},
	}
	for _, tc := range tcs {
		z := NewTokenizer(tc.input)
		actual := z.Next()
		if diff := cmp.Diff(tc.expected, actual, cmpopts.IgnoreUnexported(Token{})); diff != "" {
			t.Errorf("%s returned diff (-want, +got):\n%s", tc.desc, diff)
		}
		if tc.expectedString != actual.String() {
			t.Errorf("%s: String() got=%q want=%q", tc.desc, actual.String(), tc.expectedString)
		}
	}
}

func TestTokenization(t *testing.T) {
	tcs := []struct {
		desc, input string
		expected    []Token
	}{
		{
			desc:  "function",
			input: "amp-lightbox {  background-color: rgba(0, 0, 0, 0.9); }",
			expected: []Token{
				Token{Type: IdentToken, Value: "amp-lightbox"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: OpenCurlyToken, Value: "{"},
				Token{Type: WhitespaceToken, Value: "  "},
				Token{Type: IdentToken, Value: "background-color"},
				Token{Type: ColonToken, Value: ":"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: FunctionToken, Value: "rgba"},
				Token{Type: NumberToken, Value: "0"},
				Token{Type: CommaToken, Value: ","},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: NumberToken, Value: "0"},
				Token{Type: CommaToken, Value: ","},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: NumberToken, Value: "0"},
				Token{Type: CommaToken, Value: ","},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: NumberToken, Value: "0.9"},
				Token{Type: CloseParenToken, Value: ")"},
				Token{Type: SemicolonToken, Value: ";"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: CloseCurlyToken, Value: "}"},
				Token{Type: EOFToken},
			},
		},
		{
			desc: "longer case",
			input: ".a { color:red; background-image:url(4.png) }" +
				".b { color:black; background-image:url('http://a.com/b.png') } " +
				"@font-face {font-family: 'Medium';src: url('http://a.com/1.woff') " +
				"format('woff'),url('http://b.com/1.ttf') format('truetype')," +
				"src:url('') format('embedded-opentype');}",
			expected: []Token{
				Token{Type: DelimToken, Value: "."},
				Token{Type: IdentToken, Value: "a"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: OpenCurlyToken, Value: "{"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: IdentToken, Value: "color"},
				Token{Type: ColonToken, Value: ":"},
				Token{Type: IdentToken, Value: "red"},
				Token{Type: SemicolonToken, Value: ";"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: IdentToken, Value: "background-image"},
				Token{Type: ColonToken, Value: ":"},
				Token{Type: URLToken, Value: "4.png"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: CloseCurlyToken, Value: "}"},
				Token{Type: DelimToken, Value: "."},
				Token{Type: IdentToken, Value: "b"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: OpenCurlyToken, Value: "{"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: IdentToken, Value: "color"},
				Token{Type: ColonToken, Value: ":"},
				Token{Type: IdentToken, Value: "black"},
				Token{Type: SemicolonToken, Value: ";"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: IdentToken, Value: "background-image"},
				Token{Type: ColonToken, Value: ":"},
				Token{Type: URLToken, Value: "http://a.com/b.png"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: CloseCurlyToken, Value: "}"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: AtKeywordToken, Value: "font-face"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: OpenCurlyToken, Value: "{"},
				Token{Type: IdentToken, Value: "font-family"},
				Token{Type: ColonToken, Value: ":"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: StringToken, Value: "Medium", Extra: "'"},
				Token{Type: SemicolonToken, Value: ";"},
				Token{Type: IdentToken, Value: "src"},
				Token{Type: ColonToken, Value: ":"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: URLToken, Value: "http://a.com/1.woff"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: FunctionToken, Value: "format"},
				Token{Type: StringToken, Value: "woff", Extra: "'"},
				Token{Type: CloseParenToken, Value: ")"},
				Token{Type: CommaToken, Value: ","},
				Token{Type: URLToken, Value: "http://b.com/1.ttf"},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: FunctionToken, Value: "format"},
				Token{Type: StringToken, Value: "truetype", Extra: "'"},
				Token{Type: CloseParenToken, Value: ")"},
				Token{Type: CommaToken, Value: ","},
				Token{Type: IdentToken, Value: "src"},
				Token{Type: ColonToken, Value: ":"},
				Token{Type: URLToken},
				Token{Type: WhitespaceToken, Value: " "},
				Token{Type: FunctionToken, Value: "format"},
				Token{Type: StringToken, Value: "embedded-opentype", Extra: "'"},
				Token{Type: CloseParenToken, Value: ")"},
				Token{Type: SemicolonToken, Value: ";"},
				Token{Type: CloseCurlyToken, Value: "}"},
				Token{Type: EOFToken},
			},
		},
		{
			desc:  "raw string",
			input: `.bar{content:"English \a Version"}`,
			expected: []Token{
				Token{Type: DelimToken, Value: "."},
				Token{Type: IdentToken, Value: "bar"},
				Token{Type: OpenCurlyToken, Value: "{"},
				Token{Type: IdentToken, Value: "content"},
				Token{Type: ColonToken, Value: ":"},
				Token{Type: StringToken, Value: "English \nVersion", Extra: "\""},
				Token{Type: CloseCurlyToken, Value: "}"},
				Token{Type: EOFToken},
			},
		},
	}
	for _, tc := range tcs {
		z := NewTokenizer(tc.input)
		actual := z.All()
		if diff := cmp.Diff(tc.expected, actual, cmpopts.IgnoreUnexported(Token{})); diff != "" {
			t.Errorf("%s returned diff (-want, +got):\n%s", tc.desc, diff)
		}
	}
}

func TestSerialization(t *testing.T) {
	css := ".a { color:red; background-image:url(4.png) }" +
		".b { color:black; background-image:url('http://a.com/b.png') } " +
		"@font-face {font-family: 'Medium';src: url('http://a.com/1.woff') " +
		"format('woff'),url('http://b.com/1.ttf') format('truetype')," +
		"src:url('') format('embedded-opentype');}"
	z := NewTokenizer(css)
	var sb strings.Builder
	first := z.All()
	for _, token := range first {
		sb.WriteString(token.String())
	}
	z = NewTokenizer(sb.String())
	second := z.All()
	if diff := cmp.Diff(first, second, cmpopts.IgnoreUnexported(Token{})); diff != "" {
		t.Errorf("returned diff (-want, +got):\n%s", diff)
	}
}
