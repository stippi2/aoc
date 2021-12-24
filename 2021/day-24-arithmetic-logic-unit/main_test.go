package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = `inp z
inp x
mul z 3
eql z x`

func Test_parseInput(t *testing.T) {
	alu, program := parseInput(example)
	alu.input = []int{1,3}
	for _, instruction := range program {
		instruction.Execute()
	}
	assert.Equal(t, 1, alu.z)
}
