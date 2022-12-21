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
	HasChild(child Node) bool
	ReverseEvaluate(neededResult int, human Node) int
}

type Number struct {
	value int
}

func (n *Number) Evaluate() int {
	return n.value
}

func (n *Number) HasChild(child Node) bool {
	return n == child
}

func (n *Number) ReverseEvaluate(neededResult int, human Node) int {
	if n == human {
		return neededResult
	}
	panic("descended wrong node")
}

type Operation struct {
	a, b       Node
	op         func(a, b Node) int
	reverseOpA func(result, valueB int) int
	reverseOpB func(result, valueA int) int
}

func (o *Operation) Evaluate() int {
	return o.op(o.a, o.b)
}

func (o *Operation) HasChild(child Node) bool {
	return child == o || o.a.HasChild(child) || o.b.HasChild(child)
}

func (o *Operation) ReverseEvaluate(neededResult int, human Node) int {
	if o.a.HasChild(human) {
		return o.a.ReverseEvaluate(o.reverseOpA(neededResult, o.b.Evaluate()), human)
	} else {
		return o.b.ReverseEvaluate(o.reverseOpB(neededResult, o.a.Evaluate()), human)
	}
}

func main() {
	start := time.Now()
	root, human := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("part 1: root evaluates to: %v (%v)\n", root.Evaluate(), time.Since(start))

	start = time.Now()
	rootOp := root.(*Operation)
	var humanResult int
	if rootOp.a.HasChild(human) {
		neededResult := rootOp.b.Evaluate()
		humanResult = rootOp.a.ReverseEvaluate(neededResult, human)
	} else {
		neededResult := rootOp.a.Evaluate()
		humanResult = rootOp.b.ReverseEvaluate(neededResult, human)
	}
	fmt.Printf("part 2: human needs to yell: %v (%v)\n", humanResult, time.Since(start))
}

func parseInput(input string) (Node, Node) {
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
				operation.reverseOpA = func(result, valueB int) int { return result - valueB }
				operation.reverseOpB = func(result, valueA int) int { return result - valueA }
			case "-":
				operation.op = func(a, b Node) int { return a.Evaluate() - b.Evaluate() }
				operation.reverseOpA = func(result, valueB int) int { return result + valueB }
				operation.reverseOpB = func(result, valueA int) int { return valueA - result }
			case "*":
				operation.op = func(a, b Node) int { return a.Evaluate() * b.Evaluate() }
				operation.reverseOpA = func(result, valueB int) int { return result / valueB }
				operation.reverseOpB = func(result, valueA int) int { return result / valueA }
			case "/":
				operation.op = func(a, b Node) int { return a.Evaluate() / b.Evaluate() }
				operation.reverseOpA = func(result, valueB int) int { return result * valueB }
				operation.reverseOpB = func(result, valueA int) int { return valueA / result }
			}
		}
	}
	return allNodes["root"], allNodes["humn"]
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
