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
		switch op {
		case Add:
			total += sumColumn(grid, j)
		case Multiply:
			total += multiplyColumn(grid, j)
		}
	}

	return total
}

func Part2() int {
	r := lines.NewReader("day6/input.txt")
	lines := r.Lines()

	grid, ops := parse(lines)

	return operateGrid(grid, ops)
}

func parse(lines []string) ([][][]int, []Op) {
	grid := [][][]int{}
	lastIdx := len(lines) - 1
	opsLine := lines[lastIdx]
	ops := parseOps(opsLine)

	digitLengths := parseDigitLengths(opsLine)

	for i := range lastIdx {
		parsedLine := parseLine(lines[i], digitLengths)
		grid = append(grid, parsedLine)
	}

	return grid, ops
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

func parseLine(line string, digitLengths []int) [][]int {
	numLine := [][]int{}

	j := 0

	for _, column := range digitLengths {
		num := []int{}
		for range column {
			ch := rune(line[j])
			if ch == ' ' {
				num = append(num, 0)
			} else {
				d := int(ch) - '0'
				num = append(num, d)
			}
			j++
		}
		numLine = append(numLine, num)
		j++ // skip whitespace between columns
	}
	return numLine
}

func parseDigitLengths(line string) []int {
	lengths := []int{}

	length := 0
	for j := 1; j < len(line); j++ {
		if line[j] == ' ' {
			length++
		} else {
			lengths = append(lengths, length)
			length = 0
		}
	}
	lengths = append(lengths, length+1)

	return lengths
}

func convertColToCephs(nums [][]int) []int {
	cephs := []int{}

	for j := range len(nums[0]) {
		ceph := 0

		for i := range len(nums) {
			n := nums[i][j]
			if n == 0 {
				continue
			}
			ceph = (ceph * 10) + n
		}
		cephs = append(cephs, ceph)
	}

	return cephs
}

func sumNums(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func multiplyNums(nums []int) int {
	prod := 1
	for _, n := range nums {
		prod *= n
	}
	return prod
}

func gridToCephs(grid [][][]int) [][]int {
	cephs := [][]int{}
	colCount := len(grid[0])
	rowCount := len(grid)

	for j := range colCount {
		col := [][]int{}

		for i := range rowCount {
			col = append(col, grid[i][j])
		}

		ceph := convertColToCephs(col)
		cephs = append(cephs, ceph)
	}

	return cephs
}

func operateGrid(grid [][][]int, ops []Op) int {
	total := 0

	cephs := gridToCephs(grid)

	for i, op := range ops {
		switch op {
		case Add:
			total += sumNums(cephs[i])
		case Multiply:
			total += multiplyNums(cephs[i])
		}
	}

	return total
}
