package main

import (
	"bytes"
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

func (m *Map) String() string {
	var sb strings.Builder
	for i, c := range m.data {
		sb.WriteByte(c)
		if i%m.width == m.width-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
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
	case 'S', '.':
		return []Pos{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
	}
	return nil
}

func (m *Map) doubleSize() *Map {
	newMap := &Map{
		width:  m.width * 2,
		height: m.height * 2,
		data:   bytes.Repeat([]byte{'.'}, m.width*m.height*4),
	}

	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			current := m.get(x, y)
			newX, newY := x*2, y*2
			switch current {
			case '|':
				newMap.set(newX, newY, '|')
				newMap.set(newX, newY+1, '|')
				newMap.set(newX+1, newY, '.')
				newMap.set(newX+1, newY+1, '.')
			case '-':
				newMap.set(newX, newY, '-')
				newMap.set(newX, newY+1, '.')
				newMap.set(newX+1, newY, '-')
				newMap.set(newX+1, newY+1, '.')
			case 'L':
				newMap.set(newX, newY, 'L')
				newMap.set(newX, newY+1, '.')
				newMap.set(newX+1, newY, '-')
				newMap.set(newX+1, newY+1, '.')
			case 'J':
				newMap.set(newX, newY, 'J')
				newMap.set(newX, newY+1, '.')
				newMap.set(newX+1, newY, '.')
				newMap.set(newX+1, newY+1, '.')
			case '7':
				newMap.set(newX, newY, '7')
				newMap.set(newX, newY+1, '|')
				newMap.set(newX+1, newY, '.')
				newMap.set(newX+1, newY+1, '.')
			case 'F':
				newMap.set(newX, newY, 'F')
				newMap.set(newX, newY+1, '|')
				newMap.set(newX+1, newY, '-')
				newMap.set(newX+1, newY+1, '.')
			case '.':
				newMap.set(newX, newY, '.')
				newMap.set(newX, newY+1, '.')
				newMap.set(newX+1, newY, '.')
				newMap.set(newX+1, newY+1, '.')
			case 'S':
				newMap.set(newX, newY, 'S')
				newMap.set(newX, newY+1, 'S')
				newMap.set(newX+1, newY, 'S')
				newMap.set(newX+1, newY+1, 'S')
			}
		}
	}

	return newMap
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

func (m *Map) fill(x, y int, tile byte) {
	if m.get(x, y) != '.' {
		return
	}
	queue := []Pos{{x, y}}
	for len(queue) > 0 {
		tip := queue[0]
		queue = queue[1:]
		neighbors := m.getNeighbors(tip.x, tip.y)
		m.set(tip.x, tip.y, tile)
		for _, n := range neighbors {
			if m.get(n.x, n.y) == '.' {
				queue = append(queue, n)
			}
		}
	}
}

func partTwo(m *Map) int {
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

	// Get a cleaned up version of the map with just the loop
	converted := &Map{
		width:  m.width,
		height: m.height,
		data:   bytes.Repeat([]byte{'.'}, len(m.data)),
	}
	for _, p := range left.positions {
		converted.set(p.x, p.y, m.get(p.x, p.y))

	}

	// Create a double-sized version, so we get gaps between the loop everywhere it touches itself
	doubleSize := converted.doubleSize()

	// Fill the gaps with 'O' from the outside edges
	doubleSize.fill(doubleSize.width-1, doubleSize.height-1, 'O')
	for x := 0; x < doubleSize.width; x++ {
		doubleSize.fill(x, 0, 'O')
	}
	for y := 0; y < doubleSize.height; y++ {
		doubleSize.fill(0, y, 'O')
	}

	// Transition the fills to the original size
	for y := 0; y < doubleSize.height; y += 2 {
		for x := 0; x < doubleSize.width; x += 2 {
			if doubleSize.get(x, y) == 'O' {
				converted.set(x/2, y/2, 'O')
			}
		}
	}

	// Count the remaining '.' inside the loop
	inside := 0
	for i := 0; i < len(converted.data); i++ {
		if converted.data[i] == '.' {
			inside++
		}
	}

	return inside
}

func main() {
	now := time.Now()
	m := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(m)
	part2 := partTwo(m)
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
