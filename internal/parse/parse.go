// Package parse is used for common input processing
package parse

import (
	"iter"
	"slices"
	"strings"
)

// Lines returns non-empty lines of an input string.
func Lines(input string) []string {
	return slices.Collect(LinesSeq(input))
}

// LinesSeq returns non-empty lines of an input string as an iterator.
func LinesSeq(input string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for line := range strings.SplitSeq(input, "\n") {
			if line != "" {
				if !yield(line) {
					return
				}
			}
		}
	}
}
