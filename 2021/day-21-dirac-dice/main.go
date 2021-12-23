package main

import "fmt"

type DeterministicDice struct {
	index int
	rolls int
}

func (d *DeterministicDice) Roll() int {
	d.rolls++
	value := d.index + 1
	d.index = value % 100
	return value
}

type Player struct {
	pos int
	score int
}

type Dice interface {
	Roll() int
}

func (p *Player) turn(d Dice) {
	steps := d.Roll() + d.Roll() + d.Roll()
	p.pos = 1 + (p.pos + steps - 1) % 10
	p.score += p.pos
}

func playGame(players []Player, d Dice) (winningPlayer int) {
	for {
		for i := 0; i < len(players); i++ {
			players[i].turn(d)
			if players[i].score >= 1000 {
				return i
			}
		}
	}
}

func main() {
	// Part 1
	players := []Player{
		{pos: 4},
		{pos: 1},
	}
	dice := &DeterministicDice{}
	winner := playGame(players, dice)
	loser := (winner + 1) % 2
	fmt.Printf("loosing player: %v, score: %v, turns: %v, result: %v\n",
		loser, players[loser].score, dice.rolls, players[loser].score * dice.rolls)
}
