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

func Test_totalSyntaxErrorScore(t *testing.T) {
	totalScore := totalSyntaxErrorScore(parseInput(exampleInput))
	assert.Equal(t, 26397, totalScore)
}

func Test_totalAutocompleteScore(t *testing.T) {
	totalScore := totalAutocompleteScore(parseInput(exampleInput))
	assert.Equal(t, 288957, totalScore)
}
