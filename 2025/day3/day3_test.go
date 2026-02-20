package day3

import (
	"testing"
)

func TestLargestVoltage(t *testing.T) {
	cases := []struct {
		Line string
		K    int
		Want int
	}{
		{"981111", 2, 98},
		{"811119", 2, 89},
		{"123456", 2, 56},
		{"171151", 2, 75},
		{"123456", 3, 456},
		{"987654321111111", 12, 987654321111},
		{"811111111111119", 12, 811111111119},
		{"234234234234278", 12, 434234234278},
		{"818181911112111", 12, 888911112111},
	}
	for _, test := range cases {
		t.Run(test.Line, func(t *testing.T) {
			got := largestVoltage(test.Line, test.K)

			if got != test.Want {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}
