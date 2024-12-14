package day14

import (
	"aoc/2024/go/lib"
	"fmt"
	"strings"
)

type robot struct {
	pos lib.Vec2
	vel lib.Vec2
}

func (r *robot) simulate(width, height, seconds int) {
	r.pos.X = (r.pos.X + r.vel.X*seconds) % width
	r.pos.Y = (r.pos.Y + r.vel.Y*seconds) % height
	if r.pos.X < 0 {
		r.pos.X += width
	}
	if r.pos.Y < 0 {
		r.pos.Y += height
	}
}

func computeSafetyFactor(input string, width, height, seconds int) int {
	quadrantLT := 0
	quadrantRT := 0
	quadrantRB := 0
	quadrantLB := 0
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var r robot
		matches, _ := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.pos.X, &r.pos.Y, &r.vel.X, &r.vel.Y)
		if matches != 4 {
			panic("Failed to parse robot")

		}
		r.simulate(width, height, seconds)
		if r.pos.X < width/2 {
			if r.pos.Y < height/2 {
				quadrantLT++
			} else if r.pos.Y > height/2 {
				quadrantLB++
			}
		} else if r.pos.X > width/2 {
			if r.pos.Y < height/2 {
				quadrantRT++
			} else if r.pos.Y > height/2 {
				quadrantRB++
			}
		}
	}

	return quadrantLT * quadrantRT * quadrantRB * quadrantLB
}

func Part1() any {
	input, _ := lib.ReadInput(14)
	return computeSafetyFactor(strings.TrimSpace(input), 101, 103, 100)
}

func Part2() any {
	return "Not implemented"
}
