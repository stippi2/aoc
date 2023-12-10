package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

func Test_partOne(t *testing.T) {
	m := parseInput(input)
	assert.Equal(t, 8, partOne(m))
}

func Test_partTwo(t *testing.T) {
	_ = parseInput(input)
	assert.Equal(t, 0, partTwo())
}
