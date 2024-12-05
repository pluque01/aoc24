package day04

import (
	"log"
	"pluque01/aoc24/pkg/reader"
)

// Finds if the target can be found from a given cell in the grid in all directions
func Search2D(grid [][]rune, col int, row int, target []rune) int {
	if grid[col][row] != target[0] {
		return 0
	}
	x_directions := []int{0, 1, 1, 1, 0, -1, -1, -1}
	y_directions := []int{-1, -1, 0, 1, 1, 1, 0, -1}

	finds := 0
	for i := 0; i < 8; i++ {
		var k, rd, cd int
		rd = row + x_directions[i]
		cd = col + y_directions[i]
		for k = 1; k < len(target); k++ {
			// Check if the target is out of bounds
			if rd >= len(grid) || rd < 0 || cd >= len(grid[0]) || cd < 0 {
				break
			}
			if grid[cd][rd] != target[k] {
				break
			}
			// Keep moving in the direction
			rd += x_directions[i]
			cd += y_directions[i]
		}
		if k == len(target) {
			finds++
		}
	}
	// At the end retrun the number of finds
	return finds
}

func CountWord2D(grid [][]rune, target string) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			count += Search2D(grid, i, j, []rune(target))
		}
	}
	return count
}

func Solution1() int {
	content, err := reader.ReadCharFile("./inputs/day04.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	return CountWord2D(content, "XMAS")
}

func SearchCross2D(grid [][]rune, col int, row int, target []rune) bool {
	if len(target) == 0 || len(target)%2 == 0 {
		return false
	}
	// Check if the middle rune of the target is at the current cell
	if target[len(target)/2] != grid[row][col] {
		return false
	}

	x_directions := []int{1, 1, -1, -1}
	y_directions := []int{-1, 1, 1, -1}
	middleRuneIndex := len(target) / 2
	diagonalsCompleted := 0
	for i := 0; i < len(x_directions); i++ {
		var k, rd, cd, rdi, cdi int
		rd = row + y_directions[i]
		cd = col + x_directions[i]

		rdi = row - y_directions[i]
		cdi = col - x_directions[i]

		for k = 0; k < len(target)/2; k++ {
			// Check if the target is out of bounds
			if rd >= len(grid) || rd < 0 || cd >= len(grid[0]) || cd < 0 || rdi >= len(grid) || rdi < 0 || cdi >= len(grid[0]) || cdi < 0 {
				return false
			}
			if grid[rd][cd] != target[middleRuneIndex+(k+1)] || grid[rdi][cdi] != target[middleRuneIndex-(k+1)] {
				break
			}
			// Keep moving in the direction
			rd += x_directions[i]
			cd += y_directions[i]
			rdi -= x_directions[i]
			cdi -= y_directions[i]
		}
		if k == len(target)/2 {
			diagonalsCompleted++
		}
		if diagonalsCompleted == 2 {
			return true
		}
	}

	return false
}

func CountCrossWord2D(grid [][]rune, target string) int {
	targetRunes := []rune(target)
	if len(targetRunes) == 0 || len(targetRunes)%2 == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == targetRunes[len(targetRunes)/2] {
				if SearchCross2D(grid, j, i, targetRunes) {
					count++
				}
			}
		}
	}
	return count
}

func Solution2() int {
	content, err := reader.ReadCharFile("./inputs/day04.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	return CountCrossWord2D(content, "MAS")
}
