package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `029A
980A
179A
456A
379A`

func Test_part1(t *testing.T) {
	assert.Equal(t, 126384, calculateComplexity(example, 2))
}
