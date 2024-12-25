package day25

import (
	"aoc/2024/go/lib"
	"strings"
)

type Lock struct {
	pins [5]int
}

type Key struct {
	pins [5]int
}

func parseInput(input string) ([]*Lock, []*Key) {
	var locks []*Lock
	var keys []*Key

	for _, grid := range strings.Split(input, "\n\n") {
		lines := strings.Split(grid, "\n")
		if lines[0] == "....." {
			key := &Key{}
			for y := 1; y < 6; y++ {
				for x := 0; x < 5; x++ {
					if lines[y][x] == '#' {
						key.pins[x] = lib.Max(key.pins[x], 6-y)
					}
				}
			}
			keys = append(keys, key)
		} else {
			lock := &Lock{}
			for y := 1; y < 6; y++ {
				for x := 0; x < 5; x++ {
					if lines[y][x] == '#' {
						lock.pins[x] = y
					}
				}
			}
			locks = append(locks, lock)
		}
	}

	return locks, keys
}

func keyFitsLock(key *Key, lock *Lock) bool {
	for x := 0; x < 5; x++ {
		if lock.pins[x]+key.pins[x] > 5 {
			return false
		}
	}
	return true
}

func countFittingLockKeyCombinations(locks []*Lock, keys []*Key) int {
	sum := 0
	for _, key := range keys {
		for _, lock := range locks {
			if keyFitsLock(key, lock) {
				sum++
			}
		}
	}

	return sum
}

func Part1() any {
	input, _ := lib.ReadInput(25)
	locks, keys := parseInput(input)
	return countFittingLockKeyCombinations(locks, keys)
}

func Part2() any {
	return "Not implemented"
}
