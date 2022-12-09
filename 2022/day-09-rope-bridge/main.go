package main

import (
	"fmt"
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (t *TrackingPos) track(p Pos) {
	newPos := t.current
	diffX := p.x - newPos.x
	diffY := p.y - newPos.y

	if abs(diffX) <= 1 && abs(diffY) <= 1 {
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

func runPositions(input string, rope *Rope) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		repeat, _ := strconv.Atoi(parts[1])
		diff := Pos{}
		switch parts[0] {
		case "R":
			diff.x = 1
		case "L":
			diff.x = -1
		case "U":
			diff.y = -1
		case "D":
			diff.y = 1
		}
		for i := 0; i < repeat; i++ {
			rope.moveHead(diff)
		}
	}
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
