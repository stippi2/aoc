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

	// combinations tracks how many times a particular combination of two elements occurs.
	combinations map[string]int
	// elementCounts tracks how often individual elements occur.
	// This is necessary since the map combinations contains overlapping elements.
	elementCounts map[string]int
}

func (p *PolymerProcess) init(polymer string) {
	p.polymer = polymer
	p.insertionRules = make(map[string]string)
	p.combinations = make(map[string]int)
	p.elementCounts = make(map[string]int)
	for i := 0; i < len(polymer); i++ {
		if i < len(polymer) - 1 {
			combination := p.polymer[i : i+2]
			p.combinations[combination] = p.combinations[combination] + 1
		}
		element := polymer[i : i+1]
		p.elementCounts[element] = p.elementCounts[element] + 1
	}
}

func (p *PolymerProcess) applyRules() {
	// Track the polymer itself as long as that's reasonable
	if len(p.polymer) < 5000 {
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
	oldCombinations := make(map[string]int)
	for combination, count := range p.combinations {
		oldCombinations[combination] = count
	}
	for combination, count := range oldCombinations {
		insertion := p.insertionRules[combination]
		comboLeft := combination[:1] + insertion
		comboRight := insertion + combination[1:]

		p.combinations[comboLeft] += count
		p.combinations[comboRight] += count

		newCount := p.combinations[combination] - count
		if newCount > 0 {
			p.combinations[combination] = newCount
		} else {
			delete(p.combinations, combination)
		}

		p.elementCounts[insertion] += count
	}
}

func (p *PolymerProcess) getElementCounts() map[string]int {
	return p.elementCounts
}

func occurrences(counts map[string]int) (min, max int64, minElement, maxElement string) {
	min = math.MaxInt64
	max = math.MinInt64
	for element, c := range counts {
		count := int64(c)
		if count < min {
			min = count
			minElement = element
		}
		if count > max {
			max = count
			maxElement = element
		}
	}
	return
}

func main() {
	p := parseInput(loadInput("puzzle-input.txt"))
	for step := 0; step < 40; step++ {
		p.applyRules()
		if step == 9 || step == 39 {
			min, max, minElement, maxElement := occurrences(p.getElementCounts())
			diff := max-min
			fmt.Printf("max element (%s) minus min element (%s) after %v steps: %v\n", maxElement, minElement, step + 1, diff)
		}
	}
}

func parseInput(input string) PolymerProcess {
	inputSections := strings.Split(input, "\n\n")
	if len(inputSections) != 2 {
		panic("expected sections in the input separated by a double line-break")
	}

	process := PolymerProcess{}
	process.init(inputSections[0])

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
