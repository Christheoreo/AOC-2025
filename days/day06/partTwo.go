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
			localTotal := 0
			operation := lines[len(lines)-1][col+1]
			for i, num := range numberSlots {
				if operation == '+' {
					localTotal += num
				} else {
					if i == 0 {
						localTotal = num
					} else {
						localTotal *= num
					}
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
