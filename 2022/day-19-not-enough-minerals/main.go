package main

import (
	"fmt"
	"os"
	"strings"
)

type Blueprint struct {
	costs map[string]map[string]int
}

type Inventory struct {
	minerals map[string]int
	robots   map[string]int
}

func (i *Inventory) canBuildRobot(blueprint Blueprint, kind string) bool {
	for mineral, howMuch := range blueprint.costs[kind] {
		if i.minerals[mineral] < howMuch {
			return false
		}
	}
	return true
}

func main() {
	//	blueprints := parseInput(loadInput("puzzle-input.txt"))
}

func parseInput(input string) []Blueprint {
	lines := strings.Split(input, "\n")
	blueprints := make([]Blueprint, len(lines))
	minerals := []string{"ore", "clay", "obsidian", "geode"}
	for i, line := range lines {
		line = strings.TrimPrefix(line, fmt.Sprintf("Blueprint %d: ", i+1))
		line = strings.TrimSuffix(line, ".")
		parts := strings.Split(line, ". ")
		blueprints[i].costs = make(map[string]map[string]int)
		for j, robotKind := range minerals {
			costsString := strings.TrimPrefix(parts[j], fmt.Sprintf("Each %s robot costs ", robotKind))
			var costs1 int
			var costs2 int
			costs := make(map[string]int)
			switch robotKind {
			case "ore", "clay":
				_, _ = fmt.Sscanf(costsString, "%d ore", &costs1)
				costs["ore"] = costs1
			case "obsidian":
				_, _ = fmt.Sscanf(costsString, "%d ore and %d clay", &costs1, &costs2)
				costs["ore"] = costs1
				costs["clay"] = costs2
			case "geode":
				_, _ = fmt.Sscanf(costsString, "%d ore and %d obsidian", &costs1, &costs2)
				costs["ore"] = costs1
				costs["obsidian"] = costs2
			}
			blueprints[i].costs[robotKind] = costs
		}
	}
	return blueprints
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
