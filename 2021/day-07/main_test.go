package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `16,1,2,0,4,2,7,1,2,14`

func Test_parseSequence(t *testing.T) {
	assert.Equal(t, []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, parseSequence(exampleInput))
}

func Test_minMax(t *testing.T) {
	seq := parseSequence(exampleInput)
	min, max := minMax(seq)
	assert.Equal(t, min, 0)
	assert.Equal(t, max, 16)
}

func Test_findOptimumPos(t *testing.T) {
	positions := parseSequence(exampleInput)
	minFuel, bestPos := findOptimumPos(positions)
	assert.Equal(t, 5, bestPos)
	assert.Equal(t, 168, minFuel)
}
