package day22

import (
	"aoc/2024/go/lib"
	"fmt"
	"strconv"
	"strings"
)

type Buyer struct {
	secretNumber int
	prices       []int
	changes      []int
}

func (b *Buyer) next() {
	t := b.secretNumber * 64
	b.secretNumber = b.secretNumber ^ t
	b.secretNumber = b.secretNumber % 16777216

	t = b.secretNumber / 32
	b.secretNumber = b.secretNumber ^ t
	b.secretNumber = b.secretNumber % 16777216

	t = b.secretNumber * 2048
	b.secretNumber = b.secretNumber ^ t
	b.secretNumber = b.secretNumber % 16777216
}

func (b *Buyer) generateSequence() {
	b.prices = make([]int, 2001)
	b.prices[0] = b.secretNumber % 10

	for i := 1; i <= 2000; i++ {
		b.next()
		b.prices[i] = b.secretNumber % 10
	}

	b.changes = make([]int, 2000)
	for i := 1; i <= 2000; i++ {
		b.changes[i-1] = b.prices[i] - b.prices[i-1]
	}
}

func findBestSequence(buyers []Buyer) (int, []int) {
	bestSum := 0
	bestSeq := make([]int, 4)

	// For all possible sequences...
	for d1 := -9; d1 <= 9; d1++ {
		for d2 := -9; d2 <= 9; d2++ {
			for d3 := -9; d3 <= 9; d3++ {
				for d4 := -9; d4 <= 9; d4++ {
					seq := []int{d1, d2, d3, d4}
					sum := 0

					// For every buyer...
					for i := 0; i < len(buyers); i++ {
						if pos := findSequence(buyers[i].changes, seq); pos >= 0 {
							sum += buyers[i].prices[pos+4]
						}
					}

					if sum > bestSum {
						bestSum = sum
						copy(bestSeq, seq)
					}
				}
			}
		}
	}

	return bestSum, bestSeq
}

func findSequence(changes []int, seq []int) int {
	for i := 0; i <= len(changes)-4; i++ {
		match := true
		for j := 0; j < 4; j++ {
			if changes[i+j] != seq[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func parseInput(input string) []Buyer {
	lines := strings.Split(input, "\n")
	buyers := make([]Buyer, len(lines))
	for i, numberString := range lines {
		buyers[i].secretNumber, _ = strconv.Atoi(numberString)
	}
	return buyers
}

func sum2000thSecretNumbers(buyers []Buyer) int {
	sum := 0
	for _, buyer := range buyers {
		for i := 0; i < 2000; i++ {
			buyer.next()
		}
		sum += buyer.secretNumber
	}
	return sum
}

func Part1() any {
	input, _ := lib.ReadInput(22)
	buyers := parseInput(input)
	return sum2000thSecretNumbers(buyers)
}

func Part2() any {
	input, _ := lib.ReadInput(22)
	buyers := parseInput(input)
	for i := 0; i < len(buyers); i++ {
		buyers[i].generateSequence()
	}
	mostBananas, bestSequence := findBestSequence(buyers)
	return fmt.Sprintf("Most bananas: %v (sequence: %v)", mostBananas, bestSequence)
}
