package main

type Grid struct {
	rows [][]byte
}

func NewGrid(rows [][]byte) *Grid {
	return &Grid{
		rows: rows,
	}
}

func (g *Grid) CountRollsInAdjacentPositions(row, col int) int {
	count := 0

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

		if canLookRight {
			if g.rows[row+1][col+1] == '@' {
				count += 1
			}
		}
	}
	return count
}
