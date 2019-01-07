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
	"strings"
	"testing"

	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html"
	amppb "github.com/ampproject/amphtml/validator"
)

func TestEveryEnumHandled(t *testing.T) {
	for key := range amppb.AmpLayout_Layout_name {
		layout := amppb.AmpLayout_Layout(key)
		if _, ok := layoutMetadataMap[layout]; !ok {
			t.Errorf("layoutMetadataMap is missing enum %v", layout)
		}
	}
}

func TestGetLayoutName(t *testing.T) {
	testCases := []struct {
		layout   amppb.AmpLayout_Layout
		expected string
	}{
		{amppb.AmpLayout_UNKNOWN, ""},
		{amppb.AmpLayout_NODISPLAY, "nodisplay"},
		{amppb.AmpLayout_FIXED, "fixed"},
		{amppb.AmpLayout_FIXED_HEIGHT, "fixed-height"},
		{amppb.AmpLayout_RESPONSIVE, "responsive"},
		{amppb.AmpLayout_CONTAINER, "container"},
		{amppb.AmpLayout_FILL, "fill"},
		{amppb.AmpLayout_FLEX_ITEM, "flex-item"},
		{amppb.AmpLayout_FLUID, "fluid"},
		{amppb.AmpLayout_INTRINSIC, "intrinsic"},
	}
	for _, tc := range testCases {
		actual := getLayoutName(tc.layout)
		if actual != tc.expected {
			t.Errorf("getLayoutName(%s) = %s, want %s", tc.layout, actual, tc.expected)
		}
	}
}

func TestGetLayoutClass(t *testing.T) {
	testCases := []struct {
		layout   amppb.AmpLayout_Layout
		expected string
	}{
		{amppb.AmpLayout_UNKNOWN, ""},
		{amppb.AmpLayout_NODISPLAY, "i-amphtml-layout-nodisplay"},
		{amppb.AmpLayout_FIXED, "i-amphtml-layout-fixed"},
		{amppb.AmpLayout_FIXED_HEIGHT, "i-amphtml-layout-fixed-height"},
		{amppb.AmpLayout_RESPONSIVE, "i-amphtml-layout-responsive"},
		{amppb.AmpLayout_CONTAINER, "i-amphtml-layout-container"},
		{amppb.AmpLayout_FILL, "i-amphtml-layout-fill"},
		{amppb.AmpLayout_FLEX_ITEM, "i-amphtml-layout-flex-item"},
		{amppb.AmpLayout_FLUID, "i-amphtml-layout-fluid"},
		{amppb.AmpLayout_INTRINSIC, "i-amphtml-layout-intrinsic"},
	}
	for _, tc := range testCases {
		actual := getLayoutClass(tc.layout)
		if actual != tc.expected {
			t.Errorf("getLayoutClass(%s) = %s, want %s", tc.layout, actual, tc.expected)
		}
	}
}

