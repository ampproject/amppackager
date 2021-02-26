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

	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

// LinkTag operates on the <link> tag.
// * It will add a preconnect link tag for Google Font resources.
// * It will add a preconnect link tag to the publisher's own origin.
// * It will add "display=optional" to Google Fonts without "display" component.
func LinkTag(e *Context) error {
	preconnectAdded := false

	for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
		if !preconnectAdded && isLinkAnyGoogleFont(n) {
			addLinkGoogleFontPreconnect(n)
			preconnectAdded = true
		}
		if isLinkGoogleFont(n) {
			addDisplayOptional(n)
		}
	}

	addLinkPublisherOriginPreconnect(e.DOM.HeadNode, e.DocumentURL)
	return nil
}

// isLinkAnyGoogleFont returns true if the given node is a link tag, has attribute href with the Google Font root URL.
func isLinkAnyGoogleFont(n *html.Node) bool {
	if n.DataAtom != atom.Link {
		return false
	}
	v, ok := htmlnode.GetAttributeVal(n, "", "href")
	return ok && strings.HasPrefix(v, "https://fonts.googleapis.com/")
}

// isLinkGoogleFont returns true if the given node is a link tag, has attribute
// href with the Google Font root URL and path starting "css?" or "css2?".
func isLinkGoogleFont(n *html.Node) bool {
	if n.DataAtom != atom.Link {
		return false
	}
	v, ok := htmlnode.GetAttributeVal(n, "", "href")
	return ok && (strings.HasPrefix(v, "https://fonts.googleapis.com/css?") || strings.HasPrefix(v, "https://fonts.googleapis.com/css2?"))
}

// addDisplayOptional adds "display=optional" to Google Font link tag's href attribute.
func addDisplayOptional(n *html.Node) {
	// Confirm this is a Google Font
	if !isLinkGoogleFont(n) {
		return
	}
	if hrefVal, hrefOk := htmlnode.GetAttributeVal(n, "", "href"); hrefOk {
		if u, _ := url.Parse(strings.ToLower(hrefVal)); u != nil {
			v := u.Query()
			if g := v.Get("display"); g == "" {
				htmlnode.AppendAttribute(n, "", "href", "&display=optional")
			}
		}
	}
}

// addLinkGoogleFontPreconnect adds a preconnect link tag for Google Font resources.
func addLinkGoogleFontPreconnect(n *html.Node) {
	if n.DataAtom != atom.Link {
		return
	}
	preconnect :=
		htmlnode.Element("link",
			html.Attribute{Key: "crossorigin"},
			html.Attribute{Key: "href", Val: "https://fonts.gstatic.com/"},
			html.Attribute{Key: "rel", Val: "dns-prefetch preconnect"})
	n.Parent.InsertBefore(preconnect, n)
}

// addLinkPublisherOriginPreconnect adds a preconnect link tag for the
// publisher's own origin. This will only occur once the SXG is fully loaded
// so does not invalidate privacy preserving preload. For publishers that load
// dynamic resources, this will speed up those requests.
func addLinkPublisherOriginPreconnect(n *html.Node, u *url.URL) {
	if n.DataAtom != atom.Head {
		return
	}
	// Generates a preconnect value, which does not need anything
	// other than the origin to connect to, so to shave some bytes, strip
	// everything else.
	urlCopy := *u
	urlCopy.User = nil
	urlCopy.Path = ""
	urlCopy.ForceQuery = false
	urlCopy.RawQuery = ""
	urlCopy.Fragment = ""

	preconnect :=
		htmlnode.Element("link",
			html.Attribute{Key: "href", Val: urlCopy.String()},
			html.Attribute{Key: "rel", Val: "dns-prefetch preconnect"})
	n.AppendChild(preconnect)
}
