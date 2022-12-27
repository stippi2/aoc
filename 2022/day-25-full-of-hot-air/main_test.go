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
	}{
		{1, "1"},
		{2, "2"},
		{3, "1="},
		{4, "1-"},
		{5, "10"},
		{6, "11"},
		{7, "12"},
		{8, "2="},
		{9, "2-"},
		{10, "20"},
		{11, "21"},
		{12, "22"},
		{13, "1=="}, //1x25 -2x5 =15 + -2
		{14, "1=-"},
		{15, "1=0"},
		{16, "1=1"},
		{17, "1=2"},
		{18, "1-="}, // 1x25 -1x5 = 20 -2
		{19, "1--"}, // 1x25 -1x5 = 20 -1
		{20, "1-0"},
	}
	for _, test := range tests {
		assert.Equal(t, test.snafu, fromDecimal(test.decimal))
	}
}
