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
	"math/big"
	"strconv"
	"strings"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html"
)

// Transforms document to add prefetch link for most prevalent image.
func PreloadImage(e *Context) error {
	important_image_urls := imagesToPreload(e)
	if len(important_image_urls) == 0 {
		return nil
	}

	for _, imgurl := range important_image_urls {
		prefetch_link := htmlnode.Element("link", html.Attribute{Key: "rel", Val: "prefetch"}, html.Attribute{Key: "href", Val: imgurl})
		e.DOM.HeadNode.AppendChild(prefetch_link)
		// For v0, use the first candidate image.
		break
	}

	return nil
}

// Returns list of images on the document that qualifies for preloading.
func imagesToPreload(e *Context) []string {
	candidate_images := []string{}

	for n := e.DOM.BodyNode; n != nil; n = htmlnode.Next(n) {
		if n.Data == "amp-img" {
			if yes, imgsrc := candidateImageForPreloading(n); yes {
				candidate_images = append(candidate_images, imgsrc)
			}
		} else if n.Data == "amp-video" || n.Data == "amp-video-iframe" {
			if yes, imgsrc := candidateVideoPosterImage(n); yes {
				candidate_images = append(candidate_images, imgsrc)
			}
		}
	}
	return candidate_images
}

func candidateVideoPosterImage(i *html.Node) (bool, string) {
	poster, has_poster := htmlnode.GetAttributeVal(i, "", "poster")
	if !has_poster {
		return false, ""
	}

	layout, has_layout := htmlnode.GetAttributeVal(i, "", "layout")
	if has_layout && layout == "nodisplay" {
		return false, ""
	}

	if isTinyNode(i) {
		return false, ""
	}

	return true, poster
}

// Decides if the given image node qualifies for preloading.
func candidateImageForPreloading(i *html.Node) (bool, string) {
	imgsrc, has_src := htmlnode.GetAttributeVal(i, "", "src")
	width_int, height_int := nodeDimensions(i)

	if isTinyNode(i) {
		// Checks for aspect ratio images.
		if width_int > 0 && width_int <= 16 && height_int > 0 && height_int <= 16 && isAspectRatioDimensions(width_int, height_int) {
			return true, imgsrc
		}
		return false, ""
	}

	// Ignores images with no src attribute.
	// These cane be css images inside class definition.
	if !has_src || len(imgsrc) == 0 {
		return false, ""
	}

	// Ignores inline images.
	if strings.HasPrefix(imgsrc, "data:image") {
		return false, ""
	}

	layout_type, has_layout := htmlnode.GetAttributeVal(i, "", "layout")
	_, has_placeholder := htmlnode.GetAttributeVal(i, "", "placeholder")

	// Ignores images which are not displayed on load.
	if has_layout && layout_type == "nodisplay" {
		return false, ""
	}

	// Checks if it is placeholder image for iframe.
	// https://www.ampproject.org/docs/reference/components/amp-iframe#iframe-with-placeholder
	if has_placeholder {
		if i.Parent.Data == "amp-iframe" {
			// TODO(amaltas): Check for additional allowfullscreen attribute.
			if isTinyNode(i.Parent) {
				return false, ""
			} else {
				return true, imgsrc
			}
		} else {
			return false, ""
		}
	}

	if has_layout && (layout_type == "responsive" || layout_type == "fill") {

		if width_int == 0 && height_int == 0 {
			if isTinyNode(i.Parent) {
				return false, ""
			}
			return true, imgsrc
		}

		// Actual image dimension check is performed later.
	}

	// For other layouts with no image dimensions, take parent containers
	// dimensions into account.
	if width_int == 0 && height_int == 0 {
		// Account for parent dimension.
		width_int, height_int = nodeDimensions(i.Parent)
	}

	if (width_int > 149 || width_int == 0) && height_int > 149 {
		return true, imgsrc
	}

	return false, ""
}

// Consider a small dimension size as aspect ratio if they are relatively prime.
// TODO(amaltas): Fix it for float dimension types: 1x1.33.
func isAspectRatioDimensions(width, height int) bool {
	if width > 16 || height > 16 {
		return false
	}

	return new(big.Int).GCD(nil, nil, big.NewInt(int64(width)), big.NewInt(int64(height))).Int64() == 1
}

func isTinyNode(i *html.Node) bool {
	width, height := nodeDimensions(i)
	return (width > 0 && width < 150) || (height > 0 && height < 150)
}

func nodeDimensions(i *html.Node) (int, int) {
	// Width and Height as int type.
	width, has_width := htmlnode.GetAttributeVal(i, "", "width")
	height, has_height := htmlnode.GetAttributeVal(i, "", "height")
	width_int := 0
	height_int := 0
	var err error

	if has_width {
		if width_int, err = strconv.Atoi(width); err != nil {
			return 0, 0
		}
	}

	if has_height {
		if height_int, err = strconv.Atoi(height); err != nil {
			return 0, 0
		}
	}

	return width_int, height_int
}
