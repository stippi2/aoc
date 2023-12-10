package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input1 = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

const input2 = `OF----7F7F7F7F-7OOOO
O|F--7||||||||FJOOOO
O||OFJ||||||||L7OOOO
FJL7L7LJLJ||LJIL-7OO
L--JOL7IIILJS7F-7L7O
OOOOF-JIIF7FJ|L7L7L7
OOOOL7IF7||L7|IL7L7|
OOOOO|FJLJ|FJ|F7|OLJ
OOOOFJL-7O||O||||OOO
OOOOL---JOLJOLJLJOOO`

func Test_partOne(t *testing.T) {
	m := parseInput(input1)
	assert.Equal(t, 8, partOne(m))
}

func Test_partTwo(t *testing.T) {
	m := parseInput(input2)
	assert.Equal(t, 8, partTwo(m))
}
