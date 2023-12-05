package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Offset struct {
	from   int
	to     int
	length int
}

type ConversionMap struct {
	offsets    []Offset
	targetName string
}

func (c *ConversionMap) findOffset(from int) int {
	// Use binary search and return the index of the offset where from >= offset.from && from < (offset.from + offset.length)
	// If no such offset exists, return the index of the first offset where from < offset.from
	// If no such offset exists, return len(offsets)
	// If offsets is empty, return 0
	if c.offsets == nil {
		return 0
	}
	low := 0
	high := len(c.offsets) - 1
	for low <= high {
		mid := (low + high) / 2
		offset := c.offsets[mid]
		if offset.from <= from && from < offset.from+offset.length {
			return mid
		}
		if from < c.offsets[mid].from {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func (c *ConversionMap) addOffset(from int, to int, length int) {
	offset := Offset{from: from, to: to, length: length}
	insertion := c.findOffset(from)
	if insertion < len(c.offsets) {
		c.offsets = append(c.offsets[:insertion], append([]Offset{offset}, c.offsets[insertion:]...)...)
		return
	}
	c.offsets = append(c.offsets, offset)
}

func (c *ConversionMap) convert(from int) int {
	if c.offsets == nil || from < c.offsets[0].from {
		return from
	}
	index := c.findOffset(from)
	if index == len(c.offsets) {
		return from
	}
	offset := c.offsets[index]
	if offset.from <= from && from < offset.from+offset.length {
		return from + offset.to - offset.from
	}
	return from
}

func partOne(seeds []int, conversions map[string]*ConversionMap) int {
	values := make([]int, len(seeds))
	copy(values, seeds)
	conversionMap := conversions["seed"]
	for {
		for i, value := range values {
			values[i] = conversionMap.convert(value)
		}
		if conversionMap.targetName == "location" {
			break
		}
		conversionMap = conversions[conversionMap.targetName]
	}
	sort.Ints(values)
	return values[0]
}

func partTwo() int {
	return 0
}

func main() {
	seeds, conversions := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("Part 1: Nearest location: %d\n", partOne(seeds, conversions))
}

func parseInput(input string) ([]int, map[string]*ConversionMap) {
	conversions := make(map[string]*ConversionMap)
	var seeds []int

	input = strings.TrimSpace(input)
	sections := strings.Split(input, "\n\n")

	seedsString := strings.TrimPrefix(sections[0], "seeds: ")
	for _, seedString := range strings.Split(seedsString, " ") {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}

	for _, section := range sections[1:] {
		lines := strings.Split(section, "\n")
		mapping := strings.Split(strings.TrimSuffix(lines[0], " map:"), "-to-")
		conversion := ConversionMap{targetName: mapping[1]}
		conversions[mapping[0]] = &conversion
		for _, line := range lines[1:] {
			var to, from, length int
			matches, err := fmt.Sscanf(line, "%d %d %d", &to, &from, &length)
			if err != nil || matches != 3 {
				return nil, nil
			}
			conversion.addOffset(from, to, length)
		}
	}

	return seeds, conversions
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
