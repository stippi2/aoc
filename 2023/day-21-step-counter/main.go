package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Map struct {
	width    int
	height   int
	active   int
	tiles    [2][]int
	infinite bool
	leftUp   bool
}

func (m *Map) String() string {
	s := ""
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			tile := m.tiles[m.active][y*m.width+x]
			switch tile {
			case -1:
				s += "#"
			case 0:
				s += "."
			case 1, 2, 3, 4, 5, 6, 7, 8, 9:
				s += string('0' + byte(tile))
			default:
				s += "?"
			}
		}
		s += "\n"
	}
	return s
}

func (m *Map) clone() *Map {
	c := &Map{
		width:  m.width,
		height: m.height,
		tiles: [2][]int{
			make([]int, m.width*m.height),
			make([]int, m.width*m.height),
		},
		active:   m.active,
		infinite: m.infinite,
		leftUp:   m.leftUp,
	}
	copy(c.tiles[0], m.tiles[0])
	copy(c.tiles[1], m.tiles[1])
	return c
}

func (m *Map) getTile(x, y int) int {
	return m.tiles[m.active][y*m.width+x]
}

func (m *Map) setTile(x, y int, tile int) {
	if m.infinite {
		onNextMap := false
		if m.leftUp {
			if x < 0 {
				x = m.width - 1
				onNextMap = true
			} else if y < 0 {
				y = m.height - 1
				onNextMap = true
			} else if x >= m.width || y >= m.height {
				return
			}
		} else {
			if x >= m.width {
				x = 0
				onNextMap = true
			} else if y >= m.height {
				y = 0
				onNextMap = true
			} else if x < 0 || y < 0 {
				return
			}
		}
		if m.tiles[m.active][y*m.width+x] != -1 {
			if onNextMap {
				m.tiles[m.active][y*m.width+x] += tile
			} else if m.tiles[m.active][y*m.width+x] < tile || tile == -1 {
				m.tiles[m.active][y*m.width+x] = tile
			}
		}
	} else {
		if x < 0 || x >= m.width || y < 0 || y >= m.height {
			return
		}
		if m.tiles[m.active][y*m.width+x] != -1 {
			m.tiles[m.active][y*m.width+x] = tile
		}
	}
}

func (m *Map) step() {
	nextActive := (m.active + 1) % 2
	for i, tile := range m.tiles[m.active] {
		if tile > 0 {
			tile = 0
		}
		m.tiles[nextActive][i] = tile
	}
	oldActive := m.active
	m.active = nextActive
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			tile := m.tiles[oldActive][y*m.width+x]
			if tile > 0 {
				m.setTile(x-1, y, tile)
				m.setTile(x, y-1, tile)
				m.setTile(x+1, y, tile)
				m.setTile(x, y+1, tile)
			}
		}
	}
}

func (m *Map) countTilesReachable(steps int) int {
	for i := 0; i < steps; i++ {
		m.step()
	}
	count := 0
	for _, tile := range m.tiles[m.active] {
		if tile > 0 {
			count += tile
		}
	}
	return count
}

func partOne(m *Map, steps int) int {
	return m.countTilesReachable(steps)
}

func partTwo(m *Map, steps int) int {
	m.leftUp = true
	m2 := m.clone()
	m2.leftUp = false
	return m.countTilesReachable(steps) + m2.countTilesReachable(steps)
}

func main() {
	const input = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

	m := parseInput(input, true)

	for {
		m.step()
		fmt.Printf("%s", m)

		var in string
		_, _ = fmt.Scanln(&in)

		fmt.Printf("\033[%vA", m.height+1)
	}
}

func main__() {
	now := time.Now()
	m := parseInput(loadInput("puzzle-input.txt"), false)
	part1 := partOne(m, 64)
	m = parseInput(loadInput("puzzle-input.txt"), true)
	part2 := partTwo(m, 26501365)
	duration := time.Since(now)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseInput(input string, infinite bool) *Map {
	lines := strings.Split(input, "\n")
	width := len(lines[0])
	height := len(lines)
	if infinite && width%2 == 1 {
		width *= 2
	}
	if infinite && height%2 == 1 {
		height *= 2
	}
	m := &Map{
		width:  width,
		height: height,
		tiles: [2][]int{
			make([]int, width*height),
			make([]int, width*height),
		},
		infinite: infinite,
	}
	for y, line := range lines {
		for x, tile := range line {
			if tile == 'S' {
				m.setTile(x, y, 1)
			} else if tile == '#' {
				m.setTile(x, y, -1)
			}
		}
	}
	if width == 2*len(lines[0]) {
		// copy tiles to the right
		for y := 0; y < len(lines); y++ {
			for x := m.width / 2; x < m.width; x++ {
				if m.getTile(x-m.width/2, y) == -1 {
					m.setTile(x, y, m.getTile(x-m.width/2, y))
				}
			}
		}
	}
	if height == 2*len(lines) {
		// copy tiles to the bottom
		for y := m.height / 2; y < m.height; y++ {
			for x := 0; x < m.width; x++ {
				if m.getTile(x, y-m.height/2) == -1 {
					m.setTile(x, y, m.getTile(x, y-m.height/2))
				}
			}
		}
	}
	return m
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
