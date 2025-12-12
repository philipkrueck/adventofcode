package day3

import (
	"github.com/philipkrueck/advent-of-code/lines"
)

func Part1() int {
	lineReader := lines.NewReader("day3/input.txt")
	banks := lineReader.Lines()

	sum := 0
	for _, bank := range banks {
		sum += largestVoltage(bank, 2)
	}

	return sum
}

func Part2() int {
	return 0
}

func largestVoltage(bank string, k int) int {
	highestLeft := 0
	leftIdx := 0

	for i := 0; i <= len(bank)-k; i++ {
		digit := int(bank[i] - '0')
		if digit > highestLeft {
			highestLeft = digit
			leftIdx = i
		}
	}
	k--

	highestRight := 0

	for i := leftIdx + 1; i <= len(bank)-k; i++ {
		digit := int(bank[i]) - '0'
		if digit > highestRight {
			highestRight = digit
		}
	}

	return 10*highestLeft + highestRight
}
