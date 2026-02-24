package day02

import "testing"

func TestSafeInc(t *testing.T) {
	cases := []struct {
		levels []int
		want   bool
	}{
		{
			[]int{10, 1, 2, 4, 5},
			true,
		},
		{
			[]int{1, 10, 2, 4, 5},
			true,
		},
		{
			[]int{1, 2, 4, 10, 5},
			true,
		},
	}

	for _, tt := range cases {
		got := levelsSafeDampened(tt.levels)
		if got != tt.want {
			t.Errorf("got: %v; want: %v", got, tt.want)
		}
	}
}
