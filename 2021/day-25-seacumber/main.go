package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
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
	x = x % m.width
	for y < 0 {
		y = m.height + y
	}
	y = y % m.height
	return x, y
}

func (m *Map) set(x, y int, value uint8) {
	x, y = m.wrap(x, y)
	m.data[m.offset(x, y)] = value
}

func (m *Map) get(x, y int) uint8 {
	x, y = m.wrap(x, y)
	return m.data[m.offset(x, y)]
}

func (m *Map) step() bool {
	anyMoved := false
	// East
	for y := 0; y < m.height; y++ {
		previous := m.get(-1, y)
		for x := 0; x < m.width; x++ {
			v := m.get(x, y)
			if v == '.' && previous == '>' {
				m.set(x - 1, y, 'x')
				m.set(x, y, ')')
				anyMoved = true
			}
			previous = v
		}
	}
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			switch m.get(x, y) {
			case 'x':
				m.set(x, y, '.')
			case ')':
				m.set(x, y, '>')
			}
		}
	}
	// South
	for x := 0; x < m.width; x++ {
		previous := m.get(x, -1)
		for y := 0; y < m.height; y++ {
			v := m.get(x, y)
			if v == '.' && previous == 'v' {
				m.set(x, y - 1, 'x')
				m.set(x, y, 'V')
				anyMoved = true
			}
			previous = v
		}
	}
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			switch m.get(x, y) {
			case 'x':
				m.set(x, y, '.')
			case 'V':
				m.set(x, y, 'v')
			}
		}
	}
	return anyMoved
}

func (m *Map) String() string {
	result := ""
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width-1; x++ {
			result += fmt.Sprintf("%s", string(m.get(x, y)))
		}
		result += fmt.Sprintf("%s\n", string(m.get(m.width-1, y)))
	}
	return result[0:len(result)-1]
}

func main() {
	m := parseInput(loadInput("puzzle-input.txt"))
	start := time.Now()
	step := 0
	for {
		step++
		if !m.step() {
			break
		}
	}
	fmt.Printf("movement stopped after step: %v (took %s)\n", step, time.Since(start))
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
