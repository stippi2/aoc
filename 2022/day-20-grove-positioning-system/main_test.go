package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `1
2
-3
3
-2
0
4`

func Test_part1(t *testing.T) {
	s := parseInput(exampleInput)
	assert.Equal(t, "[0, 1], [1, 2], [2, -3], [3, 3], [4, -2], [5, 0], [6, 4]", s.String())
	s.mix()
	assert.Equal(t, "[0, 1], [1, 2], [2, -3], [6, 4], [5, 0], [3, 3], [4, -2]", s.String())

	assert.Equal(t, 3, s.findGroveCoordinates())
}

var exampleInput2 = `5
0`

func Test_part1b(t *testing.T) {
	s := parseInput(exampleInput2)
	assert.Equal(t, "[0, 5], [1, 0]", s.String())
	s.mix()
	assert.Equal(t, "[1, 0], [0, 5]", s.String())
}

func Test_wrapToSelf(t *testing.T) {
	s := parseInput(`0
2
0`)
	assert.Equal(t, "[0, 0], [1, 2], [2, 0]", s.String())
	s.mix()
	assert.Equal(t, "[0, 0], [1, 2], [2, 0]", s.String())
}
