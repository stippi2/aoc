package day08

import (
	"aoc/2025/go/lib"
	"slices"
	"strconv"
	"strings"
)

type Circuit struct {
	junctionBoxes int
}

type JunctionBox struct {
	*lib.Vec3
	circuit *Circuit
}

func parseInput(input string) []*JunctionBox {
	var boxes []*JunctionBox
	for _, line := range strings.Split(input, "\n") {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		boxes = append(boxes, &JunctionBox{
			Vec3:    &lib.Vec3{X: x, Y: y, Z: z},
			circuit: &Circuit{junctionBoxes: 1},
		})
	}
	return boxes
}

type Connection struct {
	a        *JunctionBox
	b        *JunctionBox
	distance float64
}

func multiplyBiggestCircuits(input string, maxConnections int) int {
	boxes := parseInput(input)

	var connections []Connection
	for i := range len(boxes) {
		for j := i + 1; j < len(boxes); j++ {
			connections = append(connections, Connection{
				a:        boxes[i],
				b:        boxes[j],
				distance: boxes[i].Distance(*boxes[j].Vec3),
			})
		}
	}

	slices.SortFunc(connections, func(c1, c2 Connection) int {
		return int(c2.distance - c1.distance)
	})

	seenCircuits := map[*Circuit]bool{}
	var circuits []*Circuit

	for i := range maxConnections {
		circuit := connections[i].a.circuit
		circuit.junctionBoxes++
		connections[i].b.circuit = circuit
		if !seenCircuits[circuit] {
			seenCircuits[circuit] = true
			circuits = append(circuits, circuit)
		}
	}

	slices.SortFunc(circuits, func(c1, c2 *Circuit) int {
		return c2.junctionBoxes - c1.junctionBoxes
	})

	result := 1

	for i := range 3 {
		result *= circuits[i].junctionBoxes
	}

	return result
}

func Part1() any {
	input, _ := lib.ReadInput(8)
	return multiplyBiggestCircuits(input, 1000)
}

func Part2() any {
	return 0
}
