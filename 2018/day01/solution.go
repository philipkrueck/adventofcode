// Package day01 implements 2018 day 1 of Advent of Code
package day01

import (
	_ "embed"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	freq := 0
	for _, n := range parseInput(input) {
		freq += n
	}
	return strconv.Itoa(freq)
}

func Part2(input string) string {
	nums := parseInput(input)
	return strconv.Itoa(firstRepeatedFreq(nums))
}

func parseInput(input string) []int {
	lines := parse.Lines(input)

	nums := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	return nums
}

func firstRepeatedFreq(nums []int) int {
	freq := 0
	seen := map[int]bool{0: true}
	for {
		for _, n := range nums {
			freq += n
			if seen[freq] {
				return freq
			}
			seen[freq] = true
		}
	}
}

func init() {
	const day, year = 1, 2018
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
