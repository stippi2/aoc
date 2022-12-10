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

func drawAndTrackSignal(x, tick, signal int) int {
	col := tick % 40
	if x <= col && x+2 >= col {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if col == 0 {
		fmt.Print("\n")
	}
	if tick <= 220 && (tick+20)%40 == 0 {
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
		signalStrength = drawAndTrackSignal(x, tick, signalStrength)
		tick++
		if line != "noop" {
			signalStrength = drawAndTrackSignal(x, tick, signalStrength)
			tick++
			line = strings.TrimPrefix(line, "addx ")
			value, _ := strconv.Atoi(line)
			x += value
		}
	}
	return signalStrength
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
