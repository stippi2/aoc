package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func windowContainsAllDifferentChars(sequence string, start, windowSize int) bool {
	set := make(map[uint8]bool)
	for c := 0; c < windowSize; c++ {
		set[sequence[start+c]] = true
	}
	return len(set) == windowSize
}

func findWindowOfDifferentChars(sequence string, windowSize int) int {
	for pos := 0; pos < len(sequence)-windowSize; pos++ {
		if windowContainsAllDifferentChars(sequence, pos, windowSize) {
			return pos + windowSize
		}
	}
	return -1
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

func findWindowOfDifferentCharsGoFuncs(sequence string, windowSize int) int {
	result := 0
	var wg sync.WaitGroup
	for pos := 0; pos < len(sequence)-windowSize; pos++ {
		wg.Add(1)
		go func(pos int) {
			defer wg.Done()
			if windowContainsAllDifferentChars(sequence, pos, windowSize) {
				if result == 0 || pos+windowSize < result {
					result = pos + windowSize
				}
			}
		}(pos)
		if result != 0 {
			break
		}
	}
	wg.Wait()
	return result
}

func main() {
	input := loadInput("puzzle-input.txt")
	fmt.Printf("pos after start marker: %v\n", findDifferentCharsQuick(input, 4))

	// Dry run for spinning up the OS threads
	findWindowOfDifferentCharsGoFuncs(input, 14)

	startSlowVersion := time.Now()
	posPartSlow := findWindowOfDifferentChars(input, 14)
	startQuickVersion := time.Now()
	posPartQuick := findDifferentCharsQuick(input, 14)
	startParallelVersion := time.Now()
	posPartParallel := findWindowOfDifferentCharsGoFuncs(input, 14)
	allDone := time.Now()
	fmt.Printf("pos after message: %v/%v/%v, slow version: %v, quick version: %v, parallel version: %v\n",
		posPartSlow, posPartQuick, posPartParallel,
		startQuickVersion.Sub(startSlowVersion),
		startParallelVersion.Sub(startQuickVersion),
		allDone.Sub(startParallelVersion))
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
