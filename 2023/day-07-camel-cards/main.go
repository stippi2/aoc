package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var typeOfHands = map[string]int{
	"FiveOfAKind":  7,
	"FourOfAKind":  6,
	"FullHouse":    5,
	"ThreeOfAKind": 4,
	"TwoPair":      3,
	"OnePair":      2,
	"HighCard":     1,
}

var cardValues = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var cardValuesWithJokers = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

type HandAndBid struct {
	hand       string
	bid        int
	typeOfHand string
}

func maxCount(cards map[string]int) int {
	m := 0
	for _, count := range cards {
		m = max(m, count)
	}
	return m
}

func (h HandAndBid) getType() string {
	if h.typeOfHand == "" {
		cards := make(map[string]int)
		for _, card := range h.hand {
			cards[string(card)]++
		}
		switch len(cards) {
		case 1:
			h.typeOfHand = "FiveOfAKind"
		case 2:
			switch maxCount(cards) {
			case 4:
				h.typeOfHand = "FourOfAKind"
			case 3:
				h.typeOfHand = "FullHouse"
			}
		case 3:
			switch maxCount(cards) {
			case 3:
				h.typeOfHand = "ThreeOfAKind"
			case 2:
				h.typeOfHand = "TwoPair"
			}
		case 4:
			h.typeOfHand = "OnePair"
		case 5:
			h.typeOfHand = "HighCard"
		}
	}
	return h.typeOfHand
}

func (h HandAndBid) getValue() int {
	return typeOfHands[h.getType()]
}

func (h HandAndBid) compare(other HandAndBid) int {
	if h.getValue() == other.getValue() {
		for i := 0; i < 5; i++ {
			hCard := string(h.hand[i])
			oCard := string(other.hand[i])
			if hCard != oCard {
				return cardValues[hCard] - cardValues[oCard]
			}
		}
		return 0
	}
	return h.getValue() - other.getValue()
}

func partOne(hands []HandAndBid) int {
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].compare(hands[j]) <= 0
	})
	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += hand.bid * (i + 1)
	}
	return totalWinnings
}

func partTwo(hands []HandAndBid) int {
	return 0
}

func main() {
	now := time.Now()
	handsAndBids := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(handsAndBids)
	duration := time.Since(now)
	fmt.Printf("Part 1: Total winnings: %d\n", part1)
	fmt.Printf("Time: %v\n", duration)
}

func parseInput(input string) []HandAndBid {
	lines := strings.Split(input, "\n")
	handsAndBids := make([]HandAndBid, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		handsAndBids[i].hand = parts[0]
		handsAndBids[i].bid, _ = strconv.Atoi(parts[1])
	}
	return handsAndBids
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
