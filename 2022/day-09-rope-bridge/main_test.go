package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func Test_Part1(t *testing.T) {
	rope := &Rope{}
	rope.appendKnots(2)
	runPositions(exampleInput, rope)
	assert.Equal(t, 13, len(rope.tail().visitedPositions))
}

var exampleInput2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func Test_Part2(t *testing.T) {
	rope := &Rope{}
	rope.appendKnots(10)
	runPositions(exampleInput2, rope)
	assert.Equal(t, 36, len(rope.tail().visitedPositions))
}
