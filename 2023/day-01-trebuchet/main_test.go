package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_extractNumbers(t *testing.T) {
	numbers := parseInput(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`)
	assert.Equal(t, []int{12, 38, 15, 77}, numbers)
	assert.Equal(t, 142, sum(numbers))
}
