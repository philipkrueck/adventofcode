package main

import (
	"fmt"

	"github.com/philipkrueck/advent-of-code/day2"
)

func main() {
	fmt.Println("Solution to day two part 2:", day2.Part2())

	x := 54321
	digits := x / 10
	fmt.Printf("digits of %d is %d\n", x, digits)
}
