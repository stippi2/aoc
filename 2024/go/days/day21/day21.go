package day21

import (
	"aoc/2024/go/lib"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Keypad struct {
	positions map[lib.Vec2]rune
	position  lib.Vec2
}

func newNumberKeypad() *Keypad {
	return &Keypad{
		positions: map[lib.Vec2]rune{
			{X: 0, Y: 0}: '7',
			{X: 1, Y: 0}: '8',
			{X: 2, Y: 0}: '9',

			{X: 0, Y: 1}: '4',
			{X: 1, Y: 1}: '5',
			{X: 2, Y: 1}: '6',

			{X: 0, Y: 2}: '1',
			{X: 1, Y: 2}: '2',
			{X: 2, Y: 2}: '3',

			{X: 1, Y: 3}: '0',
			{X: 2, Y: 3}: 'A',
		},
		position: lib.Vec2{X: 2, Y: 3},
	}
}

func newDirectionalKeypad() *Keypad {
	return &Keypad{
		positions: map[lib.Vec2]rune{
			{X: 1, Y: 0}: '^',
			{X: 2, Y: 0}: 'A',

			{X: 0, Y: 1}: '<',
			{X: 1, Y: 1}: 'v',
			{X: 2, Y: 1}: '>',
		},
		position: lib.Vec2{X: 2, Y: 0},
	}
}

func (k *Keypad) pressButton() rune {
	return k.positions[k.position]
}

func (k *Keypad) getPossibleMoves(position lib.Vec2) []lib.Vec2 {
	moves := []lib.Vec2{
		{X: position.X - 1, Y: position.Y},
		{X: position.X + 1, Y: position.Y},
		{X: position.X, Y: position.Y - 1},
		{X: position.X, Y: position.Y + 1},
	}
	var possibleMoves []lib.Vec2
	for _, move := range moves {
		if _, exists := k.positions[move]; exists {
			possibleMoves = append(possibleMoves, move)
		}
	}
	return possibleMoves
}

func (k *Keypad) move(direction rune) {
	switch direction {
	case '<':
		k.position.X--
	case '>':
		k.position.X++
	case '^':
		k.position.Y--
	case 'v':
		k.position.Y++
	}
}

func (k *Keypad) executeSequence(sequence string) {
	for _, action := range sequence {
		if action == 'A' {
			k.pressButton()
		} else {
			k.move(action)
		}
	}
}

// getDirection returns the button press needed to move from pos1 to pos2
func getDirection(pos1, pos2 lib.Vec2) string {
	if pos2.Y < pos1.Y {
		return "^"
	}
	if pos2.Y > pos1.Y {
		return "v"
	}
	if pos2.X < pos1.X {
		return "<"
	}
	return ">"
}

// State represents a position in our search space
type State struct {
	position lib.Vec2
	sequence string
	cost     int // Track the total cost considering directional keypads
}

// findShortestPath finds shortest sequence including the button press (A)
func (k *Keypad) findShortestPath(targetButton rune) string {
	visited := make(map[lib.Vec2]int) // map[position]lowestCostSeen
	queue := []State{{
		position: k.position,
		sequence: "",
		cost:     0,
	}}

	bestPath := ""
	lowestCost := math.MaxInt32

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// If we've seen this position with a lower cost, skip it
		if prevCost, seen := visited[current.position]; seen && prevCost < current.cost {
			continue
		}
		visited[current.position] = current.cost

		// Found target - update if it's cheaper
		if k.positions[current.position] == targetButton {
			if current.cost < lowestCost {
				lowestCost = current.cost
				bestPath = current.sequence + "A"
			}
			continue
		}

		// Add neighbors with updated costs
		for _, nextPos := range k.getPossibleMoves(current.position) {
			newCost := current.cost + 1 // Basic movement cost
			// If we're changing direction, add extra cost
			if len(current.sequence) > 0 &&
				getDirection(current.position, nextPos)[0] != current.sequence[len(current.sequence)-1] {
				newCost += 2 // Penalty for changing direction
			}

			queue = append(queue, State{
				position: nextPos,
				sequence: current.sequence + getDirection(current.position, nextPos),
				cost:     newCost,
			})
		}
	}

	if bestPath == "" {
		panic("no path found")
	}
	return bestPath
}

func findShortestSequence(keyCode string) int {
	outerPad := newDirectionalKeypad()
	middlePad := newDirectionalKeypad()
	innerPad := newNumberKeypad()

	shortestSequence := ""

	fmt.Printf("Key code %v\n", keyCode)

	for _, targetButton := range keyCode {
		// Find path to number/letter on numeric keypad
		numericPath := innerPad.findShortestPath(targetButton)
		innerPad.executeSequence(numericPath)
		fmt.Printf("  innerPad sequence %v\n", numericPath)

		// For each action needed on numeric keypad
		for _, press := range numericPath {
			middlePath := middlePad.findShortestPath(press)
			middlePad.executeSequence(middlePath)
			fmt.Printf("  %c: middlePad sequence %v\n", press, middlePath)

			for _, middlePress := range middlePath {
				outerPath := outerPad.findShortestPath(middlePress)
				outerPad.executeSequence(outerPath)
				fmt.Printf("    %c: outerPad sequence %v\n", middlePress, outerPath)
				shortestSequence += outerPath
			}
		}
	}

	fmt.Printf("Key code %v: %s\n", keyCode, shortestSequence)

	return len(shortestSequence)
}

func calculateComplexity(input string) int {
	keyCodes := strings.Split(input, "\n")
	sum := 0
	for _, keyCode := range keyCodes {
		sequenceLength := findShortestSequence(keyCode)
		numeric, _ := strconv.Atoi(strings.TrimSuffix(keyCode, "A"))
		sum += sequenceLength * numeric
	}

	return sum
}

func Part1() any {
	input, _ := lib.ReadInput(21)
	return calculateComplexity(input)
}

func Part2() any {
	return "Not implemented"
}
