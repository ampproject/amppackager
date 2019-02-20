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
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/css"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

type rewritable interface {
	// Rewrite the URLs within
	rewrite(string, *url.URL, string, nodeMap)
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

// nodeMap is map of string to Node pointer
type nodeMap map[string]*html.Node

// URLRewrite rewrites links to point to the AMP Cache and adds DNS preconnects to the <head>
// Affected links:
//  * <amp-img/amp-anim src>
//  * <amp-img/amp-anim srcset>
//  * <img src> / <img srcset> within <noscript>
//  * <image href> / <image xlink:href> which are SVG-specific images.
//  * <link rel=icon href>
//  * <amp-video poster>
//  * <use xlink:href>
//  * a background image given in the <style amp-custom> tag / style attribute
//  * any fonts given in the <style amp-custom> tag / style attribute
//  * background attributes.
func URLRewrite(e *Context) error {
	var ctx urlRewriteContext

	// First traverse the DOM, finding all nodes that need to be rewritten, building up
	// contextual information - contextual information could describe attributes of a node,
	// or an entire text node (in the case of the child of a <script> tag).
	for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
		if n.Type == html.TextNode {
			continue
		}

		if n.DataAtom == atom.Style && htmlnode.HasAttribute(n, "", amphtml.AMPCustom) && n.FirstChild != nil {
			ctx.parseStyleText(n.FirstChild)
			continue
		}

		if v, ok := htmlnode.GetAttributeVal(n, "", "style"); ok {
			ctx.parseInlineStyle(n, v)
			continue
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
			if srcOk && len(src) > 0 {
				ctx.parseSimpleImageAttr(n, "", "src")
			}

			if v, srcsetOk := htmlnode.GetAttributeVal(n, "", "srcset"); srcsetOk && len(v) > 0 {
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
	documentURL := e.DocumentURL.String()
	convertToAMPCacheURLs(ctx, documentURL, e)
	return nil
}

// convertToAMPCacheURLs examines the generated context and:
// - rewrites all the necessary URLs to point to the AMP Cache
// - adds any preconnects, unless they exist already.
func convertToAMPCacheURLs(ctx urlRewriteContext, documentURL string, e *Context) {
	preconnects := make(nodeMap)
	mainSubdomain := amphtml.ToCacheURLSubdomain(e.BaseURL.Hostname())
	for _, rw := range ctx {
		rw.rewrite(documentURL, e.BaseURL, mainSubdomain, preconnects)
	}

	// Parse any existing preconnects.
	for c := e.DOM.HeadNode.FirstChild; c != nil; c = c.NextSibling {
		if c.DataAtom != atom.Link {
			continue
		}
		if v, ok := htmlnode.GetAttributeVal(c, "", "rel"); ok {
			if href, ok := htmlnode.GetAttributeVal(c, "", "href"); ok {
				fields := strings.Fields(v)
				m := make(map[string]struct{}, len(fields))
				for _, f := range fields {
					m[f] = struct{}{}
				}
				if containsKey(m, "dns-prefetch") && containsKey(m, "preconnect") {
					// Remove the link as it will be added back later
					preconnects[href] = htmlnode.RemoveNode(&c)
				}
			}
		}
	}

	sortedPreconnects := make([]string, 0, len(preconnects))
	for k := range preconnects {
		sortedPreconnects = append(sortedPreconnects, k)
	}
	sort.Strings(sortedPreconnects)
	for _, k := range sortedPreconnects {
		n, ok := preconnects[k]
		if !ok || n == nil {
			n = htmlnode.Element("link", html.Attribute{Key: "href", Val: k}, html.Attribute{Key: "rel", Val: "dns-prefetch preconnect"})
		}
		e.DOM.HeadNode.AppendChild(n)
	}
}

func containsKey(m map[string]struct{}, k string) bool {
	_, ok := m[k]
	return ok
}

// rewrite the URLs described by the elementNodeContext. rewriteable implementation.
func (nc *elementNodeContext) rewrite(documentURL string, baseURL *url.URL,
	mainSubdomain string, preconnects nodeMap) {
	if len(nc.attrName) == 0 || len(nc.offsets) == 0 {
		return
	}
	attrVal, ok := htmlnode.GetAttributeVal(nc.node, nc.attrNS, nc.attrName)
	if !ok {
		return
	}
	replaced := replaceURLs(attrVal, nc.offsets, documentURL, baseURL,
		mainSubdomain, preconnects)
	htmlnode.SetAttribute(nc.node, nc.attrNS, nc.attrName, replaced)
}

// rewrite the URLs described by the textNodeContext. rewriteable implementation.
func (nc *textNodeContext) rewrite(documentURL string, baseURL *url.URL,
	mainSubdomain string, preconnects nodeMap) {
	nc.node.Data = replaceURLs(nc.node.Data, nc.offsets, documentURL, baseURL,
		mainSubdomain, preconnects)
}

// replaceURLs replaces all the URLs in the data string found at the various
// offsets with their AMP Cache equivalent, returning a new data string.
func replaceURLs(data string, offsets []amphtml.SubresourceOffset,
	documentURL string, baseURL *url.URL, mainSubdomain string,
	preconnects nodeMap) string {
	if len(offsets) == 0 {
		// noop
		return data
	}
	var sb strings.Builder
	pos := 0
	for _, so := range offsets {
		if pos < so.Start {
			// Add any non-URL text
			sb.WriteString(data[pos:so.Start])
			pos = so.Start
		}
		cu, err := so.GetCacheURL(documentURL, baseURL, data)
		if err != nil {
			// noop
			continue
		}
		sb.WriteString(cu.String())
		pos = so.End
		if len(mainSubdomain) > 0 && mainSubdomain != cu.Subdomain {
			preconnects[cu.OriginDomain()] = nil
		}
	}
	// Append any remaining non-URL text
	if pos < len(data) {
		sb.WriteString(data[pos:])
	}
	if sb.Len() > 0 {
		return sb.String()
	}
	return data
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

// parseSimpleImageAttr parses the specified attribute value, calculating the offset to the
// referenced subresource.
func (ctx *urlRewriteContext) parseSimpleImageAttr(n *html.Node, namespace, attrName string) {
	if v, ok := htmlnode.GetAttributeVal(n, namespace, attrName); ok && !amphtml.IsCacheURL(v) {
		nc := elementNodeContext{n, namespace, attrName, []amphtml.SubresourceOffset{amphtml.SubresourceOffset{
			SubType: amphtml.ImageType, Start: 0, End: len(v)}}}
		*ctx = append(*ctx, &nc)
	}
}

// parseStyleText parses the text node that contains stylesheet text.
func (ctx *urlRewriteContext) parseStyleText(n *html.Node) {
	if n.Type != html.TextNode {
		return
	}
	parsed, offsets := parseCSS(n.Data)
	if len(offsets) > 0 {
		*ctx = append(*ctx, &textNodeContext{n, offsets})
		n.Data = parsed
	}
}

// parseInlineStyle parses an inline style attribute, deriving the offsets for any URLs
func (ctx *urlRewriteContext) parseInlineStyle(n *html.Node, style string) {
	parsed, offsets := parseCSS(style)
	if len(offsets) > 0 {
		*ctx = append(*ctx, &elementNodeContext{n, "", "style", offsets})
		htmlnode.SetAttribute(n, "", "style", parsed)
	}
}

// Re-escape for URL serialization.
// https://www.w3.org/TR/css3-values/#strings
// https://www.w3.org/TR/css3-values/#urls
var reescapeURL = strings.NewReplacer(
	"\\", "\\\\",
	"\n", "\\A ",
	"'", "\\'",
)

// Also, prevent closing the style tag inside a data: string.
// \3C (followed by a space) is the CSS way of encoding <.
var closeStyle = regexp.MustCompile("(?i)</style")

func parseCSS(style string) (string, []amphtml.SubresourceOffset) {
	segments, err := css.ParseURLs(style)
	if err != nil {
		return style, []amphtml.SubresourceOffset{}
	}
	var sb strings.Builder
	pos := 0
	var offsets []amphtml.SubresourceOffset
	for _, segment := range segments {
		if segment.Type == css.ByteType {
			writeAndMark(&sb, &pos, segment.Data)
			continue
		}
		writeAndMark(&sb, &pos, "url('")
		if len(segment.Data) > 0 && !amphtml.IsCacheURL(segment.Data) {
			urlVal := reescapeURL.Replace(segment.Data)
			urlVal = closeStyle.ReplaceAllString(urlVal, "\\3C /style")
			slen, _ := sb.WriteString(urlVal)
			subtype := amphtml.ImageType
			if segment.Type == css.FontURLType {
				subtype = amphtml.OtherType
			}
			offsets = append(offsets, amphtml.SubresourceOffset{SubType: subtype, Start: pos, End: pos + slen})
			pos += slen
		}
		writeAndMark(&sb, &pos, "')")
	}
	return sb.String(), offsets
}

func writeAndMark(sb *strings.Builder, pos *int, s string) {
	slen, _ := sb.WriteString(s)
	*pos += slen
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
	if len(src) == 0 || strings.HasPrefix(src, "data:image/") || amphtml.IsCacheURL(src) {
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
			writeAndMark(&sb, &pos, " "+strconv.Itoa(w)+"w")
			if i < len(widths)-1 {
				writeAndMark(&sb, &pos, ", ")
			}
		}
		*ctx = append(*ctx, &nc)
	}
	if pos > 0 {
		htmlnode.SetAttribute(n, "", nc.attrName, sb.String())
	}
}
