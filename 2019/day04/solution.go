// Package day04 implements 2019 day 4 of Advent of Code
package day04

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 4, 2019
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	start, end := parse(input)
	return strconv.Itoa(countValid(start, end, isValidPart1))
}

func Part2(input string) string {
	start, end := parse(input)
	return strconv.Itoa(countValid(start, end, isValidPart2))
}

func parse(input string) (int, int) {
	start, end, found := strings.Cut(strings.TrimSpace(input), "-")
	if !found {
		panic("invalid input")
	}
	startNum, err := strconv.Atoi(start)
	if err != nil {
		panic(err)
	}
	endNum, err := strconv.Atoi(end)
	if err != nil {
		panic(err)
	}

	return startNum, endNum
}

func countValid(start, end int, valid func(int) bool) int {
	count := 0
	for i := start; i <= end; i++ {
		if valid(i) {
			count++
		}
	}
	return count
}

func isValidPart1(num int) bool {
	s := strconv.Itoa(num)
	return isNotDecreasing(s) && hasDouble(s)
}

func isValidPart2(num int) bool {
	s := strconv.Itoa(num)
	return isNotDecreasing(s) && hasExactDouble(s)
}

func isNotDecreasing(s string) bool {
	for i := range len(s) - 1 {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

func hasDouble(s string) bool {
	for i := range len(s) - 1 {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func hasExactDouble(s string) bool {
	for i := 0; i < len(s); {
		run := 1
		for i+run < len(s) && s[i+run] == s[i] {
			run++
		}
		if run == 2 {
			return true
		}
		i += run
	}
	return false
}
