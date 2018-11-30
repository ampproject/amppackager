package amphtml

import (
	"testing"
)

func TestTokenizeSrcset(t *testing.T) {
	tcs := []struct {
		desc, input string
		expected    []SubresourceURL
	}{
		{
			desc:     "absolute",
			input:    "http://www.example.com/blah.jpg 50w",
			expected: []SubresourceURL{{URLString: "http://www.example.com/blah.jpg", descriptor: "50w"}},
		},
		{
			desc:     "relative",
			input:    "blah.jpg 50w",
			expected: []SubresourceURL{{URLString: "blah.jpg", descriptor: "50w"}},
		},
		{
			desc:     "default density",
			input:    "http://www.example.com/blah.jpg",
			expected: []SubresourceURL{{URLString: "http://www.example.com/blah.jpg", descriptor: "1x"}},
		},
		{
			desc:  "multiple",
			input: "image1 2x, image2, image3 3x, image4 4x ",
			expected: []SubresourceURL{{URLString: "image1", descriptor: "2x"},
				{URLString: "image2", descriptor: "1x"},
				{URLString: "image3", descriptor: "3x"},
				{URLString: "image4", descriptor: "4x"}},
		},
		{
			desc:     "whitespace",
			input:    "  \t\n http://www.example.com/blah.jpg  \n\t\t ",
			expected: []SubresourceURL{{URLString: "http://www.example.com/blah.jpg", descriptor: "1x"}},
		},
		{
			desc:     "leading comma",
			input:    " , http://www.example.com/blah.jpg \n\t\t ",
			expected: []SubresourceURL{{URLString: "http://www.example.com/blah.jpg", descriptor: "1x"}},
		},
		{
			desc:     "comma in url",
			input:    " , http://www.example.com/image,1 \n\t\t ",
			expected: []SubresourceURL{{URLString: "http://www.example.com/image,1", descriptor: "1x"}},
		},
		{
			desc:     "commas everywhere",
			input:    ",http://www.example.com/,/,/,/,50w,",
			expected: []SubresourceURL{{URLString: "http://www.example.com/,/,/,/,50w", descriptor: "1x"}},
		},
		{
			desc:     "missing delimiter noop",
			input:    "image1 100w image2 50w",
			expected: []SubresourceURL{},
		},
		{
			desc:     "negative value noop",
			input:    "image1 100w, image2 -50w",
			expected: []SubresourceURL{},
		},
		{
			desc:     "duplicate default no-op",
			input:    "image1, image2, image3 3x, image4 4x ",
			expected: []SubresourceURL{},
		},
		{
			desc:     "duplicate explicit no-op",
			input:    "image1 2x, image2, image3 2x, image4 4x ",
			expected: []SubresourceURL{},
		},
	}
	var equals = func(a, b []SubresourceURL) bool {
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
		actual := TokenizeSrcset(tc.input)
		if !equals(actual, tc.expected) {
			t.Errorf("%s: TokenizeSrcset(%s)=%s, want=%s", tc.desc, tc.input, actual, tc.expected)
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
