package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

//  aaaa
// b    c
// b    c
//  dddd
// e    f
// e    f
//  gggg
//
// 0: abc.efg  6
// 1: ..c..f.  2 *
// 2: a.cde.g  5
// 3: a.cd.fg  5
// 4: .bcd.f.  4 *
// 5: ab.d.fg  5
// 6: ab.defg  6
// 7: a.c..f.  3 *
// 8: abcdefg  7 *
// 9: abcd.fg  6
//
// 7 - 1 =     a...... -> a
// 4 - 7 =     .b.d...

// Number of occurrences of elements across all digits:
// a = 8 *
// b = 6 *
// c = 8 *
// d = 7
// e = 4 *
// f = 9 *
// g = 7

var unscrambledDisplay = Display{
	readings: []string{
		"abcefg",  // 0
		"cf",      // 1
		"acdeg",   // 2
		"acdfg",   // 3
		"bcdf",    // 4
		"abdfg",   // 5
		"abdefg",  // 6
		"acf",     // 7
		"abcdefg", // 8
		"abcdfg",  // 9
	},
	digits:  []string{"bcdf", "acdeg"},
	mapping: make(map[string]string),
}

func Test_deduceMapping(t *testing.T) {
	d := unscrambledDisplay
	d.deduceMapping()
	assert.Equal(t, "a", d.mapping["a"])
	assert.Equal(t, "b", d.mapping["b"])
	assert.Equal(t, "c", d.mapping["c"])
	assert.Equal(t, "d", d.mapping["d"])
	assert.Equal(t, "e", d.mapping["e"])
	assert.Equal(t, "f", d.mapping["f"])
	assert.Equal(t, "g", d.mapping["g"])
}

func Test_deduceScrambledExample(t *testing.T) {
	displays := parseInput("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	d := displays[0]
	d.deduceMapping()
	assert.Equal(t, "d", d.mapping["a"])
	assert.Equal(t, "e", d.mapping["b"])
	assert.Equal(t, "a", d.mapping["c"])
	assert.Equal(t, "f", d.mapping["d"])
	assert.Equal(t, "g", d.mapping["e"])
	assert.Equal(t, "b", d.mapping["f"])
	assert.Equal(t, "c", d.mapping["g"])

	assert.Equal(t, 5353, d.descramble())
}

func Test_parseInput(t *testing.T) {
	displays := parseInput(exampleInput)
	assert.Len(t, displays, 10)
	assert.Equal(t, []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"}, displays[0].readings)
	assert.Equal(t, []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}, displays[0].digits)
}

func Test_countDigits(t *testing.T) {
	displays := parseInput(exampleInput)
	count := countDigits(displays, conditionOnesFoursSevensAndEights)
	assert.Equal(t, 26, count)
}

func Test_maskDigits(t *testing.T) {
	assert.Equal(t, "b", maskDigit("abcd", "cda"))
}
