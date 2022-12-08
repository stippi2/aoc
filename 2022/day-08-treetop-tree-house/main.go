package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	height  int
	visible int
}

type Map struct {
	trees  []Tree
	width  int
	height int
}

func (m *Map) init(width, height int) {
	m.trees = make([]Tree, width*height)
	m.width = width
	m.height = height
}

func (m *Map) get(x, y int) *Tree {
	return &m.trees[y*m.width+x]
}

func checkVisibility(tree *Tree, max int) int {
	if tree.height > max {
		tree.visible++
		return tree.height
	}
	return max
}

func (m *Map) setVisibilityTopBottom() {
	for x := 0; x < m.width; x++ {
		max := -1
		for y := 0; y < m.height; y++ {
			max = checkVisibility(m.get(x, y), max)
		}
		max = -1
		for y := m.height - 1; y >= 0; y-- {
			max = checkVisibility(m.get(x, y), max)
		}
	}
}

func (m *Map) setVisibilityLeftRight() {
	for y := 0; y < m.height; y++ {
		max := -1
		for x := 0; x < m.width; x++ {
			max = checkVisibility(m.get(x, y), max)
		}
		max = -1
		for x := m.width - 1; x >= 0; x-- {
			max = checkVisibility(m.get(x, y), max)
		}
	}
}

func (m *Map) countVisibleTrees() int {
	sum := 0
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if m.get(x, y).visible > 0 {
				sum++
			}
		}
	}
	return sum
}

func increaseScenicScore(tree, treeInSight *Tree, score int) (int, bool) {
	if treeInSight.height < tree.height {
		return score + 1, false
	} else {
		return score + 1, true
	}
}

func (m *Map) computeScenicScore(x, y int) int {
	scoreLeft := 0
	scoreRight := 0
	scoreTop := 0
	scoreBottom := 0
	tree := m.get(x, y)
	stop := false
	for xl := x - 1; xl >= 0 && !stop; xl-- {
		scoreLeft, stop = increaseScenicScore(tree, m.get(xl, y), scoreLeft)
	}
	stop = false
	for xr := x + 1; xr < m.width && !stop; xr++ {
		scoreRight, stop = increaseScenicScore(tree, m.get(xr, y), scoreRight)
	}
	stop = false
	for yt := y - 1; yt >= 0 && !stop; yt-- {
		scoreTop, stop = increaseScenicScore(tree, m.get(x, yt), scoreTop)
	}
	stop = false
	for yb := y + 1; yb < m.height && !stop; yb++ {
		scoreBottom, stop = increaseScenicScore(tree, m.get(x, yb), scoreBottom)
	}
	return scoreLeft * scoreRight * scoreTop * scoreBottom
}

func (m *Map) computeScenicScores() int {
	highScore := 0
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			score := m.computeScenicScore(x, y)
			if score > highScore {
				highScore = score
			}
		}
	}
	return highScore
}

func main() {
	m := parseInput(loadInput("puzzle-input.txt"))
	m.setVisibilityTopBottom()
	m.setVisibilityLeftRight()
	fmt.Printf("visible trees: %v\n", m.countVisibleTrees())
	fmt.Printf("highest scenic score: %v\n", m.computeScenicScores())
}

func parseInput(input string) *Map {
	lines := strings.Split(input, "\n")
	m := &Map{}
	m.init(len(lines[0]), len(lines))
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			height, _ := strconv.Atoi(string(line[x]))
			m.get(x, y).height = height
		}
	}
	return m
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
