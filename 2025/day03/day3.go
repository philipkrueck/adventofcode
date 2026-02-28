package day3

import (
	_ "embed"
	"math"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	sum := sumLinesForKDigitNumber(input, 2)
	return strconv.Itoa(sum)
}

func Part2(input string) string {
	sum := sumLinesForKDigitNumber(input, 12)
	return strconv.Itoa(sum)
}

func init() {
	const year, day = 2025, 3
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func sumLinesForKDigitNumber(input string, k int) int {
	grid := parse.Lines(input)

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
