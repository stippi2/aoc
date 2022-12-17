package main

import (
	"fmt"
	"os"
)

type Sequence struct {
	input string
	pos   int
}

func (s *Sequence) next() string {
	if s.pos == len(s.input) {
		s.pos = 0
	}
	s.pos += 1
	return s.input[s.pos-1 : s.pos]
}

const rocks = "-+L|x"

type Row struct {
	columns []bool
}

type Rock struct {
	rows []Row
	x    int
	y    int
}

func (r *Rock) width() int {
	return len(r.rows[0].columns)
}

func newRock(kind string) *Rock {
	switch kind {
	case "-":
		return &Rock{
			rows: []Row{
				{columns: []bool{true, true, true, true}},
			},
		}
	case "+":
		return &Rock{
			rows: []Row{
				{columns: []bool{false, true, false}},
				{columns: []bool{true, true, true}},
				{columns: []bool{false, true, false}},
			},
		}
	case "L":
		return &Rock{
			rows: []Row{
				{columns: []bool{false, false, true}},
				{columns: []bool{false, false, true}},
				{columns: []bool{true, true, true}},
			},
		}
	case "|":
		return &Rock{
			rows: []Row{
				{columns: []bool{true}},
				{columns: []bool{true}},
				{columns: []bool{true}},
				{columns: []bool{true}},
			},
		}
	case "x":
		return &Rock{
			rows: []Row{
				{columns: []bool{true, true}},
				{columns: []bool{true, true}},
			},
		}
	}
	return nil
}

type Chamber struct {
	rows        []Row
	removedRows int64
}

func (c *Chamber) String() string {
	result := ""
	for y := len(c.rows) - 1; y > 0; y-- {
		result += "|"
		for x := 0; x < 7; x++ {
			if c.rows[y].columns[x] {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "|\n"
	}
	result += "+-------+\n"
	return result
}

func hitTest(rock *Rock, offsetX, offsetY int, chamber *Chamber) bool {
	if offsetX != 0 {
		// Check chamber boundaries when moving left/right
		if rock.x+offsetX < 0 {
			return true
		}
		if rock.x+rock.width()+offsetX > 7 {
			return true
		}
	}
	if rock.y+offsetY <= 0 {
		// Still above highest settled rock
		return false
	}
	// -3
	// -2                            3 #  r.rows[0]
	// -1                            2 #  r.rows[1]
	// 0    c.row[len(c.rows) - 1]   1 #  r.rows[2]
	// 1    c.row[len(c.rows) - 2]   0 #  r.rows[3]
	// 2    c.row[len(c.rows) - 3]
	// 3
	for y := 0; y < len(rock.rows); y++ {
		offsetFromChamberTop := rock.y + offsetY - len(rock.rows) + y
		if offsetFromChamberTop < 0 {
			continue
		}
		chamberRow := &chamber.rows[len(chamber.rows)-offsetFromChamberTop-1]
		rockRow := &rock.rows[y]
		for x := 0; x < rock.width(); x++ {
			if chamberRow.columns[rock.x+offsetX+x] && rockRow.columns[x] {
				return true
			}
		}
	}
	return false
}

func settleRock(rock *Rock, chamber *Chamber) {
	// -3
	// -2                            3 #  r.rows[0]
	// -1                            2 #  r.rows[1]
	// 0    c.row[len(c.rows) - 1]   1 #  r.rows[2]
	// 1    c.row[len(c.rows) - 2]   0 #  r.rows[3]
	// 2    c.row[len(c.rows) - 3]
	// 3
	chamberHeight := len(chamber.rows)
	for y := len(rock.rows) - 1; y >= 0; y-- {
		offsetFromChamberTop := rock.y - len(rock.rows) + y
		var chamberRow *Row
		if offsetFromChamberTop < 0 {
			chamber.rows = append(chamber.rows, Row{columns: []bool{false, false, false, false, false, false, false}})
			chamberRow = &chamber.rows[len(chamber.rows)-1]
		} else {
			chamberRow = &chamber.rows[chamberHeight-offsetFromChamberTop-1]
		}
		rockRow := &rock.rows[y]
		for x := 0; x < rock.width(); x++ {
			chamberRow.columns[rock.x+x] = chamberRow.columns[rock.x+x] || rockRow.columns[x]
		}
	}
	if len(chamber.rows) > 1000 {
		chamber.rows = chamber.rows[500:]
		chamber.removedRows += 500
	}
}

func moveRock(rock *Rock, chamber *Chamber, jetSequence *Sequence) {
	for {
		offset := 0
		switch jetSequence.next() {
		case "<":
			offset = -1
		case ">":
			offset = 1
		}
		if !hitTest(rock, offset, 0, chamber) {
			rock.x += offset
		}
		if !hitTest(rock, 0, 1, chamber) {
			rock.y += 1
		} else {
			settleRock(rock, chamber)
			break
		}
	}
}

func simulateRocks(jetSequence *Sequence, rockCount int64) int64 {
	chamber := &Chamber{rows: []Row{{columns: []bool{true, true, true, true, true, true, true}}}}

	rockSequence := &Sequence{input: rocks}

	for i := int64(0); i < rockCount; i++ {
		rock := newRock(rockSequence.next())
		rock.x = 2
		rock.y = -3
		moveRock(rock, chamber, jetSequence)
		if i%10000000 == 0 {
			fmt.Printf("%.2f%%\n", float64(i)/float64(rockCount)*100)
		}
	}
	return int64(len(chamber.rows)-1) + chamber.removedRows
}

func main() {
	jetSequence := &Sequence{input: loadInput("puzzle-input.txt")}
	fmt.Printf("tower height after 2022 rocks: %v\n", simulateRocks(jetSequence, 2022))

	fmt.Printf("tower height after 1000000000000 rocks: %v\n", simulateRocks(jetSequence, 1000000000000))
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
