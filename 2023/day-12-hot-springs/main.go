package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Row struct {
	springs   []byte
	groups    []int
	solutions map[string]int
	failures  map[string]bool
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if b[i] != a[i] {
			return false
		}
	}
	return true
}

func finishGroup(groups []int, groupLength int) []int {
	if groupLength > 0 {
		return append(groups, groupLength)
	}
	return groups
}

func (r *Row) countMatches(springs []byte, pos, groupLength int, groups []int) int {
	//fmt.Printf("%s: found %d of %d matches\n", springs, len(groups), len(r.groups))
	if groupLength > 0 {
		// Check if there is a pending group when there should not be another one
		if len(groups) == len(r.groups) {
			return 0
		}
		// Check if current group length is already bigger than the next group
		if groupLength > r.groups[len(groups)] {
			return 0
		}
	}
	// If we reached the end, the groups must match
	if pos == len(springs) {
		if groupLength > 0 {
			groups = append(groups, groupLength)
		}
		if equals(groups, r.groups) {
			return 1
		}
		return 0
	}

	count := 0
	field := springs[pos]
	switch field {
	case '?':
		springs[pos] = '#'
		count += r.countMatches(springs, pos+1, groupLength+1, groups)
		springs[pos] = '.'
		count += r.countMatches(springs, pos+1, 0, finishGroup(groups, groupLength))
		springs[pos] = field
	case '.':
		count += r.countMatches(springs, pos+1, 0, finishGroup(groups, groupLength))
	case '#':
		count += r.countMatches(springs, pos+1, groupLength+1, groups)
	}
	return count
}

func cleanSprings(springs []byte) []byte {
	parts := strings.Split(string(springs), ".")
	var result []string
	for _, part := range parts {
		if len(part) > 0 {
			result = append(result, part)
		}
	}
	return []byte(strings.Join(result, "."))
}

func (r *Row) findSolutions() int {
	return r.countMatches(cleanSprings(r.springs), 0, 0, nil)
}

func findSolutions(rows []*Row) int {
	count := 0
	for i, row := range rows {
		fmt.Printf("#### On row %d of %d\n", i+1, len(rows))
		count += row.findSolutions()
	}
	return count
}

func partOne(rows []*Row) int {
	return findSolutions(rows)
}

func (r *Row) unfold() {
	springs := make([]byte, len(r.springs)*5+4)
	var groups []int
	for i := 0; i < 5; i++ {
		copy(springs[i*(len(r.springs)+1):], r.springs)
		if i < 4 {
			springs[(i+1)*(len(r.springs)+1)-1] = '?'
		}
		groups = append(groups, r.groups...)
	}
	r.springs = springs
	r.groups = groups
}

func partTwo(rows []*Row) int {
	for _, row := range rows {
		row.unfold()
	}
	return findSolutions(rows)
}

func main() {
	now := time.Now()
	rows := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(rows)
	fmt.Printf("Part 1: %d\n", part1)
	part2 := partTwo(rows)
	fmt.Printf("Part 2: %d\n", part2)
	duration := time.Since(now)
	fmt.Printf("Time: %v\n", duration)
}

func parseInput(input string) []*Row {
	lines := strings.Split(input, "\n")
	rows := make([]*Row, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		rows[i] = &Row{springs: []byte(parts[0])}
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
