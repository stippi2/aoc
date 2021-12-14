package main

import (
	"io/ioutil"
	"strings"
)

type PolymerProcess struct {
	polymer        string
	insertionRules map[string]string
}

func main() {
}

func parseInput(input string) PolymerProcess {
	inputSections := strings.Split(input, "\n\n")
	if len(inputSections) != 2 {
		panic("expected sections in the input separated by a double line-break")
	}

	process := PolymerProcess{
		polymer:        inputSections[0],
		insertionRules: make(map[string]string),
	}

	insertionRules := strings.Split(inputSections[1], "\n")
	for _, rule := range insertionRules {
		parts := strings.Split(rule, " -> ")
		process.insertionRules[parts[0]] = parts[1]
	}
	return process
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
