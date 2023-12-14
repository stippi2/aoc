package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Map struct {
	width   int
	height  int
	rows    []string
	columns []string
}

func (m *Map) initColumns() {
	m.columns = make([]string, m.width)
	for x := 0; x < m.width; x++ {
		var sb strings.Builder
		for y := 0; y < m.height; y++ {
			sb.WriteByte(m.rows[y][x])
		}
		m.columns[x] = sb.String()
	}
}

func (m *Map) getRow(y int) string {
	return m.rows[y]
}

func (m *Map) getColumn(x int) string {
	return m.columns[x]
}

func getItem(items *[]string, pos, flipPos, flipItem int) string {
	item := (*items)[pos]
	if pos == flipPos {
		c := item[flipItem]
		if c == '#' {
			c = '.'
		} else {
			c = '#'
		}
		item = item[:flipItem] + string(c) + item[flipItem+1:]
	}
	return item
}

func (m *Map) isMirrorAxis(axis int, items *[]string, flipPos, flipItem int) bool {
	//direction := "columns"
	//if items == &m.rows {
	//	direction = "rows"
	//}
	//fmt.Printf("  direction: %s, axis: %d, fliPos: %d, flipItem: %d\n", direction, axis, flipPos, flipItem)
	for i := 1; i <= axis; i++ {
		i1 := axis - i
		i2 := axis + i - 1
		if i1 < 0 || i2 >= len(*items) {
			break
		}
		//	fmt.Printf("axis: %d, comparing %d <-> %d\n", axis, i1, i2)
		items1 := getItem(items, i1, flipPos, flipItem)
		items2 := getItem(items, i2, flipPos, flipItem)
		if items1 != items2 {
			return false
		}
	}
	return true
}

func (m *Map) findMirrorAxisX(flipX, flipY int) []int {
	var mirrorAxes []int
	//fmt.Printf("findMirrorAxisX() flipX: %d, flipY: %d\n", flipX, flipY)
	for axis := 1; axis < m.width; axis++ {
		if m.isMirrorAxis(axis, &m.columns, flipX, flipY) {
			mirrorAxes = append(mirrorAxes, axis)
		}
	}
	return mirrorAxes
}

func (m *Map) findMirrorAxisY(flipX, flipY int) []int {
	var mirrorAxes []int
	//fmt.Printf("findMirrorAxisY() flipX: %d, flipY: %d\n", flipX, flipY)
	for axis := 1; axis < m.height; axis++ {
		if m.isMirrorAxis(axis, &m.rows, flipY, flipX) {
			mirrorAxes = append(mirrorAxes, 100*axis)
		}
	}
	return mirrorAxes
}

func (m *Map) findMirrorAxisValue(flipX, flipY int) ([]int, []int) {
	valuesX := m.findMirrorAxisX(flipX, flipY)
	valuesY := m.findMirrorAxisY(flipX, flipY)
	return valuesX, valuesY
}

func partOne(maps []*Map) int {
	sum := 0
	for _, m := range maps {
		valuesX, valuesY := m.findMirrorAxisValue(-1, -1)
		// Assume one of them will be empty
		value := 0
		if len(valuesX) > 0 {
			value = valuesX[0]
		} else if len(valuesY) > 0 {
			value = valuesY[0]
		}
		//fmt.Printf("map %d, found axis: %d\n", i+1, value)
		sum += value
	}
	return sum
}

func partTwo(maps []*Map) int {
	sum := 0
	for i, m := range maps {
		part1X, part1Y := m.findMirrorAxisValue(-1, -1)
		if part1X != nil && part1Y != nil {
			panic("Found at least two mirror axis for part 1!")
		}
		// Assume one of them will be empty
		part1 := 0
		if len(part1X) > 0 {
			part1 = part1X[0]
		} else if len(part1Y) > 0 {
			part1 = part1Y[0]
		}

		values := map[int]int{}
		for y := 0; y < m.height; y++ {
			for x := 0; x < m.width; x++ {
				valuesX, valuesY := m.findMirrorAxisValue(x, y)
				for _, valueX := range valuesX {
					values[valueX]++
				}
				for _, valueY := range valuesY {
					values[valueY]++
				}
			}
		}
		valueAdded := false
		for k := range values {
			if k == 0 {
				continue
			}
			//fmt.Printf("map %d, found value: %d %d times\n", i+1, k, v)
			if k != part1 {
				valueAdded = true
				sum += k
				break
			}
		}
		if !valueAdded {
			panic(fmt.Sprintf("No value added for map %d!", i+1))
		}
	}
	return sum
}

func main() {
	now := time.Now()
	maps := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(maps)
	part2 := partTwo(maps)
	duration := time.Since(now)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseInput(input string) []*Map {
	var maps []*Map
	sections := strings.Split(input, "\n\n")
	for _, section := range sections {
		lines := strings.Split(section, "\n")
		m := &Map{
			width:  len(lines[0]),
			height: len(lines),
			rows:   lines,
		}
		m.initColumns()
		maps = append(maps, m)
	}
	return maps
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
