package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

func Test_part1(t *testing.T) {
	computer := parseInput(example)
	computer.runProgram()
	assert.Equal(t, "4,6,3,5,6,3,5,2,1,0", computer.output)
}
