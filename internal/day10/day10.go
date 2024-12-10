package day10

import (
	"log"
	"pluque01/aoc24/pkg/math"
	"pluque01/aoc24/pkg/reader"
)

type Cell struct {
	row   int
	col   int
	value int8
}

type TopoMap struct {
	cells  [][]Cell
	width  int
	height int
}

func NewTopoMap(data *[][]int8) *TopoMap {
	cells := make([][]Cell, len(*data))
	for i, row := range *data {
		cellRows := make([]Cell, len((*data)[0]))
		for j, cell := range row {
			cellRows[j] = Cell{
				row:   i,
				col:   j,
				value: cell,
			}
		}
		cells[i] = cellRows
	}
	return &TopoMap{
		cells:  cells,
		width:  len((*data)[0]),
		height: len(*data),
	}
}
func NewTopoMapWithRune(data *[][]rune) *TopoMap {
	cells := make([][]Cell, len(*data))
	for i, row := range *data {
		cellRows := make([]Cell, len((*data)[0]))
		for j, cell := range row {
			cellRows[j] = Cell{
				row:   i,
				col:   j,
				value: int8(cell - '0'),
			}
		}
		cells[i] = cellRows
	}
	return &TopoMap{
		cells:  cells,
		width:  len((*data)[0]),
		height: len(*data),
	}
}

func (tm *TopoMap) GetCell(row, col int) *Cell {
	return &tm.cells[row][col]
}

func (tm *TopoMap) findStartingCells() *[]Cell {
	trails := make([]Cell, 0)
	for i, row := range tm.cells {
		for j, cell := range row {
			if cell.value == 0 {
				trails = append(trails, *tm.GetCell(i, j))
			}
		}
	}
	return &trails
}

func (tm *TopoMap) IsOutOfBounds(row, col int) bool {
	return row < 0 || row >= tm.height || col < 0 || col >= tm.width
}

func (tm *TopoMap) GetNeighbours(cell *Cell) *[]Cell {
	neighbours := make([]Cell, 0)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if math.Abs(i) == 1 && math.Abs(j) == 1 {
				continue
			}
			if !tm.IsOutOfBounds(cell.row+i, cell.col+j) {
				neighbours = append(neighbours, *tm.GetCell(cell.row+i, cell.col+j))
			}
		}
	}
	return &neighbours
}

func (tm *TopoMap) GetTrails(th *Cell) int {
	cellCandidates := make([]Cell, 0)
	cellCandidates = append(cellCandidates, *th)
	topCells := make(map[Cell]struct{})
	for len(cellCandidates) > 0 {
		// Get the first cell and pop it
		cell := cellCandidates[0]
		cellCandidates = cellCandidates[1:]
		// Get the neighbours
		neighbours := tm.GetNeighbours(&cell)
		for _, neighbour := range *neighbours {
			if neighbour.value == cell.value+1 {
				if neighbour.value == 9 {
					topCells[neighbour] = struct{}{}
				} else {
					cellCandidates = append(cellCandidates, neighbour)
				}
			}
		}
	}
	return len(topCells)
}
func (tm *TopoMap) GetTrails2(th *Cell) int {
	cellCandidates := make([]Cell, 0)
	cellCandidates = append(cellCandidates, *th)
	topCells := make(map[Cell]int)
	for len(cellCandidates) > 0 {
		// Get the first cell and pop it
		cell := cellCandidates[0]
		cellCandidates = cellCandidates[1:]
		// Get the neighbours
		neighbours := tm.GetNeighbours(&cell)
		for _, neighbour := range *neighbours {
			if neighbour.value == cell.value+1 {
				if neighbour.value == 9 {
					if _, ok := topCells[neighbour]; !ok {
						topCells[neighbour] = 0
					}
					topCells[neighbour]++
				} else {
					cellCandidates = append(cellCandidates, neighbour)
				}
			}
		}
	}
	totalTrails := 0
	for _, v := range topCells {
		totalTrails += v
	}
	return totalTrails
}

func Solution1() int {
	data, err := reader.ReadCharFile("./inputs/day10.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	topoMap := NewTopoMapWithRune(&data)
	trailStarters := topoMap.findStartingCells()
	trailsCount := 0
	for _, cell := range *trailStarters {
		trailsCount += topoMap.GetTrails(&cell)
	}
	return trailsCount
}
func Solution2() int {
	data, err := reader.ReadCharFile("./inputs/day10.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	topoMap := NewTopoMapWithRune(&data)
	trailStarters := topoMap.findStartingCells()
	trailsCount := 0
	for _, cell := range *trailStarters {
		trailsCount += topoMap.GetTrails2(&cell)
	}
	return trailsCount
}
