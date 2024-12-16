package day15

import (
	"log"
	"pluque01/aoc24/pkg/reader"
)

type Robot struct {
	row int
	col int
}

type Warehouse struct {
	grid  [][]rune
	robot Robot
}
type BigWarehouse struct {
	grid  [][]rune
	robot Robot
}

type Movement int

const (
	North Movement = iota
	East
	South
	West
)

var movementName = map[Movement]string{
	North: "North",
	East:  "East",
	South: "South",
	West:  "West",
}
var movementValue = map[rune]Movement{
	'^': North,
	'>': East,
	'v': South,
	'<': West,
}
var movementDelta = map[Movement][2]int{
	North: {0, -1},
	East:  {1, 0},
	South: {0, 1},
	West:  {-1, 0},
}

func (m Movement) String() string {
	return movementName[m]
}
func (m Movement) Delta() [2]int {
	return movementDelta[m]
}

func NewWarehouse(input [][]rune) *Warehouse {
	foundRobot := false
	var i, j int
	var row []rune
	var value rune
	for i, row = range input {
		for j, value = range row {
			if value == '@' {
				foundRobot = true
				break
			}
		}
		if foundRobot {
			break
		}
	}
	if !foundRobot {
		log.Fatal("Robot not found!")
	}
	return &Warehouse{
		grid:  input,
		robot: Robot{i, j},
	}
}

func (wh *Warehouse) GetCell(row, col int) rune {
	return wh.grid[row][col]
}
func (wh *BigWarehouse) GetCell(row, col int) rune {
	return wh.grid[row][col]
}

func (wh *Warehouse) MoveRobot(row, col int) {
	wh.grid[row][col] = '@'
	wh.grid[wh.robot.row][wh.robot.col] = '.'
	wh.robot.row = row
	wh.robot.col = col
}

func (wh *Warehouse) ProcessMovement(movement Movement) {
	keepLooking := false
	var firstBoxRow, firstBoxCol int
	nextRow := wh.robot.row + movement.Delta()[1]
	nextCol := wh.robot.col + movement.Delta()[0]
	cellValue := wh.GetCell(nextRow, nextCol)

	if cellValue == '.' {
		wh.MoveRobot(nextRow, nextCol)
	} else if cellValue == 'O' {
		firstBoxRow = nextRow
		firstBoxCol = nextCol
		keepLooking = true
	}
	for keepLooking {
		nextRow += movement.Delta()[1]
		nextCol += movement.Delta()[0]
		cellValue = wh.GetCell(nextRow, nextCol)
		if cellValue == '#' {
			keepLooking = false
		} else if cellValue == '.' {
			wh.grid[nextRow][nextCol] = 'O'
			wh.MoveRobot(firstBoxRow, firstBoxCol)
			keepLooking = false
		}
	}
}

func (wh *Warehouse) PrintWarehouse() {
	for _, row := range wh.grid {
		for _, value := range row {
			print(string(value))
		}
		println()
	}
}

func (wh *Warehouse) SumGPSCoordinates() int {
	total := 0
	for i, row := range wh.grid {
		for j, value := range row {
			if value == 'O' {
				total += 100*i + j
			}
		}
	}
	return total
}

func NewBigWarehouse(input [][]rune) *BigWarehouse {
	grid := make([][]rune, len(input))
	for i := range grid {
		grid[i] = make([]rune, len(input[0])*2)
	}
	var r, c int
	var robotRow, robotCol int
	for _, row := range input {
		for _, value := range row {
			switch value {
			case '#':
				grid[r][c] = '#'
				grid[r][c+1] = '#'
				break
			case '.':
				grid[r][c] = '.'
				grid[r][c+1] = '.'
				break
			case 'O':
				grid[r][c] = '['
				grid[r][c+1] = ']'
				break
			case '@':
				grid[r][c] = '@'
				grid[r][c+1] = '.'
				robotRow = r
				robotCol = c
				break
			}
			c += 2
		}
		r++
	}
	return &BigWarehouse{
		grid:  grid,
		robot: Robot{robotRow, robotCol},
	}
}

func (wh *BigWarehouse) MoveRobot(row, col int) {
	wh.grid[row][col] = '@'
	wh.grid[wh.robot.row][wh.robot.col] = '.'
	wh.robot.row = row
	wh.robot.col = col
}

type Cell struct {
	row int
	col int
}

type Box struct {
	left  Cell
	right Cell
}

