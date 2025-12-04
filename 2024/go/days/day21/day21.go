package day21

import (
	"aoc/2024/go/lib"
	"fmt"
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

func getControlSequence(sequence string, numControlPads int, optimalPaths map[string]string) string {
	var result strings.Builder
	currentPos := 'A' // Always start at A

	for _, move := range sequence {
		// Skip A since we're already there
		if move == 'A' {
			result.WriteString("A")
			continue
		}

		// Get path to the direction button
		key := fmt.Sprintf("%c-%c", currentPos, move)
		path, exists := optimalPaths[key]
		if !exists {
			panic(fmt.Sprintf("No optimal path for %s", key))
		}
		result.WriteString(path)
		result.WriteString("A") // Press the direction button

		// After pressing any button, we're back at A
		currentPos = 'A'
	}

	// For multiple control pads, we need to compute the control sequence
	// for the resulting sequence again, numControlPads-1 times
	if numControlPads > 1 {
		return getControlSequence(result.String(), numControlPads-1, optimalPaths)
	}

	return result.String()
}

func findBetterSequence(seq1, seq2 string, optimalPaths map[string]string) string {
	// Get the control sequence needed for both options
	controlSeq1 := getControlSequence(seq1, 1, optimalPaths) // We only need to check one control pad
	controlSeq2 := getControlSequence(seq2, 1, optimalPaths)

	if len(controlSeq1) < len(controlSeq2) {
		return seq1
	}
	return seq2
}

func buildOptimalPathMap() map[string]string {
	// Map for all possible paths between buttons
	//     +---+---+
	//     | ^ | A |
	// +---+---+---+
	// | < | v | > |
	// +---+---+---+
	candidatePaths := map[string][]string{
		// From ^
		"^-A": {">"},        // ^ to A
		"^-v": {"v"},        // ^ to v
		"^-<": {"<v"},       // ^ to <
		"^->": {">v", "v>"}, // ^ to >

		// From A
		"A-^": {"<"},        // A to ^
		"A-v": {"v<", "<v"}, // A to v
		"A-<": {"v<<"},      // A to <
		"A->": {"v"},        // A to >

		// From
		"<-^": {">^"},  // < to ^
		"<-A": {">>^"}, // < to A
		"<-v": {">"},   // < to v
		"<->": {">>"},  // < to >

		// From v
		"v-^": {"^"},        // v to ^
		"v-A": {"^>", ">^"}, // v to A
		"v-<": {"<"},        // v to <
		"v->": {">"},        // v to >

		// From >
		">-^": {"^<", "<^"}, // > to ^
		">-A": {"^"},        // > to A
		">-v": {"<"},        // > to v
		">-<": {"<<"},       // > to <
	}

	optimalPaths := make(map[string]string)

	// For each connection, either take the single sequence or find the better one
	for conn, sequences := range candidatePaths {
		if len(sequences) == 1 {
			optimalPaths[conn] = sequences[0]
		} else {
			optimalPaths[conn] = findBetterSequence(sequences[0], sequences[1], optimalPaths)
		}
	}

	return optimalPaths
}

func findShortestSequence(keyCode string, padCount int) int {
	return 0
}

func calculateComplexity(input string, numberOfPads int) int {
	keyCodes := strings.Split(input, "\n")
	sum := 0
	for _, keyCode := range keyCodes {
		sequenceLength := findShortestSequence(keyCode, numberOfPads)
		numeric, _ := strconv.Atoi(strings.TrimSuffix(keyCode, "A"))
		sum += sequenceLength * numeric
	}
	return sum
}

func Part1() any {
	input, _ := lib.ReadInput(21)
	return calculateComplexity(input, 2)
}

func Part2() any {
	input, _ := lib.ReadInput(21)
	return calculateComplexity(input, 25)
}
