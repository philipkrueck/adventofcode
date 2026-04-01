// Package day04 implements 2015 day 4 of Advent of Code
package day04

import (
	"crypto/md5"
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func init() {
	const day, year = 4, 2015
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}

func Part1(input string) string {
	input = parse.Lines(input)[0]
	return strconv.Itoa(findNumber(input, "00000"))
}

func Part2(input string) string {
	input = parse.Lines(input)[0]
	return strconv.Itoa(findNumber(input, "000000"))
}

func findNumber(key, prefix string) int {
	for n := 0; ; n++ {
		s := []byte(key + strconv.Itoa(n))
		hash := fmt.Sprintf("%x", md5.Sum(s))
		if strings.HasPrefix(hash, prefix) {
			return n
		}
	}
}
