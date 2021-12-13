package main

import (
	"io/ioutil"
	"strings"
)

type Origami struct {

}

func main() {
}

func parseInput(input string) *Origami {
	return nil
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
