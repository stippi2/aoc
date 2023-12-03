package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

var gears = map[Pos][]int{}

func addNumberToGear(number int, x, y int) {
	gear := gears[Pos{x, y}]
	gear = append(gear, number)
	gears[Pos{x, y}] = gear
}

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

type Symbol struct {
	symbol string
	pos    Pos
}

func symbolAt(lines []string, x, y int) *Symbol {
	if x < 0 || y < 0 || y >= len(lines) || x >= len(lines[y]) {
		return nil
	}
	if typeOfSymbol(lines[y][x]) == "symbol" {
		return &Symbol{string(lines[y][x]), Pos{x, y}}
	}
	return nil
}

func appendIfSymbol(symbols []Symbol, lines []string, x, y int) []Symbol {
	symbol := symbolAt(lines, x, y)
	if symbol != nil {
		symbols = append(symbols, *symbol)
	}
	return symbols
}

func getPartValue(lines []string, startOfNumber, endOfNumber, y int) int {
	number, _ := strconv.Atoi(lines[y][startOfNumber:endOfNumber])
	var symbols []Symbol

	if y > 0 {
		for x := startOfNumber - 1; x <= endOfNumber; x++ {
			symbols = appendIfSymbol(symbols, lines, x, y-1)
		}
	}
	if y < len(lines)-1 {
		for x := startOfNumber - 1; x <= endOfNumber; x++ {
			symbols = appendIfSymbol(symbols, lines, x, y+1)
		}
	}
	symbols = appendIfSymbol(symbols, lines, startOfNumber-1, y)
	symbols = appendIfSymbol(symbols, lines, endOfNumber, y)

	if len(symbols) == 0 {
		return 0
	}
	for _, symbol := range symbols {
		if symbol.symbol == "*" {
			addNumberToGear(number, symbol.pos.x, symbol.pos.y)
		}
	}
	return number
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

	sumGearRatios := 0
	for _, gear := range gears {
		if len(gear) == 2 {
			sumGearRatios += gear[0] * gear[1]
		}
	}
	fmt.Printf("Sum of all gear ratios: %d\n", sumGearRatios)
}

func parseInput(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, "\n")
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
