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
	"fmt"
	"math/big"
	"net/url"
	"sort"
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

// For images with aspect ratio, ignores images with aspect ratio larger than this.
const maxAspectRatioSize int = 16

// PreloadImage adds link rel="prefetch" head element to preload the most revalent image in the AMP document.
func PreloadImage(e *Context) error {
	for n := e.DOM.BodyNode; n != nil; n = htmlnode.Next(n) {
		if isNodeHiddenInLayout(n) {
			continue
		}

		if n.Data == "amp-img" {
			if imgsrcset, ok := candidateImageForPreloading(n); ok {
				srcsetToPreloadData(imgsrcset, e)
			}
		} else if n.Data == "amp-video" || n.Data == "amp-video-iframe" {
			if poster, ok := candidateVideoPosterImage(n); ok {
				posterURL, err := url.Parse(poster)
				if err != nil {
					continue
				}
				e.Preloads = append(e.Preloads, PreloadData{URL: posterURL, As: "image"})
			}
		}
	}
	return nil
}

func candidateVideoPosterImage(i *html.Node) (string, bool) {
	poster, hasPoster := htmlnode.GetAttributeVal(i, "", "poster")
	if !hasPoster || poster == "" {
		return "", false
	}

	videoWidth, videoHeight := nodeDimensions(i)
	if isTinyNode(videoWidth, videoHeight) {
		return "", false
	}

	return poster, true
}

func isNodeHiddenInLayout(n *html.Node) bool {
	return layout.ParseAMPLayout(n) == amppb.AmpLayout_NODISPLAY
}

// Converts the raw srcset attribute value and populates Context.Preloads field.
func srcsetToPreloadData(srcset string, e *Context) {
	type imageWithTargetSize struct {
		imgURL *url.URL
		size   int
	}

	srcSets := strings.FieldsFunc(strings.TrimSpace(srcset), func(c rune) bool { return c == ',' })
	srcSetsSize := len(srcSets)
	imgSet := []imageWithTargetSize{}

	for _, src := range srcSets {
		imgComponents := strings.Fields(src)
		if len(imgComponents) != 2 {
			e.Preloads = nil
			return
		}
		imgTargetSize, err := strconv.Atoi(strings.TrimSuffix(imgComponents[1], "w"))

		if err != nil {
			e.Preloads = nil
			return
		}

		urlObj, err := url.Parse(imgComponents[0])

		if err != nil {
			e.Preloads = nil
			return
		}

		imgSet = append(imgSet, imageWithTargetSize{urlObj, imgTargetSize})
	}

	// Sort the images based on their target sizes in asc order.
	sort.Slice(imgSet, func(i, j int) bool { return imgSet[i].size < imgSet[j].size })

	for i, ci := range imgSet {
		var mediaQuery string
		// srcset images should be sorted by width.
		if i == 0 {
			mediaQuery = fmt.Sprintf("(max-width: %d)", ci.size)
			// Largest image has only min width limit of second largest image.
		} else if i == srcSetsSize-1 {
			mediaQuery = fmt.Sprintf("(min-width: %d)", imgSet[i-1].size+1)
		} else {
			mediaQuery = fmt.Sprintf("(min-width: %d) and (max-width: %d)", imgSet[i-1].size+1, ci.size)
		}

		e.Preloads = append(e.Preloads, PreloadData{URL: ci.imgURL, As: "image", Media: mediaQuery})
	}
}

