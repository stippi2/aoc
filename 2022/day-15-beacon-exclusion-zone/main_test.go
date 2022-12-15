package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

func Test_parseInput(t *testing.T) {
	sensors := parseInput(exampleInput)
	assert.Len(t, sensors, 14)
	assert.Equal(t, Sensor{
		pos:    Pos{2, 18},
		beacon: Pos{-2, 15},
	}, sensors[0])
	assert.Equal(t, Sensor{
		pos:    Pos{9, 16},
		beacon: Pos{10, 16},
	}, sensors[1])
}

func Test_minMaxX(t *testing.T) {
	sensors := parseInput(exampleInput)
	line, insideRange := sensors[6].minMaxX(9)
	assert.True(t, insideRange)
	assert.Equal(t, Line{1, 16}, *line)
}

func Test_part1(t *testing.T) {
	sensors := parseInput(exampleInput)
	assert.Equal(t, 26, emptyPositionsOnLine(10, sensors))
}
