package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"aoc/2024/go/days/day07"
	"aoc/2024/go/days/day08"
	"aoc/2024/go/days/day09"
	"aoc/2024/go/days/day10"
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
	case 8:
		fmt.Printf("Part 1: %v\n", day08.Part1())
		fmt.Printf("Part 2: %v\n", day08.Part2())
	case 9:
		fmt.Printf("Part 1: %v\n", day09.Part1())
		fmt.Printf("Part 2: %v\n", day09.Part2())
	case 10:
		fmt.Printf("Part 1: %v\n", day10.Part1())
		fmt.Printf("Part 2: %v\n", day10.Part2())
	default:
		fmt.Printf("Day %d is not yet implemented\n", *day)
	}

	fmt.Printf("Time: %v\n", time.Since(start))
}
