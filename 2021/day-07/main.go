package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func minMax(sequence []int) (min int, max int) {
	min = math.MaxInt32
	max = math.MinInt32
	for _, n := range sequence {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return
}

func fuelConsumption(positions []int, pos int) (totalFuel int) {
	for _, n := range positions {
		dist := pos - n
		if dist < 0 {
			dist = -dist
		}
		fuel := 0
		for i := 1; i <= dist; i++ {
			fuel += i
		}
		totalFuel += fuel
	}
	return
}

func findOptimumPos(positions []int) (minFuel, bestPos int) {
	min, max := minMax(positions)
	minFuel = math.MaxInt32
	for pos := min; pos <= max; pos++ {
		fuel := fuelConsumption(positions, pos)
		if fuel < minFuel {
			minFuel = fuel
			bestPos = pos
		}
	}
	return
}

func main() {
	positions := parseSequence(loadInput("crab-positions.txt"))
	minFuel, bestPos := findOptimumPos(positions)
	fmt.Printf("best position: %v, fuel consumption: %v\n", bestPos, minFuel)
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
