package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"strings"
	"testing"
)

var exampleVectors = `23,-10  25,-9   27,-5   29,-6   22,-6   21,-7   9,0     27,-7   24,-5
25,-7   26,-6   25,-5   6,8     11,-2   20,-5   29,-10  6,3     28,-7
8,0     30,-6   29,-8   20,-10  6,7     6,4     6,1     14,-4   21,-6
26,-10  7,-1    7,7     8,-1    21,-9   6,2     20,-7   30,-10  14,-3
20,-8   13,-2   7,3     28,-8   29,-9   15,-3   22,-5   26,-8   25,-8
25,-6   15,-4   9,-2    15,-2   12,-2   28,-9   12,-3   24,-6   23,-7
25,-10  7,8     11,-3   26,-7   7,1     23,-9   6,0     22,-10  27,-6
8,1     22,-8   13,-4   7,6     28,-6   11,-4   12,-4   26,-9   7,4
24,-10  23,-8   30,-8   7,0     9,-1    10,-1   26,-5   22,-9   6,5
7,5     23,-6   28,-10  10,-2   11,-1   20,-9   14,-2   29,-7   13,-3
23,-5   24,-8   27,-9   30,-7   28,-5   21,-10  7,9     6,6     21,-5
27,-10  7,2     30,-9   21,-8   22,-7   24,-9   20,-6   6,9     29,-5
8,-2    27,-8   30,-5   24,-7`

func Test_aim(t *testing.T) {
	target := Target{
		minX: 20,
		maxX: 30,
		minY: -10,
		maxY: -5,
	}
	aim(target)
}

func Test_possibleVectors(t *testing.T) {
	target := Target{
		minX: 20,
		maxX: 30,
		minY: -10,
		maxY: -5,
	}
	expectedVectors := parseInput(exampleVectors)
	actualVectors := possibleVectors(target)

	assert.Equal(t, expectedVectors, actualVectors)

	fmt.Printf("expected map:\n%v\n", toMap(expectedVectors, target))
	fmt.Printf("calculated map:\n%v\n", toMap(actualVectors, target))
}

func parseInput(input string) map[Vector]bool {
	vectorMap := map[Vector]bool{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		for _, element := range strings.Split(line, " ") {
			if element == "" {
				continue
			}
			coords := strings.Split(element, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			vectorMap[Vector{x, y}] = true
		}
	}
	return vectorMap
}

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

func (m *Map) get(x, y int) string {
	value := m.data[m.offset(x, y)]
	if value == 8 {
		return "#"
	}
	if value == 4 {
		return "O"
	}
	if value == 1 {
		return "_"
	}
	return "."
}

func (m *Map) String() string {
	result := ""
	for y := m.height - 1; y >= 0; y-- {
		for x := 0; x < m.width-1; x++ {
			result += fmt.Sprintf("%s ", m.get(x, y))
		}
		result += fmt.Sprintf("%s\n", m.get(m.width-1, y))
	}
	return result
}

func toMap(vectorMap map[Vector]bool, target Target) *Map {
	if len(vectorMap) == 0 {
		return nil
	}
	minX := 0
	maxX := math.MinInt32
	minY := math.MaxInt32
	maxY := math.MinInt32
	for vector := range vectorMap {
		minX = int(math.Min(float64(vector.x), float64(minX)))
		maxX = int(math.Max(float64(vector.x), float64(maxX)))
		minY = int(math.Min(float64(vector.y), float64(minY)))
		maxY = int(math.Max(float64(vector.y), float64(maxY)))
	}

	m := Map{}
	m.init(maxX - minX + 1, maxY - minY + 1)

	offsetX := -minX
	offsetY := -minY

	for x := minX; x <= maxX; x++ {
		m.set(x + offsetX, offsetY, 1)
	}

	for vector := range vectorMap {
		m.set(vector.x + offsetX, vector.y + offsetY, 8)
	}

	for y := target.minY; y <= target.maxY; y++ {
		for x := target.minX; x <= target.maxX; x++ {
			m.set(x + offsetX, y + offsetY, 4)
		}
	}

	return &m
}
