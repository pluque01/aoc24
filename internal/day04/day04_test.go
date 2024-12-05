package day04

import "testing"

func TestCountWord2D(t *testing.T) {
	testTable := []struct {
		grid   [][]rune
		target string
		result int
	}{
		{[][]rune{
			{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
			{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
			{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
			{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
			{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
			{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
			{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
			{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
			{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
			{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
		}, "XMAS", 18},
	}

	for _, tt := range testTable {
		r := CountWord2D(tt.grid, tt.target)
		if r != tt.result {
			t.Errorf("CountWord2D(%v, %v) = %v; want %v", tt.grid, tt.target, r, tt.result)
		}
	}
}

func TestCountCrossWord2D(t *testing.T) {
	testTable := []struct {
		grid   [][]rune
		target string
		result int
	}{
		{[][]rune{
			{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
			{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
			{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
			{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
			{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
			{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
			{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
			{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
			{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
			{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
		}, "MAS", 9},
		// {[][]rune{
		// 	{'M', '.', 'M'},
		// 	{'.', 'A', '.'},
		// 	{'S', '.', 'S'},
		// }, "MAS", 1},
	}

	for _, tt := range testTable {
		r := CountCrossWord2D(tt.grid, tt.target)
		if r != tt.result {
			t.Errorf("CountWord2D(%v, %v) = %v; want %v", tt.grid, tt.target, r, tt.result)
		}
	}
}
