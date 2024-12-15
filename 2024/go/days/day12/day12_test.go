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
EEEC`))

	assert.Equal(t, 1930, calculateRegionPrice(example))
}
