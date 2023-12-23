package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

type Path struct {
	tip       Point
	visited   map[Point]bool
	length    int
	nextPaths []Path
}

type Map struct {
	width, height int
	tiles         []byte
}

func (m *Map) getTile(x, y int) byte {
	return m.tiles[y*m.width+x]
}

func (p *Path) nextPoints(m *Map) []Point {
	tile := m.getTile(p.tip.x, p.tip.y)
	var candidates []Point
	switch tile {
	case '>':
		candidates = []Point{{p.tip.x + 1, p.tip.y}}
	case '<':
		candidates = []Point{{p.tip.x - 1, p.tip.y}}
	case '^':
		candidates = []Point{{p.tip.x, p.tip.y - 1}}
	case 'v':
		candidates = []Point{{p.tip.x, p.tip.y + 1}}
	case '.':
		candidates = []Point{
			{p.tip.x - 1, p.tip.y},
			{p.tip.x + 1, p.tip.y},
			{p.tip.x, p.tip.y - 1},
			{p.tip.x, p.tip.y + 1},
		}
	}
	var nextPoints []Point
	for _, c := range candidates {
		if c.x >= 0 && c.x < m.width && c.y >= 0 && c.y < m.height && !p.visited[c] && m.getTile(c.x, c.y) != '#' {
			nextPoints = append(nextPoints, c)
		}
	}
	return nextPoints
}

func (p *Path) follow(m *Map) {
	for {
		nextPoints := p.nextPoints(m)
		if len(nextPoints) == 0 {
			break
		}
		if len(nextPoints) > 1 {
			maxLength := 0
			for _, np := range nextPoints {
				next := Path{
					tip:     np,
					visited: map[Point]bool{np: true, p.tip: true},
					length:  1,
				}
				next.follow(m)
				p.nextPaths = append(p.nextPaths, next)
				if next.length > maxLength {
					maxLength = next.length
				}
			}
			p.length += maxLength
			break
		} else {
			p.tip = nextPoints[0]
			p.visited[p.tip] = true
			p.length++
		}
	}
}

func (m *Map) findLongestPath(start, end Point) int {
	path := &Path{
		tip:     start,
		visited: map[Point]bool{start: true},
		length:  0,
	}
	path.follow(m)
	return path.length
}

func partOne(m *Map) int {
	return m.findLongestPath(Point{1, 0}, Point{m.width - 2, m.height - 1})
}

func partTwo(m *Map) int {
	return 0
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
		width:  len(lines[0]),
		height: len(lines),
		tiles:  []byte(strings.Join(lines, "")),
	}
	return m
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
