package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func Test_part1(t *testing.T) {
	assert.Equal(t, 36, sumTrailHeads(example))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 81, sumTrailHeads(example))
}
