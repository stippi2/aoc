package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elv struct {
	items []int
	sum int
}

func main() {
	elves := parseInput(loadInput("puzzle-input.txt"))
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].sum >= elves[j].sum
	})
	fmt.Printf("The most calories carried by any elv: %v\n", elves[0].sum)
	fmt.Printf("The sum of top three most calories: %v\n", elves[0].sum + elves[1].sum + elves[2].sum)
}

func parseInput(input string) []Elv {
	input = strings.TrimSpace(input)
	elvItems := strings.Split(input, "\n\n")
	elves := make([]Elv, len(elvItems))
	for i, items := range elvItems {
		for _, itemString := range strings.Split(items, "\n") {
			item, _ := strconv.Atoi(itemString)
			elves[i].items = append(elves[i].items, item)
			elves[i].sum += item
		}
	}
	return elves
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
