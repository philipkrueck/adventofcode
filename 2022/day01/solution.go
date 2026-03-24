// Package day01 implements 2022 day 1 of Advent of Code
package day01

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 1, 2022
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func parseElves(input string) []int {
	input = strings.TrimSpace(input)
	blocks := strings.Split(input, "\n\n")

	elves := make([]int, 0, len(blocks))

	for _, block := range blocks {
		var sum int
		for line := range strings.SplitSeq(block, "\n") {
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			sum += n
		}
		elves = append(elves, sum)
	}

	return elves
}

func Part1(input string) string {
	elves := parseElves(input)
	return strconv.Itoa(slices.Max(elves))
}

func Part2(input string) string {
	elves := parseElves(input)

	slices.Sort(elves)
	n := len(elves)

	top3 := elves[n-1] + elves[n-2] + elves[n-3]

	return strconv.Itoa(top3)
}
