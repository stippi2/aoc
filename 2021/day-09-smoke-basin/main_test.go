package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `2199943210
3987894921
9856789892
8767896789
9899965678`

func Test_parseInput(t *testing.T) {
	m := parseInput(exampleInput)
	assert.Equal(t, 10, m.width)
	assert.Equal(t, 5, m.height)
	assert.Equal(t, 9, m.get(1, 1))
	assert.Equal(t, 9, m.get(1, 1))
	assert.Equal(t, 8, m.get(9, 4))
}

func Test_sumRiskLevels(t *testing.T) {
	m := parseInput(exampleInput)
	assert.Equal(t, 15, m.sumRiskLevels())
}

func Test_floodBasin(t *testing.T) {
	m := parseInput(exampleInput)
	assert.Equal(t, 3, m.floodBasin(1, 0))
	assert.Equal(t, 9, m.floodBasin(9, 0))
	assert.Equal(t, 14, m.floodBasin(2, 2))
	assert.Equal(t, 9, m.floodBasin(6, 4))
}

func Test_findLargestBasins(t *testing.T) {
	m := parseInput(exampleInput)
	assert.Equal(t, 1134, m.findLargestBasins(3))
}