// Decides if the given image node qualifies for preloading and returns tuple of
// (imagesrc, true) if the node qualifies for preloading, otherwise returns
// empty string and false.
func candidateImageForPreloading(n *html.Node) (string, bool) {
	// amp-image under following containers do not qualify for preloading.
	imgsrcset, hasSrcset := htmlnode.GetAttributeVal(n, "", "srcset")

	// Ignores images with no src attribute.
	// These can be css images inside class definition.
	if !hasSrcset || len(imgsrcset) == 0 {
		return "", false
	}

	// Ignores if image src is not a https url.
	// URL rewrite transformer guarantees img srcs are https protocol.
	if !strings.HasPrefix(imgsrcset, "https://") {
		return "", false
	}

	widthInt, heightInt := nodeDimensions(n)

	// Ignores smaller images, unless they are aspect ratio dimensions.
	if isTinyNode(widthInt, heightInt) {
		// Checks for aspect ratio images.
		// Aspect ratio images larger than maxAspectRatioSize are ignored.
		// Small images of icon types inside input type container types
		// are ignored.
		if widthInt > 0 && widthInt <= maxAspectRatioSize && heightInt > 0 && heightInt <= maxAspectRatioSize && isAspectRatioDimensions(n, widthInt, heightInt) && !containerTypeInput(n) {
			return imgsrcset, true
		}
		return "", false
	}

	// Checks if it is placeholder image for iframe.
	// https://www.ampproject.org/docs/reference/components/amp-iframe#iframe-with-placeholder
	_, hasPlaceholder := htmlnode.GetAttributeVal(n, "", "placeholder")
	parentWidthInt, parentHeightInt := nodeDimensions(n.Parent)
	if hasPlaceholder {
		if n.Parent.Data == "amp-iframe" {
			if isTinyNode(parentWidthInt, parentHeightInt) {
				return "", false
			}
			return imgsrcset, true
		}
		return "", false
	}

	layoutType := layout.ParseAMPLayout(n)
	// Responsive and fill layout types generally accept parent containers dimensions.
	if layoutType == amppb.AmpLayout_RESPONSIVE || layoutType == amppb.AmpLayout_FILL {
		if widthInt == 0 && heightInt == 0 {
			if isTinyNode(parentWidthInt, parentHeightInt) {
				return "", false
			}
			return imgsrcset, true
		}

		// Actual image dimension check is performed later.
	}

	// For other layouts with no image dimensions, take parent containers
	// dimensions into account.
	if widthInt == 0 && heightInt == 0 {
		widthInt = parentWidthInt
		heightInt = parentHeightInt
	}

	// Checks image meets minimum dimension requirements.
	// Ignores the width size if it is not specified. In most layouts it
	// defaults to auto or 100% size of container.
	if (widthInt >= minImageSize || widthInt == 0) && heightInt >= minImageSize {
		return imgsrcset, true
	}

	return "", false
}

// Consider a small dimension size as aspect ratio if they are relatively prime.
// TODO(amaltas): Fix it for float dimension types: 1x1.33.
func isAspectRatioDimensions(n *html.Node, width int, height int) bool {
	// Aspect ratio doesn't work in fixed layout types.
	layoutType := layout.ParseAMPLayout(n)
	if !(layoutType == amppb.AmpLayout_FIXED || layoutType == amppb.AmpLayout_FIXED_HEIGHT) {
		return false
	}

	return new(big.Int).GCD(nil, nil, big.NewInt(int64(width)), big.NewInt(int64(height))).Int64() == 1
}

func containerTypeInput(i *html.Node) bool {
	switch i.Parent.Data {
	case
		"button",
		"input":
		return true
	}
	return false
}

// A tiny node is any container of amp-img that is smaller than 150x150.
// Node with no dimensions defaults to 0x0. A 0x0 node is not considered tiny as its dimensions
// defaults to its parent's dimensions.
// A node is small size only when width and height attribute are set and are positive value.
// Caller must check if the container's dimension are aspect ratio dimensions.
func isTinyNode(width, height int) bool {
	return (width > 0 && width < minImageSize) || (height > 0 && height < minImageSize)
}

func dimensionAsInt(d string) (int, error) {
	// Remove px suffix. Some publishers treat width/height attribute similar to CSS.
	replacer := strings.NewReplacer("px", "", "auto", "0")
	return strconv.Atoi(replacer.Replace(d))
}

func nodeDimensions(i *html.Node) (int, int) {
	var err error

	// Width and Height as int type.
	width, hasWidth := htmlnode.GetAttributeVal(i, "", "width")
	widthInt := 0
	if hasWidth {
		if widthInt, err = dimensionAsInt(width); err != nil {
			return 0, 0
		}
	}

	height, hasHeight := htmlnode.GetAttributeVal(i, "", "height")
	heightInt := 0
	if hasHeight {
		if heightInt, err = dimensionAsInt(height); err != nil {
			return 0, 0
		}
	}

	return widthInt, heightInt
}
