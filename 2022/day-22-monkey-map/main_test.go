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
	executeInstructions(m, instructions, explorer)
	printMap(m, explorer)
	assert.Equal(t, Pos{7, 5}, explorer.location)
	assert.Equal(t, 6032, explorer.getPassword())
}
