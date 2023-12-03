package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func typeOfSymbol(symbol uint8) string {
	switch symbol {
	case '.':
		return "period"
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return "number"
	default:
		return "symbol"
	}
}

func getPartValue(lines []string, startOfNumber, endOfNumber, y int) int {
	number, _ := strconv.Atoi(lines[y][startOfNumber:endOfNumber])
	fmt.Printf("Found number %d", number)
	if number == 617 {
		fmt.Printf("")
	}

	minX := startOfNumber - 1
	if minX < 0 {
		minX = 0
	}
	maxX := endOfNumber + 1
	if maxX > len(lines[y]) {
		maxX = len(lines[y])
	}

	if y > 0 {
		for x := minX; x < maxX; x++ {
			if typeOfSymbol(lines[y-1][x]) == "symbol" {
				fmt.Printf(" - is part\n")
				return number
			}
		}
	}
	if y < len(lines)-1 {
		for x := minX; x < maxX; x++ {
			if typeOfSymbol(lines[y+1][x]) == "symbol" {
				fmt.Printf(" - is part\n")
				return number
			}
		}
	}
	if startOfNumber > 0 {
		if typeOfSymbol(lines[y][startOfNumber-1]) == "symbol" {
			fmt.Printf(" - is part\n")
			return number
		}
	}
	if endOfNumber < len(lines[y]) {
		if typeOfSymbol(lines[y][endOfNumber]) == "symbol" {
			fmt.Printf(" - is part\n")
			return number
		}
	}

	fmt.Printf(" - NOT PART\n")
	return 0
}

func addPartNumbers(lines []string) int {
	sumOfParts := 0

	for y, line := range lines {
		startOfNumber := -1
		for x := range line {
			if typeOfSymbol(line[x]) == "number" {
				if startOfNumber == -1 {
					startOfNumber = x
				}
			} else {
				if startOfNumber != -1 {
					sumOfParts += getPartValue(lines, startOfNumber, x, y)
					startOfNumber = -1
				}
			}
		}
		if startOfNumber != -1 {
			sumOfParts += getPartValue(lines, startOfNumber, len(line), y)
		}
	}
	return sumOfParts
}

func main() {
	lines := parseInput(loadInput("puzzle-input.txt"))
	sumOfPowers := addPartNumbers(lines)

	fmt.Printf("Sum of all part numbers: %d\n", sumOfPowers)
}

func parseInput(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, "\n")
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
