package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const input = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func (m *Map) String() string {
	var sb strings.Builder
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if m.squareRocks[Pos{x, y}] {
				sb.WriteString("#")
			} else if m.roundRocks[Pos{x, y}] {
				sb.WriteString("O")
			} else {
				sb.WriteString(".")
			}
		}
		if y < m.height-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func Test_partOne(t *testing.T) {
	m := parseInput(input)
	assert.Equal(t, 136, partOne(m))
}

func Test_partTwo(t *testing.T) {
	m := parseInput(input)
	m.tiltCycle()
	assert.Equal(t, `.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....`, m.String())
	m.tiltCycle()
	assert.Equal(t, `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O`, m.String())
	m.tiltCycle()
	assert.Equal(t, `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O`, m.String())
	assert.Equal(t, 64, partTwo(m))
}
