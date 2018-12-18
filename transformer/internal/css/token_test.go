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

	"google3/third_party/golang/godebug/pretty"
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
		desc, input string
		expected    Token
	}{
		{
			desc:     "whitespace",
			input:    " \n \r\t \r\n \f   ",
			expected: Token{Type: WhitespaceToken, Value: " \n \n\t \n \n   "},
		},
		{
			desc:     "string",
			input:    "\"foo'bar\"",
			expected: Token{Type: StringToken, Value: "foo'bar", Extra: "\""},
		},
		{
			desc:     "string with escape",
			input:    "'\\00066 \\0006f \\0006f'",
			expected: Token{Type: StringToken, Value: "foo", Extra: "'"},
		},
		{
			desc:     "hash",
			input:    "#foo",
			expected: Token{Type: HashToken, Value: "foo"},
		},
		{
			desc:     "hash with escape",
			input:    "#\\0066 \\006f \\006f",
			expected: Token{Type: HashToken, Value: "foo"},
		},
		{
			desc:     "hash delim",
			input:    "# ",
			expected: Token{Type: DelimToken, Value: "#"},
		},
		{
			desc:     "suffix match",
			input:    "$=",
			expected: Token{Type: SuffixMatchToken, Value: "$="},
		},
		{
			desc:     "suffix delim",
			input:    "$ ",
			expected: Token{Type: DelimToken, Value: "$"},
		},
		{
			desc:     "string single quote",
			input:    "'foo\"bar'",
			expected: Token{Type: StringToken, Value: "foo\"bar", Extra: "'"},
		},
		{
			desc:     "bad string",
			input:    "'foo\n'",
			expected: Token{Type: BadStringToken, Value: "foo", Extra: "'"},
		},
		{
			desc:     "imbalanced quote (still valid string)",
			input:    "'foo",
			expected: Token{Type: StringToken, Value: "foo", Extra: "'"},
		},
		{
			desc:     "left paren",
			input:    "(",
			expected: Token{Type: OpenParenToken, Value: "("},
		},
		{
			desc:     "right paren",
			input:    ")",
			expected: Token{Type: CloseParenToken, Value: ")"},
		},
		{
			desc:     "substring match",
			input:    "*=",
			expected: Token{Type: SubstringMatchToken, Value: "*="},
		},
		{
			desc:     "asterisk delim",
			input:    "* ",
			expected: Token{Type: DelimToken, Value: "*"},
		},
		{
			desc:     "positive number",
			input:    "+123.45",
			expected: Token{Type: NumberToken, Value: "+123.45"},
		},
		{
			desc:     "dimension",
			input:    "+12px",
			expected: Token{Type: DimensionToken, Value: "+12", Extra: "px"},
		},
		{
			desc:     "comma",
			input:    ",",
			expected: Token{Type: CommaToken, Value: ","},
		},
		{
			desc:     "negative number",
			input:    "-123.45",
			expected: Token{Type: NumberToken, Value: "-123.45"},
		},
		{
			desc:     "hyphen ident",
			input:    "-abc",
			expected: Token{Type: IdentToken, Value: "-abc"},
		},
		{
			desc:     "cdc",
			input:    "-->",
			expected: Token{Type: CDCToken, Value: "-->"},
		},
		{
			desc:     "hypen delim",
			input:    "- ",
			expected: Token{Type: DelimToken, Value: "-"},
		},
		{
			desc:     "decimal number",
			input:    ".42",
			expected: Token{Type: NumberToken, Value: ".42"},
		},
		{
			desc:     "decimal delim",
			input:    ". ",
			expected: Token{Type: DelimToken, Value: "."},
		},
		{
			desc:     "comment ignored",
			input:    "/* this is throw away */",
			expected: Token{Type: EOFToken, Value: ""},
		},
		{
			desc:     "slash delim",
			input:    "/ ",
			expected: Token{Type: DelimToken, Value: "/"},
		},
		{
			desc:     "colon",
			input:    ":",
			expected: Token{Type: ColonToken, Value: ":"},
		},
		{
			desc:     "semicolon",
			input:    ";",
			expected: Token{Type: SemicolonToken, Value: ";"},
		},
		{
			desc:     "cdo",
			input:    "<!--",
			expected: Token{Type: CDOToken, Value: "<!--"},
		},
		{
			desc:     "< delim",
			input:    "< ",
			expected: Token{Type: DelimToken, Value: "<"},
		},
		{
			desc:     "at keyword",
			input:    "@font-face",
			expected: Token{Type: AtKeywordToken, Value: "font-face"},
		},
		{
			desc:     "at delim",
			input:    "@ ",
			expected: Token{Type: DelimToken, Value: "@"},
		},
		{
			desc:     "left bracket",
			input:    "[",
			expected: Token{Type: OpenSquareToken, Value: "["},
		},
		{
			desc:     "right bracket",
			input:    "]",
			expected: Token{Type: CloseSquareToken, Value: "]"},
		},
		{
			desc:     "prefix match",
			input:    "^=",
			expected: Token{Type: PrefixMatchToken, Value: "^="},
		},
		{
			desc:     "carat delim",
			input:    "^",
			expected: Token{Type: DelimToken, Value: "^"},
		},
		{
			desc:     "left curly",
			input:    "{}",
			expected: Token{Type: OpenCurlyToken, Value: "{"},
		},
		{
			desc:     "right curly",
			input:    "}",
			expected: Token{Type: CloseCurlyToken, Value: "}"},
		},
		{
			desc:     "digits",
			input:    "123.45",
			expected: Token{Type: NumberToken, Value: "123.45"},
		},
		{
			desc:     "ident",
			input:    "foo",
			expected: Token{Type: IdentToken, Value: "foo"},
		},
		{
			desc:     "dash match",
			input:    "|=",
			expected: Token{Type: DashMatchToken, Value: "|="},
		},
		{
			desc:     "column token",
			input:    "||",
			expected: Token{Type: ColumnToken, Value: "||"},
		},
		{
			desc:     "column delim",
			input:    "| ",
			expected: Token{Type: DelimToken, Value: "|"},
		},
		{
			desc:     "eof",
			input:    "",
			expected: Token{Type: EOFToken},
		},
		{
			desc:     "url",
			input:    "url(foo.gif)",
			expected: Token{Type: URLToken, Value: "foo.gif"},
		},
		{
			desc:     "url single quote",
			input:    "url('foo.gif')",
			expected: Token{Type: URLToken, Value: "foo.gif"},
		},
		{
			desc:     "url double quote",
			input:    "url(\"foo.gif\")",
			expected: Token{Type: URLToken, Value: "foo.gif"},
		},
		{
			desc:     "url with whitespace",
			input:    "url(  \n\f  'foo.gif'  \t\t )",
			expected: Token{Type: URLToken, Value: "foo.gif"},
		},
		{
			desc:     "url as escaped code points",
			input:    "\\75 \\72 \\6c('foo.gif')",
			expected: Token{Type: URLToken, Value: "foo.gif"},
		},
		{
			desc:     "url as escaped code points #2",
			input:    "\\000075 \\000072 \\00006c('foo.gif')",
			expected: Token{Type: URLToken, Value: "foo.gif"},
		},
		{
			desc:     "url with escape in value",
			input:    "url(http://a.com/b/c=d\\000026e=f_g*h)",
			expected: Token{Type: URLToken, Value: "http://a.com/b/c=d&e=f_g*h"},
		},
		{
			desc:     "unicode range start only",
			input:    "u+26",
			expected: Token{Type: UnicodeRangeToken, Value: "U+0026"},
		},
		{
			desc:     "unicode range",
			input:    "u+0-7F",
			expected: Token{Type: UnicodeRangeToken, Value: "U+0000-007F"},
		},
		{
			desc:     "unicode range",
			input:    "u+0025-00FF",
			expected: Token{Type: UnicodeRangeToken, Value: "U+0025-00FF"},
		},
				{
			desc:     "unicode range wild",
			input:    "u+4??",
			expected: Token{Type: UnicodeRangeToken, Value: "U+0400-04FF"},
		},

	}
	for _, tc := range tcs {
		z := NewTokenizer(tc.input)
		actual := z.Next()
		if diff := pretty.Compare(tc.expected, actual); diff != "" {
			t.Errorf("%s returned diff (-want, +got):\n%s", tc.desc, diff)
		}
	}
}

func TestTokenization(t *testing.T) {
	tcs := []struct {
		desc, input string
		expected    []Token
	}{
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
	}
	for _, tc := range tcs {
		z := NewTokenizer(tc.input)
		actual := []Token{}
		for {
			token := z.Next()
			actual = append(actual, token)
			if token.Type == EOFToken || token.Type == ErrorToken {
				break
			}
		}
		if diff := pretty.Compare(tc.expected, actual); diff != "" {
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
	tokens := []Token{}
	var sb strings.Builder
	for {
		token := z.Next()
		tokens = append(tokens, token)
		sb.WriteString(token.String())
		if token.Type == EOFToken || token.Type == ErrorToken {
			break
		}
	}
	z = NewTokenizer(sb.String())
	secondRound := []Token{}
	for {
		token := z.Next()
		secondRound = append(secondRound, token)
		if token.Type == EOFToken || token.Type == ErrorToken {
			break
		}
	}
	if diff := pretty.Compare(tokens, secondRound); diff != "" {
		t.Errorf("returned diff (-want, +got):\n%s", diff)
	}
}
