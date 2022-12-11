package main

import (
	"fmt"
	"sort"
)

type RemainderNumber struct {
	remainders map[int]int
}

func (n *RemainderNumber) multiplyBy(value int) {
	newRemainders := map[int]int{}
	for prime, remainder := range n.remainders {
		newRemainders[prime] = (remainder * value) % prime
	}
	n.remainders = newRemainders
}

func (n *RemainderNumber) multiplyBySelf() {
	newRemainders := map[int]int{}
	for prime, remainder := range n.remainders {
		newRemainders[prime] = (remainder * remainder) % prime
	}
	n.remainders = newRemainders
}

func (n *RemainderNumber) add(value int) {
	newRemainders := map[int]int{}
	for prime, remainder := range n.remainders {
		newRemainders[prime] = (remainder + value) % prime
	}
	n.remainders = newRemainders
}

func (n *RemainderNumber) isDivisibleBy(prime int) bool {
	// prime is one of 2, 3, 5, 7, 11, 13, 17, 19
	return n.remainders[prime] == 0
}

func toRemainderNumber(n int) *RemainderNumber {
	fn := RemainderNumber{remainders: map[int]int{}}
	for _, prime := range []int{2, 3, 5, 7, 11, 13, 17, 19} {
		fn.remainders[prime] = n % prime
	}
	return &fn
}

func convertToRemainderNumbers(values []int) []*RemainderNumber {
	var factorizeNumbers []*RemainderNumber
	for _, value := range values {
		factorizeNumbers = append(factorizeNumbers, toRemainderNumber(value))
	}
	return factorizeNumbers
}

type Monkey struct {
	items               []*RemainderNumber
	operation           func(fn *RemainderNumber)
	testDivisor         int
	targetMonkeyIfTrue  int
	targetMonkeyIfFalse int
	itemsProcessed      int64
}

func monkeyTurn(monkeys []Monkey, current int) {
	m := &monkeys[current]
	for _, item := range m.items {
		m.operation(item)
		var targetMonkey int
		if item.isDivisibleBy(m.testDivisor) {
			targetMonkey = m.targetMonkeyIfTrue
		} else {
			targetMonkey = m.targetMonkeyIfFalse
		}
		monkeys[targetMonkey].items = append(monkeys[targetMonkey].items, item)
	}
	m.itemsProcessed += int64(len(m.items))
	m.items = []*RemainderNumber{}
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
			items: convertToRemainderNumbers([]int{78, 53, 89, 51, 52, 59, 58, 85}),
			operation: func(fn *RemainderNumber) {
				fn.multiplyBy(3)
			},
			testDivisor:         5,
			targetMonkeyIfTrue:  2,
			targetMonkeyIfFalse: 7,
		},
		{
			items: convertToRemainderNumbers([]int{64}),
			operation: func(fn *RemainderNumber) {
				fn.add(7)
			},
			testDivisor:         2,
			targetMonkeyIfTrue:  3,
			targetMonkeyIfFalse: 6,
		},
		{
			items: convertToRemainderNumbers([]int{71, 93, 65, 82}),
			operation: func(fn *RemainderNumber) {
				fn.add(5)
			},
			testDivisor:         13,
			targetMonkeyIfTrue:  5,
			targetMonkeyIfFalse: 4,
		},
		{
			items: convertToRemainderNumbers([]int{67, 73, 95, 75, 56, 74}),
			operation: func(fn *RemainderNumber) {
				fn.add(8)
			},
			testDivisor:         19,
			targetMonkeyIfTrue:  6,
			targetMonkeyIfFalse: 0,
		},
		{
			items: convertToRemainderNumbers([]int{85, 91, 90}),
			operation: func(fn *RemainderNumber) {
				fn.add(4)
			},
			testDivisor:         11,
			targetMonkeyIfTrue:  3,
			targetMonkeyIfFalse: 1,
		},
		{
			items: convertToRemainderNumbers([]int{67, 96, 69, 55, 70, 83, 62}),
			operation: func(fn *RemainderNumber) {
				fn.multiplyBy(2)
			},
			testDivisor:         3,
			targetMonkeyIfTrue:  4,
			targetMonkeyIfFalse: 1,
		},
		{
			items: convertToRemainderNumbers([]int{53, 86, 98, 70, 64}),
			operation: func(fn *RemainderNumber) {
				fn.add(6)
			},
			testDivisor:         7,
			targetMonkeyIfTrue:  7,
			targetMonkeyIfFalse: 0,
		},
		{
			items: convertToRemainderNumbers([]int{88, 64}),
			operation: func(fn *RemainderNumber) {
				fn.multiplyBySelf()
			},
			testDivisor:         17,
			targetMonkeyIfTrue:  2,
			targetMonkeyIfFalse: 5,
		},
	}

	// part 1
	for i := 0; i < 10000; i++ {
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

//
//func parseInput(input string) []Monkey {
//	sections := strings.Split(input, "\n\n")
//	monkeys := make([]Monkey, len(sections))
//	for i, section := range sections {
//		lines := strings.Split(section, "\n")
//		items := strings.TrimPrefix(lines[1], "Starting items: ")
//		for _, item := range strings.Split(items, ", ") {
//			value, _ := strconv.Atoi(item)
//			monkeys[i].items = append(monkeys[i].items, value)
//		}
//		// ...
//	}
//	return monkeys
//}
//
//func loadInput(filename string) string {
//	fileContents, _ := os.ReadFile(filename)
//	return string(fileContents)
//}
