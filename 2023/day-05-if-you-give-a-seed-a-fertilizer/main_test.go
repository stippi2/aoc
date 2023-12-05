package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func Test_partOne(t *testing.T) {
	seeds, conversions := parseInput(input)
	assert.Equal(t, []int{79, 14, 55, 13}, seeds)
	assert.Equal(t, 7, len(conversions))
	toSoil := conversions["seed"]
	assert.Equal(t, 2, len(toSoil.offsets))
	assert.Equal(t, "soil", toSoil.targetName)
	assert.Equal(t, 0, toSoil.convert(0))
	assert.Equal(t, 1, toSoil.convert(1))
	assert.Equal(t, 48, toSoil.convert(48))
	assert.Equal(t, 49, toSoil.convert(49))
	assert.Equal(t, 52, toSoil.convert(50))
	assert.Equal(t, 53, toSoil.convert(51))
	assert.Equal(t, 98, toSoil.convert(96))
	assert.Equal(t, 99, toSoil.convert(97))
	assert.Equal(t, 50, toSoil.convert(98))
	assert.Equal(t, 51, toSoil.convert(99))
	assert.Equal(t, 35, partOne(seeds, conversions))
}

func Test_partTwo(t *testing.T) {
}
