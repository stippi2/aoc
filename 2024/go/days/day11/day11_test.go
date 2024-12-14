package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = "125 17"

func Test_part1(t *testing.T) {
	assert.Equal(t, 55312, countStones(example, 25))
}
