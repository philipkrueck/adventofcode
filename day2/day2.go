package day2

import (
	"fmt"
	"iter"
	"math"
	"strconv"
	"strings"

	"github.com/philipkrueck/advent-of-code/lines"
)

type ProductRange struct {
	lower, upper int
}

func Part1() int {
	return CommonLogic(invalidIdsPart1)
}

func Part2() int {
	return CommonLogic(invalidIdsPart2)
}

func CommonLogic(invalidFilter func(int, int) []int) int {
	fileContents := fileToStr("day2/input.txt")
	fmt.Println("fileContents:", fileContents)

	idRanges := strings.Split(fileContents, ",")

	fmt.Println("idRanges", idRanges)

	productRanges := []ProductRange{}

	for _, idRange := range idRanges {
		productRange := strings.Split(idRange, "-")

		lower, _ := strconv.Atoi(productRange[0])
		upper, _ := strconv.Atoi(productRange[1])

		productRanges = append(productRanges, ProductRange{
			lower,
			upper,
		},
		)
	}

	fmt.Println("productRanges", productRanges)

	sum := 0
	for _, productRange := range productRanges {
		for _, productId := range invalidFilter(productRange.lower, productRange.upper) {
			fmt.Println("id:", productId)
			sum += productId
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

func invalidIdsPart1(lower, upper int) []int {
	invalids := []int{}

	idStr := strconv.Itoa(lower)
	idDigits := len(idStr)

	if idDigits%2 != 0 {
		nextID := int(math.Pow(10, float64(idDigits)))
		idStr = strconv.Itoa(nextID)
	}

	middle := len(idStr) / 2

	leftStr := idStr[0:middle]
	left, _ := strconv.Atoi(leftStr)
	potentialInvalid, _ := strconv.Atoi(leftStr + leftStr)

	for potentialInvalid <= upper {
		if potentialInvalid >= lower {
			invalids = append(invalids, potentialInvalid)
		}

		left += 1
		leftStr = strconv.Itoa(left)
		potentialInvalid, _ = strconv.Atoi(leftStr + leftStr)
	}

	return invalids
}

func numDigits(num int) int {
	str := strconv.Itoa(num)
	return len(str)
}

func invalidIdsPart2(lower, upper int) []int {
	invalids := []int{}

	idStr := strconv.Itoa(lower)
	idDigits := len(idStr)

	if idDigits%2 != 0 {
		nextID := int(math.Pow(10, float64(idDigits)))
		idStr = strconv.Itoa(nextID)
	}

	middle := len(idStr) / 2

	leftStr := idStr[0:middle]
	left, _ := strconv.Atoi(leftStr)
	potentialInvalid, _ := strconv.Atoi(leftStr + leftStr)

	for potentialInvalid <= upper {
		if potentialInvalid >= lower {
			invalids = append(invalids, potentialInvalid)
		}

		left += 1
		leftStr = strconv.Itoa(left)
		potentialInvalid, _ = strconv.Atoi(leftStr + leftStr)
	}

	return invalids
}
