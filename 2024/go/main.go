package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"aoc/2024/go/days/day07"
)

func main() {
	day := flag.Int("day", 0, "Day of the Advent of Code")
	flag.Parse()

	if *day == 0 {
		log.Fatal("Please provide a day (--day=X)")
	}

	start := time.Now()

	fmt.Printf("Day %02d\n", *day)
	switch *day {
	case 7:
		fmt.Printf("Part 1: %v\n", day07.Part1())
		fmt.Printf("Part 2: %v\n", day07.Part2())
	default:
		fmt.Printf("Tag %d ist noch nicht implementiert\n", *day)
	}

	fmt.Printf("Zeit: %v\n", time.Since(start))
}
