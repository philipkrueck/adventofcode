// Package day03 implements 2023 day 3 of Advent of Code
package day03

import (
	_ "embed"
	"maps"
	"slices"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 3, 2023
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	lines := parse.Lines(input)

	sum := 0
	for i, line := range lines {
		for j := range len(line) {
			if line[j] == '.' || isDigit(line[j]) {
				continue
			}
			for _, n := range neighborNums(i, j, lines) {
				sum += n
			}
		}
	}

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	lines := parse.Lines(input)

	sum := 0
	for i, line := range lines {
		for j := range len(line) {
			if line[j] != '*' {
				continue
			}
			nums := neighborNums(i, j, lines)
			if len(nums) == 2 {
				sum += nums[0] * nums[1]
			}
		}
	}

	return strconv.Itoa(sum)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func neighborNums(tarI, tarJ int, lines []string) []int {
	type pos struct{ r, c int }
	nums := make(map[pos]int)

	for i := max(0, tarI-1); i <= min(len(lines)-1, tarI+1); i++ {
		for j := max(0, tarJ-1); j <= min(len(lines[i])-1, tarJ+1); j++ {
			if !isDigit(lines[i][j]) {
				continue
			}

			startCol := j
			for startCol > 0 && isDigit(lines[i][startCol-1]) {
				startCol--
			}

			p := pos{i, startCol}
			if _, exists := nums[p]; exists {
				continue
			}

			endCol := j + 1
			for endCol < len(lines[i]) && isDigit(lines[i][endCol]) {
				endCol++
			}

			num, err := strconv.Atoi(lines[i][startCol:endCol])
			if err != nil {
				panic(err)
			}

			nums[p] = num
		}
	}

	return slices.Collect(maps.Values(nums))
}
