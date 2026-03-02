// Package geom is a helper package for dealing with geometry problems.
package geom

type Point struct {
	X, Y int
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}
