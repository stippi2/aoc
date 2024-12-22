package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `1
10
100
2024`

func Test_part1(t *testing.T) {
	buyers := parseInput(example)
	assert.Equal(t, 37327623, sum2000thSecretNumbers(buyers))
}
