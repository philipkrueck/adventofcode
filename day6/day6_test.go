package day6

import (
	"testing"
)

func TestOperateOnColumns(t *testing.T) {
	grid := [][]int{
		{123, 328, 51, 64},
		{45, 64, 387, 23},
		{6, 98, 215, 314},
	}
	operators := []Op{Multiply, Add, Multiply, Add}

	got := operateColumns(grid, operators)
	want := 4277556

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestSomething(t *testing.T) {
	cases := []struct {
		Want int
	}{}

	for _, test := range cases {
		t.Run("name", func(t *testing.T) {
			got := Part1()

			if got != test.Want {
				t.Errorf("got: %d; want: %d\n", got, test.Want)
			}
		})
	}
}
