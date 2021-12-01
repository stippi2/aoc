package main

import "testing"
import "github.com/stretchr/testify/assert"

func Test_sumWindow(t *testing.T) {
	assert.Equal(t, 3, sumWindow([]int{0, 1, 2}, 1, 3))
	assert.Equal(t, 1, sumWindow([]int{0, 1, 2}, 1, 2))
	assert.Equal(t, 1, sumWindow([]int{3, 1, 2}, 1, 1))
	assert.Equal(t, 0, sumWindow([]int{}, 1, 1))
	assert.Equal(t, 1, sumWindow([]int{1, 2, 4}, 0, 1))
	assert.Equal(t, 4, sumWindow([]int{1, 2, 4}, 2, 1))
	assert.Equal(t, 3, sumWindow([]int{1, 2, 4}, 1, 2))
	assert.Equal(t, 0, sumWindow([]int{1, 2, 4}, 0, 3))
	assert.Equal(t, 7, sumWindow([]int{1, 2, 4}, 1, 3))
	assert.Equal(t, 0, sumWindow([]int{1, 2, 4}, 2, 3))
}
