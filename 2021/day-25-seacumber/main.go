package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Map struct {
	width  int
	height int
	data   []uint8
}

func (m *Map) init(width, height int) {
	m.width = width
	m.height = height
	m.data = make([]uint8, width*height)
}

func (m *Map) offset(x, y int) int {
	return m.width*y + x
}

func (m *Map) wrap(x, y int) (int, int) {
	for x < 0 {
		x = m.width + x
	}
	for y < 0 {
		y = m.height + y
	}
	return x, y
}

func (m *Map) set(x, y int, value uint8) {
	x, y = m.wrap(x, y)
	offset := m.offset(x, y)
	m.data[offset] = value
}

func (m *Map) get(x, y int) uint8 {
	x, y = m.wrap(x, y)
	return m.data[m.offset(x, y)]
}

func (m *Map) step() {
}

func (m *Map) String() string {
	result := ""
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width-1; x++ {
			result += fmt.Sprintf("%s", string(m.get(x, y)))
		}
		result += fmt.Sprintf("%s\n", string(m.get(m.width-1, y)))
	}
	return result
}

func main() {
}

func parseInput(input string) *Map {
	input = strings.TrimSpace(input)
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
		for x := 0; x < len(lines[y]); x++ {
			if x >= width {
				break
			}
			m.set(x, y, lines[y][x])
		}
	}
	return &m
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
