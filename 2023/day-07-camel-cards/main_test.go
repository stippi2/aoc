package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

const input = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func Test_partOne(t *testing.T) {
	hands := parseInput(input)
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].compare(hands[j], false) <= 0
	})
	assert.Equal(t, "32T3K", hands[0].hand)
	assert.Equal(t, "KTJJT", hands[1].hand)
	assert.Equal(t, "KK677", hands[2].hand)
	assert.Equal(t, "T55J5", hands[3].hand)
	assert.Equal(t, "QQQJA", hands[4].hand)
	assert.Equal(t, 6440, partOne(hands))
}

func Test_partTwo(t *testing.T) {
	hands := parseInput(input)
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].compare(hands[j], true) <= 0
	})
	assert.Equal(t, "32T3K", hands[0].hand)
	assert.Equal(t, "KK677", hands[1].hand)
	assert.Equal(t, "T55J5", hands[2].hand)
	assert.Equal(t, "QQQJA", hands[3].hand)
	assert.Equal(t, "KTJJT", hands[4].hand)
	assert.Equal(t, 5905, partTwo(hands))
}
