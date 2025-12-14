package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/philipkrueck/advent-of-code/lines"
)

func Part1() int {
	r := lines.NewReader("day6/input.txt")
	lines := r.Lines()

	grid, ops := parseLines(lines)

	return operateColumns(grid, ops)
}

func parseLines(lines []string) ([][]int, []Op) {
	grid := [][]int{}
	ops := []Op{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		} else if line[0] == '+' || line[0] == '*' {
			ops = parseOps(line)
		} else {
			grid = append(grid, parseNums(line))
		}
	}
	return grid, ops
}

func parseNums(line string) []int {
	nums := []int{}

	for field := range strings.FieldsSeq(line) {
		n, err := strconv.Atoi(field)
		if err != nil {
			fmt.Printf("was trying to read %q", field)
			panic("should always be able to read in a num at this point")
		}
		nums = append(nums, n)
	}

	return nums
}

func parseOps(line string) []Op {
	ops := []Op{}

	for field := range strings.FieldsSeq(line) {
		switch field {
		case "*":
			ops = append(ops, Multiply)
		case "+":
			ops = append(ops, Add)
		}
	}

	return ops
}

func multiplyColumn(grid [][]int, j int) int {
	product := 1
	for i := range grid {
		product *= grid[i][j]
	}
	return product
}

func sumColumn(grid [][]int, j int) int {
	product := 0
	for i := range grid {
		product += grid[i][j]
	}
	return product
}

type Op int

const (
	Add = iota
	Multiply
)

func operateColumns(grid [][]int, operations []Op) int {
	total := 0

	for j, op := range operations {
		if op == Add {
			total += sumColumn(grid, j)
		} else if op == Multiply {
			total += multiplyColumn(grid, j)
		}
	}

	return total
}

func Part2() int {
	// r := lines.NewReader("day6/input.txt")
	// lines := r.Lines()

	// fmt.Println("lines: ", lines)

	return 0
}
