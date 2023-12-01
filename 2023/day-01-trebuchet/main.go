package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(items []int) int {
	total := 0
	for _, item := range items {
		total += item
	}
	return total
}

func main() {
	numbers := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("The sum of the calibration numbers: %v\n", sum(numbers))
}

func parseInput(input string) []int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	fmt.Printf("Found %v lines\n", len(lines))
	result := make([]int, len(lines))
	for i, line := range lines {
		var numbers []int = nil
		for x := 0; x < len(line); x++ {
			if line[x] >= '0' && line[x] <= '9' {
				number, _ := strconv.Atoi(line[x : x+1])
				numbers = append(numbers, number)
			}
		}
		result[i] = 10*numbers[0] + numbers[len(numbers)-1]
	}
	return result
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
