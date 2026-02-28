// Package day02 implements the solutions to day 2 of advent of code 2024
package day02

import (
	"slices"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/lines"
)

func Part1() int {
	reports := loadInput()
	return countSafeReports(reports)
}

func Part2() int {
	reports := loadInput()
	return countSafeReportsDampened(reports)
}

func loadInput() [][]int {
	r := lines.NewReader("2024/day02/input.txt")
	lines := r.Lines()
	return Parse(lines)
}

func Parse(lines []string) [][]int {
	reports := make([][]int, len(lines))
	for i, line := range lines {
		nums := strings.Fields(line)
		levels := make([]int, len(nums))
		for j, n := range nums {
			level, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			levels[j] = level
		}
		reports[i] = levels
	}
	return reports
}

func countSafeReports(reports [][]int) int {
	count := 0
	for _, levels := range reports {
		if safeIncrease(levels) || safeDecrease(levels) {
			count++
		}
	}
	return count
}

func countSafeReportsDampened(reports [][]int) int {
	count := 0
	for _, levels := range reports {
		if levelsSafeDampened(levels) {
			count++
		}
	}
	return count
}

func safeIncrease(levels []int) bool {
	return levelsSafe(levels, true)
}

func safeDecrease(levels []int) bool {
	return levelsSafe(levels, false)
}

func levelsSafe(levels []int, inc bool) bool {
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		if unsafeDiff(diff, inc) {
			return false
		}
	}
	return true
}

func unsafeDiff(diff int, inc bool) bool {
	if !inc {
		diff *= -1
	}
	return diff < 1 || diff > 3
}

func levelsSafeDampened(levels []int) bool {
	for idxToSkip := range len(levels) {
		ll := slices.Concat(levels[:idxToSkip], levels[idxToSkip+1:])
		if safeIncrease(ll) || safeDecrease(ll) {
			return true
		}
	}
	return false
}
