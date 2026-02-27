// Package day01 implements the solution to Day 1 of 2015.
package day01

import (
	_ "embed"
	"strconv"

	"github.com/philipkrueck/advent-of-code/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	floor := 0
	for _, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return strconv.Itoa(floor)
}

func Part2(input string) string {
	floor := 0
	for i, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return strconv.Itoa(i + 1)
		}
	}
	panic("santa never reaches basement")
}

func init() {
	const year, day = 2015, 1
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
