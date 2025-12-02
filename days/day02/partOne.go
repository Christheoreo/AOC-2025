package main

import (
	"strconv"
	"strings"
)

func SolvePartOne(input string) int {
	answer := 0
	ranges := strings.Split(input, ",")

	m := make(map[int]bool)

	for _, r := range ranges {
		var start, end int
		parts := strings.Split(r, "-")
		start, _ = strconv.Atoi(parts[0])
		end, _ = strconv.Atoi(parts[1])

		for i := start; i <= end; i += 1 {
			if val, exists := m[i]; exists && val {
				answer += i
				continue
			}
			if isAPatternPartOne(i) {
				m[i] = true
				answer += i
				continue
			}
			m[i] = false
		}
	}

	return answer
}

func isAPatternPartOne(number int) bool {
	s := strconv.Itoa(number)

	if len(s)%2 != 0 {
		return false
	}

	halfway := len(s) / 2
	return s[0:halfway] == s[halfway:]
}
