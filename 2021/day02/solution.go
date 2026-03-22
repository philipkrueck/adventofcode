// Package day02 implements 2021 day 2 of Advent of Code
package day02

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

func init() {
	const day, year = 2, 2021
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	var pos geom.Point
	for line := range parse.LinesSeq(input) {
		cmd := parseCommand(line)
		switch cmd.dir {
		case "up":
			pos.Y -= cmd.units
		case "down":
			pos.Y += cmd.units
		case "forward":
			pos.X += cmd.units
		}
	}
	return strconv.Itoa(pos.X * pos.Y)
}

func Part2(input string) string {
	var pos geom.Point
	aim := 0
	for line := range parse.LinesSeq(input) {
		cmd := parseCommand(line)

		switch cmd.dir {
		case "up":
			aim -= cmd.units
		case "down":
			aim += cmd.units
		case "forward":
			pos.X += cmd.units
			pos.Y += aim * cmd.units
		}
	}

	return strconv.Itoa(pos.X * pos.Y)
}

type command struct {
	dir   string
	units int
}

func parseCommand(line string) command {
	dir, unitsStr, ok := strings.Cut(line, " ")
	if !ok {
		panic(fmt.Sprintf("invalid line format %q", line))
	}
	units, err := strconv.Atoi(unitsStr)
	if err != nil {
		panic(fmt.Sprintf("invalid units in %q: %v", line, err))
	}

	switch dir {
	case "up", "down", "forward":
		// valid
	default:
		panic(fmt.Sprintf("unknown command: %q", dir))
	}

	return command{dir, units}
}
