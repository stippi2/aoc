package main

import (
	"io/ioutil"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	a Point
	b Point
}

func main() {
}

func parseVentInput(input string) (lines []Line) {
	parts := strings.Split(input, "\n\n")
	return nil
}

func loadInput(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
	exitIfError(err)

	return strings.Trim(string(fileContents), "\n")
}

func exitIfError(err error) {
	if err != nil {
		panic(err)
	}
}
