package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Row struct {
	data   []byte
	groups []int
}

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

func parseInput(input string) []*Row {
	lines := strings.Split(input, "\n")
	rows := make([]*Row, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		rows[i] = &Row{data: []byte(parts[0])}
		groups := strings.Split(parts[1], ",")
		for _, group := range groups {
			value, _ := strconv.Atoi(group)
			rows[i].groups = append(rows[i].groups, value)
		}
	}
	return rows
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
