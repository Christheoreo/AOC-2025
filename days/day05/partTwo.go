package main

import (
	"strconv"
	"strings"
)

func SolvePartTwo(input string) int {
	answer := 0
	ranges := make([][2]int, 0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")

		lower, _ := strconv.Atoi(parts[0])
		higher, _ := strconv.Atoi(parts[1])

		ranges = append(ranges, [2]int{lower, higher})
	}

	for {
		madeAChange := false
		for i := 0; i < len(ranges); i += 1 {
			r := ranges[i]
			if r[0] == 0 && r[1] == 0 {
				continue
			}
			for k := 0; k < len(ranges); k += 1 {
				if i == k {
					continue
				}
				rr := ranges[k]
				if rr[0] == 0 && rr[1] == 0 {
					continue
				}
				localChange := false

				// check if rr can go inside r
				if rr[0] >= r[0] && rr[1] <= r[1] {
					ranges[k] = [2]int{0, 0}
					continue
				}

				// does the lower boundary need changing?
				if rr[0] < r[0] && rr[1] >= r[0] {
					ranges[i][0] = ranges[k][0]
					// ranges[k] = [2]int{0, 0}
					madeAChange = true
					localChange = true
				}

				// does the upper boundary need changing?
				if rr[1] > r[1] && rr[0] <= r[1] {
					ranges[i][1] = ranges[k][1]
					madeAChange = true
					localChange = true
				}

				if localChange {
					ranges[k] = [2]int{0, 0}
				}
			}
		}

		if !madeAChange {
			break
		}

	}

	for _, r := range ranges {
		if r[0] == 0 && r[1] == 0 {
			continue
		}

		diff := (r[1] - r[0]) + 1

		answer += diff
	}

	return answer
}
