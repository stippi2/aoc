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
	assert.Equal(t, 13, runRopeSimulation(exampleInput, 2))
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
	assert.Equal(t, 36, runRopeSimulation(exampleInput2, 10))
}
