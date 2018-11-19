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

package htmlnode

import (
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

const htmltext = `
<html>
<head>
<body>
0<1
<p id="A" foo="abc&#34;def">
"2"
<b empty="">
3
<i backslash="\">
&amp;4
5
<blockquote>
<br>
6
<p id="B" foo="xyz">
</body>
</html>
`

// Tests a bunch of different methods on the test HTML.
func TestHTMLNode(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(htmltext))
	if err != nil {
		t.Fatalf("Error parsing htmltext: %q", err)
	}

	// Find the first <p>
	pNode, ok := FindNode(doc, atom.P)
	if !ok {
		t.Fatalf("Expected to find <p> tag.")
	}

	idAttr, ok := FindAttribute(pNode, "", "id")
	if !ok {
		t.Fatalf("Expected to find id attribute")
	}
	if idAttr.Val != "A" {
		t.Fatalf("id=%s, want=A", idAttr.Val)
	}

	if !IsDescendantOf(pNode, atom.Body) {
		t.Fatalf("Expected <p> to be descendant of <body>")
	}
	if !IsChildOf(pNode, atom.Body) {
		t.Fatalf("Expected <p> to be child of <body>")
	}
	if !IsDescendantOf(pNode, atom.Html) {
		t.Fatalf("Expected <p> to be descendant of <html>")
	}
	if IsChildOf(pNode, atom.Html) {
		t.Fatalf("<p> should not be child of <html>")
	}

	// make sure bar attr doesn't exist on pNode
	if _, ok := FindAttribute(pNode, "", "bar"); ok {
		t.Fatalf("Unexpectedly found bar attr")
	}
	SetAttribute(pNode, "", "bar", "martini")
	// now should exist
	barAttr, ok := FindAttribute(pNode, "", "bar")
	if !ok {
		t.Fatalf("Expected fo find bar attr.")
	}
	if got, want := barAttr.Val, "martini"; got != want {
		t.Fatalf("barAttr.Val=%s, want=%s", got, want)
	}

	AppendAttribute(pNode, "", "bar", " with olives")
	if got, want := barAttr.Val, "martini with olives"; got != want {
		t.Fatalf("barAttr.Val=%s, want=%s", got, want)
	}

	AppendAttributeWithSeparator(pNode, "", "bar", "and a twist", " ")
	if got, want := barAttr.Val, "martini with olives and a twist"; got != want {
		t.Fatalf("barAttr.Val=%s, want=%s", got, want)
	}

	PrependAttribute(pNode, "", "bar", "dirty ")
	if got, want := barAttr.Val, "dirty martini with olives and a twist"; got != want {
		t.Fatalf("barAttr.Val=%s, want=%s", got, want)
	}

	RemoveAttribute(pNode, barAttr)
	// make sure bar attr doesn't exist on pNode
	if _, ok := FindAttribute(pNode, "", "bar"); ok {
		t.Fatalf("Unexpectedly found bar attr.")
	}
}

func TestElement(t *testing.T) {
	testCases := []struct {
		tag      string
		expected atom.Atom
	}{
		{"style", atom.Style},
		{"Style", atom.Style},
		{"amp-img", 0},
	}
	for _, tc := range testCases {
		e := Element(tc.tag)
		lowerTag := strings.ToLower(tc.tag)
		if e.Data != lowerTag {
			t.Errorf("Element(%s).Data=%s, want %s", tc.tag, e.Data, lowerTag)
		}
		if e.DataAtom != tc.expected {
			t.Errorf("Element(%s).DataAtom=%s, want %s", tc.tag, e.DataAtom, tc.expected)
		}
	}
}

func TestGetAttributeValOrNil(t *testing.T) {
	testCases := []struct {
		node      *html.Node
		namespace string
		expected  *string
	}{
		{Element("p", html.Attribute{Namespace: "foo", Key: "id", Val: "A"}), "foo", proto.String("A")},
		{Element("p", html.Attribute{Namespace: "foo", Key: "id", Val: "A"}), "differentnamespace", nil},
		{Element("p", html.Attribute{Key: "id", Val: "A"}), "", proto.String("A")},
		{Element("p"), "", nil},
	}
	for _, tc := range testCases {
		actual := GetAttributeValOrNil(tc.node, tc.namespace, "id")
		if tc.expected == nil {
			if actual != nil {
				t.Errorf("GetAttributeValOrNil(%v, \"id\") = %s, want nil", tc.node, *actual)
			}
		} else if *actual != *tc.expected {
			t.Errorf("GetAttributeValOrNil(%v, \"id\") = %s, want %s", tc.node, *actual, *tc.expected)
		}
	}
}

func TestNextAndPrev(t *testing.T) {
	// This creates a tree of
	// root
	// |- child1
	//    |- grandchild1
	// |- child2
	root := html.Node{Data: "root"}
	child1 := html.Node{Data: "child1"}
	child2 := html.Node{Data: "child2"}
	grandchild1 := html.Node{Data: "grandchild1"}
	root.AppendChild(&child1)
	root.AppendChild(&child2)
	child1.AppendChild(&grandchild1)

	// expected traversal order and node value
	expected := map[int]string{
		0: "root",
		1: "child1",
		2: "grandchild1",
		3: "child2",
	}

	// Next
	index := 0
	for n := &root; n != nil; n = Next(n) {
		if n.Data != expected[index] {
			t.Errorf("Next(%v, \"index\") = %d, want %s", n, index, expected[index])
		}
		index++
	}

	// Prev
	index = 3
	for n := &child2; n != nil; n = Prev(n) {
		if n.Data != expected[index] {
			t.Errorf("Prev(%v, \"index\") = %d, want %s", n, index, expected[index])
		}
		index--
	}

	// RemoveNode
	nodePtr := &child1
	removed := RemoveNode(&nodePtr)
	if nodePtr.Data != "root" {
		t.Errorf("RemoveNode(%v) = %s, want %s", child1, removed.Data, "root")
	}
	if n := Next(nodePtr); n.Data != "child2" {
		t.Errorf("Next after RemovedNode = %s, want %s", n.Data, "child2")
	}

}
