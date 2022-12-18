package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func Test_part1(t *testing.T) {
	droplet := parseInput(exampleInput)
	assert.Equal(t, 64, droplet.surfaceArea())
}

func Test_part2(t *testing.T) {
	droplet := parseInput(exampleInput)
	droplet.fillAllPockets()
	assert.Equal(t, 58, droplet.surfaceArea())
}
