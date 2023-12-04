package main

import (
	"fmt"
	"os"
	"strings"
)

type Card struct {
	winning   []string
	numbers   []string
	instances int
}

func (c *Card) contains(number string) bool {
	for _, w := range c.winning {
		if w == number {
			return true
		}
	}
	return false
}

func (c *Card) value() int {
	matches := c.countMatches()
	if matches == 0 {
		return 0
	}
	return 1 << (matches - 1)
}

func (c *Card) countMatches() int {
	matches := 0
	for _, number := range c.numbers {
		if c.contains(number) {
			matches++
		}
	}
	return matches
}

func partOne(cards []Card) int {
	sumOfWinning := 0
	for _, card := range cards {
		sumOfWinning += card.value()
	}
	return sumOfWinning
}

func partTwo(cards []Card) int {
	for i := 0; i < len(cards); i++ {
		value := cards[i].countMatches()
		for instance := 0; instance < cards[i].instances; instance++ {
			for j := i + 1; j <= i+value; j++ {
				cards[j].instances++
			}
		}
	}

	instances := 0
	for _, card := range cards {
		instances += card.instances
	}
	return instances
}

func main() {
	cards := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("Part 1: Sum of all winning values: %d\n", partOne(cards))
	fmt.Printf("Part 2: Sum of all instances: %d\n", partTwo(cards))
}

func parseNumbers(numbers string) []string {
	var result []string
	for _, number := range strings.Split(numbers, " ") {
		trimmedNumber := strings.TrimSpace(number)
		if trimmedNumber != "" {
			result = append(result, trimmedNumber)
		}
	}
	return result
}

func parseInput(input string) []Card {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	cards := make([]Card, len(lines))
	for i, line := range lines {
		prefixAndNumbers := strings.Split(line, ":")
		winningAndNumbers := strings.Split(prefixAndNumbers[1], "|")
		cards[i] = Card{
			winning:   parseNumbers(winningAndNumbers[0]),
			numbers:   parseNumbers(winningAndNumbers[1]),
			instances: 1,
		}
	}
	return cards
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
