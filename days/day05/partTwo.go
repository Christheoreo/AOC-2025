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

		r := ranges[0]

		adjusted := false
		for ii := 0; ii < len(Overlappingranges); ii += 1 {
			rr := Overlappingranges[ii]
			adjusted = false

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
				break
			}
		}

		ranges = ranges[1:]

		if !adjusted {
			Overlappingranges = append(Overlappingranges, r)
		}

		// remove r from ranges now? and for any time we adjust it above?

	}

	for {
		madeAChange := false
		for i := 0; i < len(Overlappingranges); i += 1 {
			r := Overlappingranges[i]
			if r[0] == 0 && r[1] == 0 {
				continue
			}
			for k := 0; k < len(Overlappingranges); k += 1 {
				if i == k {
					continue
				}
				rr := Overlappingranges[k]
				if rr[0] == 0 && rr[1] == 0 {
					continue
				}
				localChange := false

				// check if rr can go inside r
				if rr[0] >= r[0] && rr[1] <= r[1] {
					Overlappingranges[k] = [2]int{0, 0}
					continue
				}

				// does the lower boundary need changing?
				if rr[0] < r[0] && rr[1] >= r[0] {
					Overlappingranges[i][0] = Overlappingranges[k][0]
					// Overlappingranges[k] = [2]int{0, 0}
					madeAChange = true
					localChange = true
				}

				// does the upper boundary need changing?
				if rr[1] > r[1] && rr[0] <= r[1] {
					Overlappingranges[i][1] = Overlappingranges[k][1]
					madeAChange = true
					localChange = true
				}

				if localChange {
					Overlappingranges[k] = [2]int{0, 0}
				}
			}
		}

		if !madeAChange {
			break
		}

	}

	for _, r := range Overlappingranges {
		if r[0] == 0 && r[1] == 0 {
			continue
		}

		answer += (r[1] - r[0]) + 1
	}

	return answer
}
