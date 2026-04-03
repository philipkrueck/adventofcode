// Package day05 implements 2015 day 5 of Advent of Code
package day05

import (
	_ "embed"
	"slices"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 5, 2015
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	count := 0
	for line := range parse.LinesSeq(input) {
		if isNice1(line) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func Part2(input string) string {
	count := 0
	for line := range parse.LinesSeq(input) {
		if isNice2(line) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func isNice1(line string) bool {
	hasDouble := false
	for i := range len(line) - 1 {
		if line[i] == line[i+1] {
			hasDouble = true
		}
		if isForbiddenPair(line[i], line[i+1]) {
			return false
		}
	}

	return hasThreeVowels(line) && hasDouble
}

func hasThreeVowels(line string) bool {
	vowelCount := 0
	for i := range len(line) {
		if slices.Contains([]byte("aeiou"), line[i]) {
			vowelCount++
		}
	}
	return vowelCount >= 3
}

func isForbiddenPair(a, b byte) bool {
	d := string([]byte{a, b})
	return d == "ab" || d == "cd" || d == "pq" || d == "xy"
}

func isNice2(line string) bool {
	return hasDuplicatePair(line) && hasRepeatedWithGap(line)
}

func hasDuplicatePair(line string) bool {
	seen := make(map[[2]byte]int)

	for i := range len(line) - 1 {
		pair := [2]byte{line[i], line[i+1]}

		if prev, ok := seen[pair]; ok {
			if (prev + 2) <= i {
				return true
			}
			continue
		}
		seen[pair] = i
	}

	return false
}

func hasRepeatedWithGap(line string) bool {
	for i := range (len(line)) - 2 {
		if line[i] == line[i+2] {
			return true
		}
	}
	return false
}
