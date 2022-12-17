package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func Test_part1(t *testing.T) {
	jetSequence := &Sequence{input: exampleInput}
	assert.Equal(t, 3068, simulateRocks(jetSequence, 2022))
}
