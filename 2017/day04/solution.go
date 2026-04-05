// Package day04 implements 2017 day 4 of Advent of Code
package day04

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	return solve(input, validPhrase1)
}

func Part2(input string) string {
	return solve(input, validPhrase2)
}

func solve(input string, isValid func(string) bool) string {
	count := 0
	for line := range parse.LinesSeq(input) {
		if isValid(line) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func validPhrase1(line string) bool {
	words := make(map[string]bool)
	for word := range strings.FieldsSeq(line) {
		if words[word] {
			return false
		}
		words[word] = true
	}
	return true
}

func validPhrase2(line string) bool {
	words := make(map[string]bool)
	for word := range strings.FieldsSeq(line) {
		bytes := []byte(word)
		slices.Sort(bytes)
		sortedWord := string(bytes)

		if words[sortedWord] {
			return false
		}
		words[sortedWord] = true
	}
	return true
}

func init() {
	const day, year = 4, 2017
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
