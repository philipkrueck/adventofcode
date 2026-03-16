package day02

import "testing"

func TestRun(t *testing.T) {
	tests := []struct {
		intcode []int
		result  int
	}{
		{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, 3500},
		{[]int{1, 0, 0, 0, 99}, 2},
		{[]int{2, 3, 0, 3, 99}, 2},
		{[]int{2, 4, 4, 5, 99, 0}, 2},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 30},
	}

	for _, tt := range tests {
		if got := run(tt.intcode); got != tt.result {
			t.Errorf("got: %v; want: %v", got, tt.result)
		}
	}
}

func TestPart2(t *testing.T) {
	t.Skip("skipping...")
	tests := []struct {
		in, want string
	}{
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part2(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}
