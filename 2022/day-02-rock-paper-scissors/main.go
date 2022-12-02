package main

import (
	"fmt"
	"os"
	"strings"
)

type Mapping struct {
	rock     string
	paper    string
	scissors string
}

type Player struct {
	mapping Mapping
	moves   []string
	score   int
}

func (p *Player) appendMove(m string) {
	p.moves = append(p.moves, m)
}

func (m *Mapping) score(move string) int {
	switch move {
	case m.rock:
		return 1
	case m.paper:
		return 2
	case m.scissors:
		return 3
	}
	return 0
}

func wins(mappingA, mappingB Mapping, moveA, moveB string) bool {
	switch moveA {
	case mappingA.rock:
		return moveB == mappingB.scissors
	case mappingA.paper:
		return moveB == mappingB.rock
	case mappingA.scissors:
		return moveB == mappingB.paper
	}
	return false
}

func playRound(playerA, playerB *Player, round int) {
	moveA := playerA.moves[round]
	moveB := playerB.moves[round]
	scoreA := playerA.mapping.score(moveA)
	scoreB := playerB.mapping.score(moveB)
	playerA.score += scoreA
	playerB.score += scoreB
	if scoreA == scoreB {
		playerA.score += 3
		playerB.score += 3
		return
	}
	if wins(playerA.mapping, playerB.mapping, moveA, moveB) {
		playerA.score += 6
	} else {
		playerB.score += 6
	}
}

func main() {
	playerA, playerB := parseInput(loadInput("puzzle-input.txt"))
	playerA.mapping = Mapping{
		rock:     "A",
		paper:    "B",
		scissors: "C",
	}
	playerB.mapping = playerA.mapping
	for round := 0; round < len(playerA.moves); round++ {
		playRound(&playerA, &playerB, round)
	}
	fmt.Printf("total score player A: %v, player B: %v\n", playerA.score, playerB.score)
}

func parseInput(input string) (playerA, playerB Player) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		moves := strings.Split(line, " ")
		playerA.appendMove(moves[0])
		switch moves[1] {
		case "X": // lose
			switch moves[0] {
			case "A":
				playerB.appendMove("C")
			case "B":
				playerB.appendMove("A")
			case "C":
				playerB.appendMove("B")
			}
		case "Y": // draw
			playerB.appendMove(moves[0])
			break
		case "Z": // win
			switch moves[0] {
			case "A":
				playerB.appendMove("B")
			case "B":
				playerB.appendMove("C")
			case "C":
				playerB.appendMove("A")
			}
		}
	}
	return
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
