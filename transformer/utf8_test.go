package transformer

import (
	"testing"

	tt "github.com/ampproject/amppackager/transformer/internal/testing"
)

var minimumValidAMP = tt.Concat(
	tt.Doctype, "<html ⚡><head>",
	tt.MetaCharset, tt.MetaViewport, tt.ScriptAMPRuntime,
	tt.LinkFavicon, tt.LinkCanonical, tt.StyleAMPBoilerplate,
	tt.NoscriptAMPBoilerplate, "</head><body></body></html>",
)

// True if the code point is known to cause parse errors during HTML
// preprocessing, per
// https://html.spec.whatwg.org/multipage/parsing.html#preprocessing-the-input-stream,
// or if it is U+0000 NULL.
//
// This is easier to visually inspect and compare against the spec, so it's
// used as a slower implementation to test against.
func isHTMLInvalid(r rune) bool {
	return (
		// U+0000 NULL + https://infra.spec.whatwg.org/#control
		(r <= 0x1F && r != 0x9 && r != 0xA && r != 0xC && r != 0xD) ||
		(r >= 0x7F && r <= 0x9F) ||
		// https://infra.spec.whatwg.org/#surrogate
		(r >= 0xD800 && r <= 0xDFFF) ||
		// https://infra.spec.whatwg.org/#noncharacter
		(r >= 0xFDD0 && r <= 0xFDEF) ||
		(r >= 0xFFFE && r <= 0x10FFFF && r & 0xFFFE == 0xFFFE) ||
		// http://unicode.org/glossary/#codespace
		(r >= 0x110000))
}

func TestIsHTMLValid(t *testing.T) {
	for r := '\000'; r <= 0x110000; r++ {
		want := !isHTMLInvalid(r)
		got := isHTMLValid(r)
		if got != want {
			t.Errorf("IsHTMLValid(U+%06x) got=%t, want=%t", r, got, want)
		}
	}
}

func TestValidateUTF8ForHTMLAllowsReplacementCharacter(t *testing.T) {
	html := "\uFFFD"
	if err := validateUTF8ForHTML(html); err != nil {
		t.Errorf("validateUTF8ForHTML(U+FFFD) error=%q", err)
	}
}

func BenchmarkIsHTMLValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validateUTF8ForHTML(minimumValidAMP)
	}
}

func BenchmarkIsHTMLInvalid(b *testing.B) {
	isHTMLValid = func(r rune) bool { return !isHTMLInvalid(r) }
	defer func() { isHTMLValid = isHTMLValidInternal }()
	for i := 0; i < b.N; i++ {
		validateUTF8ForHTML(minimumValidAMP)
	}
}
