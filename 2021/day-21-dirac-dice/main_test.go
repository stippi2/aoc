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
