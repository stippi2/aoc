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

//".??..??...?##. 1,1,3"
//"..?..??...?##. 1,1,3" -> "?.??.?##"
//".?...??...?##. 1,1,3" -> "?.??.?##"

func Test_partOne(t *testing.T) {
	rows := parseInput(input)
	assert.Equal(t, 6, len(rows))
	assert.Equal(t, []int{1, 1, 3}, rows[0].groups)
	assert.Equal(t, []byte("?#?#?#?#?#?#?#?"), rows[2].springs)
	assert.Equal(t, 1, rows[0].findSolutions())
	assert.Equal(t, 4, rows[1].findSolutions())
	assert.Equal(t, 1, rows[2].findSolutions())
	assert.Equal(t, 1, rows[3].findSolutions())
	assert.Equal(t, 4, rows[4].findSolutions())
	assert.Equal(t, 10, rows[5].findSolutions())
	assert.Equal(t, 21, partOne(rows))
	//solutions := 0
	//for key, value := range rows[5].solutions {
	//	solutions += value
	//	fmt.Printf("solution: %s -> %d\n", key, value)
	//}
	//fmt.Printf("solutions total: %d\n", solutions)
}

func Test_partTwo(t *testing.T) {
	rows := parseInput(input)
	for _, row := range rows {
		row.unfold()
	}
	assert.Equal(t, []byte("???.###????.###????.###????.###????.###"), rows[0].springs)
	assert.Equal(t, []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3}, rows[0].groups)
	assert.Equal(t, 1, rows[0].findSolutions())
	assert.Equal(t, 16, rows[3].findSolutions())
	assert.Equal(t, 525152, findSolutions(rows))
}
