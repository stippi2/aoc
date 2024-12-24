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

// Convert a list of computers to a sorted, comma-separated string of their names
func computersToString(computers []*Computer) string {
	names := make([]string, len(computers))
	for i, c := range computers {
		names[i] = c.name
	}
	sort.Strings(names)
	return strings.Join(names, ",")
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

// Find largest clique starting from a specific computer
func findCliqueFromComputer(start *Computer) []*Computer {
	// Start with the computer and all its neighbors
	candidates := make([]*Computer, 0, len(start.connectedTo)+1)
	candidates = append(candidates, start)
	for _, neighbor := range start.connectedTo {
		candidates = append(candidates, neighbor)
	}

	// Current best clique
	result := make([]*Computer, 0)

	// Try all possible combinations, starting with largest
	for size := len(candidates); size >= 1; size-- {
		current := make([]*Computer, 0, size)
		found := false

		var generateCombinations func(int, int, []*Computer)
		generateCombinations = func(startIdx, remaining int, subset []*Computer) {
			if found {
				return
			}

			if remaining == 0 {
				// Check if all computers in subset are connected to each other
				isClique := true
				for i := 0; i < len(subset); i++ {
					for j := i + 1; j < len(subset); j++ {
						if _, exists := subset[i].connectedTo[subset[j].name]; !exists {
							isClique = false
							break
						}
					}
					if !isClique {
						break
					}
				}
				if isClique {
					current = make([]*Computer, len(subset))
					copy(current, subset)
					found = true
				}
				return
			}

			for i := startIdx; i <= len(candidates)-remaining; i++ {
				generateCombinations(i+1, remaining-1, append(subset, candidates[i]))
			}
		}

		generateCombinations(0, size, make([]*Computer, 0, size))

		if found {
			result = current
			break
		}
	}

	return result
}

// Find the largest clique in the network
func findLargestClique(computers map[string]*Computer) []*Computer {
	// Store all found cliques by their string representation
	foundCliques := make(map[string][]*Computer)

	// Find cliques starting from each computer
	for _, computer := range computers {
		clique := findCliqueFromComputer(computer)
		cliqueStr := computersToString(clique)
		foundCliques[cliqueStr] = clique
	}

	// Find the largest clique
	var maxClique []*Computer
	maxSize := 0
	for _, clique := range foundCliques {
		if len(clique) > maxSize {
			maxSize = len(clique)
			maxClique = clique
		}
	}

	return maxClique
}

func getPartyPassword(computers map[string]*Computer) string {
	largestClique := findLargestClique(computers)
	asStrings := make([]string, len(largestClique))
	for i, computer := range largestClique {
		asStrings[i] = computer.name
	}
	sort.Strings(asStrings)
	return strings.Join(asStrings, ",")
}

func Part1() any {
	input, _ := lib.ReadInput(23)
	computers := parseInput(input)
	return countSetsOfThreeStartingWithT(computers)
}

func Part2() any {
	input, _ := lib.ReadInput(23)
	computers := parseInput(input)
	return getPartyPassword(computers)
}
