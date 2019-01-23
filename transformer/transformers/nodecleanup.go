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

package transformers

import (
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

const (
	whitespace = " \t\r\n\f"

	// characters to be stripped out of URIs
	unsanitaryURIChars = "\t\n\r"
)

var (
	scriptJSPReplacer = strings.NewReplacer("<%", "\\u003c%", "%>", "%\\u003e")
	styleJSPReplacer  = strings.NewReplacer("<%", `\3c %`, "%>", `%\3e `)
)

// NodeCleanup cleans up the DOM tree, including, but not limited to:
//  - stripping comment nodes.
//  - stripping noscript elements.
//  - removing duplicate attributes
//  - stripping nonce attributes
//  - Escape JSP/ASP characters in <script> and <style>
//  - sanitizing URI values
//  - removing extra <title> elements
func NodeCleanup(e *Context) error {
	for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
		switch n.Type {
		case html.CommentNode:
			// Strip out comment nodes.
			htmlnode.RemoveNode(&n)
			continue

		case html.ElementNode:
			// TODO(b/79415817): Removing <noscript> is a temporary fix until we know how to handle them.
			if n.DataAtom == atom.Noscript {
				htmlnode.RemoveNode(&n)
				continue
			}

			// Deduplicate attributes from element nodes
			n.Attr = uniqueAttributes(n.Attr)

			// Strip out nonce attributes
			for i := len(n.Attr) - 1; i >= 0; i-- {
				if n.Attr[i].Key == "nonce" {
					n.Attr = append(n.Attr[:i], n.Attr[i+1:]...)
				}
			}

			// Sanitize URI attribute values.
			n.Attr = sanitizeURIAttributes(n.Attr)

			// Remove extra <title> elements
			if n.DataAtom == atom.Title {
				maybeStripTitle(&n)
			}

		case html.DoctypeNode:
			// Force doctype to be HTML 5.
			n.Data = "html"
			n.Attr = nil

		case html.TextNode:
			// Strip out whitespace only text nodes that are not in <body> or <title>.
			if len(strings.TrimLeft(n.Data, whitespace)) == 0 && !htmlnode.IsDescendantOf(n, atom.Body) && !htmlnode.IsChildOf(n, atom.Title) {
				htmlnode.RemoveNode(&n)
				continue
			}
			// Escape JSP/ASP characters
			if htmlnode.IsChildOf(n, atom.Script) {
				n.Data = scriptJSPReplacer.Replace(n.Data)
			} else if htmlnode.IsChildOf(n, atom.Style) {
				n.Data = styleJSPReplacer.Replace(n.Data)
			}
		}
	}

	// Find and fix amp-custom style after recursion above, which
	// would have removed whitespace only children nodes. This call
	// will then properly remove the empty style node.
	findAndFixStyleAMPCustom(e.DOM.HeadNode)
	return nil
}

// Returns the unique attributes (based off the attribute key), keeping
// the first one encountered.
func uniqueAttributes(attrs []html.Attribute) []html.Attribute {
	u := make([]html.Attribute, 0, len(attrs))
	m := make(map[string]struct{})
	for _, a := range attrs {
		if _, ok := m[a.Key]; !ok {
			m[a.Key] = struct{}{}
			u = append(u, a)
		}
	}
	return u
}

// Sanitizes all any possible URI values (href or src), modifying the
// input slice, and returning it as well.
func sanitizeURIAttributes(attrs []html.Attribute) []html.Attribute {
	for i := range attrs {
		if attrs[i].Key == "src" || attrs[i].Key == "href" {
			attrs[i].Val = strings.Map(func(r rune) rune {
				if strings.ContainsRune(unsanitaryURIChars, r) {
					return -1
				}
				return r
			}, attrs[i].Val)
		}
	}
	return attrs
}

// findAndFixStyleAMPCustom finds the <style amp-custom> element and
// if it is empty, removes it, or if not empty, strips all remaining
// attributes.
// There can only be one <style amp-custom> element and only within head.
// https://www.ampproject.org/docs/design/responsive_amp#add-styles-to-a-page
func findAndFixStyleAMPCustom(h *html.Node) {
	if h.DataAtom != atom.Head {
		return
	}
	for c := h.FirstChild; c != nil; c = c.NextSibling {
		if c.DataAtom == atom.Style && htmlnode.HasAttribute(c, "", amphtml.AMPCustom) {
			// Strip empty nodes
			if c.FirstChild == nil && c.LastChild == nil {
				h.RemoveChild(c)
			} else {
				// Strip remaining attributes
				c.Attr = []html.Attribute{{Key: amphtml.AMPCustom}}
			}

			// there can only be one <style amp-custom>, so return
			return
		}
	}
}

// maybeStripTitle removes the given title element if it is extraneous.
// There can only be one in head and none in body (svgs are excepted).
func maybeStripTitle(n **html.Node) {
	if (*n).DataAtom != atom.Title || htmlnode.IsDescendantOf(*n, atom.Svg) {
		return
	}

	switch {
	case htmlnode.IsDescendantOf(*n, atom.Head):
		// If we are in head, see if there are any previous title siblings,
		// and if so, strip this one.
		for c := (*n).PrevSibling; c != nil; c = c.PrevSibling {
			if c.DataAtom == atom.Title {
				htmlnode.RemoveNode(n)
				return
			}
		}
	case htmlnode.IsDescendantOf(*n, atom.Body):
		// Strip any titles found in body.
		htmlnode.RemoveNode(n)
	}
}
