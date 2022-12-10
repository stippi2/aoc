package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CycleRunner interface {
	AtTick(tick, x int)
}

type SignalStrengthCollector struct {
	signalStrength int
}

func (c *SignalStrengthCollector) AtTick(tick, x int) {
	if tick <= 220 && (tick+20)%40 == 0 {
		c.signalStrength += x * tick
	}
}

type Display struct {
	frameBuffer string
}

func (d *Display) AtTick(tick, x int) {
	col := (tick - 1) % 40
	if x-1 <= col && x+1 >= col {
		d.frameBuffer += "#"
	} else {
		d.frameBuffer += "."
	}
	if col == 39 {
		d.frameBuffer += "\n"
	}
}

func main() {
	input := loadInput("puzzle-input.txt")
	// part 1
	c := SignalStrengthCollector{}
	parseAndRunInstructions(input, &c)
	fmt.Printf("signal strength after 220 ticks: %v\n", c.signalStrength)
	// part 2
	d := Display{}
	parseAndRunInstructions(input, &d)
	fmt.Printf("display contents:\n%v", d.frameBuffer)
}

func parseAndRunInstructions(input string, runner CycleRunner) {
	x := 1
	tick := 1
	for _, line := range strings.Split(input, "\n") {
		runner.AtTick(tick, x)
		tick++
		if line != "noop" {
			runner.AtTick(tick, x)
			tick++
			valueString := strings.TrimPrefix(line, "addx ")
			value, _ := strconv.Atoi(valueString)
			x += value
		}
	}
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
