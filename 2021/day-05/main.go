package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	points []Point
}

func main() {
	//lines := parseLines(loadInput("vents-input.txt"))
}

func parseLines(input string) (lines []Line, maxX, maxY int) {
	lineStrings := strings.Split(input, "\n")
	lines = make([]Line, len(lineStrings))
	for i, lineString := range lineStrings {
		points := strings.Split(lineString, " -> ")
		for _, point := range points {
			coords := strings.Split(point, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			lines[i].points = append(lines[i].points, Point{x, y})
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	return
}

func loadInput(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
	exitIfError(err)
	return strings.TrimSpace(string(fileContents))
}

func exitIfError(err error) {
	if err != nil {
		panic(err)
	}
}
