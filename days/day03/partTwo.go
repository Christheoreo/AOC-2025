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

	// Make a map of where each index maps to the earliest index of the selectedBatteries
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
				// if int isnt in map, we know we can start from 0, else we want to skip to the earliest index
				if ok && ii < earliestIndex {
					continue
				}
				if num > selectedBatteries[ii] {
					// update the number and then reset everything after it
					selectedBatteries[ii] = num
					for iii := ii + 1; iii < len(selectedBatteries); iii += 1 {
						selectedBatteries[iii] = 0
					}
					break
				}
			}
		}
		// this feels horrinle!
		num, _ := strconv.Atoi(fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d", selectedBatteries[0], selectedBatteries[1], selectedBatteries[2], selectedBatteries[3], selectedBatteries[4], selectedBatteries[5], selectedBatteries[6], selectedBatteries[7], selectedBatteries[8], selectedBatteries[9], selectedBatteries[10], selectedBatteries[11]))
		answer += num
	}
	return answer
}
