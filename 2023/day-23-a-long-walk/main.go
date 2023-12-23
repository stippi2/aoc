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

type Segment struct {
	start, end Point
	length     int
}

type Path struct {
	tip      Point
	visited  map[Point]bool
	segments []Segment
	length   int
	foundEnd bool
}

type Map struct {
	width, height int
	tiles         []byte
}

func (m *Map) getTile(x, y int) byte {
	return m.tiles[y*m.width+x]
}

func (m *Map) nextPoints(p Point, visited map[Point]bool, climb bool) []Point {
	tile := m.getTile(p.x, p.y)
	candidates := []Point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
	if !climb {
		switch tile {
		case '>':
			candidates = []Point{{p.x + 1, p.y}}
		case '<':
			candidates = []Point{{p.x - 1, p.y}}
		case '^':
			candidates = []Point{{p.x, p.y - 1}}
		case 'v':
			candidates = []Point{{p.x, p.y + 1}}
		}
	}
	var nextPoints []Point
	for _, c := range candidates {
		if c.x >= 0 && c.x < m.width && c.y >= 0 && c.y < m.height && !visited[c] && m.getTile(c.x, c.y) != '#' {
			nextPoints = append(nextPoints, c)
		}
	}
	return nextPoints
}

func copyMap(m *map[Point]bool) map[Point]bool {
	m2 := make(map[Point]bool)
	for k, v := range *m {
		m2[k] = v
	}
	return m2
}

func (p *Path) follow(m *Map, end Point, climb bool) {
	for {
		nextPoints := m.nextPoints(p.tip, p.visited, climb)
		if len(nextPoints) == 0 {
			break
		}
		if len(nextPoints) > 1 {
			//if climb {
			//	fmt.Printf("Splitting path at %v\n", p.tip)
			//}
			maxLength := 0
			foundEnd := false
			for _, np := range nextPoints {
				next := Path{
					tip:     np,
					visited: copyMap(&p.visited),
					length:  1,
				}
				next.visited[np] = true
				next.follow(m, end, climb)
				if next.foundEnd && next.length > maxLength {
					maxLength = next.length
					foundEnd = true
				}
			}
			p.length += maxLength
			p.foundEnd = foundEnd
			break
		} else {
			p.tip = nextPoints[0]
			p.visited[p.tip] = true
			p.length++
			if p.tip == end {
				p.foundEnd = true
				break
			}
		}
	}
}

func (p *Path) followSegments(m *Map, end Point, climb bool, segmentCache map[Point][]*Segment) {
	if p.tip == end {
		p.foundEnd = true
		return
	}
	nextSegments := segmentCache[p.tip]
	maxLength := 0
	foundEnd := false
	p.visited[p.tip] = true
	for _, segment := range nextSegments {
		if p.visited[segment.end] {
			continue
		}
		next := Path{
			tip:     segment.end,
			visited: copyMap(&p.visited),
			length:  segment.length,
		}
		next.followSegments(m, end, climb, segmentCache)
		if next.foundEnd && next.length > maxLength {
			maxLength = next.length
			foundEnd = true
		}
	}
	p.length += maxLength
	p.foundEnd = foundEnd
}

func (m *Map) getSegments(from, start Point, climb bool, segmentCache map[Point][]*Segment) {
	segment := &Segment{
		start:  start,
		end:    start,
		length: 1,
	}
	visited := map[Point]bool{from: true, start: true}
	pointBeforeEnd := from
	for {
		nextPoints := m.nextPoints(segment.end, visited, climb && start != segment.end)
		if len(nextPoints) == 0 {
			break
		}
		if len(nextPoints) > 1 {
			_, contained := segmentCache[segment.end]
			if !contained {
				for _, np := range nextPoints {
					m.getSegments(segment.end, np, climb, segmentCache)
				}
			}
			break
		} else {
			pointBeforeEnd = segment.end
			segment.end = nextPoints[0]
			visited[segment.end] = true
			segment.length++
		}
	}
	if segment.length > 1 {
		segmentCache[from] = append(segmentCache[from], segment)
		segmentCache[segment.end] = append(segmentCache[segment.end],
			&Segment{start: pointBeforeEnd, end: from, length: segment.length})
	}
}

func (m *Map) findLongestPath(start, end Point, climb bool) int {
	path := &Path{
		tip:     start,
		visited: map[Point]bool{start: true},
		length:  0,
	}
	path.follow(m, end, climb)
	return path.length
}

func partOne(m *Map) int {
	return m.findLongestPath(Point{1, 0}, Point{m.width - 2, m.height - 1}, false)
}

func partTwo(m *Map) int {
	segmentCache := make(map[Point][]*Segment)
	m.getSegments(Point{1, 0}, Point{1, 1}, true, segmentCache)
	for point, segments := range segmentCache {
		fmt.Printf("Point %v:\n", point)
		for _, segment := range segments {
			fmt.Printf("  Segment from %v to %v: %d\n", segment.start, segment.end, segment.length)
		}
	}

	path := &Path{
		tip:     Point{1, 0},
		visited: map[Point]bool{Point{1, 0}: true},
		length:  0,
	}
	path.followSegments(m, Point{m.width - 2, m.height - 1}, true, segmentCache)
	return path.length
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
	lines := strings.Split(strings.TrimSpace(input), "\n")
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
