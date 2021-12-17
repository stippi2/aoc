package main

import (
	"container/heap"
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

type Path struct {
	points []Point
	tip    Point
	risk   int
}

func (p *Path) contains(point Point) bool {
	for _, i := range p.points {
		if i == point {
			return true
		}
	}
	return false
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

func (m *Map) neighbors(p Point) []Point {
	var neighbors []Point
	if p.x < m.width-1 {
		neighbors = append(neighbors, Point{p.x + 1, p.y})
	}
	if p.y < m.height-1 {
		neighbors = append(neighbors, Point{p.x, p.y + 1})
	}
	if p.x > 0 {
		neighbors = append(neighbors, Point{p.x - 1, p.y})
	}
	if p.y > 0 {
		neighbors = append(neighbors, Point{p.x, p.y - 1})
	}
	return neighbors
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
	newMap.init(m.width*count, m.height*count)
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			risk := m.get(x, y)
			for ry := 0; ry < count; ry++ {
				for rx := 0; rx < count; rx++ {
					newRisk := risk + rx + ry
					if newRisk > 9 {
						newRisk -= 9
					}
					newMap.set(x+(rx*m.width), y+(ry*m.height), newRisk)
				}
			}
		}
	}
	return &newMap
}

func copyMap(m map[Point]*Path) map[Point]*Path {
	c := make(map[Point]*Path)
	for point, path := range m {
		c[point] = path
	}
	return c
}

func (m *Map) findPath() int {
	start := Point{0, 0}
	end := Point{m.width - 1, m.height - 1}

	// bestPaths stores for each Point the lowest risk Path currently known to reach that point
	bestPaths := make(map[Point]*Path)
	bestPaths[start] = &Path{
		points: []Point{{0, 0}},
		risk:   0,
	}
	iteration := 0

	// TODO: optimization idea 1: reach the end diagonally, any path needs to be better than that
	// TODO: Use another data structure to hold the "in-flight" paths. If it is ordered for the
	// TODO: lowest risk path first, anytime we reach a new point means we have reached it via the
	// TODO: cheapest path. This means we do not need to reach this point via any other path, and
	// TODO: so we can maintain a simple "visited" map of points. We don't need to check from any
	// TODO: Path whether we already visited a new neighbor.

	for {
		// Make a copy of bestPaths map, so that we can iterate a stable map
		// and modify the copy. See https://go.dev/ref/spec#For_range
		newBestPaths := copyMap(bestPaths)

		currentWinner := bestPaths[end]
		if currentWinner != nil {
			// Remove any paths with worse risk than the current winner
			for tip, path := range bestPaths {
				if path != currentWinner && path.risk >= currentWinner.risk {
					delete(newBestPaths, tip)
				}
			}
			if len(newBestPaths) == 1 {
				return currentWinner.risk
			}
			// Sync the maps again after having updated the copy
			bestPaths = copyMap(newBestPaths)
		}

		iteration++
		fmt.Printf("iteration: %v, paths: %v\n", iteration, len(bestPaths))

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
				if bestPath != nil && bestPath.risk < path.risk+m.get(n.x, n.y) {
					// Ok, there is a better path to reach this point
					continue
				}
				pathCopy := make([]Point, len(path.points)+1)
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

// PathQueue implements a priority queue, see https://pkg.go.dev/container/heap
type PathQueue []*Path

func (q PathQueue) Len() int           { return len(q) }
func (q PathQueue) Less(i, j int) bool { return q[i].risk < q[j].risk }
func (q PathQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *PathQueue) Push(x interface{}) {
	path := x.(*Path)
	*q = append(*q, path)
}

func (q *PathQueue) Pop() interface{} {
	old := *q
	n := len(old)
	path := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*q = old[0 : n-1]
	return path
}

func (m *Map) findPathQueue() int {
	startPath := &Path{
		risk: 0,
		tip:  Point{0, 0},
	}
	end := Point{m.width - 1, m.height - 1}

	pathMap := make(map[Point]*Path)
	visited := make(map[Point]bool)

	pathMap[startPath.tip] = startPath
	queue := &PathQueue{startPath}
	heap.Init(queue)

	iteration := 0
	for queue.Len() > 0 {
		iteration++
		path := heap.Pop(queue).(*Path)
		if path.tip == end {
			return path.risk
		}
		visited[path.tip] = true
		//fmt.Printf("iteration: %v, paths: %v, tip: (%v, %v), risk: %v\n", iteration, queue.Len(),
		//	path.tip.x, path.tip.y, path.risk)

		neighbors := m.neighbors(path.tip)

		// For each of the possible directions, create a new path that includes the point taken
		// If that path is better than the path already stored to reach the new point, replace it
		for _, n := range neighbors {
			// If we visited this position already, it means we did so via a cheaper path
			if visited[n] {
				continue
			}
			risk := path.risk + m.get(n.x, n.y)
			pathToNext := pathMap[n]
			if pathToNext == nil {
				pathToNext = &Path{
					risk: risk,
					tip:  n,
				}
				pathMap[n] = pathToNext
				queue.Push(pathToNext)
			} else if risk < pathToNext.risk {
				pathToNext.risk = risk
				queue.Push(pathToNext)
				//queue.updatePath(pathToNext, n, risk)
			}
		}
	}
	return -1
}

func main() {
	m := parseInput(loadInput("puzzle-input.txt"))
	m = m.extend(5)
	start := time.Now()
	leastRisk := m.findPathQueue()
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
