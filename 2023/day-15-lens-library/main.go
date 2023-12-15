package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func hash(s string) int {
	h := 0
	for _, c := range s {
		h = ((h + int(c)) * 17) % 256
	}
	return h
}

func partOne(sequence []string) int {
	sum := 0
	for _, s := range sequence {
		sum += hash(s)
	}
	return sum
}

func partTwo(sequence []string) int {
	lenses := make(map[string]int)
	boxes := make(map[int][]string)
	for _, s := range sequence {
		if strings.Contains(s, "-") {
			parts := strings.Split(s, "-")
			index := hash(parts[0])
			labels := boxes[index]
			for i, label := range labels {
				if label == parts[0] {
					labels = append(labels[:i], labels[i+1:]...)
					boxes[index] = labels
					break
				}
			}
		} else {
			parts := strings.Split(s, "=")
			index := hash(parts[0])
			focalLength, _ := strconv.Atoi(parts[1])
			lenses[parts[0]] = focalLength
			labels := boxes[index]
			found := false
			for _, label := range labels {
				if label == parts[0] {
					boxes[index] = labels
					found = true
					break
				}
			}
			if !found {
				boxes[index] = append(labels, parts[0])
			}
		}
	}
	sum := 0
	for label, focalLength := range lenses {
		boxIndex := hash(label)
		for indexInBox, otherLabel := range boxes[boxIndex] {
			if otherLabel == label {
				sum += (1 + boxIndex) * (indexInBox + 1) * focalLength
				break
			}
		}
	}
	return sum
}

func main() {
	now := time.Now()
	sequence := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(sequence)
	part2 := partTwo(sequence)
	duration := time.Since(now)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseInput(input string) []string {
	return strings.Split(input, ",")
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
