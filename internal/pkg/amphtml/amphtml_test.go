package amphtml

import (
	"strings"
	"testing"

	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

func TestIsAMPCustomElement(t *testing.T) {
	tcs := []struct {
		desc     string
		n        *html.Node
		expected bool
	}{
		{
			"amp-img true",
			&html.Node{Type: html.ElementNode, Data: "amp-img"},
			true,
		},
		{
			"img false",
			&html.Node{Type: html.ElementNode, Data: "img", DataAtom: atom.Img},
			false,
		},
		{
			"non ElementNode false",
			&html.Node{Type: html.TextNode},
			false,
		},
	}
	for _, tc := range tcs {
		if ok := IsAMPCustomElement(tc.n); ok != tc.expected {
			t.Errorf("%s: IsAMPCustomElement=%t want=%t", tc.desc, ok, tc.expected)
		}
	}
}

func TestIsScriptAMPRuntime(t *testing.T) {
	tcs := []struct {
		desc     string
		n        *html.Node
		expected bool
	}{
		{
			"amp-img false",
			&html.Node{Type: html.ElementNode, Data: "amp-img"},
			false,
		},
		{
			"amp runtime with custom-element false",
			&html.Node{Type: html.ElementNode, Data: "script", DataAtom: atom.Script, Attr: []html.Attribute{{Key: "async"}, {Key: "custom-element"}, {Key: "src", Val: "https://cdn.ampproject.org/v0.js"}}},
			false,
		},
		{
			"amp runtime true",
			&html.Node{Type: html.ElementNode, Data: "script", DataAtom: atom.Script, Attr: []html.Attribute{{Key: "async"}, {Key: "src", Val: "https://cdn.ampproject.org/v0.js"}}},
			true,
		},
		{
			"amp4ads runtime with custom-element false",
			&html.Node{Type: html.ElementNode, Data: "script", DataAtom: atom.Script, Attr: []html.Attribute{{Key: "async"}, {Key: "custom-element"}, {Key: "src", Val: "https://cdn.ampproject.org/amp4ads-v0.js"}}},
			false,
		},
		{
			"amp4ads runtime true",
			&html.Node{Type: html.ElementNode, Data: "script", DataAtom: atom.Script, Attr: []html.Attribute{{Key: "async"}, {Key: "src", Val: "https://cdn.ampproject.org/amp4ads-v0.js"}}},
			true,
		},
	}
	for _, tc := range tcs {
		if ok := IsScriptAMPRuntime(tc.n); ok != tc.expected {
			t.Errorf("%s: IsScriptAMPRuntime()=%t want=%t", tc.desc, ok, tc.expected)
		}
	}
}

func TestIsScriptRenderDelaying(t *testing.T) {
	tcs := []struct {
		desc     string
		n        *html.Node
		expected bool
	}{
		{
			"amp-img false",
			&html.Node{Type: html.ElementNode, Data: "amp-img"},
			false,
		},
		{
			"amp-dynamic-css-classes true",
			&html.Node{Type: html.ElementNode, Data: "script", DataAtom: atom.Script, Attr: []html.Attribute{{Key: "custom-element", Val: "amp-dynamic-css-classes"}}},
			true,
		},
		{
			"amp-experiment true",
			&html.Node{Type: html.ElementNode, Data: "script", DataAtom: atom.Script, Attr: []html.Attribute{{Key: "custom-element", Val: "amp-experiment"}}},
			true,
		},
		{
			"amp-story true",
			&html.Node{Type: html.ElementNode, Data: "script", DataAtom: atom.Script, Attr: []html.Attribute{{Key: "custom-element", Val: "amp-story"}}},
			true,
		},
	}
	for _, tc := range tcs {
		if ok := IsScriptRenderDelaying(tc.n); ok != tc.expected {
			t.Errorf("%s: IsScriptRenderDelaying()=%t want=%t", tc.desc, ok, tc.expected)
		}
	}
}

func TestNewDOM(t *testing.T) {
	tcs := []struct {
		desc     string
		html     string
		expected bool
	}{
		{
			"true",
			"<html><head></head><body></body></html>",
			true,
		},
		// I can't find a false case. NewDOM might not need to check for required nodes.{
		//	"false",
		//	"<body><head><html>",
		//	false,
		//},
	}
	for _, tc := range tcs {
		n, err := html.Parse(strings.NewReader(tc.html))
		if err != nil {
			t.Errorf("%s: html.Parse(%s) failed unexpectedly. %v", tc.desc, tc.html, err)
		}
		if _, ok := NewDOM(n); ok != tc.expected {
			t.Errorf("%s: NewDOM()=%t want=%t", tc.desc, ok, tc.expected)
		}
	}
}
