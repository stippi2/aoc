package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseInput(t *testing.T) {
	n := parseInput(exampleInput)
	assert.Equal(t, "start", n.name)
	if assert.Equal(t, 2, len(n.next)) {
		assert.Equal(t, "A", n.next[0].name)
		assert.Equal(t, "b", n.next[1].name)
	}
}

var exampleInput = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`
