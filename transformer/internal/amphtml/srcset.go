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

package amphtml

import (
	"net/url"
	"regexp"
	"strings"
)

// GetSrcsetFromSrc returns a srcset attribute value of image candidate strings
// of various sizes generated from the given src absolute URL and a starting width.
func GetSrcsetFromSrc(src string, width int) (string, bool) {
	widths, ok := getWidths(width)
	if !ok {
		return "", false
	}
	req := ImageURLRequest{Input: src}
	var ss []string
	for _, w := range widths {
		req.desiredWidth = w
		ss = append(ss, req.GetCacheImageURL())
	}
	return strings.Join(ss, ", "), true
}

const defaultDensity = "1x"

var (
	// List of legitimate widths. These were computed with step size of 20% starting from 100.
	legitimateWidths = [...]int{39, 47, 56, 68, 82, 100, 120, 150, 180, 220, 270, 330, 390, 470, 560, 680, 820, 1000, 1200}

	// List of supported pixel device densities (DPRs)
	dprs = [3]float32{1.0, 2.0, 3.0}
)

// getWidths returns a slice of widths based on the input width, or false if there are
// not at least two legitimate widths.
func getWidths(w int) ([]int, bool) {
	if w < 0 {
		return nil, false
	}
	prev := -1
	var widths []int
	for _, dpr := range dprs {
		// int cast truncates. Add 0.5 to achieve rounding.
		width := roundUp(int(dpr*float32(w) + 0.5))
		if width != prev {
			widths = append(widths, width)
		}
		prev = width
	}
	if len(widths) > 1 {
		return widths, true
	}
	return nil, false
}

// roundUp rounds up to the closest legitimate width (or the largest possible) and returns it.
func roundUp(w int) int {
	for _, width := range legitimateWidths {
		if width >= w {
			return width
		}
	}
	return legitimateWidths[len(legitimateWidths)-1]
}


// Regex for leading spaces, followed by an optional comma and whitespace,
// followed by an URL*, followed by an optional space, followed by an
// optional width or pixel density**, followed by spaces, followed by an
// optional comma and whitespace.
//
// URL*: matches non-space, non-empty string which neither ends nor begins
// with a comma. The set of space characters in the srcset attribute is
// defined to include only ascii characters, so using \s, which is an
// ascii only character set, is fine. See
// https://html.spec.whatwg.org/multipage/infrastructure.html#space-character.
//
// Optional width or pixel density**: Matches the empty string or (one or
// more spaces + a non empty string containing no space or commas).
// Doesn't capture the initial space.
//
// \s*                       Match, but don't capture leading spaces
// (?:,\s*)?                 Optionally match comma and trailing space,
//                           but don't capture comma.
// ([^,\s]\S*[^,\s])         Match something like "google.com/favicon.ico"
//                           but not ",google.com/favicon.ico,"
// \s*                       Match, but don't capture spaces.
// ([\d]+.?[\d]*[w|x])?      e.g. "5w" or "5x" or "10.2x"
// \s*                       Match, but don't capture space
// (?:(,)\s*)?               Optionally match comma and trailing space,
//                           capturing comma.
var imageCandidateRE = regexp.MustCompile(`\s*(?:,\s*)?([^,\s]\S*[^,\s])\s*([\d]+.?[\d]*[w|x])?\s*(?:(,)\s*)?`)

// ConvertSrcset returns a new string from the given srcset attribute value,
// parsing the image candidates (as defined by
// https://html.spec.whatwg.org/multipage/images.html#image-candidate-string
// and rewrites URLS to point to the AMP Cache. If there is no width or
// pixel density, it defaults to 1x.
// If any portion of the input is unparseable, or if there are duplicate widths
// or pixel densities, return the input, unconverted.
func ConvertSrcset(base *url.URL, in string) string {
	matches := imageCandidateRE.FindAllStringSubmatch(in, -1)
	if len(matches) == 0 {
		return in
	}
	var sb strings.Builder
	seen := make(map[string]struct{})
	for i, m := range matches {
		d := defaultDensity
		if len(m[2]) > 0 {
			d = m[2]
		}
		if _, ok := seen[d]; ok {
			// duplicate width or pixel density
			return in
		}
		seen[d] = struct{}{}
		req := ImageURLRequest{Input: ToPortableURL(base, m[1])}
		sb.WriteString(req.GetCacheImageURL())
		sb.WriteRune(' ')
		sb.WriteString(d)
		if i < len(matches)-1 {
			if len(m[3]) == 0 {
				// missing expected comma delimiter
				return in
			}
			sb.WriteString(", ")
		}
	}
	return sb.String()
}

