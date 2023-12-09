package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func Test_partOne(t *testing.T) {
	sequences := parseInput(input)
	assert.Equal(t, 114, partOne(sequences))
}

func Test_partTwo(t *testing.T) {
	sequences := parseInput(input)
	assert.Equal(t, 2, partTwo(sequences))
}
