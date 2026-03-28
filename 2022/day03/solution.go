// Package day03 implements 2022 day 3 of Advent of Code
package day03

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
	const day, year = 3, 2022
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	sum := 0
	for line := range parse.LinesSeq(input) {
		sum += linePriority(line)
	}
	return strconv.Itoa(sum)
}

func Part2(input string) string {
	lines := parse.Lines(input)

	sum := 0
	for i := 0; i < len(lines)-2; i += 3 {
		for j := range lines[i] {
			item := lines[i][j]
			if strings.IndexByte(lines[i+1], item) >= 0 &&
				strings.IndexByte(lines[i+2], item) >= 0 {
				sum += priority(item)
				break
			}
		}
	}
	return strconv.Itoa(sum)
}

func linePriority(line string) int {
	half := len(line) / 2
	left, right := line[:half], line[half:]
	for i := range left {
		if strings.IndexByte(right, left[i]) >= 0 {
			return priority(left[i])
		}
	}
	panic("each line should have a shared item")
}

func priority(b byte) int {
	if b >= 'a' && b <= 'z' {
		return int(b-'a') + 1
	}
	return int(b-'A') + 27
}
