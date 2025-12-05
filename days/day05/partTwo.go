package main

import (
	"strconv"
	"strings"
)

func SolvePartTwo(input string) int {
	answer := 0
	ranges := make([][2]int, 0)
	Overlappingranges := make([][2]int, 0)

	lines := strings.Split(input, "\n")
	// key is the min, value is the max
	// m := make(map[int]int)
	for _, line := range lines {
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")

		lower, _ := strconv.Atoi(parts[0])
		higher, _ := strconv.Atoi(parts[1])

		ranges = append(ranges, [2]int{lower, higher})
	}

	// TODO-
	// dont loop through i, just frorever look at first rnage, deal with it and move it
	for len(ranges) > 0 {

		for i := 0; i < len(ranges); i += 1 {
			r := ranges[i]

			for ii := 0; ii < len(Overlappingranges); ii += 1 {
				rr := Overlappingranges[ii]
				adjusted := false

				// can we just ignore it?
				if r[0] >= rr[0] && r[1] <= rr[1] {
					// fitd within it, so lets ignore it
					break
				}

				if r[0] < rr[0] {
					// weve extended the bottom range
					if r[1] < rr[1] && r[1] > rr[0] {
						Overlappingranges[ii][0] = r[0]
						adjusted = true
					}

					if r[1] > rr[1] {
						Overlappingranges[ii][1] = r[1]
						adjusted = true
					}
				}

				if r[0] > rr[0] && r[0] < rr[1] {

					// lets see if we can extend range
					if r[1] > rr[1] {
						Overlappingranges[ii][1] = r[1]
						adjusted = true
					}
				}

				if adjusted {
					// delete from ranges?
				}
			}

			Overlappingranges = append(Overlappingranges, r)
			// remove r from ranges now? and for any time we adjust it above?

		}

	}

	return answer
}
