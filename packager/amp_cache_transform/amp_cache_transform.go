// Parses and interprets the AMP-Cache-Transform header specified at
// https://github.com/ampproject/amphtml/blob/master/spec/amp-cache-transform.md.

package amp_cache_transform

import (
	"io"
	"strings"

	"github.com/pkg/errors"
)

// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-3.3
type parameterisedIdentifier struct {
	id string
	// In the future, this will include params.
}

// https://tools.ietf.org/html/rfc7230#appendix-B (OWS = "optional whitespace")
func discardOWS(reader *strings.Reader) error {
	for reader.Len() > 0 {
		c, err := reader.ReadByte()
		if err != nil {
			return errors.Wrap(err, "reading OWS")
		}
		if c != ' ' && c != '\t' {
			// Return to the position before the non-whitespace char.
			reader.Seek(-1, io.SeekCurrent)
			return nil
		}
	}
	return nil
}

// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-4.2.3
func parseParameterisedList(reader *strings.Reader) ([]parameterisedIdentifier, error) {
	items := []parameterisedIdentifier{}
	for reader.Len() > 0 {
		item, err := parseParameterisedIdentifier(reader)
		if err != nil {
			return nil, errors.Wrap(err, "parsing parameterised identifier")
		}
		items = append(items, *item)
		// Steps 3 and 4 are swapped per https://github.com/httpwg/http-extensions/pull/667.
		if reader.Len() == 0 {
			return items, nil
		}
		if err := discardOWS(reader); err != nil {
			return nil, errors.Wrap(err, "discarding OWS")
		}
		comma, err := reader.ReadByte()
		if err != nil {
			return nil, errors.Wrap(err, "reading comma")
		}
		if comma != ',' {
			return nil, errors.New("expected comma")
		}
		if err := discardOWS(reader); err != nil {
			return nil, errors.Wrap(err, "discarding OWS")
		}
		if reader.Len() == 0 {
			return nil, errors.New("expected another param-id")
		}
	}
	return nil, errors.New("expected non-empty parameterised list")
}

// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-4.2.4
func parseParameterisedIdentifier(reader *strings.Reader) (*parameterisedIdentifier, error) {
	primaryID, err := parseIdentifier(reader)
	if err != nil {
		return nil, errors.Wrap(err, "parsing primary identifier")
	}
	// NOTE: The initial version of AMP-Cache-Transform does not support
	// any parameters, so we don't bother to implement a parser for them.
	// No point in distinguishing syntactic failure from semantic failure.
	// Once we add parameters, we'll need to implement parsers for the
	// subset item that they use. See Items here:
	// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-3.4
	return &parameterisedIdentifier{primaryID}, nil
}

// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-3.8
func isLCAlpha(c byte) bool {
	return c >= 'a' && c <= 'z'
}

// https://tools.ietf.org/html/rfc5234#appendix-B.1
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-3.8
func isSecondIdentifierChar(c byte) bool {
	return isLCAlpha(c) || isDigit(c) || c == '_' || c == '-' || c == '*' || c == '/'
}

// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-4.2.8
func parseIdentifier(reader *strings.Reader) (string, error) {
	var output strings.Builder
	char, err := reader.ReadByte()
	if err != nil {
		return "", errors.Wrap(err, "reading byte")
	}
	if !isLCAlpha(char) {
		return "", errors.New("expected lowercase alpha")
	}
	output.WriteByte(char)
	for reader.Len() > 0 {
		char, err := reader.ReadByte()
		if err != nil {
			return "", errors.Wrap(err, "reading byte")
		}
		if !isSecondIdentifierChar(char) {
			// Return to the position before the non-identifier char.
			reader.Seek(-1, io.SeekCurrent)
			return output.String(), nil
		}
		output.WriteByte(char)
	}
	return output.String(), nil
}

// Returns true if the given AMP-Cache-Transform request header value is one
// the packager can satisfy.
func ShouldSendSXG(header_value string) bool {
	reader := strings.NewReader(header_value)
	identifiers, err := parseParameterisedList(reader)
	if err != nil {
		// TODO(twifkak): Debug-log err and reader.Len() bytes remaining.
		return false
	}
	for _, identifier := range identifiers {
		if identifier.id == "any" || identifier.id == "google" {
			return true
		}
	}
	return false
}
