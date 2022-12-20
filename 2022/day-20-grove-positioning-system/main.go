package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	value          int
	index          int
	previous, next *Number
}

type Sequence struct {
	numbers     []*Number
	first, last *Number
	length      int
}

func (s *Sequence) append(n *Number) {
	if s.first == nil {
		s.first = n
		s.last = n
	}
	s.last.next = n
	s.first.previous = n
	n.previous = s.last
	n.next = s.first
	s.last = n
	s.length++
}

func (s *Sequence) prepend(n *Number) {
	if s.first == nil {
		s.first = n
		s.last = n
	}
	s.first.previous = n
	s.last.next = n
	n.next = s.first
	n.previous = s.last
	s.first = n
	s.length++
}

func (s *Sequence) String() string {
	result := ""
	n := s.first
	for {
		if result != "" {
			result += ", "
		}
		result += fmt.Sprintf("[%v, %v]", n.index, n.value)
		if n == s.last {
			break
		}
		n = n.next
	}
	return result
}

func (s *Sequence) remove(n *Number) {
	// Would not work if s contains only "n"
	if n == s.first {
		s.first = n.next
	}
	if n == s.last {
		s.last = n.previous
	}
	n.previous.next = n.next
	n.next.previous = n.previous
	s.length--
}

func (s *Sequence) insertAfter(n, previous *Number) {
	if previous == s.last {
		s.append(n)
		return
	}
	next := previous.next
	previous.next = n
	n.previous = previous
	n.next = next
	next.previous = n
	s.length++
}

func (s *Sequence) insertBefore(n, next *Number) {
	if next == s.first {
		s.prepend(n)
		return
	}
	previous := next.previous
	next.previous = n
	n.previous = previous
	n.next = next
	previous.next = n
	s.length++
}

func (s *Sequence) skip(n *Number, value int) *Number {
	if value < 0 {
		count := (-value) % s.length
		previous := n
		for count >= 0 {
			count--
			previous = previous.previous
		}
		return previous.next
	} else if value > 0 {
		count := value % s.length
		next := n
		for count >= 0 {
			count--
			next = next.next
		}
		return next.previous
	}
	return n
}

func (s *Sequence) mix() {
	for _, current := range s.numbers {
		if current.value < 0 {
			count := (-current.value) % (s.length - 1)
			previous := current
			for count >= 0 {
				count--
				previous = previous.previous
			}
			// Detach current
			s.remove(current)
			// Attach after previous
			s.insertAfter(current, previous)
		} else if current.value > 0 {
			count := current.value % (s.length - 1)
			next := current
			for count >= 0 {
				count--
				next = next.next
			}
			// Detach current
			s.remove(current)
			// Attach after previous
			s.insertBefore(current, next)
		}
	}
}

func (s *Sequence) find(value int) *Number {
	current := s.first
	for {
		if current.value == value {
			return current
		}
		current = current.next
		if current == s.first {
			break
		}
	}
	return nil
}

func (s *Sequence) findGroveCoordinates() int {
	n := s.find(0)
	coord1 := s.skip(n, 1000)
	coord2 := s.skip(n, 2000)
	coord3 := s.skip(n, 3000)
	return coord1.value + coord2.value + coord3.value
}

func (s *Sequence) multiply(v int) {
	for _, n := range s.numbers {
		n.value *= v
	}
}

func main() {
	sequence := parseInput(loadInput("puzzle-input.txt"))
	sequence.mix()
	fmt.Printf("part 1: grove coordinates: %v\n", sequence.findGroveCoordinates())

	sequence = parseInput(loadInput("puzzle-input.txt"))
	sequence.multiply(811589153)
	for i := 0; i < 10; i++ {
		sequence.mix()
	}
	fmt.Printf("part 2: grove coordinates: %v\n", sequence.findGroveCoordinates())
}

func parseInput(input string) *Sequence {
	lines := strings.Split(input, "\n")
	sequence := &Sequence{numbers: make([]*Number, len(lines))}
	for i, line := range lines {
		value, _ := strconv.Atoi(line)
		sequence.numbers[i] = &Number{index: i, value: value}
		sequence.append(sequence.numbers[i])
	}
	return sequence
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
