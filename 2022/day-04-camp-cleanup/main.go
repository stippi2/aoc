package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Sections struct {
	start int
	end   int
}

type Pair struct {
	one Sections
	two Sections
}

func (s *Sections) contains(other Sections) bool {
	return other.start >= s.start && other.end <= s.end
}

func (s *Sections) intersects(other Sections) bool {
	return !(other.start > s.end || other.end < s.start)
}

func sumCompletelyOverlappingPairs(pairs []Pair) int {
	sum := 0
	for _, pair := range pairs {
		if pair.one.contains(pair.two) || pair.two.contains(pair.one) {
			sum++
		}
	}
	return sum
}

func sumOverlappingPairs(pairs []Pair) int {
	sum := 0
	for _, pair := range pairs {
		if pair.one.intersects(pair.two) {
			sum++
		}
	}
	return sum
}

func main() {
	pairs := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("total sum of completely overlapping pairs: %v\n", sumCompletelyOverlappingPairs(pairs))
	fmt.Printf("total sum of overlapping pairs: %v\n", sumOverlappingPairs(pairs))
}

func parseSections(s string) Sections {
	parts := strings.Split(s, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return Sections{start, end}
}

func parseInput(input string) []Pair {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	var pairs []Pair
	for _, line := range lines {
		parts := strings.Split(line, ",")
		pairs = append(pairs, Pair{
			one: parseSections(parts[0]),
			two: parseSections(parts[1]),
		})
	}
	return pairs
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
