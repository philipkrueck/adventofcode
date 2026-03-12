// Package day01 implements 2017 day 1 of Advent of Code
package day01

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	s := strings.TrimSpace(input)
	return strconv.Itoa(captchaSum(s, 1))
}

func Part2(input string) string {
	s := strings.TrimSpace(input)
	return strconv.Itoa(captchaSum(s, len(s)/2))
}

func captchaSum(s string, step int) int {
	n := len(s)
	sum := 0

	for i := range n {
		j := (i + step) % n
		if s[i]-'0' == s[j]-'0' {
			sum += int(s[i] - '0')
		}
	}
	return sum
}

func init() {
	const day, year = 1, 2017
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
