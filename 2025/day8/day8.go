package day8

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/philipkrueck/advent-of-code/lines"
)

type Point struct {
	X, Y, Z int
}

type PointPair struct {
	A, B     Point
	Distance float64
}

func Part1() int {
	r := lines.NewReader("day8/input.txt")
	lines := r.Lines()
	points := parse(lines)
	pairs := pointPairs(points)
	sort(pairs)
	conns := createConns(pairs, 1000)
	c1, c2, c3 := longest(conns)
	return c1 * c2 * c3
}

func parse(lines []string) []Point {
	points := make([]Point, len(lines))

	for i, line := range lines {
		splitted := strings.Split(line, ",")

		x, err := strconv.Atoi(splitted[0])
		if err != nil {
			panic("input error")
		}
		y, err := strconv.Atoi(splitted[1])
		if err != nil {
			panic("input error")
		}
		z, err := strconv.Atoi(splitted[2])
		if err != nil {
			panic("input error")
		}

		points[i] = Point{x, y, z}
	}

	return points
}

func pointPairs(points []Point) []PointPair {
	pairs := []PointPair{}
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			pp := newPointPair(points[i], points[j])
			pairs = append(pairs, pp)
		}
	}
	return pairs
}

func distance(a, b Point) float64 {
	dx := float64(b.X - a.X)
	dy := float64(b.Y - a.Y)
	dz := float64(b.Z - a.Z)

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func newPointPair(a, b Point) PointPair {
	return PointPair{a, b, distance(a, b)}
}

func sort(pairs []PointPair) {
	slices.SortFunc(pairs, func(a, b PointPair) int {
		return cmp.Compare(a.Distance, b.Distance)
	})
}

func createConns(pairs []PointPair, maxConns int) [][]Point {
	conns := [][]Point{}
	m := int(math.Min(float64(maxConns), float64(len(pairs))))

	for i := 0; i < m; i++ {
		conns = addConn(conns, pairs[i])
	}

	return conns
}

func addConn(conns [][]Point, pp PointPair) [][]Point {
	aIdx, bIdx := -1, -1

	for i, conn := range conns {
		if slices.Contains(conn, pp.A) {
			aIdx = i
		}
		if slices.Contains(conn, pp.B) {
			bIdx = i
		}
	}

	if aIdx == -1 && bIdx == -1 {
		newConn := []Point{pp.A, pp.B}
		conns = append(conns, newConn)
	} else if aIdx != -1 && bIdx == -1 {
		conns[aIdx] = append(conns[aIdx], pp.B)
	} else if aIdx == -1 && bIdx != -1 {
		conns[bIdx] = append(conns[bIdx], pp.A)
	} else if aIdx != bIdx {
		conns[aIdx] = slices.Concat(conns[bIdx], conns[aIdx])
		conns = slices.Delete(conns, bIdx, bIdx+1)
	}

	return conns
}

func longest(conns [][]Point) (int, int, int) {
	connLengths := make([]int, len(conns))
	for i, conn := range conns {
		connLengths[i] = len(conn)
	}

	slices.SortFunc(connLengths, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	return connLengths[0], connLengths[1], connLengths[3]
}

func Part2() int {
	r := lines.NewReader("day8/input.txt")
	lines := r.Lines()
	points := parse(lines)
	pairs := pointPairs(points)
	sort(pairs)
	p := connectAll(pairs, len(points))

	return p.A.X * p.B.X
}

func connectAll(pairs []PointPair, totalPoints int) PointPair {
	conns := [][]Point{}

	for _, pair := range pairs {
		conns = addConn(conns, pair)

		if len(conns) == 1 && len(conns[0]) == totalPoints {
			return pair
		}
	}

	panic("should have connected all")
}
