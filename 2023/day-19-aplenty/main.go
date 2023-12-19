package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Rule interface {
	apply(part Part) (bool, string)
}

type GreaterThanRule struct {
	targetWorkflow string
	property       string
	value          int
}

func (r GreaterThanRule) apply(part Part) (bool, string) {
	switch r.property {
	case "x":
		if part.x > r.value {
			return true, r.targetWorkflow
		}
	case "m":
		if part.m > r.value {
			return true, r.targetWorkflow
		}
	case "a":
		if part.a > r.value {
			return true, r.targetWorkflow
		}
	case "s":
		if part.s > r.value {
			return true, r.targetWorkflow
		}
	}
	return false, ""
}

type LessThanRule struct {
	targetWorkflow string
	property       string
	value          int
}

func (r LessThanRule) apply(part Part) (bool, string) {
	switch r.property {
	case "x":
		if part.x < r.value {
			return true, r.targetWorkflow
		}
	case "m":
		if part.m < r.value {
			return true, r.targetWorkflow
		}
	case "a":
		if part.a < r.value {
			return true, r.targetWorkflow
		}
	case "s":
		if part.s < r.value {
			return true, r.targetWorkflow
		}
	}
	return false, ""
}

type TargetWorkflowRule struct {
	targetWorkflow string
}

func (r TargetWorkflowRule) apply(part Part) (bool, string) {
	return true, r.targetWorkflow
}

type Workflow struct {
	name  string
	rules []Rule
}

type Part struct {
	x, m, a, s int
}

func partOne(workflows []Workflow, parts []Part) int {
	return 0
}

func partTwo(workflows []Workflow, parts []Part) int {
	return 0
}

func main() {
	now := time.Now()
	workflows, parts := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(workflows, parts)
	part2 := partTwo(workflows, parts)
	duration := time.Since(now)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseRule(input string) Rule {
	var rule Rule
	if strings.Contains(input, ":") {
		targetWorkflow := strings.Split(input, ":")[1]
		input = strings.Split(input, ":")[0]
		if strings.Contains(input, "<") {
			lt := &LessThanRule{targetWorkflow: targetWorkflow}
			lt.property = strings.Split(input, "<")[0]
			lt.value, _ = strconv.Atoi(strings.Split(input, "<")[1])
			rule = lt
		} else {
			gt := &GreaterThanRule{targetWorkflow: targetWorkflow}
			gt.property = strings.Split(input, ">")[0]
			gt.value, _ = strconv.Atoi(strings.Split(input, ">")[1])
			rule = gt
		}
	} else {
		rule = &TargetWorkflowRule{targetWorkflow: input}
	}
	return rule
}

func parseWorkflow(input string) Workflow {
	workflow := Workflow{}
	// px{a<2006:qkq,m>2090:A,rfg}
	nameAndRules := strings.Split(input, "{")
	workflow.name = nameAndRules[0]
	input = strings.TrimPrefix(input, workflow.name)
	input = strings.Trim(input, "{}")
	for _, rule := range strings.Split(input, ",") {
		workflow.rules = append(workflow.rules, parseRule(rule))
	}
	return workflow
}

func parsePart(input string) Part {
	//{x=787,m=2655,a=1222,s=2876}
	input = strings.Trim(input, "{}")
	properties := strings.Split(input, ",")
	part := Part{}
	for _, property := range properties {
		kv := strings.Split(property, "=")
		switch kv[0] {
		case "x":
			part.x, _ = strconv.Atoi(kv[1])
		case "m":
			part.m, _ = strconv.Atoi(kv[1])
		case "a":
			part.a, _ = strconv.Atoi(kv[1])
		case "s":
			part.s, _ = strconv.Atoi(kv[1])
		}
	}
	return part
}

func parseInput(input string) ([]Workflow, []Part) {
	sections := strings.Split(input, "\n\n")
	var workflows []Workflow
	for _, line := range strings.Split(sections[0], "\n") {
		workflows = append(workflows, parseWorkflow(line))
	}
	var parts []Part
	for _, line := range strings.Split(sections[1], "\n") {
		parts = append(parts, parsePart(line))
	}
	return workflows, parts
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
