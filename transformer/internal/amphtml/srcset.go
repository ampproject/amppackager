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

// SrcsetWidth is a slice of at least two legitimate widths
type SrcsetWidth []int

var (
	// List of legitimate widths. These were computed with step size of 20% starting from 100.
	legitimateWidths = [...]int{39, 47, 56, 68, 82, 100, 120, 150, 180, 220, 270, 330, 390, 470, 560, 680, 820, 1000, 1200}

	// List of supported pixel device densities (DPRs)
	dprs = [3]float32{1.0, 2.0, 3.0}
)

// NewSrcsetWidth returns a SrcsetWidth based on the input width, or false if there are
// not at least two legitimate widths.
func NewSrcsetWidth(w int) (SrcsetWidth, bool) {
	prev := -1
	ssw := SrcsetWidth{}
	for _, dpr := range dprs {
		// int cast truncates. Add 0.5 to achieve rounding.
		width := roundUp(int(dpr*float32(w) + 0.5))
		if width != prev {
			ssw = append(ssw, width)
		}
		prev = width
	}
	if len(ssw) > 1 {
		return ssw, true
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
