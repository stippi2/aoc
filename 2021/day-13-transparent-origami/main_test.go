package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func Test_parseInput(t *testing.T) {
	o := parseInput(exampleInput)
	if assert.Len(t, o.dots, 18) {
		assert.Equal(t, o.dots[0], Point{6, 10})
		assert.Equal(t, o.dots[17], Point{9, 0})
	}
	if assert.Len(t, o.foldsY, 1) {
		assert.Equal(t, o.foldsY[0], 7)
	}
	if assert.Len(t, o.foldsX, 1) {
		assert.Equal(t, o.foldsX[0], 5)
	}
}
