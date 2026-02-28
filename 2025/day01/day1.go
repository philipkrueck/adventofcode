package day1

import (
	_ "embed"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	answer := 0
	dial := 50

	lines := parse.Lines(input)

	for _, line := range lines {
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
	}

	return strconv.Itoa(answer)
}

func Part2(input string) string {
	dial := 50
	answer := 0

	lines := parse.Lines(input)
	for _, line := range lines {

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
	}

	return strconv.Itoa(answer)
}

func init() {
	const year, day = 2025, 1
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
