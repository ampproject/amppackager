package util

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/pkg/errors"
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

// Conditional request headers that ServeHTTP may receive and need to be sent with fetchURL.
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Conditional_requests#Conditional_headers
var ConditionalRequestHeaders = map[string]bool{
	"If-Match":            true,
	"If-None-Match":       true,
	"If-Modified-Since":   true,
	"If-Unmodified-Since": true,
	"If-Range":            true,
}

// The following hop-by-hop headers should be removed even when not specified
// in Connection, for backwards compatibility with downstream servers that were
// written against RFC 2616, and expect gateways to behave according to
// https://tools.ietf.org/html/rfc2616#section-13.5.1. (Note: "Trailers" is a
// typo there; should be "Trailer".)
//
// Connection header should also be removed per
// https://tools.ietf.org/html/rfc7230#section-6.1.
//
// Proxy-Connection should also be deleted, per
// https://github.com/WICG/webpackage/pull/339.
var legacyHeaders = map[string]bool{
	"Connection":         true,
	"Keep-Alive":         true,
	"Proxy-Authenticate": true,
	"Proxy-Connection":   true,
	"Trailer":            true,
	"Transfer-Encoding":  true,
	"Upgrade":            true,
}

// Via is implicitly forwarded and disallowed to be included in
// config.ForwardedRequestHeaders
// TE is a hop-by-hop request header and must not be forwarded.
// Proxy-Authorization can be forwarded per rfc7235#section-4.4 but
// remove it to mitigate the risk of over-signing.
var notForwardedRequestHeader = map[string]bool{
	"Proxy-Authorization": true,
	"Te":                  true,
	"Via":                 true,
}

// Remove hop-by-hop headers, per https://tools.ietf.org/html/rfc7230#section-6.1.
func RemoveHopByHopHeaders(h http.Header) {
	if connections, ok := h[http.CanonicalHeaderKey("Connection")]; ok {
		for _, connection := range connections {
			headerNames := Comma.Split(connection, -1)
			for _, headerName := range headerNames {
				h.Del(headerName)
			}
		}
	}

	for headerName, _ := range legacyHeaders {
		h.Del(headerName)
	}
}

func haveInvalidForwardedRequestHeader(h string) string {
	if _, ok := legacyHeaders[http.CanonicalHeaderKey(h)]; ok {
		return fmt.Sprintf("have hop-by-hop header of %s", h)
	}
	if _, ok := ConditionalRequestHeaders[http.CanonicalHeaderKey(h)]; ok {
		return fmt.Sprintf("have conditional request header of %s", h)
	}
	if _, ok := notForwardedRequestHeader[http.CanonicalHeaderKey(h)]; ok {
		return fmt.Sprintf("include request header of %s", h)
	}
	return ""
}

// Escapes the input and surrounds it in quotes, so it's a valid quoted-string,
// per https://tools.ietf.org/html/rfc7230#section-3.2.6. Returns error if the
// input contains any chars outside of HTAB / SP / VCHAR
// (https://tools.ietf.org/html/rfc5234#appendix-B.1) and thus isn't even
// quotable.
func QuotedString(input string) (string, error) {
	var ret strings.Builder
	ret.WriteByte('"')
	for i := 0; i < len(input); i++ {
		b := input[i]
		if (b < 0x20 || b > 0x7e) && b != 0x09 {
			return "", errors.New("contains non-printable char")
		}
		if b == '"' || b == '\\' {
			ret.Write([]byte{'\\', b})
		} else {
			ret.WriteByte(b)
		}
	}
	ret.WriteByte('"')
	return ret.String(), nil
}
