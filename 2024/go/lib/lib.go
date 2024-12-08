package lib

import (
	"fmt"
	"os"
	"strings"
)

// ReadInput reads the input for a specific day
func ReadInput(day int) (string, error) {
	data, err := os.ReadFile(getInputPath(day))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadInputLines reads the input and splits it into lines
func ReadInputLines(day int) ([]string, error) {
	content, err := ReadInput(day)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(content), "\n"), nil
}

// GetInputPath returns the path to the input file
func getInputPath(day int) string {
	return fmt.Sprintf("../input/day%02d.txt", day)
}

// Min returns the minimum of two numbers
func Min[T ~int | ~float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two numbers
func Max[T ~int | ~float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Abs returns the absolute of a value
func Abs[T ~int | ~float64](v T) T {
	if v > 0 {
		return v
	}
	return -v
}

type Vec2 struct {
	X int
	Y int
}

type Grid struct {
	width  int
	height int
	data   []byte
}

func NewGrid(input string) *Grid {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	width := len(lines[0])
	height := len(lines)
	grid := Grid{
		width:  width,
		height: height,
		data:   make([]byte, width*height),
	}
	return &grid
}

func (m *Grid) String() string {
	var sb strings.Builder
	for i, c := range m.data {
		sb.WriteByte(c)
		if i%m.width == m.width-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

func (m *Grid) Width() int {
	return m.width
}

func (m *Grid) Height() int {
	return m.height
}

func (m *Grid) Get(x, y int) byte {
	if x < 0 || x >= m.width || y < 0 || y >= m.height {
		return ' '
	}
	return m.data[y*m.width+x]
}

func (m *Grid) Set(x, y int, tile byte) {
	offset := y*m.width + x
	m.data[offset] = tile
}
