package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = []string{
	"D2FE28",
	"38006F45291200",
	"EE00D40C823060",
	"8A004A801A8002F478",
	"620080001611562C8802118E34",
	"C0015000016115A2E0802F182340",
	"A0016C880162017C3686B18A3D4780",
}

func Test_parseInput(t *testing.T) {
	assert.Equal(t, []uint8{255, 15}, parseInput("FF0F").stream.data)
}

func Test_getLiteral(t *testing.T) {
	p := parseInput(exampleInput[0])
	assert.Equal(t, 6, p.getVersion())
	assert.Equal(t, TypeLiteral, p.getType())
	assert.Equal(t, uint64(2021), p.getLiteral())
}

func Test_mostSignificantBits(t *testing.T) {
	assert.Equal(t, uint8(0xe0), mostSignificantBits(3))
}

func Test_leastSignificantBits(t *testing.T) {
	assert.Equal(t, uint8(0x07), leastSignificantBits(3))
}

func Test_readBits(t *testing.T) {
	p := parseInput(exampleInput[0])
	assert.Equal(t, uint64(0x17), p.stream.readAt(6, 5))
}
