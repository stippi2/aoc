package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

func Test_parseInput(t *testing.T) {
	blueprints := parseInput(exampleInput)
	if assert.Len(t, blueprints, 2) {
		assert.Equal(t, Blueprint{costs: map[string]map[string]int{
			"ore":      {"ore": 4},
			"clay":     {"ore": 2},
			"obsidian": {"ore": 3, "clay": 14},
			"geode":    {"ore": 2, "obsidian": 7},
		}}, blueprints[0])
		assert.Equal(t, Blueprint{costs: map[string]map[string]int{
			"ore":      {"ore": 2},
			"clay":     {"ore": 3},
			"obsidian": {"ore": 3, "clay": 8},
			"geode":    {"ore": 3, "obsidian": 12},
		}}, blueprints[1])
	}
}

func Test_findQualityLevel(t *testing.T) {
	blueprints := parseInput(exampleInput)
	assert.Equal(t, 9, findQualityLevel(blueprints[0], 24))
	//	assert.Equal(t, 12, findQualityLevel(blueprints[1], 24))
}
