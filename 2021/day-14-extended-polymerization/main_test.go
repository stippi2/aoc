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
		temp := PolymerProcess{}
		temp.init(expected[step])
		assert.Equal(t, temp.elementCounts, p.elementCounts)
		if !assert.Equal(t, temp.combinations, p.combinations) {
			break
		}
	}
}

func Test_minMax(t *testing.T) {
	p := parseInput(exampleInput)
	for step := 0; step < 10; step++ {
		p.applyRules()
	}
	min, max, _, _ := occurrences(p.getElementCounts())
	assert.Equal(t, int64(1749), max)
	assert.Equal(t, int64(161), min)
}

/*
Template:     NNCB
After step 2: NBCCNBBBCBHCB
After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB
After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB

NN 1 -> NN--, NC++, CN++
NC 1 -> NC--, NB++, BC++
CB 1 -> CB--, CH++, HB++

-> NCNBCHB

NC 1
CN 1
NB 1
BC 1
CH 1
HB 1


*/

