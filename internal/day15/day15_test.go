package day15

import "testing"

func TestSumGPSCoordinates(t *testing.T) {
	testData := []struct {
		input    [][]rune
		expected int
	}{
		{
			input: [][]rune{
				{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
				{'#', '.', 'O', '.', 'O', '.', 'O', 'O', 'O', '#'},
				{'#', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
				{'#', 'O', 'O', '.', '.', '.', '.', '.', '.', '#'},
				{'#', 'O', 'O', '@', '.', '.', '.', '.', '.', '#'},
				{'#', 'O', '#', '.', '.', '.', '.', '.', 'O', '#'},
				{'#', 'O', '.', '.', '.', '.', '.', 'O', 'O', '#'},
				{'#', 'O', '.', '.', '.', '.', '.', 'O', 'O', '#'},
				{'#', 'O', 'O', '.', '.', '.', '.', 'O', 'O', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
			},
			expected: 10092,
		},
	}
	for _, tt := range testData {
		w := NewWarehouse(tt.input)
		actual := w.SumGPSCoordinates()
		if actual != tt.expected {
			t.Errorf("Expected %d, but got %d", tt.expected, actual)
		}
	}

}
