package main

import (
	"fmt"
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

type Pos struct {
	x, y int
}

type LavaLagoon struct {
	xMin, xMax, yMin, yMax int
	width, height          int
	bitmap                 []byte
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func (l *LavaLagoon) allocateBitmap(instructions []Instruction) {
	cursor := Pos{0, 0}
	for _, instruction := range instructions {
		switch instruction.direction {
		case 'U':
			cursor.y -= instruction.cubeCount
		case 'D':
			cursor.y += instruction.cubeCount
		case 'L':
			cursor.x -= instruction.cubeCount
		case 'R':
			cursor.x += instruction.cubeCount
		}
		l.xMin = min(l.xMin, cursor.x)
		l.yMin = min(l.yMin, cursor.y)
		l.xMax = max(l.xMax, cursor.x)
		l.yMax = max(l.yMax, cursor.y)
	}
	l.width = l.xMax - l.xMin + 1
	l.height = l.yMax - l.yMin + 1
	l.bitmap = make([]byte, l.width*l.height)
}

func (l *LavaLagoon) dig(pos Pos) {
	pos.x -= l.xMin
	pos.y -= l.yMin
	l.bitmap[pos.y*l.width+pos.x] = 1
}

func (l *LavaLagoon) isDug(pos Pos) bool {
	pos.x -= l.xMin
	pos.y -= l.yMin
	return l.bitmap[pos.y*l.width+pos.x] == 1
}

func (l *LavaLagoon) buildEdge(instructions []Instruction) {
	cursor := Pos{0, 0}
	for _, instruction := range instructions {
		start := cursor
		switch instruction.direction {
		case 'U':
			cursor.y -= instruction.cubeCount
		case 'D':
			cursor.y += instruction.cubeCount
		case 'L':
			cursor.x -= instruction.cubeCount
		case 'R':
			cursor.x += instruction.cubeCount
		}
		if cursor.x == start.x {
			step := (cursor.y - start.y) / abs(cursor.y-start.y)
			for {
				l.dig(start)
				start.y += step
				if start.y == cursor.y {
					break
				}
			}
		} else {
			step := (cursor.x - start.x) / abs(cursor.x-start.x)
			for {
				l.dig(start)
				start.x += step
				if start.x == cursor.x {
					break
				}
			}
		}
	}
}

func (l *LavaLagoon) fill() {
	start := Pos{1, 1}
	queue := []Pos{start}
	for len(queue) > 0 {
		pos := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		l.dig(pos)
		neighbors := []Pos{
			{pos.x + 1, pos.y},
			{pos.x - 1, pos.y},
			{pos.x, pos.y + 1},
			{pos.x, pos.y - 1},
		}
		for _, n := range neighbors {
			if !l.isDug(n) {
				queue = append(queue, n)
			}
		}
	}
}

func (l *LavaLagoon) area() int {
	area := 0
	for _, cube := range l.bitmap {
		if cube == 1 {
			area++
		}
	}
	return area
}

func partOne(instructions []Instruction) int {
	lagoon := LavaLagoon{}
	lagoon.allocateBitmap(instructions)
	lagoon.buildEdge(instructions)
	lagoon.fill()
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
