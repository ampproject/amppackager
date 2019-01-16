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
	"errors"
	"strings"
)

// A SegmentType is the type of a Segment.
type SegmentType uint32

const (
	// ByteType is everything not of the other types below.
	ByteType SegmentType = iota
	// ImageURLType is a URL for an image. Note that the value has been
	// stripped of the surrounding 'url' identity token. To reconstruct a stylesheet,
	// this value must be enclosed into a url token or function token with name 'url'.
	ImageURLType
	// FontURLType is similar to an image URL, but for fonts.
	FontURLType
)

// Segments is a slice of individual Segment structs.
type Segments []Segment

// Segment is a portion of CSS.
type Segment struct {
	Type SegmentType
	Data string
}

const (
	quoteOrWhitespace = "\"' \t\n\f\r"
)

// ParseURLs chops a style sheet into Segments. Each segment  is
// either a UTF8 encoded byte string, or an image or font URL.
// This is used to modify the URLs to point at a CDN.
// Note that when combining the segments back to a stylesheet,
// the client code must emit url() around URLs. This is done so that
// client code can choose the quote character as in
// url("http://foo.com") or url('http://foo.com/') or even leave out
// the quote character as in url(http://foo.com/). Note that CSS supports
// escaping quote characters within a string by prefixing with a backslash,
// so " inside a URL may be written as \".
func ParseURLs(css string) (Segments, error) {
	z := NewTokenizer(css)
	segments := Segments{}
	var sb strings.Builder
	var endOfFontFaceIdx int
	tokens := z.All()
loop:
	for i, token := range tokens {
		switch token.Type {
		case EOFToken:
			break loop
		case ErrorToken:
			return Segments{}, errors.New(token.Value)
		case AtKeywordToken:
			if token.Value == "font-face" {
				endOfFontFaceIdx = i + consumeAnAtRule(tokens[i:])
			}
		case URLToken:
			// Emit a segment which contains all non-URL CSS seen so far.
			if sb.Len() > 0 {
				segments = append(segments, Segment{ByteType, sb.String()})
				sb.Reset()
			}
			// Now emit a URL segment
			t := ImageURLType
			if endOfFontFaceIdx > i {
				t = FontURLType
			}
			segments = append(segments, Segment{t, token.Value})
			continue
		}
		sb.WriteString(token.String())
	}
	segments = append(segments, Segment{ByteType, sb.String()})
	return segments, nil
}

// consumeAnAtRule returns the index which marks the end of the at rule as per
// 5.4.2 https://www.w3.org/TR/css-syntax-3/#consume-an-at-rule
func consumeAnAtRule(tokens []Token) int {
	if len(tokens) == 0 || tokens[0].Type != AtKeywordToken {
		return -1 // should be impossible
	}
	i := 1
	for ; i < len(tokens); i++ {
		if tokens[i].Type == SemicolonToken || tokens[i].Type == EOFToken {
			return i
		}
		if tokens[i].Type == OpenCurlyToken {
			i += consumeASimpleBlock(tokens[i:])
			return i
		}
		i += consumeAComponentValue(tokens[i:])
	}
	return i
}

// consumeAComponentValue returns the index that marks the end of the component value as per
// 5.4.6 https://www.w3.org/TR/css-syntax-3/#consume-a-component-value
func consumeAComponentValue(tokens []Token) int {
	if len(tokens) == 0 {
		return -1
	}
	switch tokens[0].Type {
	case OpenCurlyToken, OpenSquareToken, OpenParenToken:
		return consumeASimpleBlock(tokens)
	case FunctionToken:
		return consumeAFunction(tokens)
	default:
		return 0
	}
}

// consumeASimpleBlock returns the index which marks the end of the block, or -1 if the tokens are empty as per
// 5.4.7 https://www.w3.org/TR/css-syntax-3/#consume-a-simple-block
func consumeASimpleBlock(tokens []Token) int {
	if len(tokens) == 0 {
		return -1
	}
	var endingTokenType TokenType
	switch tokens[0].Type {
	case OpenCurlyToken:
		endingTokenType = CloseCurlyToken
	case OpenParenToken:
		endingTokenType = CloseParenToken
	case OpenSquareToken:
		endingTokenType = CloseSquareToken
	}
	i := 1
	for ; i < len(tokens); i++ {
		if tokens[i].Type == EOFToken || tokens[i].Type == endingTokenType {
			return i
		}
		i += consumeAComponentValue(tokens[i:])
	}
	return i
}

// consumeAFunction returns the index marking the end of the function as defined by:
// 5.4.8 https://www.w3.org/TR/css-syntax-3/#consume-a-function
func consumeAFunction(tokens []Token) int {
	if len(tokens) == 0 || tokens[0].Type != FunctionToken {
		return -1 // should be impossible case.
	}
	i := 1
	for ; i < len(tokens); i++ {
		if tokens[i].Type == EOFToken || tokens[i].Type == CloseParenToken {
			return i
		}
		i += consumeAComponentValue(tokens[i:])
	}
	return i
}
