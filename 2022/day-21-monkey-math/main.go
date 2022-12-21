package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node interface {
	Evaluate() int
}

type Number struct {
	value int
}

func (n *Number) Evaluate() int {
	return n.value
}

type Operation struct {
	a, b Node
	op   func(a, b Node) int
}

func (o *Operation) Evaluate() int {
	return o.op(o.a, o.b)
}

func main() {
	start := time.Now()
	root := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("part 1: root evaluates to: %v (%v)\n", root.Evaluate(), time.Since(start))
}

func parseInput(input string) Node {
	lines := strings.Split(input, "\n")
	allNodes := make(map[string]Node)
	// Create all Node instances
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		if len(strings.Split(parts[1], " ")) == 1 {
			value, _ := strconv.Atoi(parts[1])
			allNodes[parts[0]] = &Number{value: value}
		} else {
			allNodes[parts[0]] = &Operation{}
		}
	}
	// Connect the instances
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		opParts := strings.Split(parts[1], " ")
		if len(opParts) == 3 {
			operation := allNodes[parts[0]].(*Operation)
			operation.a = allNodes[opParts[0]]
			operation.b = allNodes[opParts[2]]
			switch opParts[1] {
			case "+":
				operation.op = func(a, b Node) int { return a.Evaluate() + b.Evaluate() }
			case "-":
				operation.op = func(a, b Node) int { return a.Evaluate() - b.Evaluate() }
			case "*":
				operation.op = func(a, b Node) int { return a.Evaluate() * b.Evaluate() }
			case "/":
				operation.op = func(a, b Node) int { return a.Evaluate() / b.Evaluate() }
			}
		}
	}
	return allNodes["root"]
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
