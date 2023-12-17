package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
	"time"
)

type Map struct {
	width  int
	height int
	data   []uint8
}

func (m *Map) offset(x, y int) int {
	return y*m.width + x
}

func (m *Map) getHeatLoss(x, y int) int {
	return int(m.data[y*m.width+x])
}

type Pos struct {
	x, y int
}

type Path struct {
	direction            byte
	tip                  Pos
	heatLoss             int
	goingStraightCounter int

	distanceToTarget int
	positions        []Pos
}

func (p *Path) String() string {
	xMax := 0
	yMax := 0
	for _, pos := range p.positions {
		if pos.x > xMax {
			xMax = pos.x
		}
		if pos.y > yMax {
			yMax = pos.y
		}
	}
	var sb strings.Builder
	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			onPath := false
			for _, pos := range p.positions {
				if x == pos.x && y == pos.y {
					onPath = true
					break
				}
			}
			if onPath {
				sb.WriteByte('X')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

// PathQueue implements a priority queue, see https://pkg.go.dev/container/heap
type PathQueue []*Path

func (q *PathQueue) Len() int           { return len(*q) }
func (q *PathQueue) Less(i, j int) bool { return (*q)[i].heatLoss < (*q)[j].heatLoss }
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

type Move struct {
	direction            byte
	pos                  Pos
	goingStraightCounter int
}

func possibleMoves(m *Map, path *Path) []Move {
	var moves []Move
	var candidates []Move
	switch path.direction {
	case 0:
		candidates = []Move{
			{'E', Pos{path.tip.x + 1, path.tip.y}, 0},
			{'S', Pos{path.tip.x, path.tip.y + 1}, 0},
		}
	case 'N':
		candidates = []Move{
			{'N', Pos{path.tip.x, path.tip.y - 1}, path.goingStraightCounter + 1},
			{'E', Pos{path.tip.x + 1, path.tip.y}, 0},
			{'W', Pos{path.tip.x - 1, path.tip.y}, 0},
		}
	case 'E':
		candidates = []Move{
			{'E', Pos{path.tip.x + 1, path.tip.y}, path.goingStraightCounter + 1},
			{'N', Pos{path.tip.x, path.tip.y - 1}, 0},
			{'S', Pos{path.tip.x, path.tip.y + 1}, 0},
		}
	case 'W':
		candidates = []Move{
			{'W', Pos{path.tip.x - 1, path.tip.y}, path.goingStraightCounter + 1},
			{'N', Pos{path.tip.x, path.tip.y - 1}, 0},
			{'S', Pos{path.tip.x, path.tip.y + 1}, 0},
		}
	case 'S':
		candidates = []Move{
			{'S', Pos{path.tip.x, path.tip.y + 1}, path.goingStraightCounter + 1},
			{'E', Pos{path.tip.x + 1, path.tip.y}, 0},
			{'W', Pos{path.tip.x - 1, path.tip.y}, 0},
		}
	}
	for _, candidate := range candidates {
		if candidate.pos.x < 0 || candidate.pos.x >= m.width || candidate.pos.y < 0 || candidate.pos.y >= m.height {
			continue
		}
		if candidate.goingStraightCounter > 2 {
			continue
		}
		moves = append(moves, candidate)
	}
	return moves
}

func findPathQueue(m *Map, start, end Pos) int {
	startPath := &Path{
		tip: start,
	}

	type PathKey struct {
		pos                  Pos
		direction            byte
		goingStraightCounter int
	}

	visited := make(map[byte]map[Pos]bool)
	visited['N'] = make(map[Pos]bool)
	visited['S'] = make(map[Pos]bool)
	visited['E'] = make(map[Pos]bool)
	visited['W'] = make(map[Pos]bool)

	pathMap := make(map[PathKey]*Path)

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
			return path.heatLoss
		}
		if path.direction != 0 {
			visited[path.direction][path.tip] = true
		}
		if iteration%100000 == 0 {
			fmt.Printf("iteration: %v, paths: %v, tip: (%v, %v), heat loss: %v\n",
				iteration, queue.Len(), path.tip.x, path.tip.y, path.heatLoss)
		}

		moves := possibleMoves(m, path)

		// For each of the possible directions, create a new path that includes the point taken
		// If that path is better than the path already stored to reach the new point, replace it
		for _, move := range moves {
			heatLoss := path.heatLoss + m.getHeatLoss(move.pos.x, move.pos.y)

			pathKey := PathKey{
				pos:                  move.pos,
				direction:            move.direction,
				goingStraightCounter: move.goingStraightCounter,
			}
			pathToNext := pathMap[pathKey]

			if pathToNext == nil || heatLoss < pathToNext.heatLoss {
				positions := make([]Pos, len(path.positions)+1)
				copy(positions, path.positions)
				positions[len(positions)-1] = move.pos
				pathToNext = &Path{
					direction:            move.direction,
					tip:                  move.pos,
					goingStraightCounter: move.goingStraightCounter,
					heatLoss:             heatLoss,
					positions:            positions,
				}
				pathMap[pathKey] = pathToNext
				heap.Push(queue, pathToNext)
			}
		}
	}
	return -1
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func partOne(m *Map) int {
	return findPathQueue(m, Pos{0, 0}, Pos{m.width - 1, m.height - 1})
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
	lines := strings.Split(input, "\n")
	m := &Map{
		width:  len(lines[0]),
		height: len(lines),
		data:   make([]uint8, len(lines[0])*len(lines)),
	}
	for y, line := range lines {
		for x, c := range line {
			m.data[y*m.width+x] = uint8(c - '0')
		}
	}
	return m
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
