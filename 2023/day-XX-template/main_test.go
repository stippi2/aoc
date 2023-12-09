package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = ``

func Test_partOne(t *testing.T) {
	_ = parseInput(input)
	assert.Equal(t, 0, partOne())
}

func Test_partTwo(t *testing.T) {
	_ = parseInput(input)
	assert.Equal(t, 0, partTwo())
}
