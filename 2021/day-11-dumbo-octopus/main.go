package main

import (
	"fmt"
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
	offset := m.offset(x, y)
	if offset >= 0 && offset < len(m.data) {
		m.data[offset] = value
	}
}

func (m *Map) absorbEnergy(x, y int) {
	if x < 0 || x >= m.width || y < 0 || y >= m.height {
		return
	}
	offset := m.offset(x, y)
	if m.data[offset] > 0 {
		m.data[offset] += 1
	}
}

func (m *Map) get(x, y int) int {
	return m.data[m.offset(x, y)]
}

func (m *Map) flash(x, y int) {
	m.set(x, y, 0)
	m.absorbEnergy(x-1, y-1)
	m.absorbEnergy(x, y-1)
	m.absorbEnergy(x+1, y-1)

	m.absorbEnergy(x-1, y)
	m.absorbEnergy(x+1, y)

	m.absorbEnergy(x-1, y+1)
	m.absorbEnergy(x, y+1)
	m.absorbEnergy(x+1, y+1)
}

func (m *Map) processFlashes() int {
	flashed := 0
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			v := m.get(x, y)
			if v > 9 {
				m.flash(x, y)
				flashed++
			}
		}
	}
	return flashed
}

func (m *Map) step() (int, bool) {
	for i := 0; i < len(m.data); i++ {
		m.data[i]++
	}
	flashes := 0
	for {
		flashed := m.processFlashes()
		if flashed == 0 {
			break
		}
		flashes += flashed
	}
	synchronized := true
	for i := 0; i < len(m.data); i++ {
		if m.data[i] != 0 {
			synchronized = false
			break
		}
	}
	return flashes, synchronized
}

func (m *Map) String() string {
	result := ""
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width-1; x++ {
			result += fmt.Sprintf("%2d ", m.get(x, y))
		}
		result += fmt.Sprintf("%2d\n", m.get(m.width-1, y))
	}
	return result
}

func main() {
	m := parseInput(loadInput("puzzle-input.txt"))
	flashes := 0
	steps := 0
	for {
		flashed, synchronized := m.step()
		steps++
		if steps <= 100 {
			flashes += flashed
		}
		if synchronized {
			break
		}
	}
	fmt.Printf("total flashes after 100 steps: %v, synchronizes after: %v\n", flashes, steps)
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
