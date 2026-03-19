// Package day01 implements 2020 day 1 of Advent of Code
package day01

import (
	_ "embed"
	"slices"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

const target = 2020

func init() {
	const day, year = 1, 2020
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	nums := parseNums(input)
	seen := make(map[int]bool, len(nums))

	for _, num := range nums {
		missing := target - num
		if seen[missing] {
			return strconv.Itoa(num * missing)
		}
		seen[num] = true
	}
	return ""
}

func Part2(input string) string {
	nums := parseNums(input)
	slices.Sort(nums)

	for i := range len(nums) - 2 {
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				return strconv.Itoa(nums[i] * nums[left] * nums[right])
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return ""
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
