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
		// var aa, bb rune
		for i, c := range line {
			num, _ := strconv.Atoi(string(c))
			if i == 0 {
				a = num
				// aa = c
				continue
			}
			if num > a && i != len(line)-1 {
				a = num
				// aa = c
				b = 0
				// bb = '0'
				continue
			}
			if num > b {
				b = num
				// bb = c
			}
		}
		num, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
		fmt.Printf("Adding %d to %d\n", num, answer)
		answer += num
	}
	return answer
}
