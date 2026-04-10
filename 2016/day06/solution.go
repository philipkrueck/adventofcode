// Package day06 implements 2016 day 6 of Advent of Code
package day06

import (
	_ "embed"
	"math"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	colLetters := getColFrequencies(input)

	var sb strings.Builder
	sb.Grow(len(colLetters))

	for _, lets := range colLetters {
		maxCount, maxIdx := 0, 0
		for idx, count := range lets {
			if count > maxCount {
				maxCount = count
				maxIdx = idx
			}
		}
		sb.WriteByte('a' + byte(maxIdx))
	}

	return sb.String()
}

func Part2(input string) string {
	colLetters := getColFrequencies(input)

	var sb strings.Builder
	sb.Grow(len(colLetters))

	for _, lets := range colLetters {
		minCount, minIdx := math.MaxInt, 0
		for idx, count := range lets {
			if count > 0 && count < minCount {
				minCount = count
				minIdx = idx
			}
		}
		sb.WriteByte('a' + byte(minIdx))
	}

	return sb.String()
}

func getColFrequencies(input string) [][26]int {
	lines := parse.Lines(input)
	mesLen := len(lines[0])
	colLetters := make([][26]int, mesLen)

	for _, line := range lines {
		for col := range mesLen {
			colLetters[col][line[col]-'a']++
		}
	}

	return colLetters
}

func init() {
	const day, year = 6, 2016
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
