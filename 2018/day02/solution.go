// Package day02 implements 2018 day 2 of Advent of Code
package day02

import (
	_ "embed"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	lines := parse.Lines(input)
	twos, threes := 0, 0

	for _, line := range lines {
		var letters [26]int

		for i := range len(line) {
			letters[line[i]-'a']++
		}

		addTwo, addThree := 0, 0
		for _, v := range letters {
			switch v {
			case 2:
				addTwo = 1
			case 3:
				addThree = 1
			}
		}
		twos += addTwo
		threes += addThree
	}

	return strconv.Itoa(twos * threes)
}

func Part2(input string) string {
	lines := parse.Lines(input)

	for i := range len(lines) - 1 {
		for j := i + 1; j < len(lines); j++ {
			a, b := lines[i], lines[j]
			if ok, idx := differByOne(a, b); ok {
				return a[:idx] + a[idx+1:]
			}
		}
	}

	return ""
}

func differByOne(a, b string) (bool, int) {
	if len(a) != len(b) {
		return false, 0
	}

	diffs, idx := 0, -1
	for i := range len(a) {
		if a[i] != b[i] {
			diffs++
			if diffs > 1 {
				return false, 0
			}
			idx = i
		}
	}

	return diffs == 1, idx
}

func init() {
	const day, year = 2, 2018
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
