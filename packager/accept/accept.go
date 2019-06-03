package accept

import (
	"log"
	"mime"
	"strings"

	"github.com/WICG/webpackage/go/signedexchange/version"
	"github.com/ampproject/amppackager/packager/util"
)

// The SXG version that packager can produce. In the future, it may need to be
// able to produce multiple versions.
const AcceptedSxgVersion = "b3"

// The Content-Type for the SXG version that the signer produces.
const SxgContentType = "application/signed-exchange;v=" + AcceptedSxgVersion

// The enum of the SXG version that the signer produces, for passing to the
// signedexchange library.
var SxgVersion = version.Version1b3

// Tokenize a comma-separated string of accept patterns into a slice
func tokenize(accept string) []string {
	var tokens []string
	acceptLen := len(accept)
	if acceptLen == 0 {
		return tokens
	}

	inQuotes := false
	startIndex := 0
	for i := 0; i < acceptLen; i++ {
		char := accept[i]
		switch char {
		case '"':
			inQuotes = !inQuotes
		case ',':
			if !inQuotes {
				tokens = append(tokens, util.TrimHeaderValue(accept[startIndex:i]))
				startIndex = i + 1
			}
		case '\\':
			if !inQuotes {
				log.Printf("unable to parse Accept header: %s", accept)
				return []string{}
			}
			i++
		}
	}
	tokens = append(tokens, util.TrimHeaderValue(accept[startIndex:]))
	return tokens
}

// Determine whether a version specified by the accept header matches the
// version of signed exchange output by the packager
func hasMatchingSxgVersion(versions []string) bool {
	for _, version := range versions {
		if version == AcceptedSxgVersion {
			return true
		}
	}
	return false
}

// True if the given Accept header is one that the packager can satisfy. It
// must contain application/signed-exchange;v=$V so that the packager knows
// whether or not it can supply the correct version. "" and "*/*" are not
// satisfiable, for this reason.
func CanSatisfy(accept string) bool {
	types := tokenize(accept)
	for _, mediaRange := range types {
		mediatype, params, err := mime.ParseMediaType(mediaRange)
		if err == nil && mediatype == "application/signed-exchange" {
			if hasMatchingSxgVersion(strings.Split(params["v"], ",")) {
				return true
			}
		}
	}
	return false
}
