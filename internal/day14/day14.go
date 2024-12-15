package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pluque01/aoc24/pkg/reader"
	"regexp"
	"strconv"
)

type Tile struct {
	row int
	col int
}

type BathroomMap struct {
	tiles  [][]int
	width  int
	height int
}

type RobotMovement struct {
	Row int
	Col int
}

type Robot struct {
	initialPosition Tile
	movement        RobotMovement
}

func NewBathroomMap(width, height int) *BathroomMap {
	tiles := make([][]int, height)
	for i := range tiles {
		tiles[i] = make([]int, width)
	}
	return &BathroomMap{
		tiles:  tiles,
		width:  width,
		height: height,
	}
}

func NewRobot(line string) *Robot {
	rg := regexp.MustCompile(`-?\d+`)
	values := rg.FindAllString(line, -1)
	if len(values) != 4 {
		log.Panicf("Invalid input for robot: %s", line)
	}
	valuesInt := make([]int, len(values))
	for i, v := range values {
		var err error
		valuesInt[i], err = strconv.Atoi(v)
		if err != nil {
			log.Panicf("Invalid input for robot: %s", line)
		}
	}
	return &Robot{
		initialPosition: Tile{row: valuesInt[1], col: valuesInt[0]},
		movement:        RobotMovement{Row: valuesInt[3], Col: valuesInt[2]},
	}
}

func (bm *BathroomMap) MoveRobot(robot *Robot, steps int) {
	posRow := (robot.initialPosition.row + (robot.movement.Row * steps)) % bm.height
	posCol := (robot.initialPosition.col + (robot.movement.Col * steps)) % bm.width
	// The position must be between 0 and the width or height
	if posRow < 0 {
		posRow = bm.height + posRow
	}
	if posCol < 0 {
		posCol = bm.width + posCol
	}
	(*bm).tiles[posRow][posCol]++
}

func (bm *BathroomMap) IsPossibleTree() bool {
	for i := 0; i < bm.height; i++ {
		foundBot := false
		consecutiveNeighbors := 0
		for j := 0; j < bm.width; j++ {
			if bm.tiles[i][j] > 0 {
				consecutiveNeighbors++
				if !foundBot {
					foundBot = true
				}
			}
			if foundBot && bm.tiles[i][j] == 0 {
				break
			}
		}
		if consecutiveNeighbors >= 8 {
			return true
		}
	}
	return false
}

func (bm *BathroomMap) PrintMap() {
	for i := 0; i < bm.height; i++ {
		for j := 0; j < bm.width; j++ {
			if bm.tiles[i][j] == 0 {
				print(".")
			} else {
				print("#")
			}
		}
		println()
	}
}

func (bm *BathroomMap) CleanMap() {
	for i := 0; i < bm.height; i++ {
		for j := 0; j < bm.width; j++ {
			bm.tiles[i][j] = 0
		}
	}
}

func (bm *BathroomMap) CalculateSafetyFactor() int {
	cuadrant1 := 0
	cuadrant2 := 0
	cuadrant3 := 0
	cuadrant4 := 0
	for i := 0; i < bm.height/2; i++ {
		for j := 0; j < bm.width/2; j++ {
			cuadrant1 += bm.tiles[i][j]
		}
		for j := bm.width/2 + 1; j < bm.width; j++ {
			cuadrant2 += bm.tiles[i][j]
		}
	}
	for i := bm.height/2 + 1; i < bm.height; i++ {
		for j := 0; j < bm.width/2; j++ {
			cuadrant3 += bm.tiles[i][j]
		}
		for j := bm.width/2 + 1; j < bm.width; j++ {
			cuadrant4 += bm.tiles[i][j]
		}
	}
	return cuadrant1 * cuadrant2 * cuadrant3 * cuadrant4
}

func Solution1() int {
	data, err := reader.ReadStringByLineFile("/home/fallen/code/aoc24/inputs/day14.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	bm := NewBathroomMap(101, 103)
	for i := 0; i < len(data); i++ {
		robot := NewRobot(data[i])
		bm.MoveRobot(robot, 100)
	}
	return bm.CalculateSafetyFactor()
}

func Solution2() int {
	data, err := reader.ReadStringByLineFile("/home/fallen/code/aoc24/inputs/day14.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	bm := NewBathroomMap(101, 103)
	robots := make([]*Robot, len(data))
	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < len(data); i++ {
		robots[i] = NewRobot(data[i])
	}
	for i := 0; i < 10000; i++ {
		for _, r := range robots {
			bm.MoveRobot(r, i)
		}
		if bm.IsPossibleTree() {
			bm.PrintMap()
			fmt.Printf("Iteration: %d, Press enter to continue\n", i)
			input.Scan()
		}
		bm.CleanMap()
	}

	return 0
}
