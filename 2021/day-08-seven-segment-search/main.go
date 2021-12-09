package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
	if x < 0 || x >= m.width || y < 0 || y >= m.height {
		return math.MaxInt32
	}
	return m.data[m.offset(x, y)]
}

func (m *Map) isLowPoint(x, y int) bool {
	value := m.get(x, y)
	return value < m.get(x-1, y) &&
		value < m.get(x+1, y) &&
		value < m.get(x, y-1) &&
		value < m.get(x, y+1)
}

func (m *Map) sumRiskLevels() int {
	sum := 0
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if m.isLowPoint(x, y) {
				sum += 1 + m.get(x, y)
			}
		}
	}
	return sum
}

func main() {
	m := parseInput(loadInput("puzzle-input.txt"))
	riskLevelSum := m.sumRiskLevels()
	fmt.Printf("sum of risk levels: %v\n", riskLevelSum)
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
