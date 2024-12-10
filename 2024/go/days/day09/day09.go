package day09

import (
	"aoc/2024/go/lib"
)

func calculateDiskSize(input string) int64 {
	size := int64(0)
	for i := 0; i < len(input); i++ {
		size += int64(input[i] - '0')
	}
	return size
}

func expandBlocks(input string, disk []int) {
	currentFileID := 0
	inFile := true
	offset := 0
	for i := 0; i < len(input); i++ {
		repeat := int64(input[i] - '0')
		for repeat > 0 {
			if inFile {
				disk[offset] = currentFileID
			} else {
				disk[offset] = -1
			}
			offset++
			repeat--
		}
		if inFile {
			currentFileID++
		}
		inFile = !inFile
	}
}

func compactDisk(disk []int) {
	offsetFront := 0
	offsetBack := len(disk) - 1

	// Move the back offset to the end of the last file in case the disk ends in free space
	moveToEndOfNextFile := func() {
		for disk[offsetBack] == -1 {
			offsetBack--
		}
	}
	moveToEndOfNextFile()

	for offsetBack > offsetFront {
		for offsetBack > offsetFront && disk[offsetFront] == -1 {
			disk[offsetFront] = disk[offsetBack]
			disk[offsetBack] = -1
			offsetFront++
			offsetBack--
			moveToEndOfNextFile()
		}
		for offsetBack > offsetFront && disk[offsetFront] != -1 {
			offsetFront++
		}
	}
}

func defragmentDisk(disk []int) {
	offsetBack := len(disk) - 1
	moveToEndOfNextFile := func() {
		for disk[offsetBack] == -1 {
			offsetBack--
		}
	}
	moveToEndOfNextFile()

	moveToNextFreeSpace := func(offset int) int {
		for offset < len(disk) && disk[offset] != -1 {
			offset++
		}
		return offset
	}

	slotLength := func(offset, direction int) int {
		length := 0
		fileID := disk[offset]
		for offset >= 0 && offset < len(disk) && disk[offset] == fileID {
			length += 1
			offset += direction
		}
		return length
	}

	for offsetBack > 0 {
		fileLength := slotLength(offsetBack, -1)
		offsetFront := moveToNextFreeSpace(0)
		fileMoved := false
		for offsetFront < offsetBack-fileLength {
			spaceLength := slotLength(offsetFront, 1)
			if spaceLength >= fileLength {
				fileID := disk[offsetBack]
				for i := 0; i < fileLength; i++ {
					disk[offsetFront] = fileID
					disk[offsetBack] = -1
					offsetFront++
					offsetBack--
				}
				fileMoved = true
				moveToEndOfNextFile()
				break
			} else {
				offsetFront = moveToNextFreeSpace(offsetFront + spaceLength)
			}
		}
		if !fileMoved {
			offsetBack -= fileLength
		}
		if offsetBack < 0 {
			break
		}
		moveToEndOfNextFile()
	}
}

func calculateChecksum(disk []int) int64 {
	checksum := int64(0)
	offset := 0
	for offset < len(disk) {
		fileID := disk[offset]
		if fileID != -1 {
			checksum += int64(offset * fileID)
		}
		offset++
	}
	return checksum
}

func Part1() interface{} {
	input, _ := lib.ReadInput(9)
	disk := make([]int, calculateDiskSize(input))
	expandBlocks(input, disk)
	compactDisk(disk)
	return calculateChecksum(disk)
}

func Part2() interface{} {
	input, _ := lib.ReadInput(9)
	disk := make([]int, calculateDiskSize(input))
	expandBlocks(input, disk)
	defragmentDisk(disk)
	return calculateChecksum(disk)
}
