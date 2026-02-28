// Package day10 implements day10 of 2025 aoc
package day10

import (
	"slices"
	"strconv"
	"strings"

	"github.com/philipkrueck/adventofcode/lines"
)

type MachineConfig struct {
	Light   []bool
	Buttons [][]int
	Joltage []int
}

type LightNode struct {
	State []bool
	depth int
}
type JoltageNode struct {
	State []int
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

	cache := make(map[Joltage]int)

	for _, c := range configs {
		combos := btnCombos(c.Buttons)
		minPresses := minPressesJoltagesRec(toJoltage(c.Joltage), cache, combos)
		presses += minPresses
	}
	return presses
}

func toJoltage(old []int) (j Joltage) {
	copy(j[:], old)
	return
}

// This generates the power set
// TODO: write this as a generic power set generation algorithm
func btnCombos(buttons [][]int) [][][]int {
	numButtons := len(buttons)
	numCombos := 1 << numButtons // 2^numButtons
	combos := make([][][]int, numCombos)

	for mask := range numCombos {
		pressed := [][]int{}
		for btn := range numButtons {
			if (mask & (1 << btn)) != 0 {
				pressed = append(pressed, buttons[btn])
			}
		}
		combos[mask] = pressed
	}
	return combos
}

const MaxJoltages = 10

type Joltage [MaxJoltages]int

func minPressesJoltagesRec(j Joltage, cache map[Joltage]int, btnCombos [][][]int) int {
	if allZero(j) {
		return 0
	}

	if joltage, ok := cache[j]; ok {
		return joltage
	}

	minPresses := 10000000

	for _, combo := range btnCombos {
		newJ := applyBtnCombo(j, combo)
		if !joltageIsEven(newJ) {
			continue
		}

		count := 2*minPressesJoltagesRec(joltageHalved(newJ), cache, btnCombos) + len(combo)
		if count < minPresses {
			minPresses = count
		}
	}

	// TODO: add caching for optimization
	// cache[j] = minPresses

	return minPresses
}

func applyBtnCombo(joltage Joltage, combo [][]int) Joltage {
	for _, btn := range combo {
		for _, joltIdx := range btn {
			joltage[joltIdx]--
		}
	}
	return joltage
}

func joltageHalved(joltage Joltage) Joltage {
	for i := range joltage {
		joltage[i] /= 2
	}
	return joltage
}

func joltageIsEven(joltage Joltage) bool {
	for _, jolt := range joltage {
		if jolt%2 != 0 || jolt < 0 {
			return false
		}
	}
	return true
}

func allZero(joltage Joltage) bool {
	for _, jolt := range joltage {
		if jolt != 0 {
			return false
		}
	}
	return true
}

func applyLight(button []int, state []bool) []bool {
	nextState := slices.Clone(state)
	for _, i := range button {
		nextState[i] = !nextState[i]
	}
	return nextState
}
