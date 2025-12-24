package day12

import (
	"github.com/philipkrueck/advent-of-code/lines"
)

func Part1() int {
	r := lines.NewReader("day11/test-input.txt")
	lines := r.Lines()

	return len(lines)
}

func Part2() int {
	r := lines.NewReader("day11/input.txt")
	lines := r.Lines()

	return len(lines)
}
