package day07

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	calibrations := parseInput([]string{"42: 40 2", "3: 1 2"})
	expected := []Calibration{
		{result: 42, sequence: []int64{40, 2}},
		{result: 3, sequence: []int64{1, 2}},
	}
	assert.Equal(t, expected, calibrations)
}

func Test_sumValidCalibrations(t *testing.T) {
	inputLines := strings.Split(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`, "\n")
	assert.Equal(t, 3749, sumValidCalibrations(inputLines))
}
