package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Sequence struct {
	values []int
}

func (s *Sequence) lastValue() int {
	return s.values[len(s.values)-1]
}

func (s *Sequence) append(value int) {
	s.values = append(s.values, value)
}

func (s *Sequence) firstValue() int {
	return s.values[0]
}
func (s *Sequence) prepend(value int) {
	s.values = append([]int{value}, s.values...)
}

func (s *Sequence) isAllZeros() bool {
	for _, v := range s.values {
		if v != 0 {
			return false
		}
	}
	return true
}

func (s *Sequence) getDiffs() Sequence {
	diffs := Sequence{}
	for i := 1; i < len(s.values); i++ {
		diffs.append(s.values[i] - s.values[i-1])
	}
	return diffs
}

func (s *Sequence) getDiffsStack() []*Sequence {
	diffsStack := []*Sequence{s}
	for {
		diffs := diffsStack[len(diffsStack)-1].getDiffs()
		diffsStack = append(diffsStack, &diffs)
		if diffs.isAllZeros() {
			break
		}
	}
	return diffsStack
}

func (s *Sequence) extendRight() int {
	diffsStack := s.getDiffsStack()
	diffsStack[len(diffsStack)-1].append(0)
	for i := len(diffsStack) - 2; i >= 0; i-- {
		diffsStack[i].append(diffsStack[i].lastValue() + diffsStack[i+1].lastValue())
	}
	return s.lastValue()
}

func (s *Sequence) extendLeft() int {
	diffsStack := s.getDiffsStack()
	diffsStack[len(diffsStack)-1].prepend(0)
	for i := len(diffsStack) - 2; i >= 0; i-- {
		diffsStack[i].prepend(diffsStack[i].firstValue() - diffsStack[i+1].firstValue())
	}
	return s.firstValue()
}
func partOne(sequences []Sequence) int {
	sum := 0
	for i := range sequences {
		sum += sequences[i].extendRight()
	}
	return sum
}

func partTwo(sequences []Sequence) int {
	sum := 0
	for i := range sequences {
		sum += sequences[i].extendLeft()
	}
	return sum
}

func main() {
	now := time.Now()
	sequences := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(sequences)
	sequences = parseInput(loadInput("puzzle-input.txt"))
	part2 := partTwo(sequences)
	duration := time.Since(now)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseLine(line string) Sequence {
	sequence := Sequence{}
	for _, num := range strings.Split(line, " ") {
		value, _ := strconv.Atoi(num)
		sequence.append(value)
	}
	return sequence
}

func parseInput(input string) []Sequence {
	lines := strings.Split(input, "\n")
	sequences := make([]Sequence, len(lines))
	for i, line := range lines {
		sequences[i] = parseLine(line)
	}
	return sequences
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
