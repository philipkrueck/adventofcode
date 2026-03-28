// Package day01 implements 2023 day 1 of Advent of Code
package day01

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 1, 2023
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	sum := 0
	for line := range parse.LinesSeq(input) {
		var solution int
		for i := 0; i < len(line); i++ {
			item := line[i]
			if item >= '1' && item <= '9' {
				solution = 10 * int(item-'0')
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			item := line[i]
			if item >= '1' && item <= '9' {
				solution += int(item - '0')
				break
			}
		}
		sum += solution
	}

	return strconv.Itoa(sum)
}

var nums = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func Part2(input string) string {
	sum := 0
	for line := range parse.LinesSeq(input) {
		var first, second int
		firstIdx, secondIdx := len(line), -1

		for i, numStr := range nums {
			val := (i % 9) + 1

			if idx := strings.Index(line, numStr); idx >= 0 {
				if idx < firstIdx {
					first = val
					firstIdx = idx
				}
			}
			if idx := strings.LastIndex(line, numStr); idx >= 0 {
				if idx > secondIdx {
					second = val
					secondIdx = idx
				}
			}
		}

		n := 10*first + second
		sum += n
	}
	return strconv.Itoa(sum)
}
