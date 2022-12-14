package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

func Test_part1(t *testing.T) {
	cave := parseInput(exampleInput)
	assert.Equal(t, 24, simulateSand(Pos{500, 0}, cave))
}

func Test_part2(t *testing.T) {
	cave := parseInput(exampleInput)
	cave.floor = cave.bottom + 2
	assert.Equal(t, 93, simulateSand(Pos{500, 0}, cave))
}
