package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `Time:      7  15   30
Distance:  9  40  200`

func Test_partOne(t *testing.T) {
	races := parseInput(input)
	assert.Equal(t, 3, len(races))
	assert.Equal(t, 288, partOne(races))
}

func Test_partTwo(t *testing.T) {
	races := parseInput(input)
	race := mergeRaces(races)
	assert.Equal(t, 71530, race.time)
	assert.Equal(t, 940200, race.recordDistance)
	assert.Equal(t, 71503, partTwo(races))
}
