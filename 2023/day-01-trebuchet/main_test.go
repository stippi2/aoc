package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_extractNumbersPart1(t *testing.T) {
	numbers := parseInput(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`, tokenToValuePart1)
	assert.Equal(t, []int{12, 38, 15, 77}, numbers)
	assert.Equal(t, 142, sum(numbers))
}

func Test_extractNumbersPart2(t *testing.T) {
	numbers := parseInput(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`, tokenToValuePart2)
	assert.Equal(t, []int{29, 83, 13, 24, 42, 14, 76}, numbers)
	assert.Equal(t, 281, sum(numbers))
}
