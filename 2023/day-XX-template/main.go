package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func partOne() int {
	return 0
}

func partTwo() int {
	return 0
}

func main() {
	now := time.Now()
	_ = parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne()
	part2 := partTwo()
	duration := time.Since(now)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseInput(input string) int {
	_ = strings.Split(input, "\n")
	return 0
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
