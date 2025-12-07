package main

import (
	"fmt"
	"strings"
)

var m map[string]int = make(map[string]int)

// THIS TOOK ME LIKE 3 HOURS!!!!!
func SolvePartTwo(input string) int {
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

	t := partTwoMoveDown(0, startingIndex, data)
	return t
}

func partTwoMoveDown(row, col int, lines [][]byte) int {

	if row == len(lines)-1 {
		return 1
	}

	char := lines[row][col]
	if char != '^' {
		return partTwoMoveDown(row+1, col, lines)
	}
	key := fmt.Sprintf("%d-%d", row, col)
	if val, ok := m[key]; ok {
		return val
	}

	total := partTwoMoveDown(row, col-1, lines) + partTwoMoveDown(row, col+1, lines)
	m[key] = total
	return total
}
