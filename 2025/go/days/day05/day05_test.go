package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func Test_Part1(t *testing.T) {
	assert.Equal(t, 3, countFreshIds(example))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 14, countTotalFreshIds(example))
}
