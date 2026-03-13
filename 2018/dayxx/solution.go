// Package dayxx implements 2018 day x of Advent of Code
package dayxx

import (
	_ "embed"

	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	return ""
}

func Part2(input string) string {
	return ""
}

func init() {
	const day, year = 1, 2017
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
