package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"aoc/2024/go/days/day06"
	"aoc/2024/go/days/day07"
	"aoc/2024/go/days/day08"
	"aoc/2024/go/days/day09"
	"aoc/2024/go/days/day10"
	"aoc/2024/go/days/day11"
	"aoc/2024/go/days/day12"
	"aoc/2024/go/days/day13"
	"aoc/2024/go/days/day14"
	"aoc/2024/go/days/day15"
	"aoc/2024/go/days/day16"
	"aoc/2024/go/days/day17"
	"aoc/2024/go/days/day18"
	"aoc/2024/go/days/day19"
	"aoc/2024/go/days/day20"
	"aoc/2024/go/days/day21"
	"aoc/2024/go/days/day22"
	"aoc/2024/go/days/day23"
	"aoc/2024/go/days/day24"
	"aoc/2024/go/days/day25"
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
	case 13:
		fmt.Printf("Part 1: %v\n", day13.Part1())
		fmt.Printf("Part 2: %v\n", day13.Part2())
	case 14:
		fmt.Printf("Part 1: %v\n", day14.Part1())
		fmt.Printf("Part 2: %v\n", day14.Part2())
	case 15:
		fmt.Printf("Part 1: %v\n", day15.Part1())
		fmt.Printf("Part 2: %v\n", day15.Part2())
	case 16:
		fmt.Printf("Part 1: %v\n", day16.Part1())
		fmt.Printf("Part 2: %v\n", day16.Part2())
	case 17:
		fmt.Printf("Part 1: %v\n", day17.Part1())
		fmt.Printf("Part 2: %v\n", day17.Part2())
	case 18:
		fmt.Printf("Part 1: %v\n", day18.Part1())
		fmt.Printf("Part 2: %v\n", day18.Part2())
	case 19:
		fmt.Printf("Part 1: %v\n", day19.Part1())
		fmt.Printf("Part 2: %v\n", day19.Part2())
	case 20:
		fmt.Printf("Part 1: %v\n", day20.Part1())
		fmt.Printf("Part 2: %v\n", day20.Part2())
	case 21:
		fmt.Printf("Part 1: %v\n", day21.Part1())
		fmt.Printf("Part 2: %v\n", day21.Part2())
	case 22:
		fmt.Printf("Part 1: %v\n", day22.Part1())
		fmt.Printf("Part 2: %v\n", day22.Part2())
	case 23:
		fmt.Printf("Part 1: %v\n", day23.Part1())
		fmt.Printf("Part 2: %v\n", day23.Part2())
	case 24:
		fmt.Printf("Part 1: %v\n", day24.Part1())
		fmt.Printf("Part 2: %v\n", day24.Part2())
	case 25:
		fmt.Printf("Part 1: %v\n", day25.Part1())
		fmt.Printf("Part 2: %v\n", day25.Part2())
	default:
		fmt.Printf("Day %d is not yet implemented\n", *day)
	}

	fmt.Printf("Time: %v\n", time.Since(start))
}
