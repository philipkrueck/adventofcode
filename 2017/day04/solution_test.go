package day04

import "testing"

var in1 = `aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aa`

func TestPart1(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{in1, "1"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part1(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}

var in2 = `abcde fghij
abcde xyz ecdab
a ab abc abd abf abj
iiii oiii ooii oooi oooo
oiii ioii iioi iiio`

func TestPart2(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{in2, "3"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part2(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}
