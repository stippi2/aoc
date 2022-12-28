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

func max(values ...int) int {
	maxValue := math.MinInt
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

type Blueprint struct {
	costs  map[string]map[string]int
	useful map[string]int
}

func (b *Blueprint) usefulness(kind string) int {
	if b.useful == nil {
		b.useful = map[string]int{
			"ore":      max(b.costs["clay"]["ore"], b.costs["obsidian"]["ore"], b.costs["geode"]["ore"]),
			"clay":     b.costs["obsidian"]["clay"],
			"obsidian": b.costs["geode"]["obsidian"],
			"geode":    math.MaxInt,
		}
	}
	return b.useful[kind]
}

type State struct {
	minerals    map[string]int
	robots      map[string]int
	ignore      map[string]bool
	usedMinutes int
}

func (s *State) canBuildRobot(blueprint Blueprint, kind string) bool {
	for mineral, howMuch := range blueprint.costs[kind] {
		if s.minerals[mineral] < howMuch {
			return false
		}
	}
	return true
}

func (s *State) buildRobot(blueprint Blueprint, robot string) {
	for mineral, howMuch := range blueprint.costs[robot] {
		s.minerals[mineral] = s.minerals[mineral] - howMuch
	}
	s.robots[robot] = s.robots[robot] + 1
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
		ignore:      map[string]bool{},
		usedMinutes: s.usedMinutes,
	}
	for robotKind, howMany := range s.robots {
		result.robots[robotKind] = howMany
	}
	for mineral, howMuch := range s.minerals {
		result.minerals[mineral] = howMuch
	}
	for robotKind := range s.ignore {
		result.ignore[robotKind] = true
	}
	return result
}

func (s *State) effectivity() int {
	e := s.minerals["geode"] * 10000
	e += s.robots["geode"] * 1000
	e += s.robots["obsidian"] * 100
	e += s.robots["clay"] * 10
	e += s.robots["ore"]
	e *= 25 - s.usedMinutes
	return e
}

func (s *State) possibleOutput(timeLimit int) int {
	output := s.minerals["geode"]
	for i := timeLimit - s.usedMinutes; i > 0; i-- {
		output += s.robots["geode"] + i
	}
	return output
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
	var states []*State

	var options []string
	if state.canBuildRobot(blueprint, "geode") {
		options = append(options, "geode")
	} else {
		for _, kind := range []string{"obsidian", "clay", "ore"} {
			if !state.ignore[kind] && state.robots[kind] < blueprint.usefulness(kind) && state.canBuildRobot(blueprint, kind) {
				options = append(options, kind)
			}
		}
	}

	state.produceMinerals()

	// Do nothing option
	for _, kind := range options {
		state.ignore[kind] = true
	}
	states = append(states, state)

	for _, kind := range options {
		clone := state.clone()
		//fmt.Printf("building %s robot after minute %v\n", kind, clone.usedMinutes)
		clone.buildRobot(blueprint, kind)
		clone.ignore = map[string]bool{}
		states = append(states, clone)
	}

	return states
}

func findQualityLevel(blueprint Blueprint, timeLimit int) int {
	start := &State{
		minerals: map[string]int{},
		robots:   map[string]int{"ore": 1},
		ignore:   map[string]bool{},
	}

	queue := &StrategyQueue{start}
	heap.Init(queue)

	startTime := time.Now()
	iteration := 0

	maxGeodes := 0

	bestPerMinute := map[int]int{}

	for queue.Len() > 0 {
		iteration++
		state := heap.Pop(queue).(*State)
		if state.usedMinutes == timeLimit {
			//fmt.Printf("found best strategy after %v / %v iterations, remaining in queue: %v\n", time.Since(startTime),
			//	iteration, queue.Len())
			//return state.minerals["geode"]
			maxGeodes = max(maxGeodes, state.minerals["geode"])
			continue
		}

		if state.possibleOutput(timeLimit) < maxGeodes {
			continue
		}
		currentBest := bestPerMinute[state.usedMinutes]
		current := state.minerals["geode"]
		if current < currentBest {
			continue
		} else if current > currentBest {
			bestPerMinute[state.usedMinutes] = current
		}

		if iteration%100000 == 0 {
			fmt.Printf("iteration: %v, paths: %v, geodes: (%v), current best: %v\n",
				iteration, queue.Len(), state.minerals["geode"], maxGeodes)
		}

		for _, nextState := range possibleStates(state, blueprint) {
			heap.Push(queue, nextState)
		}
	}
	fmt.Printf("found best strategy after %v and %v iterations: %v\n", time.Since(startTime), iteration, maxGeodes)
	return maxGeodes
}

func main() {
	blueprints := parseInput(loadInput("puzzle-input.txt"))
	sum := 0
	for i, blueprint := range blueprints {
		qualityLevel := findQualityLevel(blueprint, 24)
		fmt.Printf("geodes of best strategy for blueprint %v: %v\n", i, qualityLevel)
		sum += (i + 1) * qualityLevel
	}
	fmt.Printf("sum of all quality levels: %v\n", sum)

	product := 1
	for i := 0; i < 3; i++ {
		qualityLevel := findQualityLevel(blueprints[i], 32)
		product *= qualityLevel
	}
	fmt.Printf("product of first 3 blueprints: %v\n", product)
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
