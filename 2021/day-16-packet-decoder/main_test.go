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
	value, length := p.getLiteral()
	assert.Equal(t, uint64(2021), value)
	assert.Equal(t, uint64(15), length)
}

func Test_mostSignificantBits(t *testing.T) {
	assert.Equal(t, uint8(0xe0), mostSignificantBits(3))
}

func Test_readBits(t *testing.T) {
	p := parseInput(exampleInput[0])
	assert.Equal(t, uint64(0x17), p.stream.readAt(6, 5))
}

func TestVersionAddingVisitor(t *testing.T) {
	tests := []struct {
		input string
		expectedVersionSum int
	}{
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340",23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}
	for _, test := range tests {
		v := &VersionAddingVisitor{}
		p := parseInput(test.input)
		p.visit(v)
		assert.Equal(t, test.expectedVersionSum, v.versionSum)
	}
}

func TestCalculatingVisitor(t *testing.T) {
	tests := []struct {
		input string
		expectedValue int
	}{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112",7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	}
	for _, test := range tests {
		c := &CalculatingVisitor{}
		p := parseInput(test.input)
		p.visit(c)
		assert.Equal(t, test.expectedValue, c.finalValue)
	}
}
