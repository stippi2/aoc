package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Node struct {
	name  string
	left  *Node
	right *Node
}

type Directions struct {
	directions string
	pos        int
}

func (d *Directions) next() string {
	if d.pos >= len(d.directions) {
		d.pos = 0
	}
	next := d.directions[d.pos]
	d.pos++
	return string(next)
}

func (d *Directions) reset() {
	d.pos = 0
}

func followDirections(d *Directions, node *Node, stopCondition func(node *Node, steps int) bool) int {
	steps := 0
	for {
		direction := d.next()
		if direction == "R" {
			node = node.right
		} else {
			node = node.left
		}
		steps++
		if stopCondition(node, steps) {
			break
		}
	}
	return steps
}

func partOne(d *Directions, nodes map[string]*Node) int {
	return followDirections(d, nodes["AAA"], func(node *Node, _ int) bool {
		return node.name == "ZZZ"
	})
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func findLCM(numbers []int) int {
	result := numbers[0]
	for _, num := range numbers[1:] {
		result = lcm(result, num)
	}
	return result
}

type StopNode struct {
	name            string
	posInDirections int
	steps           int
}

type NodePattern struct {
	startNode *Node
	stopNodes []StopNode
	steps     int
}

func (np *NodePattern) findStopNode(n StopNode) *StopNode {
	for _, sn := range np.stopNodes {
		if sn.name == n.name && sn.posInDirections == n.posInDirections {
			return &sn
		}
	}
	return nil
}

func partTwo(d *Directions, nodesMap map[string]*Node) int {

	var startNodes []*NodePattern
	for _, node := range nodesMap {
		if strings.HasSuffix(node.name, "A") {
			startNodes = append(startNodes, &NodePattern{startNode: node})
		}
	}
	for _, np := range startNodes {
		d.reset()
		followDirections(d, np.startNode, func(node *Node, steps int) bool {
			if strings.HasSuffix(node.name, "Z") {
				stopNode := StopNode{name: node.name, posInDirections: d.pos, steps: steps}
				earlierStopNode := np.findStopNode(stopNode)
				if earlierStopNode == nil {
					np.stopNodes = append(np.stopNodes, stopNode)
				} else {
					np.steps = earlierStopNode.steps
					return true
				}
			}
			return false
		})
	}

	var steps []int
	for _, np := range startNodes {
		fmt.Printf("%s: steps: %d, stopNodes: %d\n", np.startNode.name, np.steps, len(np.stopNodes))
		steps = append(steps, np.steps)
	}

	return findLCM(steps)
}

func main() {
	now := time.Now()
	directions, nodes := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(directions, nodes)
	part2 := partTwo(directions, nodes)
	duration := time.Since(now)
	fmt.Printf("Part 1: Steps required to reach ZZZ; %d\n", part1)
	fmt.Printf("Part 2: Steps required to simultaneously reach **Z: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseInput(input string) (*Directions, map[string]*Node) {
	sections := strings.Split(input, "\n\n")
	nodes := make(map[string]*Node)

	for _, line := range strings.Split(sections[1], "\n") {
		parts := strings.Split(line, " = ")
		name := parts[0]
		nodes[name] = &Node{name: name}
	}

	for _, line := range strings.Split(sections[1], "\n") {
		parts := strings.Split(line, " = ")
		node := nodes[parts[0]]
		connections := strings.Split(strings.Trim(parts[1], "()"), ", ")
		node.left = nodes[connections[0]]
		node.right = nodes[connections[1]]
	}

	return &Directions{directions: sections[0]}, nodes
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
