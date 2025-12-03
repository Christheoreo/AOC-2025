package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SolvePartTwo(input string) int {
	answer := 0
	lines := strings.Split(input, "\n")
	lineLength := len(lines[0])
	m := make(map[int]int)
	x := 11
	for i := lineLength - 1; i >= lineLength-12; i -= 1 {
		m[i] = x
		x -= 1
	}
	for _, line := range lines {
		var selectedBatteries [12]int
		for i, c := range line {

			num, _ := strconv.Atoi(string(c))
			earliestIndex, ok := m[i]

			for ii := 0; ii < len(selectedBatteries); ii += 1 {
				if ok && ii < earliestIndex {
					continue
				}
				if num > selectedBatteries[ii] {
					selectedBatteries[ii] = num
					for iii := ii + 1; iii < len(selectedBatteries); iii += 1 {
						selectedBatteries[iii] = 0
					}
					break
				}
			}
		}
		num, _ := strconv.Atoi(fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d", selectedBatteries[0], selectedBatteries[1], selectedBatteries[2], selectedBatteries[3], selectedBatteries[4], selectedBatteries[5], selectedBatteries[6], selectedBatteries[7], selectedBatteries[8], selectedBatteries[9], selectedBatteries[10], selectedBatteries[11]))
		fmt.Printf("Adding %d to total %d\n", num, answer)
		answer += num
	}
	return answer
}
