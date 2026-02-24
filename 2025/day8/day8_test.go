package day8

import (
	"math"
	"slices"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		Lines []string
		Want  []Point
	}{
		{
			[]string{
				"162,817,812",
				"57,618,57",
				"906,360,560",
			},
			[]Point{
				{162, 817, 812},
				{57, 618, 57},
				{906, 360, 560},
			},
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := parse(test.Lines)
			if !slices.Equal(got, test.Want) {
				t.Errorf("got: %d; want: %d", got, test.Want)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	cases := []struct {
		A, B Point
		Want float64
	}{
		{
			Point{0, 1, 0},
			Point{0, 0, 0},
			1,
		},
		{
			Point{0, 0, 0},
			Point{3, 4, 0},
			5,
		},
		{
			Point{0, 0, 0},
			Point{1, 1, 1},
			math.Sqrt(3),
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := distance(test.A, test.B)

			if got != test.Want {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}

func TestPointPairs(t *testing.T) {
	points := []Point{
		{0, 0, 0},
		{3, 4, 0},
		{0, 0, 12},
		{3, 4, 12},
	}
	want := []PointPair{
		{Point{0, 0, 0}, Point{3, 4, 0}, 5},
		{Point{0, 0, 0}, Point{0, 0, 12}, 12},
		{Point{0, 0, 0}, Point{3, 4, 12}, 13},
		{Point{3, 4, 0}, Point{0, 0, 12}, 13},
		{Point{3, 4, 0}, Point{3, 4, 12}, 12},
		{Point{0, 0, 12}, Point{3, 4, 12}, 5},
	}
	got := pointPairs(points)

	if !slices.Equal(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSortPairs(t *testing.T) {
	pairs := []PointPair{
		{Point{3, 4, 0}, Point{3, 4, 12}, 12},
		{Point{3, 4, 0}, Point{0, 0, 12}, 13},
		{Point{0, 0, 0}, Point{0, 0, 12}, 12},
		{Point{0, 0, 0}, Point{3, 4, 0}, 5},
		{Point{0, 0, 12}, Point{3, 4, 12}, 5},
		{Point{0, 0, 0}, Point{3, 4, 12}, 13},
	}
	want := []PointPair{
		{Point{0, 0, 0}, Point{3, 4, 0}, 5},
		{Point{0, 0, 12}, Point{3, 4, 12}, 5},
		{Point{3, 4, 0}, Point{3, 4, 12}, 12},
		{Point{0, 0, 0}, Point{0, 0, 12}, 12},
		{Point{3, 4, 0}, Point{0, 0, 12}, 13},
		{Point{0, 0, 0}, Point{3, 4, 12}, 13},
	}

	sort(pairs)

	if !slices.Equal(pairs, want) {
		t.Errorf("got: %v, want: %v", pairs, want)
	}
}

func TestCreateConnections(t *testing.T) {
	pairs := []PointPair{
		{Point{0, 0, 0}, Point{3, 4, 0}, 5},
		{Point{0, 0, 12}, Point{3, 4, 12}, 5},
		{Point{3, 4, 0}, Point{3, 4, 12}, 12},
		{Point{0, 0, 0}, Point{0, 0, 12}, 12},
		{Point{3, 4, 0}, Point{0, 0, 12}, 13},
		{Point{0, 0, 0}, Point{3, 4, 12}, 13},
	}
	want := [][]Point{
		{{0, 0, 0}, {3, 4, 0}, {0, 0, 12}, {3, 4, 12}},
	}

	got := createConns(pairs, 1000)

	if !slices.EqualFunc(got, want, func(a, b []Point) bool {
		return len(a) == len(b)
	}) {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestAddConnection(t *testing.T) {
	cases := []struct {
		Name  string
		Conns [][]Point
		Pp    PointPair
		Want  [][]Point
	}{
		{
			"New connection",
			[][]Point{
				{{0, 0, 0}, {3, 4, 0}},
			},
			PointPair{
				Point{0, 0, 12}, Point{3, 4, 12}, 13,
			},
			[][]Point{
				{{0, 0, 0}, {3, 4, 0}},
				{{0, 0, 12}, {3, 4, 12}},
			},
		},
		{
			"Add B to Connection of A",
			[][]Point{
				{{0, 0, 0}, {3, 4, 0}},
			},
			PointPair{
				Point{3, 4, 0}, Point{3, 4, 12}, 12,
			},
			[][]Point{
				{{0, 0, 0}, {3, 4, 0}, {3, 4, 12}},
			},
		},
		{
			"Add A to Connection of B",
			[][]Point{
				{{0, 0, 0}, {3, 4, 0}},
			},
			PointPair{
				Point{3, 4, 12}, Point{3, 4, 0}, 12,
			},
			[][]Point{
				{{0, 0, 0}, {3, 4, 0}, {3, 4, 12}},
			},
		},
		{
			"Join existing connections of A & B",
			[][]Point{
				{{0, 0, 0}, {3, 4, 0}},
				{{0, 0, 12}, {3, 4, 12}},
			},
			PointPair{
				Point{3, 4, 12}, Point{3, 4, 0}, 12,
			},
			[][]Point{
				{
					{0, 0, 0},
					{3, 4, 0},
					{0, 0, 12},
					{3, 4, 12},
				},
			},
		},
		{
			"A and B are in the same connection",
			[][]Point{
				{{0, 0, 0}, {3, 4, 0}},
			},
			PointPair{
				Point{0, 0, 0}, Point{3, 4, 0}, 5,
			},
			[][]Point{
				{{0, 0, 0}, {3, 4, 0}},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := addConn(test.Conns, test.Pp)

			if !slices.EqualFunc(got, test.Want, slices.Equal) {
				t.Errorf("got: %d, want: %d", got, test.Want)
			}
		})
	}
}
