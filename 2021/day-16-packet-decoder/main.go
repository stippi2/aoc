package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const TypeLiteral = 4

type BitStream struct {
	data      []uint8
	bitOffset int64
}

func mostSignificantBits(bits int) uint8 {
	bit := uint8(0x80)
	mask := uint8(0)
	for i := 0; i < bits; i++ {
		mask |= bit
		bit >>= 1
	}
	return mask
}

func leastSignificantBits(bits int) uint8 {
	bit := uint8(0x01)
	mask := uint8(0)
	for i := 0; i < bits; i++ {
		mask |= bit
		bit <<= 1
	}
	return mask
}

func (s *BitStream) readAt(offset, bitCount int64) uint64 {
	s.bitOffset = offset
	return s.read(bitCount)
}

func (s *BitStream) read(bitCount int64) uint64 {
	dataOffset := s.bitOffset / 8
	skipBits := int(s.bitOffset - dataOffset * 8)
	remainder := int64(8 - skipBits) - bitCount
	mask := mostSignificantBits(skipBits)
	var value uint64
	value = uint64(s.data[dataOffset] & (^mask))
	for remainder < 0 {
		value <<= 8
		dataOffset++
		value |= uint64(s.data[dataOffset])
		remainder += 8
	}
	value >>= remainder
	s.bitOffset += bitCount
	return value
}

type Packet struct {
	stream BitStream
	bitOffset int64
}

func (p *Packet) getVersion() int {
	return int(p.stream.readAt(p.bitOffset, 3))
}

func (p *Packet) getType() int {
	return int(p.stream.readAt(p.bitOffset + 3, 3))
}

func (p *Packet) getLiteral() (value uint64, length int64) {
	p.stream.bitOffset = p.bitOffset + 6
	for {
		bits := p.stream.read(5)
		value |= bits & 0xf
		length += 5
		if (bits & 0x10) == 0 {
			break
		}
		value <<= 4
	}
	return
}

func (p *Packet) getLength() int64 {
	_, length := p.getLiteral()
	return length
}

func main() {
}

func parseInput(input string) Packet {
	data := make([]uint8, len(input) / 2)
	for i := 0; i < len(input); i += 2 {
		v, err := strconv.ParseUint(input[i:i+2], 16, 8)
		if err != nil {
			panic(fmt.Sprintf("failed to parse hex (%s) at %d: %s", input[i:i+2], i, err))
		}
		data[i / 2] = uint8(v)
	}
	return Packet{stream: BitStream{data: data}}
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
