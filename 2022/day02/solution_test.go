package day02

import "testing"

var in = `A Y
B X
C Z`

func TestPart1(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{in, "15"},
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
		{in, "12"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part2(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}
