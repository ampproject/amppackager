// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package layout

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
	amppb "github.com/ampproject/amphtml/validator"
)

// cssLength encapsulates a CSS length. See
// https://developer.mozilla.org/en-US/docs/Web/CSS/length .
type cssLength struct {
	// Whether the attribute value is set
	isSet bool
	// Whether the attribute value is 'auto'. This is a special value
	// that indicates that the value gets derived from the context. In
	// practice that is only ever the case for a width.
	isAuto bool
	// Whether the attribute value is 'fluid'.
	isFluid bool
	// The numeric value.
	value float64
	// The unit, e.g. px
	unit string
}

// Constant lengths used for known natural dimensions.
// See https://github.com/ampproject/amphtml/blob/master/src/layout.js#L73
var /* const */ (
	onePx       = cssLength{true, false, false, 1, "px"}
	fortyFourPx = cssLength{true, false, false, 44, "px"}
	sixtyPx     = cssLength{true, false, false, 60, "px"}
)

var /* const */ (
	// The accepted css units
	acceptedUnits = map[string]struct{}{
		"em": {}, "px": {}, "rem": {}, "vh": {}, "vw": {}, "vmin": {}, "vmax": {}}

	// regex matches a css length comprising of a number and optional
	// unit, e.g. 12px
	regex = regexp.MustCompile("^((?:\\d*\\.)?\\d+)(\\D*)$")
)

// The default unit, when none are specified.
const defaultunit = "px"

// newCSSLength constructs and returns a pointer to a cssLength struct
// by parsing the 'input' value. 'allow_auto' determines whether
// 'auto' is accepted as a value. 'allow_fluid' determines whether
// 'fluid' is accepted as a value. The default unit is 'px' (pixels),
// if not specified. An error is returned if there was a problem
// parsing.
func newCSSLength(input *string, allowAuto, allowFluid bool) (*cssLength, error) {
	result := cssLength{unit: defaultunit}
	// Nil is considered valid.
	if input == nil {
		return &result, nil
	}
	result.isSet = true
	switch *input {
	case "auto":
		result.isAuto = true
		if !allowAuto {
			return nil, errors.New("autoAllow is false but input is auto")
		}
	case "fluid":
		result.isFluid = true
		if !allowFluid {
			return nil, errors.New("autoFluid is false but input is fluid")
		}
	default:
		m := regex.FindStringSubmatch(*input)
		if m == nil || len(m) != 3 {
			return nil, errors.Errorf("invalid input: %s", *input)
		}
		// m will have length of 3, holding the text of the leftmost
		// match of the regex, and the subexpression matches (which
		// may be empty string)
		var err error
		result.value, err = strconv.ParseFloat(m[1], 64)
		if err != nil {
			return nil, err
		}
		if m[2] != "" {
			if _, ok := acceptedUnits[m[2]]; !ok {
				return &result, errors.Errorf("Unit is not valid: %s", m[2])
			}
			result.unit = m[2]
		}
	}
	return &result, nil
}

// getNormalizedWidth returns the normalized cssLength representing
// the node's width, or an error if there was a problem
// parsing. Normalization takes into account that some elements, such
// as amp-analytics and amp-pixel, have natural dimensions (browser or
// implemention-specific defaults for width/height).
func getNormalizedWidth(n *html.Node, layout amppb.AmpLayout_Layout) (*cssLength, error) {
	width, err := newCSSLength(htmlnode.GetAttributeValOrNil(n, "", "width"), true /*allow_auto*/, false /*allow_fluid*/)
	if err != nil {
		return nil, err
	}
	if width.isSet {
		return width, nil
	}
	if layout == amppb.AmpLayout_UNKNOWN || layout == amppb.AmpLayout_FIXED {
		switch strings.ToUpper(n.Data) {
		case "AMP-ANALYTICS", "AMP-PIXEL":
			// copy the contents, so the constant isn't opened up for mutation risk
			*width = onePx
		case "AMP-SOCIAL-SHARE":
			// copy the contents, so the constant isn't opened up for mutation risk
			*width = sixtyPx
		}
	}
	return width, nil
}

// getNormalizedHeight returns the normalized cssLength representing
// the node's height, or an error if there was a problem
// parsing. Normalization takes into account that some elements, such
// as amp-analytics or amp-pixel, have natural dimensions (browser or
// implemention-specific defaults for width/height).
func getNormalizedHeight(n *html.Node, layout amppb.AmpLayout_Layout) (*cssLength, error) {
	height, err := newCSSLength(htmlnode.GetAttributeValOrNil(n, "", "height"), true /*allow_auto*/, layout == amppb.AmpLayout_FLUID /*allow_fluid*/)
	if err != nil {
		return nil, err
	}
	if height.isSet {
		return height, nil
	}
	if layout == amppb.AmpLayout_UNKNOWN ||
		layout == amppb.AmpLayout_FIXED ||
		layout == amppb.AmpLayout_FIXED_HEIGHT {
		switch strings.ToUpper(n.Data) {
		case "AMP-ANALYTICS", "AMP-PIXEL":
			// copy the contents, so the constant isn't opened up for mutation risk
			*height = onePx
		case "AMP-SOCIAL-SHARE":
			// copy the contents, so the constant isn't opened up for mutation risk
			*height = fortyFourPx
		}
	}
	return height, nil
}

// getCSSLengthStyle returns the CSS style string for the given
// cssLength struct and dimension (width or height)
func getCSSLengthStyle(input *cssLength, dimension string) string {
	if !input.isSet {
		return ""
	}
	if input.isAuto {
		return dimension + ":auto;"
	}
	f := strconv.FormatFloat(input.value, 'g', -1, 64)
	return strings.Join([]string{dimension, ":", f, input.unit, ";"}, "")
}
