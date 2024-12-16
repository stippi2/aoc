package day16

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

func Test_part1(t *testing.T) {
	lowestScore, _ := findLowestScore(example, math.MaxInt)
	assert.Equal(t, 7036, lowestScore)
}

func Test_part2(t *testing.T) {
	_, bestPositions := findLowestScore(example, math.MaxInt)
	assert.Equal(t, 45, bestPositions)
}
