package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
	"time"
)

type Map struct {
	width  int
	height int
	data   []uint8
}

func (m *Map) isEqual(other *Map) bool {
	if m.width != other.width {
		return false
	}
	if m.height != other.height {
		return false
	}
	for i := 0; i < len(m.data); i++ {
		if m.data[i] != other.data[i] {
			return false
		}
	}
	return true
}

func (m *Map) init(width, height int) {
	m.width = width
	m.height = height
	m.data = make([]uint8, width*height)
}

func (m *Map) offset(x, y int) int {
	return m.width*y + x
}

func (m *Map) set(x, y int, value uint8) {
	m.data[m.offset(x, y)] = value
}

func (m *Map) get(x, y int) uint8 {
	return m.data[m.offset(x, y)]
}

func (m *Map) String() string {
	result := ""
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			result += fmt.Sprintf("%s", string(m.get(x, y)))
		}
		result = strings.TrimRight(result, " ") + "\n"
	}
	return result[0:len(result)-1]
}

type Amphipod struct {
	x     int
	y     int
	kind  uint8
	moved bool
}

type Move struct {
	x     int
	y     int
	steps int
}

func (a Amphipod) String() string {
	return fmt.Sprintf("%s(%d, %d)", string(a.kind), a.x, a.y)
}

func (a Amphipod) energyPerStep() int {
	switch a.kind {
	case 'A': return 1
	case 'B': return 10
	case 'C': return 100
	case 'D': return 1000
	}
	return 0
}

func (a Amphipod) targetX() int {
	switch a.kind {
	case 'A': return 3
	case 'B': return 5
	case 'C': return 7
	case 'D': return 9
	}
	return 0
}

func (a Amphipod) possibleMoves(m *Map) (moves []Move) {
	// Moving from hallway into room
	room := a.targetX()
	if a.x == room && a.moved {
		return nil
	}
	if a.y == 1 {
		// Check if it's possible to move into target room
		y := m.height - 2
		for y > 1 {
			v := m.get(room, y)
			if v == '.' {
				break
			} else if v == a.kind {
				y--
			} else {
				return
			}
		}
		if y == 1 {
			return
		}
		// Check if any obstacle is along the way in the hallway
		if room > a.x {
			for x := a.x+1; x <= room; x++ {
				if m.get(x, 1) != '.' {
					return
				}
			}
		} else {
			for x := a.x-1; x >= room; x-- {
				if m.get(x, 1) != '.' {
					return
				}
			}
		}
		moves = append(moves, Move{
			x:     room,
			y:     y,
			steps: y - a.y + abs(room - a.x),
		})
	} else {
		// Moving from room into hallway
		for y := a.y-1; y >= 2; y-- {
			if m.get(a.x, y) != '.' {
				// another Amphipod is in the way
				return
			}
		}
		for x := a.x+1; x < m.width - 1; x++ {
			if m.get(x, 1) != '.' {
				break
			}
			if x == 5 || x == 7 || x == 9 {
				continue
			}
			moves = append(moves, Move{
				x:     x,
				y:     1,
				steps: a.y - 1 + x - a.x,
			})
		}
		for x := a.x-1; x > 0; x-- {
			if m.get(x, 1) != '.' {
				break
			}
			if x == 3 || x == 5 || x == 7 {
				continue
			}
			moves = append(moves, Move{
				x:     x,
				y:     1,
				steps: a.y - 1 + a.x - x,
			})
		}
	}
	return
}

type Solution struct {
	mapTemplate	*Map
	pods        []Amphipod
	energyLevel int
	previous    *Solution
}

var emptyMap = parseInput(`#############
#...........#
###.#.#.#.###
  #.#.#.#.#
  #########`)

var emptyMapPart2 = parseInput(`#############
#...........#
###.#.#.#.###
  #.#.#.#.#
  #.#.#.#.#
  #.#.#.#.#
  #########`)

func (s *Solution) generateMap() *Map {
	m := &Map{
		width: s.mapTemplate.width,
		height: s.mapTemplate.height,
		data: make([]uint8, len(s.mapTemplate.data)),
	}
	copy(m.data, s.mapTemplate.data)
	for _, pod := range s.pods {
		m.set(pod.x, pod.y, pod.kind)
	}
	return m
}

func (s *Solution) isFinal() bool {
	for _, pod := range s.pods {
		if pod.x != pod.targetX() {
			return false
		}
	}
	return true
}

