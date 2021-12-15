package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

type Map struct {
	width  int
	height int
	data   []int
}

func (m *Map) init(width, height int) {
	m.width = width
	m.height = height
	m.data = make([]int, width*height)
}

func (m *Map) offset(x, y int) int {
	return m.width*y + x
}

func (m *Map) set(x, y, value int) {
	offset := m.offset(x, y)
	if offset >= 0 && offset < len(m.data) {
		m.data[offset] = value
	}
}

func (m *Map) get(x, y int) int {
	return m.data[m.offset(x, y)]
}

func (m *Map) String() string {
	result := ""
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width-1; x++ {
			result += fmt.Sprintf("%2d ", m.get(x, y))
		}
		result += fmt.Sprintf("%2d\n", m.get(m.width-1, y))
	}
	return result
}

type Path struct {
	points []Point
	risk int
}

func (p *Path) contains(point Point) bool {
	for _, i := range p.points {
		if i == point {
			return true
		}
	}
	return false
}

func (p *Path) tip() Point {
	return p.points[len(p.points) - 1]
}

func (m *Map) neighbors(tip Point) []Point {
	var next []Point
	if tip.x < m.width - 1 {
		next = append(next, Point{tip.x + 1, tip.y})
	}
	if tip.y < m.height - 1 {
		next = append(next, Point{tip.x, tip.y + 1})
	}
	if tip.x > 0 {
		next = append(next, Point{tip.x - 1, tip.y})
	}
	if tip.y > 0 {
		next = append(next, Point{tip.x, tip.y - 1})
	}
	return next
}

func (p *Path) String() string {
	result := ""
	for _, point := range p.points {
		result += fmt.Sprintf(", (%v, %v)", point.x, point.y)
	}
	result += fmt.Sprintf("  risk: %v", p.risk)
	return result[1:]
}

func (m *Map) extend(count int) *Map {
	newMap := Map{}
	newMap.init(m.width * count, m.height * count)
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			risk := m.get(x, y)
			for ry := 0; ry < count; ry++ {
				for rx := 0; rx < count; rx++ {
					newRisk := risk + rx + ry
					if newRisk > 9 {
						newRisk -= 9
					}
					newMap.set(x+(rx * m.width), y+(ry * m.height), newRisk)
				}
			}
		}
	}
	return &newMap
}

func (m *Map) findPath() int {
	start := Point{0, 0}
	end := Point{m.width - 1, m.height - 1}

	bestPaths := make(map[Point]*Path)
	bestPaths[start] = &Path{
		points: []Point{{0, 0}},
		risk:   0,
	}
	iteration := 0

	for {
		currentWinner := bestPaths[end]
		if currentWinner != nil {
			// remove any paths with worse risk than the current winner
			for tip, path := range bestPaths {
				if path != currentWinner && path.risk >= currentWinner.risk {
					delete(bestPaths, tip)
				}
			}
			if len(bestPaths) == 1 {
				return currentWinner.risk
			}
		}
		iteration++
		fmt.Printf("iteration: %v, paths: %v\n", iteration, len(bestPaths))
		newBestPaths := make(map[Point]*Path)
		for tip, path := range bestPaths {
			//fmt.Printf("path: %v\n", path)
			newBestPaths[tip] = path
		}
		for tip, path := range bestPaths {
			neighbors := m.neighbors(tip)

			// For each of the possible directions, create a new path that includes the point taken
			// If that path is better than the path already stored to reach the new point, replace it
			for _, n := range neighbors {
				// Never enter a cycle
				if path.contains(n) {
					continue
				}
				// Compare the current best path to reach the next point
				bestPath := newBestPaths[n]
				if bestPath != nil && bestPath.risk < path.risk + m.get(n.x, n.y) {
					// Ok, there is a better path to reach this point
					continue
				}
				pathCopy := make([]Point, len(path.points) + 1)
				copy(pathCopy, path.points)
				pathCopy[len(path.points)] = n
				pathToNext := &Path{
					points: pathCopy,
					risk:   path.risk + m.get(n.x, n.y),
				}
				newBestPaths[n] = pathToNext
			}
			if tip != end {
				delete(newBestPaths, tip)
			}
		}
		bestPaths = newBestPaths
	}
}

func main() {
	m := parseInput(loadInput("puzzle-input.txt"))
	m = m.extend(5)
	start := time.Now()
	leastRisk := m.findPath()
	fmt.Printf("least risk: %v, found in %v\n", leastRisk, time.Since(start))
}

func parseInput(input string) *Map {
	lines := strings.Split(input, "\n")
	height := len(lines)
	if height == 0 {
		return nil
	}
	width := len(lines[0])
	if width == 0 {
		return nil
	}
	m := Map{}
	m.init(width, height)
	for y := 0; y < height; y++ {
		for x, char := range strings.Split(lines[y], "") {
			if x >= width {
				break
			}
			v, _ := strconv.Atoi(char)
			m.set(x, y, v)
		}
	}
	return &m
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
