package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func Test_Part1(t *testing.T) {
	assert.Equal(t, 3, countZeros(example, false))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 6, countZeros(example, true))
}
