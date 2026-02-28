// Package day02 implements the solution to Day 1 of 2015
package day02

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/philipkrueck/advent-of-code/internal/parse"
	"github.com/philipkrueck/advent-of-code/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	sum := 0
	for _, present := range parse.Lines(input) {
		l, w, h := parseLines(present)
		s1, s2, s3 := l*w, w*h, l*h
		area := 2*s1 + 2*s2 + 2*s3
		smallest := min(s1, s2, s3)
		sum += area + smallest
	}

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	return strconv.Itoa(len(input))
}

func parseLines(line string) (int, int, int) {
	parts := strings.Split(line, "x")

	l, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	w, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}
	return l, w, h
}

func init() {
	const year, day = 2015, 2
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
