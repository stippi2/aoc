package day22

import (
	"fmt"
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

func Test_part2(t *testing.T) {
	buyers := parseInput(`1
2
3
2024`)
	for i := 0; i < len(buyers); i++ {
		buyers[i].generateSequence()
	}
	mostBananas, bestSequence := findBestSequence(buyers)
	fmt.Printf("Most bananas: %v (sequence: %v)", mostBananas, bestSequence)
	assert.Equal(t, 23, mostBananas)
}
