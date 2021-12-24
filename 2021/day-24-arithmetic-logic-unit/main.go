package main

import (
	"io/ioutil"
	"strings"
)

type Instruction struct {

}

func main() {
}

func parseInput(input string) (program []Instruction){
	return nil
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
