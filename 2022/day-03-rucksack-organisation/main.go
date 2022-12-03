package main

import (
	"fmt"
	"os"
	"strings"
)

func priority(item string) int {
	if strings.ToLower(item) == item {
		return int(item[0]) - int('a') + 1
	}
	return int(item[0]) - int('A') + 27
}

func priorityOfDuplicateItem(items []string) int {
	// Go has no hash set
	seenItems := make(map[string]bool)
	for i := 0; i < len(items)/2; i++ {
		seenItems[items[i]] = true
	}
	for i := len(items) / 2; i < len(items); i++ {
		if seenItems[items[i]] {
			return priority(items[i])
		}
	}
	return 0
}

func sumContents(rucksackContents []string) int {
	sum := 0
	for i, contents := range rucksackContents {
		items := strings.Split(contents, "")
		dupPriority := priorityOfDuplicateItem(items)
		sum += dupPriority
		fmt.Printf("rucksack: %v, priority: %v\n", i, dupPriority)
	}
	return sum
}

func sumBadges(rucksackContents []string) int {
	sum := 0
	allItems := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < len(rucksackContents); i += 3 {
		for x := 0; x < len(allItems); x++ {
			containedInAll := true
			for k := 0; k < 3; k++ {
				if !strings.Contains(rucksackContents[i+k], string(allItems[x])) {
					containedInAll = false
					break
				}
			}
			if containedInAll {
				sum += priority(string(allItems[x]))
			}
		}
	}
	return sum
}

func main() {
	rucksackContents := parseInput(loadInput("puzzle-input.txt"))
	sum := sumContents(rucksackContents)
	fmt.Printf("total sum of item priorities: %v\n", sum)
	sum = sumBadges(rucksackContents)
	fmt.Printf("total sum of badges: %v\n", sum)
}

func parseInput(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, "\n")
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
