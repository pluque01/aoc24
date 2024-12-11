package day11

import (
	"log"
	"reflect"
	"testing"
)

func TestRunCycle(t *testing.T) {
	testData := struct {
		stones   []string
		expected [][]string
	}{
		[]string{"125", "17"},
		[][]string{
			[]string{"253000", "1", "7"},
			[]string{"253", "0", "2024", "14168"},
			[]string{"512072", "1", "20", "24", "28676032"},
			[]string{"512", "72", "2024", "2", "0", "2", "4", "2867", "6032"},
			[]string{"1036288", "7", "2", "20", "24", "4048", "1", "4048", "8096", "28", "67", "60", "32"},
			[]string{"2097446912", "14168", "4048", "2", "0", "2", "4", "40", "48", "2024", "40", "48", "80", "96", "2", "8", "6", "7", "6", "0", "3", "2"},
		},
	}
	ms := NewMagicStones(testData.stones)
	for _, expected := range testData.expected {
		ms.RunCycle()
		log.Printf("Stones: %v\n", ms.GetStones())
		if !reflect.DeepEqual(ms.stones, expected) {
			t.Errorf("Expected %v, got %v", expected, ms.GetStones())
		}
	}
}

func TestSolution(t *testing.T) {
	Solution1()
}
