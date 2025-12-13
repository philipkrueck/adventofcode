package day3

import (
	"math"

	"github.com/philipkrueck/advent-of-code/lines"
)

func Part1() int {
	return sumLinesForKDigitNumber(2)
}

func Part2() int {
	return sumLinesForKDigitNumber(12)
}

func sumLinesForKDigitNumber(k int) int {
	r := lines.NewReader("day3/input.txt")
	grid := r.Lines()

	sum := 0
	for _, line := range grid {
		sum += largestVoltage(line, k)
	}

	return sum
}

func largestVoltage(line string, k int) int {
	voltage := 0
	maxIdx := -1

	for k > 0 {
		max := 0
		for i := maxIdx + 1; i <= len(line)-k; i++ {
			d := int(line[i] - '0')
			if d > max {
				max = d
				maxIdx = i
			}
		}
		voltage += max * int(math.Pow10(k-1))
		k--
	}

	return voltage
}
