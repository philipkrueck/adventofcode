package day12

import (
	"testing"
)

func TestSomething(t *testing.T) {
	cases := []struct {
		Want int
	}{}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := Part1()

			if got != test.Want {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}
