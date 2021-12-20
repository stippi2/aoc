package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = `--- scanner 0 ---
-1,-1,1
-2,-2,2
-3,-3,3
-2,-3,1
5,6,-4
8,0,7

--- scanner 1 ---
1,-1,1
2,-2,2
3,-3,3
2,-1,3
-5,4,-6
-8,-7,0`

var expected = []Scanner {
	{
		beacons: []Beacon{
			{position: Position{-1,-1,1}},
			{position: Position{-2,-2,2}},
			{position: Position{-3,-3,3}},
			{position: Position{-2,-3,1}},
			{position: Position{5,6,-4}},
			{position: Position{8,0,7}},
		},
	},
	{
		beacons: []Beacon{
			{position: Position{1,-1,1}},
			{position: Position{2,-2,2}},
			{position: Position{3,-3,3}},
			{position: Position{2,-1,3}},
			{position: Position{-5,4,-6}},
			{position: Position{-8,-7,0}},
		},
	},
}

func Test_parseInput(t *testing.T) {
	assert.Equal(t, expected, parseInput(example))
}
