package main

import (
	"strconv"
	"strings"
)

func SolvePartOne(input string) int {
	answer := 0
	ranges := make([][2]int, 0)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")

			lower, _ := strconv.Atoi(parts[0])
			higher, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, [2]int{lower, higher})
			continue
		}
		num, _ := strconv.Atoi(line)

		for _, r := range ranges {
			if num >= r[0] && num <= r[1] {
				answer += 1
				break
			}
		}
	}

	return answer
}
