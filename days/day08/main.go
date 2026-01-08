package main

import (
	_ "embed"
	"fmt"
)

//go:embed inputs/example.txt
var exampleInput string

//go:embed inputs/input.txt
var input string

func main() {
	// partOneAnswer := SolvePartOne(input,1000)
	partOneAnswer := SolvePartOne(exampleInput, 10) // run in debugger to see what the value of distance is
	// partTwoAnswer := SolvePartTwo(input)
	fmt.Printf("Day 08 part 1 = %v\n", partOneAnswer)
	// fmt.Printf("Day 08 part 2 = %v\n", partTwoAnswer)
}
