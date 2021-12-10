package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Chunk struct {
	opening, closing string
	score            int
}

func closing(opening string) string {
	switch opening {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	case "<":
		return ">"
	}
	return ""
}

func score(char string) int {
	switch char {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}
	return 0
}

func isClosing(char string) bool {
	return score(char) > 0
}

func parseLine(line string) int {
	var expectedClosings []string
	for i, char := range strings.Split(line, "") {
		if isClosing(char) {
			expected := expectedClosings[len(expectedClosings)-1]
			if char != expected {
				fmt.Printf("  syntax error at %v, expected: %s\n", i, expected)
				return score(char)
			}
			expectedClosings = expectedClosings[:len(expectedClosings)-1]
		} else {
			expectedClosings = append(expectedClosings, closing(char))
		}
	}
	if len(expectedClosings) > 0 {
		fmt.Printf("  incomplete line\n")
	}
	return 0
}

func parseLines(lines []string) int {
	totalScore := 0
	for i, line := range lines {
		lineScore := parseLine(line)
		if lineScore > 0 {
			totalScore += lineScore
			fmt.Printf("syntax error on line %v\n", i)
		}
	}
	return totalScore
}

func main() {
	lines := parseInput(loadInput("puzzle-input.txt"))
	totalScore := parseLines(lines)
	fmt.Printf("total syntax score: %v\n", totalScore)
}

func parseInput(input string) (lines []string) {
	return strings.Split(input, "\n")
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
