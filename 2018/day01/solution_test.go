package day01

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"+1\n+1\n+1\n", "3"},
		{"+1\n+1\n-2\n", "0"},
		{"-1\n-2\n-3\n", "-6"},
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
		{"+1\n-1", "0"},
		{"+3\n+3\n+4\n-2\n-4", "10"},
		{"-6\n+3\n+8\n+5\n-6", "5"},
		{"+7\n+7\n-2\n-7\n-4", "14"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part2(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}
