package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19

3 15  0  2 22
9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
2  0 12  3  7
`

func Test_parseBingoInput(t *testing.T) {
	seq, boards := parseBingoInput(exampleInput)
	assert.Equal(t, []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}, seq)
	assert.Len(t, boards, 3)
	assert.Equal(t, 23, boards[0].numberAt(2, 1))
	assert.Equal(t, 7, boards[2].numberAt(4, 4))
}

func Test_playBingo(t *testing.T) {
	turn, boardIndex, score := playBingo(parseBingoInput(exampleInput))
	assert.Equal(t, 11, turn)
	assert.Equal(t, 2, boardIndex)
	assert.Equal(t, 4512, score)
}

func Test_lastBoardToComplete(t *testing.T) {
	turn, boardIndex, score := lastBoardToComplete(parseBingoInput(exampleInput))
	assert.Equal(t, 14, turn)
	assert.Equal(t, 1, boardIndex)
	assert.Equal(t, 1924, score)
}
