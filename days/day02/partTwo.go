package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SolvePartTwo(input string) int {
	answer := 0
	ranges := strings.Split(input, ",")

	// m := make(map[int]bool)

	for _, r := range ranges {
		var start, end int
		parts := strings.Split(r, "-")
		start, _ = strconv.Atoi(parts[0])
		end, _ = strconv.Atoi(parts[1])

		for i := start; i <= end; i += 1 {
			// if val, exists := m[i]; exists && val {
			// 	answer += i
			// 	continue
			// }
			if isAPatternPartTwo(i) {
				// m[i] = true
				answer += i
				continue
			}
			// m[i] = false
		}
	}

	return answer
}

func isAPatternPartTwo(number int) bool {
	s := strconv.Itoa(number)

	// deminators := getCommonDenominators(number)

	// for _, val := range deminators {

	// }
	l := len(s)

	if l%2 == 0 {
		halfway := l / 2
		if s[0:halfway] == s[halfway:] {
			fmt.Printf("%d was a pattern match!\n", number)
			return true
		}

		for i := 1; i < l; i += 1 {
			if l%i == 0 || i == 1 {
				// fmt.Printf("here")
				// lets see if this pattern repeates
				substr := s[0:i]
				ss := s[i:]
				// fmt.Printf("S = %s, I = %d, and substr = %s\n", s, i, substr)
				for len(ss) > 0 {
					if ss[0:i] == substr {
						// fmt.Printf("SS was %s but is now %s\n", ss, ss[i:])
						ss = ss[i:]
					} else {
						break
					}
				}

				if len(ss) == 0 {
					fmt.Printf("%d was a pattern match!\n", number)
					return true
				}
			}
		}
		// Even numbers may have a pattern of 2's! ddo below for even
		return false
	}

	// if its odd

	for i := 1; i < len(s); i += 2 {
		if l%i == 0 || i == 1 {
			// fmt.Printf("here")
			// lets see if this pattern repeates
			substr := s[0:i]
			ss := s[i:]
			// fmt.Printf("S = %s, I = %d, and substr = %s\n", s, i, substr)
			for len(ss) > 0 {
				if ss[0:i] == substr {
					// fmt.Printf("SS was %s but is now %s\n", ss, ss[i:])
					ss = ss[i:]
				} else {
					break
				}
			}

			if len(ss) == 0 {
				fmt.Printf("%d was a pattern match!\n", number)
				return true
			}
		}
	}
	return false

}

// 12345657890
// 12345

// func getCommonDenominators(num int) []int {
// 	arr := make([]int, 1)
// 	m := make(map[int]any)
// 	m[1] = '0'
// 	arr[0] = 1
// 	halfway := 2
// 	if num%2 == 0 {
// 		halfway = num / 2
// 	} else {
// 		halfway = (num + 1) / 2
// 	}
// 	for i := 2; i <= halfway; i += 1 {
// 		if num%i == 0 {
// 			if _, ok := m[i]; !ok {
// 				m[i] = '0'
// 				arr = append(arr, i)
// 			}
// 		}
// 	}

//		return arr
//	}
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
