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

const input3 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func Test_partOne(t *testing.T) {
	directions, nodes := parseInput(input1)
	assert.Equal(t, 2, partOne(directions, nodes))

	directions, nodes = parseInput(input2)
	assert.Equal(t, 6, partOne(directions, nodes))
}

func Test_partTwo(t *testing.T) {
	directions, nodes := parseInput(input3)
	assert.Equal(t, 6, partTwo(directions, nodes))
}