func (wh *BigWarehouse) GetBoxNeighbours(row, col int) (Cell, Cell) {
	left := Cell{row, col}
	right := Cell{row, col + 1}
	return left, right
}

func (wh *BigWarehouse) MoveBox(row, col int, direction Movement) {
	wh.grid[row+direction.Delta()[1]][col+direction.Delta()[0]] = '['
	wh.grid[row+direction.Delta()[1]][col+direction.Delta()[0]+1] = ']'
	wh.grid[row][col] = '.'
	wh.grid[row][col+1] = '.'
}

func (wh *BigWarehouse) CanBoxBeMoved(row, col int, direction Movement) bool {
	initialValue := wh.GetCell(row, col)
	var leftBox, rightBox Cell
	if initialValue == '[' {
		leftBox = Cell{row, col}
		rightBox = Cell{row, col + 1}
	} else if initialValue == ']' {
		leftBox = Cell{row, col - 1}
		rightBox = Cell{row, col}
	}
	switch direction {
	case North:
		leftNorth := wh.GetCell(leftBox.row-1, leftBox.col)
		rightNorth := wh.GetCell(rightBox.row-1, rightBox.col)
		if leftNorth == '.' && rightNorth == '.' {
			wh.grid[leftBox.row][leftBox.col] = '.'
			wh.grid[rightBox.row][rightBox.col] = '.'
			wh.grid[leftBox.row-1][leftBox.col] = '['
			wh.grid[rightBox.row-1][rightBox.col] = ']'
			return true
		} else if leftNorth == '#' || rightNorth == '#' {
			return false
		} else if leftNorth == '[' || rightNorth == ']' {
			if wh.MoveBox(leftBox.row-1, leftBox.col, North) {
				wh.grid[leftBox.row][leftBox.col] = '.'
				wh.grid[rightBox.row][rightBox.col] = '.'
				wh.grid[leftBox.row-1][leftBox.col] = '['
				wh.grid[rightBox.row-1][rightBox.col] = ']'
				return true
			}
		} else if leftNorth == ']' && rightNorth == '[' {
			if wh.MoveBox(leftBox.row-1, leftBox.col, North) && wh.MoveBox(rightBox.row-1, rightBox.col, North) {
				wh.grid[leftBox.row][leftBox.col] = '.'
				wh.grid[rightBox.row][rightBox.col] = '.'
				wh.grid[leftBox.row-1][leftBox.col] = '['
				wh.grid[rightBox.row-1][rightBox.col] = ']'
				return true
			}
		} else if leftNorth == ']' {

		}
		break
	case East:
		if wh.GetCell(rightBox.row, rightBox.col+1) == '.' {

		}
	}
}

func (wh *BigWarehouse) ProcessMovement(movement Movement) {
	keepLooking := false
	var firstBoxRow, firstBoxCol int
	nextRow := wh.robot.row + movement.Delta()[1]
	nextCol := wh.robot.col + movement.Delta()[0]
	cellValue := wh.GetCell(nextRow, nextCol)

	if cellValue == '.' {
		wh.MoveRobot(nextRow, nextCol)
	} else if cellValue == 'O' {
		firstBoxRow = nextRow
		firstBoxCol = nextCol
		keepLooking = true
	}
	for keepLooking {
		nextRow += movement.Delta()[1]
		nextCol += movement.Delta()[0]
		cellValue = wh.GetCell(nextRow, nextCol)
		if cellValue == '#' {
			keepLooking = false
		} else if cellValue == '.' {
			wh.grid[nextRow][nextCol] = 'O'
			wh.MoveRobot(firstBoxRow, firstBoxCol)
			keepLooking = false
		}
	}
}

func Solution1() int {
	data, err := reader.ReadCharFile("/home/fallen/code/aoc24/inputs/day15.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	var i int
	for i = 0; i < len(data); i++ {
		if data[i][0] != '#' {
			break
		}
	}
	wh := NewWarehouse(data[:i])
	// input := bufio.NewScanner(os.Stdin)
	for _, row := range data[i:] {
		for _, value := range row {
			if movement, ok := movementValue[value]; ok {
				// fmt.Println("Next Movement: ", movement)
				wh.ProcessMovement(movement)
				// input.Scan()
				// wh.PrintWarehouse()
			} else {
				log.Printf("Invalid character: %v", value)
			}
		}
	}
	return wh.SumGPSCoordinates()
}
