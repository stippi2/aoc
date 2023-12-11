package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
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

func (m *Map) String() string {
	var sb strings.Builder
	for y := int64(0); y < m.height; y++ {
		for x := int64(0); x < m.width; x++ {
			if m.isEmptySpace(x, y) {
				sb.WriteString(".")
			} else {
				sb.WriteString("#")
			}
		}
		if y < m.height-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func (m *Map) isEmptySpace(x int64, y int64) bool {
	for _, galaxy := range m.galaxies {
		if galaxy.x == x && galaxy.y == y {
			return false
		}
	}
	return true
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
	m.expandSpace(9)
	distancesSum, pairs := m.sumDistances()
	assert.Equal(t, 36, pairs)
	assert.Equal(t, int64(1030), distancesSum)

	m = parseInput(input)
	m.expandSpace(99)
	distancesSum, pairs = m.sumDistances()
	assert.Equal(t, 36, pairs)
	assert.Equal(t, int64(8410), distancesSum)
}
