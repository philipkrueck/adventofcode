// Package day02 implements 2019 day 2 of Advent of Code
package day02

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 2, 2019
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	mem := parseProgram(input)
	mem[1] = 12
	mem[2] = 2

	return strconv.Itoa(run(mem))
}

func Part2(input string) string {
	program := parseProgram(input)
	mem := make([]int, len(program))
	const target = 19690720

	for noun := range 100 {
		for verb := range 100 {
			copy(mem, program)
			mem[1] = noun
			mem[2] = verb
			if run(mem) == target {
				return strconv.Itoa(100*noun + verb)
			}
		}
	}
	panic("couldn't find noun-verb combination")
}

func parseProgram(input string) []int {
	ss := strings.Split(strings.TrimSpace(input), ",")
	program := make([]int, len(ss))
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		program[i] = n
	}
	return program
}

func run(mem []int) int {
	for i := 0; i < len(mem); i += 4 {
		code := mem[i]
		if code == 99 {
			return mem[0]
		}
		op1, op2, res := mem[i+1], mem[i+2], mem[i+3]
		switch code {
		case 1:
			mem[res] = mem[op1] + mem[op2]
		case 2:
			mem[res] = mem[op1] * mem[op2]
		default:
			panic("unknown program opcode")
		}
	}
	panic("didn't encounter halt code 99")
}
