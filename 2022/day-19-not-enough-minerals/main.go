package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Blueprint struct {
	costs map[string]map[string]int
}

type State struct {
	minerals    map[string]int
	robots      map[string]int
	ignore      map[string]int
	usedMinutes int
}

func (s *State) canBuildRobot(blueprint Blueprint, kind string) int {
	max := math.MaxInt
	for mineral, howMuch := range blueprint.costs[kind] {
		max = min(s.minerals[mineral]/howMuch, max)
	}
	return max
}

func (s *State) buildRobot(blueprint Blueprint, robot string) *State {
	for mineral, howMuch := range blueprint.costs[robot] {
		s.minerals[mineral] = s.minerals[mineral] - howMuch
	}
	s.robots[robot] = s.robots[robot] + 1
	return s
}

func (s *State) produceMinerals() {
	for robotKind, count := range s.robots {
		s.minerals[robotKind] = s.minerals[robotKind] + count
	}
	s.usedMinutes++
}

func (s *State) clone() *State {
	result := &State{
		minerals:    map[string]int{},
		robots:      map[string]int{},
		ignore:      map[string]int{},
		usedMinutes: s.usedMinutes,
	}
	for robotKind, howMany := range s.robots {
		result.robots[robotKind] = howMany
	}
	for mineral, howMuch := range s.minerals {
		result.minerals[mineral] = howMuch
	}
	for robotKind, howMuch := range s.ignore {
		result.ignore[robotKind] = howMuch
	}
	return result
}

func (s *State) effectivity() int {
	robots := 0
	for _, count := range s.robots {
		robots += count
	}
	return s.minerals["geode"]*1000 + robots
}

// StrategyQueue implements a priority queue, see https://pkg.go.dev/container/heap
type StrategyQueue []*State

func (q *StrategyQueue) Len() int           { return len(*q) }
func (q *StrategyQueue) Less(i, j int) bool { return (*q)[i].effectivity() > (*q)[j].effectivity() }
func (q *StrategyQueue) Swap(i, j int)      { (*q)[i], (*q)[j] = (*q)[j], (*q)[i] }

func (q *StrategyQueue) Push(x interface{}) {
	path := x.(*State)
	*q = append(*q, path)
}

func (q *StrategyQueue) Pop() interface{} {
	old := *q
	n := len(old)
	path := old[n-1]
	old[n-1] = nil // avoid memory leak
	*q = old[0 : n-1]
	return path
}

func possibleStates(state *State, blueprint Blueprint) []*State {
	state.produceMinerals()
	var states []*State
	// TODO: Should be sorted by effectiveness of the given Blueprint for the respective mineral
	robotKinds := []string{"geode", "obsidian", "clay", "ore"}
	for _, kind := range robotKinds {
		capacity := state.canBuildRobot(blueprint, kind)
		if capacity > 0 {
			if state.ignore[kind] > 0 {
				delete(state.ignore, kind)
			} else {
				clone := state.clone()
				clone.ignore[kind] = capacity
				fmt.Printf("building %s robot (%v min)\n", kind, state.usedMinutes)
				for capacity > 0 {
					clone.buildRobot(blueprint, kind)
					capacity--
				}
				states = append(states, clone)
			}
		}
	}
	// Do nothing option
	states = append(states, state)
	return states
}

func findQualityLevel(blueprint Blueprint, timeLimit int) int {
	start := &State{
		minerals: map[string]int{},
		robots:   map[string]int{"ore": 1},
		ignore:   map[string]int{},
	}

	queue := &StrategyQueue{start}
	heap.Init(queue)

	startTime := time.Now()
	iteration := 0

	for queue.Len() > 0 {
		iteration++
		state := heap.Pop(queue).(*State)
		if state.usedMinutes == timeLimit {
			fmt.Printf("found best strategy after %v / %v iterations, remaining in queue: %v\n", time.Since(startTime),
				iteration, queue.Len())
			return state.minerals["geode"]
		}

		if iteration%100000 == 0 {
			fmt.Printf("iteration: %v, paths: %v, geodes: (%v), effectivity: %v\n",
				iteration, queue.Len(), state.minerals["geode"], state.effectivity())
		}

		for _, nextState := range possibleStates(state, blueprint) {
			heap.Push(queue, nextState)
		}
	}
	return -1
}

func main() {
	blueprints := parseInput(loadInput("puzzle-input.txt"))
	for i, blueprint := range blueprints {
		fmt.Printf("geodes of best strategy for blueprint %v: %v\n", i, findQualityLevel(blueprint, 24))
	}
}

func parseInput(input string) []Blueprint {
	lines := strings.Split(input, "\n")
	blueprints := make([]Blueprint, len(lines))
	minerals := []string{"ore", "clay", "obsidian", "geode"}
	for i, line := range lines {
		line = strings.TrimPrefix(line, fmt.Sprintf("Blueprint %d: ", i+1))
		line = strings.TrimSuffix(line, ".")
		parts := strings.Split(line, ". ")
		blueprints[i].costs = make(map[string]map[string]int)
		for j, robotKind := range minerals {
			costsString := strings.TrimPrefix(parts[j], fmt.Sprintf("Each %s robot costs ", robotKind))
			var costs1 int
			var costs2 int
			costs := make(map[string]int)
			switch robotKind {
			case "ore", "clay":
				_, _ = fmt.Sscanf(costsString, "%d ore", &costs1)
				costs["ore"] = costs1
			case "obsidian":
				_, _ = fmt.Sscanf(costsString, "%d ore and %d clay", &costs1, &costs2)
				costs["ore"] = costs1
				costs["clay"] = costs2
			case "geode":
				_, _ = fmt.Sscanf(costsString, "%d ore and %d obsidian", &costs1, &costs2)
				costs["ore"] = costs1
				costs["obsidian"] = costs2
			}
			blueprints[i].costs[robotKind] = costs
		}
	}
	return blueprints
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
