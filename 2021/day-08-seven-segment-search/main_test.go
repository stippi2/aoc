package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `16,1,2,0,4,2,7,1,2,14`

func Test_parseSequence(t *testing.T) {
	assert.Equal(t, []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, parseSequence(exampleInput))
}
