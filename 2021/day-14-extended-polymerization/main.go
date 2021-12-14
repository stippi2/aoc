package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type PolymerProcess struct {
	polymer        string
	insertionRules map[string]string
}

func (p *PolymerProcess) applyRules() {
	polymer := make([]byte, len(p.polymer)*2-1)
	for i := 0; i < len(p.polymer)-1; i++ {
		key := p.polymer[i : i+2]
		insertion := p.insertionRules[key]
		polymer[i*2] = p.polymer[i]
		polymer[i*2+1] = insertion[0]
	}
	polymer[len(polymer)-1] = p.polymer[len(p.polymer)-1]
	p.polymer = string(polymer)
}

func (p *PolymerProcess) getElementCounts() map[string]int {
	counts := make(map[string]int)
	for i := 0; i < len(p.polymer); i++ {
		letter := p.polymer[i : i+1]
		counts[letter] = counts[letter] + 1
	}
	return counts
}

func occurrences(counts map[string]int) (min, max int) {
	min = math.MaxInt32
	max = math.MinInt32
	for _, count := range counts {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	return
}

func main() {
	p := parseInput(loadInput("puzzle-input.txt"))
	for step := 0; step < 10; step++ {
		p.applyRules()
	}
	min, max := occurrences(p.getElementCounts())
	fmt.Printf("max element minus min element: %v\n", max-min)
}

func parseInput(input string) PolymerProcess {
	inputSections := strings.Split(input, "\n\n")
	if len(inputSections) != 2 {
		panic("expected sections in the input separated by a double line-break")
	}

	process := PolymerProcess{
		polymer:        inputSections[0],
		insertionRules: make(map[string]string),
	}

	insertionRules := strings.Split(inputSections[1], "\n")
	for _, rule := range insertionRules {
		parts := strings.Split(rule, " -> ")
		process.insertionRules[parts[0]] = parts[1]
	}
	return process
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
