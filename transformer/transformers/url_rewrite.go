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
	"sort"
	"strconv"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

type rewritable interface {
	// Rewrite the URLs within
	rewrite(*url.URL, string, map[string]struct{})
}

type urlRewriteContext []rewritable

// elementNodeContext qualifies the attribute value of an element node, that has URLs
// which need to be rewritten. In many cases, the node attribute value
// corresponds to a single subresource, e.g. an image src, or a body background,
// in which case, there will only be a single offset location.
// For image srcset, which could be a comma delimited list, or an inline style,
// URLs will be interspersed in the attribute value and therefore there will be
// multiple offset locations.
type elementNodeContext struct {
	node             *html.Node
	attrNS, attrName string
	offsets          []amphtml.SubresourceOffset
}

// textNodeContext qualifies a text node, whose entire data value has URLs embedded
// within. This struct keeps track of where the URLs occur in the body of the text.
type textNodeContext struct {
	node    *html.Node
	offsets []amphtml.SubresourceOffset
}

// URLRewrite rewrites links to point to the AMP Cache and adds DNS preconnects to the <head>
// Affected links:
//  * <amp-img/amp-anim src>
//  * <amp-img/amp-anim srcset>
//  * <img src> / <img srcset> within <noscript>
//  * <image href> / <image xlink:href> which are SVG-specific images.
//  * <link rel=icon href>
//  * <amp-video poster>
//  * <use xlink:href>
//  * TODO(alin04): a background image given in the <style amp-custom> tag / style attribute
//  * TODO(alin04): any fonts given in the <style amp-custom> tag / style attribute
//  * background attributes.
func URLRewrite(e *Context) error {
	var ctx urlRewriteContext

	// First travers the DOM, finding all nodes that need to be rewritten, building up
	// contextual information - contextual information could describe attributes of a node,
	// or an entire text node (in the case of the child of a <script> tag).
	for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
		if n.Type == html.TextNode {
			continue
		}

		if n.DataAtom == atom.Style && htmlnode.HasAttribute(n, "", amphtml.AMPCustom) {
			// TODO(alin04): parse url tokens in css
			continue
		}

		if htmlnode.HasAttribute(n, "", "style") {
			// TODO(alin04): parse url tokens in css
		}

		// Do not rewrite links within mustache templates.
		if htmlnode.IsDescendantOf(n, atom.Template) {
			continue
		}

		// For b/78468289, rewrite the 'background' attribute on any element
		// to point into the AMP Cache. At the time of writing this code, no
		// validator rule actually allows this attribute, but we want to have
		// this in place as defense in depth in case the attribute is added
		// in the future.
		ctx.parseSimpleImageAttr(n, "", "background")

		switch n.Data {
		case "link":
			// Rewrite 'href' attribute within <link rel="icon" href=...> and variants
			// to point into the AMP Cache.
			if htmlnode.HasAttribute(n, "", "href") {
				if v, ok := htmlnode.GetAttributeVal(n, "", "rel"); ok && fieldsContain(v, "icon") {
					ctx.parseSimpleImageAttr(n, "", "href")
				}
			}

		case "amp-img", "amp-anim", "img":
			// Rewrite 'src' and 'srcset' attributes. Add 'srcset' if none.
			src, srcOk := htmlnode.GetAttributeVal(n, "", "src")
			if srcOk {
				ctx.parseSimpleImageAttr(n, "", "src")
			}

			if v, srcsetOk := htmlnode.GetAttributeVal(n, "", "srcset"); srcsetOk {
				ctx.parseExistingSrcset(n, v)
			} else if srcOk {
				ctx.parseNewSrcset(n, src)
			}

		case "amp-video", "video":
			ctx.parseSimpleImageAttr(n, "", "poster")

		case "image":
			// For b/78468289, rewrite the 'href' or `xlink:href` attribute on an
			// svg <image> tag to point into the AMP Cache.
			ctx.parseSimpleImageAttr(n, "", "href")
			ctx.parseSimpleImageAttr(n, "xlink", "href")

		case "use":
			ctx.parseSimpleImageAttr(n, "xlink", "href")
		}
	}
	// After the contextual information has been gathered, use it to rewrite the appropriate URLs
	// and adding any preconnects if necessary.
	preconnects := convertToAMPCacheURLs(ctx, e.BaseURL)
	for _, k := range preconnects {
		n := htmlnode.Element("link", html.Attribute{Key: "href", Val: k}, html.Attribute{Key: "rel", Val: "dns-prefetch preconnect"})
		e.DOM.HeadNode.AppendChild(n)
	}

	return nil
}

// convertToAMPCacheURLs examines the generated context and rewrites all the necessary
// URLs to point to the AMP Cache, returning a list of preconnects that need to be added.
func convertToAMPCacheURLs(ctx urlRewriteContext, base *url.URL) []string {
	preconnects := make(map[string]struct{})
	mainSubdomain := amphtml.ToCacheURLSubdomain(base.Hostname())
	for _, rw := range ctx {
		rw.rewrite(base, mainSubdomain, preconnects)
	}
	sortedPreconnects := make([]string, 0, len(preconnects))
	for k := range preconnects {
		sortedPreconnects = append(sortedPreconnects, k)
	}
	sort.Strings(sortedPreconnects)
	return sortedPreconnects
}

