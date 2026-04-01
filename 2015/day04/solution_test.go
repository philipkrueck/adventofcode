package day04

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"abcdef", "609043"},
		{"pqrstuv", "1048970"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part1(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}
