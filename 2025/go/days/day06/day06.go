package day06

import (
	"aoc/2025/go/lib"
	"strconv"
	"strings"
)

func parseInput(input string) ([][]int, []string) {
	lines := strings.Split(input, "\n")

	valueArrays := make([][]int, len(lines)-1)
	operations := strings.Fields(lines[len(lines)-1])

	for i := range len(lines) - 1 {
		for _, valueString := range strings.Fields(lines[i]) {
			value, _ := strconv.Atoi(valueString)
			valueArrays[i] = append(valueArrays[i], value)
		}
		if len(operations) != len(valueArrays[i]) {
			panic("Operations and values have different lengths")
		}
	}

	return valueArrays, operations
}

func performOperations(valueArrays [][]int, operations []string) int {
	sum := 0
	for i, operation := range operations {
		switch operation {
		case "*":
			result := 1
			for j := range len(valueArrays) {
				result *= valueArrays[j][i]
			}
			sum += result
		case "+":
			result := 0
			for j := range len(valueArrays) {
				result += valueArrays[j][i]
			}
			sum += result
		}
	}
	return sum
}

func Part1() any {
	input, _ := lib.ReadInput(6)
	valueArrays, operations := parseInput(input)
	return performOperations(valueArrays, operations)
}

func computeInputRightToLeft(input string) int {
	grid := lib.NewGrid(input)

	sum := 0

	var values []int
	for x := grid.Width() - 1; x >= 0; x-- {
		line := ""
		for y := 0; y < grid.Height(); y++ {
			line += string(grid.Get(x, y))
		}
		operation := ""
		if strings.HasSuffix(line, "+") {
			operation = "+"
		}
		if strings.HasSuffix(line, "*") {
			operation = "*"
		}
		line = strings.TrimSuffix(line, operation)
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		value, _ := strconv.Atoi(line)
		values = append(values, value)
		if operation != "" {
			switch operation {
			case "*":
				result := 1
				for _, value := range values {
					result *= value
				}
				sum += result
			case "+":
				result := 0
				for _, value := range values {
					result += value
				}
				sum += result
			}
			values = []int{}
		}
	}

	return sum
}

func Part2() any {
	input, _ := lib.ReadInput(6)
	return computeInputRightToLeft(input)
}
