package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Origami struct {
	dots   []Point
	foldsY []int
	foldsX []int
}

func main() {
}

func parseInput(input string) Origami {
	inputSections := strings.Split(input, "\n\n")
	if len(inputSections) != 2 {
		panic("expected sections in the input separated by a double line-break")
	}

	origami := Origami{}

	dots := strings.Split(inputSections[0], "\n")
	for _, dot := range dots {
		coords := strings.Split(dot, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		origami.dots = append(origami.dots, Point{x, y})
	}

	folds := strings.Split(inputSections[1], "\n")
	for _, fold := range folds {
		parts := strings.Split(fold, "=")
		v, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "fold along y":
			origami.foldsY = append(origami.foldsY, v)
		case "fold along x":
			origami.foldsX = append(origami.foldsX, v)
		}
	}
	return origami
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
