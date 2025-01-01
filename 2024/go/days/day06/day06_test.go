package day06

import (
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
	assert.Equal(t, 41, sumVisitedFields(example))
}

func Test_Part2(t *testing.T) {
}
