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

func Test_part2(t *testing.T) {
	lines := parseInput(input)
	addPartNumbers(lines)

	sumGearRatios := 0
	for _, gear := range gears {
		if len(gear) == 2 {
			sumGearRatios += gear[0] * gear[1]
		}
	}
	assert.Equal(t, 467835, sumGearRatios)
}
