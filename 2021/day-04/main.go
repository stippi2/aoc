package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Board struct {
	size    int
	numbers []int
	hits    []bool
}

func (b *Board) init(size int) {
	b.size = size
	b.numbers = make([]int, size*size)
	b.hits = make([]bool, size*size)
}

func (b *Board) offset(x, y int) int {
	return y*b.size + x
}

func (b *Board) setNumber(x, y, number int) {
	b.numbers[b.offset(x, y)] = number
}

func (b *Board) numberAt(x, y int) int {
	return b.numbers[b.offset(x, y)]
}

func (b *Board) markFieldsWithNumber(number int) bool {
	changed := false
	for y := 0; y < b.size; y++ {
		for x := 0; x < b.size; x++ {
			if b.numberAt(x, y) == number {
				changed = changed || !b.hits[b.offset(x, y)]
				b.hits[b.offset(x, y)] = true
			}
		}
	}
	return changed
}

func (b *Board) testCompletedRows() bool {
	for y := 0; y < b.size; y++ {
		hitsPerRow := 0
		for x := 0; x < b.size; x++ {
			if b.hits[b.offset(x, y)] {
				hitsPerRow++
			}
		}
		if hitsPerRow == 5 {
			return true
		}
	}
	for x := 0; x < b.size; x++ {
		hitsPerColumn := 0
		for y := 0; y < b.size; y++ {
			if b.hits[b.offset(x, y)] {
				hitsPerColumn++
			}
		}
		if hitsPerColumn == 5 {
			return true
		}
	}
	return false
}

func (b *Board) sumOfUnmarkedFields() int {
	sum := 0
	for offset := 0; offset < b.size*b.size; offset++ {
		if !b.hits[offset] {
			sum += b.numbers[offset]
		}
	}
	return sum
}

func (b *Board) setRow(y int, numberRow string) {
	numberRow = strings.TrimSpace(numberRow)
	numbers := strings.Split(numberRow, " ")
	x := 0
	for _, number := range numbers {
		if number == "" {
			// there are double spaces in the rows leading to empty "numbers"
			continue
		}
		v, err := strconv.Atoi(number)
		if err != nil {
			panic(fmt.Sprintf("failed to parse number at column %v ('%s'): %v", x, number, err))
		}
		b.setNumber(x, y, v)
		x++
	}
}

func playBingo(numberSequence []int, boards []Board) (turn, boardIndex, score int) {
	for t, number := range numberSequence {
		for i := 0; i < len(boards); i++ {
			if boards[i].markFieldsWithNumber(number) {
				if boards[i].testCompletedRows() {
					turn = t
					boardIndex = i
					score = boards[i].sumOfUnmarkedFields() * number
					return
				}
			}
		}
	}
	return
}

func remove(numbers []int, number int) []int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == number {
			numbers = append(numbers[:i], numbers[i+1:]...)
		}
	}
	return numbers
}

func lastBoardToComplete(numberSequence []int, boards []Board) (turn, boardIndex, score int) {
	remainingBoards := make([]int, len(boards))
	for i := 0; i < len(boards); i++ {
		remainingBoards[i] = i
	}

	for t, number := range numberSequence {
		var completedBoards []int
		for i := 0; i < len(remainingBoards); i++ {
			board := remainingBoards[i]
			if boards[board].markFieldsWithNumber(number) {
				if boards[board].testCompletedRows() {
					completedBoards = append(completedBoards, board)
				}
			}
		}
		for _, board := range completedBoards {
			remainingBoards = remove(remainingBoards, board)
			if len(remainingBoards) == 0 {
				turn = t
				boardIndex = board
				score = boards[board].sumOfUnmarkedFields() * number
				return
			}
		}
	}
	return
}

func main() {
	seq, boards := parseBingoInput(loadInput("bingo-input.txt"))

	turn, boardIndex, score := playBingo(seq, boards)
	fmt.Printf("winning board: %v (at turn %v), score: %v\n", boardIndex, turn, score)

	turn, boardIndex, score = lastBoardToComplete(seq, boards)
	fmt.Printf("last board: %v (at turn %v), score: %v\n", boardIndex, turn, score)
}

func parseBingoInput(input string) (numberSequence []int, boards []Board) {
	parts := strings.Split(input, "\n\n")
	numbers := strings.Split(strings.TrimSpace(parts[0]), ",")
	numberSequence = make([]int, len(numbers))
	for i, number := range numbers {
		var err error
		numberSequence[i], err = strconv.Atoi(number)
		if err != nil {
			panic(fmt.Sprintf("failed to parse number at index %v ('%s'): %v", i, number, err))
		}
	}

	parts = parts[1:]
	boards = make([]Board, len(parts))
	for i, part := range parts {
		rows := strings.Split(strings.TrimSpace(part), "\n")
		boards[i].init(len(rows))
		for y, row := range rows {
			boards[i].setRow(y, row)
		}
	}
	return
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
