package day2

import (
	"iter"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/lines"
)

type Range struct {
	start, end int
}

func Part1() int {
	fileContents := fileToStr("day2/input.txt")
	ranges := parseInput(fileContents)
	sum := 0

	for _, r := range ranges {

		invalids := invalidNums1(r)

		for _, invalid := range invalids {
			sum += invalid
		}
	}

	return sum
}

func Part2() int {
	fileContents := fileToStr("day2/input.txt")
	ranges := parseInput(fileContents)
	sum := 0

	for _, r := range ranges {

		invalids := invalidNums2(r)

		for _, invalid := range invalids {
			sum += invalid
		}
	}

	return sum
}

func fileToStr(fileName string) string {
	lr := lines.NewReader(fileName)
	seq := lr.Next()
	next, stop := iter.Pull(seq)
	defer stop()

	value, ok := next()
	if ok {
		return value
	}

	return ""
}

func parseInput(input string) []Range {
	idRanges := strings.Split(input, ",")

	ranges := []Range{}

	for _, idRange := range idRanges {
		singleRange := strings.Split(idRange, "-")

		start, err := strconv.Atoi(singleRange[0])
		if err != nil {
			panic("Couldn't parse input. Expected properly formatted input")
		}

		end, err := strconv.Atoi(singleRange[1])
		if err != nil {
			panic("Couldn't parse input. Expected properly formatted input")
		}

		newRange := Range{start, end}
		ranges = append(ranges, newRange)
	}

	return ranges
}

func invalidNums1(r Range) []int {
	invalids := []int{}

	sameDigitRanges := splitRange(r)

	for _, sr := range sameDigitRanges {

		n := numDigits(sr.start)
		k := int(math.Ceil(float64(n) / 2))

		if n%k != 0 {
			continue // we only care about a `k` that fits exactly twice inside n
		}

		invalids = append(invalids, invalidNums(sr, k)...)
	}
	return invalids
}

func invalidNums2(r Range) []int {
	invalids := []int{}

	sameDigitRanges := splitRange(r)

	for _, sr := range sameDigitRanges {

		n := numDigits(sr.start)

		maxK := int(math.Ceil(float64(n) / 2))

		for k := 1; k <= maxK; k++ {
			if n%k != 0 {
				continue // we only care about a `k` that fits exactly twice inside n
			}

			for _, invalid := range invalidNums(sr, k) {
				if !slices.Contains(invalids, invalid) {
					invalids = append(invalids, invalid)
				}
			}

		}

	}
	return invalids
}

func numDigits(num int) int {
	str := strconv.Itoa(num)
	return len(str) // could also use log10
}

// This function assumes that the
//
// input:
// - k is the number of repeating digits
//
// conditions:
// - nums.start & nums.end have the same number of digits
// - k must be a divisor of the numDigits(low)
func invalidNums(nums Range, k int) []int {
	n := numDigits(nums.start)
	if n != numDigits(nums.end) || n == 1 {
		return []int{}
	}

	if n%k != 0 {
		return []int{}
	}

	repeatedRange := Range{
		mostKSignificantDigits(nums.start, k),
		mostKSignificantDigits(nums.end, k),
	}

	repeatedCount := n / k

	invalids := []int{}

	for i := repeatedRange.start; i <= repeatedRange.end; i++ {
		invalidStr := strings.Repeat(strconv.Itoa(i), repeatedCount)

		invalid, err := strconv.Atoi(invalidStr)
		if err != nil {
			panic("couldn't convert str for some reason")
		}

		if invalid <= nums.end && invalid >= nums.start {
			invalids = append(invalids, invalid)
		}
	}

	return invalids
}

func mostKSignificantDigits(num int, k int) int {
	digits := numDigits(num)
	exp := int(math.Max(0, float64(digits-k)))
	return num / int(math.Pow10(exp))
}

// splits into ranges with same number of decimals
func splitRange(r Range) []Range {
	ranges := []Range{}

	currDigits := numDigits(r.start)
	maxDigits := numDigits(r.end)

	for currDigits < maxDigits {
		currEnd := nines(currDigits)
		currRange := Range{r.start, currEnd}

		ranges = append(ranges, currRange)

		r.start = currEnd + 1
		currDigits++
	}

	ranges = append(ranges, r)

	return ranges
}

func nines(count int) int {
	return int(math.Pow10(count)) - 1
}
