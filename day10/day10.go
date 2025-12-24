package day10

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/philipkrueck/advent-of-code/lines"
)

type MachineConfig struct {
	Light   []bool
	Buttons [][]int
	Joltage []int
}

// TODO: could generalize these
type LightNode struct {
	State []bool
	depth int
}
type JoltageNode struct {
	State    []int
	depth    int
	distance int
}

func Part1() int {
	r := lines.NewReader("day10/test-input.txt")
	lines := r.Lines()
	configs := parse(lines)

	return minPressesLights(configs)
}

func Part2() int {
	r := lines.NewReader("day10/input.txt")
	lines := r.Lines()
	configs := parse(lines)

	return minPressesJoltages(configs)
}

func parse(lines []string) []MachineConfig {
	configs := []MachineConfig{}
	for _, line := range lines {
		configs = append(configs, parseLine(line))
	}
	return configs
}

func parseLine(line string) MachineConfig {
	line = line[1:]
	parts := strings.Split(line, "]")
	goalStr := parts[0]
	remainingParts := strings.Split(parts[1], "{")

	buttonStr := strings.TrimSpace(remainingParts[0])
	buttonStr = strings.ReplaceAll(buttonStr, "(", "")
	buttonStr = strings.ReplaceAll(buttonStr, ")", "")
	buttonStrs := strings.Split(buttonStr, " ")

	joltageStr := remainingParts[1]
	joltageStr = strings.ReplaceAll(joltageStr, "}", "")
	joltageStrs := strings.Split(joltageStr, ",")

	goal := []bool{}
	buttons := [][]int{}
	joltage := []int{}

	for _, ch := range goalStr {
		switch ch {
		case '#':
			goal = append(goal, true)
		case '.':
			goal = append(goal, false)
		}
	}

	for _, btnStr := range buttonStrs {
		nums := strings.Split(btnStr, ",")
		button := []int{}
		for _, nStr := range nums {
			n, err := strconv.Atoi(nStr)
			if err != nil {
				panic("Bad input")
			}
			button = append(button, n)
		}
		buttons = append(buttons, button)
	}

	for _, joltageStr := range joltageStrs {
		n, err := strconv.Atoi(joltageStr)
		if err != nil {
			panic("Bad input")
		}
		joltage = append(joltage, n)
	}

	return MachineConfig{goal, buttons, joltage}
}

func minPressesLights(configs []MachineConfig) int {
	var presses int
	for _, c := range configs {
		presses += minPressesLight(c)
	}
	return presses
}

func minPressesLight(config MachineConfig) int {
	initial := make([]bool, len(config.Light))
	queue := []LightNode{{initial, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if slices.Equal(curr.State, config.Light) {
			return curr.depth
		}
		for _, b := range config.Buttons {
			nextState := applyLight(b, curr.State)
			queue = append(queue, LightNode{nextState, curr.depth + 1})
		}
	}

	return 0
}

func minPressesJoltages(configs []MachineConfig) int {
	var presses int
	for _, c := range configs {
		fmt.Printf("finding min joltage for config: %v\n\n", c)
		presses += minPressesJoltage(c)
	}
	return presses
}

type VisitedKey [16]int

func minPressesJoltage(config MachineConfig) int {
	initial := make([]int, len(config.Light))
	root := JoltageNode{initial, 0, distanceToGoal(initial, config.Joltage)}
	stack := []JoltageNode{root}
	visited := make(map[VisitedKey]bool)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		// check visited
		key := VisitedKey{}
		copy(key[:], curr.State)
		if visited[key] {
			continue
		}
		visited[key] = true

		if slices.Equal(curr.State, config.Joltage) {
			fmt.Println("solved with presses:", curr.depth)
			return curr.depth
		}

		nextJoltages := []JoltageNode{}

		for _, b := range config.Buttons {
			nextState := applyJoltage(b, curr.State)

			if !validState(nextState, config.Joltage) {
				continue
			}
			node := JoltageNode{
				nextState,
				curr.depth + 1,
				distanceToGoal(nextState, config.Joltage),
			}
			nextJoltages = append(nextJoltages, node)
		}
		slices.SortFunc(nextJoltages, func(a, b JoltageNode) int {
			return b.distance - a.distance
		})
		stack = append(stack, nextJoltages...)
	}

	return 0
}

func applyLight(button []int, state []bool) []bool {
	nextState := slices.Clone(state)
	for _, i := range button {
		nextState[i] = !nextState[i]
	}
	return nextState
}

func applyJoltage(button []int, state []int) []int {
	nextState := slices.Clone(state)
	for _, i := range button {
		nextState[i] = nextState[i] + 1
	}
	return nextState
}

func validState(state []int, goal []int) bool {
	for i := range state {
		if state[i] > goal[i] {
			return false
		}
	}
	return true
}

func distanceToGoal(state []int, goal []int) int {
	var distance int
	for i := range state {
		distance += goal[i] - state[i]
	}
	return distance
}
