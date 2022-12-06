package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Part1(t *testing.T) {
	assert.Equal(t, 7, findWindowOfDifferentChars("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4))
	assert.Equal(t, 5, findWindowOfDifferentChars("bvwbjplbgvbhsrlpgdmjqwftvncz", 4))
	assert.Equal(t, 11, findWindowOfDifferentChars("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 19, findWindowOfDifferentChars("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))
	assert.Equal(t, 23, findWindowOfDifferentChars("bvwbjplbgvbhsrlpgdmjqwftvncz", 14))
	assert.Equal(t, 29, findWindowOfDifferentChars("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14))
}

func Test_Part1Quick(t *testing.T) {
	assert.Equal(t, 7, findDifferentCharsQuick("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4))
	assert.Equal(t, 5, findDifferentCharsQuick("bvwbjplbgvbhsrlpgdmjqwftvncz", 4))
	assert.Equal(t, 11, findDifferentCharsQuick("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4))
}

func Test_Part2Quick(t *testing.T) {
	assert.Equal(t, 19, findDifferentCharsQuick("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))
	assert.Equal(t, 23, findDifferentCharsQuick("bvwbjplbgvbhsrlpgdmjqwftvncz", 14))
	assert.Equal(t, 29, findDifferentCharsQuick("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14))
}
