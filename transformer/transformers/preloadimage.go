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
	"sort"
	"strconv"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html"
)

const maxHeroImages int = 2

// HeroImage represents the necessary data to inject a <link ref=preload> and optional <img> tag.
type HeroImage struct {
	src    string
	srcset string
	ampImg *html.Node
}

// mediaQuerySource represents an individual source from a srcset with a unique media query that will only load the href when appropriate.
type mediaQuerySource struct {
	href  string
	media string
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

	return nil
}

func prioritizeHeroImage(e *Context, heroImage HeroImage) {
	for _, link := range buildLinkPreloads(heroImage) {
		e.DOM.HeadNode.AppendChild(link)
	}

	if ampImg := heroImage.ampImg; ampImg != nil {
		img := buildImg(ampImg)
		htmlnode.SetAttribute(ampImg, "", "i-amphtml-ssr", "")
		ampImg.AppendChild(img)
	}
}

func buildLinkPreloads(heroImage HeroImage) []*html.Node {
	// One of these is guaranteed.
	if heroImage.srcset != "" {
		if medias, ok := srcsetToMediaQueries(heroImage.srcset); ok {
			links := make([]*html.Node, len(medias))
			for i, mediaSrc := range medias {
				links[i] = htmlnode.Element("link",
					html.Attribute{Key: "rel", Val: "preload"},
					html.Attribute{Key: "as", Val: "image"},
					html.Attribute{Key: "href", Val: mediaSrc.href},
					html.Attribute{Key: "media", Val: mediaSrc.media},
				)
			}

			return links
		}
	}

	if heroImage.src == "" {
		return []*html.Node{}
	}

	return []*html.Node{
		htmlnode.Element("link",
			html.Attribute{Key: "rel", Val: "preload"},
			html.Attribute{Key: "as", Val: "image"},
			html.Attribute{Key: "href", Val: heroImage.src},
		),
	}
}

func buildImg(ampImg *html.Node) *html.Node {
	img := htmlnode.Element("img",
		html.Attribute{Key: "class", Val: "i-amphtml-fill-content i-amphtml-replaced-content"},
		html.Attribute{Key: "decoding", Val: "async"})
	attrsToCopy := [...]string{
		"alt",
		"attribution",
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

// Converts the raw srcset value into multiple sources with an appropriate media query to select that source.
// TODO(jridgewell, amaltas): Support pixel density with -webkit-max-device-pixel-ratio and -webkit-min-device-pixel-ratio
func srcsetToMediaQueries(srcset string) ([]mediaQuerySource, bool) {
	type Source struct {
		href string
		size int
	}

	srcSets := strings.Split(strings.TrimSpace(srcset), ",")
	length := len(srcSets)
	sources := make([]Source, length)
	medias := make([]mediaQuerySource, length)

	if length == 0 {
		return medias, false
	}

	for i, src := range srcSets {
		imgComponents := strings.Fields(src)
		if len(imgComponents) != 2 {
			return medias, false
		}

		source := Source{imgComponents[0], 0}

		if strings.HasSuffix(imgComponents[1], "w") {
			size, err := strconv.Atoi(strings.TrimSuffix(imgComponents[1], "w"))
			if err != nil {
				return medias, false
			}

			source.size = size
		} else {
			return medias, false
		}

		sources[i] = source
	}

	// Sort the images based on their target sizes in asc order.
	sort.Slice(sources, func(i, j int) bool {
		return sources[i].size < sources[j].size
	})

	for i, ci := range sources {
		var mediaQuery string
		if i == 0 {
			mediaQuery = fmt.Sprintf("(max-width: %dpx)", ci.size)
		} else if i == length-1 {
			mediaQuery = fmt.Sprintf("(min-width: %dpx)", sources[i-1].size+1)
		} else {
			mediaQuery = fmt.Sprintf("(min-width: %dpx) and (max-width: %dpx)", sources[i-1].size+1, ci.size)
		}

		medias[i] = mediaQuerySource{ci.href, "screen and " + mediaQuery}
	}
	return medias, true
}
