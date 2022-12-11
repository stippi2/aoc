package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items               []int
	operation           func(old int) int
	testDivisor         int
	targetMonkeyIfTrue  int
	targetMonkeyIfFalse int
	itemsProcessed      int
}

func monkeyTurn(monkeys []Monkey, current int) {
	m := &monkeys[current]
	for _, item := range m.items {
		item = m.operation(item)
		item /= 3
		var targetMonkey int
		if item%m.testDivisor == 0 {
			targetMonkey = m.targetMonkeyIfTrue
		} else {
			targetMonkey = m.targetMonkeyIfFalse
		}
		monkeys[targetMonkey].items = append(monkeys[targetMonkey].items, item)
	}
	m.itemsProcessed += len(m.items)
	m.items = []int{}
}

func monkeyRound(monkeys []Monkey) {
	for i := 0; i < len(monkeys); i++ {
		monkeyTurn(monkeys, i)
	}
}

func main() {
	//monkeys := loadInput("puzzle-input.txt")
	monkeys := []Monkey{
		{
			items: []int{78, 53, 89, 51, 52, 59, 58, 85},
			operation: func(old int) int {
				return old * 3
			},
			testDivisor:         5,
			targetMonkeyIfTrue:  2,
			targetMonkeyIfFalse: 7,
		},
		{
			items: []int{64},
			operation: func(old int) int {
				return old + 7
			},
			testDivisor:         2,
			targetMonkeyIfTrue:  3,
			targetMonkeyIfFalse: 6,
		},
		{
			items: []int{71, 93, 65, 82},
			operation: func(old int) int {
				return old + 5
			},
			testDivisor:         13,
			targetMonkeyIfTrue:  5,
			targetMonkeyIfFalse: 4,
		},
		{
			items: []int{67, 73, 95, 75, 56, 74},
			operation: func(old int) int {
				return old + 8
			},
			testDivisor:         19,
			targetMonkeyIfTrue:  6,
			targetMonkeyIfFalse: 0,
		},
		{
			items: []int{85, 91, 90},
			operation: func(old int) int {
				return old + 4
			},
			testDivisor:         11,
			targetMonkeyIfTrue:  3,
			targetMonkeyIfFalse: 1,
		},
		{
			items: []int{67, 96, 69, 55, 70, 83, 62},
			operation: func(old int) int {
				return old * 2
			},
			testDivisor:         3,
			targetMonkeyIfTrue:  4,
			targetMonkeyIfFalse: 1,
		},
		{
			items: []int{53, 86, 98, 70, 64},
			operation: func(old int) int {
				return old + 6
			},
			testDivisor:         7,
			targetMonkeyIfTrue:  7,
			targetMonkeyIfFalse: 0,
		},
		{
			items: []int{88, 64},
			operation: func(old int) int {
				return old * old
			},
			testDivisor:         17,
			targetMonkeyIfTrue:  2,
			targetMonkeyIfFalse: 5,
		},
	}
	// part 1
	for i := 0; i < 20; i++ {
		monkeyRound(monkeys)
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].itemsProcessed >= monkeys[j].itemsProcessed
	})
	for i := 0; i < len(monkeys); i++ {
		fmt.Printf("items processed: %v\n", monkeys[i].itemsProcessed)
	}
	fmt.Printf("monkey business: %v\n", monkeys[0].itemsProcessed*monkeys[1].itemsProcessed)
}

func parseInput(input string) []Monkey {
	sections := strings.Split(input, "\n\n")
	monkeys := make([]Monkey, len(sections))
	for i, section := range sections {
		lines := strings.Split(section, "\n")
		items := strings.TrimPrefix(lines[1], "Starting items: ")
		for _, item := range strings.Split(items, ", ") {
			value, _ := strconv.Atoi(item)
			monkeys[i].items = append(monkeys[i].items, value)
		}
		// ...
	}
	return monkeys
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
