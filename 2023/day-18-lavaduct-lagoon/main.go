package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Instruction struct {
	direction byte
	cubeCount int
	edgeColor string
}

type Point struct {
	x, y float64
}

type LavaLagoon struct {
	points    []Point
	edgeCubes int64
}

func (l *LavaLagoon) buildEdgePartOne(instructions []Instruction) {
	cursor := Point{0, 0}
	l.points = []Point{cursor}
	l.edgeCubes = 1
	for _, instruction := range instructions {
		switch instruction.direction {
		case 'U':
			cursor.y -= float64(instruction.cubeCount)
		case 'D':
			cursor.y += float64(instruction.cubeCount)
		case 'L':
			cursor.x -= float64(instruction.cubeCount)
		case 'R':
			cursor.x += float64(instruction.cubeCount)
		}
		l.edgeCubes += int64(instruction.cubeCount)
		l.points = append(l.points, cursor)
	}
}

func (l *LavaLagoon) buildEdgePartTwo(instructions []Instruction) {
	cursor := Point{0, 0}
	l.points = []Point{cursor}
	l.edgeCubes = 1
	for _, instruction := range instructions {
		distance, _ := strconv.ParseInt(instruction.edgeColor[:5], 16, 64)
		direction := instruction.edgeColor[5]
		// 0 means R, 1 means D, 2 means L, and 3 means U
		switch direction {
		case '3':
			cursor.y -= float64(distance)
		case '1':
			cursor.y += float64(distance)
		case '2':
			cursor.x -= float64(distance)
		case '0':
			cursor.x += float64(distance)
		}
		l.edgeCubes += distance
		l.points = append(l.points, cursor)
	}
}

func (l *LavaLagoon) area() int64 {
	n := len(l.points)
	if n < 3 {
		return 0
	}

	// Shoelace formula
	area := 0.0
	for i := 0; i < n-1; i++ {
		area += (l.points[i].x * l.points[i+1].y) - (l.points[i+1].x * l.points[i].y)
	}
	area += (l.points[n-1].x * l.points[0].y) - (l.points[0].x * l.points[n-1].y)

	// The edge cubes increase the area in the following way:
	// For most of the edge cubes, one half is within the area of the polygon,
	// the other is outside of it. So we need to add one half of the edge cubes.
	// For the corner cubes, they contribute either 3/4 or 1/4, depending on whether they are an inner or outer edge.
	// Their number almost evens out.
	// But regardless of how many corner cubes there are, to form a closed area,
	// we must account for 4 additional quarters of a cube.
	return int64(math.Abs(area/2.0)) + (l.edgeCubes)/2 + 1
}

func partOne(instructions []Instruction) int64 {
	lagoon := LavaLagoon{}
	lagoon.buildEdgePartOne(instructions)
	return lagoon.area()
}

func partTwo(instructions []Instruction) int64 {
	lagoon := LavaLagoon{}
	lagoon.buildEdgePartTwo(instructions)
	return lagoon.area()
}

func main() {
	now := time.Now()
	instructions := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(instructions)
	part2 := partTwo(instructions)
	duration := time.Since(now)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseLine(line string) Instruction {
	parts := strings.Split(line, " ")
	direction := parts[0][0]
	cubeCount, _ := strconv.Atoi(parts[1])
	edgeColor := strings.Trim(parts[2], "(#)")
	return Instruction{direction, cubeCount, edgeColor}
}

func parseInput(input string) []Instruction {
	lines := strings.Split(input, "\n")
	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		instructions[i] = parseLine(line)
	}
	return instructions
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
