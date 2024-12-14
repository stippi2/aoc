package day14

import (
	"aoc/2024/go/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func Test_part1(t *testing.T) {
	assert.Equal(t, 12, computeSafetyFactor(example, 11, 7, 100))
}

func Test_simulate(t *testing.T) {
	r := robot{
		pos: lib.Vec2{X: 2, Y: 4},
		vel: lib.Vec2{X: 2, Y: -3},
	}
	r.simulate(11, 7, 4)
	expected := robot{
		pos: lib.Vec2{X: 10, Y: 6},
		vel: lib.Vec2{X: 2, Y: -3},
	}
	assert.Equal(t, expected, r)
}
