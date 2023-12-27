package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

func Test_partOne(t *testing.T) {
	slabs := parseInput(input)
	result := partOne(slabs)
	assert.Equal(t, 6, slabs[len(slabs)-1].z2)
	assert.Equal(t, 5, result)
}

func Test_partTwo(t *testing.T) {
	_ = parseInput(input)
	assert.Equal(t, 0, partTwo())
}
