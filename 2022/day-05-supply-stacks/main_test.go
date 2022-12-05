package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func Test_Part1(t *testing.T) {
	stacks, instructions := parseInput(exampleInput, 3)
	assert.Equal(t, []CrateStack{
		{
			[]Crate{{"Z"}, {"N"}},
		},
		{
			[]Crate{{"M"}, {"C"}, {"D"}},
		},
		{
			[]Crate{{"P"}},
		},
	}, stacks)
	assert.Equal(t, []MoveInstruction{
		{from: 2, to: 1, count: 1},
		{from: 1, to: 3, count: 3},
		{from: 2, to: 1, count: 2},
		{from: 1, to: 2, count: 1},
	}, instructions)
	applyMoveInstructions(stacks, instructions)
	topCrates := getTopCrates(stacks)
	assert.Equal(t, "CMZ", topCrates)
}
