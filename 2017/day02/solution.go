// Package day02 implements 2017 day 2 of Advent of Code
package day02

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	sum := 0
	for _, line := range parse.Lines(input) {
		row := parseRow(line)
		sum += rowRange(row)
	}
	return strconv.Itoa(sum)
}

func Part2(input string) string {
	sum := 0
	for _, line := range parse.Lines(input) {
		row := parseRow(line)
		sum += rowQuotient(row)
	}
	return strconv.Itoa(sum)
}

func parseRow(line string) []int {
	ss := strings.Fields(line)
	nums := make([]int, len(ss))
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}

func rowRange(row []int) int {
	min, max := row[0], row[0]
	for _, x := range row[1:] {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return max - min
}

func rowQuotient(row []int) int {
	for i, x := range row {
		for j, y := range row {
			if i == j {
				continue
			}
			if x%y == 0 {
				return x / y
			}
		}
	}
	panic("no divisible pair found")
}

func init() {
	const day, year = 2, 2017
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
