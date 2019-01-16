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
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// A TokenType is a type of Token
type TokenType int

const (
	// ErrorToken means an error occurred during tokenization.
	ErrorToken = iota
	// EOFToken marks the end of the input.
	EOFToken

	// CSS3 tokens as per section 4.

	// IdentToken is an identity token
	IdentToken
	// FunctionToken is a function token
	FunctionToken
	// AtKeywordToken is an @ keyword
	AtKeywordToken
	// HashToken is a hash
	HashToken
	// StringToken is a quoted string
	StringToken
	// BadStringToken is a badly parsed quoted string
	BadStringToken
	// URLToken is a URL
	URLToken
	// BadURLToken denotes a badly formed URL
	BadURLToken
	// DelimToken is a single code point
	DelimToken
	// NumberToken is a number
	NumberToken
	// PercentageToken is a percentage
	PercentageToken
	// DimensionToken stores a number with a dimension
	DimensionToken
	// UnicodeRangeToken sets a range of hex numbers
	UnicodeRangeToken
	// IncludeMatchToken is ~=
	IncludeMatchToken
	// DashMatchToken is |=
	DashMatchToken
	// PrefixMatchToken is ^=
	PrefixMatchToken
	// SuffixMatchToken is $=
	SuffixMatchToken
	// SubstringMatchToken is *=
	SubstringMatchToken
	// ColumnToken : U+007C VERTICAL LINE (|)
	ColumnToken
	// WhitespaceToken : a U+000A LINE FEED (\n), U+0009 CHARACTER TABULATION (\t), or U+0020 SPACE ( ).
	WhitespaceToken
	// CDOToken : <!--
	CDOToken
	// CDCToken : -->
	CDCToken
	// ColonToken : U+003A COLON (:)
	ColonToken
	// SemicolonToken : U+003B SEMICOLON (;)
	SemicolonToken
	// CommaToken : U+002C COMMA (,)
	CommaToken
	// OpenSquareToken : U+005B LEFT SQUARE BRACKET ([)
	OpenSquareToken
	// CloseSquareToken : U+005D RIGHT SQUARE BRACKET (])
	CloseSquareToken
	// OpenParenToken : U+0028 LEFT PARENTHESIS (()
	OpenParenToken
	// CloseParenToken : U+0029 RIGHT PARENTHESIS ())
	CloseParenToken
	// OpenCurlyToken : U+007B LEFT CURLY BRACKET ({)
	OpenCurlyToken
	// CloseCurlyToken : U+007D RIGHT CURLY BRACKET (})
	CloseCurlyToken
)

// A Token consists of a TokenType, Value, and optional contextual data. startPos, endPos are the indices
// in the input CSS being parsed, that this token corresponds to, e.g. Tokenizer.input[startPos:endPos].
type Token struct {
	Type             TokenType
	Value            string
	Extra            string
	parent           *Tokenizer
	startPos, endPos int
}

// String returns the raw string representation for the token.
// The original input may not be preserved, due to preprocessing.
//
// This differs from the token's value. For example, the input of "@font-face"
// is tokenized as an AtKeywordToken, with a value of "font-face" (without the @ sign). However,
// its String() will be "@font-face".
func (t *Token) String() string {
	return t.parent.input[t.startPos:t.endPos]
}

// Tokenizer returns a stream of CSS Tokens.
type Tokenizer struct {
	input       string
	pos, length int
}

var preprocessReplacer = strings.NewReplacer(
	"\r\n", "\n",
	"\f", "\n",
	"\r", "\n",
	"\u0000", string(unicode.ReplacementChar),
)

// NewTokenizer returns a new CSS Tokenizer for the string, which is assumed to be UTF-8 encoded.
func NewTokenizer(input string) *Tokenizer {
	// 3.3 https://www.w3.org/TR/css-syntax-3/#input-preprocessing
	preprocessed := preprocessReplacer.Replace(input)
	return &Tokenizer{input: preprocessed, length: len(preprocessed)}
}

// Next scans the next token and returns it
func (z *Tokenizer) Next() Token {
	mark := z.pos
	t := z.consumeAToken()
	t.parent = z
	t.startPos = mark
	t.endPos = z.pos
	return t
}

// All returns all the tokens
func (z *Tokenizer) All() []Token {
	ret := []Token{}
	for {
		token := z.Next()
		ret = append(ret, token)
		if token.Type == EOFToken || token.Type == ErrorToken {
			break
		}
	}
	return ret
}

