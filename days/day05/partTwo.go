package main

import (
	"math"
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
		foundACrossOver := false
		for i := 0; i < len(ranges); i += 1 {

			if ranges[i][0] == 0 && ranges[i][1] == 0 {
				continue
			}

			for k := 0; k < len(ranges); k += 1 {
				if i == k || (ranges[k][0] == 0 && ranges[k][1] == 0) {
					continue
				}

				r := ranges[i]
				rr := ranges[k]

				if hasAnyCrossover(r, rr) {
					foundACrossOver = true
					ranges[i] = [2]int{
						int(
							math.Min(
								float64(r[0]),
								float64(rr[0]),
							),
						),
						int(
							math.Max(
								float64(r[1]),
								float64(rr[1]),
							),
						),
					}
					ranges[k] = [2]int{0, 0}
				}
			}
		}

		// if we loop through and didnt find a single cross over, exit!
		if !foundACrossOver {
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

func hasAnyCrossover(r1, r2 [2]int) bool {
	if (r1[0] >= r2[0] && r1[0] <= r2[1]) ||
		(r1[1] >= r2[0] && r1[1] <= r2[1]) ||
		(r2[0] >= r1[0] && r2[0] <= r1[1]) ||
		(r2[1] >= r1[0] && r2[1] <= r1[1]) {
		return true
	}
	return false
}