func TestApplyLayout(t *testing.T) {
	testCases := []struct {
		desc     string
		node     *html.Node
		expected string
	}{
		{
			"Appends class properly",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "layout", Val: "nodisplay"}, html.Attribute{Key: "class", Val: "myclass"}),
			`<amp-img layout="nodisplay" class="myclass i-amphtml-layout-nodisplay" hidden="hidden" i-amphtml-layout="nodisplay"></amp-img>`,
		},
		{
			"Nodisplay",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "layout", Val: "nodisplay"}),
			`<amp-img layout="nodisplay" class="i-amphtml-layout-nodisplay" hidden="hidden" i-amphtml-layout="nodisplay"></amp-img>`,
		},
		{
			"Fixed",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "height", Val: "100"}, html.Attribute{Key: "layout", Val: "fixed"}, html.Attribute{Key: "width", Val: "300"}),
			`<amp-img height="100" layout="fixed" width="300" class="i-amphtml-layout-fixed i-amphtml-layout-size-defined" style="width:300px;height:100px;" i-amphtml-layout="fixed"></amp-img>`,
		},
		{
			"Fixed height",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "height", Val: "100"}, html.Attribute{Key: "layout", Val: "fixed-height"}),
			`<amp-img height="100" layout="fixed-height" class="i-amphtml-layout-fixed-height i-amphtml-layout-size-defined" style="height:100px;" i-amphtml-layout="fixed-height"></amp-img>`,
		},
		{
			"Responsive",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "height", Val: "100"}, html.Attribute{Key: "layout", Val: "responsive"}, html.Attribute{Key: "width", Val: "300"}),
			`<amp-img height="100" layout="responsive" width="300" class="i-amphtml-layout-responsive i-amphtml-layout-size-defined" i-amphtml-layout="responsive"><i-amphtml-sizer style="display:block;padding-top:33.3333%;"></i-amphtml-sizer></amp-img>`,
		},
		{
			"Fill",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "layout", Val: "fill"}),
			`<amp-img layout="fill" class="i-amphtml-layout-fill i-amphtml-layout-size-defined" i-amphtml-layout="fill"></amp-img>`,
		},
		{
			"Container",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "layout", Val: "container"}),
			`<amp-img layout="container" class="i-amphtml-layout-container" i-amphtml-layout="container"></amp-img>`,
		},
		{
			"Flex item with width",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "layout", Val: "flex-item"}, html.Attribute{Key: "width", Val: "300"}),
			`<amp-img layout="flex-item" width="300" class="i-amphtml-layout-flex-item i-amphtml-layout-size-defined" style="width:300px;" i-amphtml-layout="flex-item"></amp-img>`,
		},
		{
			"Flex item with height",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "layout", Val: "flex-item"}, html.Attribute{Key: "height", Val: "100"}),
			`<amp-img layout="flex-item" height="100" class="i-amphtml-layout-flex-item i-amphtml-layout-size-defined" style="height:100px;" i-amphtml-layout="flex-item"></amp-img>`,
		},
		{
			"Flex item with width and height",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "layout", Val: "flex-item"}, html.Attribute{Key: "height", Val: "100"}, html.Attribute{Key: "width", Val: "300"}),
			`<amp-img layout="flex-item" height="100" width="300" class="i-amphtml-layout-flex-item i-amphtml-layout-size-defined" style="width:300px;height:100px;" i-amphtml-layout="flex-item"></amp-img>`,
		},
		{
			"No changes to amp-audio",
			htmlnode.Element("amp-audio"),
			"<amp-audio></amp-audio>",
		},
		{
			"Non amp custom tag skipped",
			htmlnode.Element("img"),
			"<img/>",
		},
		{
			"Style attributes preserved and added",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "height", Val: "400"}, html.Attribute{Key: "layout", Val: "fixed"}, html.Attribute{Key: "style", Val: "display:none;height:200px;position:relative;width:150px"}, html.Attribute{Key: "width", Val: "300"}),
			`<amp-img height="400" layout="fixed" style="display:none;height:200px;position:relative;width:150px;width:300px;height:400px;" width="300" class="i-amphtml-layout-fixed i-amphtml-layout-size-defined" i-amphtml-layout="fixed"></amp-img>`,
		},
	}
	for _, tc := range testCases {
		if err := ApplyLayout(tc.node); err != nil {
			t.Errorf("%s: ApplyLayout got unexpected error %v", tc.desc, err)
			continue
		}
		var output strings.Builder
		if err := html.Render(&output, tc.node); err != nil {
			t.Errorf("%s: html.Render got unexepcted error %v", tc.desc, err)
			continue
		}
		if output.String() != tc.expected {
			t.Errorf("%s: ApplyLayout =\n%s, want\n%s", tc.desc, output.String(), tc.expected)
		}
	}
}

func TestTemplateAncestorIgnored(t *testing.T) {
	template := htmlnode.Element("template")
	ampImg := htmlnode.Element("amp-img", html.Attribute{Key: "layout", Val: "nodisplay"})
	template.AppendChild(ampImg)

	expected := `<amp-img layout="nodisplay"></amp-img>`

	if err := ApplyLayout(ampImg); err != nil {
		t.Fatalf("ApplyLayout got unexpected error %v", err)
	}
	var output strings.Builder
	if err := html.Render(&output, ampImg); err != nil {
		t.Fatalf("html.Render got unexepcted error %v", err)
	}
	if output.String() != expected {
		t.Errorf("ApplyLayout =\n%s, want\n%s", output.String(), expected)
	}
}

func TestUnsupported(t *testing.T) {
	testCases := []struct {
		desc string
		node *html.Node
	}{
		{
			"Fluid",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "layout", Val: "fluid"}),
		},
		{
			"Intrinsic",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "layout", Val: "intrinsic"}),
		},
		{
			"fluid width",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "layout", Val: "fixed-height"}, html.Attribute{Key: "width", Val: "fluid"}),
		},
		{
			"fluid height, but not fluid layout",
			htmlnode.Element(
				"amp-img",
				html.Attribute{Key: "layout", Val: "fixed-height"}, html.Attribute{Key: "height", Val: "fluid"}),
		},
	}
	for _, tc := range testCases {
		if err := ApplyLayout(tc.node); err == nil {
			t.Errorf("%s: ApplyLayout unexpectedly succeeded", tc.desc)
		}
	}
}
