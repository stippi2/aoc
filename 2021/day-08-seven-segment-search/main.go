package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func countDigits(digits [][]string, condition func(digit string) bool) int {
	sum := 0
	for _, display := range digits {
		for _, digit := range display {
			if condition(digit) {
				sum++
			}
		}
	}
	return sum
}

func conditionOnesFoursSevensAndEights(digit string) bool {
	switch len(digit) {
	case 2, 4, 3, 7:
		return true
	}
	return false
}

func main() {
	_, digits := parseInput(loadInput("puzzle-input.txt"))
	count := countDigits(digits, conditionOnesFoursSevensAndEights)
	fmt.Printf("number of 1, 4, 7, and 8: %v\n", count)
}

func parseInput(input string) (signals, digits [][]string) {
	lines := strings.Split(input, "\n")
	signals = make([][]string, len(lines))
	digits = make([][]string, len(lines))
	for i, line := range lines {
		signalsDigits := strings.Split(line, " | ")
		signals[i] = strings.Split(signalsDigits[0], " ")
		digits[i] = strings.Split(signalsDigits[1], " ")
	}
	return
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
