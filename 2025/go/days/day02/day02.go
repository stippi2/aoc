package day02

import (
	"aoc/2025/go/lib"
	"strconv"
	"strings"
)

var pow10Table = []int{
	1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000,
	1000000000, 10000000000, 100000000000, 1000000000000,
	10000000000000, 100000000000000, 1000000000000000,
	10000000000000000, 100000000000000000, 1000000000000000000,
}

func pow10(n int) int {
	if n < len(pow10Table) {
		return pow10Table[n]
	}
	result := pow10Table[len(pow10Table)-1]
	for i := len(pow10Table) - 1; i < n; i++ {
		result *= 10
	}
	return result
}

func digitCount(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func sumRepeatNumbers(lo, hi int) int {
	sum := 0

	digits := digitCount(lo)
	if digits%2 == 1 {
		digits++
	}

	for {
		halfDigits := digits / 2
		multiplier := pow10(halfDigits)

		minHalf := pow10(halfDigits - 1)
		if halfDigits == 1 {
			minHalf = 1
		}

		half := minHalf
		if digits == digitCount(lo) {
			half = lo / multiplier
		}

		for {
			num := half*multiplier + half

			if num > hi {
				return sum
			}
			if num >= lo {
				sum += num
			}
			half++

			if half >= multiplier {
				break
			}
		}

		digits += 2
	}
}

func sumInvalidIds(input string) int {
	ranges := strings.Split(input, ",")
	sum := 0
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		min, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		max, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		sum += sumRepeatNumbers(min, max)
	}
	return sum
}

func Part1() any {
	input, _ := lib.ReadInput(2)
	return sumInvalidIds(input)
}

func Part2() any {
	return 0
}
