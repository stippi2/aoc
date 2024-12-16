package day16

import (
	"aoc/2024/go/lib"
	"math"
	"strings"
)

type Reindeer struct {
	score    int
	facing   string
	position lib.Vec2
	visited  map[lib.Vec2]bool
}

func (r *Reindeer) clone() *Reindeer {
	c := &Reindeer{
		score:    r.score,
		facing:   r.facing,
		position: r.position,
		visited:  make(map[lib.Vec2]bool),
	}
	for key, value := range r.visited {
		c.visited[key] = value
	}
	return c
}

func (r *Reindeer) forward() *Reindeer {
	c := r.clone()
	switch r.facing {
	case "up":
		c.position.Y--
	case "down":
		c.position.Y++
	case "left":
		c.position.X--
	case "right":
		c.position.X++
	}
	c.score += 1
	c.visited[r.position] = true
	return c
}

func (r *Reindeer) left() *Reindeer {
	c := r.clone()
	switch r.facing {
	case "up":
		c.facing = "left"
	case "down":
		c.facing = "right"
	case "left":
		c.facing = "down"
	case "right":
		c.facing = "up"
	}
	c.score += 1000
	return c
}

func (r *Reindeer) right() *Reindeer {
	c := r.clone()
	switch r.facing {
	case "up":
		c.facing = "right"
	case "down":
		c.facing = "left"
	case "left":
		c.facing = "up"
	case "right":
		c.facing = "down"
	}
	c.score += 1000
	return c
}

func findStart(grid *lib.Grid) lib.Vec2 {
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			if grid.Get(x, y) == 'S' {
				return lib.Vec2{X: x, Y: y}
			}
		}
	}
	panic("Did not find start position")
}

func findLowestScore(input string) int {
	grid := lib.NewGrid(input)
	queue := []*Reindeer{
		{
			score:    0,
			facing:   "right",
			position: findStart(grid),
			visited:  make(map[lib.Vec2]bool),
		},
	}
	lowestScore := math.MinInt

	for len(queue) > 0 {

	}

	return lowestScore
}

func Part1() any {
	input, _ := lib.ReadInput(16)
	return findLowestScore(strings.TrimSpace(input))
}

func Part2() any {
	return "Not implemented"
}
