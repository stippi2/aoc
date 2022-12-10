package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := loadInput("puzzle-input.txt")
	fmt.Printf("signal strength after 220 ticks: %v\n", getSignalStrength(input))
}

func increaseSignalStrength(x, tick, signal int) int {
	if (tick+20)%40 == 0 {
		signal += x * tick
	}
	return signal
}

func getSignalStrength(input string) int {
	x := 1
	tick := 1
	signalStrength := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		signalStrength = increaseSignalStrength(x, tick, signalStrength)
		tick++
		if line != "noop" {
			signalStrength = increaseSignalStrength(x, tick, signalStrength)
			tick++
			line = strings.TrimPrefix(line, "addx ")
			value, _ := strconv.Atoi(line)
			x += value
		}
		if tick > 220 {
			break
		}
	}
	return signalStrength
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
