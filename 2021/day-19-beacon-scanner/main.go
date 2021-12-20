package main

import (
	"io/ioutil"
	"strings"
)

func main() {
}

func parseInput(input string) {
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
