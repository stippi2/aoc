package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func Test_partOne(t *testing.T) {
	sequence := parseInput(input)
	assert.Equal(t, 1320, partOne(sequence))
}

func Test_partTwo(t *testing.T) {
	sequence := parseInput(input)
	assert.Equal(t, 145, partTwo(sequence))
}
