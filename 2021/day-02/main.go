package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Position struct {
	aim        int
	horizontal int
	depth      int
}

// forward increases your horizontal position by X units.
// And it increases your depth by your aim multiplied by X.
func (p *Position) forward(howMuch int) {
	p.horizontal += howMuch
	p.depth += p.aim * howMuch
}

func (p *Position) up(howMuch int) {
	p.aim -= howMuch
}

func (p *Position) down(howMuch int) {
	p.aim += howMuch
}

func (p *Position) calc() int {
	return p.horizontal * p.depth
}

func (p *Position) processCommand(cmd string) error {
	cmdAndValue := strings.Split(cmd, " ")
	if len(cmdAndValue) != 2 {
		return fmt.Errorf("invalid command: %s", cmd)
	}
	value, err := strconv.Atoi(cmdAndValue[1])
	if err != nil {
		return fmt.Errorf("invalid command value: %s: %w", cmdAndValue[1], err)
	}
	switch cmdAndValue[0] {
	case "forward":
		p.forward(value)
	case "up":
		p.up(value)
	case "down":
		p.down(value)
	default:
		return fmt.Errorf("unexpected command: %s", cmdAndValue[0])
	}
	return nil
}

func processCommands(p *Position, commands []string) {
	for _, cmd := range commands {
		if cmd != "" {
			err := p.processCommand(cmd)
			exitIfError(err)
		}
	}
}

func main() {
	commands := loadCommands("pos-commands.txt")
	var pos Position
	processCommands(&pos, commands)
	fmt.Printf("position value: %v\n", pos.calc())
}

func loadCommands(filename string) []string {
	fileContents, err := ioutil.ReadFile(filename)
	exitIfError(err)

	return strings.Split(string(fileContents), "\n")
}

func exitIfError(err error) {
	if err != nil {
		panic(err)
	}
}
