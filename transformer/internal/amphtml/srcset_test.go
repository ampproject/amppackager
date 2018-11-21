package amphtml

import (
	"testing"
)

func TestNewSrcsetWidth(t *testing.T) {
	tcs := []struct {
		input    int
		expected SrcsetWidth
		ok       bool
	}{
		{1, nil, false},
		{20, SrcsetWidth{39, 47, 68}, true},
		{500, SrcsetWidth{560, 1000, 1200}, true},
		{600, SrcsetWidth{680, 1200}, true},
		{1000, SrcsetWidth{1000, 1200}, true},
		{1001, nil, false},
	}
	for _, tc := range tcs {
		actual, ok := NewSrcsetWidth(tc.input)
		if ok != tc.ok {
			t.Errorf("NewSrcsetWidth(%d)=%t want=%t", tc.input, ok, tc.ok)
		}
		if ok {
			if actual == nil {
				t.Errorf("NewSrcsetWidth(%d) returned ok, but nil srcset", tc.input)
			} else if !equal(actual, tc.expected) {
				t.Errorf("NewSrcsetWidth(%d)=%v want=%v", tc.input, actual, tc.expected)
			}
		}
	}
}

func equal(a, b SrcsetWidth) bool {
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
