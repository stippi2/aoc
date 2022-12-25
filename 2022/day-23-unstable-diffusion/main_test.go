package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

func Test_simulateOneDiffusionRound(t *testing.T) {
	m := parseInput(exampleInput)
	assert.Equal(t, exampleInput, m.String())
	simulateDiffusion(m, 1)
	assert.Equal(t, `.....#...
...#...#.
.#..#.#..
.....#..#
..#.#.##.
#..#.#...
#.#.#.##.
.........
..#..#...`, m.String())
}

func Test_part1(t *testing.T) {
	m := parseInput(exampleInput)
	simulateDiffusion(m, 10)
	assert.Equal(t, 110, countEmptyPositions(m))
}

func Test_part2(t *testing.T) {
	m := parseInput(exampleInput)
	assert.Equal(t, 20, simulateDiffusionUntilSettled(m))
}
