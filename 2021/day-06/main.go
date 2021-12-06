package main

import (
	"io/ioutil"
	"strings"
)

func main() {
}

func parseLanternFishAges(input string) (numberSequence []int) {
	return nil
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.Trim(string(fileContents), "\n")
}
