package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Range struct {
	min, max int
}

func (r Range) isValid() bool {
	return r.min <= r.max
}

type RangeSet struct {
	ranges map[string]*Range
}

func (rs *RangeSet) clone() *RangeSet {
	return &RangeSet{ranges: map[string]*Range{
		"x": {rs.ranges["x"].min, rs.ranges["x"].max},
		"m": {rs.ranges["m"].min, rs.ranges["m"].max},
		"a": {rs.ranges["a"].min, rs.ranges["a"].max},
		"s": {rs.ranges["s"].min, rs.ranges["s"].max},
	}}
}

func (rs *RangeSet) String() string {
	return fmt.Sprintf("x: %d-%d, m: %d-%d, a: %d-%d, s: %d-%d",
		rs.ranges["x"].min, rs.ranges["x"].max,
		rs.ranges["m"].min, rs.ranges["m"].max,
		rs.ranges["a"].min, rs.ranges["a"].max,
		rs.ranges["s"].min, rs.ranges["s"].max)
}

type Rule interface {
	apply(part Part) (bool, string)
	applyToRangeSet(rangeSet *RangeSet) (bool, string)
	gateRangeSet(rangeSet *RangeSet)
}

type GreaterThanRule struct {
	targetWorkflow string
	property       string
	value          int
}

func (gtr GreaterThanRule) apply(part Part) (bool, string) {
	switch gtr.property {
	case "x":
		if part.x > gtr.value {
			return true, gtr.targetWorkflow
		}
	case "m":
		if part.m > gtr.value {
			return true, gtr.targetWorkflow
		}
	case "a":
		if part.a > gtr.value {
			return true, gtr.targetWorkflow
		}
	case "s":
		if part.s > gtr.value {
			return true, gtr.targetWorkflow
		}
	}
	return false, ""
}

func (gtr GreaterThanRule) applyToRangeSet(rangeSet *RangeSet) (bool, string) {
	r := rangeSet.ranges[gtr.property]
	if !r.isValid() || r.max <= gtr.value {
		return false, ""
	}
	r.min = max(r.min, gtr.value+1)
	return true, gtr.targetWorkflow
}

func (gtr GreaterThanRule) gateRangeSet(rangeSet *RangeSet) {
	r := rangeSet.ranges[gtr.property]
	r.max = min(r.max, gtr.value)
}

type LessThanRule struct {
	targetWorkflow string
	property       string
	value          int
}

func (ltr LessThanRule) apply(part Part) (bool, string) {
	switch ltr.property {
	case "x":
		if part.x < ltr.value {
			return true, ltr.targetWorkflow
		}
	case "m":
		if part.m < ltr.value {
			return true, ltr.targetWorkflow
		}
	case "a":
		if part.a < ltr.value {
			return true, ltr.targetWorkflow
		}
	case "s":
		if part.s < ltr.value {
			return true, ltr.targetWorkflow
		}
	}
	return false, ""
}

func (ltr LessThanRule) applyToRangeSet(rangeSet *RangeSet) (bool, string) {
	r := rangeSet.ranges[ltr.property]
	if !r.isValid() || r.min >= ltr.value {
		return false, ""
	}
	r.max = min(r.max, ltr.value-1)
	return true, ltr.targetWorkflow
}

func (ltr LessThanRule) gateRangeSet(rangeSet *RangeSet) {
	r := rangeSet.ranges[ltr.property]
	r.min = max(r.min, ltr.value)
}

type TargetWorkflowRule struct {
	targetWorkflow string
}

func (r TargetWorkflowRule) apply(_ Part) (bool, string) {
	return true, r.targetWorkflow
}

func (r TargetWorkflowRule) applyToRangeSet(_ *RangeSet) (bool, string) {
	return true, r.targetWorkflow
}

func (r TargetWorkflowRule) gateRangeSet(_ *RangeSet) {
}

type Workflow struct {
	name  string
	rules []Rule
}

type Part struct {
	x, m, a, s int
}

func partOne(workflows map[string]*Workflow, parts []Part) int {
	var accepted []Part

	for _, part := range parts {
		workflow := workflows["in"]
		decisionMade := false
		for {
			for _, rule := range workflow.rules {
				applied, targetWorkflow := rule.apply(part)
				if applied {
					switch targetWorkflow {
					case "A":
						accepted = append(accepted, part)
						decisionMade = true
					case "R":
						decisionMade = true
					default:
						workflow = workflows[targetWorkflow]
					}
					break
				}
			}
			if decisionMade {
				break
			}
		}
	}

	sumProperties := 0
	for _, part := range accepted {
		sumProperties += part.x + part.m + part.a + part.s
	}

	return sumProperties
}

func assembleAcceptedRanges(workflows map[string]*Workflow, workflowName string, rangeSet *RangeSet) []*RangeSet {
	workflow := workflows[workflowName]
	var resultSets []*RangeSet
	for _, rule := range workflow.rules {
		clonedRangeSet := rangeSet.clone()
		applied, targetWorkflow := rule.applyToRangeSet(rangeSet)
		if applied {
			switch targetWorkflow {
			case "A":
				resultSets = append(resultSets, rangeSet)
			case "R":
			default:
				resultSets = append(resultSets, assembleAcceptedRanges(workflows, targetWorkflow, rangeSet)...)
			}
			// Continue the other rules with "reverted" rangeSet
			rule.gateRangeSet(clonedRangeSet)
			rangeSet = clonedRangeSet
		}
	}
	return resultSets
}

func partTwo(workflows map[string]*Workflow) int {
	rangeSet := &RangeSet{ranges: map[string]*Range{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}}

	acceptedRangeSets := assembleAcceptedRanges(workflows, "in", rangeSet)
	for _, rs := range acceptedRangeSets {
		fmt.Printf("%s\n", rs)
	}
	sum := 0
	for _, rs := range acceptedRangeSets {
		fmt.Printf("%s\n", rs)
		product := 1
		product *= rs.ranges["x"].max - rs.ranges["x"].min + 1
		product *= rs.ranges["m"].max - rs.ranges["m"].min + 1
		product *= rs.ranges["a"].max - rs.ranges["a"].min + 1
		product *= rs.ranges["s"].max - rs.ranges["s"].min + 1
		sum += product
	}
	return sum
}

func main() {
	now := time.Now()
	workflows, parts := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(workflows, parts)
	part2 := partTwo(workflows)
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

func parseInput(input string) (map[string]*Workflow, []Part) {
	sections := strings.Split(input, "\n\n")
	workflows := make(map[string]*Workflow)
	for _, line := range strings.Split(sections[0], "\n") {
		workflow := parseWorkflow(line)
		workflows[workflow.name] = &workflow
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
