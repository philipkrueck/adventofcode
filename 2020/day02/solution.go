// Package day02 implements 2020 day 2 of Advent of Code
package day02

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 2, 2020
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

type Policy struct {
	min, max int
	b        byte
	password string
}

func (p Policy) isValidPart1() bool {
	count := strings.Count(p.password, string(p.b))
	return count >= p.min && count <= p.max
}

func (p Policy) isValidPart2() bool {
	return (p.password[p.min-1] == p.b) != (p.password[p.max-1] == p.b)
}

func Part1(input string) string {
	valids := 0

	for line := range parse.LinesSeq(input) {
		p := parsePolicy(line)
		if p.isValidPart1() {
			valids++
		}
	}
	return strconv.Itoa(valids)
}

func Part2(input string) string {
	valids := 0

	for line := range parse.LinesSeq(input) {
		p := parsePolicy(line)
		if p.isValidPart2() {
			valids++
		}
	}
	return strconv.Itoa(valids)
}

func parsePolicy(line string) Policy {
	dash := strings.IndexByte(line, '-')
	space := strings.IndexByte(line, ' ')

	min, _ := strconv.Atoi(line[0:dash])
	max, _ := strconv.Atoi(line[dash+1 : space])

	return Policy{
		min,
		max,
		line[space+1],
		strings.TrimSpace(line[space+4:]),
	}
}
