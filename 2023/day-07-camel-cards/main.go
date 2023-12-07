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
	hand                 string
	bid                  int
	typeOfHand           string
	typeOfHandWithJokers string
}

func maxCount(cards map[string]int) (int, string) {
	m := 0
	c := ""
	for card, count := range cards {
		if count > m {
			m = count
			c = card
		}
	}
	return m, c
}

func (h HandAndBid) getType(withJokers bool) string {
	typeOfHand := ""
	if withJokers {
		typeOfHand = h.typeOfHand
	} else {
		typeOfHand = h.typeOfHandWithJokers
	}
	if typeOfHand == "" {
		cards := make(map[string]int)
		for _, card := range h.hand {
			cards[string(card)]++
		}

		if withJokers {
			jokers := cards["J"]
			if jokers > 0 {
				delete(cards, "J")
			}
			_, bestCard := maxCount(cards)
			cards[bestCard] += jokers
		}

		sameCardCount, _ := maxCount(cards)

		switch len(cards) {
		case 1:
			typeOfHand = "FiveOfAKind"
		case 2:
			switch sameCardCount {
			case 4:
				typeOfHand = "FourOfAKind"
			case 3:
				typeOfHand = "FullHouse"
			}
		case 3:
			switch sameCardCount {
			case 3:
				typeOfHand = "ThreeOfAKind"
			case 2:
				typeOfHand = "TwoPair"
			}
		case 4:
			typeOfHand = "OnePair"
		case 5:
			typeOfHand = "HighCard"
		}
		if withJokers {
			h.typeOfHand = typeOfHand
		} else {
			h.typeOfHandWithJokers = typeOfHand
		}
	}
	return typeOfHand
}

func (h HandAndBid) getValue(withJokers bool) int {
	return typeOfHands[h.getType(withJokers)]
}

func (h HandAndBid) compare(other HandAndBid, withJokers bool) int {
	if h.getValue(withJokers) == other.getValue(withJokers) {
		for i := 0; i < 5; i++ {
			hCard := string(h.hand[i])
			oCard := string(other.hand[i])
			if hCard != oCard {
				if withJokers {
					return cardValuesWithJokers[hCard] - cardValuesWithJokers[oCard]
				} else {
					return cardValues[hCard] - cardValues[oCard]
				}
			}
		}
		return 0
	}
	return h.getValue(withJokers) - other.getValue(withJokers)
}

func getTotalWinnings(hands []HandAndBid, withJokers bool) int {
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].compare(hands[j], withJokers) <= 0
	})
	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += hand.bid * (i + 1)
	}
	return totalWinnings
}

func partOne(hands []HandAndBid) int {
	return getTotalWinnings(hands, false)
}

func partTwo(hands []HandAndBid) int {
	return getTotalWinnings(hands, true)
}

func main() {
	now := time.Now()
	handsAndBids := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(handsAndBids)
	part2 := partTwo(handsAndBids)
	duration := time.Since(now)
	fmt.Printf("Part 1: Total winnings: %d\n", part1)
	fmt.Printf("Part 2: Total winnings with jokers: %d\n", part2)
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
