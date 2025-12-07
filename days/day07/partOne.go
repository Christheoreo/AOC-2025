package main

import (
	"fmt"
	"strings"
)

func SolvePartOne(input string) int {
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

	m := partOneMoveDown(0, startingIndex, data, make(map[string]interface{}))

	return len(m)
}

func partOneMoveDown(row, col int, lines [][]byte, m map[string]interface{}) map[string]interface{} {

	if row == len(lines)-1 {
		return m
	}
	char := lines[row][col]
	if char != '^' {
		return partOneMoveDown(row+1, col, lines, m)
	}
	key := fmt.Sprintf("%d-%d", row, col)
	_, ok := m[key]

	if ok {
		// we've already been here before, we can just return
		return m
	}
	m[key] = '^'

	mm := partOneMoveDown(row, col-1, lines, m)
	return partOneMoveDown(row, col+1, lines, mm)
}
