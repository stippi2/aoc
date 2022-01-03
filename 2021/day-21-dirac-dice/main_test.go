package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeterministicDice_Roll(t *testing.T) {
	dice := &DeterministicDice{}
	number := dice.Roll()
	assert.Equal(t, 1, dice.index)
	assert.Equal(t, 1, number)

	dice.index = 99
	number = dice.Roll()
	assert.Equal(t, 0, dice.index)
	assert.Equal(t, 100, number)
	assert.Equal(t, 1, dice.Roll())
}

func Test_playGame(t *testing.T) {
	players := []Player{
		{pos: 4},
		{pos: 8},
	}
	dice := &DeterministicDice{}
	winner := playGame(players, dice)
	assert.Equal(t, 0, winner)
	assert.Equal(t, 993, dice.rolls)
	assert.Equal(t, 745, players[1].score)
}

func Test_playGameDiracDice(t *testing.T) {
	p1 := Player{pos: 4}
	p2 := Player{pos: 8}
	wins1, wins2 := countWinningUniverses(p1, p2)
	assert.Equal(t, 444356092776315, wins1)
	assert.Equal(t, 341960390180808, wins2)
}
