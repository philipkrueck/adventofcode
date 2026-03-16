// Package day01 implements 2019 day 1 of Advent of Code
package day01

import (
	_ "embed"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 1, 2019
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	sum := 0
	for line := range parse.LinesSeq(input) {
		n := mustAtoi(line)
		sum += calcFuel(n)
	}
	return strconv.Itoa(sum)
}

func Part2(input string) string {
	sum := 0
	for line := range parse.LinesSeq(input) {
		n := mustAtoi(line)
		for fuel := calcFuel(n); fuel > 0; fuel = calcFuel(fuel) {
			sum += fuel
		}
	}

	return strconv.Itoa(sum)
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func calcFuel(mass int) int {
	return mass/3 - 2
}
