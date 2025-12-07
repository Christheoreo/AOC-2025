package main

import (
	"strings"
)

var reachedEnd int = 0

func SolvePartTwo(input string) int {
	// answer := 0
	lines := strings.Split(input, "\n")

	startingIndex := strings.IndexByte(lines[0], 'S')

	data := make([][]byte, 0)

	for i := 1; i < len(lines); i += 1 {
		row := make([]byte, len(lines[i]))

		for ii, char := range lines[i] {
			row[ii] = byte(char)
		}
		data = append(data, row)
	}

	partTwoMoveDown(0, startingIndex, data)

	return reachedEnd
}

func partTwoMoveDown(row, col int, lines [][]byte) {

	if row == len(lines)-1 || col < 0 || col > len(lines[0])-1 {
		// fmt.Printf("XX - %d\n", timelines)
		reachedEnd += 1
		return
	}
	char := lines[row][col]
	if char != '^' {
		partTwoMoveDown(row+1, col, lines)
		return
	}

	partTwoMoveDown(row, col-1, lines)
	partTwoMoveDown(row, col+1, lines)
}

// answer is how many ways can wr get to the bottom

// so, start from the bottom and work out - store
