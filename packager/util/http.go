package util

import (
	"regexp"
	"strings"
)

// A comma, as defined in https://tools.ietf.org/html/rfc7230#section-7, with
// OWS defined in https://tools.ietf.org/html/rfc7230#appendix-B. This is
// commonly used as a separator in header field value definitions.
var Comma *regexp.Regexp = regexp.MustCompile(`[ \t]*,[ \t]*`)

// Trim optional whitespace from a header value, adhering to
// https://tools.ietf.org/html/rfc7230#section-7 with OWS defined in
// https://tools.ietf.org/html/rfc7230#appendix-B.
func TrimHeaderValue(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return r == ' ' || r == '\t'
	})
}
