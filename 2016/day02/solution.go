package day02

import (
	_ "embed"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	lines := parse.Lines(input)

	keypad := [3][3]byte{
		{'1', '2', '3'},
		{'4', '5', '6'},
		{'7', '8', '9'},
	}

	row, col := 1, 1
	code := make([]byte, len(lines))

	for i, line := range lines {
		for _, c := range line {
			switch c {
			case 'U':
				if row > 0 {
					row--
				}
			case 'D':
				if row < 2 {
					row++
				}
			case 'L':
				if col > 0 {
					col--
				}
			case 'R':
				if col < 2 {
					col++
				}
			default:
				panic("invalid direction")
			}
		}
		code[i] = keypad[row][col]
	}

	return string(code)
}

func Part2(input string) string {
	lines := parse.Lines(input)

	keypad := [5][5]byte{
		{' ', ' ', '1', ' ', ' '},
		{' ', '2', '3', '4', ' '},
		{'5', '6', '7', '8', '9'},
		{' ', 'A', 'B', 'C', ' '},
		{' ', ' ', 'D', ' ', ' '},
	}

	row, col := 2, 0
	code := make([]byte, len(lines))

	for i, line := range lines {
		for _, c := range line {
			switch c {
			case 'U':
				if row > 0 && keypad[row-1][col] != ' ' {
					row--
				}
			case 'D':
				if row < 4 && keypad[row+1][col] != ' ' {
					row++
				}
			case 'L':
				if col > 0 && keypad[row][col-1] != ' ' {
					col--
				}
			case 'R':
				if col < 4 && keypad[row][col+1] != ' ' {
					col++
				}
			default:
				panic("invalid direction")
			}
		}
		code[i] = keypad[row][col]
	}

	return string(code)
}

func init() {
	const year, day = 2016, 2
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
