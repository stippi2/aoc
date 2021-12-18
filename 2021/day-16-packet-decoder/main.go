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
	bitOffset uint64
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

func (s *BitStream) readAt(offset, bitCount uint64) uint64 {
	s.bitOffset = offset
	return s.read(bitCount)
}

func (s *BitStream) readByte(dataOffset uint64) uint8 {
	if uint64(len(s.data)) <= dataOffset {
		return 0
	}
	return s.data[dataOffset]
}

func (s *BitStream) read(bitCount uint64) uint64 {
	dataOffset := s.bitOffset / 8
	skipBits := int(s.bitOffset - dataOffset * 8)
	remainder := int64(8) - int64(skipBits) - int64(bitCount)
	mask := mostSignificantBits(skipBits)
	var value uint64
	value = uint64(s.readByte(dataOffset) & (^mask))
	for remainder < 0 {
		value <<= 8
		dataOffset++
		value |= uint64(s.readByte(dataOffset))
		remainder += 8
	}
	value >>= remainder
	s.bitOffset += bitCount
	return value
}

type Packet struct {
	stream BitStream
	bitOffset uint64
}

func (p *Packet) readAt(bitOffset, bitCount uint64) uint64 {
	return p.stream.readAt(p.bitOffset + bitOffset, bitCount)
}

func (p *Packet) getVersion() int {
	return int(p.readAt(0, 3))
}

func (p *Packet) getType() int {
	return int(p.readAt(3, 3))
}

func (p *Packet) getLiteral() (value, length uint64) {
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

type PacketVisitor interface {
	Visit(p *Packet)
}

func (p *Packet) visit(visitor PacketVisitor) uint64 {
	visitor.Visit(p)
	t := p.getType()
	if t == TypeLiteral {
		_, length := p.getLiteral()
		return 6 + length
	} else {
		lengthType := p.readAt(6, 1)
		offset := uint64(7)
		if lengthType == 0 {
			remaining := p.readAt(offset, 15)
			offset += 15
			for remaining > 0 {
				packet := Packet{
					stream:    p.stream,
					bitOffset: p.bitOffset + offset,
				}
				length := packet.visit(visitor)
				remaining -= length
				offset += length
			}
		} else {
			packetCount := p.readAt(offset, 11)
			offset += 11
			for packetCount > 0 {
				packet := Packet{
					stream:    p.stream,
					bitOffset: p.bitOffset + offset,
				}
				offset += packet.visit(visitor)
				packetCount--
			}
		}
		return offset
	}
}

type VersionAddingVisitor struct {
	versionSum int
}

func (v *VersionAddingVisitor) Visit(p *Packet) {
	v.versionSum += p.getVersion()
}


func main() {
	v := &VersionAddingVisitor{}
	p := parseInput(loadInput("puzzle-input.txt"))
	p.visit(v)
	fmt.Printf("sum of packet versions: %v\n", v.versionSum)
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
