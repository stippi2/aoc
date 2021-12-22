package main

import (
	"fmt"
	"github.com/stippi2/gopart"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	sweeps := loadSweeps("sweeps.txt")
	increasingSweeps := countIncreasingSweeps(sweeps)
	increasingSweepWindows := countIncreasingSweeps(sumPartitions(sweeps, 3, 1))
	fmt.Printf("number of increasing sweeps: %v, windowed: %v\n", increasingSweeps, increasingSweepWindows)
}

func countIncreasingSweeps(sweeps []int) int {
	increasingSweeps := 0
	for i := 1; i < len(sweeps); i++ {
		if sweeps[i-1] < sweeps[i] {
			increasingSweeps++
		}
	}
	return increasingSweeps
}

func sumPartitions(values []int, partitionSize, step int) []int {
	var sums []int
	for idxRange := range gopart.PartitionWithStep(len(values), partitionSize, step) {
		if idxRange.High-idxRange.Low < partitionSize {
			// Ignore remainder partition
			break
		}
		sum := 0
		for i := idxRange.Low; i < idxRange.High; i++ {
			sum += values[i]
		}
		sums = append(sums, sum)
	}
	return sums
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
