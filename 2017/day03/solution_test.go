package day03

import (
	"fmt"
	"testing"

	"github.com/philipkrueck/adventofcode/internal/geom"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"1", "0"},
		{"12", "3"},
		{"23", "2"},
		{"1024", "31"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part1(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}

func TestPart1Alternative(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"1", "0"},
		{"12", "3"},
		{"23", "2"},
		{"1024", "31"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part1Alternative(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}

func TestPos(t *testing.T) {
	tests := []struct {
		n   int
		pos geom.Point
	}{
		{1, geom.Point{X: 0, Y: 0}},
		{2, geom.Point{X: 1, Y: 0}},
		{3, geom.Point{X: 1, Y: 1}},
		{4, geom.Point{X: 0, Y: 1}},
		{5, geom.Point{X: -1, Y: 1}},
		{6, geom.Point{X: -1, Y: 0}},
		{7, geom.Point{X: -1, Y: -1}},
		{8, geom.Point{X: 0, Y: -1}},
		{9, geom.Point{X: 1, Y: -1}},
		{10, geom.Point{X: 2, Y: -1}},
		{11, geom.Point{X: 2, Y: 0}},
		{12, geom.Point{X: 2, Y: 1}},
		{13, geom.Point{X: 2, Y: 2}},
		{17, geom.Point{X: -2, Y: 2}},
		{21, geom.Point{X: -2, Y: -2}},
		{26, geom.Point{X: 3, Y: -2}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n_%v", tt.n), func(t *testing.T) {
			if got := pos(tt.n); got != tt.pos {
				t.Errorf("got: %v; want: %v", got, tt.pos)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"1", "2"},
		{"2", "4"},
		{"4", "5"},
		{"5", "10"},
		{"23", "25"},
		{"54", "57"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part2(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}
