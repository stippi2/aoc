package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = []string{
	`v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`,
	`....>.>v.>
v.v>.>v.v.
>v>>..>v..
>>v>v>.>.v
.>v.v...v.
v>>.>vvv..
..v...>>..
vv...>>vv.
>.v.v..v.v`,
	`>.v.v>>..v
v.v.>>vv..
>v>.>.>.v.
>>v>v.>v>.
.>..v....v
.>v>>.v.v.
v....v>v>.
.vv..>>v..
v>.....vv.`,
	`v>v.v>.>v.
v...>>.v.v
>vv>.>v>..
>>v>v.>.v>
..>....v..
.>.>v>v..v
..v..v>vv>
v.v..>>v..
.v>....v..`,
	`v>..v.>>..
v.v.>.>.v.
>vv.>>.v>v
>>.>..v>.>
..v>v...v.
..>>.>vv..
>.v.vv>v.v
.....>>vv.
vvv>...v..`,
	`vv>...>v>.
v.v.v>.>v.
>.v.>.>.>v
>v>.>..v>>
..v>v.v...
..>.>>vvv.
.>...v>v..
..v.v>>v.v
v.v.>...v.`,
}

func Test_parseInput(t *testing.T) {
	m := parseInput(example[0])
	assert.Equal(t, 10, m.width)
	assert.Equal(t, 9, m.height)
	assert.Equal(t, example[0], fmt.Sprintf("%s", m))
}

func Test_step(t *testing.T) {
	m := parseInput(example[0])
	for step := 1; step < len(example); step++ {
		m.step()
		assert.Equal(t, example[step], fmt.Sprintf("%s", m))
	}
}

func Test_stepSimple(t *testing.T) {
	tests := []struct {
		input string
		expected string
	}{
		{
			input:    "...>>>>>...",
			expected: "...>>>>.>..",
		},
		{
			input:    "...>>>>.>..",
			expected: "...>>>.>.>.",
		},
		{
			input:    `..........
.>v....v..
.......>..
..........`,
			expected: `..........
.>........
..v....v>.
..........`,
		},
	}
	for _, test := range tests {
		m := parseInput(test.input)
		m.step()
		assert.Equal(t, test.expected, m.String())
	}
}

func Test_noneMoved(t *testing.T) {
	m := parseInput(example[0])
	step := 0
	for {
		step++
		if !m.step() {
			break
		}
	}
	assert.Equal(t, 58, step)
}
