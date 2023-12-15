package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Pos struct {
	x int
	y int
}

type Map struct {
	width       int
	height      int
	roundRocks  map[Pos]bool
	squareRocks map[Pos]bool
}

func addRoundRocksWeight(result, rowsToSouth, roundRocks int) int {
	for roundRocks > 0 {
		result += rowsToSouth
		rowsToSouth--
		roundRocks--
	}
	return result
}

func partOne(m *Map) int {
	result := 0
	for x := 0; x < m.width; x++ {
		y := 0
		yStart := y
		roundRocks := 0
		for y < m.height {
			if m.squareRocks[Pos{x, y}] {
				rowsToSouth := m.height - yStart
				result = addRoundRocksWeight(result, rowsToSouth, roundRocks)
				roundRocks = 0
				yStart = y + 1
			} else if m.roundRocks[Pos{x, y}] {
				roundRocks++
			}
			y++
		}
		rowsToSouth := m.height - yStart
		result = addRoundRocksWeight(result, rowsToSouth, roundRocks)
	}
	return result
}

func (m *Map) tiltNorth() *Map {
	result := &Map{
		width:       m.width,
		height:      m.height,
		roundRocks:  map[Pos]bool{},
		squareRocks: m.squareRocks,
	}
	for x := 0; x < m.width; x++ {
		y := 0
		yStart := y
		roundRocks := 0
		for y < m.height {
			if m.squareRocks[Pos{x, y}] {
				rowsToSouth := m.height - yStart
				for roundRocks > 0 {
					rowsToSouth--
					roundRocks--
				}
				roundRocks = 0
				yStart = y + 1
			} else if m.roundRocks[Pos{x, y}] {
				roundRocks++
			}
			y++
		}
		rowsToSouth := m.height - yStart
		for roundRocks > 0 {
			rowsToSouth--
			roundRocks--
		}
	}
	return result
}

func partTwo() int {
	return 0
}

func main() {
	now := time.Now()
	m := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(m)
	part2 := partTwo()
	duration := time.Since(now)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseInput(input string) *Map {
	lines := strings.Split(input, "\n")
	m := &Map{
		width:       len(lines[0]),
		height:      len(lines),
		roundRocks:  map[Pos]bool{},
		squareRocks: map[Pos]bool{},
	}
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				m.squareRocks[Pos{x, y}] = true
			} else if char == 'O' {
				m.roundRocks[Pos{x, y}] = true
			}
		}
	}
	return m
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
