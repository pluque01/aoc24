package day09

import (
	"log"
	"pluque01/aoc24/pkg/reader"
)

func ParseDisk(input *[]rune) *[]rune {
	output := []rune{}
	currentID := 0
	for i := 0; i < len(*input); i++ {
		// First the block file
		if i%2 == 0 {
			for j := 0; j < int((*input)[i]-'0'); j++ {
				output = append(output, rune(currentID+'0'))
			}
			currentID++
		}
		// Then the free space
		if i%2 == 1 {
			for j := 0; j < int((*input)[i]-'0'); j++ {
				output = append(output, '.')
			}
		}
	}
	return &output
}

func findLastBlockIndexFromPosition(input *[]rune, position int) int {
	if position >= len(*input) || position < 0 {
		return 0
	}
	for i := position; i >= 0; i-- {
		if (*input)[i] != rune('.') {
			return i
		}
	}
	return 0
}

func CompactDisk(input *[]rune) {
	j := findLastBlockIndexFromPosition(input, len(*input)-1)
	for i := 0; i < j; i++ {
		if (*input)[i] == '.' {
			(*input)[i] = (*input)[j]
			(*input)[j] = '.'
			j = findLastBlockIndexFromPosition(input, j-1)
		}
	}
}

func GetChecksum(input *[]rune) int {
	checksum := 0
	for i := 0; i < len(*input); i++ {
		if (*input)[i] != '.' {
			checksum += (int((*input)[i]-'0') * i)
		}
	}
	return checksum
}

func findFirstFreeSpaceIndexWithSize(input *[]rune, size int) int {
	candidateSize := 0
	candidateIndex := 0
	for i := 0; i < len(*input); i++ {
		if (*input)[i] == '.' {
			if candidateSize == 0 {
				candidateIndex = i
			}
			candidateSize++
			if candidateSize == size {
				return candidateIndex
			}
		} else if candidateSize > 0 {
			candidateSize = 0
		}
	}
	return len(*input)
}

func findLastBlockFromPosition(input *[]rune, position int) (int, int) {
	lastBlockIndex := findLastBlockIndexFromPosition(input, position)
	lastBLock := (*input)[lastBlockIndex]
	if lastBlockIndex == 0 {
		return 0, 0
	}
	size := 1
	for i := lastBlockIndex - 1; i >= 0; i-- {
		if (*input)[i] != lastBLock {
			break
		}
		size++
	}
	return lastBlockIndex, size
}

func moveBlock(input *[]rune, blockIndex int, size int, freeSpaceIndex int) {
	for i := 0; i < size; i++ {
		(*input)[freeSpaceIndex+i] = (*input)[blockIndex-i]
		(*input)[blockIndex-i] = '.'
	}
}

func CompactDiskWithoutFragmentation(input *[]rune) {
	startingIndex := len(*input) - 1
	blockIndex, size := findLastBlockFromPosition(input, startingIndex)
	for blockIndex > 0 {
		freeSpaceIndex := findFirstFreeSpaceIndexWithSize(input, size)
		if freeSpaceIndex < blockIndex {
			moveBlock(input, blockIndex, size, freeSpaceIndex)
		}
		startingIndex = blockIndex - size
		blockIndex, size = findLastBlockFromPosition(input, startingIndex)
	}
}

func Solution1() int {
	data, err := reader.ReadCharFile("./inputs/day09.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	if len(data) > 1 {
		log.Printf("Input file has too many lines, merging all of them")
		// Merge all lines into a single line
		merged := []rune{}
		for i := 0; i < len(data); i++ {
			merged = append(merged, data[i]...)
		}
		data = [][]rune{merged}
	}
	disk := ParseDisk(&data[0])
	CompactDisk(disk)
	return GetChecksum(disk)
}

func Solution2() int {
	data, err := reader.ReadCharFile("./inputs/day09.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	if len(data) > 1 {
		log.Printf("Input file has too many lines, merging all of them")
		// Merge all lines into a single line
		merged := []rune{}
		for i := 0; i < len(data); i++ {
			merged = append(merged, data[i]...)
		}
		data = [][]rune{merged}
	}
	disk := ParseDisk(&data[0])
	CompactDiskWithoutFragmentation(disk)
	return GetChecksum(disk)

	// 3986571717772 too low
}
