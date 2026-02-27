package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/philipkrueck/advent-of-code/2015"
	"github.com/philipkrueck/advent-of-code/internal/registry"
)

func main() {
	year := flag.Int("year", 2025, "year to run (e.g. 2025)")
	day := flag.Int("day", 1, "day to run (1-25; 1-12 for 2025, 0 = all days)")
	part := flag.Int("part", 0, "part to run (1 or 2)")
	flag.Parse()

	if *day == 0 {
		runAllDays(*year)
		return
	}

	runOneDay(*year, *day, *part)
}

func runAllDays(year int) {
	for d := 1; d <= 12; d++ {
		runOneDay(year, d, 0)
	}
}

func runOneDay(year, day, part int) {
	switch part {
	case 0:
		runPart(year, day, 1)
		runPart(year, day, 2)
	case 1, 2:
		runPart(year, day, part)
	default:
		fmt.Fprintf(os.Stderr, "invalid part: %d (must be 0, 1 or 2)", part)

	}
}

func runPart(year int, day int, part int) {
	runner, input, ok := registry.Get(year, day, part)

	if !ok {
		fmt.Fprintf(os.Stderr, "no solution registered for %d day %2d part %d\n", year, day, part)
		return
	}

	fmt.Printf("%d Day %2d Part %d: ", year, day, part)

	start := time.Now()
	ans := runner(input)
	dur := time.Since(start)

	fmt.Printf("%8v (in: %s)\n", ans, dur)
}
