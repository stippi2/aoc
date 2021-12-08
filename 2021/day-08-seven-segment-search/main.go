package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
}

func parseSequence(input string) (numbers []int) {
	numberStrings := strings.Split(input, ",")
	numbers = make([]int, len(numberStrings))
	for i, numberString := range numberStrings {
		numbers[i], _ = strconv.Atoi(numberString)
	}
	return
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.Trim(string(fileContents), "\n")
}
