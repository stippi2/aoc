package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const TypeSum = 0
const TypeProduct = 1
const TypeMinimum = 2
const TypeMaximum = 3
const TypeLiteral = 4
const TypeGreater = 5
const TypeLess = 6
const TypeEqual = 7

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

func (s *BitStream) readAt(offset, bitCount uint64) uint64 {
	s.bitOffset = offset
	return s.read(bitCount)
}

func (s *BitStream) readByte(dataOffset uint64) uint8 {
	if uint64(len(s.data)) <= dataOffset {
		panic("reading out of bounds from bitstream")
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
	Enter(p *Packet)
	Leave()
}

func (p *Packet) visit(visitor PacketVisitor) uint64 {
	visitor.Enter(p)
	defer visitor.Leave()
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

func (v *VersionAddingVisitor) Enter(p *Packet) {
	v.versionSum += p.getVersion()
}

func (v *VersionAddingVisitor) Leave() {
}

type Operation struct {
	op       int
	operands []uint64
}

func (o *Operation) Evaluate() uint64 {
	switch o.op {
	case TypeSum:
		v := uint64(0)
		for _, operand := range o.operands {
			v += operand
		}
		return v
	case TypeProduct:
		v := uint64(1)
		for _, operand := range o.operands {
			v *= operand
		}
		return v
	case TypeMinimum:
		min := uint64(math.MaxUint64)
		for _, operand := range o.operands {
			v := operand
			if v < min {
				min = v
			}
		}
		return min
	case TypeMaximum:
		max := uint64(0)
		for _, operand := range o.operands {
			v := operand
			if v > max {
				max = v
			}
		}
		return max
	case TypeGreater:
		if o.operands[0] > o.operands[1] {
			return 1
		}
		return 0
	case TypeLess:
		if o.operands[0] < o.operands[1] {
			return 1
		}
		return 0
	case TypeEqual:
		if o.operands[0] == o.operands[1] {
			return 1
		}
		return 0
	case TypeLiteral:
		return o.operands[0]
	}
	panic(fmt.Sprintf("unknown operation type %v", o.op))
}

type CalculatingVisitor struct {
	operationStack []*Operation
	finalValue uint64
}

func (v *CalculatingVisitor) Enter(p *Packet) {
	operation := &Operation{op: p.getType()}
	v.operationStack = append(v.operationStack, operation)

	if p.getType() == TypeLiteral {
		value, _ := p.getLiteral()
		operation.operands = append(operation.operands, value)
	}
}

func (v *CalculatingVisitor) Leave() {
	last := len(v.operationStack) - 1
	completedOperation := v.operationStack[last]
	value := completedOperation.Evaluate()

	if last > 0 {
		previous := v.operationStack[last - 1]
		previous.operands = append(previous.operands, value)
		v.operationStack = v.operationStack[:last]
	} else {
		v.finalValue = value
	}
}

func main() {
	p := parseInput(loadInput("puzzle-input.txt"))

	v := &VersionAddingVisitor{}
	p.visit(v)
	fmt.Printf("part 1, sum of packet versions: %v\n", v.versionSum)

	c := &CalculatingVisitor{}
	p.visit(c)
	fmt.Printf("part 2, evaluation: %v\n", c.finalValue)
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
