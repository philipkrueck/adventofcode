package day3

import (
	"testing"
)

func TestLargestVoltage(t *testing.T) {
	cases := []struct {
		Bank string
		K    int
		Want int
	}{
		{"981111", 2, 98},
		{"811119", 2, 89},
		{"123456", 2, 56},
		{"171151", 2, 75},
	}
	for _, test := range cases {
		t.Run(test.Bank, func(t *testing.T) {
			got := largestVoltage(test.Bank, test.K)

			if got != test.Want {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}
