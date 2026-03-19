package day01

import "testing"

func TestPart1(t *testing.T) {
	t.Skip("skipping...")
	tests := []struct {
		in, want string
	}{
		{"1721\n979\n366\n299\n675\n1456", "514579"},
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
		{"1721\n979\n366\n299\n675\n1456", "241861950"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part2(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}
