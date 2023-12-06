package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time           int
	recordDistance int
}

func countWaysToWin(race Race) int {
	left := 0
	for left < race.time {
		distance := left * (race.time - left)
		if distance > race.recordDistance {
			break
		}
		left++
	}
	right := race.time
	for right > 0 {
		distance := right * (race.time - right)
		if distance > race.recordDistance {
			break
		}
		right--
	}
	return right - left + 1
}

func partOne(races []Race) int {
	product := 1
	for _, race := range races {
		product *= countWaysToWin(race)
	}
	return product
}

func mergeRaces(races []Race) Race {
	time := ""
	distance := ""
	for _, race := range races {
		time += fmt.Sprintf("%d", race.time)
		distance += fmt.Sprintf("%d", race.recordDistance)
	}
	timeInt, _ := strconv.Atoi(time)
	distanceInt, _ := strconv.Atoi(distance)
	return Race{time: timeInt, recordDistance: distanceInt}
}

func partTwo(races []Race) int {
	race := mergeRaces(races)
	return countWaysToWin(race)
}

func main() {
	races := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("Part 1: Ways to win: %d\n", partOne(races))
	fmt.Printf("Part 2: Ways to win merged race: %d\n", partTwo(races))
}

func toInts(input string) []int {
	fields := strings.Fields(input)
	ints := make([]int, len(fields))
	for i, field := range fields {
		ints[i], _ = strconv.Atoi(field)
	}
	return ints
}

func parseInput(input string) []Race {
	lines := strings.Split(input, "\n")
	times := toInts(strings.TrimPrefix(lines[0], "Time: "))
	distances := toInts(strings.TrimPrefix(lines[1], "Distance: "))
	if len(times) != len(distances) {
		panic("times and distances must be the same length")
	}
	races := make([]Race, len(times))
	for i := range times {
		races[i] = Race{time: times[i], recordDistance: distances[i]}
	}
	return races
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
