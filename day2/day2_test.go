package day2

import (
	"slices"
	"testing"
)

var casesPart1 = []struct {
	r    Range
	want []int
}{
	{Range{11, 22}, []int{11, 22}},
	{Range{95, 115}, []int{99}},
	{Range{998, 1010}, []int{1010}},
	{Range{1188511880, 1188511890}, []int{1188511885}},
	{Range{1698522, 1698528}, []int{}},
	{Range{222220, 222224}, []int{222222}},
	{Range{446443, 446449}, []int{446446}},
	{Range{38593856, 38593862}, []int{38593859}},
	{Range{565653, 565659}, []int{}},
	{Range{1000, 1009}, []int{}},
	{Range{1, 13}, []int{11}},
}

func TestInvalidIdsPart1(t *testing.T) {
	for _, test := range casesPart1 {
		got := invalidNums1(test.r)

		if !slices.Equal(got, test.want) {
			t.Errorf("got: %v, want: %v", got, test.want)
		}
	}
}

var numDigitsCases = []struct {
	num, digits int
}{
	{1, 1},
	{9, 1},
	{11, 2},
	{99, 2},
	{111, 3},
	{999, 3},
}

func TestNumDigits(t *testing.T) {
	for _, test := range numDigitsCases {
		got := numDigits(test.num)
		want := test.digits

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
}

var invalidDigitsCases = []struct {
	desc string
	k    int
	nums Range
	want []int
}{
	// range with different number of digits
	{
		"different len on nums.low and nums.high should return empty slice",
		1,
		Range{99, 100},
		[]int{},
	},
	{
		"k not divisible by numDigits should empty slice",
		2,
		Range{10000, 20000},
		[]int{},
	},
	{
		"single digit range should return empty slice",
		1,
		Range{1, 9},
		[]int{},
	},
	{
		"valid input should return successful",
		1,
		Range{10, 60},
		[]int{11, 22, 33, 44, 55},
	},
	{
		"valid input should return successful",
		2,
		Range{8000, 8500},
		[]int{8080, 8181, 8282, 8383, 8484},
	},
	{
		"valid input should return successful",
		1,
		Range{1000, 6000},
		[]int{1111, 2222, 3333, 4444, 5555},
	},
	{
		"valid input should return successful",
		2,
		Range{565653, 565659},
		[]int{565656},
	},
	{
		"valid input should return successful",
		5,
		Range{1188511880, 1188511890},
		[]int{1188511885},
	},
	{
		"valid input should return successful",
		3,
		Range{565653, 565659},
		[]int{},
	},
}

func TestInvalidNums(t *testing.T) {
	for _, test := range invalidDigitsCases {
		got := invalidNums(test.nums, test.k)

		if !slices.Equal(got, test.want) {
			t.Errorf("got: %d, want: %d", got, test.want)
		}
	}
}

var mostKSignificantDigitsCases = []struct {
	num, k, want int
}{
	{123456, 1, 1},
	{123456, 2, 12},
	{123456, 3, 123},
	{123456, 4, 1234},
	{123456, 5, 12345},
	{123456, 6, 123456},
	{123456, 7, 123456},
	{100000, 4, 1000},
}

func TestMostKSignificantBytes(t *testing.T) {
	for _, test := range mostKSignificantDigitsCases {
		got := mostKSignificantDigits(test.num, test.k)

		if got != test.want {
			t.Errorf("got: %d, want: %d", got, test.want)
		}
	}
}

var splitRangeCases = []struct {
	in   Range
	want []Range
}{
	{
		in: Range{95, 1113},
		want: []Range{
			{95, 99}, {100, 999}, {1000, 1113},
		},
	},
	{
		in: Range{11, 55},
		want: []Range{
			{11, 55},
		},
	},
}

func TestSplitRange(t *testing.T) {
	for _, test := range splitRangeCases {
		got := splitRange(test.in)

		if !slices.Equal(got, test.want) {
			t.Errorf("got %v, want: %v", got, test.want)
		}
	}
}

var parseInputCases = []struct {
	input string
	want  []Range
}{
	{
		input: "11-55",
		want: []Range{
			{11, 55},
		},
	},
	{
		input: "11-22,95-115,998-1012,1188511880-1188511890",
		want: []Range{
			{11, 22},
			{95, 115},
			{998, 1012},
			{1188511880, 1188511890},
		},
	},
}

func TestParseInput(t *testing.T) {
	for _, test := range parseInputCases {
		got := parseInput(test.input)

		if !slices.Equal(got, test.want) {
			t.Errorf("got: %v, want: %v", got, test.want)
		}
	}
}
