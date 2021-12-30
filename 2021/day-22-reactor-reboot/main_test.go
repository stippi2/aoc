package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var smallExample = `on x=10..12,y=10..12,z=10..12
on x=11..13,y=11..13,z=11..13
off x=9..11,y=9..11,z=9..11
on x=10..10,y=10..10,z=10..10`

func Test_parseInput(t *testing.T) {
	expected := []Volume{
		{on: true, min: Position{10, 10, 10}, max: Position{12, 12, 12}},
		{on: true, min: Position{11, 11, 11}, max: Position{13, 13, 13}},
		{on: false, min: Position{9, 9, 9}, max: Position{11, 11, 11}},
		{on: true, min: Position{10, 10, 10}, max: Position{10, 10, 10}},
	}
	actual := parseInput(smallExample)
	assert.Equal(t, expected, actual)
}
