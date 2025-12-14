package day5

import (
	"testing"
)

func TestCountInNumbers(t *testing.T) {
	nums := []int{1, 5, 8, 11, 17, 32}
	intervals := []Interval{
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18},
	}

	got := countNumbers(nums, intervals)
	want := 3

	if got != want {
		t.Errorf("got: %d; want: %d\n", got, want)
	}
}

func TestIntervalWidth(t *testing.T) {
	cases := []struct {
		Interval Interval
		Want     int
	}{
		{Interval{3, 5}, 3},
		{Interval{3, 10}, 8},
	}

	for _, test := range cases {
		t.Run("name", func(t *testing.T) {
			got := test.Interval.Width()

			if got != test.Want {
				t.Errorf("got: %d; want: %d\n", got, test.Want)
			}
		})
	}
}

func TestIntervalsOverlap(t *testing.T) {
	cases := []struct {
		IntervalA, IntervalB Interval
		Want                 bool
	}{
		{Interval{3, 5}, Interval{6, 8}, false},
		{Interval{3, 5}, Interval{5, 8}, true},
		{Interval{6, 8}, Interval{4, 6}, true},
		{Interval{7, 9}, Interval{4, 6}, false},
		{Interval{100, 110}, Interval{105, 109}, true},
		{Interval{105, 109}, Interval{100, 110}, true},
		{Interval{105, 109}, Interval{120, 125}, false},
	}

	for _, test := range cases {
		t.Run("name", func(t *testing.T) {
			got := test.IntervalA.Overlaps(test.IntervalB)

			if got != test.Want {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}

func TestIntervalMerge(t *testing.T) {
	cases := []struct {
		IntervalA, IntervalB, Want Interval
	}{
		{Interval{3, 5}, Interval{5, 8}, Interval{3, 8}},
		{Interval{105, 109}, Interval{100, 110}, Interval{100, 110}},
		{Interval{6, 8}, Interval{4, 6}, Interval{4, 8}},
	}

	for _, test := range cases {
		t.Run("name", func(t *testing.T) {
			got := test.IntervalA.Merge(test.IntervalB)

			if got != test.Want {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}
