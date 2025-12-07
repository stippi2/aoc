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
	return strings.TrimSpace(string(data)), nil
}

// ReadInputLines reads the input and splits it into lines
func ReadInputLines(day int) ([]string, error) {
	content, err := ReadInput(day)
	if err != nil {
		return nil, err
	}
	return strings.Split(content, "\n"), nil
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

// GCD finds the greatest common divisor
func GCD(a, b int) int {
	a = Abs(a)
	b = Abs(b)

	// GCD(0,b) = b, GCD(a,0) = a
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

type Vec2 struct {
	X int
	Y int
}

func (v *Vec2) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

func (v *Vec2) Add(o Vec2) Vec2 {
	return Vec2{X: v.X + o.X, Y: v.Y + o.Y}
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
	for y, line := range lines {
		for x := range width {
			if x < len(line) {
				grid.Set(x, y, line[x])
			} else {
				grid.Set(x, y, ' ')
			}
		}
	}
	return &grid
}

func NewGridFilled(width, height int, fill byte) *Grid {
	grid := Grid{
		width:  width,
		height: height,
		data:   make([]byte, width*height),
	}
	grid.Fill(fill)
	return &grid
}

func (g *Grid) Fill(fill byte) {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			g.Set(x, y, fill)
		}
	}
}

func (g *Grid) ContainsString(s string) bool {
	for y := 0; y < g.height; y++ {
		line := string(g.data[y*g.width : (y+1)*g.width])
		if strings.Contains(line, s) {
			return true
		}
	}
	return false
}

func (g *Grid) String() string {
	var sb strings.Builder
	for i, c := range g.data {
		sb.WriteByte(c)
		if i%g.width == g.width-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

func (g *Grid) Width() int {
	return g.width
}

func (g *Grid) Height() int {
	return g.height
}

func (g *Grid) Get(x, y int) byte {
	if !g.Contains(x, y) {
		return ' '
	}
	return g.data[y*g.width+x]
}

func (g *Grid) Set(x, y int, tile byte) {
	if g.Contains(x, y) {
		offset := y*g.width + x
		g.data[offset] = tile
	}
}

func (g *Grid) Contains(x, y int) bool {
	return x >= 0 && x < g.width && y >= 0 && y < g.height
}

// Item represents any type that has a weight for priority queue ordering
type Item interface {
	Weight() float64
}

// PriorityQueue implements a generic priority queue
type PriorityQueue[T Item] []T

func (pq *PriorityQueue[T]) Len() int { return len(*pq) }

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return (*pq)[i].Weight() < (*pq)[j].Weight()
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(T)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
