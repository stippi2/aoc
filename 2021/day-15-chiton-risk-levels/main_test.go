package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

var smallExample = `19111
19191
11191`

func Test_parseInput(t *testing.T) {
	m := parseInput(exampleInput)
	assert.Equal(t, 10, m.width)
	assert.Equal(t, 10, m.height)
	assert.Equal(t, 1, m.get(0, 0))
	assert.Equal(t, 8, m.get(8, 9))
}

func Test_findPath(t *testing.T) {
	m := parseInput(exampleInput)
	assert.Equal(t, 40, m.findPath())
}

func Test_findPathSmall(t *testing.T) {
	m := parseInput(smallExample)
	assert.Equal(t, 10, m.findPath())
}
