package main

import "strings"

type Grid struct {
	rows [][]byte
}

func NewGrid(rows [][]byte) *Grid {
	return &Grid{
		rows: rows,
	}
}

func (g *Grid) CountRollsInAdjacentPositions(row, col int) int {
	// 8 Adject posisitions
	count := 0
	// can we go up?
	canLookLeft := col > 0
	canLookRight := col < len(g.rows[0])-1
	canLookUp := row > 0
	canLookDown := row < len(g.rows)-1
	if canLookUp {
		if canLookLeft {
			if g.rows[row-1][col-1] == '@' {
				count += 1
			}
		}
		// top middle
		if g.rows[row-1][col] == '@' {
			count += 1
		}
		// top right
		if canLookRight {
			if g.rows[row-1][col+1] == '@' {
				count += 1
			}
		}
	}

	if canLookLeft {
		if g.rows[row][col-1] == '@' {
			count += 1
		}
	}

	if canLookRight {
		if g.rows[row][col+1] == '@' {
			count += 1
		}
	}

	if canLookDown {
		if canLookLeft {
			if g.rows[row+1][col-1] == '@' {
				count += 1
			}
		}
		// bottom middle
		if g.rows[row+1][col] == '@' {
			count += 1
		}
		// bottom right
		if canLookRight {
			if g.rows[row+1][col+1] == '@' {
				count += 1
			}
		}
	}
	return count
}

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
