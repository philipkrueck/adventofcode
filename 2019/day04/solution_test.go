package day04

import (
	"strconv"
	"testing"
)

func TestIsValidPart1(t *testing.T) {
	tests := []struct {
		in   int
		want bool
	}{
		{111111, true},
		{223450, false},
		{123789, false},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.in), func(t *testing.T) {
			if got := isValidPart1(tt.in); got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestIsValidPart2(t *testing.T) {
	tests := []struct {
		in   int
		want bool
	}{
		{112233, true},
		{123444, false},
		{123789, false},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.in), func(t *testing.T) {
			if got := isValidPart2(tt.in); got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}
