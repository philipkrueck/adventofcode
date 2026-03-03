package day01

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/geom"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

type instruction struct {
	right  bool
	blocks int
}

func Part1(input string) string {
	pos := geom.Point{}
	dir := geom.UnitUp

	insts := parseInstructions(input)
	for _, inst := range insts {
		pos, dir = move(pos, dir, inst)
	}

	dist := pos.Manhattan(geom.Point{})
	return strconv.Itoa(dist)
}

func Part2(input string) string {
	pos := geom.Point{}
	dir := geom.UnitUp
	visited := make(map[geom.Point]bool)

	insts := parseInstructions(input)
	for _, inst := range insts {
		dir = rotate(dir, inst.right)
		for range inst.blocks {
			visited[pos] = true
			pos = pos.Add(dir)
			if visited[pos] {
				dist := pos.Manhattan(geom.Point{})
				return strconv.Itoa(dist)
			}
		}
	}

	panic("no position re-encountered")
}

func init() {
	const year, day = 2016, 1
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func parseInstructions(in string) []instruction {
	parts := strings.Split(strings.TrimSpace(in), ", ")
	insts := make([]instruction, len(parts))

	for i, r := range parts {
		var right bool
		switch r[0] {
		case 'R':
			right = true
		case 'L':
			right = false
		default:
			panic("directional char must be 'L' or 'R'")
		}

		blocks, err := strconv.Atoi(r[1:])
		if err != nil {
			panic(err)
		}

		insts[i] = instruction{right, blocks}

	}

	return insts
}

func move(pos geom.Point, dir geom.Point, inst instruction) (geom.Point, geom.Point) {
	nextDir := rotate(dir, inst.right)
	dir = nextDir
	nextPos := pos.Add(dir.Scale(inst.blocks))
	return nextPos, nextDir
}

var unitDirs = []geom.Point{geom.UnitUp, geom.UnitRight, geom.UnitDown, geom.UnitLeft}

func rotate(dir geom.Point, right bool) geom.Point {
	for i, ud := range unitDirs {
		if dir == ud {
			if right {
				return unitDirs[(i+1)%4]
			} else {
				return unitDirs[(i+3)%4]
			}
		}
	}
	panic("invalid direction")
}
