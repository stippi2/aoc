package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var exampleInput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func Test_sumContents(t *testing.T) {
	rucksackContents := parseInput(exampleInput)
	sum := sumContents(rucksackContents)
	assert.Equal(t, 157, sum)
}

func Test_priorityOfDuplicateItem(t *testing.T) {
	assert.Equal(t, 16, priorityOfDuplicateItem(strings.Split("vJrwpWtwJgWrhcsFMMfFFhFp", "")))
	assert.Equal(t, 48, priorityOfDuplicateItem(strings.Split("ttffrVJWtWpgtQnZGVnNSLTHSZ", "")))
}

func Test_sumBadge(t *testing.T) {
	rucksackContents := parseInput(exampleInput)
	sum := sumBadges(rucksackContents)
	assert.Equal(t, 70, sum)
}

func Test_itemPriority(t *testing.T) {
	assert.Equal(t, priority("c"), 3)
	assert.Equal(t, priority("a"), 1)
	assert.Equal(t, priority("z"), 26)
	assert.Equal(t, priority("A"), 27)
	assert.Equal(t, priority("Z"), 52)
}
