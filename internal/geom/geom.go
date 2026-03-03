// Package geom is a helper package for dealing with geometry problems.
package geom

import (
	"fmt"
)

type Point struct {
	X, Y int
}

var (
	UnitUp    = Point{Y: 1}
	UnitDown  = Point{Y: -1}
	UnitRight = Point{X: 1}
	UnitLeft  = Point{X: -1}
)

func (p Point) String() string {
	return fmt.Sprintf("(%v,%v)", p.X, p.Y)
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Scale(f int) Point {
	return Point{f * p.X, f * p.Y}
}

func (p Point) Manhattan(q Point) int {
	return abs(q.X-p.X) + abs(q.Y-p.Y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
