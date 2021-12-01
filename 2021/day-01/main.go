package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	sweeps := loadSweeps("sweeps.txt")
	increasingSweeps := countIncreasingSweeps(sweeps, 1)
	increasingSweepWindows := countIncreasingSweeps(sweeps, 3)
	fmt.Printf("number of increasing sweeps: %v, windowed: %v\n", increasingSweeps, increasingSweepWindows)
}

func countIncreasingSweeps(sweeps []int, windowSize int) int {
	increasingSweeps := 0

	start := windowSize/2 + 1
	end := len(sweeps) - windowSize/2

	for i := start; i < end; i++ {
		sumCurrent := sumWindow(sweeps, i, windowSize)
		sumPrevious := sumWindow(sweeps, i-1, windowSize)
		if sumCurrent > sumPrevious {
			increasingSweeps++
		}
	}
	return increasingSweeps
}

func sumWindow(sweeps []int, center, windowSize int) int {
	start := center - windowSize/2
	if start < 0 {
		return 0
	}
	end := start + windowSize
	if end > len(sweeps) {
		return 0
	}
	sum := 0
	for i := start; i < end; i++ {
		sum += sweeps[i]
	}
	return sum
}

func loadSweeps(filename string) []int {
	fileContents, err := ioutil.ReadFile(filename)
	exitIfError(err)

	sweepStrings := strings.Split(string(fileContents), "\n")
	var sweeps []int
	for i, sweepString := range sweepStrings {
		if sweepString != "" {
			sweep, err := strconv.Atoi(sweepString)
			if err != nil {
				exitIfError(fmt.Errorf("invalid sweep '%s' at line %v: %w", sweepString, i+1, err))
			} else {
				sweeps = append(sweeps, sweep)
			}
		}
	}
	return sweeps
}

func exitIfError(err error) {
	if err != nil {
		panic(err)
	}
}
