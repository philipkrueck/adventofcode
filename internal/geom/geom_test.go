package geom

import (
	"fmt"
	"testing"
)

func TestManhatten(t *testing.T) {
	tests := []struct {
		p, q Point
		want int
	}{
		{Point{0, 0}, Point{0, 0}, 0},
		{Point{0, 0}, Point{2, 3}, 5},
		{Point{0, 0}, Point{-2, -3}, 5},
		{Point{0, 0}, Point{2, -3}, 5},
		{Point{-1, -1}, Point{10, 2}, 14},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v->%v", tt.p, tt.q), func(t *testing.T) {
			if got := tt.p.Manhattan(tt.q); got != tt.want {
				t.Errorf("got: %d, want: %d", got, tt.want)
			}
		})
	}
}
