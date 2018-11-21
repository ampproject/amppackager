package amphtml

import (
	"net/url"
	"testing"
)

func TestGetSrcsetFromSrc(t *testing.T) {
	tcs := []struct {
		desc, input, expected string
		width                 int
		expectedOk            bool
	}{
		{
			desc:       "valid",
			input:      "http://www.example.com/blah.jpg",
			width:      92,
			expected:   `https://www-example-com.cdn.ampproject.org/ii/w100/www.example.com/blah.jpg 100w, https://www-example-com.cdn.ampproject.org/ii/w220/www.example.com/blah.jpg 220w, https://www-example-com.cdn.ampproject.org/ii/w330/www.example.com/blah.jpg 330w`,
			expectedOk: true,
		},
		{
			desc:       "bad width",
			input:      "http://www.example.com/blah.jpg",
			width:      1,
			expectedOk: false,
		},
	}
	for _, tc := range tcs {
		actual, ok := GetSrcsetFromSrc(tc.input, tc.width)
		if ok != tc.expectedOk {
			t.Errorf("GetSrcsetFromSrc ok got=%t, want=%t", ok, tc.expectedOk)
		}
		if ok && actual != tc.expected {
			t.Errorf("GetSrcsetFromSrc(%s)=%s, want=%s", tc.input, actual, tc.expected)
		}
	}
}

func TestConvertSrcset(t *testing.T) {
	tcs := []struct {
		desc, input, expected string
	}{
		{
			desc:     "rewritten",
			input:    "http://www.example.com/blah.jpg 50w",
			expected: "https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 50w",
		},
		{
			desc:     "relative rewritten",
			input:    "blah.jpg 50w",
			expected: "https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 50w",
		},
		{
			desc:     "secure",
			input:    "https://www.example.com/blah.jpg 50w",
			expected: "https://www-example-com.cdn.ampproject.org/i/s/www.example.com/blah.jpg 50w",
		},
		{
			desc:     "default density",
			input:    "http://www.example.com/blah.jpg",
			expected: "https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 1x",
		},
		{
			desc:     "multiple",
			input:    "image1 2x, image2, image3 3x, image4 4x ",
			expected: "https://www-example-com.cdn.ampproject.org/i/www.example.com/image1 2x, https://www-example-com.cdn.ampproject.org/i/www.example.com/image2 1x, https://www-example-com.cdn.ampproject.org/i/www.example.com/image3 3x, https://www-example-com.cdn.ampproject.org/i/www.example.com/image4 4x",
		},
		{
			desc:     "whitespace",
			input:    "  \t\n http://www.example.com/blah.jpg  \n\t\t ",
			expected: "https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 1x",
		},
		{
			desc:     "leading comma",
			input:    " , http://www.example.com/blah.jpg \n\t\t ",
			expected: "https://www-example-com.cdn.ampproject.org/i/www.example.com/blah.jpg 1x",
		},
		{
			desc:     "comma in url",
			input:    " , http://www.example.com/image,1 \n\t\t ",
			expected: "https://www-example-com.cdn.ampproject.org/i/www.example.com/image,1 1x",
		},
		{
			desc:     "commas everywhere",
			input:    ",http://www.example.com/,/,/,/,50w,",
			expected: "https://www-example-com.cdn.ampproject.org/i/www.example.com/,/,/,/,50w 1x",
		},
		{
			desc:     "missing delimiter noop",
			input:    "image1 100w image2 50w",
			expected: "image1 100w image2 50w",
		},
		{
			desc:     "negative value noop",
			input:    "image1 100w, image2 -50w",
			expected: "image1 100w, image2 -50w",
		},
	}
	baseURL, _ := url.Parse("http://www.example.com")
	for _, tc := range tcs {
		actual := ConvertSrcset(baseURL, tc.input)
		if actual != tc.expected {
			t.Errorf("%s: ConvertSrcset(%s)=%s, want=%s", tc.desc, tc.input, actual, tc.expected)
		}
	}
}

func TestGetWidths(t *testing.T) {
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
	for _, tc := range tcs {
		actual, ok := getWidths(tc.input)
		if ok != tc.ok {
			t.Errorf("getWidths(%d)=%t want=%t", tc.input, ok, tc.ok)
		}
		if ok {
			if actual == nil {
				t.Errorf("getWidths(%d) returned ok, but nil srcset", tc.input)
			} else if !equal(actual, tc.expected) {
				t.Errorf("getWidths(%d)=%v want=%v", tc.input, actual, tc.expected)
			}
		}
	}
}

func equal(a, b []int) bool {
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
