package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Map struct {
	width  int
	height int
	data   []byte
}

type Pos struct {
	x int
	y int
}

func (m *Map) get(x, y int) byte {
	if x < 0 || x >= m.width || y < 0 || y >= m.height {
		return ' '
	}
	return m.data[y*m.width+x]
}

func (m *Map) set(x, y int, tile byte) {
	offset := y*m.width + x
	m.data[offset] = tile
}

func (m *Map) getNeighbors(x, y int) []Pos {
	tile := m.get(x, y)
	switch tile {
	case 'F':
		return []Pos{{x + 1, y}, {x, y + 1}}
	case '7':
		return []Pos{{x - 1, y}, {x, y + 1}}
	case 'J':
		return []Pos{{x, y - 1}, {x - 1, y}}
	case 'L':
		return []Pos{{x, y - 1}, {x + 1, y}}
	case '|':
		return []Pos{{x, y - 1}, {x, y + 1}}
	case '-':
		return []Pos{{x - 1, y}, {x + 1, y}}
	case 'S':
		return []Pos{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
	}
	return nil
}

func contains(positions []Pos, pos Pos) bool {
	for _, p := range positions {
		if p == pos {
			return true
		}
	}
	return false
}

func (m *Map) getStart() Pos {
	for i := 0; i < len(m.data); i++ {
		if m.data[i] == 'S' {
			return Pos{i % m.width, i / m.width}
		}
	}
	return Pos{-1, -1}
}

type Path struct {
	positions []Pos
}

func (p *Path) tip() Pos {
	return p.positions[len(p.positions)-1]
}

func (p *Path) previous() Pos {
	return p.positions[len(p.positions)-2]
}

func getPaths(m *Map) (start Pos, left, right *Path) {
	start = m.getStart()
	var paths []*Path
	for _, n := range m.getNeighbors(start.x, start.y) {
		if contains(m.getNeighbors(n.x, n.y), start) {
			paths = append(paths, &Path{[]Pos{start, n}})
		}
	}
	if len(paths) != 2 {
		panic("Expected two paths")
	}
	return start, paths[0], paths[1]
}

func partOne(m *Map) int {
	start, left, _ := getPaths(m)
	for {
		neighbors := m.getNeighbors(left.tip().x, left.tip().y)
		for _, n := range neighbors {
			if left.previous() != n {
				left.positions = append(left.positions, n)
				break
			}
		}
		if left.tip() == start {
			break
		}
	}

	return len(left.positions) / 2
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
	fmt.Printf("Part 1: Longest distance to S: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseInput(input string) *Map {
	lines := strings.Split(input, "\n")
	return &Map{
		width:  len(lines[0]),
		height: len(lines),
		data:   []byte(strings.Join(lines, "")),
	}
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
