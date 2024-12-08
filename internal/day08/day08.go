package day08

import (
	"fmt"
	"log"
	"pluque01/aoc24/pkg/reader"
)

type Cell struct {
	row    int
	column int
}
type Map struct {
	antenas   map[rune][]Cell
	antinodes map[Cell]struct{}
	width     int
	height    int
}

func NewMap(m [][]rune) *Map {
	antenas := make(map[rune][]Cell)
	for i, row := range m {
		for j, cell := range row {
			if cell != '.' {
				antenas[cell] = append(antenas[cell], Cell{i, j})
			}
		}
	}
	return &Map{
		antenas:   antenas,
		antinodes: make(map[Cell]struct{}),
		width:     len(m[0]),
		height:    len(m),
	}
}

func (m *Map) isOutOfBounds(c *Cell) bool {
	return c.row < 0 || c.row >= m.height || c.column < 0 || c.column >= m.width
}

func (m *Map) addAntinode(c *Cell) {
	m.antinodes[*c] = struct{}{}
}

func getDistance(a Cell, b Cell) (int, int) {
	return b.row - a.row, b.column - a.column
}

func (m *Map) FindAntinodes() {
	for frequency := range m.antenas {
		for i := 0; i < len(m.antenas[frequency]); i++ {
			for j := i + 1; j < len(m.antenas[frequency]); j++ {
				drow, dcolumn := getDistance(m.antenas[frequency][i], m.antenas[frequency][j])
				c := Cell{m.antenas[frequency][i].row - drow, m.antenas[frequency][i].column - dcolumn}
				if !m.isOutOfBounds(&c) {
					m.addAntinode(&c)
				}
				c = Cell{m.antenas[frequency][j].row + drow, m.antenas[frequency][j].column + dcolumn}
				if !m.isOutOfBounds(&c) {
					m.addAntinode(&c)
				}
			}
		}
	}
}

func (m *Map) FindAntinodesRecalculated() {
	for frequency := range m.antenas {
		for i := 0; i < len(m.antenas[frequency]); i++ {
			for j := i + 1; j < len(m.antenas[frequency]); j++ {
				drow, dcolumn := getDistance(m.antenas[frequency][i], m.antenas[frequency][j])
				c := Cell{m.antenas[frequency][i].row + drow, m.antenas[frequency][i].column + dcolumn}
				for !m.isOutOfBounds(&c) {
					m.addAntinode(&c)
					c.row += drow
					c.column += dcolumn
				}
				c = Cell{m.antenas[frequency][j].row - drow, m.antenas[frequency][j].column - dcolumn}
				for !m.isOutOfBounds(&c) {
					m.addAntinode(&c)
					c.row -= drow
					c.column -= dcolumn
				}
			}
		}
	}
}

func (m *Map) CountAntinodes() int {
	return len(m.antinodes)
}

func (m *Map) GetWidth() int {
	return m.width
}

func (m *Map) GetHeight() int {
	return m.height
}

func Solution1() int {
	data, err := reader.ReadCharFile("./inputs/day08.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	m := NewMap(data)
	m.FindAntinodes()
	fmt.Printf("Width: %d, Height: %d\n", m.GetWidth(), m.GetHeight())
	return m.CountAntinodes()
}

func Solution2() int {
	data, err := reader.ReadCharFile("./inputs/day08.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	m := NewMap(data)
	m.FindAntinodesRecalculated()
	return m.CountAntinodes()
}
