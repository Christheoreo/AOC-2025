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
	partOneAnswer := SolvePartOne(input)
	partTwoAnswer := SolvePartTwo(input)
	fmt.Printf("Day 01 part 1 = %v\n", partOneAnswer)
	fmt.Printf("Day 01 part 2 = %v\n", partTwoAnswer)
}
