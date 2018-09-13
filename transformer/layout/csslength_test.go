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
	"fmt"
	"testing"

	"golang.org/x/net/html"
	amppb "github.com/ampproject/amphtml/validator"
)

func TestValidCSSLength(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected cssLength
	}{
		{"BasicCssLength", "10.1em", cssLength{true, false, false, 10.1, "em"}},
		{"EmptyUnitIsPx", "10", cssLength{true, false, false, 10, "px"}},
	}

	for _, tc := range testCases {
		actual, err := newCSSLength(&tc.input, false, false)
		if err != nil {
			t.Errorf("%s: newCSSLength(%s) got unexpected error: %v", tc.desc, tc.input, err)
			continue
		}
		if *actual != tc.expected {
			t.Errorf("%s: newCSSLength(%s) = %v, want %v", tc.desc, tc.input, actual, tc.expected)
		}
	}
}

func TestInvalidCSSLength(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
	}{
		{"EmptyStringIsInvalid", ""},
		{"Garbage", "100%"},
		{"MoreGarbage", "not a number"},
		{"YetMoreGarbage", "1.1.1"},
		{"InvalidUnit", "5inches"},
		{"MissingNumberWithValidUnitIsGarbage", "px"},
		{"ScreensizeInRomanIsGarbage", "ix unciae"},
		{"BadRegex", "px9px"},
	}

	for _, tc := range testCases {
		actual, err := newCSSLength(&tc.input, false, false)
		if err == nil {
			t.Errorf("%s: newCSSLength(%s) = %v, want error", tc.desc, tc.input, *actual)
		}
	}
}

func TestSupportedUnits(t *testing.T) {
	testCases := []string{
		"px", "em", "rem", "vh", "vmin", "vmax",
	}

	for _, tc := range testCases {
		input := fmt.Sprintf("10%s", tc)
		actual, err := newCSSLength(&input, false, false)
		if err != nil {
			t.Errorf("newCSSLength(%s) got an unexpected error: %q", tc, err)
			continue
		}
		if expected := (cssLength{true, false, false, 10, tc}); *actual != expected {
			t.Errorf("newCSSLength(%s) = %v, want %v", tc, *actual, expected)
		}
	}
}

func TestNullPtrInCssLengthIsValid(t *testing.T) {
	actual, err := newCSSLength(nil, false, false)
	if err != nil {
		t.Errorf("newCSSLength(nil) got an unexpected error: %q", err)
		return
	}
	if expected := (cssLength{false, false, false, 0, "px"}); *actual != expected {
		t.Errorf("newCSSLength(nil) = %v, want %v", *actual, expected)
	}
}

func TestAutoIfAllowAuto(t *testing.T) {
	testCases := []struct {
		desc, input                            string
		allowAuto, expectedValid, expectedAuto bool
	}{
		{"AllowAutoFalse with input != auto", "1", false, true, false},
		{"AllowAutoTrue with input != auto", "1", true, true, false},
		{"AllowAutoFalse with input == auto", "auto", false, false, true},
		{"AllowAutoTrue with input == auto", "auto", true, true, true},
	}

	for _, tc := range testCases {
		actual, err := newCSSLength(&tc.input, tc.allowAuto, false)
		if (err != nil) == tc.expectedValid {
			t.Errorf("%s: newCSSLength = error %v, want valid %t", tc.desc, err, tc.expectedValid)
		}
		if actual != nil && actual.isAuto != tc.expectedAuto {
			t.Errorf("%s: newCSSLength got isAuto = %t, want %t", tc.desc, actual.isAuto, tc.expectedAuto)
		}
	}
}

func TestFluidIfAllowFluid(t *testing.T) {
	testCases := []struct {
		desc, input                              string
		allowFluid, expectedValid, expectedFluid bool
	}{
		{"AllowFluidFalse with input != fluid", "1", false, true, false},
		{"AllowFluidTrue with input != fluid", "1", true, true, false},
		{"AllowFluidFalse with input == fluid", "fluid", false, false, true},
		{"AllowFluidTrue with input == fluid", "fluid", true, true, true},
	}

	for _, tc := range testCases {
		actual, err := newCSSLength(&tc.input, false, tc.allowFluid)
		if (err != nil) == tc.expectedValid {
			t.Errorf("%s newCSSLength = error %v, want valid %t", tc.desc, err, tc.expectedValid)
		}
		if actual != nil && actual.isFluid != tc.expectedFluid {
			t.Errorf("%s new CSSLength got isFluid = %t, want %t", tc.desc, actual.isFluid, tc.expectedFluid)
		}
	}
}

