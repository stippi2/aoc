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

func (s *Sequence) skip(amount int) {
	s.pos = (s.pos + amount) % len(s.input)
}

const rocks = "-+L|x"

type Row struct {
	columns []bool
}

func (r *Row) String() string {
	result := ""
	for x := 0; x < 7; x++ {
		if r.columns[x] {
			result += "#"
		} else {
			result += "."
		}
	}
	return result
}

func (r *Row) equals(other *Row) bool {
	for x := 0; x < 7; x++ {
		if r.columns[x] != other.columns[x] {
			return false
		}
	}
	return true
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
	rows []Row
}

func (c *Chamber) String() string {
	result := ""
	for y := len(c.rows) - 1; y > 0; y-- {
		result += "|"
		result += c.rows[y].String()
		result += "|\n"
	}
	result += "+-------+\n"
	return result
}

func (c *Chamber) height() int {
	return len(c.rows) - 1
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

func newChamber() *Chamber {
	return &Chamber{rows: []Row{{columns: []bool{true, true, true, true, true, true, true}}}}
}

func simulateRocks(chamber *Chamber, jetSequence, rockSequence *Sequence, rockCount int) int {
	for i := 0; i < rockCount; i++ {
		rock := newRock(rockSequence.next())
		rock.x = 2
		rock.y = -3
		moveRock(rock, chamber, jetSequence)
	}
	return chamber.height()
}

func (c *Chamber) isRepeatedSequence() bool {
	offset := (len(c.rows) - 1) / 2
	sequenceRepeats := true
	for i := 1; i < offset+1; i++ {
		if !c.rows[i].equals(&c.rows[i+offset]) {
			sequenceRepeats = false
			break
		}
	}
	return sequenceRepeats
}

func isRepeatedHeightChangeSequence(heightChanges []int) bool {
	if len(heightChanges)%2 != 0 {
		return false
	}
	offset := len(heightChanges) / 2
	sequenceRepeats := true
	for i := 0; i < offset; i++ {
		if heightChanges[i] != heightChanges[i+offset] {
			sequenceRepeats = false
			break
		}
	}
	return sequenceRepeats
}

func reverseInplace(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func findRepeatableSequence(heightChanges []int) []int {
	reverseInplace(heightChanges)
	var repeatableSequence []int
	for i := 0; i < len(heightChanges); i++ {
		if repeatableSequence != nil {
			repeats := true
			for j := 0; j < len(repeatableSequence); j++ {
				if heightChanges[i+j] != repeatableSequence[j] {
					repeats = false
					break
				}
			}
			if repeats {
				return repeatableSequence
			}
		}
		repeatableSequence = append(repeatableSequence, heightChanges[i])
	}
	return nil
}

func partOne() {
	jetSequence := &Sequence{input: loadInput("puzzle-input.txt")}
	rockSequence := &Sequence{input: rocks}
	chamber := newChamber()
	fmt.Printf("tower height after 2022 rocks: %v\n", simulateRocks(chamber, jetSequence, rockSequence, 2022))
}

func partTwo() {
	jetSequence := &Sequence{input: loadInput("puzzle-input.txt")}
	rockSequence := &Sequence{input: rocks}
	chamber := newChamber()

	var heightChanges []int
	probeRockCount := 10000
	lastHeight := 0
	for i := 0; i < probeRockCount; i++ {
		height := simulateRocks(chamber, jetSequence, rockSequence, 1)
		if i%len(rocks) == 0 {
			heightChanges = append(heightChanges, height-lastHeight)
		}
	}
	repeatableSequence := findRepeatableSequence(heightChanges)
	rocksPerCycle := len(repeatableSequence) * len(rocks)
	heightPerCycle := 0
	for _, heightDiff := range repeatableSequence {
		heightPerCycle += heightDiff
	}
	fmt.Printf("found repeatable sequence with height %v, rocks per cycle: %v\n", heightPerCycle, rocksPerCycle)
	//for {
	//	lastHeight := chamber.height()
	//	heightPerCycle = simulateRocks(chamber, jetSequence, rockSequence, len(rocks))
	//	rocksPerCycle += len(rocks)
	//	heightChanges = append(heightChanges, heightPerCycle-lastHeight)
	//	if isRepeatedHeightChangeSequence(heightChanges) {
	//		break
	//	}
	//	//if heightPerCycle > 20 {
	//	//	fmt.Printf("%s\n", chamber)
	//	//}
	//}

	fmt.Printf("cycle height: %v / rocks: %v\n", heightPerCycle, rocksPerCycle)

}

func main() {
	//	partOne()
	partTwo()
	//	fmt.Printf("tower height after 1000000000000 rocks: %v\n", simulateRocks(jetSequence, 1000000000000))
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
