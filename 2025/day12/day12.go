package day12

import (
	"strconv"

	"github.com/philipkrueck/advent-of-code/lines"
)

// Based on Reddit comment: https://www.reddit.com/r/adventofcode/comments/1pkjynl/2025_day_12_day_12_solutions/

func Part1() int {
	r := lines.NewReader("2025/day12/input.txt")
	lines := r.Lines()
	input := parse(lines)

	return countFit(input)
}

func countFit(input Input) (count int) {
	for _, treeArea := range input.treeAreas {
		if willFit(input.presents, treeArea) {
			count++
		}
	}
	return
}

type Input struct {
	presents  [6]Present
	treeAreas []TreeArea
}

type TreeArea struct {
	area          Area
	presentCounts [6]int
}

type Present struct {
	squares int
}

type Area struct {
	w, h int
}

func countPresents(lines []string, startIdx, endIdx int) (counts int) {
	for i := startIdx; i < endIdx; i++ {
		for j := 0; j < 3; j++ {
			if lines[i][j] == '#' {
				counts++
			}
		}
	}
	return
}

func parse(lines []string) Input {
	presents := [6]Present{
		{countPresents(lines, 1, 4)},
		{countPresents(lines, 6, 9)},
		{countPresents(lines, 11, 14)},
		{countPresents(lines, 16, 19)},
		{countPresents(lines, 21, 24)},
		{countPresents(lines, 26, 29)},
	}

	treeAreas := []TreeArea{}
	for l := 30; l < 1030; l++ {
		treeAreas = append(treeAreas, parseTreeArea(lines[l]))
	}

	return Input{
		presents,
		treeAreas,
	}
}

func parseTreeArea(line string) TreeArea {
	wStr := line[:2]
	w, err := strconv.Atoi(wStr)
	if err != nil {
		panic("width assumption incorrect")
	}

	hStr := line[3:5]
	h, err := strconv.Atoi(hStr)
	if err != nil {
		panic("height assumption incorrect")
	}

	presentCounts := [6]int{
		presentCount(line, 7, 9),
		presentCount(line, 10, 12),
		presentCount(line, 13, 15),
		presentCount(line, 16, 18),
		presentCount(line, 19, 21),
		presentCount(line, 22, 24),
	}

	return TreeArea{
		Area{w, h},
		presentCounts,
	}
}

func presentCount(line string, start, end int) int {
	nStr := line[start:end]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		panic("present count index assumption is off")
	}
	return n
}

// how many 3x3 presents fit in a given area
func numFullPresents(area Area) int {
	return (area.w / 3) * (area.h / 3)
}

func (area Area) squares() int {
	return area.w * area.h
}

func willFit(presents [6]Present, treeArea TreeArea) bool {
	totalPresents := 0
	for _, count := range treeArea.presentCounts {
		totalPresents += count
	}

	if totalPresents > numFullPresents(treeArea.area) {
		// definitely impossible
		return false
	}

	totalSquares := 0
	for _, present := range presents {
		totalSquares += present.squares
	}

	if totalSquares > treeArea.area.squares() {
		// definitely impossible
		return false
	}

	// assume there is some way to fit the presents
	return true
}