func TestGetNormalizedHeight(t *testing.T) {
	/* const */ heightAttr := []html.Attribute{{Key: "height", Val: "720"}}

	testCases := []struct {
		desc     string
		node     *html.Node
		layout   amppb.AmpLayout_Layout
		expected float64
	}{
		{
			"amp-analytics fixed nil",
			&html.Node{Type: html.ElementNode, Data: "amp-analytics"},
			amppb.AmpLayout_FIXED, 1,
		},
		{
			"amp-analytics fixed_height nil",
			&html.Node{Type: html.ElementNode, Data: "amp-analytics"},
			amppb.AmpLayout_FIXED_HEIGHT, 1,
		},
		{
			"amp-analytics container nil",
			&html.Node{Type: html.ElementNode, Data: "amp-container"},
			amppb.AmpLayout_CONTAINER, 0,
		},
		{
			"amp-pixel fixed nil",
			&html.Node{Type: html.ElementNode, Data: "amp-pixel"},
			amppb.AmpLayout_FIXED, 1,
		},
		{
			"amp-pixel unknown nil",
			&html.Node{Type: html.ElementNode, Data: "amp-pixel"},
			amppb.AmpLayout_UNKNOWN, 1,
		},
		{
			"amp-pixel fixed-height nil",
			&html.Node{Type: html.ElementNode, Data: "amp-pixel"},
			amppb.AmpLayout_FIXED_HEIGHT, 1,
		},
		{
			"amp-pixel container nil",
			&html.Node{Type: html.ElementNode, Data: "amp-pixel"},
			amppb.AmpLayout_CONTAINER, 0,
		},
		{
			"amp-social-share fixed nil",
			&html.Node{Type: html.ElementNode, Data: "amp-social-share"},
			amppb.AmpLayout_FIXED, 44,
		},
		{
			"amp-social-share unknown nil",
			&html.Node{Type: html.ElementNode, Data: "amp-social-share"},
			amppb.AmpLayout_UNKNOWN, 44,
		},
		{
			"amp-social-share fixed-height nil",
			&html.Node{Type: html.ElementNode, Data: "amp-social-share"},
			amppb.AmpLayout_FIXED_HEIGHT, 44,
		},
		{
			"amp-social-share container nil",
			&html.Node{Type: html.ElementNode, Data: "amp-container"},
			amppb.AmpLayout_CONTAINER, 0,
		},
		{
			"amp-img fixed nil",
			&html.Node{Type: html.ElementNode, Data: "amp-img"},
			amppb.AmpLayout_FIXED, 0,
		},
		{
			"amp-img fixed",
			&html.Node{Type: html.ElementNode, Data: "amp-img", Attr: heightAttr},
			amppb.AmpLayout_FIXED, 720,
		},
		{
			"amp-img container",
			&html.Node{Type: html.ElementNode, Data: "amp-img", Attr: heightAttr},
			amppb.AmpLayout_CONTAINER, 720,
		},
	}

	for _, tc := range testCases {
		height, err := getNormalizedHeight(tc.node, tc.layout)
		if err != nil {
			t.Errorf("%s: getNormalizedHeight got unexpected error %v", tc.desc, err)
		}
		if height.value != tc.expected {
			t.Errorf("%s: getNormalizedHeight got value=%f, want %f", tc.desc, height.value, tc.expected)
		}
		if height.unit != "px" {
			t.Errorf("%s: getNormalizedHeight got unit=%s, want px", tc.desc, height.unit)
		}
	}
}

