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
	"strconv"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"github.com/ampproject/amppackager/transformer/layout"
	"golang.org/x/net/html"
	amppb "github.com/ampproject/amphtml/validator"
)

// Images smaller than 150 pixels are ignored for preloading.
// This number is chosen after manually reviewing 10k samples.
const minImageSize int = 150

// preloadImageInferSize infers an appropriate AMP element to use as a hero image. This requires the
// element have defined size of at least 150px.
func preloadImageInferSize(n *html.Node) (HeroImage, bool) {
	for n != nil {
		if isNodeHiddenInLayout(n) {
			n = htmlnode.NextSkippingChildren(n)
			continue
		}

		if n.Data == "amp-img" {
			if heroImage, ok := inferredSizeImageForPreloading(n); ok {
				return heroImage, true
			}
			n = htmlnode.NextSkippingChildren(n)
			continue
		}

		if n.Data == "amp-video" || n.Data == "amp-video-iframe" {
			if heroImage, ok := inferredSizeVideoPosterImage(n); ok {
				return heroImage, true
			}
			if heroImage, ok := inferredSizeWithPlaceholderImage(n); ok {
				return heroImage, true
			}
			n = htmlnode.NextSkippingChildren(n)
			continue
		}

		if n.Data == "amp-iframe" {
			if heroImage, ok := inferredSizeWithPlaceholderImage(n); ok {
				return heroImage, true
			}
			n = htmlnode.NextSkippingChildren(n)
			continue
		}

		if n.Data == "template" {
			n = htmlnode.NextSkippingChildren(n)
			continue
		}

		n = htmlnode.Next(n)
	}

	return HeroImage{}, false
}

func isNodeHiddenInLayout(n *html.Node) bool {
	return layout.ParseAMPLayout(n) == amppb.AmpLayout_NODISPLAY
}

// For a given <amp-video> node or any node that has poster attribute, and
// qualifies as hero image, returns the HeroImage.
func inferredSizeVideoPosterImage(i *html.Node) (HeroImage, bool) {
	poster, hasPoster := ValidateSrc(htmlnode.GetAttributeVal(i, "", "poster"))
	if !hasPoster {
		return HeroImage{}, false
	}

	videoWidth, _, videoHeight, _ := nodeDimensions(i)
	if isTinyNode(videoWidth, videoHeight) {
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
func inferredSizeWithPlaceholderImage(n *html.Node) (HeroImage, bool) {
	if n.FirstChild == nil {
		return HeroImage{}, false
	}

	width, _, height, _ := nodeDimensions(n)
	if isTinyNode(width, height) {
		return HeroImage{}, false
	}

	for c := n.FirstChild; n != nil; n = n.NextSibling {
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
func inferredSizeImageForPreloading(n *html.Node) (HeroImage, bool) {
	// amp-image under following containers do not qualify for preloading.
	src, hasSrc := ValidateSrc(htmlnode.GetAttributeVal(n, "", "src"))
	srcset, hasSrcset := ParseAndValidateSrcset(htmlnode.GetAttributeVal(n, "", "srcset"))

	// Ignores images with no src attribute.
	if !hasSrc && !hasSrcset {
		return HeroImage{}, false
	}

	layoutType := layout.ParseAMPLayout(n)
	width, hasWidth, height, hasHeight := nodeDimensions(n)

	if !hasWidth && !hasHeight {
		// Responsive and fill layout types generally accept parent containers dimensions.
		if layoutType == amppb.AmpLayout_RESPONSIVE || layoutType == amppb.AmpLayout_FILL {
			width, hasWidth, height, hasHeight = nodeDimensionsFromParent(n)
		}
	}

	// Ignores smaller images, unless they are aspect ratio dimensions.
	if !(hasWidth && hasHeight) || isTinyNode(width, height) {
		return HeroImage{}, false
	}

	return HeroImage{
		src:    src,
		srcset: srcset,
		ampImg: n,
	}, true
}

// A tiny node is any container of amp-img that is smaller than 150x150.
// A node is small size only when width and height attribute are set and are positive value.
// Caller must check if the container's dimension are aspect ratio dimensions.
func isTinyNode(width, height int) bool {
	return width < minImageSize || height < minImageSize
}

func dimensionAsInt(d string) (int, error) {
	// Remove px suffix. Some publishers treat width/height attribute similar to CSS.
	replacer := strings.NewReplacer("px", "", "auto", "0")
	return strconv.Atoi(replacer.Replace(d))
}

func nodeDimensions(i *html.Node) (int, bool, int, bool) {
	var err error

	// Width and Height as int type.
	width, hasWidth := htmlnode.GetAttributeVal(i, "", "width")
	height, hasHeight := htmlnode.GetAttributeVal(i, "", "height")
	widthInt := 0
	heightInt := 0

	if hasWidth {
		if widthInt, err = dimensionAsInt(width); err != nil {
			return 0, hasWidth, 0, hasHeight
		}
	}

	if hasHeight {
		if heightInt, err = dimensionAsInt(height); err != nil {
			return 0, hasWidth, 0, hasHeight
		}
	}

	return widthInt, hasWidth, heightInt, hasHeight
}

// Given a node, determines its dimensions based on parent node's dimensions.
// Used when a node has no width/height attribute. This doesn't check layout,
// assumes node's layout supports inheriting parent's dimensions, like
// responsive or fill.
func nodeDimensionsFromParent(n *html.Node) (int, bool, int, bool) {
	for n.Parent != nil {
		n = n.Parent

		width, hasWidth, height, hasHeight := nodeDimensions(n)
		if hasWidth || hasHeight {
			return width, hasWidth, height, hasHeight
		}
	}

	return 0, false, 0, false
}
