package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func Test_partOne(t *testing.T) {
	m := parseInput(input)
	assert.Equal(t, 2, m.getHeatLoss(0, 0))
	assert.Equal(t, 102, partOne(m))
}

func Test_partTwo(t *testing.T) {
	_ = parseInput(input)
	assert.Equal(t, 0, partTwo())
}
