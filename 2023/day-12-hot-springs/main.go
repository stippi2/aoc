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

func (r *Row) maxSprings() int {
	count := 0
	for _, group := range r.groups {
		count += group
	}
	return count
}

func (r *Row) countSprings() int {
	count := 0
	for _, spring := range r.springs {
		if spring == '#' {
			count++
		}
	}
	return count
}

func (r *Row) countCandidates() int {
	count := 0
	for _, spring := range r.springs {
		if spring == '?' {
			count++
		}
	}
	return count
}

func (r *Row) matches(springs []byte) bool {
	count := 0
	var groups []int
	for _, spring := range springs {
		if spring == '#' || spring == '?' {
			count++
		} else {
			if count > 0 {
				groups = append(groups, count)
				count = 0
			}
		}
	}
	if count > 0 {
		groups = append(groups, count)
	}
	if len(groups) != len(r.groups) {
		return false
	}
	for i := range groups {
		if groups[i] != r.groups[i] {
			return false
		}
	}
	return true
}

func copySprings(springs []byte) []byte {
	c := make([]byte, len(springs))
	copy(c, springs)
	return c
}

func compress(springs []byte) string {
	var compressed []string
	for _, chunk := range strings.Split(string(springs), ".") {
		if len(chunk) > 0 {
			compressed = append(compressed, chunk)
		}
	}
	result := strings.Join(compressed, ".")
	return result
}

func (r *Row) generateSolutions(current []byte, remaining int, startIndex int, iteration *int) (int, int) {
	*iteration++
	if remaining == 0 {
		if r.matches(current) {
			fmt.Printf("  found solution: %s\n", string(current))
			return 1, 0
		}
		return 0, 1
	}

	compressed := compress(current)
	if priorSolutions, ok := r.solutions[compressed]; ok {
		fmt.Printf("  found solutions for compressed: %s, %d\n", compressed, priorSolutions)
		return priorSolutions, 0
	}
	if r.failures[compressed] {
		fmt.Printf("  found failure for compressed: %s\n", compressed)
		return 0, 1
	}

	if *iteration%10000 == 0 {
		fmt.Printf("  iteration %d, tesing %s\n", *iteration, string(current))
	}

	solutions := 0
	failures := 0

	for i := startIndex; i < len(current); i++ {
		if current[i] == '?' {
			current[i] = '.'
			s, f := r.generateSolutions(current, remaining-1, i+1, iteration)
			solutions += s
			failures += f
			current[i] = '?'
		}
	}

	if solutions > 0 && failures == 0 {
		// Only cache something if it had only solutions and no failures
		r.solutions[compressed] = solutions
	} else if solutions == 0 && failures > 0 {
		r.failures[compressed] = true
	}

	return solutions, failures
}

func (r *Row) findSolutions() int {
	iteration := 0
	r.solutions = make(map[string]int)
	r.failures = make(map[string]bool)
	solutions, _ := r.generateSolutions(copySprings(r.springs), r.countCandidates()-(r.maxSprings()-r.countSprings()), 0, &iteration)
	return solutions
}

func findSolutions(rows []*Row) int {
	count := 0
	for _, row := range rows {
		//fmt.Printf("#### On row %d of %d\n", i, len(rows))
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
	part2 := partTwo(rows)
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
