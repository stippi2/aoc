package day01

import (
	"aoc/2025/go/lib"
	"strconv"
	"strings"
)

func parseLine(line string) int {
	direction := -1
	switch line[0] {
	case 'L':
		direction = -1
	case 'R':
		direction = 1
	default:
		panic("Unexpected direction")
	}
	count, err := strconv.Atoi(line[1:])
	if err != nil {
		panic("Failed to parse count")
	}
	return direction * count
}

func countZeros(input string, countCrossings bool) int {
	lines := strings.Split(input, "\n")
	dial := 50
	password := 0
	for _, line := range lines {
		count := parseLine(line)
		threeSixties := count / 100
		if countCrossings {
			password += lib.Abs(threeSixties)
		}
		count -= threeSixties * 100
		newDial := dial + count
		if newDial > 99 {
			newDial -= 100
			if countCrossings && newDial != 0 && dial != 0 {
				password++
			}
		} else if newDial < 0 {
			newDial += 100
			if countCrossings && newDial != 0 && dial != 0 {
				password++
			}
		}
		if newDial == 0 {
			password++
		}

		dial = newDial
	}

	return password
}

func Part1() any {
	input, _ := lib.ReadInput(1)
	return countZeros(input, false)
}

func Part2() any {
	input, _ := lib.ReadInput(1)
	return countZeros(input, true)
}
