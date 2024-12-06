package day06

import (
	"fmt"
	"log"
	"pluque01/aoc24/pkg/reader"
)

func SimulateMovement(m *[][]rune, x int, y int) {
	direction := []struct {
		x int
		y int
	}{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}
	currentDirection := 0

	outOfBounds := func(x, y int) bool {
		return x < 0 || y < 0 || x >= len(*m) || y >= len((*m)[0])
	}
	nx := x + direction[currentDirection].x
	ny := y + direction[currentDirection].y
	for !outOfBounds(nx, ny) {
		if (*m)[ny][nx] == '#' {
			currentDirection = (currentDirection + 1) % 4
		} else {
			x = nx
			y = ny
			(*m)[y][x] = 'X'
		}
		nx = x + direction[currentDirection].x
		ny = y + direction[currentDirection].y
	}
}

func countVisitedCells(m *[][]rune) int {
	count := 0
	for _, row := range *m {
		for _, cell := range row {
			if cell == 'X' {
				count++
			}
		}
	}
	return count
}

func FindStartingPosition(m *[][]rune) (int, int, error) {
	for y, row := range *m {
		for x, cell := range row {
			if cell == '^' {
				return x, y, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("Starting position not found")
}

type visitedCell struct {
	x         int
	y         int
	direction int
}

func isLoop(m *[][]rune, x int, y int) bool {
	direction := []struct {
		x int
		y int
	}{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}
	currentDirection := 0
	visitedCells := make(map[visitedCell]struct{})
	visitedCells[visitedCell{x, y, currentDirection}] = struct{}{}

	outOfBounds := func(x, y int) bool {
		return x < 0 || y < 0 || x >= len(*m) || y >= len((*m)[0])
	}
	nx := x + direction[currentDirection].x
	ny := y + direction[currentDirection].y
	for !outOfBounds(nx, ny) {
		if _, ok := visitedCells[visitedCell{nx, ny, currentDirection}]; ok {
			return true
		}
		if (*m)[ny][nx] == '#' {
			currentDirection = (currentDirection + 1) % 4
		} else {
			x = nx
			y = ny
			visitedCells[visitedCell{x, y, currentDirection}] = struct{}{}
		}
		nx = x + direction[currentDirection].x
		ny = y + direction[currentDirection].y
	}
	return false
}

func CountLoops(m *[][]rune, x int, y int) int {
	countLoops := 0
	for i, row := range *m {
		for j, cell := range row {
			if cell == '.' {
				(*m)[i][j] = '#'
				if isLoop(m, x, y) {
					countLoops++
				}
				(*m)[i][j] = '.'
			}
		}
	}
	return countLoops
}

func Solution1() int {
	m, err := reader.ReadCharFile("./inputs/day06.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	x, y, err := FindStartingPosition(&m)
	if err != nil {
		log.Fatalf("Failed to find starting position: %v", err)
	}
	SimulateMovement(&m, x, y)
	return countVisitedCells(&m)
}

func Solution2() int {
	m, err := reader.ReadCharFile("./inputs/day06.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	x, y, err := FindStartingPosition(&m)
	return CountLoops(&m, x, y)
}
