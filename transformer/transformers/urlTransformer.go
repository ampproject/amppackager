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

var /* const */ anyTagAttrs = []string{"src"}
var /* const */ ampInstallServiceWorkerTagAttrs = []string{"data-iframe-src", "data-no-service-worker-fallback-shell-url"}
var /* const */ ampStoryTagAttrs = []string{"background-audio", "bookend-config-src", "poster-landscape-src", "poster-square-src", "publisher-logo-src"}
var /* const */ ampStoryPageTagAttrs = []string{"background-audio"}
var /* const */ formTagAttrs = []string{"action", "action-xhr"}
var /* const */ imgTagAttrs = []string{"longdesc"}

// baseInfo encapsulates data about the <base> tag of the document.
type baseInfo struct {
	url    *url.URL
	target string
}

// URLTransformer operates on URL attributes, rewriting URLs as needed
// depending on whether the document is being served from the AMP
// Cache or not, relative to the base URL (which is derived from the
// document URL and the <base> tag).
//
// TODO(b/112361534): Handle the final URL of the document after all
// redirects.
//
// URLs will be rewritten to either be absolute (i.e. point to the
// origin document when served from the AMP Cache.) or not. In the
// latter case (aka portable), the URL could be absolute (yes
// confusing), a fragment-relative URL, or the exact text of a data
// URL.
//
// URLs are also canonicalized:
// * leading and trailing whitespace are trimmed.
//
// * The following attributes may be rewritten:
//   * Any tag (except amp-img [1]) with attribute:
//     * href
//     * src
//   * Any <amp-install-serviceworker> with attribute:
//     * data-iframe-src
//     * data-no-service-worker-fallback-shell-url
//   * Any <amp-story> tag with attribute:
//     * background-audio
//     * bookend-config-src
//     * poster-landscape-src
//     * poster-square-src
//     * publisher-logo-src
//   * Any <amp-story-page> tag with attribute:
//     * background-audio
//   * Any <form> tag with attribute:
//     * action
//     * action-xhr
//   * Any <img> tag with attribute:
//     * longdesc
//
//     [1]. TODO(b/112417267): Handle amp-img rewriting.
//
func URLTransformer(e *Engine) {
	dom, ok := amphtml.NewDOM(e.Doc)
	if !ok {
		return
	}
	baseInfo := extractBase(dom.HeadNode, e.DocumentURL)

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
		urlTransform(top, baseInfo)
	}
}

// extractBase returns a baseInfo struct that encapsulates the base
// URL and target values derived from the <base> tag, if it exists.
//
// The resulting baseInfo.url is an absolute URL using the document
// URL as the absolute base URI and base href as the reference to
// resolve.  baseInfo.target defaults to "_top" if not specified, or
// if the target isn't allowed.
func extractBase(n *html.Node, d *url.URL) *baseInfo {
	b := baseInfo{d, "_top"}
	if n, ok := htmlnode.FindNode(n, atom.Base); ok {
		if v, ok := htmlnode.GetAttributeVal(n, "href"); ok {
			u, err := url.Parse(v)
			if err != nil {
				b.url = u
			}
		}
		if v, ok := htmlnode.GetAttributeVal(n, "target"); ok && isAllowedTarget(v) {
			b.target = v
		}
	}
	return &b
}

// isAllowedTarget returns true if the given string is either "_blank" or "_top"
func isAllowedTarget(t string) bool {
	return strings.EqualFold(t, "_blank") || strings.EqualFold(t, "_top")
}

