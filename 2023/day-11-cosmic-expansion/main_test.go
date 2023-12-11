package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func Test_partOne(t *testing.T) {
	m := parseInput(input)
	assert.Equal(t, 9, len(m.galaxies))
	distancesSum, pairs := partOne(m)
	assert.Equal(t, 13, m.width)
	assert.Equal(t, 12, m.height)
	assert.Equal(t, `....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......`, m.String())
	assert.Equal(t, 36, pairs)
	assert.Equal(t, 374, distancesSum)
}

func Test_expandSpace(t *testing.T) {
	m := parseInput(`.#
..
.#`)
	m.expandSpace(1)
	assert.Equal(t, 3, m.width)
	assert.Equal(t, 4, m.height)
}

func Test_partTwo(t *testing.T) {
	m := parseInput(input)
	m.expandSpace(99)
	distancesSum, pairs := m.sumDistances()
	assert.Equal(t, 36, pairs)
	assert.Equal(t, int64(8410), distancesSum)
}
