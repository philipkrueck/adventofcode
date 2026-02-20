package day01

import (
	"slices"
	"strconv"
	"strings"

	"github.com/philipkrueck/advent-of-code/lines"
)

func Part1() int {
	l1, l2 := loadInput()

	slices.Sort(l1)
	slices.Sort(l2)

	distance := 0
	for i := range l1 {
		distance += abs(l1[i] - l2[i])
	}
	return distance
}

func Part2() int {
	l1, l2 := loadInput()

	freq := make(map[int]int, len(l2))
	for _, el := range l2 {
		freq[el]++
	}

	similarity := 0
	for _, el := range l1 {
		similarity += el * freq[el]
	}

	return similarity
}

func loadInput() ([]int, []int) {
	r := lines.NewReader("2024/day01/input.txt")
	lines := r.Lines()
	return parse(lines)
}

func parse(lines []string) ([]int, []int) {
	l1, l2 := make([]int, len(lines)), make([]int, len(lines))

	for i, line := range lines {
		nums := strings.Fields(line)
		if len(nums) != 2 {
			panic("every line should have exactly 2 nums")
		}

		n1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		l1[i] = n1
		l2[i] = n2
	}

	return l1, l2
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
