package main

import (
	"strconv"
	"strings"
)

func SolvePartOne(input string) int {
	answer := 0
	numberSlots := make([][]int, 0)
	operationSlots := make([]string, 0)

	lines := strings.Split(input, "\n")

	for _, line := range lines {

		if strings.Contains(line, "*") {
			parts := strings.Split(line, " ")

			for _, part := range parts {
				p := strings.TrimSpace(part)
				if p != "+" && p != "*" {
					continue
				}
				operationSlots = append(operationSlots, p)
			}
			continue
		}
		parts := strings.Split(line, " ")

		nums := make([]int, 0)

		for _, part := range parts {
			if num, err := strconv.Atoi(strings.TrimSpace(part)); err == nil {
				nums = append(nums, num)
			}
		}

		numberSlots = append(numberSlots, nums)
	}

	for i := 0; i < len(operationSlots); i += 1 {
		operation := operationSlots[i]

		total := 0
		for rowIndex, row := range numberSlots {
			if rowIndex == 0 {
				total = row[i]
				continue
			}

			if operation == "+" {
				total += row[i]
			} else {
				total *= row[i]
			}
		}
		answer += total

	}

	return answer
}
