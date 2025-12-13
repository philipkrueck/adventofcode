package day4

import (
	"fmt"
	"testing"
)

func TestIsPaper(t *testing.T) {
	grid := []string{"@"}

	cases := []struct {
		Grid []string
		I, J int
		Want bool
	}{
		{grid, -1, -1, false},
		{grid, -1, 0, false},
		{grid, -1, +1, false},
		{grid, 0, -1, false},
		{grid, 0, 1, false},
		{grid, 1, -1, false},
		{grid, 1, 0, false},
		{grid, 1, +1, false},
		{grid, 0, 0, true},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%v at (i=%d,j=%d)", test.Grid, test.I, test.J), func(t *testing.T) {
			got := isPaper(test.Grid, test.I, test.J)

			if got != test.Want {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}

var grid = []string{
	"..@@.@@@@.",
	"@@@.@.@.@@",
	"@@@@@.@.@@",
	"@.@@@@..@.",
	"@@.@@@@.@@",
	".@@@@@@@.@",
	".@.@.@.@@@",
	"@.@@@.@@@@",
	".@@@@@@@@.",
	"@.@.@@@.@.",
}

func TestIsAccessible(t *testing.T) {
	cases := []struct {
		Grid []string
		I, J int
		Want bool
	}{
		{grid, 0, 0, false},
		{grid, 0, 1, false},
		{grid, 0, 2, true},
		{grid, 0, 3, true},
		{grid, 0, 4, false},
		{grid, 0, 5, true},
		{grid, 0, 6, true},
		{grid, 0, 7, false},
		{grid, 0, 8, true},
		{grid, 0, 9, false},
		{grid, 1, 0, true},
		{grid, 1, 1, false},
		{grid, 1, 2, false},
		{grid, 1, 3, false},
		{grid, 1, 4, false},
		{grid, 1, 5, false},
		{grid, 1, 6, false},
		{grid, 1, 7, false},
		{grid, 1, 8, false},
		{grid, 1, 9, false},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("element at (i=%d,j=%d)", test.I, test.J), func(t *testing.T) {
			got := isAccessible(test.Grid, test.I, test.J)

			if got != test.Want {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}

func TestSumAccessibleRolls(t *testing.T) {
	got := sumAccessibleRolls(grid)
	want := 13

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
