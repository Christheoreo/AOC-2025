package main

import "strings"

func SolvePartTwo(input string) int {
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

	for {
		positionsToRemove := make([][2]int, 0)
		for row := 0; row < len(rows); row += 1 {
			for col := 0; col < len(rows[0]); col += 1 {
				if rows[row][col] != '@' {
					continue
				}
				count := g.CountRollsInAdjacentPositions(row, col)
				if count < 4 {
					positionsToRemove = append(positionsToRemove, [2]int{row, col})
					answer += 1
				}
			}
		}

		// We can't remove anymore!
		if len(positionsToRemove) == 0 {
			break
		}

		for _, posToRemove := range positionsToRemove {
			rows[posToRemove[0]][posToRemove[1]] = '.'
		}

		g = NewGrid(rows)

	}

	return answer
}
