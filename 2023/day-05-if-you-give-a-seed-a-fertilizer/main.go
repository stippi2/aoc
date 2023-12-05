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
	offsets []Offset
	reverse []Offset
	source  string
	target  string
}

func findOffset(offsets []Offset, from int) int {
	// Use binary search and return the index of the offset where from >= offset.from && from < (offset.from + offset.length)
	// If no such offset exists, return the index of the first offset where from < offset.from
	// If no such offset exists, return len(offsets)
	// If offsets is empty, return 0
	if offsets == nil {
		return 0
	}
	low := 0
	high := len(offsets) - 1
	for low <= high {
		mid := (low + high) / 2
		offset := offsets[mid]
		if offset.from <= from && from < offset.from+offset.length {
			return mid
		}
		if from < offsets[mid].from {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func addOffset(offsets []Offset, from int, to int, length int) []Offset {
	offset := Offset{from: from, to: to, length: length}
	insertion := findOffset(offsets, from)
	if insertion < len(offsets) {
		return append(offsets[:insertion], append([]Offset{offset}, offsets[insertion:]...)...)
	}
	return append(offsets, offset)
}

func (c *ConversionMap) addOffset(from int, to int, length int) {
	c.offsets = addOffset(c.offsets, from, to, length)
	c.reverse = addOffset(c.reverse, to, from, length)
}

func convert(offsets []Offset, from int) int {
	if offsets == nil || from < offsets[0].from {
		return from
	}
	index := findOffset(offsets, from)
	if index == len(offsets) {
		return from
	}
	offset := offsets[index]
	if offset.from <= from && from < offset.from+offset.length {
		return from + offset.to - offset.from
	}
	return from
}

func (c *ConversionMap) convert(from int) int {
	return convert(c.offsets, from)
}

func (c *ConversionMap) reverseConvert(from int) int {
	return convert(c.reverse, from)
}

func partOne(seeds []int, conversions map[string]*ConversionMap) int {
	values := make([]int, len(seeds))
	copy(values, seeds)
	conversionMap := conversions["seed"]
	for {
		for i, value := range values {
			values[i] = conversionMap.convert(value)
		}
		if conversionMap.target == "location" {
			break
		}
		conversionMap = conversions[conversionMap.target]
	}
	sort.Ints(values)
	return values[0]
}

func isSeed(seeds []int, value int) bool {
	for i := 0; i < len(seeds); i += 2 {
		if value >= seeds[i] && value <= seeds[i]+seeds[i+1] {
			return true
		}
	}
	return false
}

func partTwo(seeds []int, conversions map[string]*ConversionMap) int {
	reverseConversions := make(map[string]*ConversionMap)
	for _, conversion := range conversions {
		reverseConversions[conversion.target] = conversion
	}

	location := 0
	for {
		conversionMap := reverseConversions["location"]
		value := location
		for {
			value = conversionMap.reverseConvert(value)
			if conversionMap.source == "seed" {
				if isSeed(seeds, value) {
					return location
				}
				break
			}
			conversionMap = reverseConversions[conversionMap.source]
		}
		location++
	}
}

func main() {
	seeds, conversions := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("Part 1: Nearest location: %d\n", partOne(seeds, conversions))
	fmt.Printf("Part 2: Nearest location (seed ranges): %d\n", partTwo(seeds, conversions))
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
		conversion := ConversionMap{source: mapping[0], target: mapping[1]}
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
