package day7

import (
	"fmt"
	"strings"

	"github.com/philipkrueck/adventofcode/lines"
)

func Part1() int {
	r := lines.NewReader("day7/input.txt")
	lines := r.ByteLines()

	return countSplits(lines)
}

func countSplits(lines [][]byte) int {
	count := 0

	splitBeams(lines)

	for i := 2; i < len(lines); i++ {
		prevLine := lines[i-1]
		line := lines[i]

		for j := range line {
			if line[j] == '^' && prevLine[j] == '|' &&
				line[j-1] == '|' && line[j+1] == '|' {
				count++
			}
		}
	}

	return count
}

func splitBeams(lines [][]byte) {
	for i := range len(lines) - 1 {
		line, nextLine := lines[i], lines[i+1]

		for j := range len(line) {
			ch, chb := rune(line[j]), rune(nextLine[j])
			if ch == '|' || ch == 'S' {
				switch chb {
				case '^':
					nextLine[j-1], nextLine[j+1] = '|', '|'
				case '.':
					nextLine[j] = '|'
				}
			}
		}
	}
}

func toStr(bytes [][]byte) string {
	lines := []string{}
	for _, byteLine := range bytes {
		lines = append(lines, string(byteLine))
	}
	joined := strings.Join(lines, "\n")
	return fmt.Sprintf("%s\n\n", joined)
}

func Part2() int {
	r := lines.NewReader("day7/input.txt")
	lines := r.Lines()

	grid := parse(lines)

	transform(grid)

	return sum(grid[len(grid)-1])
}

func print(grid [][]int) {
	for _, line := range grid {
		for _, n := range line {
			fmt.Printf("%02d ", n)
		}
		fmt.Println()
	}
}

func parse(lines []string) [][]int {
	grid := make([][]int, len(lines))
	for i, line := range lines {
		row := make([]int, len(line))
		for j, ch := range line {
			row[j] = chToNum(ch)
		}
		grid[i] = row
	}
	return grid
}

func chToNum(ch rune) int {
	switch ch {
	case 'S':
		return 1
	case '^':
		return -1
	default:
		return 0
	}
}

func transform(grid [][]int) {
	for i := 1; i < len(grid); i++ {
		currLine, prevLine := grid[i], grid[i-1]

		for j := range grid[i] {
			if currLine[j] == 0 && prevLine[j] > 0 {
				currLine[j] = prevLine[j]
			}
		}

		for j := range grid[i] {
			if currLine[j] == -1 {
				currLine[j-1] += prevLine[j]
				currLine[j+1] += prevLine[j]
			}
		}
	}
}

func sum(nums []int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return sum
}
