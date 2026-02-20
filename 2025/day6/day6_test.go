package day6

import (
	"slices"
	"testing"
)

var gridOne = [][][]int{
	{{1, 2, 3}, {3, 2, 8}, {0, 5, 1}, {6, 4, 0}},
	{{0, 4, 5}, {6, 4, 0}, {3, 8, 7}, {2, 3, 0}},
	{{0, 0, 6}, {9, 8, 0}, {2, 1, 5}, {3, 1, 4}},
}

var gridTwo = [][][]int{
	{{8, 4}, {7, 2, 0}, {0, 0, 0, 7}, {0, 8}, {0, 3, 6}},
	{{7, 1}, {6, 3, 0}, {8, 2, 7, 9}, {0, 6}, {0, 7, 5}},
	{{6, 4}, {5, 6, 3}, {4, 1, 6, 6}, {7, 4}, {0, 2, 1}},
	{{1, 1}, {7, 6, 3}, {2, 5, 7, 4}, {5, 4}, {7, 2, 2}},
}

func TestOperateGrid(t *testing.T) {
	cases := []struct {
		Grid [][][]int
		Ops  []Op
		Want int
	}{
		{
			gridOne,
			[]Op{Multiply, Add, Multiply, Add},
			3263827,
		},
		{
			gridTwo,
			[]Op{Multiply, Multiply, Add, Multiply, Multiply},
			804444283,
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := operateGrid(test.Grid, test.Ops)

			if got != test.Want {
				t.Errorf("got: %d; want: %d\n", got, test.Want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	cases := []struct {
		Lines    []string
		WantOps  []Op
		WantGrid [][][]int
	}{
		{
			[]string{
				"123 328  51 64 ",
				" 45 64  387 23 ",
				"  6 98  215 314",
				"*   +   *   +  ",
			},
			[]Op{Multiply, Add, Multiply, Add},
			gridOne,
		},
		{
			[]string{
				"84 72     7  8  36",
				"71 63  8279  6  75",
				"64 563 4166 74  21",
				"11 763 2574 54 722",
				"*  *   +    *  *  ",
			},
			[]Op{Multiply, Multiply, Add, Multiply, Multiply},
			gridTwo,
		},
	}

	for _, test := range cases {
		t.Run("name", func(t *testing.T) {
			gotGrid, gotOps := parse(test.Lines)

			if !slices.Equal(gotOps, test.WantOps) {
				t.Errorf("got ops: %d, want ops: %d", gotOps, test.WantOps)
			}

			if !areEqual3(gotGrid, test.WantGrid) {
				t.Errorf("got %d, want: %d", gotGrid, test.WantGrid)
			}
		})
	}
}

func TestParseLine(t *testing.T) {
	cases := []struct {
		Line         string
		DigitLengths []int
		Want         [][]int
	}{
		{
			"123 328  51 64 ",
			[]int{3, 3, 3, 3},
			[][]int{{1, 2, 3}, {3, 2, 8}, {0, 5, 1}, {6, 4, 0}},
		},
		{
			" 45 64  387 23 ",
			[]int{3, 3, 3, 3},
			[][]int{{0, 4, 5}, {6, 4, 0}, {3, 8, 7}, {2, 3, 0}},
		},
		{
			"  6 98  215 314",
			[]int{3, 3, 3, 3},
			[][]int{{0, 0, 6}, {9, 8, 0}, {2, 1, 5}, {3, 1, 4}},
		},
		{
			"84 72     7  8  36",
			[]int{2, 3, 4, 2, 3},
			[][]int{{8, 4}, {7, 2, 0}, {0, 0, 0, 7}, {0, 8}, {0, 3, 6}},
		},
		{
			"71 63  8279  6  75",
			[]int{2, 3, 4, 2, 3},
			[][]int{{7, 1}, {6, 3, 0}, {8, 2, 7, 9}, {0, 6}, {0, 7, 5}},
		},
		{
			"64 563 4166 74  21",
			[]int{2, 3, 4, 2, 3},
			[][]int{{6, 4}, {5, 6, 3}, {4, 1, 6, 6}, {7, 4}, {0, 2, 1}},
		},
		{
			"11 763 2574 54 722",
			[]int{2, 3, 4, 2, 3},
			[][]int{{1, 1}, {7, 6, 3}, {2, 5, 7, 4}, {5, 4}, {7, 2, 2}},
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := parseLine(test.Line, test.DigitLengths)

			if !areEqual2(got, test.Want) {
				t.Errorf("got %d, want: %d", got, test.Want)
			}
		})
	}
}

func TestParseDigitLengths(t *testing.T) {
	cases := []struct {
		OpsLine string
		Want    []int
	}{
		{
			"*   +   *   +  ",
			[]int{3, 3, 3, 3},
		},
		{
			"*  +   *    + ",
			[]int{2, 3, 4, 2},
		},
		{
			"*  *   +   ",
			[]int{2, 3, 4},
		},
		{
			"*   +   *  * ",
			[]int{3, 3, 2, 2},
		},
		{
			"*  *   +    *  *  ",
			[]int{2, 3, 4, 2, 3},
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := parseDigitLengths(test.OpsLine)

			if !slices.Equal(got, test.Want) {
				t.Errorf("got %d, want: %d", got, test.Want)
			}
		})
	}
}

func areEqual2(a, b [][]int) bool {
	return slices.EqualFunc(a, b, func(a1, b2 []int) bool {
		return slices.Equal(a1, b2)
	})
}

func areEqual3(a, b [][][]int) bool {
	return slices.EqualFunc(a, b, func(a1, b2 [][]int) bool {
		return slices.EqualFunc(a1, b2, slices.Equal)
	})
}

func TestConvertGridToCephs(t *testing.T) {
	cases := []struct {
		Grid  [][][]int
		Cephs [][]int
	}{
		{
			gridOne,
			[][]int{
				{1, 24, 356},
				{369, 248, 8},
				{32, 581, 175},
				{623, 431, 4},
			},
		},
		{
			gridTwo,
			[][]int{
				{8761, 4141},
				{7657, 2366, 33},
				{842, 215, 767, 7964},
				{75, 8644},
				{7, 3722, 6512},
			},
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := gridToCephs(test.Grid)

			if !areEqual2(got, test.Cephs) {
				t.Errorf("got: %v, want: %v", got, test.Cephs)
			}
		})
	}
}

func TestConvertCephs(t *testing.T) {
	cases := []struct {
		In   [][]int
		Want []int
	}{
		{
			[][]int{{1, 2, 3}, {0, 4, 5}, {0, 0, 6}},
			[]int{1, 24, 356},
		},
		{
			[][]int{{3, 2, 8}, {6, 4, 0}, {9, 8, 0}},
			[]int{369, 248, 8},
		},
		{
			[][]int{{0, 5, 1}, {3, 8, 7}, {2, 1, 5}},
			[]int{32, 581, 175},
		},
		{
			[][]int{{6, 4, 0}, {2, 3, 0}, {3, 1, 4}},
			[]int{623, 431, 4},
		},
		{
			[][]int{{1, 2, 3, 4}, {0, 5, 6, 7}},
			[]int{1, 25, 36, 47},
		},
		{
			[][]int{{8, 4}, {7, 1}, {6, 4}, {1, 1}},
			[]int{8761, 4141},
		},
		{
			[][]int{{7, 2, 0}, {6, 3, 0}, {5, 6, 3}, {7, 6, 3}},
			[]int{7657, 2366, 33},
		},
		{
			[][]int{{0, 0, 0, 7}, {8, 2, 7, 9}, {4, 1, 6, 6}, {2, 5, 7, 4}},
			[]int{842, 215, 767, 7964},
		},
		{
			[][]int{{0, 3, 6}, {0, 7, 5}, {0, 2, 1}, {7, 2, 2}},
			[]int{7, 3722, 6512},
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := convertColToCephs(test.In)

			if !slices.Equal(got, test.Want) {
				t.Errorf("got: %d; want: %d\n", got, test.Want)
			}
		})
	}
}

func TestSumNums(t *testing.T) {
	cases := []struct {
		Nums []int
		Want int
	}{
		{[]int{623, 431, 4}, 1058},
		{[]int{369, 248, 8}, 625},
		{[]int{842, 215, 767, 7964}, 9788},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := sumNums(test.Nums)

			if got != test.Want {
				t.Errorf("got: %d; want: %d\n", got, test.Want)
			}
		})
	}
}

func TestMultiplyNums(t *testing.T) {
	cases := []struct {
		Nums []int
		Want int
	}{
		{[]int{1, 24, 356}, 8544},
		{[]int{32, 581, 175}, 3253600},
		{[]int{8761, 4141}, 36279301},
		{[]int{7657, 2366, 33}, 597843246},
		{[]int{7, 3722, 6512}, 169663648},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := multiplyNums(test.Nums)

			if got != test.Want {
				t.Errorf("got: %d; want: %d\n", got, test.Want)
			}
		})
	}
}
