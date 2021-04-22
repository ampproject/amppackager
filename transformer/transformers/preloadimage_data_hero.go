// Copyright 2020 Google LLC
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
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html"
)

// preloadImageDataHero finds appropriate AMP elements that have opted in with a `data-hero`
// attribute. There are no size restrictions for these elements.
func preloadImageDataHero(n *html.Node) (HeroImage, bool, *html.Node) {
	for n != nil {
		if n.Data == "amp-img" {
			next := htmlnode.NextSkippingChildren(n)
			if heroImage, ok := dataHeroImageForPreloading(n); ok {
				return heroImage, true, next
			}
			n = next
			continue
		}

		if n.Data == "amp-video" || n.Data == "amp-video-iframe" {
			next := htmlnode.NextSkippingChildren(n)
			if heroImage, ok := dataHeroVideoPosterImage(n); ok {
				return heroImage, true, next
			}
			if heroImage, ok := dataHeroWithPlaceholderImage(n); ok {
				return heroImage, true, next
			}
			n = next
			continue
		}

		if n.Data == "amp-iframe" {
			next := htmlnode.NextSkippingChildren(n)
			if heroImage, ok := dataHeroWithPlaceholderImage(n); ok {
				return heroImage, true, next
			}
			n = next
			continue
		}

		if n.Data == "template" {
			n = htmlnode.NextSkippingChildren(n)
			continue
		}

		n = htmlnode.Next(n)
	}

	return HeroImage{}, false, nil
}

// For a given <amp-video> node or any node that has poster attribute, and
// qualifies as hero image, returns the HeroImage.
func dataHeroVideoPosterImage(i *html.Node) (HeroImage, bool) {
	if !htmlnode.HasAttribute(i, "", "data-hero") {
		return HeroImage{}, false
	}

	poster, hasPoster := ValidateSrc(htmlnode.GetAttributeVal(i, "", "poster"))
	if !hasPoster {
		return HeroImage{}, false
	}

	return HeroImage{
		src:    poster,
		srcset: "",
		ampImg: nil,
	}, true
}

// For a given node returns a placeholder image if the placeholder qualifies
// for hero image.
func dataHeroWithPlaceholderImage(n *html.Node) (HeroImage, bool) {
	if !htmlnode.HasAttribute(n, "", "data-hero") {
		return HeroImage{}, false
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data != "amp-img" || !htmlnode.HasAttribute(c, "", "placeholder") {
			continue
		}

		layout, hasLayout := htmlnode.GetAttributeVal(c, "", "layout")
		if !hasLayout || layout != "fill" {
			continue
		}

		src, hasSrc := ValidateSrc(htmlnode.GetAttributeVal(c, "", "src"))
		srcset, hasSrcset := ParseAndValidateSrcset(htmlnode.GetAttributeVal(c, "", "srcset"))
		if hasSrc || hasSrcset {
			return HeroImage{
				src:    src,
				srcset: srcset,
				ampImg: c,
			}, true
		}
	}

	return HeroImage{}, false
}

// Checks if amp-img qualifies to be a hero image. Returns HeroImage if the
// node is a hero image.
func dataHeroImageForPreloading(n *html.Node) (HeroImage, bool) {
	if !htmlnode.HasAttribute(n, "", "data-hero") {
		return HeroImage{}, false
	}

	src, hasSrc := ValidateSrc(htmlnode.GetAttributeVal(n, "", "src"))
	srcset, hasSrcset := ParseAndValidateSrcset(htmlnode.GetAttributeVal(n, "", "srcset"))

	// Ignores images with no src attribute.
	if !hasSrc && !hasSrcset {
		return HeroImage{}, false
	}

	return HeroImage{
		src:    src,
		srcset: srcset,
		ampImg: n,
	}, true
}
