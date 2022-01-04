package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`

var examplePartTwo = `#############
#...........#
###B#C#B#D###
  #D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########`

func Test_parseInput(t *testing.T) {
	assert.Equal(t, example, parseInput(example).String())
}

func Test_getAmphipods(t *testing.T) {
	m := parseInput(example)
	pods := getAmphipods(m)
	expected := []Amphipod{
		{
			x:    3,
			y:    2,
			kind: 'B',
		},
		{
			x:    5,
			y:    2,
			kind: 'C',
		},
		{
			x:    7,
			y:    2,
			kind: 'B',
		},
		{
			x:    9,
			y:    2,
			kind: 'D',
		},
		{
			x:    3,
			y:    3,
			kind: 'A',
		},
		{
			x:    5,
			y:    3,
			kind: 'D',
		},
		{
			x:    7,
			y:    3,
			kind: 'C',
		},
		{
			x:    9,
			y:    3,
			kind: 'A',
		},
	}
	assert.Equal(t, expected, pods)
}

func Test_solve(t *testing.T) {
	m := parseInput(example)
	energy := solve(m, emptyMap)
	assert.Equal(t, 12521, energy)
}

func Test_solvePartTwo(t *testing.T) {
	m := parseInput(examplePartTwo)
	energy := solve(m, emptyMapPart2)
	assert.Equal(t, 44169, energy)
}

func Test_possibleMoves(t *testing.T) {
	tests := []struct{
		mapInput      string
		expectedMoves []Move
	}{
		{
			mapInput: `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`,
  			expectedMoves: []Move{
				{4, 1, 2},
				{6, 1, 4},
				{8, 1, 6},
				{10, 1, 8},
				{11, 1, 9},
				{2, 1, 2},
				{1, 1, 3},
			},
		},
		{
			mapInput: `#############
#B..........#
###.#C#B#D###
  #A#D#C#A#
  #########`,
			expectedMoves: nil,
		},
		{
			mapInput: `#############
#B..C.......#
###A#.#.#D###
  #A#B#C#D#
  #########`,
			expectedMoves: nil,
		},
		{
			mapInput: `#############
#B..........#
###A#.#.#D###
  #A#C#C#D#
  #########`,
			expectedMoves: nil,
		},
		{
			mapInput: `#############
#B..........#
###A#.#.#D###
  #A#B#C#D#
  #########`,
			expectedMoves: []Move{
				{5, 2, 5},
			},
		},
		{
			mapInput: `#############
#B..........#
###A#.#.#D###
  #A#.#C#D#
  #########`,
			expectedMoves: []Move{
				{5, 3, 6},
			},
		},
		{
			mapInput: `#############
#.....B.....#
###A#.#.#D###
  #A#.#C#D#
  #########`,
			expectedMoves: []Move{
				{5, 3, 3},
			},
		},
		{
			mapInput: `#############
#.....C....B#
###B#.#.#D###
  #D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########`,
			expectedMoves: nil,
		},
	}

	for _, test := range tests {
		m := parseInput(test.mapInput)
		pod := getAmphipods(m)[0]
		moves := pod.possibleMoves(m)
		assert.Equal(t, test.expectedMoves, moves)
	}
}
