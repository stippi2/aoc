package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Node struct {
	name string
	next []*Node
}

func findEnd(n *Node, path string) int {
	if n.name == strings.ToLower(n.name) && strings.Contains(path, n.name) {
		return 0
	}
	if path != "" {
		path = path + "," + n.name
	} else {
		path = n.name
	}
	if n.name == "end" {
		fmt.Printf("path: %v\n", path)
		return 1
	}
	ends := 0
	for _, next := range n.next {
		ends += findEnd(next, path)
	}
	return ends
}

func main() {
	start := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("unique paths: %v\n", findEnd(start, ""))
}

func getNode(nodeMap map[string]*Node, name string) *Node {
	node := nodeMap[name]
	if node == nil {
		node = &Node{name: name}
		nodeMap[name] = node
	}
	return node
}

func parseInput(input string) *Node {
	allNodes := make(map[string]*Node)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		nodeNames := strings.Split(line, "-")
		a := getNode(allNodes, nodeNames[0])
		b := getNode(allNodes, nodeNames[1])
		a.next = append(a.next, b)
		if nodeNames[0] != "start" && nodeNames[1] != "end" {
			b.next = append(b.next, a)
		}
	}
	return allNodes["start"]
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
