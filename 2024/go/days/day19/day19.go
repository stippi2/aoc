package day19

import (
	"aoc/2024/go/lib"
	"fmt"
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

// func getSplitPositions(design string, towels []string) []int {
// 	splitPositions := []int{0}
// 	for i := 1; i < len(design); i++ {
// 		stripePair := design[i-1 : i+1]
// 		hasPossibleMatch := false

// 		for _, towel := range towels {
// 			// Find all occurances of the pair in the towel
// 			for j := 0; j < len(towel)-1; j++ {
// 				if towel[j:j+2] == stripePair {
// 					// Check if the part before the pair matches the design
// 					prefix := towel[:j]
// 					if len(prefix) <= i-1 && design[i-1-len(prefix):i-1] == prefix {
// 						// Check if the part after the pair matches the design
// 						suffix := towel[j+2:]
// 						if len(suffix) <= len(design)-i-1 && design[i+1:i+1+len(suffix)] == suffix {
// 							hasPossibleMatch = true
// 							break
// 						}
// 					}
// 				}
// 			}
// 			if hasPossibleMatch {
// 				break
// 			}
// 		}

// 		if !hasPossibleMatch {
// 			splitPositions = append(splitPositions, i)
// 		}
// 	}
// 	splitPositions = append(splitPositions, len(design))
// 	return splitPositions
// }

func getSplitPositions(design string, towels []string) []int {
	splitPositions := []int{0}
	for i := 1; i < len(design); i++ {
		// Check if this position is covered by any single towel
		isCovered := false
		for _, towel := range towels {
			// For each position in the towel that could cover position i
			for j := 1; j < len(towel); j++ {
				if i >= j && i+len(towel)-j <= len(design) {
					if design[i-j:i+len(towel)-j] == towel {
						isCovered = true
						break
					}
				}
			}
			if isCovered {
				break
			}
		}
		if !isCovered {
			splitPositions = append(splitPositions, i)
		}
	}
	splitPositions = append(splitPositions, len(design))
	return splitPositions
}

func countArrangements(design string, towels []string, cache map[string]int) int {
	if cached, exists := cache[design]; exists {
		return cached
	}

	if len(design) == 0 {
		return 1
	}

	possibleArrangements := 0
	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			possibleArrangements += countArrangements(design[len(towel):], towels, cache)
		}
	}

	cache[design] = possibleArrangements
	return possibleArrangements
}

func countPossibleArrangements(towels, designs []string) int {
	sum := 0
	sort.Slice(towels, func(i, j int) bool {
		return len(towels[i]) > len(towels[j])
	})
	cache := make(map[string]int)
	for _, design := range designs {
		if !designIsPossible(design, towels) {
			fmt.Printf("design %s is not possible\n", design)
			continue
		}

		splitPositions := getSplitPositions(design, towels)

		var designParts []string
		for j := 0; j < len(splitPositions)-1; j++ {
			designParts = append(designParts, design[splitPositions[j]:splitPositions[j+1]])
		}
		fmt.Printf("design %s, possible splits: %v, parts: %v", design, splitPositions, designParts)
		arrangements := 1
		for _, designPart := range designParts {
			count, exists := cache[designPart]
			if !exists {
				count = countArrangements(designPart, towels, cache)
				cache[designPart] = count
			}
			arrangements *= count
			if arrangements == 0 {
				break
			}
		}
		fmt.Printf(" found arrangments: %v\n", arrangements)
		sum += arrangements
	}
	return sum
}

func Part2() any {
	input, _ := lib.ReadInput(19)
	return countPossibleArrangements(parseInput(input))
}
