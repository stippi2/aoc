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

func concatInts(a, b int64) int64 {
	multiplier := int64(1)
	temp := b
	for temp > 0 {
		multiplier *= 10
		temp /= 10
	}
	return a*multiplier + b
}

func solve(result, value int64, sequence []int64, includeConcatination bool) bool {
	if value > result {
		return false
	}
	if len(sequence) == 0 {
		return result == value
	}
	next := sequence[0]
	if solve(result, value+next, sequence[1:], includeConcatination) {
		return true
	}
	if solve(result, value*next, sequence[1:], includeConcatination) {
		return true
	}
	if includeConcatination && solve(result, concatInts(value, next), sequence[1:], true) {
		return true
	}
	return false
}

func (c *Calibration) validate(includeConcatination bool) bool {
	return solve(c.result, c.sequence[0], c.sequence[1:], includeConcatination)
}

func sumValidCalibrations(inputLines []string, includeConcatination bool) int64 {
	var sum int64 = 0
	calibrations := parseInput(inputLines)
	for _, c := range calibrations {
		if c.validate(includeConcatination) {
			sum += c.result
		}
	}
	return sum
}

func Part1() interface{} {
	inputLines, _ := lib.ReadInputLines(7)
	sum := sumValidCalibrations(inputLines, false)
	return fmt.Sprintf("Sum of valid calibrations: %v", sum)
}

func Part2() interface{} {
	inputLines, _ := lib.ReadInputLines(7)
	sum := sumValidCalibrations(inputLines, true)
	return fmt.Sprintf("Sum of valid calibrations: %v", sum)
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
