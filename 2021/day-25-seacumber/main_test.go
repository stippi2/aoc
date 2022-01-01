package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>
`

func Test_parseInput(t *testing.T) {
	m := parseInput(example)
	assert.Equal(t, 10, m.width)
	assert.Equal(t, 9, m.height)
	assert.Equal(t, example, fmt.Sprintf("%s", m))
}
