// Package day06 implements 2017 day 6 of Advent of Code
package day06

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

func Part1(input string) string {
	rounds, _ := findPrevious(input)
	return strconv.Itoa(rounds)
}

func Part2(input string) string {
	_, roundsAgo := findPrevious(input)
	return strconv.Itoa(roundsAgo)
}

func findPrevious(input string) (rounds int, roundsAgo int) {
	banks := parseInput(input)

	seen := make(map[string]int)
	seen[hash(banks)] = 0

	for {
		largest, largestIdx := 0, 0
		for i, v := range banks {
			if v > largest {
				largest = v
				largestIdx = i
			}
		}

		banks[largestIdx] = 0
		q := largest / len(banks)
		r := largest % len(banks)

		for i := range banks {
			banks[i] += q
		}

		for i := 1; i <= r; i++ {
			banks[(largestIdx+i)%len(banks)]++
		}

		rounds++
		hashBanks := hash(banks)

		if prevRound, ok := seen[hashBanks]; ok {
			return rounds, rounds - prevRound
		}

		seen[hashBanks] = rounds
	}
}

func parseInput(input string) []int {
	banksStr := strings.Fields(input)

	banks := make([]int, 0, len(banksStr))
	for _, bankStr := range banksStr {
		b, err := strconv.Atoi(bankStr)
		if err != nil {
			panic(err)
		}
		banks = append(banks, b)
	}
	return banks
}

func hash(banks []int) string {
	var s strings.Builder

	s.Grow(len(banks) * 4)

	for i, v := range banks {
		if i > 0 {
			s.WriteByte('_')
		}
		s.WriteString(strconv.Itoa(v))
	}

	return s.String()
}

func init() {
	const day, year = 6, 2017
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
