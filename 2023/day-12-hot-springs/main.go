package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type CacheKey struct {
	springs     string
	pos         int
	groupLength int
	groupsFound int
}

type Row struct {
	springs []byte
	groups  []int
	cache   map[CacheKey]int
}

func (r *Row) countMatchesAtHash(springs []byte, pos, groupLength, groupsFound int) int {
	if groupsFound+1 <= len(r.groups) && groupLength+1 <= r.groups[groupsFound] {
		// Only continue, if there are still more groups to be found, and if the current group is not already too long
		return r.countMatches(springs, pos+1, groupLength+1, groupsFound)
	}
	return 0
}

func (r *Row) countMatchesAtDot(springs []byte, pos, groupLength, groupsFound int) int {
	if groupLength == 0 {
		// (Dot at start of row, or after another dot)
		// Not currently accumulating a group, just continue
		return r.countMatches(springs, pos+1, 0, groupsFound)
	}
	if groupsFound+1 <= len(r.groups) && groupLength == r.groups[groupsFound] {
		// (First dot after a group)
		// Only continue, if there are still more groups to be found, and if the current group has the correct size
		return r.countMatches(springs, pos+1, 0, groupsFound+1)
	}
	return 0
}

func (r *Row) countMatches(springs []byte, pos, groupLength, groupsFound int) int {
	cacheKey := CacheKey{string(springs), pos, groupLength, groupsFound}
	if count, ok := r.cache[cacheKey]; ok {
		return count
	}

	count := 0

	// If we reached the end, the groups sizes must match
	if pos == len(springs) {
		// The following check can be avoided by appending "." to the springs, see cleanSprings()
		// It means we always arrive here with a closed group.
		//if groupLength > 0 && r.groups[groupsFound] == groupLength {
		//	groupsFound++
		//}
		if groupsFound == len(r.groups) {
			count = 1
		} else {
			count = 0
		}
	} else {
		switch springs[pos] {
		case '#':
			count += r.countMatchesAtHash(springs, pos, groupLength, groupsFound)
		case '.':
			count += r.countMatchesAtDot(springs, pos, groupLength, groupsFound)
		case '?':
			count += r.countMatchesAtHash(springs, pos, groupLength, groupsFound)
			count += r.countMatchesAtDot(springs, pos, groupLength, groupsFound)
		}
	}

	r.cache[cacheKey] = count
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
	return []byte(strings.Join(result, ".") + ".")
}

func (r *Row) findSolutions() int {
	r.cache = make(map[CacheKey]int)
	return r.countMatches(cleanSprings(r.springs), 0, 0, 0)
}

func findSolutions(rows []*Row) int {
	count := 0
	for i, row := range rows {
		fmt.Printf("#### On row %d of %d\n", i+1, len(rows))
		count += row.findSolutions()
		row.cache = nil
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
