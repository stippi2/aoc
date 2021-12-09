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
	return m.height*y + x
}

func (m *Map) set(x, y, value int) {
	m.data[m.offset(x, y)] = value
}

func (m *Map) get(x, y int) int {
	return m.data[m.offset(x, y)]
}

func main() {
}

func parseInput(input string) Map {
	lines := strings.Split(input, "\n")
	var m Map
	m.init(len(lines[0]), len(lines))
	for y, line := range lines {
		for x, c := range strings.Split(line, "") {
			value, err := strconv.Atoi(c)
			if err == nil {
				m.set(x, y, value)
			}
		}
	}
	return m
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
