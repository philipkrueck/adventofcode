// Package day05 implements 2016 day 5 of Advent of Code
package day05

import (
	"crypto/md5"
	_ "embed"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/internal/registry"
)

//go:embed input.txt
var rawInput string

const hexChars = "0123456789abcdef"

func Part1(input string) string {
	key := strings.TrimSpace(input)
	buf := make([]byte, len(key), len(key)+20)
	copy(buf, key)

	var sb strings.Builder
	sb.Grow(8)

	for n := 0; sb.Len() < 8; n++ {
		s := strconv.AppendInt(buf[:len(key)], int64(n), 10)

		hash := md5.Sum(s)

		if hashStartsWithZeros(hash) {
			sb.WriteByte(hexChars[hash[2]])
		}
	}

	return sb.String()
}

func Part2(input string) string {
	key := strings.TrimSpace(input)
	buf := make([]byte, len(key), len(key)+20)
	copy(buf, key)

	solution := []byte("________")
	found := 0

	for n := 0; found < 8; n++ {
		s := strconv.AppendInt(buf[:len(key)], int64(n), 10)
		hash := md5.Sum(s)

		if !hashStartsWithZeros(hash) {
			continue
		}

		idx := int(hexChars[hash[2]] - '0')
		if idx < 0 || idx > 7 || solution[idx] != '_' {
			continue
		}

		solution[idx] = hexChars[hash[3]>>4]
		found++
	}

	return string(solution)
}

func hashStartsWithZeros(hash [16]byte) bool {
	// 16 byte hash == 32 char string
	return hash[0] == 0 && hash[1] == 0 && hash[2] < 16
}

func init() {
	const day, year = 5, 2016
	registry.Register(year, day, 1, Part1, rawInput)
	registry.Register(year, day, 2, Part2, rawInput)
}