// consumeWhitespace returns a whitespace Token.
func (z *Tokenizer) consumeWhitespace() Token {
	start := z.pos
	for _, r := range z.input[start:] {
		if isWhitespace(r) {
			z.pos++
		} else {
			break
		}
	}
	return Token{Type: WhitespaceToken, Value: z.input[start:z.pos]}
}

// Put the rune back into the input stream.
func (z *Tokenizer) reconsume() {
	_, w := utf8.DecodeLastRuneInString(z.input[0:z.pos])
	z.pos -= w
}

// consume returns the next rune in the input stream. Repeated calls to consume() will return subsequent runes
// or a RuneError when the end is reached (or there is a problem decoding).
func (z *Tokenizer) consume() (rune, int) {
	if z.pos >= z.length {
		return utf8.RuneError, 0
	}
	r, w := utf8.DecodeRuneInString(z.input[z.pos:])
	z.pos += w
	return r, w
}

// peek returns the rune without consuming it. Repeated calls to peek() will return the same rune.
func (z *Tokenizer) peek() rune {
	return z.peekAt(0)
}

// peekAt returns the rune at the idx (relative to the current position) without consuming it.
// Repeated calls with same index value will return the same rune.
func (z *Tokenizer) peekAt(idx int) rune {
	if z.pos+idx > z.length {
		return utf8.RuneError
	}
	idx += z.pos
	r, _ := utf8.DecodeRuneInString(z.input[idx:])
	return r
}

// https://www.w3.org/TR/css-syntax-3/#digit
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// https://www.w3.org/TR/css-syntax-3/#hex-digit
func isHex(r rune) bool {
	if isDigit(r) {
		return true
	}
	if r >= 'A' && r <= 'F' {
		return true
	}
	return r >= 'a' && r <= 'f'
}

