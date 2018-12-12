package amphtml

import (
	"testing"
)

func TestParseSrcset(t *testing.T) {
	tcs := []struct {
		desc, input, normalized string
		expected                []SubresourceOffset
	}{
		{
			desc:       "absolute",
			input:      "http://www.example.com/blah.jpg 50w",
			normalized: "http://www.example.com/blah.jpg 50w",
			expected:   []SubresourceOffset{{End: 31}},
		},
		{
			desc:       "relative",
			input:      "blah.jpg 50w",
			normalized: "blah.jpg 50w",
			expected:   []SubresourceOffset{{End: 8}},
		},
		{
			desc:       "default density",
			input:      "http://www.example.com/blah.jpg",
			normalized: "http://www.example.com/blah.jpg 1x",
			expected:   []SubresourceOffset{{End: 31}},
		},
		{
			desc:       "multiple",
			input:      "image1 2x, image2, image3 3x, image4 4x ",
			normalized: "image2 1x, image1 2x, image3 3x, image4 4x",
			expected: []SubresourceOffset{{Start: 0, End: 6},
				{Start: 11, End: 17},
				{Start: 22, End: 28},
				{Start: 33, End: 39}},
		},
		{
			desc:       "whitespace",
			input:      "  \t\n http://www.example.com/blah.jpg  \n\t\t ",
			normalized: "http://www.example.com/blah.jpg 1x",
			expected:   []SubresourceOffset{{End: 31}},
		},
		{
			desc:       "leading comma",
			input:      " , http://www.example.com/blah.jpg \n\t\t ",
			normalized: "http://www.example.com/blah.jpg 1x",
			expected:   []SubresourceOffset{{End: 31}},
		},
		{
			desc:       "comma in url",
			input:      " , http://www.example.com/image,1 \n\t\t ",
			normalized: "http://www.example.com/image,1 1x",
			expected:   []SubresourceOffset{{End: 30}},
		},
		{
			desc:       "commas everywhere",
			input:      ",http://www.example.com/,/,/,/,50w,",
			normalized: "http://www.example.com/,/,/,/,50w 1x",
			expected:   []SubresourceOffset{{End: 33}},
		},
		{
			desc:       "missing delimiter noop",
			input:      "image1 100w image2 50w",
			normalized: "image1 100w image2 50w",
			expected:   []SubresourceOffset{},
		},
		{
			desc:       "negative value noop",
			input:      "image1 100w, image2 -50w",
			normalized: "image1 100w, image2 -50w",
			expected:   []SubresourceOffset{},
		},
		{
			desc:       "duplicate default no-op",
			input:      "image1, image2, image3 3x, image4 4x ",
			normalized: "image1, image2, image3 3x, image4 4x ",
			expected:   []SubresourceOffset{},
		},
		{
			desc:       "duplicate explicit no-op",
			input:      "image1 2x, image2, image3 2x, image4 4x ",
			normalized: "image1 2x, image2, image3 2x, image4 4x ",
			expected:   []SubresourceOffset{},
		},
	}
	var equals = func(a, b []SubresourceOffset) bool {
		if len(a) != len(b) {
			return false
		}
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}

	for _, tc := range tcs {
		normalized, actual := ParseSrcset(tc.input)
		if normalized != tc.normalized {
			t.Errorf("%s: ParseSrcset(%q)=%q, want=%q", tc.desc, tc.input, normalized, tc.normalized)
		}
		if !equals(actual, tc.expected) {
			t.Errorf("%s: ParseSrcset(%q)=%v, want=%v", tc.desc, tc.input, actual, tc.expected)
		}
	}
}

func TestGetSrcsetWidths(t *testing.T) {
	tcs := []struct {
		input    int
		expected []int
		ok       bool
	}{
		{-20, nil, false},
		{1, nil, false},
		{20, []int{39, 47, 68}, true},
		{500, []int{560, 1000, 1200}, true},
		{600, []int{680, 1200}, true},
		{1000, []int{1000, 1200}, true},
		{1001, nil, false},
	}
	var equals = func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}
	for _, tc := range tcs {
		actual, ok := GetSrcsetWidths(tc.input)
		if ok != tc.ok {
			t.Errorf("GetSrcsetWidths(%d)=%t want=%t", tc.input, ok, tc.ok)
		}
		if ok {
			if actual == nil {
				t.Errorf("GetSrcsetWidths(%d) returned ok, but nil srcset", tc.input)
			} else if !equals(actual, tc.expected) {
				t.Errorf("GetSrcsetWidths(%d)=%v want=%v", tc.input, actual, tc.expected)
			}
		}
	}
}
