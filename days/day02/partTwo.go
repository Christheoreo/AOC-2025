package main

import (
	"strconv"
	"strings"
)

func SolvePartTwo(input string) int {
	answer := 0
	ranges := strings.Split(input, ",")

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i += 1 {
			if isAPatternPartTwo(i) {
				answer += i
				continue
			}
		}
	}

	return answer
}

func isAPatternPartTwo(number int) bool {
	s := strconv.Itoa(number)
	l := len(s)

	// will be 2 for odd numbers, 1 for even
	incrAmount := 2
	if l%2 == 0 {
		incrAmount = 1
	}

	for i := 1; i < len(s); i += incrAmount {
		if l%i == 0 {
			substr := s[0:i]
			ss := s[i:]
			for len(ss) > 0 {
				if ss[0:i] == substr {
					// remove the first part and lets see if the pattern exists again
					ss = ss[i:]
				} else {
					break
				}
			}
			// if the whole string is empty, then a pattern existed!
			if len(ss) == 0 {
				return true
			}
		}
	}
	return false
}
