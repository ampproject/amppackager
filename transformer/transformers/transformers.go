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
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"github.com/ampproject/amppackager/transformer/printer"
	rpb "github.com/ampproject/amppackager/transformer/request"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

// An Engine stores the root DOM Node, configurable transformers to
// modify the DOM, and contextual data used for the transformers.
type Engine struct {
	Doc          *html.Node
	DocumentURL  *url.URL
	Transformers []func(*Engine)
	Request      *rpb.Request
}

// Transform executes required transformations on the Engine's Node and
// invokes any additionally configured transformers.
func (e *Engine) Transform() {
	requiredTransform(e.Doc)

	// Invoke the configured transformers
	for _, f := range e.Transformers {
		f(e)
	}
}

const (
	whitespace = " \t\r\n\f"

	// characters to be stripped out of URIs
	unsanitaryURIChars = "\t\n\r"
)

// Performs required transformations on the given node.
func requiredTransform(n *html.Node) {
	switch n.Type {
	case html.CommentNode:
		// Strip out comment nodes.
		n.Parent.RemoveChild(n)
		return

	case html.ElementNode:
		// Deduplicate attributes from element nodes
		n.Attr = uniqueAttributes(n.Attr)

		// Strip out nonce attributes
		for i := len(n.Attr) - 1; i >= 0; i-- {
			if n.Attr[i].Key == "nonce" {
				n.Attr = append(n.Attr[:i], n.Attr[i+1:]...)
			}
		}

		// Fix style amp-custom nodes.
		fixStyleAMPCustom(n)

		// Sanitize URI attribute values.
		n.Attr = sanitizeURIAttributes(n.Attr)

		// Remove extra <title> elements
		if n.DataAtom == atom.Title {
			stripExtraTitles(n)
		}

	case html.DoctypeNode:
		// Force doctype to be HTML 5.
		n.Data = "html"
		n.Attr = nil

	case html.TextNode:
		if n.Parent.Data == "noscript" && parseNoscriptContents(n) {
			return
		}

		// Strip out whitespace only text nodes, except in <body> or <title>.
		if !htmlnode.IsDescendantOf(n, atom.Body) && !htmlnode.IsChildOf(n, atom.Title) {
			s := strings.Trim(n.Data, whitespace)
			if len(s) == 0 {
				n.Parent.RemoveChild(n)
				return
			}
		}
	}

	var next *html.Node
	for c := n.FirstChild; c != nil; c = next {
		// Track the next sibling because if the node is removed in the recursive
		// call, it becomes orphaned and the pointer to NextSibling is lost.
		next = c.NextSibling
		requiredTransform(c)
	}
}

// Parse the contents of <noscript> tag, returning true if parsing was done.
// The golang tokenizer treats <noscript> children as one big TextNode, so
// retokenize to extract the elements.
// See https://golang.org/issue/16318
func parseNoscriptContents(n *html.Node) bool {
	parent := n.Parent
	if n.Type == html.TextNode && parent != nil && parent.Data == "noscript" {
		// Pass in <noscript>'s parent as the context. Passing <noscript> in
		// will result in the same behavior (one big TextNode), so remove
		// noscript from the context and use its parent (either head or body).
		parsed, err := html.ParseFragment(strings.NewReader(n.Data), parent.Parent)
		if err != nil {
			log.Fatal(err)
		}
		parent.RemoveChild(n)
		for _, c := range parsed {
			parent.AppendChild(c)
		}
		return true
	}
	return false
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

// fixStyleAMPCustom removes empty style amp-custom nodes, or if not empty,
// strips all remaining attributes.
func fixStyleAMPCustom(n *html.Node) {
	if n.DataAtom != atom.Style {
		return
	}
	if !htmlnode.HasAttribute(n, amphtml.AMPCustom) {
		return
	}

	// Strip empty nodes
	if n.FirstChild == nil && n.LastChild == nil {
		n.Parent.RemoveChild(n)
		return
	}
	// Strip remaining attributes
	n.Attr = []html.Attribute{{Key: amphtml.AMPCustom}}
}

// stripExtraTitles removes extraneous title elements. There can only be one
// in head and none in body (svgs are excepted).
func stripExtraTitles(n *html.Node) {
	if n.DataAtom != atom.Title || htmlnode.IsDescendantOf(n, atom.Svg) {
		return
	}

	switch {
	case htmlnode.IsDescendantOf(n, atom.Head):
		// If we are in head, the provided node is the one title element we
		// want to keep. Strip all others in head.
		for c := n.NextSibling; c != nil; {
			next := c.NextSibling
			if c.DataAtom == atom.Title {
				n.Parent.RemoveChild(c)
			}
			c = next
		}
	case htmlnode.IsDescendantOf(n, atom.Body):
		// Strip any titles found in body.
		n.Parent.RemoveChild(n)
	}
}

// Print renders the HTML to STDOUT.
func (e *Engine) Print() {
	err := printer.Print(os.Stdout, e.Doc)
	if err != nil {
		log.Fatal(err)
	}
}
