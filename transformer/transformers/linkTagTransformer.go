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
	"net/url"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

// LinkTagTransformer operates on the <link> tag.
// * It will rename author supplied resource hints from rel= to disabled-rel=.
// * It will add a preconnect link tag for Google Font resources.
func LinkTagTransformer(e *Engine) {
	if _, ok := amphtml.NewDOM(e.Doc); !ok {
		return
	}

	preconnectAdded := false

	var stk htmlnode.Stack
	stk.Push(e.Doc)
	for len(stk) > 0 {
		top := stk.Pop()
		// Traverse the childen in reverse order so the iteration of
		// the DOM tree traversal is in the proper sequence.
		// E.g. Given <a><b/><c/></a>, we will visit a, b, c.
		// An alternative is to traverse childen in forward order and
		// utilize a queue instead.
		for c := top.LastChild; c != nil; c = c.PrevSibling {
			stk.Push(c)
		}
		linkTagTransform(top, &preconnectAdded)
	}
}

// linkTagTransform does the actual work on each node.
func linkTagTransform(n *html.Node, preconnectAdded *bool) {
	if htmlnode.HasAttribute(n, "rel") {
		renameAuthorSuppliedResourceHints(n)
	}
	if !*preconnectAdded && isLinkGoogleFont(n) {
		addLinkGoogleFontPreconnect(n)
		*preconnectAdded = true
	}
}

// isGoogleFontHostname returns true if the given string, after being parsed as
// a URL has the hostname of "fonts.googleapis.com".
func isGoogleFontHostname(s string) bool {
	u, err := url.Parse(s)
	return err == nil && strings.EqualFold(u.Hostname(), "fonts.googleapis.com")
}

// isLinkGoogleFont returns true if the given node is a link tag, has attribute
// href with the Google Font hostname.
func isLinkGoogleFont(n *html.Node) bool {
	if n.DataAtom != atom.Link {
		return false
	}
	v, ok := htmlnode.GetAttributeVal(n, "href")
	return ok && isGoogleFontHostname(v)
}

// addLinkGoogleFontPreconnect adds a preconnect link tag for Google Font resources.
func addLinkGoogleFontPreconnect(n *html.Node) {
	if n.DataAtom != atom.Link {
		return
	}
	preconnect := htmlnode.Element("link", html.Attribute{Key: "crossorigin"}, html.Attribute{Key: "href", Val: "https://fonts.gstatic.com"}, html.Attribute{Key: "rel", Val: "dns-prefetch preconnect"})
	n.Parent.InsertBefore(preconnect, n)
}

// renameAuthorSuppliedResourceHints renames author supplied resource hints from
// rel= to disabled=rel=.
func renameAuthorSuppliedResourceHints(n *html.Node) {
	r, ok := htmlnode.FindAttribute(n, "", "rel")
	if !ok {
		return
	}
	s := strings.Split(strings.ToLower(r.Val), " ")
	for _, h := range s {
		switch h {
		case "dns-prefetch", "preconnect", "prefetch", "preload", "prerender":
			htmlnode.SetAttribute(n, "", "disabled-rel", r.Val)
			htmlnode.RemoveAttribute(n, r)
			return
		}
	}
}
