package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

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

func syntaxScore(char string) int {
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

func autocompleteScore(char string) int {
	switch char {
	case ")":
		return 1
	case "]":
		return 2
	case "}":
		return 3
	case ">":
		return 4
	}
	return 0
}

func isClosing(char string) bool {
	return syntaxScore(char) > 0
}

func parseLine(line string) (int, []string) {
	var expectedClosings []string
	for _, char := range strings.Split(line, "") {
		if isClosing(char) {
			expected := expectedClosings[len(expectedClosings)-1]
			if char != expected {
				return syntaxScore(char), expectedClosings
			}
			expectedClosings = expectedClosings[:len(expectedClosings)-1]
		} else {
			expectedClosings = append(expectedClosings, closing(char))
		}
	}
	return 0, expectedClosings
}

func totalSyntaxErrorScore(lines []string) int {
	totalScore := 0
	for _, line := range lines {
		score, _ := parseLine(line)
		totalScore += score
	}
	return totalScore
}

func computeAutocompleteScore(completions []string) int {
	score := 0
	for i := len(completions) - 1; i >= 0; i-- {
		score = 5*score + autocompleteScore(completions[i])
	}
	return score
}

func totalAutocompleteScore(lines []string) int {
	var scores []int
	for _, line := range lines {
		score, expectedCompletions := parseLine(line)
		if score == 0 {
			scores = append(scores, computeAutocompleteScore(expectedCompletions))
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	lines := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("total syntax score: %v\n", totalSyntaxErrorScore(lines))
	fmt.Printf("total autocompletion score: %v\n", totalAutocompleteScore(lines))
}

func parseInput(input string) (lines []string) {
	return strings.Split(input, "\n")
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
