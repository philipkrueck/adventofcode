package day03

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{">", "2"},
		{"^>v<", "4"},
		{"^v^v^v^v^v", "2"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part1(tt.in); got != tt.want {
				t.Errorf("got: %q, want: %q", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"^v", "3"},
		{"^v>", "4"},
		{"^>v<", "3"},
		{"^v^v^v^v^v", "11"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part2(tt.in); got != tt.want {
				t.Errorf("got: %q, want: %q", got, tt.want)
			}
		})
	}
}
