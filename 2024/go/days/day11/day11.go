package day11

import (
	"aoc/2024/go/lib"
	"fmt"
	"strconv"
	"strings"
)

func countStones(input string, blinks int) int64 {
	stones := make(map[string]int64)

	for _, engraving := range strings.Split(input, " ") {
		stones[engraving] += 1
	}

	for blink := 0; blink < blinks; blink++ {
		newStones := make(map[string]int64)
		for engraving, count := range stones {
			if engraving == "0" {
				newStones["1"] += count
			} else if len(engraving)%2 == 0 {
				length := len(engraving)
				engravingL := engraving[:length/2]
				engravingR := engraving[length/2:]
				numL, _ := strconv.ParseInt(engravingL, 10, 32)
				numR, _ := strconv.ParseInt(engravingR, 10, 32)
				newStones[fmt.Sprintf("%v", numL)] += count
				newStones[fmt.Sprintf("%v", numR)] += count
			} else {
				num, _ := strconv.ParseInt(engraving, 10, 32)
				num *= 2024
				newStones[fmt.Sprintf("%v", num)] += count
			}
		}
		stones = newStones
	}

	sumStones := int64(0)
	for _, count := range stones {
		sumStones += count
	}
	return sumStones
}

func Part1() any {
	input, _ := lib.ReadInput(11)
	return countStones(strings.TrimSpace(input), 25)
}

func Part2() any {
	input, _ := lib.ReadInput(11)
	return countStones(strings.TrimSpace(input), 75)
}
