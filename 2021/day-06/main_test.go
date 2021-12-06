package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `3,4,3,1,2`

func Test_parseLanternFishAges(t *testing.T) {
	assert.Len(t, exampleInput, 5)
}
