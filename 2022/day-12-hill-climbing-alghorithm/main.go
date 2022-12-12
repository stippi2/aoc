package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
	"time"
)

type Pos struct {
	x, y int
}

type Path struct {
	positions []Pos
	tip       Pos
	cost      int
}

type Elevation struct {
	height int
}

type Map struct {
	grid   []Elevation
	width  int
	height int
}

func (m *Map) init(width, height int) {
	m.grid = make([]Elevation, width*height)
	m.width = width
	m.height = height
}

func (m *Map) offset(x, y int) int {
	return y*m.width + x
}

func (m *Map) get(x, y int) *Elevation {
	return &m.grid[m.offset(x, y)]
}

// PathQueue implements a priority queue, see https://pkg.go.dev/container/heap
type PathQueue []*Path

func (q *PathQueue) Len() int           { return len(*q) }
func (q *PathQueue) Less(i, j int) bool { return (*q)[i].cost < (*q)[j].cost }
func (q *PathQueue) Swap(i, j int)      { (*q)[i], (*q)[j] = (*q)[j], (*q)[i] }

func (q *PathQueue) Push(x interface{}) {
	path := x.(*Path)
	*q = append(*q, path)
}

func (q *PathQueue) Pop() interface{} {
	old := *q
	n := len(old)
	path := old[n-1]
	old[n-1] = nil // avoid memory leak
	*q = old[0 : n-1]
	return path
}

func isPossibleDirection(m *Map, p Pos, currentElevation int) bool {
	if p.x < 0 || p.x >= m.width || p.y < 0 || p.y >= m.height {
		return false
	}
	newElevation := m.get(p.x, p.y).height
	return newElevation-currentElevation <= 1
}

func possibleDirections(m *Map, p Pos) []Pos {
	var neighbors []Pos
	elevation := m.get(p.x, p.y).height
	candidates := []Pos{{p.x + 1, p.y}, {p.x, p.y + 1}, {p.x - 1, p.y}, {p.x, p.y - 1}}
	for _, candidate := range candidates {
		if isPossibleDirection(m, candidate, elevation) {
			neighbors = append(neighbors, candidate)
		}
	}
	return neighbors
}

func findPathQueue(m *Map, start, end Pos) int {
	startPath := &Path{
		positions: []Pos{start},
		tip:       start,
	}

	pathMap := make([]*Path, m.width*m.height)
	visited := make(map[Pos]bool)

	pathMap[m.offset(start.x, start.y)] = startPath
	queue := &PathQueue{startPath}
	heap.Init(queue)

	startTime := time.Now()
	iteration := 0

	for queue.Len() > 0 {
		iteration++
		path := heap.Pop(queue).(*Path)
		if path.tip == end {
			fmt.Printf("found end after %v iterations, paths in queue: %v\n", iteration, queue.Len())
			fmt.Printf("found end after %v / %v iterations, paths in map: %v\n", time.Since(startTime),
				iteration, queue.Len())
			return len(path.positions) - 1 // "steps" excludes starting point
		}
		visited[path.tip] = true
		if iteration%100000 == 0 {
			fmt.Printf("iteration: %v, paths: %v, tip: (%v, %v), risk: %v\n",
				iteration, queue.Len(), path.tip.x, path.tip.y, path.cost)
		}

		neighbors := possibleDirections(m, path.tip)

		// For each of the possible directions, create a new path that includes the point taken
		// If that path is better than the path already stored to reach the new point, replace it
		for _, n := range neighbors {
			// If we visited this position already, it means we did so via a cheaper path
			if visited[n] {
				continue
			}
			cost := path.cost + 1
			pathMapOffset := m.offset(n.x, n.y)
			pathToNext := pathMap[pathMapOffset]
			if pathToNext == nil {
				pathToNext = &Path{
					positions: append(path.positions, n),
					cost:      cost,
					tip:       n,
				}
				pathMap[pathMapOffset] = pathToNext
				heap.Push(queue, pathToNext)
			} else if cost < pathToNext.cost {
				pathToNext.cost = cost
				heap.Push(queue, pathToNext)
			}
		}
	}
	return -1
}

func main() {
	// Part 1
	m, start, dest := parseInput(loadInput("puzzle-input.txt"))
	shortestPathFromStart := findPathQueue(m, start, dest)
	fmt.Printf("shortest path from starting point: %v\n", shortestPathFromStart)
	// Part 2
	shortest := shortestPathFromStart
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if m.get(x, y).height == 0 {
				pathLength := findPathQueue(m, Pos{x, y}, dest)
				if pathLength > 0 && pathLength < shortest {
					shortest = pathLength
				}
			}
		}
	}
	fmt.Printf("shortest path from any low-point: %v\n", shortest)
}

func parseInput(input string) (m *Map, start, dest Pos) {
	lines := strings.Split(input, "\n")
	m = &Map{}
	m.init(len(lines[0]), len(lines))
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			height := 0
			if line[x] == 'S' {
				start = Pos{x, y}
			} else if line[x] == 'E' {
				height = int('z'-'a') + 1
				dest = Pos{x, y}
			} else {
				height = int(line[x] - 'a')
			}
			m.get(x, y).height = height
		}
	}
	return
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