// urlTransform does the actual work of rewriting URLs
func urlTransform(n *html.Node, b *baseInfo) {
	// Skip text nodes
	if n.Type == html.TextNode {
		return
	}

	// TODO(b/112417267): Handle amp-img rewriting.
	if strings.EqualFold(n.Data, "amp-img") {
		return
	}

	// Make attributes with URLs portable on any tag
	rewritePortableURLs(n, b.url, anyTagAttrs)

	switch n.DataAtom {
	case atom.Form:
		// Make attributes with URLs absolute on <form> tag.
		rewriteAbsoluteURLs(n, b.url, formTagAttrs)
	case atom.Img:
		// Make attributes with URLs portable on <img> tag.
		rewritePortableURLs(n, b.url, imgTagAttrs)
	default:
		switch n.Data {
		case "amp-install-serviceworker":
			// Make attributes with URLs portable on <amp-install-serviceworker> tag.
			rewritePortableURLs(n, b.url, ampInstallServiceWorkerTagAttrs)
		case amphtml.AMPStory:
			// Make attributes with URLs portable on <amp-story> tag.
			rewritePortableURLs(n, b.url, ampStoryTagAttrs)
		case "amp-story-page":
			// Make attributes with URLs portable on <amp-story-page> tag.
			rewritePortableURLs(n, b.url, ampStoryPageTagAttrs)
		}
	}

	// Tags with href attribute.
	if href, ok := htmlnode.FindAttribute(n, "", "href"); ok {
		// Remove the base tag href with the following rationale:
		//
		// 1) The <base href> can be harmful. When handling things like image
		//    source sets which are re-hosted and served from
		//    https://cdn.ampproject.org, paths starting with "/" are rewritten
		//    into the stored html document with the intent that "/" is relative
		//    to the root of cdn.ampproject.org. If a base href were present, it
		//    would change the meaning of the relative links.
		//
		// 2) Other hrefs are absolutified in the document relative to the base
		//    href. Thus, it is not necessary to maintain the base href for
		//    browser URL resolution.
		if n.DataAtom == atom.Base {
			htmlnode.RemoveAttribute(n, href)
			if len(n.Attr) == 0 {
				n.Parent.RemoveChild(n)
				return
			}
		} else if v, ok := htmlnode.GetAttributeVal(n, "rel"); ok && n.DataAtom == atom.Link && v == "canonical" {
			// If the origin doc is self-canonical, it should be an absolute URL
			// and not portable (which would result in canonical = "#").
			// Maintain the original canonical, and absolutify it. See b/36102624
			in := htmlnode.IsDescendantOf(n, atom.Template)
			htmlnode.SetAttribute(n, "", "href", rewriteURL(b.url, in, href.Val, true))
		} else if n.DataAtom == atom.A {
			in := htmlnode.IsDescendantOf(n, atom.Template)
			portableHref := rewriteURL(b.url, in, href.Val, false)
			// Set a default target
			// 1. If the href is not a fragment AND
			// 2. If there is no target OR If there is a target and it is not an allowed target
			if !strings.HasPrefix(portableHref, "#") {
				if v, ok := htmlnode.GetAttributeVal(n, "target"); !ok || (ok && !isAllowedTarget(v)) {
					htmlnode.SetAttribute(n, "", "target", b.target)
				}
			}
			htmlnode.SetAttribute(n, "", "href", portableHref)
		} else {
			// Make a PortableUrl for any remaining tags with href.
			in := htmlnode.IsDescendantOf(n, atom.Template)
			htmlnode.SetAttribute(n, "", "href", rewriteURL(b.url, in, href.Val, false))
		}
	}
}

// rewriteAbsoluteURLs rewrites URLs in the given slice of attributes
// to be absolute for the base URL provided.
func rewriteAbsoluteURLs(n *html.Node, base *url.URL, tagAttrs []string) {
	rewriteURLs(n, base, tagAttrs, true)
}

// rewritePortableURLs rewrites URLs in the given slice of attributes
// to be portable relative to the base URL provided.
func rewritePortableURLs(n *html.Node, base *url.URL, tagAttrs []string) {
	rewriteURLs(n, base, tagAttrs, false)
}

func rewriteURLs(n *html.Node, base *url.URL, tagAttrs []string, absolute bool) {
	for _, attr := range tagAttrs {
		if v, ok := htmlnode.GetAttributeVal(n, attr); ok {
			in := htmlnode.IsDescendantOf(n, atom.Template)
			if in {
				return
			}
			htmlnode.SetAttribute(n, "", attr, rewriteURL(base, in, v, absolute))
		}
	}
}

func rewriteURL(base *url.URL, inTemplate bool, url string, absolute bool) string {
	if inTemplate {
		// For b/26741101, do not rewrite URLs within mustache templates
		return url
	}
	orig := url
	url = strings.TrimSpace(url)
	if url == "" {
		return orig
	}

	// For b/27292423:
	// In general, if the origin doc was fetched on http:// and has a relative
	// URL to a resource, we must assume that the resource may only be
	// available on http. However: if the subresource has a protocol-relative
	// path (beginning with '//') this is a clear statement that either
	// HTTP or HTTPS can work. Special-case the protocol-relative case.
	if strings.HasPrefix(url, "//") {
		return "https:" + url
	}
	u, err := base.Parse(url)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		return url
	}

	uVal := u.String()
	if absolute {
		return uVal
	}

	if uVal == base.String()+u.Fragment {
		// Keep links to page-local fragments relative.
		// Otherwise, we'll turn "#dogs" into "https://origin.com/page.html#dogs"
		// and send the user away when we could have kept them on the page they
		// already loaded for a better experience.
		//
		// This also handles the case where base == url, and neither has
		// a fragment. In which case we emit '#' which links to the top of the page.
		return "#" + u.Fragment
	}
	return uVal
}