func TestGetNormalizeWidth(t *testing.T) {
	/* const */ widthAttr := []html.Attribute{{Key: "width", Val: "480"}}

	testCases := []struct {
		desc     string
		node     *html.Node
		layout   amppb.AmpLayout_Layout
		expected float64
	}{
		{
			"amp-analytics fixed",
			&html.Node{Type: html.ElementNode, Data: "amp-analytics"},
			amppb.AmpLayout_FIXED, 1,
		},
		{
			"amp-analytics unknown nil",
			&html.Node{Type: html.ElementNode, Data: "amp-analytics"},
			amppb.AmpLayout_UNKNOWN, 1,
		},
		{
			"amp-analytics fixed_height nil",
			&html.Node{Type: html.ElementNode, Data: "amp-analytics"},
			amppb.AmpLayout_FIXED_HEIGHT, 0,
		},
		{
			"amp-analytics container nil",
			&html.Node{Type: html.ElementNode, Data: "amp-analytics"},
			amppb.AmpLayout_CONTAINER, 0,
		},
		{
			"amp-pixel fixed nil",
			&html.Node{Type: html.ElementNode, Data: "amp-pixel"},
			amppb.AmpLayout_FIXED, 1,
		},
		{
			"amp-pixel unknown nil",
			&html.Node{Type: html.ElementNode, Data: "amp-pixel"},
			amppb.AmpLayout_UNKNOWN, 1,
		},
		{
			"amp-pixel fixed-height nil",
			&html.Node{Type: html.ElementNode, Data: "amp-pixel"},
			amppb.AmpLayout_FIXED_HEIGHT, 0,
		},
		{
			"amp-pixel container nil",
			&html.Node{Type: html.ElementNode, Data: "amp-pixel"},
			amppb.AmpLayout_CONTAINER, 0,
		},
		{
			"amp-social-share fixed nil",
			&html.Node{Type: html.ElementNode, Data: "amp-social-share"},
			amppb.AmpLayout_FIXED, 60,
		},
		{
			"amp-social-share unknown nil",
			&html.Node{Type: html.ElementNode, Data: "amp-social-share"},
			amppb.AmpLayout_UNKNOWN, 60,
		},
		{
			"amp-social-share fixed-height nil",
			&html.Node{Type: html.ElementNode, Data: "amp-social-share"},
			amppb.AmpLayout_FIXED_HEIGHT, 0,
		},
		{
			"amp-social-share container",
			&html.Node{Type: html.ElementNode, Data: "amp-social-share"},
			amppb.AmpLayout_CONTAINER, 0,
		},
		{
			"amp-img fixed nil",
			&html.Node{Type: html.ElementNode, Data: "amp-img"},
			amppb.AmpLayout_FIXED, 0,
		},
		{
			"amp-img fixed",
			&html.Node{Type: html.ElementNode, Data: "amp-img", Attr: widthAttr},
			amppb.AmpLayout_FIXED, 480,
		},
		{
			"amp-img container",
			&html.Node{Type: html.ElementNode, Data: "amp-img", Attr: widthAttr},
			amppb.AmpLayout_CONTAINER, 480,
		},
	}

	for _, tc := range testCases {
		width, err := getNormalizedWidth(tc.node, tc.layout)
		if err != nil {
			t.Errorf("%s: getNormalizedWidth got unexpected error %v", tc.desc, err)
		}
		if width.value != tc.expected {
			t.Errorf("%s: getNormalizedWidgth got value=%f, want %f", tc.desc, width.value, tc.expected)
		}
		if width.unit != "px" {
			t.Errorf("%s: getNormalizedWidgth got unit=%s, want px", tc.desc, width.unit)
		}
	}
}

func TestGetCSSLengthStyle(t *testing.T) {
	testCases := []struct {
		input     cssLength
		dimension string
		expected  string
	}{
		{
			cssLength{true, false, false, 30, "px"},
			"height",
			"height:30px;",
		},
		{
			cssLength{true, false, false, 10.1, "em"},
			"width",
			"width:10.1em;",
		},
		{
			cssLength{true, true, false, 0, "px"},
			"width",
			"width:auto;",
		},
		{
			cssLength{true, false, false, 0, "px"},
			"height",
			"height:0px;",
		},
		{
			cssLength{false, false, false, 0, "px"},
			"height",
			"",
		},
	}
	for _, tc := range testCases {
		if actual := getCSSLengthStyle(&tc.input, tc.dimension); actual != tc.expected {
			t.Errorf("getCSSLengthStyle(%v) = %s, want %s", &tc.input, actual, tc.expected)
		}
	}
}
