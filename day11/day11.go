package day11

import (
	"strings"

	"github.com/philipkrueck/advent-of-code/lines"
)

func Part1() int {
	r := lines.NewReader("day11/input.txt")
	lines := r.Lines()
	conns := parse(lines)

	return countPaths(conns)
}

func Part2() int {
	r := lines.NewReader("day11/input.txt")
	lines := r.Lines()
	conns := parse(lines)

	return len(conns) * 0
}

func parse(lines []string) map[string][]string {
	conns := make(map[string][]string)

	for _, l := range lines {
		k, v := parseLine(l)
		conns[k] = v
	}

	return conns
}

func parseLine(line string) (string, []string) {
	keySplit := strings.Split(line, ":")
	device := keySplit[0]
	outputs := strings.Split(strings.TrimSpace(keySplit[1]), " ")

	return device, outputs
}

func countPaths(nodes map[string][]string) int {
	var count int

	queue := []string{"you"}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == "out" {
			count++
			continue
		}

		queue = append(queue, nodes[node]...)
	}

	return count
}
