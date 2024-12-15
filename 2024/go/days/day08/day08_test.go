package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func Test_countAntiNodes(t *testing.T) {
	assert.Equal(t, 14, countAntiNodes(example, false))
}

func Test_countAntiNodesRepeating(t *testing.T) {
	assert.Equal(t, 34, countAntiNodes(example, true))
}