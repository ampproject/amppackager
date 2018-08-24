package transformer

import (
	"strings"

	"github.com/ampproject/amppackager/internal/pkg/amphtml"
	"github.com/ampproject/amppackager/internal/pkg/htmlnode"
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
	httpEquiv, ok := htmlnode.GetAttributeVal(n, "http-equiv")
	if ok && strings.EqualFold(httpEquiv, "x-dns-prefetch-control") {
		return true
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
	name, ok := htmlnode.GetAttributeVal(n, "name")
	name = strings.ToLower(name)
	if ok && (strings.HasPrefix(name, "amp-") ||
		strings.HasPrefix(name, "amp4ads-") ||
		strings.HasPrefix(name, "dc.") ||
		strings.HasPrefix(name, "i-amphtml-") ||
		strings.HasPrefix(name, "twitter:") ||
		name == "apple-itunes-app" ||
		name == "copyright" ||
		name == "referrer" ||
		name == "viewport") {
		return false
	}

	// Keep <meta property=...> tags if property:
	//   (1) Has prefix "al:"
	//   (2) Has prefix "fb:"
	//   (3) Has prefix "og:"
	property, ok := htmlnode.GetAttributeVal(n, "property")
	property = strings.ToLower(property)
	if ok && (strings.HasPrefix(property, "al:") ||
		strings.HasPrefix(property, "fb:") ||
		strings.HasPrefix(property, "og:")) {
		return false
	}

	// Keep <meta itemprop> when name attribute is not present.
	// For amp-subscriptions (see b/75965615).
	if htmlnode.HasAttribute(n, "itemprop") && !htmlnode.HasAttribute(n, "name") {
		return false
	}

	return true
}
