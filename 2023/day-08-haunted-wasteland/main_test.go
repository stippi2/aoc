package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input1 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

const input2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

func Test_partOne(t *testing.T) {
	directions, nodes := parseInput(input1)
	assert.Equal(t, 2, partOne(directions, nodes))

	directions, nodes = parseInput(input2)
	assert.Equal(t, 6, partOne(directions, nodes))
}

func Test_partTwo(t *testing.T) {
	_, _ = parseInput(input1)
	assert.Equal(t, 0, partTwo())
}
