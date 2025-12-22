package day9

import (
	"maps"
	"slices"
	"testing"
)

func TestLargestArea(t *testing.T) {
	points := []Point{
		{7, 1},
		{11, 1},
		{11, 7},
		{9, 7},
		{9, 5},
		{2, 5},
		{2, 3},
		{7, 3},
	}
	want := 50

	got := largestArea(points)

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestArea(t *testing.T) {
	cases := []struct {
		X, Y Point
		Want int
	}{
		{
			Point{2, 5}, Point{11, 1}, 50,
		},
		{
			Point{2, 5}, Point{9, 7}, 24,
		},
		{
			Point{7, 1}, Point{11, 7}, 35,
		},
		{
			Point{7, 3}, Point{2, 3}, 6,
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := area(test.X, test.Y)

			if got != test.Want {
				t.Errorf("got: %d, want: %d", got, test.Want)
			}
		})
	}
}

var outer = map[Point]bool{
	{8, 0}:  true,
	{9, 0}:  true,
	{10, 0}: true,
	{11, 0}: true,
	{12, 0}: true,
	{12, 1}: true,
	{12, 2}: true,
	{12, 3}: true,
	{12, 4}: true,
	{12, 5}: true,
	{12, 6}: true,
	{12, 7}: true,
	{12, 8}: true,
	{11, 8}: true,
	{10, 8}: true,
	{9, 8}:  true,
	{8, 8}:  true,
	{8, 7}:  true,
	{8, 6}:  true,
	{7, 6}:  true,
	{6, 6}:  true,
	{5, 6}:  true,
	{4, 6}:  true,
	{3, 6}:  true,
	{2, 6}:  true,
	{1, 6}:  true,
	{1, 5}:  true,
	{1, 4}:  true,
	{1, 3}:  true,
	{1, 2}:  true,
	{2, 2}:  true,
	{3, 2}:  true,
	{4, 2}:  true,
	{5, 2}:  true,
	{6, 2}:  true,
	{6, 1}:  true,
	{6, 0}:  true,
	{7, 0}:  true,
}

var outer2 = map[Point]bool{
	{6, 9}:   true,
	{7, 9}:   true,
	{8, 9}:   true,
	{9, 9}:   true,
	{9, 10}:  true,
	{9, 11}:  true,
	{10, 11}: true,
	{11, 11}: true,
	{12, 11}: true,
	{13, 11}: true,
	{14, 11}: true,
	{14, 12}: true,
	{14, 13}: true,
	{14, 14}: true,
	{13, 14}: true,
	{12, 14}: true,
	{11, 14}: true,
	{10, 14}: true,
	{9, 14}:  true,
	{8, 14}:  true,
	{8, 15}:  true,
	{8, 16}:  true,
	{7, 16}:  true,
	{6, 16}:  true,
	{5, 16}:  true,
	{4, 16}:  true,
	{3, 16}:  true,
	{2, 16}:  true,
	{1, 16}:  true,
	{1, 15}:  true,
	{1, 14}:  true,
	{1, 13}:  true,
	{1, 12}:  true,
	{1, 11}:  true,
	{1, 10}:  true,
	{2, 10}:  true,
	{3, 10}:  true,
	{4, 10}:  true,
	{4, 9}:   true,
	{5, 9}:   true,
}

