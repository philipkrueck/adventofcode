// Package day03 implements 2019 day 3 of Advent of Code
package day03

import (
	_ "embed"
	"iter"
	"maps"
	"math"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/geom"
	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 3, 2019
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	lines := parse.Lines(input)
	seen := make(map[geom.Point]bool)

	for p := range traceWire(lines[0]) {
		seen[p] = true
	}

	closest := math.MaxInt
	for p := range traceWire(lines[1]) {
		if ok := seen[p]; ok {
			dist := p.Manhattan(geom.Point{})
			if dist < closest {
				closest = dist
			}
		}
	}

	return strconv.Itoa(closest)
}

func Part2(input string) string {
	lines := parse.Lines(input)
	seen := make(map[geom.Point]int)

	maps.Insert(seen, traceWire(lines[0]))

	closest := math.MaxInt
	for p, steps := range traceWire(lines[1]) {
		if otherSteps := seen[p]; otherSteps > 0 {
			if steps+otherSteps < closest {
				closest = steps + otherSteps
			}
		}
	}

	return strconv.Itoa(closest)
}

func getDir(b byte) geom.Point {
	switch b {
	case 'U':
		return geom.Point{Y: 1}
	case 'D':
		return geom.Point{Y: -1}
	case 'R':
		return geom.Point{X: 1}
	case 'L':
		return geom.Point{X: -1}
	default:
		return geom.Point{}
	}
}

// yield every point a wire visits alongside its accumulated steps
func traceWire(wire string) iter.Seq2[geom.Point, int] {
	return func(yield func(geom.Point, int) bool) {
		curr := geom.Point{}
		steps := 0

		for inst := range strings.SplitSeq(wire, ",") {
			dir := getDir(inst[0])

			dist, err := strconv.Atoi(inst[1:])
			if err != nil {
				panic(err)
			}

			for range dist {
				curr = curr.Add(dir)
				steps++
				if !yield(curr, steps) {
					return
				}
			}
		}
	}
}
