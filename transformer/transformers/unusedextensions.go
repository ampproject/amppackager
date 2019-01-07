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
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html"
)

// elementGrandfatheredExtensions are names of elements that indicate usage
// of an equally named extension. e.g. If the <amp-iframe> element is present,
// then the amp-iframe extension is in use. Used by insertMatchingExtensions.
var /* const */ elementGrandfatheredExtensions = map[string]string{"amp-accordion": "", "amp-ad": "", "amp-anim": "", "amp-apester-media": "", "amp-audio": "", "amp-brid-player": "", "amp-brightcove": "", "amp-call-tracking": "", "amp-carousel": "", "amp-dailymotion": "", "amp-experiment": "", "amp-facebook": "", "amp-fit-text": "", "amp-font": "", "amp-fx-flying-carpet": "", "amp-gfycat": "", "amp-iframe": "", "amp-image-lightbox": "", "amp-instagram": "", "amp-install-serviceworker": "", "amp-izlesene": "", "amp-jwplayer": "", "amp-kaltura-player": "", "amp-lightbox": "", "amp-list": "", "amp-live-list": "", "amp-o2-player": "", "amp-pinterest": "", "amp-reach-player": "", "amp-selector": "", "amp-sidebar": "", "amp-social-share": "", "amp-soundcloud": "", "amp-springboard-player": "", "amp-sticky-ad": "", "amp-twitter": "", "amp-user-notification": "", "amp-vimeo": "", "amp-vine": "", "amp-youtube": ""}

// differentElementGrandfatheredExtensions are names of extensions that indicate
// usage by a differently named tag or tag with attribute. e.g. If the <form>
// element is present, then the amp-form extension is in use.
var /* const */ differentElementGrandfatheredExtensions = map[string]string{"amp-access": "", "amp-form": "", "amp-mustache": ""}

// UnusedExtensions removes script tags for unused grandfathered extensions.
func UnusedExtensions(e *Context) error {
	extensionsUsed := make(map[string]string)
	for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
		insertMatchingExtensions(n, extensionsUsed)
	}
	for c := e.DOM.HeadNode.FirstChild; c != nil; c = c.NextSibling {
		if amphtml.IsScriptAMPExtension(c) {
			var ext string
			if v, ok := htmlnode.GetAttributeVal(c, "", amphtml.AMPCustomElement); ok {
				ext = v
			} else if v, ok := htmlnode.GetAttributeVal(c, "", amphtml.AMPCustomTemplate); ok {
				ext = v
			}
			if len(ext) > 0 && (isStringKeyInMap(ext, elementGrandfatheredExtensions) || isStringKeyInMap(ext, differentElementGrandfatheredExtensions)) && !isStringKeyInMap(ext, extensionsUsed) {
				htmlnode.RemoveNode(&c)
			}
		}
	}
	return nil
}

// insertMatchingExtensions inserts all extensions that might be activated
// by the inclusion of this element. It's okay if it has false positives
// (that just means we won't be as aggressive about removing their script
// tags), but it's not okay if it has false negatives (we may remove an
// extension that's needed).
//
// This logic should match the requires_extension fields in the
// validator-*.protoascii files that correspond to GRANDFATHERED
// extension_specs.
func insertMatchingExtensions(n *html.Node, e map[string]string) {
	if n.Type != html.ElementNode {
		return
	}
	switch n.Data {
	case "script":
		if v, ok := htmlnode.GetAttributeVal(n, "", "id"); ok && strings.EqualFold(v, "amp-access") {
			e["amp-access"] = ""
		}
	case "amp-embed":
		e["amp-ad"] = ""
	case "form":
		e["amp-form"] = ""
	case "template":
		e["amp-mustache"] = ""
	default:
		if _, ok := elementGrandfatheredExtensions[n.Data]; ok {
			e[n.Data] = ""
		}
	}
	return
}

// isStringKeyInMap returns true if the given string is a key in the given map.
func isStringKeyInMap(s string, m map[string]string) bool {
	_, ok := m[s]
	return ok
}
