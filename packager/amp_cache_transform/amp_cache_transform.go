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

	"github.com/ampproject/amppackager/transformer"
	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/pkg/errors"
)

// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-3.3
type parameterisedIdentifier struct {
	id string
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
		comma, err := reader.ReadByte()
		if err != nil {
			return nil, errors.Wrap(err, "reading ','")
		}
		if comma != ',' {
			return nil, errors.New("expected ','")
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
		semicolon, err := reader.ReadByte()
		if err != nil {
			return nil, errors.Wrap(err, "reading ';'")
		}
		if semicolon != ';' {
			if err := reader.UnreadByte(); err != nil {
				return nil, errors.Wrap(err, "unreading ';'")
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
		if name != "v" {
			return nil, errors.Errorf("invalid AMP-Cache-Transform param %q", name)
		}

		// ... and a parameter value.
		// NOTE: The current version of the AMP-Cache-Transform spec
		// does not use any null-valued parameters. So, for now, a
		// missing '=' is a parse failure.
		equals, err := reader.ReadByte()
		if err != nil {
			return nil, errors.Wrap(err, "reading '='")
		}
		if equals != '=' {
			return nil, errors.New("expected '='")
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

// https://tools.ietf.org/html/draft-ietf-httpbis-header-structure-07#section-4.2.7
func parseString(reader *strings.Reader) (string, error) {
	quote, err := reader.ReadByte()
	if err != nil {
		return "", errors.Wrap(err, "reading '\"'")
	}
	if quote != '"' {
		return "", errors.New("expected '\"'")
	}
	var value strings.Builder
	for {
		char, err := reader.ReadByte()
		if err != nil {
			return "", errors.Wrap(err, "reading char")
		}
		if char <= 0x1f || char == 0x7f {
			return "", errors.Errorf("invalid char %d", char)
		}
		if char == '"' {
			break
		}
		if char == '\\' {
			char, err = reader.ReadByte()
			if err != nil {
				return "", errors.Wrap(err, "reading backslash-escaped char")
			}
			if char != '"' && char != '\\' {
				return "", errors.Errorf("unexpected backslash-escaped char %c", char)
			}
		}
		value.WriteByte(char)  // "The returned error is always nil." https://golang.org/pkg/strings/#Builder.WriteByte
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
			if err := reader.UnreadByte(); err != nil {
				return "", errors.Wrap(err, "unreading byte")
			}
			return output.String(), nil
		}
		output.WriteByte(char)
	}
	return output.String(), nil
}

// https://github.com/ampproject/amphtml/blob/master/spec/amp-cache-transform.md#version-negotation
var vRangeRE = regexp.MustCompile(`[ \t]*\.\.[ \t]*`)

// https://github.com/ampproject/amphtml/blob/master/spec/amp-cache-transform.md#version-negotation
func parseVersions(v_spec string) ([]*rpb.VersionRange, error) {
	vRanges := strings.FieldsFunc(v_spec, func(c rune) bool {
		return c == ',' || c == ' ' || c == '\t'
	})
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
	for _, identifier := range identifiers {
		if identifier.id == "any" || identifier.id == "google" {
			var requested []*rpb.VersionRange
			if v, ok := identifier.params["v"]; ok {
				requested, err = parseVersions(v)
				if err != nil {
					log.Printf("Failed to parse versions from %q with error %v\n", header_value, err)
					return "", 0
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
