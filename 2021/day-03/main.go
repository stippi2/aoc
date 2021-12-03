package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type BitCounts struct {
	ones  int
	zeros int
}

var indexOutOfRangeErr = errors.New("index out of range")

func getBitCountsAtColumn(binaryNumbers []string, column int) (BitCounts, error) {
	result := BitCounts{}
	if len(binaryNumbers) == 0 {
		return result, fmt.Errorf("no input binary numbers")
	}
	if len(binaryNumbers[0]) <= column {
		return result, indexOutOfRangeErr
	}
	for _, binaryNumber := range binaryNumbers {
		if binaryNumber == "" {
			continue
		}

		switch binaryNumber[column] {
		case '0':
			result.zeros++
		case '1':
			result.ones++
		}
	}
	return result, nil
}

func getUsedBits(bitCounts []BitCounts, most bool) string {
	result := ""
	for _, digit := range bitCounts {
		if digit.ones > digit.zeros {
			if most {
				result += "1"
			} else {
				result += "0"
			}
		} else {
			if most {
				result += "0"
			} else {
				result += "1"
			}
		}
	}
	return result
}

func getBitCounts(binaryNumbers []string, columnCount int) []BitCounts {
	var bitCountsByColumns []BitCounts
	for column := 0; column < columnCount; column++ {
		bitCounts, err := getBitCountsAtColumn(binaryNumbers, column)
		if err != nil {
			if err == indexOutOfRangeErr {
				break
			}
			exitIfError(err)
		}
		bitCountsByColumns = append(bitCountsByColumns, bitCounts)
	}
	return bitCountsByColumns
}

func filterBy(binaryNumbers []string, most bool, keepBit uint8, column int) string {
	bitCounts, err := getBitCountsAtColumn(binaryNumbers, column)
	exitIfError(err)

	var filtered []string
	for _, number := range binaryNumbers {
		if number == "" {
			continue
		}
		keep := false
		if bitCounts.zeros == bitCounts.ones {
			keep = keepBit == number[column]
		} else {
			switch number[column] {
			case '0':
				if most {
					keep = bitCounts.zeros > bitCounts.ones
				} else {
					keep = bitCounts.zeros < bitCounts.ones
				}
			case '1':
				if most {
					keep = bitCounts.ones > bitCounts.zeros
				} else {
					keep = bitCounts.ones < bitCounts.zeros
				}
			}
		}
		if keep {
			filtered = append(filtered, number)
		}
	}

	if len(filtered) == 1 {
		return filtered[0]
	}
	if len(filtered) == 0 {
		return ""
	}
	return filterBy(filtered, most, keepBit, column+1)
}

func getGammaAndEpsilon(binaryNumbers []string, columnCount int) (gamma, epsilon string) {
	bitCounts := getBitCounts(binaryNumbers, columnCount)
	gamma = getUsedBits(bitCounts, true)
	epsilon = getUsedBits(bitCounts, false)
	return
}

func getOxygenAndCo2Scrubber(binaryNumbers []string) (oxygen, co2scrubber string) {
	oxygen = filterBy(binaryNumbers, true, '1', 0)
	co2scrubber = filterBy(binaryNumbers, false, '0', 0)
	return
}

func toDecimal(binary string) int {
	result := 0
	for index := 1; index <= len(binary); index++ {
		bit := binary[len(binary)-index]
		if bit == '1' {
			result += int(math.Pow(2, float64(index-1)))
		}
	}
	return result
}

func main() {
	binaryNumbers := loadInput("binary-input.txt")
	gamma, epsilon := getGammaAndEpsilon(binaryNumbers, 12)
	gammaDecimal := toDecimal(gamma)
	epsilonDecimal := toDecimal(epsilon)
	fmt.Printf("gamma: %s (%v), epsilon: %s (%v)\n", gamma, gammaDecimal, epsilon, epsilonDecimal)
	fmt.Printf("power: %v\n", gammaDecimal*epsilonDecimal)

	oxygen, co2scrubber := getOxygenAndCo2Scrubber(binaryNumbers)
	oxygenDecimal := toDecimal(oxygen)
	co2scrubberDecimal := toDecimal(co2scrubber)
	fmt.Printf("oxygen: %s (%v), co2 scrubber: %s (%v)\n", oxygen, oxygenDecimal, co2scrubber, co2scrubberDecimal)
	fmt.Printf("life support: %v\n", oxygenDecimal*co2scrubberDecimal)
}

func loadInput(filename string) []string {
	fileContents, err := ioutil.ReadFile(filename)
	exitIfError(err)

	return strings.Split(string(fileContents), "\n")
}

func exitIfError(err error) {
	if err != nil {
		panic(err)
	}
}
