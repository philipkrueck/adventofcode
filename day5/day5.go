package day5

import (
	"slices"
	"strconv"
	"strings"

	"github.com/philipkrueck/advent-of-code/lines"
)

type Interval struct {
	Min, Max int
}

func (i Interval) Contains(n int) bool {
	return n >= i.Min && n <= i.Max
}

func (a Interval) Overlaps(b Interval) bool {
	return (a.Min <= b.Min && b.Min <= a.Max) ||
		(b.Min <= a.Min && a.Min <= b.Max)
}

func (a Interval) Merge(b Interval) Interval {
	return Interval{min(a.Min, b.Min), max(a.Max, b.Max)}
}

func (i Interval) Width() int {
	return (i.Max - i.Min) + 1
}

func Part1() int {
	r := lines.NewReader("day5/input.txt")
	lines := r.Lines()

	intervals, nums := parseLines(lines)

	return countNumbers(nums, intervals)
}

func Part2() int {
	r := lines.NewReader("day5/input.txt")
	lines := r.Lines()

	intervals := parseIntervals(lines)
	sortIntervals(intervals)

	merged := mergeIntervals(intervals)
	count := countWidths(merged)

	return count
}

func countWidths(intervals []Interval) int {
	count := 0

	for _, i := range intervals {
		count += i.Width()
	}

	return count
}

func mergeIntervals(intervals []Interval) []Interval {
	merged := []Interval{}

	i, j := 0, 1

	for j < len(intervals) {
		curr, next := intervals[i], intervals[j]
		if curr.Overlaps(next) {
			curr = curr.Merge(next)

			if j == len(intervals)-1 {
				merged = append(merged, curr)
			} else {
				intervals[i] = curr
			}

			j++
		} else {
			merged = append(merged, curr)
			i = j
			j++
		}
	}

	return merged
}

func sortIntervals(intervals []Interval) {
	slices.SortFunc(intervals, func(a Interval, b Interval) int {
		return a.Min - b.Min
	})
}

func parseIntervals(lines []string) []Interval {
	intervals := []Interval{}

	i := 0

	for lines[i] != "" {
		res := strings.Split(lines[i], "-")

		min, err := strconv.Atoi(res[0])
		if err != nil {
			panic("shoudn't fail at this point")
		}
		max, err := strconv.Atoi(res[1])
		if err != nil {
			panic("shoudn't fail at this point")
		}

		interval := Interval{min, max}
		intervals = append(intervals, interval)

		i++
	}

	return intervals
}

func parseLines(lines []string) ([]Interval, []int) {
	intervals, nums := []Interval{}, []int{}

	i := 0

	for lines[i] != "" {

		res := strings.Split(lines[i], "-")

		min, err := strconv.Atoi(res[0])
		if err != nil {
			panic("shoudn't fail at this point")
		}
		max, err := strconv.Atoi(res[1])
		if err != nil {
			panic("shoudn't fail at this point")
		}

		interval := Interval{min, max}
		intervals = append(intervals, interval)

		i++
	}

	i++

	for i < len(lines) {
		num, err := strconv.Atoi(lines[i])
		if err != nil {
			panic("shouldn't panic here")
		}
		nums = append(nums, num)
		i++
	}

	return intervals, nums
}

func countNumbers(nums []int, intervals []Interval) int {
	count := 0

	for _, n := range nums {
		for _, i := range intervals {
			if i.Contains(n) {
				count++
				// fmt.Printf("%d is in %v\n", n, i)
				break
			}
			// fmt.Printf("%d is not in %v\n", n, i)
		}
	}

	return count
}
