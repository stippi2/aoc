package day23

import (
	"aoc/2024/go/lib"
	"sort"
	"strings"
)

type Computer struct {
	name        string
	connectedTo map[string]*Computer
}

func (c *Computer) getThreeInterconnected() map[string]bool {
	setsOfThree := make(map[string]bool)
	for _, neighborA := range c.connectedTo {
		for _, neighborB := range c.connectedTo {
			if neighborB == neighborA {
				continue
			}
			if neighborA.connectedTo[neighborB.name] == neighborB {
				list := []string{c.name, neighborA.name, neighborB.name}
				sort.Strings(list)
				setsOfThree[strings.Join(list, ",")] = true
			}
		}
	}
	return setsOfThree
}

func parseInput(input string) map[string]*Computer {
	computers := make(map[string]*Computer)

	getOrCreate := func(name string) *Computer {
		computer := computers[name]
		if computer == nil {
			computer = &Computer{name: name, connectedTo: make(map[string]*Computer)}
			computers[name] = computer
		}
		return computer
	}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "-")
		a := getOrCreate(parts[0])
		b := getOrCreate(parts[1])
		a.connectedTo[b.name] = b
		b.connectedTo[a.name] = a
	}
	return computers
}

func countSetsOfThreeStartingWithT(computers map[string]*Computer) int {
	setsOfThree := make(map[string]bool)
	for _, computer := range computers {
		for setOfThree := range computer.getThreeInterconnected() {
			if strings.HasPrefix(setOfThree, "t") || strings.Contains(setOfThree, ",t") {
				setsOfThree[setOfThree] = true
			}
		}
	}
	return len(setsOfThree)
}

func Part1() any {
	input, _ := lib.ReadInput(23)
	computers := parseInput(input)
	return countSetsOfThreeStartingWithT(computers)
}

func Part2() any {
	return "Not implemented"
}
