package day01

import "testing"

func TestPart1(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"(())", "0"},
		{"()()", "0"},
		{"(((", "3"},
		{"(()(()(", "3"},
		{"))(((((", "3"},
		{"())", "-1"},
		{"))(", "-1"},
		{")))", "-3"},
		{")())())", "-3"},
	}

	for _, tt := range cases {
		got := Part1(tt.in)
		if got != tt.want {
			t.Errorf("got: %q; want: %q", got, tt.want)
		}
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{")", "1"},
		{"()())", "5"},
	}
	for _, tt := range cases {
		got := Part2(tt.in)
		if got != tt.want {
			t.Errorf("got: %q; want: %q", got, tt.want)
		}
	}
}
