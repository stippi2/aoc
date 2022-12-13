package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NoDecision int = 0
	Sorted         = 1
	Unsorted       = 2
)

func compare(left, right any) int {
	lInt, lIsInt := left.(int)
	rInt, rIsInt := right.(int)
	// If both are integers:
	if lIsInt && rIsInt {
		if lInt < rInt {
			return Sorted
		} else if lInt > rInt {
			return Unsorted
		} else {
			return NoDecision
		}
	}
	// Mismatched types:
	if lIsInt {
		return compare([]any{lInt}, right)
	}
	if rIsInt {
		return compare(left, []any{rInt})
	}
	// Compare lists:
	leftList := left.([]any)
	rightList := right.([]any)
	leftLength := len(leftList)
	rightLength := len(rightList)
	for i := 0; i < leftLength; i++ {
		if i >= rightLength {
			return Unsorted
		}
		innerCompare := compare(leftList[i], rightList[i])
		if innerCompare != NoDecision {
			return innerCompare
		}
	}
	if rightLength > leftLength {
		return Sorted
	}
	return NoDecision
}

type Pair struct {
	left, right any
}

func sumIndicesOrderedPairs(pairs []Pair) int {
	sumIndices := 0
	for i := 0; i < len(pairs); i++ {
		comparison := compare(pairs[i].left, pairs[i].right)
		if comparison == Sorted {
			fmt.Printf("sorted: %v\n", i+1)
			sumIndices += i + 1
		}
		if comparison == NoDecision {
			fmt.Printf("no decision: %v\n", i+1)
		}
	}
	return sumIndices
}

func main() {
	pairs := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("sum of the indices of the ordered pairs: %v\n", sumIndicesOrderedPairs(pairs))
}

func parseItem(n string) any {
	isArray := strings.HasPrefix(n, "[")
	if isArray {
		n = strings.TrimPrefix(n, "[")
		n = strings.TrimSuffix(n, "]")
	}

	items := []any{}

	start := 0
	for start < len(n) {
		var item any
		if n[start] == '[' {
			level := 1
			end := len(n)
			for j := start + 1; j < len(n); j++ {
				switch n[j : j+1] {
				case "[":
					level++
				case "]":
					level--
				}
				if level == 0 {
					end = j + 1
					break
				}
			}
			item = parseItem(n[start:end])
			start = end
		} else if n[start] >= '0' && n[start] <= '9' {
			end := start + 1
			for j := start + 1; j < len(n); j++ {
				if n[j] < '0' || n[j] > '9' {
					end = j
					break
				}
			}
			number, err := strconv.Atoi(n[start:end])
			if err != nil {
				panic(fmt.Sprintf("failed to parse '%s' as number: %s", n[start:end], err))
			}
			item = number
			start = end
		} else {
			start++
		}
		if item != nil {
			items = append(items, item)
		}
	}
	if !isArray {
		if len(items) > 1 {
			panic("not an array?!")
		}
		return items[0]
	}
	return items
}

func parseInput(input string) []Pair {
	pairSections := strings.Split(input, "\n\n")
	pairs := make([]Pair, len(pairSections))
	for i, section := range pairSections {
		parts := strings.Split(section, "\n")
		pairs[i].left = parseItem(parts[0])
		pairs[i].right = parseItem(parts[1])
	}
	return pairs
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
