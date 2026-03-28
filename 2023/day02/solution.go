// Package day02 implements 2023 day 2 of Advent of Code
package day02

import (
	_ "embed"

	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 2, 2023
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	return "1"
}

func Part2(input string) string {
	return ""
}
