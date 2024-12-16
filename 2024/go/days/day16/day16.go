package day16

import (
	"aoc/2024/go/lib"
	"fmt"
	"math"
	"strings"
)

type Reindeer struct {
	score    int
	facing   string
	position lib.Vec2
	visited  map[lib.Vec2]bool
}

func (r *Reindeer) clone(position lib.Vec2, facing string, grid *lib.Grid) *Reindeer {
	if grid.Get(position.X, position.Y) == '#' {
		return nil
	}
	c := &Reindeer{
		score:    r.score,
		facing:   facing,
		position: position,
		visited:  make(map[lib.Vec2]bool),
	}
	for key, value := range r.visited {
		c.visited[key] = value
	}
	return c
}

func nextPosition(position lib.Vec2, facing string) lib.Vec2 {
	nextPosition := position
	switch facing {
	case "up":
		nextPosition.Y--
	case "down":
		nextPosition.Y++
	case "left":
		nextPosition.X--
	case "right":
		nextPosition.X++
	}
	return nextPosition
}

func (r *Reindeer) forward(grid *lib.Grid) *Reindeer {
	nextPosition := nextPosition(r.position, r.facing)
	if c := r.clone(nextPosition, r.facing, grid); c != nil {
		c.score += 1
		c.visited[r.position] = true
		c.visited[c.position] = true
		return c
	}
	return nil
}

func (r *Reindeer) left(grid *lib.Grid) *Reindeer {
	facing := r.facing
	switch facing {
	case "up":
		facing = "left"
	case "down":
		facing = "right"
	case "left":
		facing = "down"
	case "right":
		facing = "up"
	}
	nextPosition := nextPosition(r.position, facing)
	if c := r.clone(nextPosition, facing, grid); c != nil {
		c.score += 1000
		c.position = r.position
		return c
	}
	return nil
}

func (r *Reindeer) right(grid *lib.Grid) *Reindeer {
	facing := r.facing
	switch facing {
	case "up":
		facing = "right"
	case "down":
		facing = "left"
	case "left":
		facing = "up"
	case "right":
		facing = "down"
	}
	nextPosition := nextPosition(r.position, facing)
	if c := r.clone(nextPosition, facing, grid); c != nil {
		c.score += 1000
		c.position = r.position
		return c
	}
	return nil
}

func findPosition(grid *lib.Grid, value byte) lib.Vec2 {
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			if grid.Get(x, y) == value {
				return lib.Vec2{X: x, Y: y}
			}
		}
	}
	panic("Did not find start position")
}

func findLowestScore(input string, expectScore int) (int, int) {
	grid := lib.NewGrid(input)

	queue := []*Reindeer{
		{
			score:    0,
			facing:   "right",
			position: findPosition(grid, 'S'),
		},
	}

	visited := make(map[string]map[lib.Vec2]*Reindeer)
	for _, facing := range []string{"up", "down", "left", "right"} {
		visited[facing] = make(map[lib.Vec2]*Reindeer)
	}
	visited[queue[0].facing][queue[0].position] = queue[0]

	goal := findPosition(grid, 'E')
	lowestScore := expectScore
	bestPositions := make(map[lib.Vec2]bool)

	for len(queue) > 0 {
		reindeer := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if reindeer.position == goal {
			if reindeer.score < lowestScore {
				for v := range bestPositions {
					delete(bestPositions, v)
				}
				for v := range reindeer.visited {
					bestPositions[v] = true
				}
				fmt.Printf("found goal! score: %d\n", reindeer.score)
				lowestScore = reindeer.score
			} else if reindeer.score == lowestScore {
				fmt.Printf("found goal again! score: %d\n", reindeer.score)
				for v := range reindeer.visited {
					bestPositions[v] = true
				}
			}
		}
		if reindeer.score > lowestScore {
			continue
		}

		nextReindeers := []*Reindeer{
			reindeer.left(grid),
			reindeer.right(grid),
			reindeer.forward(grid),
		}

		for _, nextReindeer := range nextReindeers {
			if nextReindeer != nil && lowestScore >= nextReindeer.score {
				otherReindeer := visited[nextReindeer.facing][nextReindeer.position]
				if otherReindeer == nil || otherReindeer.score >= nextReindeer.score {
					queue = append(queue, nextReindeer)
					visited[nextReindeer.facing][nextReindeer.position] = nextReindeer
				}
			}
		}
	}

	return lowestScore, len(bestPositions)
}

func Part1() any {
	input, _ := lib.ReadInput(16)
	lowestScore, bestPositions := findLowestScore(strings.TrimSpace(input), math.MaxInt)
	return fmt.Sprintf("lowest Score: %d, best positions: %d", lowestScore, bestPositions)
}

func Part2() any {
	return "see above"
}
