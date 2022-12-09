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

type TrackingPos struct {
	current          Pos
	visitedPositions map[Pos]int
}

func (t *TrackingPos) setPos(p Pos) {
	visited := t.visitedPositions[p]
	t.visitedPositions[p] = visited + 1
	t.current = p
}

func (t *TrackingPos) track(p Pos) {
	newPos := t.current
	diffX := p.x - newPos.x
	diffY := p.y - newPos.y

	if math.Abs(float64(diffX)) <= 1 && math.Abs(float64(diffY)) <= 1 {
		return
	}

	if diffX > 0 {
		newPos.x++
	} else if diffX < 0 {
		newPos.x--
	}
	if diffY > 0 {
		newPos.y++
	} else if diffY < 0 {
		newPos.y--
	}

	t.setPos(newPos)
}

func newTrackablePos() *TrackingPos {
	pos := &TrackingPos{
		visitedPositions: map[Pos]int{},
	}
	pos.visitedPositions[pos.current] = 1
	return pos
}

type Rope struct {
	knots []TrackingPos
}

func (r *Rope) moveHead(diff Pos) {
	head := &r.knots[0]
	head.setPos(Pos{head.current.x + diff.x, head.current.y + diff.y})
	for i := 1; i < len(r.knots); i++ {
		r.knots[i].track(r.knots[i-1].current)
	}
}

func (r *Rope) appendKnots(count int) {
	for i := 0; i < count; i++ {
		r.knots = append(r.knots, *newTrackablePos())
	}
}

func (r *Rope) tail() *TrackingPos {
	return &r.knots[len(r.knots)-1]
}

func partOne() {
	rope := &Rope{}
	rope.appendKnots(2)
	runPositions(loadInput("puzzle-input.txt"), rope)
	fmt.Printf("unique visited tail positions: %v\n", len(rope.tail().visitedPositions))
}

func partTwo() {
	rope := &Rope{}
	rope.appendKnots(10)
	runPositions(loadInput("puzzle-input.txt"), rope)
	fmt.Printf("unique visited tail positions: %v\n", len(rope.tail().visitedPositions))
}

func main() {
	partOne()
	partTwo()
}

var motion = map[string]Pos{
	"L": {-1, 0},
	"R": {1, 0},
	"U": {0, -1},
	"D": {0, 1},
}

func runPositions(input string, rope *Rope) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		repeat, _ := strconv.Atoi(parts[1])
		diff := motion[parts[0]]
		for i := 0; i < repeat; i++ {
			rope.moveHead(diff)
		}
	}
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
