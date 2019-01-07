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
	"strconv"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"github.com/pkg/errors"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
	amppb "github.com/ampproject/amphtml/validator"
)

// String constants to avoid magic numbers
const (
	enumSeparator      = "_"
	attributeSeparator = "-"
	// The CSS class used for layouts that have size definitions
	layoutSizeDefinedClass = "i-amphtml-layout-size-defined"
)

// cssDimensions encapsulates both width and height.
type cssDimensions struct {
	width  *cssLength
	height *cssLength
}

// layoutMetadata provides more information about an AmpLayout_Layout enum
type layoutMetadata struct {
	// Whether the layout is supported for server-side rendering
	isSupported bool
	// Whether the layout has size defined
	hasSizeDefinition bool
}

// layoutMetadataMap describes for each AmpLayout enum value, its level of
// support within the AMP layout algorithm.
//
// This map must be kept up-to-date to include every value! There is a
// corresponding test that ensures this.
var /* const */ layoutMetadataMap = map[amppb.AmpLayout_Layout]layoutMetadata{
	amppb.AmpLayout_UNKNOWN:      {false, false},
	amppb.AmpLayout_NODISPLAY:    {true, false},
	amppb.AmpLayout_FIXED:        {true, true},
	amppb.AmpLayout_FIXED_HEIGHT: {true, true},
	amppb.AmpLayout_RESPONSIVE:   {true, true},
	amppb.AmpLayout_CONTAINER:    {true, false},
	amppb.AmpLayout_FILL:         {true, true},
	amppb.AmpLayout_FLEX_ITEM:    {true, true},
	amppb.AmpLayout_FLUID:        {false, true},
	amppb.AmpLayout_INTRINSIC:    {false, true},
}

// ApplyLayout applies the AMP layout algorithm to the given custom element
// (prefixed with "amp-"), by performing the same calculations that
// the AMP runtime would perform, and recording the results in the
// style attribute. Returns an error if the layout isn't supported.
//
// <amp-audio> is an exception because it requires knowing the
// dimensions of the browser. Therefore no layout is applied to it.
// Also, any descendants of <template> are ignored.
func ApplyLayout(n *html.Node) error {
	if !amphtml.IsAMPCustomElement(n) || n.Data == amphtml.AMPAudio || htmlnode.IsDescendantOf(n, atom.Template) {
		return nil
	}

	inputLayout := parseAMPLayout(n)
	dimensions, err := getNormalizedDimensions(n, inputLayout)
	if err != nil {
		return err
	}
	actualLayout, err := getNormalizedLayout(
		inputLayout, dimensions,
		htmlnode.GetAttributeValOrNil(n, "", "sizes"),
		htmlnode.GetAttributeValOrNil(n, "", "heights"))
	if err != nil {
		return err
	}
	apply(n, actualLayout, dimensions)
	return nil
}

// Parses the layout attribute value of the given node and returns the
// corresponding AmpLayout_Layout enum.
func parseAMPLayout(n *html.Node) amppb.AmpLayout_Layout {
	v, ok := htmlnode.GetAttributeVal(n, "", "layout")
	if !ok || v == "" {
		return amppb.AmpLayout_UNKNOWN
	}
	key := strings.Replace(strings.ToUpper(v), attributeSeparator, enumSeparator, -1)
	if val, ok := amppb.AmpLayout_Layout_value[key]; ok {
		return amppb.AmpLayout_Layout(val)
	}
	return amppb.AmpLayout_UNKNOWN
}

// getLayoutName returns the name of the AmpLayout_Layout enum suitable
// for AMP HTML (i.e. lowercased, underscore instead of hyphens).
func getLayoutName(layout amppb.AmpLayout_Layout) string {
	if layout == amppb.AmpLayout_UNKNOWN {
		return ""
	}
	return strings.Replace(strings.ToLower(layout.String()), enumSeparator, attributeSeparator, -1)
}

// getLayoutClass returns the CSS class for the AmpLayout_Layout enum
func getLayoutClass(layout amppb.AmpLayout_Layout) string {
	s := getLayoutName(layout)
	if s == "" {
		return s
	}
	return amphtml.IAMPHTMLLayout + "-" + s
}

