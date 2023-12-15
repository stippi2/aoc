package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func hash(s string) int {
	h := 0
	for _, c := range s {
		h = ((h + int(c)) * 17) % 256
	}
	return h
}

func partOne(sequence []string) int {
	sum := 0
	for _, s := range sequence {
		sum += hash(s)
	}
	return sum
}

func partTwo() int {
	return 0
}

func main() {
	now := time.Now()
	sequence := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(sequence)
	part2 := partTwo()
	duration := time.Since(now)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseInput(input string) []string {
	return strings.Split(input, ",")
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
