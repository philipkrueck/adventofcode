// Package day03 implements 2017 day 3 of Advent of Code
package day03

import (
	_ "embed"
	"math"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/geom"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	n, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		panic(err)
	}

	return strconv.Itoa(dist(n))
}

func Part1Alternative(input string) string {
	n, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		panic("Couldn't parse input")
	}

	dist := geom.Point{}.Manhattan(pos(n))

	return strconv.Itoa(dist)
}

func Part2(input string) string {
	n, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		panic("Couldn't parse input")
	}

	return strconv.Itoa(larger(n))
}

func dist(n int) int {
	if n == 1 {
		return 0
	}

	ring := 1
	side, max := 0, 0
	for {
		side = 2*ring + 1
		max = side * side
		if max >= n {
			break
		}
		ring++
	}

	bestD := math.MaxInt
	for i := range 4 {
		p := max - ring - i*(side-1)
		d := abs(p - n)
		if d < bestD {
			bestD = d
		}
	}

	return ring + bestD
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func pos(n int) (p geom.Point) {
	curr := 1

	dirs := []geom.Point{geom.UnitRight, geom.UnitUp, geom.UnitLeft, geom.UnitDown}

	dir := 0
	for turn := 0; curr < n; turn++ {
		segLen := (turn / 2) + 1

		if curr+segLen >= n {
			diff := n - curr
			p = p.Add(dirs[dir].Scale(diff))
			break
		}

		curr += segLen
		p = p.Add(dirs[dir].Scale(segLen))

		dir = (dir + 1) % len(dirs)
	}

	return p
}

func larger(n int) int {
	p := geom.Point{}
	grid := map[geom.Point]int{p: 1}

	dirs := []geom.Point{geom.UnitRight, geom.UnitUp, geom.UnitLeft, geom.UnitDown}
	adj := []geom.Point{
		{X: -1, Y: 1},
		{X: 0, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 0},
		{X: 1, Y: -1},
		{X: 0, Y: -1},
		{X: -1, Y: -1},
		{X: -1, Y: 0},
	}

	dir := 0
	for turn := 0; ; turn++ {
		segLen := (turn / 2) + 1

		for range segLen {
			p = p.Add(dirs[dir])

			sum := 0
			for _, a := range adj {
				sum += grid[p.Add(a)]
			}

			if sum > n {
				return sum
			}

			grid[p] = sum
		}
		dir = (dir + 1) % len(dirs)
	}
}

func init() {
	const day, year = 3, 2017
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
