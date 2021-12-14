package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Node struct {
	name string
	next []*Node
}

type CaveSystem struct {
	start    *Node
	allNodes map[string]*Node
}

type Rule interface {
	Allow(n *Node, path string) bool
}

type SmallCavesOnce struct {}

func (r *SmallCavesOnce) Allow(n *Node, path string) bool {
	return n.name != strings.ToLower(n.name) || !strings.Contains(path, n.name)
}

type OneSmallCaveTwice struct {
	singleSmallCave string
}

func (r *OneSmallCaveTwice) Allow(n *Node, path string) bool {
	if n.name != strings.ToLower(n.name) {
		return true
	}
	if n.name == r.singleSmallCave {
		visited := 0
		for _, node := range strings.Split(path, ",") {
			if node == n.name {
				visited++
			}
		}
		return visited < 2
	}
	return !strings.Contains(path, n.name)
}

func findPathsWithRule(n *Node, path string, rule Rule) []string {
	if !rule.Allow(n, path) {
		return nil
	}
	if path != "" {
		path = path + "," + n.name
	} else {
		path = n.name
	}
	if n.name == "end" {
		return []string{path}
	}
	var paths []string
	for _, next := range n.next {
		paths = append(paths, findPathsWithRule(next, path, rule)...)
	}
	return paths
}

func findPathsPart1(caves CaveSystem) []string {
	return findPathsWithRule(caves.start, "", &SmallCavesOnce{})
}

func findPathsPart2(caves CaveSystem) []string {
	uniquePaths := make(map[string]bool)
	for name := range caves.allNodes {
		if name == strings.ToLower(name) {
			if name == "start" || name == "end" {
				continue
			}
			paths := findPathsWithRule(caves.start, "", &OneSmallCaveTwice{name})
			for _, path := range paths {
				uniquePaths[path] = true
			}
		}
	}
	var allPaths []string
	for path := range uniquePaths {
		allPaths = append(allPaths, path)
	}
	sort.Strings(allPaths)
	return allPaths
}

func main() {
	caves := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("unique paths visiting small caves once: %v\n", len(findPathsPart1(caves)))
	fmt.Printf("unique paths visiting exactly one small caves twice: %v\n", len(findPathsPart2(caves)))
}

func getNode(nodeMap map[string]*Node, name string) *Node {
	node := nodeMap[name]
	if node == nil {
		node = &Node{name: name}
		nodeMap[name] = node
	}
	return node
}

func parseInput(input string) CaveSystem {
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
	return CaveSystem{start: allNodes["start"], allNodes: allNodes}
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
