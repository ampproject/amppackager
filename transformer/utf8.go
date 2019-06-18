package transformer

import (
	"unicode/utf8"

	"github.com/pkg/errors"
)

// False if the code point is known to cause parse errors during HTML
// preprocessing, per
// https://html.spec.whatwg.org/multipage/parsing.html#preprocessing-the-input-stream
//
// Also false for U+0000 NULL, as that causes parse errors everywhere except
// CDATA, and for defense in depth we don't assume that all parsers interpret
// this properly.
func isHTMLValidInternal(r rune) bool {
	// In order to reduce the average number of comparisons per code point,
	// test for validity (OR of ANDs) rather than invalidity (AND of ORs),
	// and check popular ranges first.
	return (
		// Invalid chars:
		// U+0000 NULL, per above logic.
		// U+0001 through U+001F, except 0x9, 0xA, 0xC, 0xD, per https://infra.spec.whatwg.org/#control.
		(r > 0x1F && r < 0x7F) || r == 0x9 || r == 0xA || r == 0xC || r == 0xD ||
		// U+007F through U+009F, per https://infra.spec.whatwg.org/#control.
		(r > 0x9F && r < 0xD800) ||
		// U+D800 through U+DFFF, per https://infra.spec.whatwg.org/#surrogate.
		(r > 0xDFFF && r < 0xFDD0) ||
		// U+FDD0 through U+FDEF, per https://infra.spec.whatwg.org/#noncharacter.
		(r > 0xFDEF && r < 0xFFFE) ||
		// U+??FFFE and U+??FFFF, per https://infra.spec.whatwg.org/#noncharacter.
		(r > 0xFFFF && r < 0x10FFFE && r & 0xFFFE != 0xFFFE))
		// There are no codepoints greater than U+10FFFF.
}

// Overrideable for test.
var isHTMLValid = isHTMLValidInternal

// Returns an error if the given string is not well-formed UTF-8, or contains
// characters known to cause parse errors in HTML. This requirement is imposed
// by the AMPHTML validator, so it doesn't make sense to create a SXG.
func validateUTF8ForHTML(html string) error {
	pos := 0
	for pos < len(html) {
		r, width := utf8.DecodeRuneInString(html[pos:])
		// Check that the code point wasn't ill-formed. utf8.RuneError
		// == '\uFFFD' so we need to check for a mismatched width, too.
		if r == utf8.RuneError && width < 2 {
			return errors.Errorf("invalid UTF-8 at byte position %d", pos)
		}
		if !isHTMLValid(r) {
			return errors.Errorf("character U+%04x at position %d is not allowed in AMPHTML", r, pos)
		}
		pos += width
	}
	return nil
}
