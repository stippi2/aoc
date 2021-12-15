package main

import (
	"io/ioutil"
	"strings"
)

type RiskMap struct {
}

func main() {
}

func parseInput(input string) *RiskMap {
	return nil
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
