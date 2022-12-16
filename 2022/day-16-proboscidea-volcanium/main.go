package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
	"time"
)

type Node struct {
	label          string
	flowRate       int
	connectedNodes []*Node
}

type Path struct {
	elapsedTime      int
	pressureReleased int
	actions          []string
	valvesToOpen     []*Node
	previous         *Node
	tip              *Node
}

func (p *Path) openValue(node *Node, timeLimit int) {
	var valuesToOpen []*Node
	for _, n := range p.valvesToOpen {
		if n != node {
			valuesToOpen = append(valuesToOpen, n)
		}
	}
	p.elapsedTime++
	p.pressureReleased += node.flowRate * (timeLimit - p.elapsedTime)
	p.valvesToOpen = valuesToOpen
}

func (p *Path) canOpenValue(node *Node, timeLimit int) bool {
	if p.elapsedTime == timeLimit {
		return false
	}
	for _, n := range p.valvesToOpen {
		if n == node {
			return true
		}
	}
	return false
}

func (p *Path) potential() int {
	potential := p.pressureReleased
	for _, v := range p.valvesToOpen {
		potential += v.flowRate * (30 - p.elapsedTime)
	}
	return potential
}

func (p *Path) String() string {
	s := ""
	for _, action := range p.actions {
		s += "\n" + action
	}
	for _, valve := range p.valvesToOpen {
		s += "\nnot yet opened: " + valve.label
	}
	return s
}

func explore(path *Path, node *Node, openVale bool, timeLimit int) *Path {
	path = &Path{
		elapsedTime:      path.elapsedTime + 1,
		pressureReleased: path.pressureReleased,
		actions:          append([]string{}, path.actions...),
		valvesToOpen:     append([]*Node{}, path.valvesToOpen...),
		previous:         path.tip,
		tip:              node,
	}
	path.actions = append(path.actions, "visit "+node.label)
	if openVale && path.canOpenValue(node, timeLimit) {
		path.openValue(node, timeLimit)
		path.actions = append(path.actions, "open "+node.label)
	}
	return path
}

// PathQueue implements a priority queue, see https://pkg.go.dev/container/heap
type PathQueue []*Path

func (q *PathQueue) Len() int {
	return len(*q)
}

func (q *PathQueue) Less(i, j int) bool {
	//if (*q)[i].pressureReleased == (*q)[j].pressureReleased {
	//	return (*q)[i].elapsedTime < (*q)[j].elapsedTime
	//}
	//return (*q)[i].pressureReleased > (*q)[j].pressureReleased
	return (*q)[i].potential() > (*q)[j].potential()
}

func (q *PathQueue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

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

func maximumPressureRelease(startPath *Path, timeLimit int) int {
	queue := &PathQueue{startPath}
	heap.Init(queue)

	startTime := time.Now()
	iteration := 0

	//	elapsedTimeToMostPressureReleased := make(map[int]*Path)
	var bestPath *Path

	for queue.Len() > 0 {
		iteration++
		path := heap.Pop(queue).(*Path)
		if path.elapsedTime == timeLimit || len(path.valvesToOpen) == 0 {
			//fmt.Printf("found path with pressure release %v after %v / %v iterations, paths in map: %v\n",
			//	path.pressureReleased, time.Since(startTime), iteration, queue.Len())
			//fmt.Printf("path: %s\n", path)
			//return path.pressureReleased
			if bestPath == nil || path.pressureReleased > bestPath.pressureReleased {
				bestPath = path
			}
			continue
		}
		if iteration%100000 == 0 {
			fmt.Printf("iteration: %v, paths: %v, tip: (%v), pressure released: %v\n",
				iteration, queue.Len(), path.tip.label, path.pressureReleased)
		}

		var nextPaths []*Path
		for _, n := range path.tip.connectedNodes {
			if n == path.previous {
				continue
			}
			if path.canOpenValue(n, timeLimit-1) {
				nextPaths = append(nextPaths, explore(path, n, true, timeLimit))
			}
			nextPaths = append(nextPaths, explore(path, n, false, timeLimit))
		}

		// For each of the possible directions, create a new path that includes the node taken
		for _, p := range nextPaths {
			//pathAtTime := elapsedTimeToMostPressureReleased[p.elapsedTime]
			//if pathAtTime == nil {
			//	elapsedTimeToMostPressureReleased[p.elapsedTime] = p
			//	heap.Push(queue, p)
			//} else if p.pressureReleased > pathAtTime.pressureReleased {
			//	elapsedTimeToMostPressureReleased[p.elapsedTime] = p
			//	heap.Push(queue, p)
			//}
			heap.Push(queue, p)
		}
	}
	fmt.Printf("found path with pressure release %v after %v / %v iterations, paths in map: %v\n",
		bestPath.pressureReleased, time.Since(startTime), iteration, queue.Len())
	fmt.Printf("path: %s\n", bestPath)
	return bestPath.pressureReleased
}

func main() {
	start := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("highest achievable pressure release within 30 minutes: %v\n", maximumPressureRelease(start, 30))
}

func parseInput(input string) *Path {
	nodes := make(map[string]*Node)
	connections := make(map[string]string)
	var valvesToOpen []*Node
	// Create all nodes, remember the connections in a string -> string map
	for _, line := range strings.Split(input, "\n") {
		node := &Node{}
		parts := strings.Split(line, "; ")
		matches, err := fmt.Sscanf(parts[0], "Valve %s has flow rate=%d", &node.label, &node.flowRate)
		if matches != 2 || err != nil {
			panic(fmt.Sprintf("failed to parse valve line '%s': %v", line, err))
		}
		// Trim all lowercase characters and whitespace
		tunnels := strings.Trim(parts[1], " abcdefghijklmnopqrstuvwxyz")
		connections[node.label] = tunnels
		nodes[node.label] = node
		if node.flowRate > 0 {
			valvesToOpen = append(valvesToOpen, node)
		}
	}
	// Establish the connections
	for label, connectedLabels := range connections {
		node := nodes[label]
		for _, connectedLabel := range strings.Split(connectedLabels, ", ") {
			node.connectedNodes = append(node.connectedNodes, nodes[connectedLabel])
		}
	}
	return &Path{
		tip:          nodes["AA"],
		valvesToOpen: valvesToOpen,
	}
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
