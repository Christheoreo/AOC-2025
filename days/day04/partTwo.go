package main

import "strings"

// type Grid2 struct {
// 	rows [][]byte
// }

// func NewGrid2(rows [][]byte) *Grid2 {
// 	return &Grid2{
// 		rows: rows,
// 	}
// }

// func (g *Grid2) CountRollsInAdjacentPositions(row, col int) int {
// 	// 8 Adject posisitions
// 	count := 0
// 	// can we go up?
// 	canLookLeft := col > 0
// 	canLookRight := col < len(g.rows[0])-1
// 	canLookUp := row > 0
// 	canLookDown := row < len(g.rows)-1
// 	if canLookUp {
// 		if canLookLeft {
// 			if g.rows[row-1][col-1] == '@' {
// 				count += 1
// 			}
// 		}
// 		// top middle
// 		if g.rows[row-1][col] == '@' {
// 			count += 1
// 		}
// 		// top right
// 		if canLookRight {
// 			if g.rows[row-1][col+1] == '@' {
// 				count += 1
// 			}
// 		}
// 	}

// 	if canLookLeft {
// 		if g.rows[row][col-1] == '@' {
// 			count += 1
// 		}
// 	}

// 	if canLookRight {
// 		if g.rows[row][col+1] == '@' {
// 			count += 1
// 		}
// 	}

// 	if canLookDown {
// 		if canLookLeft {
// 			if g.rows[row+1][col-1] == '@' {
// 				count += 1
// 			}
// 		}
// 		// bottom middle
// 		if g.rows[row+1][col] == '@' {
// 			count += 1
// 		}
// 		// bottom right
// 		if canLookRight {
// 			if g.rows[row+1][col+1] == '@' {
// 				count += 1
// 			}
// 		}
// 	}
// 	return count
// }

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

		if len(positionsToRemove) == 0 {
			break
		}

		for _, x := range positionsToRemove {
			rows[x[0]][x[1]] = '.'
		}

		g = NewGrid(rows)

	}

	return answer
}
