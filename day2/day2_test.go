package day2

import (
	"reflect"
	"testing"
)

var casesPart1 = []struct {
	lower, upper int
	invalidIds   []int
}{
	{11, 22, []int{11, 22}},
	{95, 115, []int{99}},
	{998, 1010, []int{1010}},
	{1188511880, 1188511890, []int{1188511885}},
	{1698522, 1698528, []int{}},
	{222220, 222224, []int{222222}},
	{446443, 446449, []int{446446}},
	{38593856, 38593862, []int{38593859}},
	{565653, 565659, []int{}},
}

var casesPart2 = []struct {
	lower, upper int
	invalidIds   []int
}{
	{11, 22, []int{11, 22}},
	// {95, 115, []int{99, 111}},
	// {998, 1010, []int{1010}},
	// {1188511880, 1188511890, []int{1188511885}},
	// {1698522, 1698528, []int{}},
	// {222220, 222224, []int{222222}},
	// {446443, 446449, []int{446446}},
	// {38593856, 38593862, []int{38593859}},
	// {565653, 565659, []int{}},
}

func TestInvalidIdsPart1(t *testing.T) {
	for _, test := range casesPart1 {
		got := invalidIdsPart1(test.lower, test.upper)
		want := test.invalidIds

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
}

func TestInvalidIdsPart2(t *testing.T) {
	for _, test := range casesPart2 {
		got := invalidIdsPart2(test.lower, test.upper)
		want := test.invalidIds

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
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