// https://www.w3.org/TR/css-syntax-3/#name-start-code-point
func isNameStart(r rune) bool {
	if r == '_' {
		return true
	}
	if r >= utf8.RuneSelf {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return r >= 'a' && r <= 'z'
}

// https://www.w3.org/TR/css-syntax-3/#name-code-point
func isName(r rune) bool {
	return isNameStart(r) || isDigit(r) || r == '-'
}

// https://www.w3.org/TR/css-syntax-3/#non-printable-code-point
func isNonPrintable(r rune) bool {
	if r >= '\u0000' && r <= '\u0008' {
		return true
	}
	if r == '\u000b' || r == '\u007f' {
		return true
	}
	return r >= '\u000e' && r <= '\u001f'
}

// isWhitespace returns true if the rune is whitespace as defined by
// https://www.w3.org/TR/css-syntax-3/#whitespace
func isWhitespace(r rune) bool {
	return r == '\n' || r == '\t' || r == ' '
}

// consumeAToken
// 4.3.1 https://www.w3.org/TR/css-syntax-3/#consume-a-token
func (z *Tokenizer) consumeAToken() Token {
	r, w := z.consume()
	if r == utf8.RuneError {
		if w == 0 {
			return Token{Type: EOFToken}
		}
		return Token{Type: ErrorToken, Value: "encoding error"}
	}

	if isWhitespace(r) {
		z.reconsume()
		return z.consumeWhitespace()
	}
	switch r {
	case '"', '\'':
		return z.consumeAString(r)
	case '#':
		if isName(z.peek()) || isValidEscape(z.peek(), z.peekAt(1)) {
			// TODO(alin04): Support hashtoken id type.
			return Token{Type: HashToken, Value: z.consumeAName()}
		}
		return Token{Type: DelimToken, Value: string(r)}
	case '$':
		if z.peek() == '=' {
			z.consume()
			return Token{Type: SuffixMatchToken, Value: "$="}
		}
		return Token{Type: DelimToken, Value: string(r)}
	case '(':
		return Token{Type: OpenParenToken, Value: string(r)}
	case ')':
		return Token{Type: CloseParenToken, Value: string(r)}
	case '*':
		if z.peek() == '=' {
			z.consume()
			return Token{Type: SubstringMatchToken, Value: "*="}
		}
		return Token{Type: DelimToken, Value: string(r)}
	case '+':
		if isANumber(r, z.peek(), z.peekAt(1)) {
			z.reconsume()
			return z.consumeANumeric()
		}
		return Token{Type: DelimToken, Value: string(r)}
	case ',':
		return Token{Type: CommaToken, Value: string(r)}
	case '-':
		if isANumber(r, z.peek(), z.peekAt(1)) {
			z.reconsume()
			return z.consumeANumeric()
		}
		if isAnIdentifier(r, z.peek(), z.peekAt(1)) {
			z.reconsume()
			return z.consumeAnIdentLike()
		}
		if z.peek() == '-' && z.peekAt(1) == '>' {
			z.consume()
			z.consume()
			return Token{Type: CDCToken, Value: "-->"}
		}
		return Token{Type: DelimToken, Value: string(r)}
	case '.':
		if isANumber(r, z.peek(), z.peekAt(1)) {
			z.reconsume()
			return z.consumeANumeric()
		}
		return Token{Type: DelimToken, Value: string(r)}
	case '/':
		if z.peek() == '*' {
			z.consume()
			for {
				eat, _ := z.consume()
				if eat == utf8.RuneError {
					break
				}
				if eat == '*' && z.peek() == '/' {
					z.consume()
					break
				}
			}
			return z.consumeAToken()
		}
		return Token{Type: DelimToken, Value: string(r)}
	case ':':
		return Token{Type: ColonToken, Value: string(r)}
	case ';':
		return Token{Type: SemicolonToken, Value: string(r)}
	case '<':
		if z.peek() == '!' && z.peekAt(1) == '-' && z.peekAt(2) == '-' {
			z.consume()
			z.consume()
			z.consume()
			return Token{Type: CDOToken, Value: "<!--"}
		}
		return Token{Type: DelimToken, Value: string(r)}
	case '@':
		if isAnIdentifier(z.peek(), z.peekAt(1), z.peekAt(2)) {
			return Token{Type: AtKeywordToken, Value: z.consumeAName()}
		}
		return Token{Type: DelimToken, Value: string(r)}
	case '[':
		return Token{Type: OpenSquareToken, Value: string(r)}
	case '\u005c': // U+005C REVERSE SOLIDUS (\)
		if isValidEscape(r, z.peek()) {
			z.reconsume()
			return z.consumeAnIdentLike()
		}
		return Token{Type: DelimToken, Value: "\\"}
	case ']':
		return Token{Type: CloseSquareToken, Value: string(r)}
	case '^':
		if z.peek() == '=' {
			z.consume()
			return Token{Type: PrefixMatchToken, Value: "^="}
		}
		return Token{Type: DelimToken, Value: string(r)}
	case '{':
		return Token{Type: OpenCurlyToken, Value: string(r)}
	case '}':
		return Token{Type: CloseCurlyToken, Value: string(r)}
	}

	if isDigit(r) {
		z.reconsume()
		return z.consumeANumeric()
	}

	if r == 'U' || r == 'u' {
		if z.peek() == '+' && (isHex(z.peekAt(1)) || z.peekAt(1) == '?') {
			z.consume()
			return z.consumeAUnicodeRange()
		}
		z.reconsume()
		return z.consumeAnIdentLike()
	}

	if isNameStart(r) {
		z.reconsume()
		return z.consumeAnIdentLike()
	}

	switch r {
	case '|':
		next := z.peek()
		if next == '=' {
			z.consume()
			return Token{Type: DashMatchToken, Value: "|="}
		} else if next == '|' {
			z.consume()
			return Token{Type: ColumnToken, Value: "||"}
		}
		return Token{Type: DelimToken, Value: string(r)}
	case '~':
		if z.peek() == '=' {
			return Token{Type: IncludeMatchToken, Value: "~="}
		}
		return Token{Type: DelimToken, Value: string(r)}
	}
	return Token{Type: DelimToken, Value: string(r)}
}

// 4.3.2 https://www.w3.org/TR/css-syntax-3/#consume-a-numeric-token
func (z *Tokenizer) consumeANumeric() Token {
	repr, _ := z.consumeANumber()
	r := z.peek()
	if r == utf8.RuneError {
		return Token{Type: NumberToken, Value: repr}
	}
	if isAnIdentifier(r, z.peekAt(1), z.peekAt(2)) {
		return Token{Type: DimensionToken, Value: repr, Extra: z.consumeAName()}
	}
	if r == '%' {
		z.consume()
		return Token{Type: PercentageToken, Value: repr}
	}
	return Token{Type: NumberToken, Value: repr}
}

// 4.3.3 https://www.w3.org/TR/css-syntax-3/#consume-an-ident-like-token
func (z *Tokenizer) consumeAnIdentLike() Token {
	name := z.consumeAName()
	if strings.EqualFold(name, "url") && z.peek() == '(' {
		z.consume()
		return z.consumeAURL()
	}
	if z.peek() == '(' {
		z.consume()
		return Token{Type: FunctionToken, Value: name}
	}
	return Token{Type: IdentToken, Value: name}
}

// consumeAString
// 4.3.4 https://www.w3.org/TR/css-syntax-3/#consume-a-string-token
func (z *Tokenizer) consumeAString(endingCodePoint rune) Token {
	t := Token{Type: StringToken, Extra: string(endingCodePoint)}
	var sb strings.Builder
	for {
		r, _ := z.consume()
		if r == utf8.RuneError {
			break
		}
		if r == endingCodePoint {
			break
		} else if r == '\n' {
			z.reconsume()
			t.Type = BadStringToken
			break
		} else if r == '\u005c' { // U+005C REVERSE SOLIDUS (\)
			if z.pos+1 == z.length {
				// do nothing
				continue
			}
			next := z.peek()
			if next == '\n' {
				z.consume()
			} else if isValidEscape(r, next) {
				sb.WriteRune(z.consumeAnEscape())
			} else {
				// not an escape, so append
				sb.WriteRune(r)
			}
		} else {
			sb.WriteRune(r)
		}
	}
	t.Value = sb.String()
	return t
}

// 4.3.5 https://www.w3.org/TR/css-syntax-3/#consume-a-url-token
func (z *Tokenizer) consumeAURL() Token {
	ret := Token{Type: URLToken}
	z.consumeWhitespace()
	next := z.peek()
	switch next {
	case utf8.RuneError:
		return ret
	case '"', '\'':
		z.consume()
		token := z.consumeAString(next)
		if token.Type == BadStringToken {
			return z.consumeRemnantsOfBadURL()
		}
		ret.Value = token.Value
		z.consumeWhitespace()
		if z.peek() == ')' || z.peek() == utf8.RuneError {
			z.consume()
			return ret
		}
		return z.consumeRemnantsOfBadURL()
	default:
		for {
			r, _ := z.consume()
			if r == ')' || r == utf8.RuneError {
				return ret
			}
			if isWhitespace(r) {
				z.consumeWhitespace()
				if z.peek() == ')' || z.peek() == utf8.RuneError {
					z.consume()
					return ret
				}
				return z.consumeRemnantsOfBadURL()
			}
			if r == '"' || r == '\'' || r == '(' || isNonPrintable(r) {
				return z.consumeRemnantsOfBadURL()
			}
			if r == '\u005c' {
				if isValidEscape(r, z.peek()) {
					cp := z.consumeAnEscape()
					ret.Value += string(cp)
				} else {
					return z.consumeRemnantsOfBadURL()
				}
			} else {
				ret.Value += string(r)
			}
		}
	}
}

const maxHexDigits = 6

// 4.3.6 https://www.w3.org/TR/css-syntax-3/#consume-a-unicode-range-token
// TODO(alin04): The UnicodeRange Token doesn't store separate values for start and end as per the spec.
func (z *Tokenizer) consumeAUnicodeRange() Token {
	start := make([]byte, maxHexDigits)
	end := make([]byte, maxHexDigits)
	var foundQ bool
	i := 0
	for ; i < maxHexDigits; i++ {
		r, _ := z.consume()
		if r == utf8.RuneError {
			break
		} else if isHex(r) {
			utf8.EncodeRune(start[i:], r)
			utf8.EncodeRune(end[i:], r)
		} else if r == '?' {
			foundQ = true
			start[i] = '0'
			end[i] = 'F'
		} else {
			z.reconsume()
			break
		}
	}
	startCP, err := strconv.ParseInt(string(start[:i]), 16, 32)
	if err != nil {
		return Token{Type: ErrorToken, Value: err.Error()}
	}
	if !foundQ && z.peek() == '-' {
		z.consume()
		i = 0
		for ; i < maxHexDigits; i++ {
			r, _ := z.consume()
			if r == utf8.RuneError {
				break
			} else if isHex(r) {
				utf8.EncodeRune(end[i:], r)
			} else {
				z.reconsume()
				break
			}
		}
	}
	endCP, err := strconv.ParseInt(string(end[:i]), 16, 32)
	if err != nil {
		return Token{Type: ErrorToken, Value: err.Error()}
	}
	if startCP == endCP {
		return Token{Type: UnicodeRangeToken, Value: fmt.Sprintf("U+%04X", startCP)}
	}
	return Token{Type: UnicodeRangeToken, Value: fmt.Sprintf("U+%04X-%04X", startCP, endCP)}
}

// 4.3.7 https://www.w3.org/TR/css-syntax-3/#consume-an-escaped-code-point
func (z *Tokenizer) consumeAnEscape() rune {
	r, _ := z.consume()
	if r == utf8.RuneError {
		return r
	} else if isHex(r) {
		digits := make([]byte, maxHexDigits)
		utf8.EncodeRune(digits, r)
		i := 1
		for ; i < maxHexDigits; i++ {
			r, _ = z.consume()
			if r == utf8.RuneError {
				break
			} else if isHex(r) {
				utf8.EncodeRune(digits[i:], r)
			} else {
				z.reconsume()
				break
			}
		}
		if isWhitespace(z.peek()) {
			z.consume()
		}
		cp, err := strconv.ParseInt(string(digits[:i]), 16, 32)
		if err != nil || cp == 0 || cp > unicode.MaxRune || unicode.Is(unicode.Cs, rune(cp)) {
			return unicode.ReplacementChar
		}
		return rune(cp)
	}
	return r
}

// 4.3.8 https://www.w3.org/TR/css-syntax-3/#starts-with-a-valid-escape
func isValidEscape(r1, r2 rune) bool {
	if r1 != '\u005c' {
		return false
	}
	if r2 == '\n' {
		return false
	}
	return true
}

// 4.3.9 https://www.w3.org/TR/css-syntax-3/#would-start-an-identifier
func isAnIdentifier(r1, r2, r3 rune) bool {
	if r1 == '-' {
		r1 = r2
		r2 = r3
	}
	if r1 == utf8.RuneError {
		return false
	}
	if isNameStart(r1) {
		return true
	}
	return isValidEscape(r1, r2)
}

// 4.3.10 https://www.w3.org/TR/css-syntax-3/#starts-with-a-number
func isANumber(r1, r2, r3 rune) bool {
	if r1 == '+' || r1 == '-' {
		r1 = r2
		r2 = r3
	}
	if r1 == '.' {
		r1 = r2
	}
	return isDigit(r1)
}

// 4.3.11 https://www.w3.org/TR/css-syntax-3/#consume-a-name
func (z *Tokenizer) consumeAName() string {
	var sb strings.Builder
	for {
		r, _ := z.consume()
		if r == utf8.RuneError {
			break
		}

		if isName(r) {
			sb.WriteRune(r)
		} else if isValidEscape(r, z.peek()) {
			sb.WriteRune(z.consumeAnEscape())
		} else {
			z.reconsume()
			break
		}
	}
	return sb.String()
}

type numberType int8

const (
	integer numberType = iota
	number
)

// 4.3.12 https://www.w3.org/TR/css-syntax-3/#consume-a-number
// TODO(alin04): This doesn't return the actual numeric value.
func (z *Tokenizer) consumeANumber() (string, numberType) {
	var repr strings.Builder
	nType := integer

	if z.peek() == '+' || z.peek() == '-' {
		z.consumeAndWrite(&repr, 1)
	}
	z.consumeDigits(&repr)
	if z.peek() == '.' && isDigit(z.peekAt(1)) {
		z.consumeAndWrite(&repr, 2)
		nType = number
		z.consumeDigits(&repr)
	}
	if z.peek() == 'E' || z.peek() == 'e' {
		toConsume := 1
		r := z.peekAt(1)
		if r == '+' || r == '-' {
			toConsume++
			r = z.peekAt(2)
		}
		if isDigit(r) {
			nType = number
			z.consumeAndWrite(&repr, toConsume)
			z.consumeDigits(&repr)
		}
	}
	return repr.String(), nType
}

// 4.3.14 https://www.w3.org/TR/css-syntax-3/#consume-the-remnants-of-a-bad-url
func (z *Tokenizer) consumeRemnantsOfBadURL() Token {
	for {
		r, _ := z.consume()
		if r == ')' || r == utf8.RuneError {
			return Token{Type: BadURLToken}
		}
		if isValidEscape(r, z.peek()) {
			z.consumeAnEscape()
		}
	}
}

// consumeAndWrite consumes n runes and appends to the string builder.
func (z *Tokenizer) consumeAndWrite(sb *strings.Builder, n int) {
	for i := 0; i < n; i++ {
		r, _ := z.consume()
		sb.WriteRune(r)
	}
}

func (z *Tokenizer) consumeDigits(sb *strings.Builder) {
	for {
		if isDigit(z.peek()) {
			z.consumeAndWrite(sb, 1)
		} else {
			break
		}
	}
}
