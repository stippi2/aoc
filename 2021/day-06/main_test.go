package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `3,4,3,1,2`
var agesAfter18Days = `6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8`

// 0 = 3
// 1 = 5
// 2 = 3
// 3 = 2
// 4 = 2
// 5 = 1
// 6 = 5
// 7 = 1
// 8 = 4

func Test_parseLanternFishAges(t *testing.T) {
	assert.Equal(t, []int{3, 4, 3, 1, 2}, parseLanternFishAges(exampleInput))
}

func Test_initAgeCounts(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1, 2, 1, 0, 0, 0, 0}, initAgeCounts(parseLanternFishAges(exampleInput)))
}

func Test_simulateAgingAndReproduction(t *testing.T) {
	countsPerAge := initAgeCounts(parseLanternFishAges(exampleInput))
	for i := 0; i < 18; i++ {
		simulateAgingAndReproduction(countsPerAge)
	}
	assert.Equal(t, initAgeCounts(parseLanternFishAges(agesAfter18Days)), countsPerAge)
	assert.Equal(t, int64(26), countLanternFish(countsPerAge))
}
