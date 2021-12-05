package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func Test_parseVentInput(t *testing.T) {
	lines, maxX, maxY := parseLines(exampleInput)
	assert.Len(t, lines, 10)
	assert.Equal(t, Point{0, 9}, lines[0].points[0])
	assert.Equal(t, Point{8, 2}, lines[9].points[1])
	assert.Equal(t, 9, maxX)
	assert.Equal(t, 9, maxY)
}

func Test_countPoints(t *testing.T) {
	dangerMap := buildDangerMap(parseLines(exampleInput))
	assert.Equal(t, 12, dangerMap.countPoints(2))
}
