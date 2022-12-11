package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_add(t *testing.T) {
	n := toRemainderNumber(45)
	n.add(4)
	assert.True(t, n.isDivisibleBy(7))
}

func Test_trackRemainderIdea(t *testing.T) {
	v := 92
	r := v % 7
	assert.Equal(t, (v+5)%7, (r+5)%7)
	assert.Equal(t, (v*3)%7, (r*3)%7)
	assert.Equal(t, (v*2)%7, (r*2)%7)
	assert.Equal(t, (v*v)%7, (r*v)%7)
	assert.Equal(t, (v*v)%7, (r*r)%7)
}
