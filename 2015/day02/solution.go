// Package day02 implements the solution to Day 1 of 2015
package day02

import (
	_ "embed"
	"fmt"

	"github.com/philipkrueck/advent-of-code/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	return fmt.Sprintf("%d", len(input))
}

func Part2(input string) string {
	return fmt.Sprintf("%d", len(input))
}

func init() {
	const year, day = 2015, 2
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
