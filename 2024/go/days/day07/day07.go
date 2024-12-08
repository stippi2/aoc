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

func (c *Calibration) validate() bool {
	return true
}

func sumValidCalibrations(inputLines []string) int64 {
	var sum int64 = 0
	calibrations := parseInput(inputLines)
	for _, c := range calibrations {
		if c.validate() {
			sum += c.result
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
	return fmt.Sprint("Sum of valid calibrations: %ld", sum)
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
