package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

func Test_parseInput(t *testing.T) {
	assert.True(t, true)
}
