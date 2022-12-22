package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func Test_part1(t *testing.T) {
	jetSequence := &Sequence{input: exampleInput}
	rockSequence := &Sequence{input: rocks}
	chamber := newChamber()
	assert.Equal(t, 3068, simulateRocks(chamber, jetSequence, rockSequence, 2022))
}

func Test_isRepeatedSequence(t *testing.T) {
	chamber := newChamber()
	assert.Len(t, chamber.rows, 1)
	for i := 0; i < 2; i++ {
		chamber.rows = append(chamber.rows, Row{columns: []bool{true, false, false, false, false, false, false}})
		chamber.rows = append(chamber.rows, Row{columns: []bool{false, true, false, false, false, false, false}})
		chamber.rows = append(chamber.rows, Row{columns: []bool{false, false, true, false, false, false, false}})
	}
	assert.True(t, chamber.isRepeatedSequence())
}

func Test_isRepeatedHeightChangeSequence(t *testing.T) {
	assert.True(t, isRepeatedHeightChangeSequence([]int{1, 2, 3, 1, 2, 3}))
	assert.False(t, isRepeatedHeightChangeSequence([]int{1, 2, 3, 1, 2}))
	assert.False(t, isRepeatedHeightChangeSequence([]int{1, 2, 3, 1, 2, 3, 4}))
	assert.False(t, isRepeatedHeightChangeSequence([]int{1, 2, 3, 1, 2, 4}))
}
