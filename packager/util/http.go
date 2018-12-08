package util

import (
	"regexp"
)

// A comma, as defined in https://tools.ietf.org/html/rfc7230#section-7, with
// OWS defined in https://tools.ietf.org/html/rfc7230#appendix-B. This is
// commonly used as a separator in header field value definitions.
var Comma *regexp.Regexp = regexp.MustCompile(`[ \t]*,[ \t]*`)
