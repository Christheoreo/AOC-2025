package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SolvePartOne(input string) int {
	answer := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		var a, b int
		for i, c := range line {
			num, _ := strconv.Atoi(string(c))
			if i == 0 {
				a = num
				continue
			}
			if num > a && i != len(line)-1 {
				a = num
				b = 0
				continue
			}
			if num > b {
				b = num
			}
		}
		num, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
		answer += num
	}
	return answer
}
