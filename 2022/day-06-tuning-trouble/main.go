package main

import (
	"fmt"
	"os"
)

func findWindowOfDifferentChars(sequence string, windowSize int) int {
	pos := 0
	for pos < len(sequence)-windowSize {
		set := make(map[uint8]bool)
		for c := 0; c < windowSize; c++ {
			set[sequence[pos+c]] = true
		}
		if len(set) == windowSize {
			return pos + windowSize
		}
		pos++
	}
	return pos
}

func findDifferentCharsQuick(sequence string, windowSize int) int {
	set := make(map[uint8]int)
	for pos := 0; pos < len(sequence); pos++ {
		stored := set[sequence[pos]]
		set[sequence[pos]] = stored + 1
		if pos >= windowSize {
			windowStart := set[sequence[pos-windowSize]]
			if windowStart == 1 {
				delete(set, sequence[pos-windowSize])
			} else {
				set[sequence[pos-windowSize]] = windowStart - 1
			}
		}
		if len(set) == windowSize {
			return pos + 1
		}
	}
	return -1
}

func main() {
	input := loadInput("puzzle-input.txt")
	fmt.Printf("pos after start marker: %v\n", findDifferentCharsQuick(input, 4))
	fmt.Printf("pos after message: %v\n", findDifferentCharsQuick(input, 14))
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
