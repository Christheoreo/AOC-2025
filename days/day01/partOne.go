package main

import (
	"strconv"
	"strings"
)

func SolvePartOne(input string) int {
	lines := strings.Split(input, "\n")

	answer := 0
	pointer := 50

	// sort the data

	for _, line := range lines {
		direction := line[0]
		numberS := line[1:]
		num, _ := strconv.Atoi(numberS)
		if direction == 'L' {
			pointer -= num
			for pointer < 0 {
				pointer += 100
			}

		} else {
			// R
			pointer += num
			for pointer > 99 {
				pointer -= 100
			}
		}

		if pointer == 0 {
			answer += 1
		}
	}

	return answer
}
