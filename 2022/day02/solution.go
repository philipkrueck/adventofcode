// Package day02 implements 2022 day 2 of Advent of Code
package day02

import (
	_ "embed"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 2, 2022
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

type shape int

const (
	rock = iota
	paper
	scissors
)

func (s shape) score() int {
	return int(s) + 1
}

func (s shape) beats() shape {
	return (s + 2) % 3
}

func (s shape) loosesTo() shape {
	return (s + 1) % 3
}

func parseShape(s byte) shape {
	switch s {
	case 'X', 'A':
		return rock
	case 'Y', 'B':
		return paper
	case 'Z', 'C':
		return scissors
	default:
		panic("unknown shape")
	}
}

func playScore(op, me shape) int {
	if op == me {
		return 3
	}

	if me.beats() == op {
		return 6
	}

	return 0
}

func Part1(input string) string {
	lines := parse.Lines(input)

	score := 0
	for _, line := range lines {
		if len(line) != 3 {
			panic("malformatted line")
		}
		p1, p2 := parseShape(line[0]), parseShape(line[2])
		score += p2.score() + playScore(p1, p2)
	}
	return strconv.Itoa(score)
}

func Part2(input string) string {
	lines := parse.Lines(input)

	score := 0
	for _, line := range lines {
		if len(line) != 3 {
			panic("malformatted line")
		}

		op := parseShape(line[0])

		var me shape
		switch line[2] {
		case 'X':
			me = op.beats()
		case 'Y':
			me = op
		case 'Z':
			me = op.loosesTo()
		default:
			panic("unknown strategy")
		}

		score += me.score() + playScore(op, me)
	}
	return strconv.Itoa(score)
}
