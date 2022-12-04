package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func Test_sumCompletelyOverlappingPairs(t *testing.T) {
	pairs := parseInput(exampleInput)
	sum := sumCompletelyOverlappingPairs(pairs)
	assert.Equal(t, 2, sum)
}

func Test_sumOverlappingPairs(t *testing.T) {
	pairs := parseInput(exampleInput)
	sum := sumOverlappingPairs(pairs)
	assert.Equal(t, 4, sum)
}
