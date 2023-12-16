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

type PosAndDirection struct {
	pos       Pos
	direction Pos
}

type Map struct {
	width  int
	height int
	tiles  map[Pos]byte
}

func contains(positions []PosAndDirection, pos PosAndDirection) bool {
	for _, p := range positions {
		if p == pos {
			return true
		}
	}
	return false
}

type VisitedTiles struct {
	north map[Pos]bool
	south map[Pos]bool
	east  map[Pos]bool
	west  map[Pos]bool
}

func (vt *VisitedTiles) add(pos, direction Pos) {
	switch direction {
	case Pos{0, -1}:
		vt.north[pos] = true
	case Pos{0, 1}:
		vt.south[pos] = true
	case Pos{1, 0}:
		vt.east[pos] = true
	case Pos{-1, 0}:
		vt.west[pos] = true
	}
}

func (vt *VisitedTiles) countVisitedTiles() int {
	visited := map[Pos]bool{}
	for pos := range vt.north {
		visited[pos] = true
	}
	for pos := range vt.south {
		visited[pos] = true
	}
	for pos := range vt.east {
		visited[pos] = true
	}
	for pos := range vt.west {
		visited[pos] = true
	}
	return len(visited)
}

func (vt *VisitedTiles) contains(pos, direction Pos) bool {
	switch direction {
	case Pos{0, -1}:
		return vt.north[pos]
	case Pos{0, 1}:
		return vt.south[pos]
	case Pos{1, 0}:
		return vt.east[pos]
	case Pos{-1, 0}:
		return vt.west[pos]
	}
	return false
}

func (m *Map) traceLight(start, direction Pos, visitedTiles *VisitedTiles) {
	tip := start
	for {
		tip = Pos{tip.x + direction.x, tip.y + direction.y}
		// Detect beam leaving the map
		if tip.x < 0 || tip.x >= m.width || tip.y < 0 || tip.y >= m.height {
			break
		}
		// Detect cycles
		if visitedTiles.contains(tip, direction) {
			break
		}

		visitedTiles.add(tip, direction)

		tile, isObject := m.tiles[tip]
		if !isObject {
			continue
		}
		switch tile {
		case '|':
			if direction.y == 0 {
				m.traceLight(tip, Pos{0, -1}, visitedTiles)
				m.traceLight(tip, Pos{0, 1}, visitedTiles)
				return
			}
		case '-':
			if direction.x == 0 {
				m.traceLight(tip, Pos{-1, 0}, visitedTiles)
				m.traceLight(tip, Pos{1, 0}, visitedTiles)
				return
			}
		case '/':
			if direction.x == 0 && direction.y == 1 {
				direction.x = -1
				direction.y = 0
			} else if direction.x == 0 && direction.y == -1 {
				direction.x = 1
				direction.y = 0
			} else if direction.x == 1 && direction.y == 0 {
				direction.x = 0
				direction.y = -1
			} else if direction.x == -1 && direction.y == 0 {
				direction.x = 0
				direction.y = 1
			}
		case '\\':
			if direction.x == 0 && direction.y == 1 {
				direction.x = 1
				direction.y = 0
			} else if direction.x == 0 && direction.y == -1 {
				direction.x = -1
				direction.y = 0
			} else if direction.x == 1 && direction.y == 0 {
				direction.x = 0
				direction.y = 1
			} else if direction.x == -1 && direction.y == 0 {
				direction.x = 0
				direction.y = -1
			}
		}
	}
}

func partOne(m *Map) int {
	visitedTiles := VisitedTiles{
		north: make(map[Pos]bool),
		south: make(map[Pos]bool),
		east:  make(map[Pos]bool),
		west:  make(map[Pos]bool),
	}
	m.traceLight(Pos{-1, 0}, Pos{1, 0}, &visitedTiles)
	return visitedTiles.countVisitedTiles()
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
	lines := strings.Split(strings.TrimSpace(input), "\n")
	m := &Map{
		width:  len(lines[0]),
		height: len(lines),
		tiles:  make(map[Pos]byte),
	}
	for y, line := range lines {
		for x, tile := range line {
			if tile != '.' {
				m.tiles[Pos{x, y}] = byte(tile)
			}
		}
	}
	return m
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
