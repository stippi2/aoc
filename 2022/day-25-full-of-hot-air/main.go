package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func fromSnafu(s string) int {
	value := 0
	for i := len(s) - 1; i >= 0; i-- {
		pot := int(math.Ceil(math.Pow(5, float64(len(s)-1-i))))
		switch s[i] {
		case '2':
			pot *= 2
		case '1':
			pot *= 1
		case '0':
			pot *= 0
		case '-':
			pot *= -1
		case '=':
			pot *= -2
		}
		value += pot
	}
	return value
}

func fromDecimal(value int) string {
	return ""
}

func sum(values []int) int {
	s := 0
	for _, v := range values {
		s += v
	}
	return s
}

func main() {
	values := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("sum: %v\n", sum(values))
}

func parseInput(input string) []int {
	var values []int
	for _, snafu := range strings.Split(input, "\n") {
		values = append(values, fromSnafu(snafu))
	}
	return values
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