// getNormalizedDimensions returns the normalized CSS dimensions
// (width / height) for the element.
func getNormalizedDimensions(n *html.Node, layout amppb.AmpLayout_Layout) (cssDimensions, error) {
	var result cssDimensions
	var err error
	result.width, err = getNormalizedWidth(n, layout)
	if err != nil {
		return result, err
	}

	result.height, err = getNormalizedHeight(n, layout)
	return result, err
}

// getNormalizedLayout returns the normalized AmpLayout based on the
// provided dimensions, or an error if it is not supported.
func getNormalizedLayout(layout amppb.AmpLayout_Layout, dimensions cssDimensions, sizes, heights *string) (amppb.AmpLayout_Layout, error) {
	var result amppb.AmpLayout_Layout
	if layout != amppb.AmpLayout_UNKNOWN {
		result = layout
	} else if !dimensions.width.isSet && !dimensions.height.isSet {
		result = amppb.AmpLayout_CONTAINER
	} else if (dimensions.height.isSet && dimensions.height.isFluid) || (dimensions.width.isSet && dimensions.width.isFluid) {
		result = amppb.AmpLayout_FLUID
	} else if dimensions.height.isSet && (!dimensions.width.isSet || dimensions.width.isAuto) {
		result = amppb.AmpLayout_FIXED_HEIGHT
	} else if dimensions.height.isSet && dimensions.width.isSet && (sizes != nil || heights != nil) {
		result = amppb.AmpLayout_RESPONSIVE
	} else {
		result = amppb.AmpLayout_FIXED
	}

	meta, ok := layoutMetadataMap[result]
	if !ok || !meta.isSupported {
		return result, errors.Errorf("layout %q is not supported", result)
	}
	return result, nil
}

// apply modifies the DOM node, adding appropriate styles for the
// layout and dimensions.
func apply(n *html.Node, layout amppb.AmpLayout_Layout, dimensions cssDimensions) {
	class := getLayoutClass(layout)
	meta, ok := layoutMetadataMap[layout]
	if ok && meta.hasSizeDefinition {
		class = class + " " + layoutSizeDefinedClass
	}
	htmlnode.AppendAttributeWithSeparator(n, "", "class", class, " ")

	var styles string
	switch layout {
	case amppb.AmpLayout_NODISPLAY:
		htmlnode.SetAttribute(n, "", "hidden", "hidden")
	case amppb.AmpLayout_FIXED, amppb.AmpLayout_FLEX_ITEM:
		styles = getCSSLengthStyle(dimensions.width, "width") +
			getCSSLengthStyle(dimensions.height, "height")
	case amppb.AmpLayout_FIXED_HEIGHT:
		styles = getCSSLengthStyle(dimensions.height, "height")
	case amppb.AmpLayout_RESPONSIVE:
		// Do nothing here but emit <i-amphtml-sizer> later.
	case amppb.AmpLayout_FILL, amppb.AmpLayout_CONTAINER:
		// Do nothing
	}
	// Appends the given styles so they take precedence over any
	// user-defined styles (!important is reserved for AMP Runtime). Currently
	// the only potential properties added are `display`, `height`, and `width`.
	if styles != "" {
		htmlnode.AppendAttributeWithSeparator(n, "", "style", styles, ";")
	}

	if a, ok := htmlnode.FindAttribute(n, "", "style"); ok && a.Val == "" {
		// Remove empty style attribute
		htmlnode.RemoveAttribute(n, a)
	}
	htmlnode.SetAttribute(n, "", amphtml.IAMPHTMLLayout, getLayoutName(layout))

	// Add sizer info if necessary
	if layout != amppb.AmpLayout_RESPONSIVE ||
		!dimensions.width.isSet ||
		dimensions.width.value == 0 ||
		!dimensions.height.isSet ||
		dimensions.height.unit != dimensions.width.unit {
		return
	}
	percent := dimensions.height.value / dimensions.width.value * 100
	padding := strconv.FormatFloat(percent, 'f', 4, 64)
	sizerStyle := "display:block;padding-top:" + padding + "%;"
	sizer := htmlnode.Element(
		"i-amphtml-sizer", html.Attribute{Key: "style", Val: sizerStyle})
	n.InsertBefore(sizer, n.FirstChild)
}
