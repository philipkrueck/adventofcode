package day05

import "testing"

var in1 = `ugknbfddgicrmopn
aaa
jchzalrnumimnmhp
haegwjzuvuyypxyu
dvszwmarrgswjxmb`

func TestPart1(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{in1, "2"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part1(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}

var in2 = `qjhvhtzxzqqjkmpb
xxyxx
uurcxstgmygtbstg
ieodomkazucvgmuy`

func TestPart2(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{in2, "2"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part2(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}

func TestDuplicatePair(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"xyxy", true},
		{"aabcdefgaa", true},
		{"aaa", false},
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", true},
		{"ieodomkazucvgmuy", false},
		{"aaabbaa", true},
		{"aaaa", true},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := hasDuplicatePair(tt.in); got != tt.want {
				t.Errorf("got: %t; want: %t", got, tt.want)
			}
		})
	}
}

func TestDuplicateLetterWithExactlyOneBetween(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"xyx", true},
		{"abcdefeghi", true},
		{"aaa", true},
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", true},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := hasRepeatedWithGap(tt.in); got != tt.want {
				t.Errorf("got: %t; want: %t", got, tt.want)
			}
		})
	}
}
