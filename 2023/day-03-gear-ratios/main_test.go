package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func Test_part1(t *testing.T) {
	lines := parseInput(input)
	sumOfPowers := addPartNumbers(lines)
	assert.Equal(t, 4361, sumOfPowers)
}
