package day7

import (
	"fmt"
	"slices"
	"testing"
)

func TestSplitBeams(t *testing.T) {
	cases := []struct {
		Lines [][]byte
		Want  [][]byte
	}{
		{
			[][]byte{
				{'.', '|', '.'},
				{'.', '^', '.'},
			},
			[][]byte{
				{'.', '|', '.'},
				{'|', '^', '|'},
			},
		},
		{
			[][]byte{
				{'|', '.', '|'},
				{'.', '.', '.'},
			},
			[][]byte{
				{'|', '.', '|'},
				{'|', '.', '|'},
			},
		},
		{
			[][]byte{
				{'.', 'S', '.'},
				{'.', '.', '.'},
			},
			[][]byte{
				{'.', 'S', '.'},
				{'.', '|', '.'},
			},
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			splitBeams(test.Lines)

			if !slices.EqualFunc(test.Lines, test.Want, slices.Equal) {
				t.Errorf("\ngot: \n%vwant: \n%v", toStr(test.Lines), toStr(test.Want))
			}
		})
	}
}

func TestParse(t *testing.T) {
	lines := []string{
		".S.",
		"...",
		"^.^",
	}
	want := [][]int{
		{0, 1, 0},
		{0, 0, 0},
		{-1, 0, -1},
	}

	got := parse(lines)

	if !slices.EqualFunc(got, want, slices.Equal) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestTransform(t *testing.T) {
	cases := []struct {
		Grid, Want [][]int
	}{
		{
			[][]int{
				{0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, -1, 0, 0},
				{0, 0, 0, 0, 0},
				{4, -1, 0, -1, 0},
				{0, 0, 0, 0, 0},
			},
			[][]int{
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 1, -1, 1, 0},
				{0, 1, 0, 1, 0},
				{5, -1, 2, -1, 1},
				{5, 0, 2, 0, 1},
			},
		},
		{
			[][]int{
				{0, 3, 3, 1, 0},

				{0, 0, 0, -1, 0},
			},
			[][]int{
				{0, 3, 3, 1, 0},

				{0, 3, 4, -1, 1},
			},
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			transform(test.Grid)

			if !slices.EqualFunc(test.Grid, test.Want, slices.Equal) {
				fmt.Println("GOT:")
				print(test.Grid)
				fmt.Println()
				fmt.Println("WANT:")
				print(test.Want)

				t.Error()
				// t.Errorf("got: %v, want: %v", test.Grid, test.Want)

			}
		})
	}
}
