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

var /* const */ anyTagAttrs = []string{"background", "poster", "src"}
var /* const */ ampInstallServiceWorkerTagAttrs = []string{"data-iframe-src", "data-no-service-worker-fallback-shell-url"}
var /* const */ ampStoryTagAttrs = []string{"background-audio", "bookend-config-src", "poster-landscape-src", "poster-square-src", "publisher-logo-src"}
var /* const */ ampStoryPageTagAttrs = []string{"background-audio"}
var /* const */ formTagAttrs = []string{"action", "action-xhr"}
var /* const */ imgTagAttrs = []string{"longdesc"}

// AbsoluteURL operates on URL attributes. It rewrites URLs as Absolute URLs.
// These are based on the URL in the attribute and the base url of the
// document. A base URL is the final URL of the document after all redirects.
// If the attribute URL is relative, then it is relative to the base_url.
// There is special handling for URLs that contain fragments.
//
// URLs are also canonicalized:
// * leading and trailing whitespace are trimmed.
//
// * The following attributes may be rewritten:
//   * Any tag with attribute:
//     * background
//     * href
//     * poster
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
// URLs in stylesheets and srcsets are handled by the ExternalUrlRewrite
// transformer.
//
func AbsoluteURL(e *Context) error {
	target := extractBaseTarget(e.DOM.HeadNode)
	documentURL := e.DocumentURL.String()

	for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
		// Skip text nodes and anything inside mustache templates
		if n.Type == html.TextNode || htmlnode.IsDescendantOf(n, atom.Template) {
			continue
		}

		// Make attributes with URLs portable on any tag
		rewriteAbsoluteURLs(n, documentURL, e.BaseURL, anyTagAttrs)

		switch n.DataAtom {
		case atom.Form:
			// Make attributes with URLs absolute on <form> tag.
			rewriteAbsoluteURLs(n, documentURL, e.BaseURL, formTagAttrs)
		case atom.Img:
			// Make attributes with URLs portable on <img> tag.
			rewriteAbsoluteURLs(n, documentURL, e.BaseURL, imgTagAttrs)
		default:
			switch n.Data {
			case "amp-install-serviceworker":
				// Make attributes with URLs portable on <amp-install-serviceworker> tag.
				rewriteAbsoluteURLs(n, documentURL, e.BaseURL, ampInstallServiceWorkerTagAttrs)
			case amphtml.AMPStory:
				// Make attributes with URLs portable on <amp-story> tag.
				rewriteAbsoluteURLs(n, documentURL, e.BaseURL, ampStoryTagAttrs)
			case "amp-story-page":
				// Make attributes with URLs portable on <amp-story-page> tag.
				rewriteAbsoluteURLs(n, documentURL, e.BaseURL, ampStoryPageTagAttrs)
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
			switch n.DataAtom {
			case atom.Base:
				htmlnode.RemoveAttribute(n, href)
				if len(n.Attr) == 0 {
					htmlnode.RemoveNode(&n)
				}
			case atom.Link:
				if v, ok := htmlnode.GetAttributeVal(n, "", "rel"); ok && v == "canonical" {
					// Some users set <link rel=canonical href=#> which is silly, but
					// can lead the cache to leave this as "#" which means that canonical
					// is not preserved correctly. Fragments don't make any sense for
					// canonical, so we handle this case specially.
					//
					// Protocol relative is also silly for canonical but we don't want to
					// try to second guess the publisher here.
					absoluteURL, err := e.BaseURL.Parse(href.Val)
					if err == nil {
						absoluteURL.Fragment = ""
						htmlnode.SetAttribute(n, "", "href", absoluteURL.String())
					}
				} else {
					htmlnode.SetAttribute(n, "", "href",
						amphtml.ToAbsoluteURL(documentURL, e.BaseURL, href.Val))
				}
			case atom.A:
				newValue := amphtml.ToAbsoluteURL(documentURL, e.BaseURL, href.Val)
				// Set a default target
				// 1. If the href is not a fragment AND
				// 2. If there is no target OR
				// 3. If there is a target and it is not an allowed target
				if !strings.HasPrefix(newValue, "#") {
					if v, ok := htmlnode.GetAttributeVal(n, "", "target"); !ok || (ok && !isAllowedTarget(v)) {
						htmlnode.SetAttribute(n, "", "target", target)
					}
				}
				htmlnode.SetAttribute(n, "", "href", newValue)
			default:
				// Absoluteify any remaining tags with an href attribute.
				htmlnode.SetAttribute(n, "", "href",
					amphtml.ToAbsoluteURL(documentURL, e.BaseURL, href.Val))
			}
		}
		if _, ok := htmlnode.FindAttribute(n, "", "srcset"); ok {
			rewriteSrcsetURLs(n, documentURL, e.BaseURL)
		}
	}
	return nil
}

// extractBaseTarget returns the target value derived from the <base> tag, if it exists,
// and is allowed. Otherwise, returns "_top".
func extractBaseTarget(head *html.Node) string {
	if n, ok := htmlnode.FindNode(head, atom.Base); ok {
		if v, ok := htmlnode.GetAttributeVal(n, "", "target"); ok && isAllowedTarget(v) {
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
func rewriteAbsoluteURLs(n *html.Node, documentURL string, baseURL *url.URL,
	tagAttrs []string) {
	for _, attr := range tagAttrs {
		if v, ok := htmlnode.GetAttributeVal(n, "", attr); ok {
			htmlnode.SetAttribute(n, "", attr,
				amphtml.ToAbsoluteURL(documentURL, baseURL, v))
		}
	}
}

func rewriteSrcsetURLs(n *html.Node, documentURL string, baseURL *url.URL) {
	if v, ok := htmlnode.GetAttributeVal(n, "", "srcset"); ok {
		normalized, offsets := amphtml.ParseSrcset(v)
		var sb strings.Builder
		var pos int
		for _, element := range offsets {
			sb.WriteString(normalized[pos:element.Start])
			sb.WriteString(amphtml.ToAbsoluteURL(
				documentURL, baseURL, normalized[element.Start:element.End]))
			pos = element.End
		}
		sb.WriteString(normalized[pos:])
		htmlnode.SetAttribute(n, "", "srcset", sb.String())
	}
}
