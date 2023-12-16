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
	width           int
	height          int
	roundRocks      map[Pos]bool
	squareRocks     map[Pos]bool
	cyclesCompleted int
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

func (m *Map) tiltWest() map[Pos]bool {
	result := map[Pos]bool{}
	for y := 0; y < m.height; y++ {
		x := 0
		xStart := x
		roundRocks := 0
		for x < m.width {
			if m.squareRocks[Pos{x, y}] {
				for roundRocks > 0 {
					result[Pos{xStart, y}] = true
					xStart++
					roundRocks--
				}
				roundRocks = 0
				xStart = x + 1
			} else if m.roundRocks[Pos{x, y}] {
				roundRocks++
			}
			x++
		}
		for roundRocks > 0 {
			result[Pos{xStart, y}] = true
			xStart++
			roundRocks--
		}
	}
	return result
}

func (m *Map) tiltEast() map[Pos]bool {
	result := map[Pos]bool{}
	for y := 0; y < m.height; y++ {
		x := m.width - 1
		xStart := x
		roundRocks := 0
		for x >= 0 {
			if m.squareRocks[Pos{x, y}] {
				for roundRocks > 0 {
					result[Pos{xStart, y}] = true
					xStart--
					roundRocks--
				}
				roundRocks = 0
				xStart = x - 1
			} else if m.roundRocks[Pos{x, y}] {
				roundRocks++
			}
			x--
		}
		for roundRocks > 0 {
			result[Pos{xStart, y}] = true
			xStart--
			roundRocks--
		}
	}
	return result
}

func (m *Map) tiltNorth() map[Pos]bool {
	result := map[Pos]bool{}
	for x := 0; x < m.width; x++ {
		y := 0
		yStart := y
		roundRocks := 0
		for y < m.height {
			if m.squareRocks[Pos{x, y}] {
				for roundRocks > 0 {
					result[Pos{x, yStart}] = true
					yStart++
					roundRocks--
				}
				roundRocks = 0
				yStart = y + 1
			} else if m.roundRocks[Pos{x, y}] {
				roundRocks++
			}
			y++
		}
		for roundRocks > 0 {
			result[Pos{x, yStart}] = true
			yStart++
			roundRocks--
		}
	}
	return result
}

func (m *Map) tiltSouth() map[Pos]bool {
	result := map[Pos]bool{}
	for x := 0; x < m.width; x++ {
		y := m.height - 1
		yStart := y
		roundRocks := 0
		for y >= 0 {
			if m.squareRocks[Pos{x, y}] {
				for roundRocks > 0 {
					result[Pos{x, yStart}] = true
					yStart--
					roundRocks--
				}
				roundRocks = 0
				yStart = y - 1
			} else if m.roundRocks[Pos{x, y}] {
				roundRocks++
			}
			y--
		}
		for roundRocks > 0 {
			result[Pos{x, yStart}] = true
			yStart--
			roundRocks--
		}
	}
	return result
}

func equals(a, b map[Pos]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for pos := range a {
		if !b[pos] {
			return false
		}
	}
	return true
}

func (m *Map) tiltCycle() {
	m.roundRocks = m.tiltNorth()
	m.roundRocks = m.tiltWest()
	m.roundRocks = m.tiltSouth()
	m.roundRocks = m.tiltEast()
	m.cyclesCompleted++
}

func runTiltCycles(m *Map, repeatCount int, detectCycle bool) {
	roundRocksSnapshots := []map[Pos]bool{m.roundRocks}
	foundCycle := !detectCycle
	for i := m.cyclesCompleted; i < repeatCount; i++ {
		fmt.Printf("Tilt cycle %d\n", i)
		m.tiltCycle()
		if !foundCycle {
			for cycleLength := 1; cycleLength < len(roundRocksSnapshots); cycleLength++ {
				if equals(m.roundRocks, roundRocksSnapshots[len(roundRocksSnapshots)-cycleLength]) {
					fmt.Printf("Found cycle at %d, cycle length: %d\n", i, cycleLength)
					i += ((repeatCount - i) / cycleLength) * cycleLength
					foundCycle = true
					break
				}
			}
			roundRocksSnapshots = append(roundRocksSnapshots, m.roundRocks)
		}
	}
}

func (m *Map) calcLoadNorth() int {
	result := 0
	for rock := range m.roundRocks {
		result += m.height - rock.y
	}
	return result
}

func partTwo(m *Map) int {
	runTiltCycles(m, 1000000000, true)
	return m.calcLoadNorth()
}

func main() {
	now := time.Now()
	m := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(m)
	part2 := partTwo(m)
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
