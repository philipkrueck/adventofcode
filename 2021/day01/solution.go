// Package day01 implements 2021 day x of Advent of Code
package day01

import (
	_ "embed"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 1, 2021
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	nums := parseNums(input)

	count := 0
	for i := range len(nums) - 1 {
		if nums[i] < nums[i+1] {
			count++
		}
	}
	return strconv.Itoa(count)
}

func Part2(input string) string {
	nums := parseNums(input)

	count := 0
	for i := range len(nums) - 3 {
		if nums[i] < nums[i+3] {
			count++
		}
	}

	return strconv.Itoa(count)
}

func parseNums(input string) []int {
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
