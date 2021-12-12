package main

import (
	"io/ioutil"
	"strings"
)

type Node struct {
	name string
	next []*Node
}

func main() {
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
	}
	return allNodes["start"]
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
