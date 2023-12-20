package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Pulse struct {
	isHigh bool
	from   string
	to     string
}

type Module interface {
	InitInput(input string)
	Receive(pulse Pulse) []Pulse
	Outputs() []string
}

type ModuleConfiguration struct {
	modules    map[string]Module
	pulsesHigh int
	pulsesLow  int
}

type Broadcaster struct {
	name    string
	outputs []string
}

type FlipFlop struct {
	Broadcaster
	state bool
}

type Conjunction struct {
	Broadcaster
	inputs map[string]bool
}

func (p Pulse) String() string {
	highOrLow := "low"
	if p.isHigh {
		highOrLow = "high"
	}
	return fmt.Sprintf("%s -%s-> %s", p.from, highOrLow, p.to)
}

func (b *Broadcaster) InitInput(_ string) {
}

func (b *Broadcaster) Outputs() []string {
	return b.outputs
}

func (b *Broadcaster) Receive(pulse Pulse) []Pulse {
	var resultPulses []Pulse
	for _, output := range b.outputs {
		resultPulses = append(resultPulses, Pulse{
			isHigh: pulse.isHigh,
			from:   b.name,
			to:     output,
		})
	}
	return resultPulses
}

func (f *FlipFlop) Receive(pulse Pulse) []Pulse {
	var resultPulses []Pulse
	if !pulse.isHigh {
		f.state = !f.state
		for _, output := range f.outputs {
			resultPulses = append(resultPulses, Pulse{
				isHigh: f.state,
				from:   f.name,
				to:     output,
			})
		}
	}
	return resultPulses
}

func (c *Conjunction) InitInput(input string) {
	c.inputs[input] = false
}

func (c *Conjunction) Receive(pulse Pulse) []Pulse {
	c.inputs[pulse.from] = pulse.isHigh
	sendLow := true
	for _, isHigh := range c.inputs {
		if !isHigh {
			sendLow = false
			break
		}
	}
	var resultPulses []Pulse
	for _, output := range c.outputs {
		resultPulses = append(resultPulses, Pulse{
			isHigh: !sendLow,
			from:   c.name,
			to:     output,
		})
	}
	return resultPulses
}

func (m *ModuleConfiguration) pushButton() {
	m.processPulses([]Pulse{{
		isHigh: false,
		from:   "button",
		to:     "broadcaster",
	}})
}

func (m *ModuleConfiguration) processPulses(pulses []Pulse) {
	var nextPulses []Pulse
	for _, pulse := range pulses {
		if pulse.isHigh {
			m.pulsesHigh++
		} else {
			m.pulsesLow++
		}
		//fmt.Printf("%s\n", pulse)
		module := m.modules[pulse.to]
		if module == nil {
			if !pulse.isHigh {
				fmt.Printf("single low pulse delivered to %s, low: %v, high: %v\n", pulse.to, m.pulsesLow, m.pulsesHigh)
			}
			continue
		}
		nextPulses = append(nextPulses, module.Receive(pulse)...)
	}
	if len(nextPulses) > 0 {
		m.processPulses(nextPulses)
	}
}

func partOne(m *ModuleConfiguration) int {
	for i := 0; i < 1000; i++ {
		m.pushButton()
	}
	return m.pulsesLow * m.pulsesHigh
}

func partTwo(m *ModuleConfiguration) int {
	for i := 0; i < 100000000; i++ {
		m.pushButton()
	}
	return m.pulsesLow * m.pulsesHigh
}

func main() {
	now := time.Now()
	moduleConfiguration := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(moduleConfiguration)
	moduleConfiguration = parseInput(loadInput("puzzle-input.txt"))
	part2 := partTwo(moduleConfiguration)
	duration := time.Since(now)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseInput(input string) *ModuleConfiguration {
	lines := strings.Split(input, "\n")
	modules := make(map[string]Module)
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		name := strings.TrimLeft(parts[0], "%&")
		outputs := strings.Split(parts[1], ", ")
		baseModule := Broadcaster{
			name:    name,
			outputs: outputs,
		}
		if strings.HasPrefix(parts[0], "%") {
			modules[name] = &FlipFlop{
				Broadcaster: baseModule,
			}
		} else if strings.HasPrefix(parts[0], "&") {
			modules[name] = &Conjunction{
				Broadcaster: baseModule,
				inputs:      make(map[string]bool),
			}
		} else {
			modules[name] = &baseModule
		}
	}

	for inputName, module := range modules {
		for _, name := range module.Outputs() {
			output := modules[name]
			if output != nil {
				output.InitInput(inputName)
			}
		}
	}

	return &ModuleConfiguration{
		modules: modules,
	}
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
