package day13

import (
	"aoc/2024/go/lib"
	"fmt"
	"math"
	"strings"
)

type ClawMachine struct {
	buttonA lib.Vec2
	buttonB lib.Vec2
	prize   lib.Vec2
}

func parseInput(input string) []ClawMachine {
	machineStrings := strings.Split(input, "\n\n")
	clawMachines := make([]ClawMachine, len(machineStrings))
	for i, machineString := range machineStrings {
		machine := &clawMachines[i]
		matches, _ := fmt.Sscanf(machineString, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d",
			&machine.buttonA.X, &machine.buttonA.Y,
			&machine.buttonB.X, &machine.buttonB.Y,
			&machine.prize.X, &machine.prize.Y,
		)
		if matches != 6 {
			panic("Failed to parse claw machine")
		}
	}
	return clawMachines
}

func calculateMinTokens(machine ClawMachine) int {
	maxA := lib.Min(
		(machine.prize.X+machine.buttonA.X-1)/machine.buttonA.X,
		(machine.prize.Y+machine.buttonA.Y-1)/machine.buttonA.Y)
	maxB := lib.Min(
		(machine.prize.X+machine.buttonB.X-1)/machine.buttonB.X,
		(machine.prize.Y+machine.buttonB.Y-1)/machine.buttonB.Y)
	type ButtonPresses struct {
		countA int
		countB int
	}
	var solutions []ButtonPresses
	for i := 0; i <= maxA; i++ {
		for j := 0; j <= maxB; j++ {
			if machine.buttonA.X*i+machine.buttonB.X*j == machine.prize.X &&
				machine.buttonA.Y*i+machine.buttonB.Y*j == machine.prize.Y {
				solutions = append(solutions, ButtonPresses{countA: i, countB: j})
			}
		}
	}
	if len(solutions) == 0 {
		return 0
	}
	minTokens := math.MaxInt
	for _, solution := range solutions {
		tokens := solution.countA*3 + solution.countB
		if tokens < minTokens {
			minTokens = tokens
		}
	}
	return minTokens
}

func calculateMinTokensForMaxPrizes(input string) int {
	machines := parseInput(input)
	sumTokens := 0
	for _, machine := range machines {
		sumTokens += calculateMinTokens(machine)
	}
	return sumTokens
}

func Part1() any {
	input, _ := lib.ReadInput(13)
	return calculateMinTokensForMaxPrizes(input)
}

func Part2() any {
	return "Not implemented"
}
