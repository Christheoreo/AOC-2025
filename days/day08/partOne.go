package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type ListX [][3]int

type Distance struct {
	BoxAIndex int
	BoxBIndex int
	Distance  float64
}

type JunctionBox struct {
	X, Y, Z     int
	ID          int
	CircuitID   int
	IsConnected bool
}

func NewJunctionBox(line string, index int) JunctionBox {
	parts := strings.Split(line, ",")

	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])

	return JunctionBox{
		X:         x,
		Y:         y,
		Z:         z,
		ID:        index,
		CircuitID: index,
	}
}

func (c *JunctionBox) Ouput() string {
	return fmt.Sprintf("%d,%d,%d", c.X, c.Y, c.Z)
}

func SolvePartOne(input string, connectionsLimit int) int {
	// answer := 0
	lines := strings.Split(input, "\n")
	boxes := make([]JunctionBox, 0)
	distances := make([]Distance, 0)

	for index, line := range lines {
		box := NewJunctionBox(line, index)
		boxes = append(boxes, box)
	}

	for i := 0; i < len(boxes); i += 1 {
		for k := i + 1; k < len(boxes); k += 1 {
			a := boxes[i]
			b := boxes[k]

			if i == k {
				continue
			}

			distance := getDistance(a, b)

			distances = append(distances, Distance{
				BoxAIndex: i,
				BoxBIndex: k,
				Distance:  distance,
			})

		}
	}

	slices.SortFunc(distances, func(a Distance, b Distance) int {
		// mp(a, b) should return a negative number when a < b, a positive number when a > b and zero when a == b or a and b are incomparable in the sense of a strict weak ordering
		if a.Distance < b.Distance {
			return -1
		}
		if a.Distance > b.Distance {
			return 1
		}
		return 0
	})
	connections := 0

	for i := 0; i < len(distances); i += 1 {

		if connections == connectionsLimit {
			break
		}
		boxA := boxes[distances[i].BoxAIndex]
		boxB := boxes[distances[i].BoxBIndex]
		// They are in the same circuit, so skip

		if boxA.CircuitID == boxB.CircuitID {
			continue
		}

		if !boxA.IsConnected && !boxB.IsConnected {
			connections += 1
			boxes[distances[i].BoxAIndex].IsConnected = true
			boxes[distances[i].BoxBIndex].IsConnected = true
			boxes[distances[i].BoxBIndex].CircuitID = boxes[distances[i].BoxAIndex].CircuitID
			continue
		}

		if boxA.IsConnected && !boxB.IsConnected {
			connections += 1
			boxes[distances[i].BoxBIndex].IsConnected = true
			boxes[distances[i].BoxBIndex].CircuitID = boxes[distances[i].BoxAIndex].CircuitID
			continue
		}

		if !boxA.IsConnected && boxB.IsConnected {
			connections += 1
			boxes[distances[i].BoxAIndex].IsConnected = true
			boxes[distances[i].BoxAIndex].CircuitID = boxes[distances[i].BoxBIndex].CircuitID
			continue
		}
		// cIDToChange := boxes[distances[i].BoxBIndex].CircuitID
		// boyth connected, but to different circuits, do we move this to one circuit?
		// for k := 0; k < len(boxes); k += 1 {
		// 	if boxes[k].CircuitID == cIDToChange {
		// 		boxes[k].CircuitID = boxes[distances[i].BoxAIndex].CircuitID
		// 	}
		// }

		fmt.Println("Got here!!")

	}

	m := make(map[int]int)

	for _, v := range boxes {
		if _, ok := m[v.CircuitID]; !ok {
			m[v.CircuitID] = 0
		}
		m[v.CircuitID] += 1
	}

	// get the three largest circuits
	var a, b, c int
	for _, x := range m {

		if x > a {
			c = b
			b = a
			a = x
			continue
		}

		if x > b {
			c = b
			b = x
			continue
		}
		if x > c {
			c = x
		}
	}
	fmt.Println(connections)
	// Sort the coordinates
	return a * b * c
}

func getDistance(a, b JunctionBox) float64 {
	distanceSquared := math.Pow(float64(a.X)-float64(b.X), 2) + math.Pow(float64(a.Y)-float64(b.Y), 2) + math.Pow(float64(a.Z)-float64(b.Z), 2)

	return math.Sqrt(distanceSquared)
}
