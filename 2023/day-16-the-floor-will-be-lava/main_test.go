package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func Test_partOne(t *testing.T) {
	m := parseInput(input)
	assert.Equal(t, 46, partOne(m))
}

func Test_partTwo(t *testing.T) {
	m := parseInput(input)
	assert.Equal(t, 51, partTwo(m))
}
