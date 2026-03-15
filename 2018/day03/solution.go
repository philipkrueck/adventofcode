// Package day03 implements 2018 day 3 of Advent of Code
package day03

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

type rectangle struct {
	id            int
	left, top     int
	width, height int
}

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	lines := parse.Lines(input)
	recs := make([]rectangle, len(lines))
	for i, l := range lines {
		recs[i] = mustParseRectangle(l)
	}

	grid := [1028][1028]int{}

	count := 0
	for _, rec := range recs {
		for y := range rec.height {
			for x := range rec.width {
				posX, posY := rec.left+x, rec.top+y
				grid[posY][posX]++

				if grid[posY][posX] == 2 {
					count++
				}
			}
		}
	}
	return strconv.Itoa(count)
}

func Part2(input string) string {
	lines := parse.Lines(input)
	recs := make([]rectangle, len(lines))
	for i, l := range lines {
		recs[i] = mustParseRectangle(l)
	}

	hasIntersected := make([]bool, len(recs)+1)

	for i := range len(recs) - 1 {
		for j := i + 1; j < len(recs); j++ {
			if overlap(recs[i], recs[j]) {
				hasIntersected[recs[i].id] = true
				hasIntersected[recs[j].id] = true
			}
		}
	}

	for _, rec := range recs {
		if !hasIntersected[rec.id] {
			return strconv.Itoa(rec.id)
		}
	}

	return ""
}

func mustParseRectangle(line string) rectangle {
	var r rectangle
	_, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &r.id, &r.left, &r.top, &r.width, &r.height)
	if err != nil {
		panic(err)
	}
	return r
}

func overlap(r1, r2 rectangle) bool {
	return r1.left < r2.left+r2.width &&
		r2.left < r1.left+r1.width &&
		r1.top < r2.top+r2.height &&
		r2.top < r1.top+r1.height
}

func init() {
	const day, year = 3, 2018
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
