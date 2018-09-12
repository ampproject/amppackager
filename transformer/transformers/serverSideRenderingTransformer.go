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
	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"github.com/ampproject/amppackager/transformer/layout"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

// ServerSideRenderingTransformer implements server-side rendering,
// described in http://go/amp-ssr. In short, it calculates the AMP layout
// server-side and expresses it by annotating the document with style
// attributes etc. And if possible, it removes the boilerplate.
func ServerSideRenderingTransformer(e *Engine) {
	dom, ok := amphtml.NewDOM(e.Doc)
	if !ok {
		return
	}

	// Simple check to ensure server-side rendering is only applied once.
	if _, ok := htmlnode.FindAttribute(dom.HTMLNode, "", amphtml.IAMPHTMLLayout); ok {
		return
	}
	htmlnode.SetAttribute(dom.HTMLNode, "", amphtml.IAMPHTMLLayout, "")

	// Assume the boilerplate can be removed, unless proven otherwise.
	remove := true

	transform(dom.BodyNode, &remove)

	// Emit the amp-runtime marker to indicate that server side
	// rendering has been applied.
	ampRuntimeMarker := htmlnode.Element("style", html.Attribute{Key: "amp-runtime"})
	dom.HeadNode.InsertBefore(ampRuntimeMarker, dom.HeadNode.FirstChild)

	// Also check the <head> tag if boilerplate is needed or not.
	remove = remove && canRemoveBoilerplateRecursive(dom.HeadNode)
	if remove {
		htmlnode.SetAttribute(dom.HTMLNode, "", "i-amphtml-no-boilerplate", "")

		// Find the boilerplate and remove it
		removeBoilerplate(dom.HeadNode)
	}
}

// transform recursively calls ApplyLayout to each AMP custom element,
// and at the same time, checks if the boilerplate can be removed.
func transform(n *html.Node, remove *bool) {
	// Skip tags inside a template tag.
	if htmlnode.IsDescendantOf(n, atom.Template) {
		return
	}

	if amphtml.IsAMPCustomElement(n) {
		if *remove {
			*remove = canRemoveBoilerplate(n)
		}

		// TODO(sedano): remove when SSR overwrites declarations.
		if _, ok := htmlnode.FindAttribute(n, "", "style"); ok {
			return
		}

		// If ApplyLayout encounters any unsupported layout, the
		// boilerplate cannot be removed.
		if err := layout.ApplyLayout(n); err != nil {
			*remove = false
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		transform(c, remove)
	}
}

// canRemoveBoilerplate checks if any attributes or tags exist on node
// n that need the boilerplate, and returns 'false' (meaning the
// boilerplate is required and cannot be removed).
func canRemoveBoilerplate(n *html.Node) bool {
	if n.Type != html.ElementNode || htmlnode.IsDescendantOf(n, atom.Template) {
		return true
	}

	if amphtml.IsAMPCustomElement(n) && htmlnode.IsDescendantOf(n, atom.Body) {
		// amp-experiment is a render delaying extension iff the tag is used in
		// the doc.
		if n.Data == amphtml.AMPExperiment {
			return false
		}
		// amp-audio requires knowing the dimensions of the browser. Do not
		// remove the boilerplate or apply layout if amp-audio is present in the
		// document.
		if n.Data == amphtml.AMPAudio {
			return false
		}
		if _, ok := htmlnode.FindAttribute(n, "", "heights"); ok {
			return false
		}
		if _, ok := htmlnode.FindAttribute(n, "", "media"); ok {
			return false
		}
		if _, ok := htmlnode.FindAttribute(n, "", "sizes"); ok {
			return false
		}
		// TODO(sedano): remove when SSR overwrites declarations.
		if _, ok := htmlnode.FindAttribute(n, "", "style"); ok {
			return false
		}
	}

	if n.DataAtom == atom.Script && htmlnode.IsDescendantOf(n, atom.Head) {
		if a, ok := htmlnode.FindAttribute(n, "", amphtml.AMPCustomElement); ok {
			// TODO(b/77581738): Remove amp-story from here.
			if a.Val == amphtml.AMPDynamicCSSClasses || a.Val == amphtml.AMPStory {
				return false
			}
		}
	}
	return true
}

// canRemoveBoilerplateRecursive recursively calls
// canRemoveBoilerplate on every node (depth-first) of n, returning as
// soon as anything is found that requires keeping the boilerplate.
func canRemoveBoilerplateRecursive(n *html.Node) bool {
	remove := canRemoveBoilerplate(n)

	for c := n.FirstChild; c != nil && remove; c = c.NextSibling {
		remove = canRemoveBoilerplateRecursive(c)
	}
	return remove
}

// removeBoilerplate removes the AMP boilerplate script (and noscript) tags
// from the given head node.
func removeBoilerplate(n *html.Node) {
	if n.DataAtom != atom.Head {
		return
	}
	for c := n.FirstChild; c != nil; {
		next := c.NextSibling
		if c.DataAtom == atom.Noscript {
			n.RemoveChild(c)
		} else if c.DataAtom == atom.Style {
			if _, ok := htmlnode.FindAttribute(c, "", amphtml.AMPBoilerplate); ok {
				n.RemoveChild(c)
			} else if _, ok := htmlnode.FindAttribute(c, "", amphtml.AMP4AdsBoilerplate); ok {
				n.RemoveChild(c)
			} else if _, ok := htmlnode.FindAttribute(c, "", amphtml.AMP4EmailBoilerplate); ok {
				n.RemoveChild(c)
			}
		}
		c = next
	}
}
