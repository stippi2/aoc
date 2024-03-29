package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Pos struct {
	x, y int
}

func (p Pos) add(vector Pos) Pos {
	return Pos{p.x + vector.x, p.y + vector.y}
}

type Path struct {
	positions  []Pos
	tip        Pos
	end        Pos
	backTracks int
}

func (p *Path) nextMinute() int {
	return len(p.positions)
}

func (p *Path) cost() int {
	totalDistance := p.end.x*p.end.x + p.end.y*p.end.y
	distanceToEnd := (p.end.x-p.tip.x)*(p.end.x-p.tip.x) + (p.end.y-p.tip.y)*(p.end.y-p.tip.y)
	return len(p.positions)*totalDistance + distanceToEnd/100 + p.backTracks*totalDistance/10
}

type Blizzard struct {
	pos       Pos
	direction Pos
}

type Map struct {
	width, height  int
	emptyPositions []map[Pos]bool
	blizzards      []*Blizzard
}

func (m *Map) offset(p Pos) int {
	return p.y*m.width + p.x
}

func (m *Map) nextMinute() {
	occupiedByBlizzard := make(map[Pos]bool)
	for _, blizzard := range m.blizzards {
		if blizzard.direction.x != 0 || blizzard.direction.y != 0 {
			blizzard.pos = blizzard.pos.add(blizzard.direction)
			if blizzard.pos.x == 0 {
				blizzard.pos.x = m.width - 2
			}
			if blizzard.pos.x == m.width-1 {
				blizzard.pos.x = 1
			}
			if blizzard.pos.y == 0 {
				blizzard.pos.y = m.height - 2
			}
			if blizzard.pos.y == m.height-1 {
				blizzard.pos.y = 1
			}
		}
		occupiedByBlizzard[blizzard.pos] = true
	}
	emptyPositions := make(map[Pos]bool)
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if !occupiedByBlizzard[Pos{x, y}] {
				emptyPositions[Pos{x, y}] = true
			}
		}
	}
	m.emptyPositions = append(m.emptyPositions, emptyPositions)
}

func (m *Map) String() string {
	positions := make(map[Pos]string)
	for _, blizzard := range m.blizzards {
		p := positions[blizzard.pos]
		if p != "" {
			value, err := strconv.Atoi(p)
			if err != nil {
				value = 2
			} else {
				value++
			}
			positions[blizzard.pos] = strconv.Itoa(value)
		} else {
			switch blizzard.direction {
			case Pos{0, 0}:
				positions[blizzard.pos] = "#"
			case Pos{-1, 0}:
				positions[blizzard.pos] = "<"
			case Pos{1, 0}:
				positions[blizzard.pos] = ">"
			case Pos{0, -1}:
				positions[blizzard.pos] = "^"
			case Pos{0, 1}:
				positions[blizzard.pos] = "v"
			}
		}
	}
	result := ""
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			p := positions[Pos{x, y}]
			if p == "" {
				result += "."
			} else {
				result += p
			}
		}
		result += "\n"
	}
	return strings.TrimSpace(result)
}

// PathQueue implements a priority queue, see https://pkg.go.dev/container/heap
type PathQueue []*Path

func (q *PathQueue) Len() int           { return len(*q) }
func (q *PathQueue) Less(i, j int) bool { return (*q)[i].cost() < (*q)[j].cost() }
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

func isPossiblePosition(m *Map, p Pos, currentMinute int) bool {
	if p.x < 0 || p.x >= m.width || p.y < 0 || p.y >= m.height {
		return false
	}
	for len(m.emptyPositions) <= currentMinute {
		m.nextMinute()
	}
	return m.emptyPositions[currentMinute][p]
}

func possiblePositions(m *Map, p Pos, currentMinute int) []Pos {
	var positions []Pos
	candidates := []Pos{{p.x + 1, p.y}, {p.x, p.y + 1}, {p.x - 1, p.y}, {p.x, p.y - 1}, {p.x, p.y}}
	for _, candidate := range candidates {
		if isPossiblePosition(m, candidate, currentMinute) {
			positions = append(positions, candidate)
		}
	}
	return positions
}

func findPathQueue(m *Map, start, end Pos, minuteOffset int) int {
	startPath := &Path{
		positions: []Pos{start},
		tip:       start,
		end:       end,
	}

	visited := make(map[Pos]int)

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
			return path.nextMinute() - 1
		}
		visited[path.tip] = visited[path.tip] + 1
		if iteration%100000 == 0 {
			fmt.Printf("iteration: %v, paths: %v, tip: (%v, %v), cost: %v\n",
				iteration, queue.Len(), path.tip.x, path.tip.y, path.cost())
		}

		positions := possiblePositions(m, path.tip, path.nextMinute()+minuteOffset)

		for _, n := range positions {
			pathToNext := &Path{
				positions:  append(path.positions, n),
				tip:        n,
				end:        end,
				backTracks: visited[n],
			}

			heap.Push(queue, pathToNext)
		}
	}
	return -1
}

func main() {
	m := parseInput(loadInput("puzzle-input.txt"))
	minutesToGoal := findPathQueue(m, Pos{1, 0}, Pos{m.width - 2, m.height - 1}, 0)
	fmt.Printf("part one, minutes to goal: %v\n", minutesToGoal)
	minutesBack := findPathQueue(m, Pos{m.width - 2, m.height - 1}, Pos{1, 0}, minutesToGoal)
	minutesBackToGoal := findPathQueue(m, Pos{1, 0}, Pos{m.width - 2, m.height - 1}, minutesToGoal+minutesBack)
	fmt.Printf("part two, minutes to goal, back to start, and back to goal: %v\n", minutesToGoal+minutesBack+minutesBackToGoal)
}

func parseInput(input string) *Map {
	lines := strings.Split(input, "\n")
	m := &Map{
		width:  len(lines[0]),
		height: len(lines),
	}
	emptyPositions := make(map[Pos]bool)
	for y, line := range strings.Split(input, "\n") {
		for x := 0; x < len(line); x++ {
			p := Pos{x, y}
			switch line[x] {
			case '#':
				m.blizzards = append(m.blizzards, &Blizzard{p, Pos{0, 0}})
			case '^':
				m.blizzards = append(m.blizzards, &Blizzard{p, Pos{0, -1}})
			case 'v':
				m.blizzards = append(m.blizzards, &Blizzard{p, Pos{0, 1}})
			case '>':
				m.blizzards = append(m.blizzards, &Blizzard{p, Pos{1, 0}})
			case '<':
				m.blizzards = append(m.blizzards, &Blizzard{p, Pos{-1, 0}})
			case '.':
				emptyPositions[p] = true
			}
		}
	}
	m.emptyPositions = append(m.emptyPositions, emptyPositions)
	return m
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
