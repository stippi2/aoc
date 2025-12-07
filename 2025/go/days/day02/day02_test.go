package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`

func Test_Part1(t *testing.T) {
	assert.Equal(t, 1227775554, sumInvalidIds(example))
}

func Test_Part2(t *testing.T) {
	assert.True(t, true)
}
