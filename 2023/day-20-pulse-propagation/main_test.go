package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input1 = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

const input2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

func Test_partOne(t *testing.T) {
	moduleConfiguration := parseInput(input1)
	//moduleConfiguration.pushButton()
	assert.Equal(t, 32000000, partOne(moduleConfiguration))
}

func Test_partOneOutputExample(t *testing.T) {
	moduleConfiguration := parseInput(input2)
	//moduleConfiguration.pushButton()
	//fmt.Println("-----")
	//moduleConfiguration.pushButton()
	//fmt.Println("-----")
	//moduleConfiguration.pushButton()
	//fmt.Println("-----")
	//moduleConfiguration.pushButton()
	assert.Equal(t, 11687500, partOne(moduleConfiguration))
}

func Test_partTwo(t *testing.T) {
	_ = parseInput(input1)
	assert.Equal(t, 0, partTwo())
}
