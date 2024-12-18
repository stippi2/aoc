package day18

import (
	"aoc/2024/go/lib"
	"container/heap"
	"fmt"
	"math"
	"strings"
)

type Path struct {
	tip    lib.Vec2
	length int
}

func (p *Path) Weight() float64 {
	return float64(p.length)
}

func findShortestPath(input string, gridSize, simulationLength int) int {
	grid := lib.NewGridFilled(gridSize, gridSize, '.')

	currupted := make(map[lib.Vec2]bool)
	for i, line := range strings.Split(input, "\n") {
		if i == simulationLength {
			break
		}
		var pos lib.Vec2
		_, _ = fmt.Sscanf(line, "%d,%d", &pos.X, &pos.Y)
		currupted[pos] = true
		grid.Set(pos.X, pos.Y, '#')
	}

	fmt.Print(grid.String())

	start := &Path{tip: lib.Vec2{X: 0, Y: 0}}
	goal := lib.Vec2{X: gridSize - 1, Y: gridSize - 1}

	pq := &lib.PriorityQueue[*Path]{}
	heap.Init(pq)
	pq.Push(start)

	visited := make(map[lib.Vec2]int)
	visited[start.tip] = 0

	shortestPath := math.MaxInt

	for pq.Len() > 0 {
		path := heap.Pop(pq).(*Path)
		visited[path.tip] = path.length
		if path.tip == goal {
			if path.length < shortestPath {
				fmt.Printf("found goal! length: %d\n", path.length)
				shortestPath = path.length
			}
		}
		if path.length > shortestPath {
			continue
		}

		nextPositions := []lib.Vec2{
			{X: path.tip.X - 1, Y: path.tip.Y},
			{X: path.tip.X + 1, Y: path.tip.Y},
			{X: path.tip.X, Y: path.tip.Y - 1},
			{X: path.tip.X, Y: path.tip.Y + 1},
		}

		for _, next := range nextPositions {
			if next.X >= 0 && next.X < gridSize && next.Y >= 0 && next.Y < gridSize && !currupted[next] {
				previous, found := visited[next]
				if !found || previous > path.length+1 {
					pq.Push(&Path{tip: next, length: path.length + 1})
				}
			}
		}
	}

	return shortestPath
}

func Part1() any {
	input, _ := lib.ReadInput(18)
	return findShortestPath(input, 71, 1000)
}

func Part2() any {
	return "Not implemented"
}
