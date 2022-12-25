package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

func Test_part1(t *testing.T) {
	m := parseInput(exampleInput)
	assert.Equal(t, 18, findPathQueue(m, Pos{1, 0}, Pos{m.width - 2, m.height - 1}))
}

func Test_advanceBlizzards(t *testing.T) {
	m := parseInput(exampleInput)
	m.nextMinute()
	assert.Equal(t, `#.######
#.>3.<.#
#<..<<.#
#>2.22.#
#>v..^<#
######.#`, m.String())
	m.nextMinute()
	assert.Equal(t, `#.######
#.2>2..#
#.^22^<#
#.>2.^>#
#.>..<.#
######.#`, m.String())
	assert.True(t, m.emptyPositions[2][Pos{1, 2}])
}
