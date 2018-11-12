package accept

import (
	"mime"
	"regexp"

	"github.com/WICG/webpackage/go/signedexchange/version"
)

// The SXG version that packager can produce. In the future, it may need to be
// able to produce multiple versions.
const AcceptedSxgVersion = "b2"

// The Content-Type for the SXG version that the signer produces.
const SxgContentType = "application/signed-exchange;v=" + AcceptedSxgVersion

// The enum of the SXG version that the signer produces, for passing to the
// signedexchange library.
var SxgVersion = version.Version1b2

// A comma, as would appear in an Accept header. Comma-separation is defined
// in https://tools.ietf.org/html/rfc7230#section-7, with OWS defined in
// https://tools.ietf.org/html/rfc7230#appendix-B.
//
// There is an edge case on which this fails:
//   Accept: application/signed-exchange;junk="some,thing";v=b2
// However, in practice, browsers don't send media types with quoted
// commas in them:
//   https://developer.mozilla.org/en-US/docs/Web/HTTP/Content_negotiation/List_of_default_Accept_values
// So we'll live with this deficiency for the sake of not forking
// mime.ParseMediaType.
var comma *regexp.Regexp = regexp.MustCompile(`[ \t]*,[ \t]*`)

// True if the given Accept header is one that the packager can satisfy. It
// must contain application/signed-exchange;v=$V so that the packager knows
// whether or not it can supply the correct version. "" and "*/*" are not
// satisfiable, for this reason.
func CanSatisfy(accept string) bool {
	types := comma.Split(accept, -1)
	for _, mediaRange := range types {
		mediatype, params, err := mime.ParseMediaType(mediaRange)
		if err == nil && mediatype == "application/signed-exchange" && params["v"] == AcceptedSxgVersion {
			return true
		}
	}
	return false
}
