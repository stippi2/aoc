package main

import (
	"fmt"
	"math"
	"sort"
)

type FactorizedNumber struct {
	primeFactors map[int]int
}

func (n *FactorizedNumber) multiplyBy(prime int) {
	n.primeFactors[prime] = n.primeFactors[prime] + 1
}

func (n *FactorizedNumber) multiplyBySelf() {
	newPrimeFactors := map[int]int{}
	for prime, potency := range n.primeFactors {
		newPrimeFactors[prime] = potency * 2
	}
	n.primeFactors = newPrimeFactors
}

func (n *FactorizedNumber) add(value int) {
	// 4, 5, 6, 7, 8

	// 2*2 * 7*7 + 8  (196 + 8 = 204) = 2*2 * 3 * 17

	// 2*7+8
	// 2*11

	// 2 * 7 + 8      (14 + 8 = 22)
	// 2 * 11

	// 3*3 * 5 + 4    (45 + 4 = 49)
	// 7*7

	// 2*2 * 11       (44 + 5 = 49)
	// 7*7

	defactorized := 1
	for prime, potency := range n.primeFactors {
		defactorized *= int(math.Pow(float64(prime), float64(potency)))
	}
	defactorized += value
	newFactorized := toFactorizedNumber(defactorized)
	n.primeFactors = newFactorized.primeFactors
}

func (n *FactorizedNumber) isDivisibleBy(prime int) bool {
	// prime is one of 2, 3, 5, 7, 11, 13, 17, 19
	return n.primeFactors[prime] > 0
}

func (n *FactorizedNumber) String() string {
	result := ""
	for prime, potency := range n.primeFactors {
		if result != "" {
			result += "*"
		}
		result += fmt.Sprintf("%v^%v", prime, potency)
	}
	return result
}

func toFactorizedNumber(n int) *FactorizedNumber {
	// From: https://siongui.github.io/2017/05/09/go-find-all-prime-factors-of-integer-number/
	fn := FactorizedNumber{primeFactors: map[int]int{}}
	// Get the number of 2s that divide n
	for n%2 == 0 {
		fn.multiplyBy(2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			fn.multiplyBy(i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number greater than 2
	if n > 2 {
		fn.multiplyBy(n)
	}

	return &fn
}

func convertToFactorizedNumbers(values []int) []*FactorizedNumber {
	var factorizeNumbers []*FactorizedNumber
	for _, value := range values {
		factorizeNumbers = append(factorizeNumbers, toFactorizedNumber(value))
	}
	return factorizeNumbers
}

type Monkey struct {
	items               []*FactorizedNumber
	operation           func(fn *FactorizedNumber)
	testDivisor         int
	targetMonkeyIfTrue  int
	targetMonkeyIfFalse int
	itemsProcessed      int
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
	m.itemsProcessed += len(m.items)
	m.items = []*FactorizedNumber{}
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
			items: convertToFactorizedNumbers([]int{78, 53, 89, 51, 52, 59, 58, 85}),
			operation: func(fn *FactorizedNumber) {
				fn.multiplyBy(3)
			},
			testDivisor:         5,
			targetMonkeyIfTrue:  2,
			targetMonkeyIfFalse: 7,
		},
		{
			items: convertToFactorizedNumbers([]int{64}),
			operation: func(fn *FactorizedNumber) {
				fn.add(7)
			},
			testDivisor:         2,
			targetMonkeyIfTrue:  3,
			targetMonkeyIfFalse: 6,
		},
		{
			items: convertToFactorizedNumbers([]int{71, 93, 65, 82}),
			operation: func(fn *FactorizedNumber) {
				fn.add(5)
			},
			testDivisor:         13,
			targetMonkeyIfTrue:  5,
			targetMonkeyIfFalse: 4,
		},
		{
			items: convertToFactorizedNumbers([]int{67, 73, 95, 75, 56, 74}),
			operation: func(fn *FactorizedNumber) {
				fn.add(8)
			},
			testDivisor:         19,
			targetMonkeyIfTrue:  6,
			targetMonkeyIfFalse: 0,
		},
		{
			items: convertToFactorizedNumbers([]int{85, 91, 90}),
			operation: func(fn *FactorizedNumber) {
				fn.add(4)
			},
			testDivisor:         11,
			targetMonkeyIfTrue:  3,
			targetMonkeyIfFalse: 1,
		},
		{
			items: convertToFactorizedNumbers([]int{67, 96, 69, 55, 70, 83, 62}),
			operation: func(fn *FactorizedNumber) {
				fn.multiplyBy(2)
			},
			testDivisor:         3,
			targetMonkeyIfTrue:  4,
			targetMonkeyIfFalse: 1,
		},
		{
			items: convertToFactorizedNumbers([]int{53, 86, 98, 70, 64}),
			operation: func(fn *FactorizedNumber) {
				fn.add(6)
			},
			testDivisor:         7,
			targetMonkeyIfTrue:  7,
			targetMonkeyIfFalse: 0,
		},
		{
			items: convertToFactorizedNumbers([]int{88, 64}),
			operation: func(fn *FactorizedNumber) {
				fn.multiplyBySelf()
			},
			testDivisor:         17,
			targetMonkeyIfTrue:  2,
			targetMonkeyIfFalse: 5,
		},
	}

	for i, monkey := range monkeys {
		fmt.Printf("Monkey %v: [", i)
		for j, item := range monkey.items {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Print(item)
		}
		fmt.Print("]\n\n")
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
