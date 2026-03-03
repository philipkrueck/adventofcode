package day01

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"R2, L3", "5"},
		{"R2, R2, R2", "2"},
		{"R5, L5, R5, R3", "12"},
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
		{"R8, R4, R4, R8", "4"},
		{"L1, L1, L2, L2, L2, L3", "1"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part2(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}
