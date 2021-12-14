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
	p := parseInput(exampleInput)
	assert.Equal(t, "NNCB", p.polymer)
	assert.Len(t, p.insertionRules, 16)
	assert.Equal(t, p.insertionRules["HC"], "B")
}
