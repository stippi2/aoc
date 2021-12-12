package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseInput(t *testing.T) {
	n := parseInput(exampleInput)
	assert.Equal(t, "start", n.name)
}

var exampleInput = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`
