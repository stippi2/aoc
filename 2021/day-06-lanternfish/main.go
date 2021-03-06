package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func countLanternFish(countsPerAge []int64) int64 {
	count := int64(0)
	for i := 0; i < 9; i++ {
		count += countsPerAge[i]
	}
	return count
}

func simulateAgingAndReproduction(countsPerAge []int64) {
	reproductionCount := countsPerAge[0]
	// simulate aging by shifting the array
	// (number of fish at index 1 (age = 1) are placed at index 0 (age = 0)
	// ... and so forth
	for i := 1; i < 9; i++ {
		countsPerAge[i-1] = countsPerAge[i]
	}
	// simulate the reproduction by adding the number of fish with (age == 0) to age == 8
	countsPerAge[8] = reproductionCount
	// rotate the fish with day == 0 back to day == 6
	countsPerAge[6] += reproductionCount
}

func initAgeCounts(individualAges []int) (countsPerAge []int64) {
	countsPerAge = make([]int64, 9)
	for _, individualAge := range individualAges {
		if individualAge >= 0 && individualAge < len(countsPerAge) {
			countsPerAge[individualAge]++
		}
	}
	return
}

func main() {
	countsPerAge := initAgeCounts(parseLanternFishAges(loadInput("lanternfish-ages.txt")))
	startTime := time.Now()
	for i := 0; i < 256; i++ {
		simulateAgingAndReproduction(countsPerAge)
	}
	duration := time.Since(startTime)
	fmt.Printf("Number of lantern fish after 256 days: %v (took %v)\n", countLanternFish(countsPerAge), duration)
}

func parseLanternFishAges(input string) (individualAges []int) {
	numbers := strings.Split(input, ",")
	individualAges = make([]int, len(numbers))
	for i, numberString := range numbers {
		individualAges[i], _ = strconv.Atoi(numberString)
	}
	return
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.Trim(string(fileContents), "\n")
}
