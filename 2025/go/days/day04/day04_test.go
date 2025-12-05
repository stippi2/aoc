package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func Test_Part1(t *testing.T) {
	assert.Equal(t, 13, countAccessiblePaperRolls(example))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 43, countRemovablePaperRolls(example))
}
