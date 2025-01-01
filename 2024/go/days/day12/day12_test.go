package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func Test_part1(t *testing.T) {
	assert.Equal(t, 4*10+4*8+4*10+3*8+1*4, calculateRegionPrice(`AAAA
BBCD
BBCC
EEEC`, false))

	assert.Equal(t, 1930, calculateRegionPrice(example, false))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 236, calculateRegionPrice(`EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`, true))

	assert.Equal(t, 368, calculateRegionPrice(`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`, true))

	assert.Equal(t, 1206, calculateRegionPrice(example, true))
}
