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

// URL operates on URL attributes, rewriting URLs as needed
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
func URL(e *Context) error {
	target := extractBaseTarget(e.DOM.HeadNode)

	for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
		// Skip text nodes
		if n.Type == html.TextNode {
			continue
		}

		// TODO(b/112417267): Handle amp-img rewriting.
		if strings.EqualFold(n.Data, "amp-img") {
			continue
		}

		// Make attributes with URLs portable on any tag
		rewritePortableURLs(n, e.BaseURL, anyTagAttrs)

		switch n.DataAtom {
		case atom.Form:
			// Make attributes with URLs absolute on <form> tag.
			rewriteAbsoluteURLs(n, e.BaseURL, formTagAttrs)
		case atom.Img:
			// Make attributes with URLs portable on <img> tag.
			rewritePortableURLs(n, e.BaseURL, imgTagAttrs)
		default:
			switch n.Data {
			case "amp-install-serviceworker":
				// Make attributes with URLs portable on <amp-install-serviceworker> tag.
				rewritePortableURLs(n, e.BaseURL, ampInstallServiceWorkerTagAttrs)
			case amphtml.AMPStory:
				// Make attributes with URLs portable on <amp-story> tag.
				rewritePortableURLs(n, e.BaseURL, ampStoryTagAttrs)
			case "amp-story-page":
				// Make attributes with URLs portable on <amp-story-page> tag.
				rewritePortableURLs(n, e.BaseURL, ampStoryPageTagAttrs)
			}
		}

		// Tags with href attribute.
		if href, ok := htmlnode.FindAttribute(n, "", "href"); ok {
			in := htmlnode.IsDescendantOf(n, atom.Template)

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
			switch n.DataAtom {
			case atom.Base:
				htmlnode.RemoveAttribute(n, href)
				if len(n.Attr) == 0 {
					htmlnode.RemoveNode(&n)
				}
			case atom.Link:
				if v, ok := htmlnode.GetAttributeVal(n, "rel"); ok && v == "canonical" {
					// If the origin doc is self-canonical, it should be an absolute URL
					// and not portable (which would result in canonical = "#").
					// Maintain the original canonical, and absolutify it. See b/36102624
					htmlnode.SetAttribute(n, "", "href", amphtml.RewriteAbsoluteURL(e.BaseURL, in, href.Val))
				} else {
					htmlnode.SetAttribute(n, "", "href", amphtml.RewritePortableURL(e.BaseURL, in, href.Val))
				}
			case atom.A:
				portableHref := amphtml.RewritePortableURL(e.BaseURL, in, href.Val)
				// Set a default target
				// 1. If the href is not a fragment AND
				// 2. If there is no target OR If there is a target and it is not an allowed target
				if !strings.HasPrefix(portableHref, "#") {
					if v, ok := htmlnode.GetAttributeVal(n, "target"); !ok || (ok && !isAllowedTarget(v)) {
						htmlnode.SetAttribute(n, "", "target", target)
					}
				}
				htmlnode.SetAttribute(n, "", "href", portableHref)
			default:
				// Make a PortableUrl for any remaining tags with href.
				htmlnode.SetAttribute(n, "", "href", amphtml.RewritePortableURL(e.BaseURL, in, href.Val))
			}
		}
	}
	return nil
}

// extractBaseTarget returns the target value derived from the <base> tag, if it exists,
// and is allowed. Otherwise, returns "_top".
func extractBaseTarget(head *html.Node) string {
	if n, ok := htmlnode.FindNode(head, atom.Base); ok {
		if v, ok := htmlnode.GetAttributeVal(n, "target"); ok && isAllowedTarget(v) {
			return v
		}
	}
	return "_top"
}

// isAllowedTarget returns true if the given string is either "_blank" or "_top"
func isAllowedTarget(t string) bool {
	return strings.EqualFold(t, "_blank") || strings.EqualFold(t, "_top")
}

// rewriteAbsoluteURLs rewrites URLs in the given slice of attributes
// to be absolute for the base URL provided.
func rewriteAbsoluteURLs(n *html.Node, base *url.URL, tagAttrs []string) {
	if htmlnode.IsDescendantOf(n, atom.Template) {
		return
	}
	for _, attr := range tagAttrs {
		if v, ok := htmlnode.GetAttributeVal(n, attr); ok {
			htmlnode.SetAttribute(n, "", attr, amphtml.RewriteAbsoluteURL(base, false, v))
		}
	}
}

// rewritePortableURLs rewrites URLs in the given slice of attributes
// to be portable relative to the base URL provided.
func rewritePortableURLs(n *html.Node, base *url.URL, tagAttrs []string) {
	if htmlnode.IsDescendantOf(n, atom.Template) {
		return
	}
	for _, attr := range tagAttrs {
		if v, ok := htmlnode.GetAttributeVal(n, attr); ok {
			htmlnode.SetAttribute(n, "", attr, amphtml.RewritePortableURL(base, false, v))
		}
	}
}
