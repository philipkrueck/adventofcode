package day12

import (
	"testing"
)

func TestFullPresents(t *testing.T) {
	cases := []struct {
		Area Area
		Want int
	}{
		{
			Area{2, 3}, 0,
		},
		{
			Area{3, 3}, 1,
		},
		{
			Area{6, 3}, 2,
		},
		{
			Area{6, 6}, 4,
		},
		{
			Area{8, 7}, 4,
		},
		{
			Area{2, 3}, 0,
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := numFullPresents(test.Area)

			if got != test.Want {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}
