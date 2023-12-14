package main

import (
	"fmt"
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
		sb.WriteString("\n")
	}
	return sb.String()
}

func Test_partOne(t *testing.T) {
	m := parseInput(input)
	fmt.Printf("Map:\n%s\n", m)
	assert.Equal(t, 136, partOne(m))
}

func Test_partTwo(t *testing.T) {
	_ = parseInput(input)
	assert.Equal(t, 0, partTwo())
}
