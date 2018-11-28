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
	"strconv"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

type nodeContext struct {
	node    *html.Node
	baseURL *url.URL
}

// URLRewrite rewrites links to point to the AMP Cache and (TODO) adds DNS preconnects to the <head>
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
	// TODO(alin04): Populate the preconnects
	var preconnects []string
	for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
		if n.Type == html.TextNode {
			continue
		}

		nc := nodeContext{n, e.BaseURL}

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
		nc.rewriteSimpleImgAttr("", "background")

		switch n.Data {
		case "link":
			// Rewrite 'href' attribute within <link rel="icon" href=...> and variants
			// to point into the AMP Cache.
			if htmlnode.HasAttribute(n, "", "href") {
				if v, ok := htmlnode.GetAttributeVal(n, "", "rel"); ok && fieldsContain(v, "icon") {
					nc.rewriteSimpleImgAttr("", "href")
				}
			}

		case "amp-img", "amp-anim", "img":
			// Rewrite 'src' and 'srcset' attributes. Add 'srcset' if none.
			src, srcOk := htmlnode.GetAttributeVal(nc.node, "", "src")
			if srcOk {
				nc.rewriteSimpleImgAttr("", "src")
			}

			if v, srcsetOk := htmlnode.GetAttributeVal(n, "", "srcset"); srcsetOk {
				htmlnode.SetAttribute(n, "", "srcset", amphtml.ConvertSrcset(e.BaseURL, v))
			} else if srcOk {
				nc.addSrcset(src)
			}

		case "amp-video", "video":
			nc.rewriteSimpleImgAttr("", "poster")

		case "image":
			// For b/78468289, rewrite the 'href' or `xlink:href` attribute on an
			// svg <image> tag to point into the AMP Cache.
			nc.rewriteSimpleImgAttr("", "href")
			nc.rewriteSimpleImgAttr("xlink", "href")

		case "use":
			nc.rewriteSimpleImgAttr("xlink", "href")
		}
	}

	for _, preconnect := range preconnects {
		n := htmlnode.Element("link", html.Attribute{Key: "href", Val: preconnect}, html.Attribute{Key: "rel", Val: "dns-prefetch preconnect"})
		e.DOM.HeadNode.AppendChild(n)
	}

	return nil
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

// rewriteSimpleImgAttr rewrites the specified attribute value to point into the AMP cache.
func (nc *nodeContext) rewriteSimpleImgAttr(namespace, attrName string) {
	if v, ok := htmlnode.GetAttributeVal(nc.node, namespace, attrName); ok {
		req := amphtml.ImageURLRequest{Input: amphtml.ToPortableURL(nc.baseURL, v)}
		htmlnode.SetAttribute(nc.node, namespace, attrName, req.GetCacheImageURL())
	}
}

// Do not add srcset for responsive layout if the width attribute is smaller
// than this value. In the responsive value, width and height might be used
// for indicating the aspect ratio instead of the actual render dimension.
// This happens often when the width and height have small values. Value of
// 300 is chosen based on the assumption that it is large enough to be the
// render dimension, however, we may need to adjust this value if the assumption
// is found invalid later.
const minWidthToAddSrcsetInResponsiveLayout = 300

// addSrcset adds a srcset attribute, if applicable.
func (nc *nodeContext) addSrcset(src string) {
	if strings.HasPrefix(src, "data:image/") {
		return
	}
	var width int
	if widthVal, ok := htmlnode.GetAttributeVal(nc.node, "", "width"); ok {
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
	layout, layoutOk := htmlnode.GetAttributeVal(nc.node, "", "layout")
	isResponsiveLayout := (layoutOk && layout == "responsive") ||
		(!layoutOk && htmlnode.HasAttribute(nc.node, "", "height") && htmlnode.HasAttribute(nc.node, "", "sizes"))
	// In responsive layout, width and height might be used for indicating
	// the aspect ratio instead of the actual render dimensions. This usually
	// happens for dimensions of small values.
	if isResponsiveLayout && width < minWidthToAddSrcsetInResponsiveLayout {
		return
	}

	absolute := amphtml.ToPortableURL(nc.baseURL, src)
	if sc, ok := amphtml.GetSrcsetFromSrc(absolute, width); ok {
		htmlnode.SetAttribute(nc.node, "", "srcset", sc)
	}
}
