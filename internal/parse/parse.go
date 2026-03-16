// Package parse is used for common input processing
package parse

import (
	"iter"
	"strings"
)

func Lines(input string) []string {
	lines := strings.Split(input, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		return lines[:len(lines)-1]
	}
	return lines
}

func LinesSeq(input string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for line := range strings.SplitSeq(input, "\n") {
			if line != "" {
				yield(line)
			}
		}
	}
}
