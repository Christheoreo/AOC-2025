package main

import "strings"

func SolvePartOne(input string) int {
	answer := 0

	rows := make([][]byte, 0)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		row := make([]byte, 0)
		for _, b := range line {
			row = append(row, byte(b))
		}
		rows = append(rows, row)
	}

	g := NewGrid(rows)

	for row := 0; row < len(rows); row += 1 {
		for col := 0; col < len(rows[0]); col += 1 {
			if rows[row][col] != '@' {
				continue
			}
			count := g.CountRollsInAdjacentPositions(row, col)
			if count < 4 {
				answer += 1
			}
		}
	}

	return answer
}
