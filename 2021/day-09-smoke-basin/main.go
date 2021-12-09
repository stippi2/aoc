package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Display struct {
	// readings contains strings with varying number of elements for all ten digits
	readings []string
	// digits contains 4 digits (out of the set of readings) representing the current 4-digit value of the display
	digits []string
	// mapping is the deduced mapping of signals to digit components
	//  aaaa
	// b    c
	// b    c
	//  dddd
	// e    f
	// e    f
	//  gggg
	// If it can be deduced that the signal "e" should be mapped to the digit component "a",
	// then this key-value pair would be contained in mapping.
	mapping map[string]string
}

func maskDigit(digit, mask string) string {
	for _, char := range strings.Split(mask, "") {
		digit = strings.ReplaceAll(digit, char, "")
	}
	return digit
}

func (d *Display) deduceMapping() {
	// Deduce the mapping for "a"
	one := d.findFirstDigit(2)
	seven := d.findFirstDigit(3)
	a := maskDigit(seven, one)
	d.mapping["a"] = a

	// Removing the elements of 7 from those of 4 leaves "bd"
	four := d.findFirstDigit(4)
	bd := maskDigit(four, seven)

	// countPerElement tracks how often a given element occurs across all readings
	// In an unscrambled display, the distribution is like this:
	// a = 8
	// b = 6
	// c = 8
	// d = 7
	// e = 4
	// f = 9
	// g = 7
	countPerElement := make(map[string]int)
	for _, reading := range d.readings {
		for _, element := range strings.Split(reading, "") {
			countPerElement[element]++
		}
	}

	// Deduce the elements by occurrences
	for element, count := range countPerElement {
		switch count {
		case 4:
			d.mapping["e"] = element
		case 6:
			d.mapping["b"] = element
		case 7:
			if strings.Contains(bd, element) {
				d.mapping["d"] = element
			} else {
				d.mapping["g"] = element
			}
		case 8:
			if element != d.mapping["a"] {
				d.mapping["c"] = element
			}
		case 9:
			d.mapping["f"] = element
		}
	}
}

func (d *Display) findFirstDigit(length int) string {
	for _, reading := range d.readings {
		if len(reading) == length {
			return reading
		}
	}
	return ""
}

func countDigits(displays []Display, condition func(digit string) bool) int {
	sum := 0
	for _, display := range displays {
		for _, digit := range display.digits {
			if condition(digit) {
				sum++
			}
		}
	}
	return sum
}

func conditionOnesFoursSevensAndEights(digit string) bool {
	switch len(digit) {
	case 2, 4, 3, 7:
		return true
	}
	return false
}

func (d *Display) scramble(unscrambled string) string {
	scrambled := ""
	for _, element := range strings.Split(unscrambled, "") {
		scrambled += d.mapping[element]
	}
	return scrambled
}

func (d *Display) digitToInt() map[string]int {
	digitToInt := make(map[string]int)
	// 0: abc.efg  6
	// 1: ..c..f.  2 *
	// 2: a.cde.g  5
	// 3: a.cd.fg  5
	// 4: .bcd.f.  4 *
	// 5: ab.d.fg  5
	// 6: ab.defg  6
	// 7: a.c..f.  3 *
	// 8: abcdefg  7 *
	// 9: abcd.fg  6
	digitToInt[d.scramble("abcefg")] = 0
	digitToInt[d.scramble("cf")] = 1
	digitToInt[d.scramble("acdeg")] = 2
	digitToInt[d.scramble("acdfg")] = 3
	digitToInt[d.scramble("bcdf")] = 4
	digitToInt[d.scramble("abdfg")] = 5
	digitToInt[d.scramble("abdefg")] = 6
	digitToInt[d.scramble("acf")] = 7
	digitToInt[d.scramble("abcdefg")] = 8
	digitToInt[d.scramble("abcdfg")] = 9
	return digitToInt
}

func hasSameChars(a, b string) bool {
	// Assumes each char is only contained once in each string!
	if len(a) != len(b) {
		return false
	}
	found := 0
	for _, char := range strings.Split(a, "") {
		if strings.Contains(b, char) {
			found++
		}
	}
	return found == len(a)
}

func findValue(scrambledDigit string, mapping map[string]int) int {
	for key, value := range mapping {
		if hasSameChars(key, scrambledDigit) {
			return value
		}
	}
	panic("did not find value in mapping")
}

func (d *Display) descramble() int {
	mapping := d.digitToInt()
	result := 0
	result += findValue(d.digits[0], mapping) * 1000
	result += findValue(d.digits[1], mapping) * 100
	result += findValue(d.digits[2], mapping) * 10
	result += findValue(d.digits[3], mapping) * 1
	return result
}

func main() {
	displays := parseInput(loadInput("puzzle-input.txt"))
	count := countDigits(displays, conditionOnesFoursSevensAndEights)
	fmt.Printf("number of 1, 4, 7, and 8: %v\n", count)

	sum := 0
	for _, d := range displays {
		d.deduceMapping()
		sum += d.descramble()
	}
	fmt.Printf("sum of all output values: %v\n", sum)
}

func parseInput(input string) (displays []Display) {
	lines := strings.Split(input, "\n")
	displays = make([]Display, len(lines))
	for i, line := range lines {
		signalsDigits := strings.Split(line, " | ")
		displays[i].readings = strings.Split(signalsDigits[0], " ")
		displays[i].digits = strings.Split(signalsDigits[1], " ")
		displays[i].mapping = make(map[string]string)
	}
	return
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
