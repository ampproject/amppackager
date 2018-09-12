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

// MetaTagTransformer operates on the <meta> tag.
// * It will strip some meta tags.
// * It will relocate all meta tags found inside the body into the head.
//
// It does *not* sort the meta tags. This is done by ReorderHeadTransformer.
// TODO(sedano): The naming is repetitive with the package name as this is
// transformer.MetaTagTransformer. Consider when porting is done to remove the
// duplicative Transformer (e.g. this becomes transformer.MetaTag).
func MetaTagTransformer(e *Engine) {
	dom, ok := amphtml.NewDOM(e.Doc)
	if !ok {
		return
	}

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
		metaTagTransform(top, dom.HeadNode)
	}
}

// metaTagTransform does the actual work on each node.
func metaTagTransform(n, h *html.Node) {
	// Skip non-meta tags.
	if n.DataAtom != atom.Meta {
		return
	}

	if shouldStripMetaTag(n) {
		n.Parent.RemoveChild(n)
		return
	}

	// Relocate meta tags in body, which are not stripped, into head.
	if htmlnode.IsDescendantOf(n, atom.Body) {
		n.Parent.RemoveChild(n)
		h.AppendChild(n)
	}
}

// shouldStripMetaTag returns true if the given node should be removed
// from the document. See b/34885784 for details.
func shouldStripMetaTag(n *html.Node) bool {
	if n.DataAtom != atom.Meta {
		return false
	}

	// Keep <meta charset> tag.
	if htmlnode.HasAttribute(n, "charset") {
		return false
	}

	// Keep <meta> tags that have attribute http-equiv except if
	// value=x-dns-prefetch-control
	if v, ok := htmlnode.GetAttributeVal(n, "http-equiv"); ok {
		return strings.EqualFold(v, "x-dns-prefetch-control")
	}

	// Keep <meta> tags that don't have attributes content, itemprop, name,
	// and property.
	if !htmlnode.HasAttribute(n, "content") && !htmlnode.HasAttribute(n, "itemprop") && !htmlnode.HasAttribute(n, "name") && !htmlnode.HasAttribute(n, "property") {
		return false
	}

	// Keep <meta name=...> tags if name:
	//   - Has prefix "amp-"
	//   - Has prefix "amp4ads-"
	//   - Has prefix "dc."
	//   - Has prefix "i-amphtml-"
	//   - Has prefix "twitter:"
	//   - Is "apple-itunes-app"
	//   - Is "copyright"
	//   - Is "referrer"
	//   - Is "viewport"
	if v, ok := htmlnode.GetAttributeVal(n, "name"); ok {
		v = strings.ToLower(v)
		if strings.HasPrefix(v, "amp-") ||
			strings.HasPrefix(v, "amp4ads-") ||
			strings.HasPrefix(v, "dc.") ||
			strings.HasPrefix(v, "i-amphtml-") ||
			strings.HasPrefix(v, "twitter:") ||
			v == "apple-itunes-app" ||
			v == "copyright" ||
			v == "referrer" ||
			v == "viewport" {
			return false
		}
	}

	// Keep <meta property=...> tags if property:
	//   (1) Has prefix "al:"
	//   (2) Has prefix "fb:"
	//   (3) Has prefix "og:"
	if v, ok := htmlnode.GetAttributeVal(n, "property"); ok {
		v = strings.ToLower(v)
		if strings.HasPrefix(v, "al:") ||
			strings.HasPrefix(v, "fb:") ||
			strings.HasPrefix(v, "og:") {
			return false
		}
	}

	// Keep <meta itemprop> when name attribute is not present.
	// For amp-subscriptions (see b/75965615).
	if htmlnode.HasAttribute(n, "itemprop") && !htmlnode.HasAttribute(n, "name") {
		return false
	}

	return true
}
