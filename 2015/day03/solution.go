package day03

import (
	_ "embed"
	"strconv"

	"github.com/philipkrueck/adventofcode/internal/geom"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	visited := make(map[geom.Point]bool, len(input))
	santa := geom.Point{}
	visited[santa] = true

	for _, c := range input {
		santa = santa.Add(direction(c))
		visited[santa] = true
	}

	return strconv.Itoa(len(visited))
}

func Part2(input string) string {
	visited := make(map[geom.Point]bool, len(input))
	santa, robo := geom.Point{}, geom.Point{}
	visited[santa] = true

	for i, c := range input {
		dir := direction(c)

		if i%2 == 0 {
			santa = santa.Add(dir)
			visited[santa] = true
		} else {
			robo = robo.Add(dir)
			visited[robo] = true
		}
	}

	return strconv.Itoa(len(visited))
}

func direction(dir rune) geom.Point {
	switch dir {
	case '^':
		return geom.Point{Y: 1}
	case 'v':
		return geom.Point{Y: -1}
	case '>':
		return geom.Point{X: 1}
	case '<':
		return geom.Point{X: -1}
	default:
		return geom.Point{}
	}
}

func init() {
	const year, day = 2015, 3
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
