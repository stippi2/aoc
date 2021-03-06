package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Origami struct {
	dots   []Point
	folds  []Folder
}

type Folder interface {
	Fold(p Point) Point
}

type FolderY struct {
	line int
}

type FolderX struct {
	line int
}

func (f *FolderX) Fold(p Point) Point {
	if p.x > f.line {
		diff := p.x - f.line
		return Point{f.line - diff, p.y}
	} else {
		return p
	}
}

//func (f *FolderX) String() string {
//	return fmt.Sprintf("fold along x at %v", f.line)
//}

func (f *FolderY) Fold(p Point) Point {
	if p.y > f.line {
		diff := p.y - f.line
		return Point{p.x, f.line - diff}
	} else {
		return p
	}
}

//func (f *FolderY) String() string {
//	return fmt.Sprintf("fold along y at %v", f.line)
//}

func contains(points []Point, point Point) bool {
	for _, p := range points {
		if p == point {
			return true
		}
	}
	return false
}

func (o *Origami) fold(f Folder) {
	var dots []Point
	for _, dot := range o.dots {
		dot = f.Fold(dot)
		if !contains(dots, dot) {
			dots = append(dots, dot)
		}
	}
	o.dots = dots
}

func (o *Origami) applyFolds() {
	for _, folder := range o.folds {
//		fmt.Printf("%v\n", folder)
		o.fold(folder)
	}
}

func (o *Origami) sortDots() []Point {
	sort.Slice(o.dots, func(i, j int) bool {
		if o.dots[i].y < o.dots[j].y {
			return true
		}
		if o.dots[i].y == o.dots[j].y {
			return o.dots[i].x < o.dots[j].x
		}
		return false
	})
	return o.dots
}

func (o *Origami) String() string {
	o.sortDots()
	maxX := math.MinInt32
	maxY := math.MinInt32
	for _, dot := range o.dots {
		if dot.x > maxX {
			maxX = dot.x
		}
		if dot.y > maxY {
			maxY = dot.y
		}
	}
	result := ""
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if contains(o.dots, Point{x, y}) {
				result += "#"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func main() {
	o := parseInput(loadInput("puzzle-input.txt"))
	// Part 1
	o.fold(o.folds[0])
	fmt.Printf("dots visible after first fold: %v\n", len(o.dots))
	// Part 2
	o.folds = o.folds[1:]
	o.applyFolds()
	fmt.Printf("transparent origami:\n%s\n", &o)
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
			origami.folds = append(origami.folds, &FolderY{v})
		case "fold along x":
			origami.folds = append(origami.folds, &FolderX{v})
		}
	}
	return origami
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
