package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Galaxy struct {
	x int64
	y int64
}

type Map struct {
	width    int64
	height   int64
	galaxies []*Galaxy
}

func (m *Map) String() string {
	var sb strings.Builder
	for y := int64(0); y < m.height; y++ {
		for x := int64(0); x < m.width; x++ {
			if m.isEmptySpace(x, y) {
				sb.WriteString(".")
			} else {
				sb.WriteString("#")
			}
		}
		if y < m.height-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func (m *Map) isEmptySpace(x int64, y int64) bool {
	for _, galaxy := range m.galaxies {
		if galaxy.x == x && galaxy.y == y {
			return false
		}
	}
	return true
}

func (m *Map) isEmptySpaceY(y int64) bool {
	for _, galaxy := range m.galaxies {
		if galaxy.y == y {
			return false
		}
	}
	return true
}

func (m *Map) isEmptySpaceX(x int64) bool {
	for _, galaxy := range m.galaxies {
		if galaxy.x == x {
			return false
		}
	}
	return true
}

func (m *Map) expandSpace(howMuch int64) {
	for y := int64(0); y < m.height; y++ {
		if m.isEmptySpaceY(y) {
			for _, galaxy := range m.galaxies {
				if galaxy.y > y {
					orig := galaxy.y
					galaxy.y += howMuch
					if galaxy.y < orig {
						panic("overflow")
					}
				}
			}
			m.height += howMuch
			y += howMuch
		}
	}
	for x := int64(0); x < m.width; x++ {
		if m.isEmptySpaceX(x) {
			for _, galaxy := range m.galaxies {
				if galaxy.x > x {
					orig := galaxy.x
					galaxy.x += howMuch
					if galaxy.x < orig {
						panic("overflow")
					}
				}
			}
			m.width += howMuch
			x += howMuch
		}
	}
}

func abs(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}

func (m *Map) sumDistances() (int64, int) {
	distanceSum := int64(0)
	pairs := 0
	for _, galaxy := range m.galaxies {
		for _, otherGalaxy := range m.galaxies {
			if galaxy != otherGalaxy {
				pairs++
				orig := distanceSum
				distanceSum += abs(otherGalaxy.x - galaxy.x)
				if distanceSum < orig {
					panic("overflow")
				}
				orig = distanceSum
				distanceSum += abs(otherGalaxy.y - galaxy.y)
				if distanceSum < orig {
					panic("overflow")
				}
			}
		}
	}
	return distanceSum / 2, pairs / 2
}

func partOne(m *Map) (int64, int) {
	m.expandSpace(1)
	return m.sumDistances()
}

func partTwo(m *Map) (int64, int) {
	m.expandSpace(999999)
	return m.sumDistances()
}

func main() {
	now := time.Now()
	m := parseInput(loadInput("puzzle-input.txt"))
	part1, _ := partOne(m)
	part2, _ := partTwo(m)
	duration := time.Since(now)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseInput(input string) *Map {
	lines := strings.Split(input, "\n")
	m := &Map{
		width:  int64(len(lines[0])),
		height: int64(len(lines)),
	}
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				m.galaxies = append(m.galaxies, &Galaxy{x: int64(x), y: int64(y)})
			}
		}
	}
	return m
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
