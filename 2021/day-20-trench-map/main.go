package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Image struct {
	pixel map[Point]bool
	minX, minY, maxX, maxY int
	outside bool
}

func newImage() *Image {
	return &Image{
		pixel: make(map[Point]bool),
		minX:  math.MaxInt32,
		minY:  math.MaxInt32,
		maxX:  math.MinInt32,
		maxY:  math.MinInt32,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (i *Image) set(x, y int) {
	i.pixel[Point{x, y}] = true
	i.maxX = max(i.maxX, x)
	i.maxY = max(i.maxY, y)
	i.minX = min(i.minX, x)
	i.minY = min(i.minY, y)
}

func (i *Image) get(x, y int) bool {
	if x >= i.minX && x <= i.maxX && y >= i.minY && y <= i.maxY {
		return i.pixel[Point{x, y}]
	}
	return i.outside
}


func (i *Image) algorithmIndex(x, y int) int {
	number := ""
	for y1 := y - 1; y1 <= y + 1; y1++ {
		for x1 := x - 1; x1 <= x + 1; x1++ {
			if i.get(x1, y1) {
				number += "1"
			} else {
				number += "0"
			}
		}
	}
	v, err := strconv.ParseInt(number, 2, 16)
	if err != nil {
		panic(fmt.Sprintf("failed to parse binary '%s': %s", number, err))
	}
	return int(v)
}

func (i *Image) enhance(algorithm string) *Image {
	result := newImage()
	for y := i.minY - 1; y <= i.maxY + 1; y++ {
		for x := i.minX - 1; x <= i.maxX + 1; x++ {
			algoIndex := i.algorithmIndex(x, y)
			if algorithm[algoIndex:algoIndex+1] == "#" {
				result.set(x, y)
			}
		}
	}
	algorithmLen := len(algorithm)
	if i.outside {
		result.outside = algorithm[algorithmLen - 1:algorithmLen] == "#"
	} else {
		result.outside = algorithm[0:1] == "#"
	}
	return result
}

func (i *Image) String() string {
	output := ""
	for y := i.minY; y <= i.maxY; y++ {
		for x := i.minX; x <= i.maxX; x++ {
			if i.pixel[Point{x, y}] {
				output += "#"
			} else {
				output += "."
			}
		}
		output += "\n"
	}
	return output
}

func main() {
	// Part 1
	a, image := parseInput(loadInput("puzzle-input.txt"))
	image = image.enhance(a)
	image = image.enhance(a)
	fmt.Printf("number of lit pixels after enhancing twice: %v\n", len(image.pixel))

	// Part 2
	for i := 2; i < 50; i++ {
		image = image.enhance(a)
	}
	fmt.Printf("number of lit pixels after enhancing 50 times: %v\n", len(image.pixel))
}

func parseAlgorithm(input string) string {
	return strings.ReplaceAll(input, "\n", "")
}

func parseImage(input string) *Image {
	image := newImage()
	for y, line := range strings.Split(input, "\n") {
		for x, pixel := range strings.Split(line, "") {
			if pixel == "#" {
				image.set(x, y)
			}
		}
	}
	return image
}

func parseInput(input string) (algorithm string, image *Image) {
	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		panic("unexpected number of input parts")
	}
	algorithm = parseAlgorithm(parts[0])
	image = parseImage(parts[1])
	return
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
