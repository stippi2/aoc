package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`

func Test_part1(t *testing.T) {
	m, instructions := parseInput(exampleInput)
	explorer := m.startingPos()
	fmt.Printf("start pos: %s\n", explorer.location)
	executeInstructions(m, instructions, explorer, handleWrapPartOne)
	printMap(m, explorer)
	assert.Equal(t, Pos{7, 5}, explorer.location)
	assert.Equal(t, 6032, explorer.getPassword())
}

func Test_part2(t *testing.T) {
	tests := []struct {
		label            string
		explorer         Explorer
		expectedLocation Pos
		expectedFacing   Pos
	}{
		{
			label: "FC",
			explorer: Explorer{
				location: Pos{149, 5},
				facing:   Pos{1, 0},
			},
			expectedLocation: Pos{99, 144},
			expectedFacing:   Pos{-1, 0},
		},
		{
			label: "FD",
			explorer: Explorer{
				location: Pos{110, 49},
				facing:   Pos{0, 1},
			},
			expectedLocation: Pos{99, 60},
			expectedFacing:   Pos{-1, 0},
		},
		{
			label: "DF",
			explorer: Explorer{
				location: Pos{99, 60},
				facing:   Pos{1, 0},
			},
			expectedLocation: Pos{110, 49},
			expectedFacing:   Pos{0, -1},
		},
		{
			label: "AF",
			explorer: Explorer{
				location: Pos{20, 199},
				facing:   Pos{0, 1},
			},
			expectedLocation: Pos{120, 0},
			expectedFacing:   Pos{0, 1},
		},
		{
			label: "FA",
			explorer: Explorer{
				location: Pos{120, 0},
				facing:   Pos{0, -1},
			},
			expectedLocation: Pos{20, 199},
			expectedFacing:   Pos{0, -1},
		},
		{
			label: "CA-1",
			explorer: Explorer{
				location: Pos{50, 149},
				facing:   Pos{0, 1},
			},
			expectedLocation: Pos{49, 150},
			expectedFacing:   Pos{-1, 0},
		},
		{
			label: "CA-2",
			explorer: Explorer{
				location: Pos{99, 149},
				facing:   Pos{0, 1},
			},
			expectedLocation: Pos{49, 199},
			expectedFacing:   Pos{-1, 0},
		},
		{
			label: "BE",
			explorer: Explorer{
				location: Pos{0, 102},
				facing:   Pos{-1, 0},
			},
			expectedLocation: Pos{50, 47},
			expectedFacing:   Pos{1, 0},
		},
		{
			label: "EA",
			explorer: Explorer{
				location: Pos{80, 0},
				facing:   Pos{0, -1},
			},
			expectedLocation: Pos{0, 180},
			expectedFacing:   Pos{1, 0},
		},
	}
	for _, test := range tests {
		newLocation, newFacing := handleWrapPartTwo(nil, &test.explorer)
		assert.Equal(t, test.expectedLocation, newLocation, test.label)
		assert.Equal(t, test.expectedFacing, newFacing, test.label)
	}
}
