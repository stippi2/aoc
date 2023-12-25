package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

func Test_partOne(t *testing.T) {
	hailstones := parseInput(input)
	assert.Equal(t, 2, partOne(hailstones, 7, 27))
}

func Test_partTwo(t *testing.T) {
	_ = parseInput(input)
	assert.Equal(t, 0, partTwo())
}
