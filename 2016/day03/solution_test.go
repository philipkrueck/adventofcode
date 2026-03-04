package day03

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"5 10 25", "0"},
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
		{
			"101 301 501\n102 302 502\n103 303 503\n201 401 601\n202 402 602\n203 403 603", "6",
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part2(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}
