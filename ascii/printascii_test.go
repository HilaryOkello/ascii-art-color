package ascii

import (
	"testing"
)

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

func Test_shouldPrintWithColor(t *testing.T) {
	type args struct {
		args    *PrintArgs
		indices []int
		track   int
		i       int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Color but No Substring",
			args: args{args: &PrintArgs{Color: "31", Str: "Hello World", Substr: ""}, indices: []int{}, track: 0, i: 0},
			want: true,
		},
		{
			name: "Substring Match 01",
			args: args{args: &PrintArgs{Color: "31", Str: "hello world", Substr: "hello"}, indices: []int{0}, track: 0, i: 3},
			want: true,
		},
		{
			name: "Substring Match But Index Not Colored",
			args: args{args: &PrintArgs{Color: "31", Str: "hello world", Substr: "world"}, indices: []int{6}, track: 0, i: 3},
			want: false,
		},
		{
			name: "Multiple Indices",
			args: args{args: &PrintArgs{Color: "31", Str: "hello world hello", Substr: "hello"}, indices: []int{0, 12}, track: 1, i: 13},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shouldPrintWithColor(tt.args.args, tt.args.indices, tt.args.track, tt.args.i); got != tt.want {
				t.Errorf("shouldPrintWithColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
