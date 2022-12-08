package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `30373
25512
65332
33549
35390`

func Test_Part1(t *testing.T) {
	m := parseInput(exampleInput)
	assert.Equal(t, 5, m.width)
	assert.Equal(t, 5, m.height)
	m.setVisibilityLeftRight()
	m.setVisibilityTopBottom()
	assert.Equal(t, 21, m.countVisibleTrees())
}

func Test_Part2(t *testing.T) {
	m := parseInput(exampleInput)
	assert.Equal(t, 8, m.computeScenicScores())
}
