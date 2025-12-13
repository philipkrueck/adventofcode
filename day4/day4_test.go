package day4

import (
	"fmt"
	"testing"
)

func TestIsPaper(t *testing.T) {
	grid := []string{"@"}

	cases := []struct {
		Grid Grid
		Idx  Index
		Want bool
	}{
		{grid, Index{-1, -1}, false},
		{grid, Index{-1, 0}, false},
		{grid, Index{-1, +1}, false},
		{grid, Index{0, -1}, false},
		{grid, Index{0, 1}, false},
		{grid, Index{1, -1}, false},
		{grid, Index{1, 0}, false},
		{grid, Index{1, +1}, false},
		{grid, Index{0, 0}, true},
	}

	for _, test := range cases {
		got := test.Grid.HasPaper(test.Idx)
		t.Run(fmt.Sprintf("%v at (idx=%d)", test.Grid, test.Idx), func(t *testing.T) {
			if got != test.Want {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}

var grid = Grid([]string{
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
})

func TestIsAccessible(t *testing.T) {
	cases := []struct {
		Grid Grid
		Idx  Index
		Want bool
	}{
		{grid, Index{0, 0}, false},
		{grid, Index{0, 1}, false},
		{grid, Index{0, 2}, true},
		{grid, Index{0, 3}, true},
		{grid, Index{0, 4}, false},
		{grid, Index{0, 5}, true},
		{grid, Index{0, 6}, true},
		{grid, Index{0, 7}, false},
		{grid, Index{0, 8}, true},
		{grid, Index{0, 9}, false},
		{grid, Index{1, 0}, true},
		{grid, Index{1, 1}, false},
		{grid, Index{1, 2}, false},
		{grid, Index{1, 3}, false},
		{grid, Index{1, 4}, false},
		{grid, Index{1, 5}, false},
		{grid, Index{1, 6}, false},
		{grid, Index{1, 7}, false},
		{grid, Index{1, 8}, false},
		{grid, Index{1, 9}, false},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("element at (idx=%d)", test.Idx), func(t *testing.T) {
			got := test.Grid.IsAccessible(test.Idx)

			if got != test.Want {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}

// func TestAccessibleIndices(t *testing.T) {
// 	got := accessibleIndices(grid)
// 	want :=
//
// 	if got != want {
// 		t.Errorf("got: %d, want: %d", got, want)
// 	}
// }

func TestSumAccessibleRolls(t *testing.T) {
	got := grid.CountAccessible()
	want := 13

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
