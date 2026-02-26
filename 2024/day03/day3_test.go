package day03

import "testing"

func TestParseLine(t *testing.T) {
	in := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	want := 2*4 + 5*5 + 11*8 + 8*5

	got := parseLine(in)
	if got != want {
		t.Errorf("got: %d; want: %d", got, want)
	}
}

func TestParseLineSwitch(t *testing.T) {
	cases := []struct {
		in      string
		wantSum int
		wantDo  bool
	}{
		{
			"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			2*4 + 8*5,
			true,
		},
		{
			"don't()xmul(2,4)&mul[3,7]!^_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			8 * 5,
			true,
		},
	}

	for _, tt := range cases {
		gotSum, gotDo := parseLineSwitch(tt.in, true)
		if gotSum != tt.wantSum || gotDo != tt.wantDo {
			t.Errorf("gotSum: %d; wantSum: %d; gotDo: %v; wantDo: %v", gotSum, tt.wantSum, gotDo, tt.wantDo)
		}

	}
}
