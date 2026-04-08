// Package day04 implements 2016 day 4 of Advent of Code
package day04

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/parse"
	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

type room struct {
	words    []string
	id       int
	checksum string
}

func Part1(input string) string {
	rooms := parseRooms(input)
	sum := 0
	for _, room := range rooms {
		if room.isValid() {
			sum += room.id
		}
	}
	return strconv.Itoa(sum)
}

func Part2(input string) string {
	rooms := parseRooms(input)

	for _, room := range rooms {
		if room.decrypt() == "northpole object storage" {
			return strconv.Itoa(room.id)
		}
	}
	return "no solution"
}

func (r room) isValid() bool {
	var counts [26]int

	for _, group := range r.words {
		for i := range group {
			counts[group[i]-'a']++
		}
	}

	letters := make([]byte, 0, 26)
	for i := range counts {
		letters = append(letters, byte(i)+'a')
	}

	sort.Slice(letters, func(i int, j int) bool {
		countI := counts[letters[i]-'a']
		countJ := counts[letters[j]-'a']
		if countI != countJ {
			return countI > countJ
		}
		return letters[i] < letters[j]
	})

	return r.checksum == string(letters[:5])
}

func (r room) decrypt() string {
	shift := byte(r.id % 26)

	totalLen := 0
	for _, word := range r.words {
		totalLen += len(word)
	}
	totalLen += len(r.words) - 1

	var sb strings.Builder
	sb.Grow(totalLen)

	for i, word := range r.words {
		if i > 0 {
			sb.WriteByte(' ')
		}
		for j := range word {
			curr := word[j] - 'a'
			shifted := (curr + shift) % 26
			sb.WriteByte('a' + shifted)
		}
	}

	return sb.String()
}

func parseRooms(input string) []room {
	lines := parse.Lines(input)
	rooms := make([]room, 0, len(lines))
	for _, line := range lines {
		rooms = append(rooms, parseRoom(line))
	}
	return rooms
}

func parseRoom(line string) room {
	letterEndIdx := strings.IndexByte(line, '[')

	splits := strings.Split(line[:letterEndIdx], "-")

	roomID, err := strconv.Atoi(splits[len(splits)-1])
	if err != nil {
		panic(err)
	}

	return room{
		id:       roomID,
		words:    splits[:len(splits)-1],
		checksum: line[letterEndIdx+1 : letterEndIdx+6],
	}
}

func init() {
	const day, year = 4, 2016
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
