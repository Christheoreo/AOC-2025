package main

import (
	"strconv"
	"strings"
)

func SolvePartTwo(input string) int {
	lines := strings.Split(input, "\n")

	answer := 0
	pointer := 50
	for _, line := range lines {
		direction := line[0]
		numberS := line[1:]
		num, _ := strconv.Atoi(numberS)
		startedAtZero := pointer == 0
		if direction == 'L' {
			pointer -= num
			rotated := false
			for pointer < 0 {
				rotated = true
				pointer += 100
				// only add it we went past 0
				if !startedAtZero || (startedAtZero && num < -99) {
					answer += 1
				}
			}

			if !rotated && pointer == 0 {
				answer += 1
			}

		} else {
			// R
			pointer += num
			for pointer > 99 {
				pointer -= 100
				answer += 1
			}
		}

	}

	return answer
}
