package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func Test_partOne(t *testing.T) {
	maps := parseInput(input)
	assert.Len(t, maps, 2)
	assert.Equal(t, 405, partOne(maps))
}

func Test_partOneEdge(t *testing.T) {
	maps := parseInput(`#########.##.##
##.....#.####.#
#.#.##...####..
...#####..##...
#.....#...##...
.##..#..#.##.#.
.##..#...#..#..
####..##.#..#.#
####..##.#..#.#`)
	assert.True(t, maps[0].isMirrorAxis(8, &maps[0].rows, -1, -1))
	x, y := maps[0].findMirrorAxisValue(-1, -1)
	assert.Equal(t, 0, x)
	assert.Equal(t, 800, y)
}

func Test_partTwo(t *testing.T) {
	maps := parseInput(input)
	assert.True(t, maps[0].isMirrorAxis(3, &maps[0].rows, 0, 0))
	assert.Equal(t, 400, partTwo(maps))
}

func Test_partTwoEdge(t *testing.T) {
	maps := parseInput(`.#..###..##.#
.#..###..##.#
##..#.###.###
#.##...#..##.
..####...#..#
##..#..#.#...
###..#.....##
....##.##..#.
....##.#...#.`)
	assert.True(t, maps[0].isMirrorAxis(8, &maps[0].rows, 7, 8))
	assert.Equal(t, 800, partTwo(maps))
}
