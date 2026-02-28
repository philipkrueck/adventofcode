// Package day03 implements the solutions to day 3 of advent of code 2024
package day03

import (
	"regexp"
	"strconv"

	"github.com/philipkrueck/adventofcode/lines"
)

var (
	// matches mul(a,b) and captures a,b
	mulRe = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// matches do(), don't() or mul(a,b)
	instrRe = regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d)+,(\d)+\)`)

	// matches and captures numbers
	numRe = regexp.MustCompile(`(\d+)`)
)

func Part1() int {
	r := lines.NewReader("2024/day03/input.txt")
	lines := r.Lines()

	sum := 0
	for _, l := range lines {
		sum += parseLine(l)
	}

	return sum
}

func Part2() int {
	r := lines.NewReader("2024/day03/input.txt")
	lines := r.Lines()

	sum, do := 0, true
	for _, l := range lines {
		lineSum, nextDo := parseLineSwitch(l, do)
		do = nextDo
		sum += lineSum
	}

	return sum
}

func parseLine(in string) int {
	matches := mulRe.FindAllStringSubmatch(in, -1)

	sum := 0
	for _, m := range matches {
		a, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		sum += a * b
	}
	return sum
}

func parseLineSwitch(in string, do bool) (int, bool) {
	instructions := instrRe.FindAllString(in, -1)

	sum := 0
	for _, instr := range instructions {
		switch instr {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			if !do {
				continue
			}
			ops := numRe.FindAllString(instr, -1)

			a, err := strconv.Atoi(ops[0])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(ops[1])
			if err != nil {
				panic(err)
			}

			sum += a * b
		}
	}

	return sum, do
}
