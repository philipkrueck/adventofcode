// Package day02 implements 2023 day 2 of Advent of Code
package day02

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 2, 2023
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

type game struct {
	id      int
	subsets []subset
}

type subset struct {
	red, green, blue int
}

const maxRed, maxGreen, maxBlue = 12, 13, 14

func Part1(input string) string {
	sum := 0

	for line := range parse.LinesSeq(input) {
		game := parseGame(line)
		if game.isValid() {
			sum += game.id
		}

	}

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	sum := 0

	for line := range parse.LinesSeq(input) {
		game := parseGame(line)
		sum += game.minCubesPower()
	}

	return strconv.Itoa(sum)
}

func (game game) minCubesPower() int {
	var minBlue, minRed, minGreen int

	for _, subset := range game.subsets {
		minBlue = max(minBlue, subset.blue)
		minRed = max(minRed, subset.red)
		minGreen = max(minGreen, subset.green)
	}

	return minBlue * minRed * minGreen
}

func (game game) isValid() bool {
	for _, subset := range game.subsets {
		if subset.blue > maxBlue || subset.green > maxGreen || subset.red > maxRed {
			return false
		}
	}
	return true
}

func parseGame(line string) game {
	var game game

	colon := strings.IndexByte(line, ':')
	id, err := strconv.Atoi(line[5:colon])
	if err != nil {
		panic(err)
	}
	game.id = id

	subsetsStrs := strings.Split(line[colon+2:], "; ")
	game.subsets = make([]subset, 0, len(subsetsStrs))

	for _, subsetStr := range subsetsStrs {
		var subset subset
		cubes := strings.SplitSeq(subsetStr, ", ")
		for cube := range cubes {
			numStr, color, ok := strings.Cut(cube, " ")
			if !ok {
				panic("should find num and color")
			}

			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}

			switch color {
			case "red":
				subset.red = num
			case "green":
				subset.green = num
			case "blue":
				subset.blue = num
			}

		}
		game.subsets = append(game.subsets, subset)
	}

	return game
}
