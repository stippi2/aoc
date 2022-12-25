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
