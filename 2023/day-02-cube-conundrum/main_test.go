package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func Test_part1(t *testing.T) {
	games := parseInput(input)
	maxCubes := CubeSet{red: 12, green: 13, blue: 14}
	possibleGames := getPossibleGames(games, maxCubes)
	assert.Equal(t, 8, sumIds(possibleGames))
}

func Test_part2(t *testing.T) {
	games := parseInput(input)
	sumOfPowers := 0
	for _, game := range games {
		sumOfPowers += game.requiredCubes().power()
	}
	assert.Equal(t, 2286, sumOfPowers)
}
