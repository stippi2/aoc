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

func Test_applyRules(t *testing.T) {
	p := parseInput(exampleInput)
	expected := []string{
		"NCNBCHB",
		"NBCCNBBBCBHCB",
		"NBBBCNCCNBBNBNBBCHBHHBCHB",
		"NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB",
	}
	for step := 0; step < len(expected); step++ {
		p.applyRules()
		assert.Equal(t, expected[step], p.polymer)
	}
}

func Test_minMax(t *testing.T) {
	p := parseInput(exampleInput)
	for step := 0; step < 10; step++ {
		p.applyRules()
	}
	min, max := occurrences(p.getElementCounts())
	assert.Equal(t, 1749, max)
	assert.Equal(t, 161, min)
}
