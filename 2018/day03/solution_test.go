package day03

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2\n", "4"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part1(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}

func TestMustParseRectangle(t *testing.T) {
	tests := []struct {
		line string
		rec  rectangle
	}{
		{
			"#1 @ 2,3: 4x5",
			rectangle{1, 2, 3, 4, 5},
		},
		{
			"#108 @ 769,356: 22x13",
			rectangle{108, 769, 356, 22, 13},
		},
	}

	for _, tt := range tests {
		if got := mustParseRectangle(tt.line); got != tt.rec {
			t.Errorf("\ngot:  %#v \nwant: %#v", got, tt.rec)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		// {"#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2\n", "3"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part2(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}

var (
	r1 = rectangle{0, 1, 3, 4, 4}
	r2 = rectangle{0, 3, 1, 4, 4}
	r3 = rectangle{0, 5, 5, 2, 2}
)

func TestOverlap(t *testing.T) {
	tests := []struct {
		r1, r2    rectangle
		intersect bool
	}{
		{r1, r2, true},
		{r2, r1, true},
		{r1, r3, false},
		{r3, r1, false},
		{r2, r3, false},
		{r3, r2, false},
	}

	for _, tt := range tests {
		if got := overlap(tt.r1, tt.r2); got != tt.intersect {
			t.Errorf("got: %v; want: %v\n", got, tt.intersect)
		}
	}
}
