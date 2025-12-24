package day10

import (
	"slices"
	"testing"
)

var testConfigsRaw []string = []string{
	"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
	"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
	"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
}

var testConfigs []MachineConfig = []MachineConfig{
	{
		[]bool{false, true, true, false},
		[][]int{{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1}},
		[]int{3, 5, 4, 7},
	},
	{
		[]bool{false, false, false, true, false},
		[][]int{{0, 2, 3, 4}, {2, 3}, {0, 4}, {0, 1, 2}, {1, 2, 3, 4}},
		[]int{7, 5, 12, 7, 2},
	},
	{
		[]bool{false, true, true, true, false, true},
		[][]int{{0, 1, 2, 3, 4}, {0, 3, 4}, {0, 1, 2, 4, 5}, {1, 2}},
		[]int{10, 11, 11, 5, 10, 5},
	},
}

func TestParseLine(t *testing.T) {
	cases := []struct {
		Line string
		Want MachineConfig
	}{
		{
			testConfigsRaw[0],
			testConfigs[0],
		},
		{
			testConfigsRaw[1],
			testConfigs[1],
		},
		{
			testConfigsRaw[2],
			testConfigs[2],
		},
		{
			"[..##.] (0,3,4) (0,2) (1,3,4) (0,3) (0,2,3,4) (0,1,2,4) {67,15,36,39,35}",
			MachineConfig{
				[]bool{false, false, true, true, false},
				[][]int{{0, 3, 4}, {0, 2}, {1, 3, 4}, {0, 3}, {0, 2, 3, 4}, {0, 1, 2, 4}},
				[]int{67, 15, 36, 39, 35},
			},
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := parseLine(test.Line)

			if !same(got, test.Want) {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}

func same(c1, c2 MachineConfig) bool {
	return slices.Equal(c1.Light, c2.Light) &&
		slices.EqualFunc(c1.Buttons, c2.Buttons, slices.Equal) &&
		slices.Equal(c1.Joltage, c2.Joltage)
}

func TestApply(t *testing.T) {
	cases := []struct {
		Initial []bool
		Button  []int
		Want    []bool
	}{
		{
			[]bool{false, false, false},
			[]int{0, 1, 2},
			[]bool{true, true, true},
		},
		{
			[]bool{true, false, true},
			[]int{0, 2},
			[]bool{false, false, false},
		},
		{
			[]bool{false, true, false},
			[]int{2, 0},
			[]bool{true, true, true},
		},
		{
			[]bool{false, true, false},
			[]int{},
			[]bool{false, true, false},
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := applyLight(test.Button, test.Initial)

			if !slices.Equal(got, test.Want) {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}

func TestMinPresses(t *testing.T) {
	cases := []struct {
		Config MachineConfig
		Want   int
	}{
		{
			testConfigs[0],
			2,
		},
		{
			testConfigs[1],
			3,
		},
		{
			testConfigs[2],
			2,
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := minPressesLight(test.Config)

			if got != test.Want {
				t.Errorf("got: %v, want: %v", got, test.Want)
			}
		})
	}
}
