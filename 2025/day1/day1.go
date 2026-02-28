package day1

import (
	"fmt"
	"strconv"

	"github.com/philipkrueck/adventofcode/lines"
)

func Part1() int {
	answer := 0
	dial := 50

	lineReader := lines.NewReader("2025/day1/input.txt")

	for line := range lineReader.Next() {
		direction, suffix := line[0], line[1:]
		clicks, _ := strconv.Atoi(suffix)

		clicks = clicks % 100

		if direction == 'L' {
			clicks = -clicks
		}

		dial += clicks

		if dial < 0 {
			dial += 100
		}

		if dial > 99 {
			dial -= 100
		}

		if dial == 0 {
			answer++
		}

		fmt.Println("curr dial:", dial)
	}

	return answer
}

func Part2() int {
	dial := 50
	answer := 0
	lineReader := lines.NewReader("2025/day1/input.txt")

	for line := range lineReader.Next() {
		direction, suffix := line[0], line[1:]

		clicks, _ := strconv.Atoi(suffix)

		answer = answer + (clicks / 100)
		clicks = clicks % 100

		if direction == 'L' {
			clicks = -clicks
		}

		prevDial := dial
		dial += clicks

		if dial == 0 {
			answer++
		} else if dial < 0 {
			dial += 100
			if prevDial != 0 {
				answer++
			}
		} else if dial > 99 {
			answer++
			dial -= 100
		}

		fmt.Println("curr:", dial)
		fmt.Println(line)
	}

	return answer
}
