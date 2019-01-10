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
	"regexp"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

// eventRE is a regexp to match event attributes, e.g. onClick
var eventRE = func() *regexp.Regexp {
	r := regexp.MustCompile("^on[A-Za-z].*")
	return r
}()

// StripJS removes non-AMP javascript from the DOM.
// - For <script> elements, remove where any of the following is true:
//     - has a src attribute whose value is not prefixed by https://cdn.ampproject.org/ (case-insensitive match).
//     - It has no src attribute and no type attribute (case-insensitive match).
//     - It has a type attribute whose value is neither application/json nor application/ld+json (case-insensitive match on both name and value).
//
// - For all other elements, remove any event attribute that matches "on[A-Za-z].*".
func StripJS(e *Context) error {
	for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
		if n.Type != html.ElementNode {
			continue
		}

		if n.DataAtom == atom.Script {
			srcVal, srcOk := htmlnode.GetAttributeVal(n, "", "src")
			var isCacheSrc bool
			if srcOk {
				if !strings.HasPrefix(strings.ToLower(srcVal), amphtml.AMPCacheRootURL) {
					htmlnode.RemoveNode(&n)
					continue
				}
				isCacheSrc = true
			}
			typeVal, typeOk := htmlnode.GetAttributeVal(n, "", "type")
			if !srcOk && !typeOk {
				htmlnode.RemoveNode(&n)
				continue
			}
			if typeOk {
				switch strings.ToLower(typeVal) {
				case "application/json", "application/ld+json":
					// ok to keep
				case "text/javascript":
					// ok to keep only for AMP Cache scripts.
					if !isCacheSrc {
						htmlnode.RemoveNode(&n)
					}
				default:
					htmlnode.RemoveNode(&n)
				}
			}
		} else {
			for _, attr := range n.Attr {
				if attr.Namespace == "" {
					if match := eventRE.MatchString(attr.Key); match {
						htmlnode.RemoveAttribute(n, &attr)
					}
				}
			}
		}
	}
	return nil
}
