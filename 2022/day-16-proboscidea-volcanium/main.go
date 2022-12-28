package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

type Node struct {
	label          string
	flowRate       int
	connectedNodes []*Node
	distance       map[string]int
}

type NodePath struct {
	visited map[*Node]bool
	tip     *Node
}

func findDistance(fromNode *Node, toNode *Node) int {
	var pathQueue []*NodePath
	for _, node := range fromNode.connectedNodes {
		pathQueue = append(pathQueue, &NodePath{tip: node, visited: map[*Node]bool{node: true}})
	}
	distance := math.MaxInt
	for len(pathQueue) > 0 {
		path := pathQueue[0]
		pathQueue = pathQueue[1:]
		if len(path.visited) >= distance {
			continue
		}
		if path.tip == toNode {
			if len(path.visited) < distance {
				distance = len(path.visited)
			}
			continue
		}
		for _, node := range path.tip.connectedNodes {
			if !path.visited[node] {
				nextPath := &NodePath{tip: node, visited: map[*Node]bool{node: true}}
				for visited := range path.visited {
					nextPath.visited[visited] = true
				}
				pathQueue = append(pathQueue, nextPath)
			}
		}
	}
	return distance
}

type Path struct {
	timeRemaining    int
	pressureReleased int
	actions          []string
	valvesToOpen     []*Node
	previous         *Node
	tip              *Node
}

func (p *Path) canOpenValue(node *Node) bool {
	if p.timeRemaining == 0 {
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
		potential += v.flowRate * (p.timeRemaining - p.tip.distance[v.label] - 1)
	}
	return potential
}

func (p *Path) String() string {
	s := ""
	for _, action := range p.actions {
		s += "\n" + action
	}
	if len(p.valvesToOpen) > 0 {
		s += "\nnot yet opened: "
		for i, valve := range p.valvesToOpen {
			if i > 0 {
				s += ", "
			}
			s += valve.label
		}
	}
	return s
}

func (p *Path) clone() *Path {
	return &Path{
		timeRemaining:    p.timeRemaining,
		pressureReleased: p.pressureReleased,
		actions:          append([]string{}, p.actions...),
		valvesToOpen:     append([]*Node{}, p.valvesToOpen...),
		previous:         p.previous,
		tip:              p.tip,
	}
}

func explore(path *Path, node *Node) *Path {
	newPath := path.clone()
	newPath.timeRemaining--
	newPath.previous = path.tip // Prevent going back to this node immediately
	newPath.tip = node
	newPath.actions = append(newPath.actions, "visit "+node.label)
	return newPath
}

func openValve(path *Path, node *Node) *Path {
	if path.tip != node {
		panic(fmt.Sprintf("cannot open node %s, path: %s", node.label, path))
	}
	var valuesToOpen []*Node
	for _, n := range path.valvesToOpen {
		if n != node {
			valuesToOpen = append(valuesToOpen, n)
		}
	}
	newPath := path.clone()
	newPath.timeRemaining--
	newPath.valvesToOpen = valuesToOpen
	newPath.previous = nil // It's ok to go back to the actual previous node if we just opened the tip
	newPath.pressureReleased += node.flowRate * newPath.timeRemaining
	newPath.actions = append(newPath.actions, "open "+node.label)
	return newPath
}

// PathQueue implements a priority queue, see https://pkg.go.dev/container/heap
type PathQueue []*Path

func (q *PathQueue) Len() int {
	return len(*q)
}

func (q *PathQueue) Less(i, j int) bool {
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
	startPath.timeRemaining = timeLimit
	queue := &PathQueue{startPath}
	heap.Init(queue)

	startTime := time.Now()
	iteration := 0

	for queue.Len() > 0 {
		iteration++
		path := heap.Pop(queue).(*Path)
		if path.timeRemaining == 0 || len(path.valvesToOpen) == 0 {
			fmt.Printf("found path with pressure release %v after %v / %v iterations, paths in map: %v\n",
				path.pressureReleased, time.Since(startTime), iteration, queue.Len())
			fmt.Printf("path: %s\n", path)
			return path.pressureReleased
		}
		if iteration%1000 == 0 {
			fmt.Printf("iteration: %v, paths: %v, tip: (%v), pressure released: %v, potential: %v, elapsed minutes: %v\n",
				iteration, queue.Len(), path.tip.label, path.pressureReleased, path.potential(), timeLimit-path.timeRemaining)
		}

		for _, n := range path.tip.connectedNodes {
			if n == path.previous {
				// No immediate backtracking
				continue
			}
			pathToNode := explore(path, n)
			heap.Push(queue, pathToNode)
			if pathToNode.canOpenValue(n) {
				heap.Push(queue, openValve(pathToNode, n))
			}
		}
	}
	return 0
}

func sortNodes(valves []*Node, tip *Node) {
	sort.Slice(valves, func(i, j int) bool {
		return (valves[i].flowRate - tip.distance[valves[i].label]) > (valves[j].flowRate - tip.distance[valves[j].label])
	})
}

func maximumPressureReleaseWithElephant(startPath *Path, timeLimit int) int {
	var myValves []*Node
	myTip := startPath.tip
	var elephantValves []*Node
	elephantTip := startPath.tip
	valves := append([]*Node{}, startPath.valvesToOpen...)
	index := 0
	for len(valves) > 0 {
		if index%2 == 0 {
			sortNodes(valves, myTip)
			myValves = append(myValves, valves[0])
			myTip = valves[0]
		} else {
			sortNodes(valves, elephantTip)
			elephantValves = append(elephantValves, valves[0])
			elephantTip = valves[0]
		}
		index++
		valves = valves[1:]
	}
	myPath := startPath.clone()
	myPath.valvesToOpen = myValves

	elephantPath := startPath.clone()
	elephantPath.valvesToOpen = elephantValves

	return maximumPressureRelease(myPath, timeLimit) + maximumPressureRelease(elephantPath, timeLimit)
}

func main() {
	startTime := time.Now()
	start := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("building node tree: %v\n", time.Since(startTime))
	fmt.Printf("highest achievable pressure release within 30 minutes: %v\n", maximumPressureRelease(start, 30))

	startTime = time.Now()
	fmt.Printf("highest achievable pressure release within 26 minutes with elephant: %v\n", maximumPressureReleaseWithElephant(start, 26))
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
	// Establish the distances
	for _, fromNode := range nodes {
		fromNode.distance = map[string]int{}
		for _, toNode := range nodes {
			if fromNode == toNode {
				continue
			}
			fromNode.distance[toNode.label] = findDistance(fromNode, toNode)
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
