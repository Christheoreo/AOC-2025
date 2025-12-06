package main

import (
	"strconv"
	"strings"
)

func SolvePartTwo(input string) int {
	answer := 0
	numberSlots := make([]int, 0)

	lines := strings.Split(input, "\n")
	lineLength := len(lines[0])

	// go to -1 as pos 0 will always be an operation
	for col := lineLength - 1; col >= -1; col -= 1 {
		total := 0
		if col >= 0 {
			for i := 0; i < len(lines)-1; i += 1 {
				char := lines[i][col]
				if char == ' ' {
					continue
				}
				num, _ := strconv.Atoi(string(char))
				if total == 0 {
					total = num
				} else {
					total *= 10
					total += num
				}
			}
		}
		if total == 0 {
			// we found  nothing!
			localTotal := numberSlots[0]

			// doing this instead of the loop then a operation check so that we just do the check once
			if lines[len(lines)-1][col+1] != '+' {
				for i := 1; i < len(numberSlots); i += 1 {
					localTotal *= numberSlots[i]
				}
			} else {
				for i := 1; i < len(numberSlots); i += 1 {
					localTotal += numberSlots[i]
				}
			}
			answer += localTotal
			numberSlots = make([]int, 0)
			continue
		}
		numberSlots = append(numberSlots, total)
	}

	return answer
}
