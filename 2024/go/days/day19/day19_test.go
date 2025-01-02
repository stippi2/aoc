package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

func Test_part1(t *testing.T) {
	assert.Equal(t, 6, countPossibleDesigns(parseInput(example)))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 16, countPossibleArrangements(parseInput(example)))
}
