package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `A Y
B X
C Z`

func Test_parseInput(t *testing.T) {
	playerA, playerB := parseInput(exampleInput)
	assert.Len(t, playerA.moves, 3)
	assert.Len(t, playerB.moves, 3)
	assert.Equal(t, playerA.moves, []string{"A", "B", "C"})
	assert.Equal(t, playerB.moves, []string{"Y", "X", "Z"})
}

func Test_playRound(t *testing.T) {
	playerA, playerB := parseInput(exampleInput)
	playerA.mapping = Mapping{
		rock:     "A",
		paper:    "B",
		scissors: "C",
	}
	playerB.mapping = playerA.mapping

	playRound(&playerA, &playerB, 0)
	assert.Equal(t, playerB.score, 4)

	playRound(&playerA, &playerB, 1)
	assert.Equal(t, playerB.score, 4+1)

	playRound(&playerA, &playerB, 2)
	assert.Equal(t, playerB.score, 4+1+7)
}
