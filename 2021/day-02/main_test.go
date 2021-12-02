package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_processCommands(t *testing.T) {
	exampleInput := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	var pos Position
	processCommands(&pos, exampleInput)
	assert.Equal(t, 900, pos.calc())
}
