package day4

import (
	"github.com/philipkrueck/advent-of-code/lines"
)

func Part1() int {
	r := lines.NewReader("day4/input.txt")
	grid := Grid(r.Lines())

	return grid.CountAccessible()
}

func Part2() int {
	return 0
}

type Index struct {
	I, j int
}

type Grid []string

func (grid Grid) CountAccessible() int {
	var n int
	for i, line := range grid {
		for j := range line {
			if grid.IsAccessible(Index{i, j}) {
				n++
			}
		}
	}
	return n
}

func (g Grid) InBounds(idx Index) bool {
	return idx.I >= 0 && idx.j >= 0 && idx.I < len(g) && idx.j < len(g[idx.I])
}

func (g Grid) At(idx Index) rune {
	return rune(g[idx.I][idx.j])
}

func (grid Grid) IsAccessible(idx Index) bool {
	if grid.At(idx) != '@' {
		return false
	}

	neighbors := 0

	for x := idx.I - 1; x <= idx.I+1; x++ {
		for y := idx.j - 1; y <= idx.j+1; y++ {
			if x == idx.I && y == idx.j {
				continue
			}

			if grid.HasPaper(Index{x, y}) {
				neighbors++
			}
		}
	}

	return neighbors < 4
}

func (grid Grid) HasPaper(idx Index) bool {
	if !grid.InBounds(idx) {
		return false
	}

	return grid.At(idx) == '@'
}
