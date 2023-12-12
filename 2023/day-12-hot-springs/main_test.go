package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func Test_partOne(t *testing.T) {
	rows := parseInput(input)
	assert.Equal(t, 6, len(rows))
	assert.Equal(t, []int{1, 1, 3}, rows[0].groups)
	assert.Equal(t, []byte("?#?#?#?#?#?#?#?"), rows[2].data)
}

func Test_partTwo(t *testing.T) {
	_ = parseInput(input)
	assert.Equal(t, 0, partTwo())
}
