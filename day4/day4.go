package day4

import (
	"github.com/philipkrueck/advent-of-code/lines"
)

func Part1() int {
	r := lines.NewReader("day4/input.txt")
	lines := r.Lines()

	return sumAccessibleRolls(lines)
}

func Part2() int {
	return 0
}

func sumAccessibleRolls(grid []string) (count int) {
	for i, line := range grid {
		for j := range line {
			if isAccessible(grid, i, j) {
				count++
			}
		}
	}

	return count
}

func isPaper(grid []string, i, j int) bool {
	insideGrid := i >= 0 && j >= 0 && i < len(grid) && j < len(grid[i])

	if !insideGrid {
		return false
	}

	return grid[i][j] == '@'
}

func isAccessible(grid []string, i, j int) bool {
	if grid[i][j] != '@' {
		return false
	}

	surroundingRolls := 0

	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x == i && y == j {
				continue
			}

			if isPaper(grid, x, y) {
				surroundingRolls++
			}
		}
	}

	return surroundingRolls < 4
}
