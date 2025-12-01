package main

import (
	"strconv"
	"strings"
)

type Lock struct {
	pointer              int
	AmountOfLandedOnZero int
}

func NewLock() *Lock {
	return &Lock{
		pointer:              50,
		AmountOfLandedOnZero: 0,
	}
}

func (l *Lock) turnLeft() {
	l.pointer -= 1
	if l.pointer == 0 {
		l.AmountOfLandedOnZero += 1
	}
	if l.pointer == -1 {
		l.pointer = 99
	}
}

func (l *Lock) turnRight() {
	l.pointer += 1
	if l.pointer == 100 {
		l.AmountOfLandedOnZero += 1
		l.pointer = 0
	}
}

func (l *Lock) TurnLeftX(x int) {
	for i := 0; i < x; i += 1 {
		l.turnLeft()
	}
}

func (l *Lock) TurnRightX(x int) {
	for i := 0; i < x; i += 1 {
		l.turnRight()
	}
}

// This is a BAD solution, but a working one regardless!
func SolvePartTwo(input string) int {
	lines := strings.Split(input, "\n")
	l := NewLock()
	for _, line := range lines {
		direction := line[0]
		numberS := line[1:]
		num, _ := strconv.Atoi(numberS)
		if direction == 'L' {
			l.TurnLeftX(num)
		} else {
			// R
			l.TurnRightX(num)
		}
	}

	return l.AmountOfLandedOnZero
}

// Something was off here, will try this again later!
// func SolvePartTwo(input string) int {
// 	lines := strings.Split(input, "\n")

// 	answer := 0
// 	pointer := 50
// 	for _, line := range lines {
// 		direction := line[0]
// 		numberS := line[1:]
// 		num, _ := strconv.Atoi(numberS)
// 		// startedAtZero := pointer == 0
// 		if direction == 'L' {
// 			pointer -= num
// 			if pointer == 0 {
// 				answer += 1
// 				continue
// 			}

// 			if pointer < -99 {
// 				rotation := pointer % -100
// 				answer += rotation
// 				pointer += (rotation * 100)
// 			}

// 		} else {
// 			// R
// 			pointer += num
// 			for pointer > 99 {
// 				pointer -= 100
// 				answer += 1
// 			}
// 		}

// 	}

// 	return answer
// }
