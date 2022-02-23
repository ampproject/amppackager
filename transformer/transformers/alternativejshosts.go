// Copyright 2022 Google LLC
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

// AlternativeJsHosts AMP pages can choose to host the AMP Javascript on their own
// domain names, or even a third-party. One of these is recommended by the OpenJS
// Foundation - ampjs.org. Regardless of the domain name used on the publisher origin,
// when we serve the document on Google's Cache, we must rewrite it first to use
// cdn.ampproject.org.
func AlternativeJsHosts(e *Context) error {
	for c := e.DOM.HeadNode.FirstChild; c != nil; c = htmlnode.Next(c) {
		if isAsyncScriptTag(c) {
			// We need to rewrite the src URL.
			if src, ok := htmlnode.FindAttribute(c, "", "src"); ok {
				if parsedURL, err := url.Parse(src.Val); err != nil {
					htmlnode.RemoveAttribute(c, src)
				} else {
					// Handle /v0.js and /v0.mjs, as well as the ads equivalents.
					if strings.HasSuffix(parsedURL.Path, "/v0.js") {
						htmlnode.SetAttribute(c, "", "src", "https://cdn.ampproject.org/v0.js")
					} else if strings.HasSuffix(parsedURL.Path, "/v0.mjs") {
						htmlnode.SetAttribute(c, "", "src", "https://cdn.ampproject.org/v0.mjs")
					} else if strings.HasSuffix(parsedURL.Path, "/amp4ads-v0.js") {
						htmlnode.SetAttribute(c, "", "src", "https://cdn.ampproject.org/amp4ads-v0.js")
					} else if strings.HasSuffix(parsedURL.Path, "/amp4ads-v0.mjs") {
						htmlnode.SetAttribute(c, "", "src", "https://cdn.ampproject.org/amp4ads-v0.mjs")
					} else {
						// Handle all of the various component script tags:
						if endPart, ok := ExtractScriptNameAndVersionPart(parsedURL.Path); ok {
							htmlnode.SetAttribute(c, "", "src", "https://cdn.ampproject.org/v0/"+endPart)
						}
					}
				}
			}
		}
	}
	return nil
}

// ExtractScriptNameAndVersionPart extracts the suffix string like
// "/v0/amp-analytics-0.1.js" or false on error.
func ExtractScriptNameAndVersionPart(urlPath string) (string, bool) {
	// url_path should look something like "/v0/amp-analytics-0.1.js". If it
	// doesn't, we have an issue.
	if !strings.HasPrefix(urlPath, "/v0/amp-") {
		return "", false
	}
	if !strings.HasSuffix(urlPath, ".js") && !strings.HasSuffix(urlPath, ".mjs") {
		return "", false
	}

	// If everything is correct, this part will look like some variant of
	// "amp-analytics-0.1.js".
	endPart := urlPath[len("/v0/"):]

	// We constrain it to certain characters. This isn't perfect. The string
	// could be "foo-bar.js" and we'd accept it.
	for _, c := range endPart {
		// Strict. We only accept lower case:
		if c >= 'a' && c <= 'z' {
			continue
		}
		if c >= '0' && c <= '9' {
			continue
		}
		if c == '-' || c == '.' {
			continue
		}
		// Unexpected character:
		return "", false
	}
	return endPart, true
}

func isAsyncScriptTag(n *html.Node) bool {
	return (n.DataAtom == atom.Script &&
		htmlnode.HasAttribute(n, "", "async") &&
		htmlnode.HasAttribute(n, "", "src"))
}