func TestAllowed(t *testing.T) {
	cases := []struct {
		A, B  Point
		Outer map[Point]bool
		Want  bool
	}{
		{
			Point{2, 3},
			Point{9, 5},
			outer,
			true,
		},
		{
			Point{7, 1},
			Point{11, 7},
			outer,
			false,
		},
		{
			Point{11, 1},
			Point{9, 7},
			outer,
			true,
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := allowed(test.A, test.B, outer)

			if got != test.Want {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}

var corners = []Point{
	{7, 1},
	{11, 1},
	{11, 7},
	{9, 7},
	{9, 5},
	{2, 5},
	{2, 3},
	{7, 3},
}

var corners2 = []Point{
	{5, 10},
	{8, 10},
	{8, 12},
	{13, 12},
	{13, 13},
	{7, 13},
	{7, 15},
	{2, 15},
	{2, 11},
	{5, 11},
}

func TestLargestAllowedArea(t *testing.T) {
	cases := []struct {
		RedPoints []Point
		Want      int
	}{
		{
			slices.Clone(corners),
			24,
		},
		{
			slices.Clone(corners2),
			30,
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := largestAllowedArea(test.RedPoints)

			if got != test.Want {
				t.Errorf("got: %v; want: %v", got, test.Want)
			}
		})
	}
}

func TestFindOuterArea(t *testing.T) {
	cases := []struct {
		RedPoints []Point
		Want      map[Point]bool
	}{
		{
			slices.Clone(corners),
			outer,
		},
		{
			slices.Clone(corners2),
			outer2,
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := findOuterArea(test.RedPoints)

			if !maps.Equal(got, test.Want) {
				t.Errorf("got: %v; want: %v", got, test.Want)
			}
		})
	}
}

func TestOuterBetween(t *testing.T) {
	cases := []struct {
		A, B         Point
		Dir, NextDir Direction
		Want         []Point
	}{
		{
			Point{1, 1},
			Point{4, 1},
			Right, Down,
			[]Point{
				{2, 0}, {3, 0}, {4, 0}, {5, 0}, {5, 1},
			},
		},
		{
			Point{1, 3},
			Point{4, 3},
			Right, Up,
			[]Point{
				{2, 2}, {3, 2},
			},
		},
		{
			Point{4, 3},
			Point{1, 3},
			Left, Down,
			[]Point{
				{3, 4}, {2, 4},
			},
		},
		{
			Point{4, 3},
			Point{1, 3},
			Left, Up,
			[]Point{
				{3, 4}, {2, 4}, {1, 4}, {0, 4}, {0, 3},
			},
		},
		{
			Point{2, 1},
			Point{2, 4},
			Down, Left,
			[]Point{
				{3, 2}, {3, 3}, {3, 4}, {3, 5}, {2, 5},
			},
		},
		{
			Point{2, 4},
			Point{2, 1},
			Up, Left,
			[]Point{
				{1, 3}, {1, 2},
			},
		},
		{
			Point{2, 4},
			Point{2, 1},
			Up, Right,
			[]Point{
				{1, 3}, {1, 2}, {1, 1}, {1, 0}, {2, 0},
			},
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := outerBetween(test.A, test.B, test.Dir, test.NextDir)

			if !slices.Equal(got, test.Want) {
				t.Errorf("got: %v; want: %v", got, test.Want)
			}
		})
	}
}

func TestDir(t *testing.T) {
	cases := []struct {
		A, B Point
		Want Direction
	}{
		{
			Point{7, 1},
			Point{11, 1},
			Right,
		},
		{
			Point{11, 1},
			Point{7, 1},
			Left,
		},
		{
			Point{9, 7},
			Point{9, 5},
			Up,
		},
		{
			Point{9, 5},
			Point{9, 7},
			Down,
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := determineDir(test.A, test.B)

			if got != test.Want {
				t.Errorf("got: %v; want: %v", got, test.Want)
			}
		})
	}
}

func TestRectanglePoints(t *testing.T) {
	cases := []struct {
		A, B Point
		Want [4]Point
	}{
		{
			Point{7, 1},
			Point{2, 3},
			[4]Point{
				{2, 1}, {7, 1}, {2, 3}, {7, 3},
			},
		},
		{
			Point{7, 1},
			Point{11, 7},
			[4]Point{
				{7, 1}, {11, 1}, {7, 7}, {11, 7},
			},
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := rectanglePoints(test.A, test.B)

			if got != test.Want {
				t.Errorf("got: %v; want: %v", got, test.Want)
			}
		})
	}
}
