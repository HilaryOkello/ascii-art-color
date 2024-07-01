package ascii

import "testing"

func TestGetIndices(t *testing.T) {
	testCases := []struct {
		str      string
		substr   string
		expected []int
	}{
		{
			str:      "hello world",
			substr:   "lo",
			expected: []int{3},
		},
		{
			str:      "hello world",
			substr:   "o",
			expected: []int{4, 7},
		},
		{
			str:      "hello world",
			substr:   "world",
			expected: []int{6},
		},
		{
			str:      "hello world",
			substr:   "",
			expected: []int{},
		},
		{
			str:      "hello world",
			substr:   "z",
			expected: []int{},
		},
		{
			str:      "aaaaaa",
			substr:   "aa",
			expected: []int{0, 2, 4},
		},
		{
			str:      "hello",
			substr:   "hello",
			expected: []int{0},
		},
		{
			str:      "",
			substr:   "empty",
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.str+"_"+tc.substr, func(t *testing.T) {
			result := GetIndices(tc.str, tc.substr)
			if !equalSlices(result, tc.expected) {
				t.Errorf("GetIndices(%q, %q) = %v; want %v", tc.str, tc.substr, result, tc.expected)
			}
		})
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
