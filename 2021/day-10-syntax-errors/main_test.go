package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func Test_parseLines(t *testing.T) {
	totalScore := parseLines(parseInput(exampleInput))
	assert.Equal(t, 26397, totalScore)
}
