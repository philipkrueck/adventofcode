package lines

import (
	"bufio"
	"iter"
	"log"
	"os"
)

type LineReader struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewReader(fileName string) LineReader {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	lineReader := LineReader{file, scanner}

	return lineReader
}

func (lr *LineReader) Lines() (lines []string) {
	for line := range lr.Next() {
		lines = append(lines, line)
	}
	return lines
}

func (lr *LineReader) ByteLines() (bytes [][]byte) {
	for line := range lr.Next() {
		bytes = append(bytes, []byte(line))
	}
	return bytes
}

func (lr LineReader) Next() iter.Seq[string] {
	return func(yield func(string) bool) {
		for lr.scanner.Scan() {
			line := lr.scanner.Text()
			if !yield(line) {
				break
			}
		}
		if err := lr.scanner.Err(); err != nil {
			panic(err)
		}
		lr.file.Close()
	}
}
