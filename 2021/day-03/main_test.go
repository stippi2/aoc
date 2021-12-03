package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func Test_getGammaAndEpsilon(t *testing.T) {
	gamma, epsilon := getGammaAndEpsilon(exampleInput, 5)
	assert.Equal(t, "10110", gamma)
	assert.Equal(t, "01001", epsilon)
}

func Test_toDecimal(t *testing.T) {
	assert.Equal(t, 22, toDecimal("10110"))
	assert.Equal(t, 9, toDecimal("01001"))
}

func Test_getOxygenAndCo2Scrubber(t *testing.T) {
	oxygen, co2scrubber := getOxygenAndCo2Scrubber(exampleInput)
	assert.Equal(t, "10111", oxygen)
	assert.Equal(t, "01010", co2scrubber)
}
