package main

import (
	"io/ioutil"
	"strings"
)

func main() {
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
