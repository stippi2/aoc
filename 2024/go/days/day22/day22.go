package day22

import (
	"aoc/2024/go/lib"
	"strconv"
	"strings"
)

type Buyer struct {
	secretNumber int
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
	return "Not implemented"
}
