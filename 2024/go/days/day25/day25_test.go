package day25

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`

func Test_parseInput(t *testing.T) {
	test := `#####
.####
.####
.####
.#.#.
.#...
.....

.....
#....
#....
#...#
#.#.#
#.###
#####`

	lock := Lock{}
	lock.pins[0] = 0
	lock.pins[1] = 5
	lock.pins[2] = 3
	lock.pins[3] = 4
	lock.pins[4] = 3

	key := Key{}
	key.pins[0] = 5
	key.pins[1] = 0
	key.pins[2] = 2
	key.pins[3] = 1
	key.pins[4] = 3

	locks, keys := parseInput(test)
	assert.Equal(t, lock, *locks[0])
	assert.Equal(t, key, *keys[0])
}

func Test_keyFitsLock(t *testing.T) {
	test := `#####
.####
.####
.####
.#.#.
.#...
.....

.....
.....
.....
#....
#.#..
#.#.#
#####`

	locks, keys := parseInput(test)
	assert.True(t, keyFitsLock(keys[0], locks[0]))
}

func Test_part1(t *testing.T) {
	locks, keys := parseInput(example)
	assert.Equal(t, 3, countFittingLockKeyCombinations(locks, keys))
}
