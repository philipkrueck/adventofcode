// Package day03 implements 2020 day 3 of Advent of Code
package day03

import (
	_ "embed"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 3, 2020
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	lines := parse.Lines(input)
	return strconv.Itoa(countTrees(lines, 3, 1))
}

func Part2(input string) string {
	lines := parse.Lines(input)

	slopes := []struct {
		dx, dy int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	total := 1
	for _, slope := range slopes {
		total *= countTrees(lines, slope.dx, slope.dy)
	}
	return strconv.Itoa(total)
}

func countTrees(lines []string, moveX, moveY int) int {
	x, y := 0, 0
	count := 0

	for y < len(lines) {
		if lines[y][x] == '#' {
			count++
		}
		x = (x + moveX) % len(lines[y])
		y += moveY
	}
	return count
}
