package day10

import (
	"testing"
)

var testTopo = [][]int8{
	{8, 9, 0, 1, 0, 1, 2, 3},
	{7, 8, 1, 2, 1, 8, 7, 4},
	{8, 7, 4, 3, 0, 9, 6, 5},
	{9, 6, 5, 4, 9, 8, 7, 4},
	{4, 5, 6, 7, 8, 9, 0, 3},
	{3, 2, 0, 1, 9, 0, 1, 2},
	{0, 1, 3, 2, 9, 8, 0, 1},
	{1, 0, 4, 5, 6, 7, 3, 2},
}

func TestGetTrails(t *testing.T) {
	testTable := []struct {
		input    Cell
		expected int
	}{
		{Cell{0, 2, 0}, 5},
		{Cell{0, 4, 0}, 6},
		{Cell{2, 4, 0}, 5},
		{Cell{4, 6, 0}, 3},
	}

	tm := NewTopoMap(&testTopo)
	for _, test := range testTable {
		output := (*tm).GetTrails(&test.input)
		if output != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, output)
		}
	}
}
