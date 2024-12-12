package day12

import (
	"log"
	"pluque01/aoc24/pkg/reader"
)

type Cell struct {
	row int
	col int
}

type Region struct {
	cells     map[Cell]struct{}
	plantType rune
}

type Garden struct {
	regions []Region
	data    [][]rune
	width   int
	height  int
}

func findCellWithoutRegion(cellsWithRegion [][]bool) (Cell, bool) {
	for i, row := range cellsWithRegion {
		for j, cell := range row {
			if !cell {
				return Cell{i, j}, true
			}
		}
	}
	return Cell{}, false
}

func GetNeighbors(cell Cell, width, height int) *[]Cell {
	neighbors := make([]Cell, 0)
	if cell.row > 0 {
		neighbors = append(neighbors, Cell{cell.row - 1, cell.col})
	}
	if cell.row < height-1 {
		neighbors = append(neighbors, Cell{cell.row + 1, cell.col})
	}
	if cell.col > 0 {
		neighbors = append(neighbors, Cell{cell.row, cell.col - 1})
	}
	if cell.col < width-1 {
		neighbors = append(neighbors, Cell{cell.row, cell.col + 1})
	}
	return &neighbors
}

func NewGarden(data [][]rune) *Garden {
	height := len(data)
	width := len(data[0])

	// Create a 2D array of bools to keep track of which cells are already
	// assigned to a region.
	cellsWithRegion := make([][]bool, height)
	for i := range cellsWithRegion {
		cellsWithRegion[i] = make([]bool, width)
	}

	regions := make([]Region, 0)

	// Get the first element without a region and create a new region for it
	cell, found := findCellWithoutRegion(cellsWithRegion)
	for found {
		region := Region{
			cells:     make(map[Cell]struct{}),
			plantType: data[cell.row][cell.col],
		}

		cellsInQueue := make(map[Cell]struct{})
		cellsToProcess := []Cell{cell}
		for len(cellsToProcess) > 0 {
			// pop the first element
			c := cellsToProcess[0]
			cellsToProcess = cellsToProcess[1:]
			// fmt.Printf("Processing cell: %v\n", c)
			// if the cell is the same type add to the region
			region.cells[c] = struct{}{}
			cellsWithRegion[c.row][c.col] = true
			// Add the neighbors to the list of cells to process
			neighbors := GetNeighbors(c, width, height)
			for _, n := range *neighbors {
				if data[n.row][n.col] == region.plantType {
					// if _, ok := cellsProcessed[n]; !ok {
					if _, ok := cellsInQueue[n]; !ok {
						if !cellsWithRegion[n.row][n.col] {
							// fmt.Printf("Cell %v not processed yet\n", n)
							cellsToProcess = append(cellsToProcess, n)
							cellsInQueue[n] = struct{}{}
						}
					}
				}
			}
			// cellsProcessed[c] = struct{}{}
		}
		// Add the region to the garden
		regions = append(regions, region)
		size := region.GetArea()
		log.Printf("Added region with size: %v\n", size)
		// Find the next cell without a region
		cell, found = findCellWithoutRegion(cellsWithRegion)
	}

	return &Garden{
		regions: regions,
		data:    data,
		width:   width,
		height:  height,
	}
}

func (r *Region) GetArea() int {
	return len(r.cells)
}

func (r *Region) GetPerimeter(width, height int) int {
	totalPerimeter := 0
	for cell := range r.cells {
		cellPerimeter := 4
		neighbors := GetNeighbors(cell, width, height)
		for _, neighbor := range *neighbors {
			if _, ok := r.cells[neighbor]; ok {
				cellPerimeter--
			}
		}
		totalPerimeter += cellPerimeter
	}
	return totalPerimeter
}

func (r *Region) GetSides(width, height int) int {
	directions := []struct {
		row int
		col int
	}{
		{0, -1}, // left
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // bottom
	}
	totalSides := 0
	for cell := range r.cells {
		for i := 0; i < 4; i++ {
			n1 := Cell{cell.row + directions[i].row, cell.col + directions[i].col}
			// if no adjacent cell
			if _, ok := r.cells[n1]; !ok {
				n2 := Cell{cell.row + directions[(i+1)%4].row, cell.col + directions[(i+1)%4].col}
				if _, ok := r.cells[n2]; !ok {
					totalSides++
				} else {
					// Check diagonal
					n3 := Cell{
						cell.row + directions[i].row + directions[(i+1)%4].row,
						cell.col + directions[i].col + directions[(i+1)%4].col,
					}
					// if there is cell on diagonal
					if _, ok := r.cells[n3]; ok {
						totalSides++
					}
				}
			}
		}
	}
	return totalSides
}

func (g *Garden) GetTotalPrice() int {
	totalPrice := 0
	for _, region := range g.regions {
		totalPrice += region.GetArea() * region.GetPerimeter(g.width, g.height)
	}
	return totalPrice
}

func (g *Garden) GetTotalPrice2() int {
	totalPrice := 0
	for _, region := range g.regions {
		totalPrice += region.GetArea() * region.GetSides(g.width, g.height)
	}
	return totalPrice
}

func Solution1() int {
	data, err := reader.ReadCharFile("/home/fallen/code/aoc24/inputs/day12.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	garden := NewGarden(data)
	return garden.GetTotalPrice()
}

func Solution2() int {
	data, err := reader.ReadCharFile("/home/fallen/code/aoc24/inputs/day12.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	garden := NewGarden(data)
	return garden.GetTotalPrice2()
}
