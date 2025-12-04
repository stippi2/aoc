package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"aoc/2025/go/days/day01"
	"aoc/2025/go/days/day02"
	"aoc/2025/go/days/day03"
	"aoc/2025/go/days/day04"
	"aoc/2025/go/days/day05"
	"aoc/2025/go/days/day06"
	"aoc/2025/go/days/day07"
	"aoc/2025/go/days/day08"
	"aoc/2025/go/days/day09"
	"aoc/2025/go/days/day10"
	"aoc/2025/go/days/day11"
	"aoc/2025/go/days/day12"
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
	case 1:
		fmt.Printf("Part 1: %v\n", day01.Part1())
		fmt.Printf("Part 2: %v\n", day01.Part2())
	case 2:
		fmt.Printf("Part 1: %v\n", day02.Part1())
		fmt.Printf("Part 2: %v\n", day02.Part2())
	case 3:
		fmt.Printf("Part 1: %v\n", day03.Part1())
		fmt.Printf("Part 2: %v\n", day03.Part2())
	case 4:
		fmt.Printf("Part 1: %v\n", day04.Part1())
		fmt.Printf("Part 2: %v\n", day04.Part2())
	case 5:
		fmt.Printf("Part 1: %v\n", day05.Part1())
		fmt.Printf("Part 2: %v\n", day05.Part2())
	case 6:
		fmt.Printf("Part 1: %v\n", day06.Part1())
		fmt.Printf("Part 2: %v\n", day06.Part2())
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
	case 11:
		fmt.Printf("Part 1: %v\n", day11.Part1())
		fmt.Printf("Part 2: %v\n", day11.Part2())
	case 12:
		fmt.Printf("Part 1: %v\n", day12.Part1())
		fmt.Printf("Part 2: %v\n", day12.Part2())

	default:
		fmt.Printf("Day %d is not yet implemented\n", *day)
	}

	fmt.Printf("Time: %v\n", time.Since(start))
}
