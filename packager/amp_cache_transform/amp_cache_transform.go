// Parses and interprets the AMP-Cache-Transform header specified at
// https://github.com/ampproject/amphtml/blob/master/spec/amp-cache-transform.md.

package amp_cache_transform

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/ampproject/amppackager/packager/util"
	"github.com/ampproject/amppackager/transformer"
	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/pkg/errors"
)

// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-3.3
type parameterisedIdentifier struct {
	id     string
	params map[string]string
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

const parameterisedListSeparator = ','

// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-4.2.3
func parseParameterisedList(reader *strings.Reader) ([]parameterisedIdentifier, error) {
	items := []parameterisedIdentifier{}
	for reader.Len() > 0 {
		item, err := parseParameterisedIdentifier(reader)
		if err != nil {
			return nil, errors.Wrap(err, "parsing parameterised identifier")
		}
		items = append(items, *item)
		// Steps 3 and 4 are swapped per
		// https://github.com/httpwg/http-extensions/pull/667 and
		// https://tools.ietf.org/html/draft-thomson-postel-was-wrong-00.
		if reader.Len() == 0 {
			return items, nil
		}
		if err := discardOWS(reader); err != nil {
			return nil, errors.Wrap(err, "discarding OWS")
		}
		if char, err := reader.ReadByte(); err != nil {
			return nil, errors.Wrapf(err, "reading '%c'", parameterisedListSeparator)
		} else if char != parameterisedListSeparator {
			return nil, errors.Errorf("expected '%c'", parameterisedListSeparator)
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

const (
	parameterSeparator = ';'
	parameterValueSeparator = '='
)

// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-4.2.4
func parseParameterisedIdentifier(reader *strings.Reader) (*parameterisedIdentifier, error) {
	primaryID, err := parseIdentifier(reader)
	if err != nil {
		return nil, errors.Wrap(err, "parsing primary identifier")
	}
	// NOTE: The current version of AMP-Cache-Transform only uses
	// string-valued parameters. Future versions of the spec may require
	// parsing other Item types.
	params := map[string]string{}
	for {
		if err := discardOWS(reader); err != nil {
			return nil, errors.Wrap(err, "discarding OWS")
		}

		// If we've reach the end of the parameter list, exit success.
		if reader.Len() == 0 {
			break
		}
		if char, err := reader.ReadByte(); err != nil {
			return nil, errors.Wrapf(err, "reading '%c'", parameterSeparator)
		} else if char != parameterSeparator {
			if err := reader.UnreadByte(); err != nil {
				return nil, errors.Wrapf(err, "unreading '%c'", parameterSeparator)
			}
			break
		}

		// Otherwise, parse a parameter name.
		if err := discardOWS(reader); err != nil {
			return nil, errors.Wrap(err, "discarding OWS")
		}
		name, err := parseIdentifier(reader)
		if err != nil {
			return nil, errors.Wrap(err, "parsing param name")
		}
		if _, has := params[name]; has {
			return nil, errors.Errorf("param %q already present", name)
		}

		// ... and a parameter value.
		// NOTE: The current version of the AMP-Cache-Transform spec
		// does not use any null-valued parameters. So, for now, a
		// missing '=' is a parse failure.
		if char, err := reader.ReadByte(); err != nil {
			return nil, errors.Wrapf(err, "reading '%c'", parameterValueSeparator)
		} else if char != parameterValueSeparator {
			return nil, errors.Errorf("expected '%c'", parameterValueSeparator)
		}
		// NOTE: In the future, this may need to be parseItem() to
		// support other Item types:
		// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-4.2.5
		value, err := parseString(reader)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing value for param %s", name)
		}

		params[name] = value
	}
	return &parameterisedIdentifier{primaryID, params}, nil
}

const (
	stringDelimiter = '"'
	stringEscape = '\\'
)

func invalidStringChar(char byte) bool {
	return char <= 0x1f || char == 0x7f
}

// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-4.2.7
func parseString(reader *strings.Reader) (string, error) {
	if char, err := reader.ReadByte(); err != nil {
		return "", errors.Wrapf(err, "reading '%c'", stringDelimiter)
	} else if char != stringDelimiter {
		return "", errors.Errorf("expected '%c'", stringDelimiter)
	}
	var value strings.Builder
	for {
		char, err := reader.ReadByte()
		if err != nil {
			return "", errors.Wrap(err, "reading char")
		}
		if char == stringEscape {
			if char, err = reader.ReadByte(); err != nil {
				return "", errors.Wrap(err, "reading backslash-escaped char")
			} else if char != stringDelimiter && char != stringEscape {
				return "", errors.Errorf("unexpected backslash-escaped char %c", char)
			} else {
				value.WriteByte(char) // "The returned error is always nil." https://golang.org/pkg/strings/#Builder.WriteByte
			}
		} else if char == stringDelimiter {
			break
		} else if invalidStringChar(char) {
			return "", errors.Errorf("invalid char %d", char)
		} else {
			value.WriteByte(char)
		}
	}
	return value.String(), nil
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
	if char, err := reader.ReadByte(); err != nil {
		return "", errors.Wrap(err, "reading byte")
	} else if !isLCAlpha(char) {
		return "", errors.New("expected lowercase alpha")
	} else {
		output.WriteByte(char)
	}
	for reader.Len() > 0 {
		if char, err := reader.ReadByte(); err != nil {
			return "", errors.Wrap(err, "reading byte")
		} else if !isSecondIdentifierChar(char) {
			// Return to the position before the non-identifier char.
			if err := reader.UnreadByte(); err != nil {
				return "", errors.Wrap(err, "unreading byte")
			}
			return output.String(), nil
		} else {
			output.WriteByte(char)
		}
	}
	return output.String(), nil
}

// https://github.com/ampproject/amphtml/blob/master/spec/amp-cache-transform.md#version-negotation
var vRangeRE = regexp.MustCompile(`[ \t]*\.\.[ \t]*`)

// https://github.com/ampproject/amphtml/blob/master/spec/amp-cache-transform.md#version-negotation
func parseVersions(vSpec string) ([]*rpb.VersionRange, error) {
	vRanges := util.Comma.Split(vSpec, -1)
	ret := []*rpb.VersionRange{}
	for _, vRange := range vRanges {
		bounds := vRangeRE.Split(vRange, -1)
		switch len(bounds) {
		case 1:
			version, err := strconv.ParseInt(bounds[0], 10, 64)
			if err != nil {
				return nil, errors.Errorf("parsing v_range %q", vRange)
			}
			ret = append(ret, &rpb.VersionRange{Min: version, Max: version})
		case 2:
			min, err := strconv.ParseInt(bounds[0], 10, 64)
			if err != nil {
				return nil, errors.Errorf("parsing v_range %q", vRange)
			}
			max, err := strconv.ParseInt(bounds[1], 10, 64)
			if err != nil {
				return nil, errors.Errorf("parsing v_range %q", vRange)
			}
			ret = append(ret, &rpb.VersionRange{Min: min, Max: max})
		default:
			return nil, errors.Errorf("parsing v_range %q", vRange)
		}
	}
	// Sort descending.
	sort.SliceStable(ret, func(i, j int) bool {
		return ret[i].Max > ret[j].Max
	})
	return ret, nil
}

// The valid set of "destination AMP cache" identifiers for which this packager
// can serve a request. Eventually, this should be parsed from caches.json
// (issue #156).
var validIdentifiers = map[string]bool {"any": true, "google": true}

const versionParamName = "v"

// If the given AMP-Cache-Transform request header value is one the packager
// can satisfy, returns the corresponding AMP-Cache-Transform response header
// it should send, plus the transform version it should use. Else, returns
// empty string.
func ShouldSendSXG(header_value string) (string, int64) {
	reader := strings.NewReader(header_value)
	identifiers, err := parseParameterisedList(reader)
	if err != nil {
		log.Printf("Failed to parse AMP-Cache-Transform %q with error %v\n", header_value, err)
		return "", 0
	}

IdentifierLoop:
	for _, identifier := range identifiers {
		if _, ok := validIdentifiers[identifier.id]; ok {
			var requested []*rpb.VersionRange
			for name, value := range identifier.params {
				if name == versionParamName {
					requested, err = parseVersions(value)
					if err != nil {
						log.Printf("Failed to parse versions from %q with error %v\n", header_value, err)
						continue IdentifierLoop
					}
				} else {
					log.Printf("Invalid param name %q in %q\n", name, header_value)
					continue IdentifierLoop
				}
			}
			version, err := transformer.SelectVersion(requested)
			if err != nil {
				log.Printf("Failed to select version from %q with error %v\n", header_value, err)
				continue
			}
			return fmt.Sprintf(`%s;v="%d"`, identifier.id, version), version
		}
	}
	return "", 0
}
