package day07

import (
	"aoc/2024/go/lib"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Calibration struct {
	result   int64
	sequence []int64
}

func solve(result, value int64, sequence []int64) bool {
	if len(sequence) == 0 {
		return result == value
	}
	next := sequence[0]
	if value+next < value || value*next < value {
		log.Fatal("Integer overflow")
	}
	return solve(result, value+next, sequence[1:]) || solve(result, value*next, sequence[1:])
}

func (c *Calibration) validate() bool {
	var product int64 = 1
	for _, value := range c.sequence {
		product *= value
	}
	if product < c.result {
		return false
	}

	var sum int64 = 0
	for _, value := range c.sequence {
		sum += value
	}
	if sum > c.result {
		return false
	}

	return solve(c.result, c.sequence[0], c.sequence[1:])
}

func sumValidCalibrations(inputLines []string) int64 {
	var sum int64 = 0
	calibrations := parseInput(inputLines)
	for _, c := range calibrations {
		if c.validate() {
			oldSum := sum
			sum += c.result
			if oldSum > sum {
				log.Fatal("Integer overflow")
			}
		}
	}
	return sum
}

func Part1() interface{} {
	inputLines, err := lib.ReadInputLines(7)
	if err != nil {
		log.Fatalf("Failed to read input from day 7")
	}
	sum := sumValidCalibrations(inputLines)
	return fmt.Sprintf("Sum of valid calibrations: %v", sum)
}

func Part2() interface{} {
	return "Not implemented"
}

func parseInput(lines []string) []Calibration {
	var calibrations []Calibration = nil
	for i, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			log.Fatalf("Expected two parts in line %d", i)
		}
		result, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse %s as int64", parts[0])
		}
		var sequence []int64 = nil
		for _, number := range strings.Split(parts[1], " ") {
			value, _ := strconv.ParseInt(number, 10, 64)
			sequence = append(sequence, value)
		}
		calibrations = append(calibrations, Calibration{result, sequence})
	}

	return calibrations
}
