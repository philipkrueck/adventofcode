package day02

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"2x3x4", "58"},
		{"1x1x10", "43"},
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
		{"2x3x4", "34"},
		{"1x1x10", "14"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part2(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}

func TestParsePresent(t *testing.T) {
	tests := []struct {
		in                  string
		wantL, wantW, wantH int
	}{
		{
			"2x3x4",
			2, 3, 4,
		},
		{
			"1x1x10",
			1, 1, 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			gotL, gotW, gotH := parseLines(tt.in)
			if gotL != tt.wantL || gotW != tt.wantW || gotH != tt.wantH {
				t.Errorf("Got: (%d,%d,%d); Want: (%d,%d,%d)", gotL, gotW, gotH, tt.wantL, tt.wantW, tt.wantH)
			}
		})
	}
}
