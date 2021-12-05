package main

import (
	"fmt"
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

type DangerMap struct {
	danger []int
	width  int
	height int
}

func (d *DangerMap) init(maxX, maxY int) {
	d.width = maxX + 1
	d.height = maxY + 1
	d.danger = make([]int, d.height*d.width)
}

func (d *DangerMap) offset(p Point) int {
	return d.width*p.y + p.x
}

func (d *DangerMap) increaseDanger(p Point) {
	d.danger[d.offset(p)]++
}

func (d *DangerMap) markVentLine(l Line) {
	a := l.points[0]
	b := l.points[1]
	// assumes abs(diffX) == abs(diffY) or one diff == 0
	diffX := b.x - a.x
	diffY := b.y - a.y
	d.increaseDanger(a)
	for {
		if diffX > 0 {
			a.x++
		} else if diffX < 0 {
			a.x--
		}
		if diffY > 0 {
			a.y++
		} else if diffY < 0 {
			a.y--
		}
		d.increaseDanger(a)
		if a == b {
			break
		}
	}
}

func buildDangerMap(lines []Line, maxX, maxY int) *DangerMap {
	dangerMap := DangerMap{}
	dangerMap.init(maxX, maxY)
	for _, line := range lines {
		dangerMap.markVentLine(line)
	}
	return &dangerMap
}

func (d *DangerMap) countPoints(minDanger int) (count int) {
	for i := 0; i < len(d.danger); i++ {
		if d.danger[i] >= minDanger {
			count++
		}
	}
	return
}

func main() {
	dangerMap := buildDangerMap(parseLines(loadInput("vents-input.txt")))
	dangerPoints := dangerMap.countPoints(2)
	fmt.Printf("points with danger greater 2: %v\n", dangerPoints)
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
