// Package day06 implements 2015 day 6 of Advent of Code
package day06

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/geom"
	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

type op int

const (
	turnOn op = iota
	turnOff
	toggle
)

type instruction struct {
	op         op
	start, end geom.Point
}

const size = 1000

func Part1(input string) string {
	grid := make([]bool, size*size)

	apply(parseInstructions(input), func(i int, in instruction) {
		switch in.op {
		case turnOn:
			grid[i] = true
		case turnOff:
			grid[i] = false
		case toggle:
			grid[i] = !grid[i]
		}
	})

	count := 0
	for _, on := range grid {
		if on {
			count++
		}
	}

	return strconv.Itoa(count)
}

func Part2(input string) string {
	grid := make([]int, size*size)

	apply(parseInstructions(input), func(i int, in instruction) {
		switch in.op {
		case turnOn:
			grid[i]++
		case turnOff:
			if grid[i] > 0 {
				grid[i]--
			}
		case toggle:
			grid[i] += 2
		}
	})

	sum := 0
	for _, v := range grid {
		sum += v
	}

	return strconv.Itoa(sum)
}

func apply(insts []instruction, fn func(i int, in instruction)) {
	for _, instr := range insts {
		for y := instr.start.Y; y <= instr.end.Y; y++ {
			row := y * size
			for x := instr.start.X; x <= instr.end.X; x++ {
				fn(row+x, instr)
			}
		}
	}
}

func parseInstructions(input string) []instruction {
	lines := parse.Lines(input)
	insts := make([]instruction, 0, len(lines))

	for _, line := range lines {
		var instr instruction

		var remainder string
		if strings.HasPrefix(line, "toggle ") {
			instr.op = toggle
			remainder = strings.TrimPrefix(line, "toggle ")
		} else if strings.HasPrefix(line, "turn on ") {
			instr.op = turnOn
			remainder = strings.TrimPrefix(line, "turn on ")
		} else if strings.HasPrefix(line, "turn off ") {
			remainder = strings.TrimPrefix(line, "turn off ")
			instr.op = turnOff
		} else {
			panic("unknown prefix")
		}

		_, err := fmt.Sscanf(remainder, "%d,%d through %d,%d",
			&instr.start.X, &instr.start.Y,
			&instr.end.X, &instr.end.Y)
		if err != nil {
			panic(err)
		}

		insts = append(insts, instr)
	}

	return insts
}

func init() {
	const day, year = 6, 2015
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
