package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CubeSet struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id    int
	cubes []CubeSet
}

func (g *Game) addCubeSet(s string) {
	cubeSet := CubeSet{}
	re := regexp.MustCompile(`(\d+)\s*(red|blue|green)`)
	for _, match := range re.FindAllStringSubmatch(s, -1) {
		value, _ := strconv.Atoi(match[1])
		switch match[2] {
		case "red":
			cubeSet.red = value
		case "blue":
			cubeSet.blue = value
		case "green":
			cubeSet.green = value
		}
	}
	g.cubes = append(g.cubes, cubeSet)
}

func filter(games []Game, f func(Game) bool) []Game {
	filtered := make([]Game, 0)
	for _, game := range games {
		if f(game) {
			filtered = append(filtered, game)
		}
	}
	return filtered
}

func getPossibleGames(games []Game, maxCubes CubeSet) []Game {
	return filter(games, func(g Game) bool {
		for _, cubeSet := range g.cubes {
			if cubeSet.red > maxCubes.red || cubeSet.green > maxCubes.green || cubeSet.blue > maxCubes.blue {
				return false
			}
		}
		return true
	})
}

func sumIds(games []Game) int {
	sum := 0
	for _, game := range games {
		sum += game.id
	}
	return sum
}

func main() {
	games := parseInput(loadInput("puzzle-input.txt"))
	maxCubes := CubeSet{red: 12, green: 13, blue: 14}
	possibleGames := getPossibleGames(games, maxCubes)
	fmt.Printf("Possible games: %d\n", sumIds(possibleGames))
}

func parseInput(input string) []Game {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	games := make([]Game, len(lines))
	for i, line := range lines {
		games[i] = Game{id: i + 1}
		prefix := fmt.Sprintf("Game %d: ", i+1)
		rest := strings.TrimPrefix(line, prefix)
		cubeSets := strings.Split(rest, "; ")
		for _, cubeSet := range cubeSets {
			games[i].addCubeSet(cubeSet)
		}
	}
	return games
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
