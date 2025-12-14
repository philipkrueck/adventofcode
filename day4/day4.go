package day4

import (
	"fmt"
	"slices"
	"strings"

	"github.com/philipkrueck/advent-of-code/lines"
)

func convertToGrid(lines []string) Grid {
	grid := Grid(make([][]rune, len(lines)))

	for i, line := range lines {
		// 2. Convert the string line to a rune slice
		grid[i] = []rune(line)
	}

	return grid
}

func (g Grid) Printable() string {
	lines := []string{}

	for _, row := range g {
		lines = append(lines, string(row))
	}

	printable := "\n" + strings.Join(lines, "\n")

	return printable
}

func Part1() int {
	r := lines.NewReader("day4/input.txt")
	lines := r.Lines()

	grid := convertToGrid(lines)

	return grid.CountAccessible()
}

func Part2() int {
	r := lines.NewReader("day4/input.txt")
	lines := r.Lines()

	grid := convertToGrid(lines)

	return grid.RemovePossible()
}

type Index struct {
	I, J int
}

type Grid [][]rune

func (g Grid) CountAccessible() int {
	indices := g.AccessibleIndices()
	return len(indices)
}

func (g Grid) RemovePossible() int {
	count := 0

	prevRemoved := []Index{}
	removed := g.RemoveAccessible()

	for !slices.Equal(removed, prevRemoved) {
		fmt.Printf("Removed %v rolls of paper:\n", (len(removed)))
		fmt.Printf("%v\n\n\n", g.Printable())

		count += len(removed)
		prevRemoved = removed
		removed = g.RemoveAccessible()
	}

	return count
}

func (g Grid) RemoveAccessible() []Index {
	indices := g.AccessibleIndices()

	for _, idx := range indices {
		g.Set(idx, 'x')
	}

	return indices
}

func (g Grid) Set(idx Index, v rune) {
	g[idx.I][idx.J] = v
}

func (g Grid) AccessibleIndices() []Index {
	var indices []Index

	for i, line := range g {
		for j := range line {
			idx := Index{i, j}
			if g.IsAccessible(idx) {
				indices = append(indices, idx)
			}
		}
	}

	return indices
}

func (g Grid) InBounds(idx Index) bool {
	return idx.I >= 0 && idx.J >= 0 && idx.I < len(g) && idx.J < len(g[idx.I])
}

func (g Grid) At(idx Index) rune {
	return rune(g[idx.I][idx.J])
}

func (g Grid) IsAccessible(idx Index) bool {
	if g.At(idx) != '@' {
		return false
	}

	neighbors := 0

	for x := idx.I - 1; x <= idx.I+1; x++ {
		for y := idx.J - 1; y <= idx.J+1; y++ {
			if x == idx.I && y == idx.J {
				continue
			}

			if g.HasPaper(Index{x, y}) {
				neighbors++
			}
		}
	}

	return neighbors < 4
}

func (g Grid) HasPaper(idx Index) bool {
	if !g.InBounds(idx) {
		return false
	}

	return g.At(idx) == '@'
}
