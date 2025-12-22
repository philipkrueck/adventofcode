package day9

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/philipkrueck/advent-of-code/lines"
)

type Point struct {
	X, Y int
}

func Part1() int {
	r := lines.NewReader("day9/input.txt")
	lines := r.Lines()
	points := parse(lines)

	return largestArea(points)
}

func Part2() int {
	r := lines.NewReader("day9/input.txt")
	lines := r.Lines()
	points := parse(lines)

	return largestAllowedArea(points)
}

func parse(lines []string) []Point {
	points := make([]Point, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x := mustAtoi(parts[0])
		y := mustAtoi(parts[1])
		points[i] = Point{x, y}
	}
	return points
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return n
}

func largestArea(points []Point) int {
	var largest int
	for i := range len(points) - 1 {
		for j := i + 1; j < len(points); j++ {
			if a := area(points[i], points[j]); a > largest {
				largest = a
			}
		}
	}

	return largest
}

func area(a, b Point) int {
	w := math.Abs(float64(a.X-b.X)) + 1
	h := math.Abs(float64(a.Y-b.Y)) + 1

	return int(w * h)
}

func largestAllowedArea(points []Point) int {
	var largest int
	outer := findOuterArea(points)
	for i := range len(points) - 1 {
		for j := i + 1; j < len(points); j++ {
			pA, pB := points[i], points[j]
			if !allowed(pA, pB, outer) {
				continue
			}
			if a := area(pA, pB); a > largest {
				largest = a
			}
		}
	}

	return largest
}

func rectanglePoints(a, b Point) [4]Point {
	minX := min(a.X, b.X)
	minY := min(a.Y, b.Y)
	maxX := max(a.X, b.X)
	maxY := max(a.Y, b.Y)

	points := [4]Point{
		{minX, minY},
		{maxX, minY},
		{minX, maxY},
		{maxX, maxY},
	}

	return points
}

func allowed(a, b Point, outer map[Point]bool) bool {
	c := rectanglePoints(a, b)

	for j := c[0].X; j <= c[1].X; j++ {
		if outer[Point{j, c[0].Y}] || outer[Point{j, c[2].Y}] {
			return false
		}
	}

	for i := c[0].Y; i <= c[2].Y; i++ {
		if outer[Point{c[0].X, i}] || outer[Point{c[1].X, i}] {
			return false
		}
	}

	return true
}

func findOuterArea(points []Point) map[Point]bool {
	outer := make(map[Point]bool)

	for i := range points {
		var a, b, c Point
		var bIdx, cIdx int
		if i == len(points)-2 {
			bIdx = len(points) - 1
			cIdx = 0
		} else if i == len(points)-1 {
			bIdx = 0
			cIdx = 1
		} else {
			bIdx = i + 1
			cIdx = i + 2
		}
		a, b, c = points[i], points[bIdx], points[cIdx]
		currDir := determineDir(a, b)
		nextDir := determineDir(b, c)

		fmt.Printf("%v -> %v: %v\n", a, b, sPrintDir(currDir))

		for _, p := range outerBetween(a, b, currDir, nextDir) {
			outer[p] = true
		}
	}

	return outer
}

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

func sPrintDir(dir Direction) string {
	switch dir {
	case Left:
		return "left"
	case Right:
		return "right"
	case Up:
		return "up"
	case Down:
		return "down"
	}

	return "idk"
}

func outerBetween(a, b Point, dir, nextDir Direction) []Point {
	outer := []Point{}

	switch dir {
	case Right:
		y := a.Y - 1
		for x := a.X + 1; x < b.X; x++ {
			outer = append(outer, Point{x, y})
		}
	case Left:
		y := a.Y + 1
		for x := a.X - 1; x > b.X; x-- {
			outer = append(outer, Point{x, y})
		}
	case Down:
		x := a.X + 1
		for y := a.Y + 1; y < b.Y; y++ {
			outer = append(outer, Point{x, y})
		}
	case Up:
		x := a.X - 1
		for y := a.Y - 1; y > b.Y; y-- {
			outer = append(outer, Point{x, y})
		}
	}

	if dir == Right && nextDir == Down {
		outer = append(outer,
			Point{b.X, b.Y - 1},
			Point{b.X + 1, b.Y - 1},
			Point{b.X + 1, b.Y},
		)
	} else if dir == Left && nextDir == Up {
		outer = append(outer,
			Point{b.X, b.Y + 1},
			Point{b.X - 1, b.Y + 1},
			Point{b.X - 1, b.Y},
		)
	} else if dir == Down && nextDir == Left {
		outer = append(outer,
			Point{b.X + 1, b.Y},
			Point{b.X + 1, b.Y + 1},
			Point{b.X, b.Y + 1},
		)
	} else if dir == Up && nextDir == Right {
		outer = append(outer,
			Point{b.X - 1, b.Y},
			Point{b.X - 1, b.Y - 1},
			Point{b.X, b.Y - 1},
		)
	}

	return outer
}

func determineDir(a, b Point) Direction {
	if a.Y == b.Y {
		if a.X < b.X {
			return Right
		} else {
			return Left
		}
	}
	if a.X == b.X {
		if a.Y < b.Y {
			return Down
		} else {
			return Up
		}
	}

	panic("shouldn't reach")
}
