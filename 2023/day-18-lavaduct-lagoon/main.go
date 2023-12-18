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
	edgeCubes int
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
		l.edgeCubes += instruction.cubeCount
		l.points = append(l.points, cursor)
	}
}

func (l *LavaLagoon) area() int {
	n := len(l.points)
	if n < 3 {
		return 0
	}

	area := 0.0
	for i := 0; i < n-1; i++ {
		area += (l.points[i].x * l.points[i+1].y) - (l.points[i+1].x * l.points[i].y)
	}
	area += (l.points[n-1].x * l.points[0].y) - (l.points[0].x * l.points[n-1].y)

	return int(math.Abs(area/2.0)) + (l.edgeCubes)/2 + 1
}

func partOne(instructions []Instruction) int {
	lagoon := LavaLagoon{}
	lagoon.buildEdgePartOne(instructions)
	return lagoon.area()
}

func partTwo() int {
	return 0
}

func main() {
	now := time.Now()
	instructions := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(instructions)
	part2 := partTwo()
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
