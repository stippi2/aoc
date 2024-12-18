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

func calculateMinTokensSmart(machine ClawMachine) int {
	// D = A_x*B_y - A_y*B_x
	determinant := machine.buttonA.X*machine.buttonB.Y - machine.buttonA.Y*machine.buttonB.X
	if determinant == 0 {
		fmt.Printf("Determinant is 0\n")
		return 0
	}
	// a = (X_p*B_y - Y_p*B_x)/D
	// b = (A_x*Y_p - A_y*X_p)/D
	countA := float64(machine.prize.X*machine.buttonB.Y-machine.prize.Y*machine.buttonB.X) / float64(determinant)
	countB := float64(machine.buttonA.X*machine.prize.Y-machine.buttonA.Y*machine.prize.X) / float64(determinant)
	if math.Round(countA) != countA || math.Round(countB) != countB {
		fmt.Printf("Non integer result A: %.4f, B: %.4f\n", countA, countB)
		return 0
	}
	if countA < 0 || countB < 0 {
		fmt.Printf("Negative results A: %.4f, B: %.4f\n", countA, countB)
		return 0
	}
	return int(countA)*3 + int(countB)
}

func calculateMinTokensForMaxPrizes(input string, prizeOffset int) int {
	machines := parseInput(input)
	sumTokens := 0
	for _, machine := range machines {
		machine.prize.X += prizeOffset
		machine.prize.Y += prizeOffset
		sumTokens += calculateMinTokensSmart(machine)
	}
	return sumTokens
}

func Part1() any {
	input, _ := lib.ReadInput(13)
	return calculateMinTokensForMaxPrizes(input, 0)
}

func Part2() any {
	input, _ := lib.ReadInput(13)
	return calculateMinTokensForMaxPrizes(input, 10000000000000)
}
