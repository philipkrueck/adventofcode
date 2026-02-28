package day11

import (
	"strings"

	"github.com/philipkrueck/adventofcode/lines"
)

func Part1() int {
	r := lines.NewReader("day11/input.txt")
	lines := r.Lines()
	conns := parse(lines)
	cache := make(map[string]int)

	return countPaths(conns, "you", cache)
}

func Part2() int {
	r := lines.NewReader("day11/input.txt")
	lines := r.Lines()
	conns := parse(lines)
	start := Node{"svr", false, false}
	cache := make(map[Node]int)

	return countPathsPart2(conns, start, cache)
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

type Node struct {
	node     string
	fft, dac bool
}

func countPaths(nodes map[string][]string, node string, cache map[string]int) int {
	if node == "out" {
		return 1
	}

	if count, ok := cache[node]; ok {
		return count
	}

	var count int
	for _, neighbor := range nodes[node] {
		count += countPaths(nodes, neighbor, cache)
	}

	cache[node] = count
	return count
}

func countPathsPart2(nodes map[string][]string, node Node, cache map[Node]int) int {
	switch node.node {
	case "out":
		if node.dac && node.fft {
			return 1
		}
	case "dac":
		node.dac = true
	case "fft":
		node.fft = true
	}

	if count, ok := cache[node]; ok {
		return count
	}

	var count int
	for _, neighbor := range nodes[node.node] {
		neighborNode := Node{neighbor, node.fft, node.dac}
		count += countPathsPart2(nodes, neighborNode, cache)
	}

	cache[node] = count
	return count
}
