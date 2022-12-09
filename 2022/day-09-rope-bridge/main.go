package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

func (p *Pos) track(to Pos) {
	diffX := to.x - p.x
	diffY := to.y - p.y
	if math.Abs(float64(diffX)) <= 1 && math.Abs(float64(diffY)) <= 1 {
		return
	}

	if diffX > 0 {
		p.x++
	} else if diffX < 0 {
		p.x--
	}
	if diffY > 0 {
		p.y++
	} else if diffY < 0 {
		p.y--
	}
}

type Rope struct {
	knots []Pos
}

func (r *Rope) moveHead(diff Pos) {
	r.knots[0].x += diff.x
	r.knots[0].y += diff.y
	for i := 1; i < len(r.knots); i++ {
		r.knots[i].track(r.knots[i-1])
	}
}

func (r *Rope) tail() Pos {
	return r.knots[len(r.knots)-1]
}

func main() {
	input := loadInput("puzzle-input.txt")
	fmt.Printf("unique visited tail positions (rope length 2): %v\n", runRopeSimulation(input, 2))
	fmt.Printf("unique visited tail positions (rope length 10): %v\n", runRopeSimulation(input, 10))
}

var motion = map[string]Pos{
	"L": {-1, 0},
	"R": {1, 0},
	"U": {0, -1},
	"D": {0, 1},
}

func runRopeSimulation(input string, knotCount int) int {
	rope := Rope{knots: make([]Pos, knotCount)}
	visitedPositions := map[Pos]bool{rope.tail(): true}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		repeat, _ := strconv.Atoi(parts[1])
		diff := motion[parts[0]]
		for i := 0; i < repeat; i++ {
			rope.moveHead(diff)
			visitedPositions[rope.tail()] = true
		}
	}
	return len(visitedPositions)
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
