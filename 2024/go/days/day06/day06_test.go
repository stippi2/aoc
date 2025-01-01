package day06

import (
	"aoc/2024/go/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func Test_Part1(t *testing.T) {
	grid := lib.NewGrid(example)
	start := findStart(grid)
	fieldsVisited, _ := trackGuardFields(grid, start, nil)
	assert.Equal(t, 41, len(fieldsVisited))
}

func Test_Part2(t *testing.T) {
	grid := lib.NewGrid(example)
	start := findStart(grid)
	fieldsVisited, _ := trackGuardFields(grid, start, nil)
	assert.Equal(t, 6, countLoopCausingObstacles(grid, start, fieldsVisited))
}
