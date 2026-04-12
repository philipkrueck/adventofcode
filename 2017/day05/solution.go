// Package day05 implements 2017 day 5 of Advent of Code
package day05

import (
	_ "embed"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	return strconv.Itoa(solve(parseInput(input), false))
}

func Part2(input string) string {
	return strconv.Itoa(solve(parseInput(input), true))
}

func solve(insts []int, isPart2 bool) int {
	offset, count := 0, 0
	for offset < len(insts) {
		jump := insts[offset]

		if isPart2 && jump >= 3 {
			insts[offset]--
		} else {
			insts[offset]++
		}

		offset += jump
		count++
	}

	return count
}

func parseInput(input string) []int {
	lines := parse.Lines(input)
	insts := make([]int, 0, len(lines))

	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		insts = append(insts, n)
	}

	return insts
}

func init() {
	const day, year = 5, 2017
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
