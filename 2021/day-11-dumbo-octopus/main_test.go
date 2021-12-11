package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

var smallInput = `11111
19991
19191
19991
11111`

var expectedAfterOneStep = `34543
40004
50005
40004
34543`

func Test_parseInput(t *testing.T) {
	m := parseInput(exampleInput)
	assert.Equal(t, 10, m.width)
	assert.Equal(t, 10, m.height)
	assert.Equal(t, 5, m.get(0, 0))
	assert.Equal(t, 6, m.get(9, 9))
}
