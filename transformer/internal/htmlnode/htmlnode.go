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

	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

// RemoveNode removes the node and adjusts the Node pointer to continue
// iterating over the remaining tree nodes. Returns the removed node.
func RemoveNode(n **html.Node) *html.Node {
	current := *n
	// Update pointer to previous.
	*n = Prev(*n)
	current.Parent.RemoveChild(current)
	return current
}

// Prev returns the prev node in depth first order.
func Prev(n *html.Node) *html.Node {
	c := n.PrevSibling
	if c == nil {
		return n.Parent
	}
	for c.LastChild != nil {
		c = c.LastChild
	}
	return c
}

// Next returns the next node in depth first order.
func Next(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	if n.FirstChild != nil {
		return n.FirstChild
	}

	c := n.NextSibling
	p := n.Parent
	for c == nil && p != nil {
		c = p.NextSibling
		p = p.Parent
	}
	return c
}

// FindNode returns the (first) specified child node of the given atom
// type or ok=false if there are none.
func FindNode(n *html.Node, atom atom.Atom) (*html.Node, bool) {
	if n.DataAtom == atom {
		return n, true
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := FindNode(c, atom)
		if ok {
			return result, ok
		}
	}
	return nil, false
}

// IsDescendantOf returns true if the node is a descendant of the
// given atom, regardless of distance (e.g. the node can be child, or
// grandchild, or great grandchild, etc.).
func IsDescendantOf(n *html.Node, atom atom.Atom) bool {
	for p := n.Parent; p != nil; p = p.Parent {
		if p.DataAtom == atom {
			return true
		}
	}
	return false
}

// IsChildOf returns true if the node is a direct descendant
// (immediate child) of the given atom.
func IsChildOf(n *html.Node, atom atom.Atom) bool {
	return n.Parent.DataAtom == atom
}

// FindAttribute returns a pointer to the attribute of node n with the
// given namespace and key or ok=false if there is none.
func FindAttribute(n *html.Node, namespace, key string) (*html.Attribute, bool) {
	for i := range n.Attr {
		if n.Attr[i].Namespace == namespace && n.Attr[i].Key == key {
			return &n.Attr[i], true
		}
	}
	return nil, false
}

// GetAttributeVal returns the value for the attribute named with
// 'key' or ok=false if the attribute doesn't exist.
func GetAttributeVal(n *html.Node, namespace, key string) (string, bool) {
	if a, ok := FindAttribute(n, namespace, key); ok {
		return a.Val, true
	}
	return "", false
}

// GetAttributeValOrNil returns a pointer to the value for the
// attribute named with 'key' or nil if the attribute doesn't
// exist. There are cases when it is necessary to differentiate
// between nil versus empty (which is imposible from Go primitives).
func GetAttributeValOrNil(n *html.Node, namespace, key string) *string {
	if v, ok := GetAttributeVal(n, namespace, key); ok {
		return &v
	}
	return nil
}

// HasAttribute returns true if the node has the attribute named with 'key'.
func HasAttribute(n *html.Node, namespace, key string) bool {
	_, ok := FindAttribute(n, namespace, key)
	return ok
}

// SetAttribute overrides the value of the attribute on node n with
// namespace and key with val. If the attribute doesn't exist, it adds it.
func SetAttribute(n *html.Node, namespace, key, val string) {
	if a, ok := FindAttribute(n, namespace, key); ok {
		a.Val = val
	} else {
		n.Attr = append(n.Attr, html.Attribute{
			Namespace: namespace,
			Key:       key,
			Val:       val,
		})
	}
}

// AppendAttribute appends to any existing value of the attribute on
// node n with namespace and key with val. If the attribute doesn't
// exist, it adds it.
func AppendAttribute(n *html.Node, namespace, key, val string) {
	addAttributeValue(n, namespace, key, val, false, "")
}

// AppendAttributeWithSeparator appends to any existing value of the
// attribute on node n with namespace and key with val. If the
// attribute doesn't exist, it adds it. If the attribute already exists,
// uses the given separator before appending the new value.
func AppendAttributeWithSeparator(n *html.Node, namespace, key, val, sep string) {
	addAttributeValue(n, namespace, key, val, false, sep)
}

// PrependAttribute prepends to any existing value of the attribute on
// node n with namespace and key with val. If the attribute doesn't
// exist, it adds it.
func PrependAttribute(n *html.Node, namespace, key, val string) {
	addAttributeValue(n, namespace, key, val, true, "")
}

// addAttributeValue appends (or prepends) to any existing value of the
// attribute on node n with namespace and key with val. If the
// attribute doesn't exist, it adds it.
func addAttributeValue(n *html.Node, namespace, key, val string, prepend bool, sep string) {
	if a, ok := FindAttribute(n, namespace, key); ok {
		if len(a.Val) == 0 {
			a.Val = val
		} else if prepend {
			a.Val = val + sep + a.Val
		} else {
			// Only add separator if needed.
			if strings.EqualFold(a.Val[len(a.Val)-1:], sep) {
				a.Val = a.Val + val
			} else {
				a.Val = a.Val + sep + val
			}
		}
	} else {
		n.Attr = append(n.Attr, html.Attribute{
			Namespace: namespace,
			Key:       key,
			Val:       val,
		})
	}
}

// RemoveAttribute removes the given attribute from node n. If it doesn't
// exist, this does nothing.
func RemoveAttribute(n *html.Node, a *html.Attribute) {
	if a == nil {
		return
	}
	for i := len(n.Attr) - 1; i >= 0; i-- {
		// if memory address is same, or if the contents are equal, remove it
		if &n.Attr[i] == a || n.Attr[i] == *a {
			n.Attr = append(n.Attr[:i], n.Attr[i+1:]...)
		}
	}
}

// AppendChildren appends the array of nodes to node n.
func AppendChildren(p *html.Node, c ...*html.Node) {
	for _, n := range c {
		p.AppendChild(n)
	}
}

// RemoveAllChildren removes all children from node n.
func RemoveAllChildren(n *html.Node) {
	for c := n.FirstChild; c != nil; c = n.FirstChild {
		n.RemoveChild(c)
	}
}

// Text returns an html.Node containing the given string.
func Text(t string) *html.Node {
	return &html.Node{
		Type: html.TextNode,
		Data: t,
	}
}

// Element returns an html.Node for an element with tag and attributes
func Element(tag string, attrs ...html.Attribute) *html.Node {
	s := strings.ToLower(tag)
	return &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Lookup([]byte(s)),
		Data:     s,
		Attr:     attrs,
	}
}
