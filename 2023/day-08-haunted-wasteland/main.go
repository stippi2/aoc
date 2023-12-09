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

func partOne(d Directions, nodes map[string]*Node) int {
	node := nodes["AAA"]
	steps := 0
	for {
		direction := d.next()
		if direction == "R" {
			node = node.right
		} else {
			node = node.left
		}
		steps++
		if node.name == "ZZZ" {
			break
		}
	}
	return steps
}

func partTwo(d Directions, nodesMap map[string]*Node) int {
	var nodes []*Node
	for _, node := range nodesMap {
		if strings.HasSuffix(node.name, "A") {
			nodes = append(nodes, node)
		}
	}
	steps := 0
	for {
		direction := d.next()
		for i, node := range nodes {
			if direction == "R" {
				nodes[i] = node.right
			} else {
				nodes[i] = node.left
			}
		}
		steps++
		allNodesAtZ := true
		for _, node := range nodes {
			if !strings.HasSuffix(node.name, "Z") {
				allNodesAtZ = false
				break
			}
		}
		if allNodesAtZ {
			break
		}
	}
	return steps
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

func parseInput(input string) (Directions, map[string]*Node) {
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

	return Directions{directions: sections[0]}, nodes
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}