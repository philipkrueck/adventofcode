package day11

import (
	"maps"
	"slices"
	"testing"
)

var testInput []string = []string{
	"aaa: you hhh",
	"you: bbb ccc",
	"bbb: ddd eee",
	"ccc: ddd eee fff",
	"ddd: ggg",
	"eee: out",
	"fff: out",
	"ggg: out",
	"hhh: ccc fff iii",
	"iii: out",
}

var testNodes map[string][]string = map[string][]string{
	"aaa": {"you", "hhh"},
	"you": {"bbb", "ccc"},
	"bbb": {"ddd", "eee"},
	"ccc": {"ddd", "eee", "fff"},
	"ddd": {"ggg"},
	"eee": {"out"},
	"fff": {"out"},
	"ggg": {"out"},
	"hhh": {"ccc", "fff", "iii"},
	"iii": {"out"},
}

func TestLine(t *testing.T) {
	cases := []struct {
		Line    string
		Device  string
		Outputs []string
	}{
		{
			testInput[0],
			"aaa",
			testNodes["aaa"],
		},
		{
			testInput[1],
			"you",
			testNodes["you"],
		},
		{
			testInput[2],
			"bbb",
			testNodes["bbb"],
		},
		{
			testInput[3],
			"ccc",
			testNodes["ccc"],
		},

		{
			testInput[4],
			"ddd",
			testNodes["ddd"],
		},
		{
			testInput[5],
			"eee",
			testNodes["eee"],
		},
		{
			"smd: kfd vev ury nyf hbo xpq tvd quk faw jwc mnu",
			"smd",
			[]string{"kfd", "vev", "ury", "nyf", "hbo", "xpq", "tvd", "quk", "faw", "jwc", "mnu"},
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			device, outputs := parseLine(test.Line)

			if device != test.Device || !slices.Equal(outputs, test.Outputs) {
				t.Errorf("got: (%v -> %v); want: (%v -> %v)\n", device, outputs, test.Device, test.Outputs)
			}
		})
	}
}

func TestParse(t *testing.T) {
	got := parse(testInput)

	if !maps.EqualFunc(got, testNodes, slices.Equal) {
		t.Errorf("got: %v; want: %v", got, testNodes)
	}
}
