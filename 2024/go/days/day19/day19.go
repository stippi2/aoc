package day19

import (
	"aoc/2024/go/lib"
	"sort"
	"strings"
)

func parseInput(input string) (towels, designs []string) {
	parts := strings.Split(input, "\n\n")
	towels = strings.Split(parts[0], ", ")
	designs = strings.Split(parts[1], "\n")
	return
}

func designIsPossible(design string, towels []string) bool {
	if len(design) == 0 {
		return true
	}
	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			if designIsPossible(design[len(towel):], towels) {
				return true
			}
		}
	}
	return false
}

func countPossibleDesigns(towels, designs []string) int {
	sum := 0
	sort.Slice(towels, func(i, j int) bool {
		return len(towels[i]) > len(towels[j])
	})
	for _, design := range designs {
		var relevantTowels []string
		for _, towel := range towels {
			if strings.Contains(design, towel) {
				relevantTowels = append(relevantTowels, towel)
			}
		}
		if designIsPossible(design, relevantTowels) {
			sum++
		}
	}
	return sum
}

func Part1() any {
	input, _ := lib.ReadInput(19)
	return countPossibleDesigns(parseInput(input))
}

func countArrangements(design string, towels []string) int {
	if len(design) == 0 {
		return 1
	}
	possibleArrangements := 0
	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			possibleArrangements += countArrangements(design[len(towel):], towels)
		}
	}
	return possibleArrangements
}

func countPossibleArrangements(towels, designs []string) int {
	sum := 0
	sort.Slice(towels, func(i, j int) bool {
		return len(towels[i]) > len(towels[j])
	})
	for _, design := range designs {
		var relevantTowels []string
		for _, towel := range towels {
			if strings.Contains(design, towel) {
				relevantTowels = append(relevantTowels, towel)
			}
		}
		sum += countArrangements(design, relevantTowels)
	}
	return sum
}

func Part2() any {
	input, _ := lib.ReadInput(19)
	return countPossibleArrangements(parseInput(input))
}
