package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = "123 328  51 64 \n" +
	" 45 64  387 23 \n" +
	"  6 98  215 314\n" +
	"*   +   *   +  "

func Test_Part1(t *testing.T) {
	valueArrays, operations := parseInput(example)
	assert.Equal(t, 4277556, performOperations(valueArrays, operations))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 3263827, computeInputRightToLeft(example))
}
