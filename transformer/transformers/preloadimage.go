// Copyright 2019 Google LLC
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

const maxHeroImages int = 2

// A map which translates <amp-img> attributes (keys) to <link rel=preload> attributes (values).
// Any HeroImage which has a <amp-img> node will also inherit these attribute values.
var preloadAttributes = map[string]string{
	"sizes":          "imagesizes",
	"crossorigin":    "crossorigin",
	"referrerpolicy": "referrerpolicy",
}

// HeroImage represents the necessary data to inject a <link ref=preload> and optional <img> tag.
type HeroImage struct {
	src    string
	srcset string
	ampImg *html.Node
}

// PreloadImage adds link rel="preload" to head element to preload the most revalent image in the AMP document,
// and inserts an img tag if the image is an amp-img.
func PreloadImage(e *Context) error {
	body := e.DOM.BodyNode
	current := body
	count := 0
	for i := 0; i < maxHeroImages; i++ {
		heroImage, found, next := preloadImageDataHero(current)
		if !found {
			break
		}
		count++
		current = next
		prioritizeHeroImage(e, heroImage)
	}

	// If any elements were opted-in, then we do not need to infer a hero image.
	if count == 0 {
		heroImage, found := preloadImageInferSize(body)
		if found {
			prioritizeHeroImage(e, heroImage)
		}
	}

	// Finally, inject a loading=lazy img for all remaining amp-img elements.
	if e.Version >= 5 {
		lazyLoadRemainingAmpImgs(body)
	}

	return nil
}

func prioritizeHeroImage(e *Context, heroImage HeroImage) {
	if heroImage.src != "" || heroImage.srcset != "" {
		link := htmlnode.Element("link",
			html.Attribute{Key: "rel", Val: "preload"},
			html.Attribute{Key: "as", Val: "image"},
		)
		if heroImage.src != "" {
			htmlnode.SetAttribute(link, "", "href", heroImage.src)
		}
		if heroImage.srcset != "" {
			htmlnode.SetAttribute(link, "", "imagesrcset", heroImage.srcset)
		}
		if ampImg := heroImage.ampImg; ampImg != nil {
			for name, linkName := range preloadAttributes {
				if value, ok := htmlnode.GetAttributeVal(ampImg, "", name); ok {
					htmlnode.SetAttribute(link, "", linkName, value)
				}
			}
		}
		e.DOM.HeadNode.AppendChild(link)
	}

	if ampImg := heroImage.ampImg; ampImg != nil {
		ampImg.AppendChild(buildImg(ampImg))
		htmlnode.SetAttribute(ampImg, "", "i-amphtml-ssr", "")
	}
}

func buildImg(ampImg *html.Node) *html.Node {
	img := htmlnode.Element("img",
		html.Attribute{Key: "class", Val: "i-amphtml-fill-content i-amphtml-replaced-content"},
		html.Attribute{Key: "decoding", Val: "async"})
	attrsToCopy := [...]string{
		"alt",
		"attribution",
		"crossorigin",
		"object-fit",
		"object-position",
		"referrerpolicy",
		"src",
		"srcset",
		"sizes",
		"title",
	}

	for _, attr := range attrsToCopy {
		val, has := htmlnode.GetAttributeVal(ampImg, "", attr)
		if has {
			htmlnode.SetAttribute(img, "", attr, val)
		}
	}

	return img
}

// ParseAndValidateSrcset parses each source in the srcset, ensuring it points to a HTTPS URL, and normalizes the srcset.
func ParseAndValidateSrcset(in string, has bool) (string, bool) {
	if !has {
		return "", false
	}

	parsed, offsets := amphtml.ParseSrcset(in)
	if len(offsets) == 0 {
		return "", false
	}

	for _, offset := range offsets {
		if !isValidURL(parsed[offset.Start:offset.End]) {
			return "", false
		}
	}

	return parsed, true
}

func isValidURL(url string) bool {
	return strings.HasPrefix(url, "https:")
}

// ValidateSrc ensures the src points to a HTTPS URL.
func ValidateSrc(in string, has bool) (string, bool) {
	if !has || !isValidURL(in) {
		return "", false
	}
	return in, true
}

func lazyLoadRemainingAmpImgs(n *html.Node) {
	for n != nil {
		if n.Data == "amp-img" {
			// If the amp-img already has the ssr attribute, then it's a hero image with an already injected img.
			if !htmlnode.HasAttribute(n, "", "i-amphtml-ssr") {
				img := buildImg(n)
				htmlnode.SetAttribute(img, "", "loading", "lazy")
				n.AppendChild(img)
				htmlnode.SetAttribute(n, "", "i-amphtml-ssr", "")
			}

			n = htmlnode.NextSkippingChildren(n)
		} else if n.Data == "template" {
			n = htmlnode.NextSkippingChildren(n)
		} else {
			n = htmlnode.Next(n)
		}
	}
}