// rewrite the URLs described by the elementNodeContext
func (nc *elementNodeContext) rewrite(base *url.URL, mainSubdomain string, preconnects map[string]struct{}) {
	if len(nc.attrName) == 0 || len(nc.offsets) == 0 {
		return
	}
	attrVal, ok := htmlnode.GetAttributeVal(nc.node, nc.attrNS, nc.attrName)
	if !ok {
		return
	}
	var sb strings.Builder
	pos := 0
	for _, so := range nc.offsets {
		if pos < so.Start {
			// Add any non-URL text
			sb.WriteString(attrVal[pos:so.Start])
		}
		cu, err := so.GetCacheURL(base, attrVal)
		if err != nil {
			// noop
			continue
		}
		sb.WriteString(cu.String())
		pos = so.End
		if len(mainSubdomain) > 0 && mainSubdomain != cu.Subdomain {
			preconnects[cu.OriginDomain()] = struct{}{}
		}
	}
	// Append any remaining non-URL text
	if pos < len(attrVal) {
		sb.WriteString(attrVal[pos:])
	}
	if sb.Len() > 0 {
		htmlnode.SetAttribute(nc.node, nc.attrNS, nc.attrName, sb.String())
	}
}

// fieldsContain returns true if needle is a field in the haystack (case-insensitive).
// So "needle", "foo needle", "needle foo", "foo needle foo" return true,
// but "fooneedle" returns false.
func fieldsContain(haystack, needle string) bool {
	for _, s := range strings.Fields(haystack) {
		if strings.EqualFold(s, needle) {
			return true
		}
	}
	return false
}

// parseSimpleImageAttr parses the specified attribute value, deriving the offset to the
// referenced subresource.
func (ctx *urlRewriteContext) parseSimpleImageAttr(n *html.Node, namespace, attrName string) {
	if v, ok := htmlnode.GetAttributeVal(n, namespace, attrName); ok {
		nc := elementNodeContext{n, namespace, attrName, []amphtml.SubresourceOffset{amphtml.SubresourceOffset{
			SubType: amphtml.ImageType, Start: 0, End: len(v)}}}
		*ctx = append(*ctx, &nc)
	}
}

// parseExistingSrcset normalizes and parses the srcset attribute.
func (ctx *urlRewriteContext) parseExistingSrcset(n *html.Node, srcset string) {
	normalized, offsets := amphtml.ParseSrcset(srcset)
	nc := elementNodeContext{n, "", "srcset", offsets}
	*ctx = append(*ctx, &nc)
	htmlnode.SetAttribute(n, "", "srcset", normalized)
}

// Do not add srcset for responsive layout if the width attribute is smaller
// than this value. In the responsive value, width and height might be used
// for indicating the aspect ratio instead of the actual render dimension.
// This happens often when the width and height have small values. Value of
// 300 is chosen based on the assumption that it is large enough to be the
// render dimension, however, we may need to adjust this value if the assumption
// is found invalid later.
const minWidthToAddSrcsetInResponsiveLayout = 300

// parseNewSrcset parses a new srcset derived from the src.
func (ctx *urlRewriteContext) parseNewSrcset(n *html.Node, src string) {
	if len(src) == 0 || strings.HasPrefix(src, "data:image/") {
		return
	}
	var width int
	if widthVal, ok := htmlnode.GetAttributeVal(n, "", "width"); ok {
		var err error
		// TODO(b/113271759): Handle width values that include 'px' (probably others).
		if width, err = strconv.Atoi(widthVal); err != nil {
			// invalid width
			return
		}
	} else {
		// no width
		return
	}
	// Determine if the layout is "responsive".
	// https://www.ampproject.org/docs/guides/responsive/control_layout.html
	layout, layoutOk := htmlnode.GetAttributeVal(n, "", "layout")
	isResponsiveLayout := (layoutOk && layout == "responsive") ||
		(!layoutOk && htmlnode.HasAttribute(n, "", "height") && htmlnode.HasAttribute(n, "", "sizes"))
	// In responsive layout, width and height might be used for indicating
	// the aspect ratio instead of the actual render dimensions. This usually
	// happens for dimensions of small values.
	if isResponsiveLayout && width < minWidthToAddSrcsetInResponsiveLayout {
		return
	}

	var sb strings.Builder
	var pos int
	nc := elementNodeContext{node: n, attrName: "srcset"}
	if widths, ok := amphtml.GetSrcsetWidths(width); ok {
		for i, w := range widths {
			slen, _ := sb.WriteString(src)
			nc.offsets = append(nc.offsets, amphtml.SubresourceOffset{Start: pos, End: pos + slen, DesiredImageWidth: w})
			pos += slen
			if i < len(widths)-1 {
				slen, _ = sb.WriteString(", ")
				pos += slen
			}
		}
		*ctx = append(*ctx, &nc)
	}
	if pos > 0 {
		htmlnode.SetAttribute(n, "", nc.attrName, sb.String())
	}
}
