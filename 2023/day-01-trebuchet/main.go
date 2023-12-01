package main

import (
	"fmt"
	"os"
	"strings"
)

func sum(items []int) int {
	total := 0
	for _, item := range items {
		total += item
	}
	return total
}

var tokenToValuePart1 = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var tokenToValuePart2 = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func main() {
	numbersPart1 := parseInput(loadInput("puzzle-input.txt"), tokenToValuePart1)
	fmt.Printf("The sum of the calibration numbers (part 1): %v\n", sum(numbersPart1))

	numbersPart2 := parseInput(loadInput("puzzle-input.txt"), tokenToValuePart2)
	fmt.Printf("The sum of the calibration numbers (part 2): %v\n", sum(numbersPart2))
}

func parseInput(input string, tokens map[string]int) []int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	result := make([]int, len(lines))
	for i, line := range lines {
		var numbers []int = nil
		for x := 0; x < len(line); x++ {
			rest := line[x:]
			for token, value := range tokens {
				if strings.HasPrefix(rest, token) {
					numbers = append(numbers, value)
					// That would look sensible, but it's wrong:
					// Consider "oneight" as in the example. We would miss the "eight" token.
					// x += len(token) - 1
				}
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
