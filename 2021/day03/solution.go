// Package day03 implements 2021 day 3 of Advent of Code
package day03

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

func init() {
	const day, year = 3, 2021
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	lines := parse.Lines(input)
	width := len(lines[0])

	var epsilon, gamma strings.Builder

	epsilon.Grow(width)
	gamma.Grow(width)

	for j := range width {
		count := 0
		for i := range len(lines) {
			if lines[i][j] == '1' {
				count++
			}
		}

		if count*2 >= len(lines) {
			epsilon.WriteByte('1')
			gamma.WriteByte('0')
		} else {
			epsilon.WriteByte('0')
			gamma.WriteByte('1')
		}
	}

	epsilonDec := binToDecimal(epsilon.String())
	gammaDec := binToDecimal(gamma.String())

	return strconv.Itoa(epsilonDec * gammaDec)
}

func Part2(input string) string {
	lines := parse.Lines(input)

	oxygen := search(slices.Clone(lines), true)
	co2 := search(lines, false)

	oxygenDec := binToDecimal(oxygen)
	co2Dec := binToDecimal(co2)

	return strconv.Itoa(int(oxygenDec) * int(co2Dec))
}

func search(lines []string, mostCommon bool) string {
	width := len(lines[0])
	for j := range width {
		if len(lines) == 1 {
			break
		}

		ones := 0
		for _, line := range lines {
			if line[j] == '1' {
				ones++
			}
		}

		keepOne := 2*ones >= len(lines)
		if !mostCommon {
			keepOne = !keepOne
		}

		var toKeep byte = '0'
		if keepOne {
			toKeep = '1'
		}

		n := 0
		for _, line := range lines {
			if line[j] == toKeep {
				lines[n] = line
				n++
			}
		}

		lines = lines[:n]
	}

	return lines[0]
}

func binToDecimal(s string) int {
	res, err := strconv.ParseInt(s, 2, 0)
	if err != nil {
		panic(err)
	}
	return int(res)
}
