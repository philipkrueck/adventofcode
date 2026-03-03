package day02

import "testing"

const testInstructrs = "ULL\nRRDDD\nLURDL\nUUUUD"

func TestPart1(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{testInstructrs, "1985"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part1(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{testInstructrs, "5DB3"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part2(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}
