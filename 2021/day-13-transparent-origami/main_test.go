package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

var exampleInput = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func Test_parseInput(t *testing.T) {
	o := parseInput(exampleInput)
	if assert.Len(t, o.dots, 18) {
		assert.Equal(t, o.dots[0], Point{6, 10})
		assert.Equal(t, o.dots[17], Point{9, 0})
	}
	assert.Len(t, o.folds, 2)
}

func Test_applyFolds(t *testing.T) {
	o := parseInput(exampleInput)
	o.applyFolds()
	assert.Len(t, o.dots, 16)
}

func (o *Origami) sortDots() []Point {
	sort.Slice(o.dots, func(i, j int) bool {
		if o.dots[i].y < o.dots[j].y {
			return true
		}
		if o.dots[i].y == o.dots[j].y {
			return o.dots[i].x < o.dots[j].x
		}
		return false
	})
	return o.dots
}

func Test_applyFirstFold(t *testing.T) {
	o := parseInput(exampleInput)

	o.fold(o.folds[0])
	assert.Equal(t, []Point{
		{0, 0},
		{2, 0},
		{3, 0},
		{6, 0},
		{9, 0},
	}, o.sortDots()[0:5])

	o.fold(o.folds[1])
	assert.Equal(t, []Point{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{0, 1},
		{4, 1},
		{0, 2},
		{4, 2},
		{0, 3},
		{4, 3},
		{0, 4},
		{1, 4},
		{2, 4},
		{3, 4},
		{4, 4},
	}, o.sortDots())
}
