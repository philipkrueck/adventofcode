// Package day03 implements the solution to 2016 day 3 of Advent of Code.
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

type Triangle [3]int

func Part1(input string) string {
	lines := parse.Lines(input)
	count := 0

	for _, line := range lines {
		t := parseLine(line)
		if validTri(t) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func Part2(input string) string {
	lines := parse.Lines(input)
	if len(lines)%3 != 0 {
		panic("Expected len(lines) to be divisible by 3")
	}

	count := 0
	for i := 0; i < len(lines); i += 3 {
		block := [3][3]int{
			parseLine(lines[i]),
			parseLine(lines[i+1]),
			parseLine(lines[i+2]),
		}
		tris := []Triangle{
			{block[0][0], block[1][0], block[2][0]},
			{block[0][1], block[1][1], block[2][1]},
			{block[0][2], block[1][2], block[2][2]},
		}
		for _, t := range tris {
			if validTri(t) {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}

func init() {
	const year, day = 2016, 3
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func parseLine(line string) (t Triangle) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		panic("each line should have exactly 3 numbers")
	}

	for i, part := range parts {
		a, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		t[i] = a
	}

	return t
}

func validTri(t Triangle) bool {
	return (t[0]+t[1]) > t[2] && (t[0]+t[2]) > t[1] && (t[1]+t[2]) > t[0]
}
