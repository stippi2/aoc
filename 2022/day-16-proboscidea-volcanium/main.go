package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type Node struct {
	label          string
	flowRate       int
	connectedNodes []*Node
	distance       map[string]int
}

func removeNode(nodes []*Node, node *Node) []*Node {
	var newNodes []*Node
	for _, n := range nodes {
		if n != node {
			newNodes = append(newNodes, n)
		}
	}
	return newNodes
}

type NodePath struct {
	visited map[*Node]bool
	tip     *Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findDistance(fromNode, toNode *Node) int {
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
	timeRemaining    []int
	pressureReleased int
	actions          []string
	valvesToOpen     []*Node
	previous         []*Node
	tip              []*Node
}

func (p *Path) canOpenValue(node *Node, tipIndex int) bool {
	if p.timeRemaining[tipIndex] == 0 {
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
		maxPotential := 0
		for i, tip := range p.tip {
			maxPotential = max(maxPotential, p.timeRemaining[i]-tip.distance[v.label]-1)
		}
		potential += v.flowRate * maxPotential
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
		timeRemaining:    append([]int{}, p.timeRemaining...),
		pressureReleased: p.pressureReleased,
		actions:          append([]string{}, p.actions...),
		valvesToOpen:     append([]*Node{}, p.valvesToOpen...),
		previous:         append([]*Node{}, p.previous...),
		tip:              append([]*Node{}, p.tip...),
	}
}

func explore(path *Path, tipIndex int, node *Node) *Path {
	newPath := path.clone()
	newPath.timeRemaining[tipIndex]--
	newPath.previous[tipIndex] = path.tip[tipIndex] // Prevent going back to this node immediately
	newPath.tip[tipIndex] = node
	newPath.actions = append(newPath.actions, fmt.Sprintf("[%v] visit %s", tipIndex, node.label))
	return newPath
}

func travelTo(path *Path, tipIndex int, node *Node) *Path {
	newPath := path.clone()
	newPath.timeRemaining[tipIndex] -= path.tip[tipIndex].distance[node.label]
	newPath.previous[tipIndex] = path.tip[tipIndex] // Prevent going back to this node immediately
	newPath.tip[tipIndex] = node
	newPath.actions = append(newPath.actions, fmt.Sprintf("[%v] travel to %s", tipIndex, node.label))
	return newPath
}

func openValve(path *Path, tipIndex int, node *Node) *Path {
	if path.tip[tipIndex] != node {
		panic(fmt.Sprintf("cannot open node %s, path: %s", node.label, path))
	}
	newPath := path.clone()
	newPath.timeRemaining[tipIndex]--
	newPath.valvesToOpen = removeNode(newPath.valvesToOpen, node)
	newPath.previous[tipIndex] = nil // It's ok to go back to the actual previous node if we just opened the tip
	newPath.pressureReleased += node.flowRate * newPath.timeRemaining[tipIndex]
	newPath.actions = append(newPath.actions, fmt.Sprintf("[%v] open %s", tipIndex, node.label))
	return newPath
}

func (p *Path) timeLeft() int {
	sum := 0
	for i := 0; i < len(p.timeRemaining); i++ {
		sum += p.timeRemaining[i]
	}
	return sum
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

func maximumPressureRelease(startPath *Path, timeLimit, workers int) int {
	for i := 0; i < workers; i++ {
		startPath.timeRemaining = append(startPath.timeRemaining, timeLimit)
		startPath.previous = append(startPath.previous, nil)
	}
	for i := 1; i < workers; i++ {
		startPath.tip = append(startPath.tip, startPath.tip[0])
	}
	queue := &PathQueue{startPath}
	heap.Init(queue)

	startTime := time.Now()
	iteration := 0

	for queue.Len() > 0 {
		iteration++
		path := heap.Pop(queue).(*Path)
		if path.timeLeft() == 0 || len(path.valvesToOpen) == 0 {
			fmt.Printf("found path with pressure release %v after %v / %v iterations, paths in map: %v\n",
				path.pressureReleased, time.Since(startTime), iteration, queue.Len())
			fmt.Printf("path: %s\n", path)
			return path.pressureReleased
		}
		if iteration%1000 == 0 {
			fmt.Printf("iteration: %v, paths: %v, pressure released: %v, potential: %v, elapsed minutes: %v\n",
				iteration, queue.Len(), path.pressureReleased, path.potential(), timeLimit*workers-path.timeLeft())
		}

		if workers == 1 {
			for _, n := range path.tip[0].connectedNodes {
				if n == path.previous[0] {
					// No immediate backtracking
					continue
				}
				pathToNode := explore(path, 0, n)
				heap.Push(queue, pathToNode)
				if pathToNode.canOpenValue(n, 0) {
					heap.Push(queue, openValve(pathToNode, 0, n))
				}
			}
		} else if workers == 2 {
			for _, n1 := range path.tip[0].connectedNodes {
				if n1 == path.previous[0] {
					continue
				}
				nextPath := explore(path, 0, n1)
				for _, n2 := range path.tip[1].connectedNodes {
					if n2 == path.previous[1] {
						continue
					}
					heap.Push(queue, explore(nextPath, 1, n2))
				}
				if nextPath.canOpenValue(nextPath.tip[1], 1) {
					heap.Push(queue, openValve(nextPath, 1, nextPath.tip[1]))
				}
			}
			if path.canOpenValue(path.tip[0], 0) {
				nextPath := openValve(path, 0, path.tip[0])
				for _, n2 := range path.tip[1].connectedNodes {
					if n2 == path.previous[1] {
						continue
					}
					heap.Push(queue, explore(nextPath, 1, n2))
				}
				if nextPath.canOpenValue(nextPath.tip[1], 1) {
					heap.Push(queue, openValve(nextPath, 1, nextPath.tip[1]))
				}
			}
		}
	}
	return 0
}

func main() {
	startTime := time.Now()
	start := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("building node tree: %v\n", time.Since(startTime))
	fmt.Printf("highest achievable pressure release within 30 minutes: %v\n", maximumPressureRelease(start, 30, 1))

	startTime = time.Now()
	fmt.Printf("highest achievable pressure release within 26 minutes with elephant: %v\n", maximumPressureRelease(start, 26, 2))
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
		tip:          []*Node{nodes["AA"]},
		valvesToOpen: valvesToOpen,
	}
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
