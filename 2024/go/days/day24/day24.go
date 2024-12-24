package day24

import (
	"aoc/2024/go/lib"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Wire struct {
	name  string
	value string
}

type Gate struct {
	operation string
	inputA    *Wire
	inputB    *Wire
	output    *Wire
}

func (g *Gate) getOutput() string {
	if g.inputA.value != "" && g.inputB.value != "" {
		switch g.operation {
		case "AND":
			if g.inputA.value == "1" && g.inputB.value == "1" {
				return "1"
			} else {
				return "0"
			}
		case "OR":
			if g.inputA.value == "1" || g.inputB.value == "1" {
				return "1"
			} else {
				return "0"
			}
		case "XOR":
			if g.inputA.value != g.inputB.value {
				return "1"
			} else {
				return "0"
			}
		}
	}
	return ""
}

func parseInput(input string) (map[string]*Wire, []*Gate, map[string]string) {
	wires := make(map[string]*Wire)
	var gates []*Gate

	getOrCreateWire := func(name string) *Wire {
		wire := wires[name]
		if wire == nil {
			wire = &Wire{name: name}
			wires[name] = wire
		}
		return wire
	}

	parts := strings.Split(input, "\n\n")

	// parse gates
	for _, line := range strings.Split(parts[1], "\n") {
		inputsAndOutput := strings.Split(line, " -> ")
		output := inputsAndOutput[1]
		inputAndOperations := strings.Split(inputsAndOutput[0], " ")
		gates = append(gates, &Gate{
			inputA:    getOrCreateWire(inputAndOperations[0]),
			inputB:    getOrCreateWire(inputAndOperations[2]),
			operation: inputAndOperations[1],
			output:    getOrCreateWire(output),
		})
	}

	// parse initial values
	initialValues := make(map[string]string)
	for _, line := range strings.Split(parts[0], "\n") {
		nameValue := strings.Split(line, ": ")
		initialValues[nameValue[0]] = nameValue[1]
	}

	return wires, gates, initialValues
}

func getZWiresValue(input string) int64 {
	wires, gates, initialValues := parseInput(input)
	for name, value := range initialValues {
		wire := wires[name]
		wire.value = value
	}

	gatesToProcess := gates
	for len(gatesToProcess) > 0 {
		var gatesWithNoOutputs []*Gate
		for _, gate := range gatesToProcess {
			if output := gate.getOutput(); output != "" {
				gate.output.value = output
			} else {
				gatesWithNoOutputs = append(gatesWithNoOutputs, gate)
			}
		}
		gatesToProcess = gatesWithNoOutputs
	}

	var zWires []*Wire
	for name, wire := range wires {
		if strings.HasPrefix(name, "z") {
			zWires = append(zWires, wire)
		}
	}
	sort.Slice(zWires, func(i, j int) bool {
		// Reverse order to have most significant bit first in resulting string
		return zWires[i].name > zWires[j].name
	})

	valueString := ""
	for _, wire := range zWires {
		valueString += wire.value
	}

	value, err := strconv.ParseInt(valueString, 2, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse %s, error: %v", valueString, err))
	}
	return value
}

func Part1() any {
	input, _ := lib.ReadInput(24)
	return getZWiresValue(input)
}

func Part2() any {
	return "Not implemented"
}
