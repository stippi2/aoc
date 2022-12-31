package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`

func Test_part1(t *testing.T) {
	numbers := parseInput(exampleInput)
	decimal := sum(numbers)
	assert.Equal(t, 4890, decimal)
	//	assert.Equal(t, "2=-1=0", fromDecimal(decimal))
}

func Test_fromDecimal(t *testing.T) {
	tests := []struct {
		decimal int
		snafu   string
		quintal string
	}{
		{1, "1", "1"},
		{2, "2", "2"},
		{3, "1=", "3"},
		{4, "1-", "4"},
		{5, "10", "10"},
		{6, "11", "11"},
		{7, "12", "12"},
		{8, "2=", "13"},
		{9, "2-", "14"},
		{10, "20", "20"},
		{11, "21", "21"},
		{12, "22", "22"},
		{13, "1==", "23"}, //1x25 -2x5 =15 + -2
		{14, "1=-", "24"},
		{15, "1=0", "30"},
		{16, "1=1", "31"},
		{17, "1=2", "32"},
		{18, "1-=", "33"}, // 1x25 -1x5 = 20 -2
		{19, "1--", "34"}, // 1x25 -1x5 = 20 -1
		{20, "1-0", "40"},
		{50 + 15 + 4, "1=--", "234"},
	}
	for _, test := range tests {
		assert.Equal(t, test.snafu, fromDecimal(test.decimal))
	}
}