func (s *Solution) apply(pod Amphipod, move Move) *Solution {
	result := &Solution{
		mapTemplate: s.mapTemplate,
		pods:        make([]Amphipod, len(s.pods)),
		energyLevel: s.energyLevel,
		previous:    s,
	}
	result.energyLevel += pod.energyPerStep() * move.steps
	for i, p := range s.pods {
		if p == pod {
			result.pods[i] = Amphipod{
				x:     move.x,
				y:     move.y,
				kind:  pod.kind,
				moved: true,
			}
		} else {
			result.pods[i] = p
		}
	}
	sortPods(result.pods)
	return result
}

func (s *Solution) String() string {
	result := ""
	for _, p := range s.pods {
		result += p.String()
	}
	return result
}

func sortPods(a []Amphipod) {
	sort.Slice(a, func (i, j int) bool {
		if a[i].x < a[j].x {
			return true
		}
		if a[i].x == a[j].x {
			if a[i].y < a[j].y {
				return true
			}
			if a[i].kind == a[j].kind {
				if a[i].kind < a[j].kind {
					return true
				}
				return false
			}
			return false
		}
		return false
	})
}

// SolutionQueue implements a priority queue, see https://pkg.go.dev/container/heap
type SolutionQueue []*Solution

func (q SolutionQueue) Len() int           { return len(q) }
func (q SolutionQueue) Less(i, j int) bool { return q[i].energyLevel < q[j].energyLevel }
func (q SolutionQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *SolutionQueue) Push(x interface{}) {
	path := x.(*Solution)
	*q = append(*q, path)
}

func (q *SolutionQueue) Pop() interface{} {
	old := *q
	n := len(old)
	path := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*q = old[0 : n-1]
	return path
}

func solve(m, mapTemplate *Map) int {
	solutions := &SolutionQueue{
		{
			mapTemplate: mapTemplate,
			pods:        getAmphipods(m),
			energyLevel: 0,
		},
	}
	heap.Init(solutions)
	// A mapping between a certain configuration of pod locations to the currently lowest energy level to achieve it
	configurations := make(map[string]int)
	iteration := 0
	rejected := 0
	for solutions.Len() > 0 {
		iteration++
		solution := heap.Pop(solutions).(*Solution)
		if iteration % 100000 == 0 {
			fmt.Printf("solutions: %d, rejected: %d, level of first: %d, map:\n%s\n",
				solutions.Len(), rejected, solution.energyLevel, solution.generateMap().String())
		}

		bestLevel, hasKey := configurations[solution.String()]
		if hasKey && bestLevel < solution.energyLevel {
			// No need to continue with this solution,
			// as there is another one getting to this configuration with less energy
			rejected++
			continue
		}

		currentMap := solution.generateMap()
		//fmt.Printf("  map: %s\n", currentMap)
		hasMoves := false
		for _, pod := range solution.pods {
			moves := pod.possibleMoves(currentMap)
			if len(moves) > 0 {
				hasMoves = true
				//fmt.Printf("pod: %s has moves: %d\n", &pod, len(moves))
				for _, move := range moves {
					derived := solution.apply(pod, move)
					currentBestLevel, contained := configurations[derived.String()]
					if !contained || currentBestLevel > derived.energyLevel {
						configurations[derived.String()] = derived.energyLevel
						heap.Push(solutions, derived)
					} else {
						rejected++
					}
				}
			}
		}
		if !hasMoves && solution.isFinal() {
			//previous := solution
			//for previous != nil {
			//	fmt.Printf("--- %d:\n%s\n", previous.energyLevel, previous.generateMap().String())
			//	previous = previous.previous
			//}
			return solution.energyLevel
		}
	}
	return math.MaxInt64
}

func getAmphipods(m *Map) []Amphipod {
	var amphipods []Amphipod
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			v := m.get(x, y)
			switch v {
			case 'A', 'B', 'C', 'D':
				amphipods = append(amphipods, Amphipod{
					x:    x,
					y:    y,
					kind: v,
				})
			}
		}
	}
	return amphipods
}

func main() {
	m := parseInput(loadInput("puzzle-input.txt"))
	start := time.Now()
	energy := solve(m, emptyMapPart2)
	fmt.Printf("lowest energy to move: %d (%s)\n", energy, time.Since(start))
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func parseInput(input string) *Map {
	input = strings.TrimSpace(input)
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
		x := 0
		for ; x < len(lines[y]); x++ {
			if x >= width {
				break
			}
			m.set(x, y, lines[y][x])
		}
		for ; x < width; x++ {
			m.set(x, y, ' ')
		}
	}
	return &m
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
