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

// O1's suggestion to use Aho-Corasick-based Trie.
// But Claude's much cleaner implementation and didn't contain bugs.

type TrieNode struct {
	children    map[int]*TrieNode // Key is the price change (-9 to +9)
	failureLink *TrieNode         // Where to jump on mismatch
	pattern     []int             // If this node is a pattern end
	isEnd       bool              // Marks pattern end
	output      map[string]int    // Stores position for each found pattern
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[int]*TrieNode),
		output:   make(map[string]int),
	}
}

type AhoCorasick struct {
	root *TrieNode
}

func NewAhoCorasick() *AhoCorasick {
	return &AhoCorasick{
		root: NewTrieNode(),
	}
}

// Adds a pattern to the trie
func (ac *AhoCorasick) AddPattern(pattern []int) {
	node := ac.root
	for _, change := range pattern {
		if _, exists := node.children[change]; !exists {
			node.children[change] = NewTrieNode()
		}
		node = node.children[change]
	}
	node.isEnd = true
	node.pattern = pattern
}

// Builds the failure links
func (ac *AhoCorasick) BuildFailureLinks() {
	queue := []*TrieNode{}

	// For Level 1: Failure links go to root
	for _, child := range ac.root.children {
		child.failureLink = ac.root
		queue = append(queue, child)
	}

	// BFS for the remaining levels
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for change, child := range current.children {
			queue = append(queue, child)

			// Follow the failure link of the current node
			failureNode := current.failureLink
			for failureNode != nil {
				if nextNode, exists := failureNode.children[change]; exists {
					child.failureLink = nextNode
					break
				}
				failureNode = failureNode.failureLink
			}
			if failureNode == nil {
				child.failureLink = ac.root
			}
		}
	}
}

// Searches for all patterns in a change sequence
func (ac *AhoCorasick) Search(changes []int) map[string]int {
	results := make(map[string]int)
	current := ac.root

	for pos, change := range changes {
		// As long as no matching transition is found, follow the failure links
		for current != nil && current.children[change] == nil {
			current = current.failureLink
		}

		if current == nil {
			current = ac.root
			continue
		}

		current = current.children[change]

		// If we are at a pattern end, store the position
		if current.isEnd {
			key := fmt.Sprintf("%v", current.pattern)
			if _, exists := results[key]; !exists {
				results[key] = pos - len(current.pattern) + 1
			}
		}
	}

	return results
}

func findBestSequenceAC(buyers []Buyer) (int, []int) {
	// Create the Aho-Corasick automaton
	ac := NewAhoCorasick()

	// Add all possible 4-sequences
	for d1 := -9; d1 <= 9; d1++ {
		for d2 := -9; d2 <= 9; d2++ {
			for d3 := -9; d3 <= 9; d3++ {
				for d4 := -9; d4 <= 9; d4++ {
					ac.AddPattern([]int{d1, d2, d3, d4})
				}
			}
		}
	}

	// Build the failure links
	ac.BuildFailureLinks()

	// For each sequence we store the sum of prices
	sequenceSums := make(map[string]int)

	// Search through the changes of each buyer
	for _, buyer := range buyers {
		matches := ac.Search(buyer.changes)
		// For each found pattern we add the price
		for pattern, pos := range matches {
			sequenceSums[pattern] += buyer.prices[pos+4]
		}
	}

	// Find the best sequence
	bestSum := 0
	bestSeq := make([]int, 4)
	for pattern, sum := range sequenceSums {
		if sum > bestSum {
			bestSum = sum
			// Parse pattern string back to slice
			fmt.Sscanf(pattern, "[%d %d %d %d]", &bestSeq[0], &bestSeq[1], &bestSeq[2], &bestSeq[3])
		}
	}

	return bestSum, bestSeq
}

func Part2() any {
	input, _ := lib.ReadInput(22)
	buyers := parseInput(input)
	for i := 0; i < len(buyers); i++ {
		buyers[i].generateSequence()
	}
	// mostBananas, bestSequence := findBestSequence(buyers)
	mostBananas, bestSequence := findBestSequenceAC(buyers)
	return fmt.Sprintf("Most bananas: %v (sequence: %v)", mostBananas, bestSequence)
}
