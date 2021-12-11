package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Map struct {
	width  int
	height int
	data   []int
}

func (m *Map) init(width, height int) {
	m.width = width
	m.height = height
	m.data = make([]int, width*height)
}

func (m *Map) offset(x, y int) int {
	return m.width*y + x
}

func (m *Map) set(x, y, value int) {
	m.data[m.offset(x, y)] = value
}

func (m *Map) get(x, y int) int {
	return m.data[m.offset(x, y)]
}

func main() {
}

func parseInput(input string) *Map {
	lines := strings.Split(input, "\n")
	height := len(lines)
	if height == 0 {
		return nil
	}
	width := len(lines[0])
	if width == 0 {
		return nil
	}
	m := Map{}
	m.init(width, height)
	for y := 0; y < height; y++ {
		for x, char := range strings.Split(lines[y], "") {
			if x >= width {
				break
			}
			v, _ := strconv.Atoi(char)
			m.set(x, y, v)
		}
	}
	return &m
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
