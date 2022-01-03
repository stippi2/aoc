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

// 1,1,1 -> 3 1

// 2,1,1 -> 4 3
// 1,2,1 -> 4
// 1,1,2 -> 4

// 2,2,1 -> 5 6
// 1,2,2 -> 5
// 2,1,2 -> 5
// 1,1,3 -> 5
// 1,3,1 -> 5
// 3,1,1 -> 5

// 2,2,2 -> 6 7
// 1,2,3 -> 6
// 2,1,3 -> 6
// 2,3,1 -> 6
// 1,3,2 -> 6
// 3,1,2 -> 6
// 3,2,1 -> 6

// 2,2,3 -> 7 6
// 3,2,2 -> 7
// 2,3,2 -> 7
// 3,3,1 -> 7
// 1,3,3 -> 7
// 3,1,3 -> 7

// 2,3,3 -> 8 3
// 3,2,3 -> 8
// 3,3,2 -> 8

// 3,3,3 -> 9 1

func countWinningUniverses(p1, p2 Player) (int, int) {
	if p2.score >= 21 {
		return 0, 1
	}

	rollFrequencies := []struct {
		roll      int
		frequency int
	}{
		{3,1},
		{4,3},
		{5,6},
		{6,7},
		{7,6},
		{8,3},
		{9,1},
	}

	universes1 := 0
	universes2 := 0
	for _, rf := range rollFrequencies {
		// Recurse with swapped players
		pos := 1 + (p1.pos + rf.roll - 1) % 10
		movedP1 := Player{
			pos:   pos,
			score: p1.score + pos,
		}
		u2, u1 := countWinningUniverses(p2, movedP1)
		universes1 += rf.frequency * u1
		universes2 += rf.frequency * u2
	}

	return universes1, universes2
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

	// Part 2
	p1 := Player{pos: 4}
	p2 := Player{pos: 1}
	wins1, wins2 := countWinningUniverses(p1, p2)
	if wins1 > wins2 {
		fmt.Printf("player 1 wins in more universes: %d, player 2 in %d\n", wins1, wins2)
	} else {
		fmt.Printf("player 2 wins in more universes: %d, player 1 in %d\n", wins2, wins1)
	}
}
