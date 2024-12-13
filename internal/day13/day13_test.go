package day13

import "testing"

func TestSolveEquation(t *testing.T) {
	testTable := []struct {
		eq1         string
		eq2         string
		sol         string
		expectedA   int64
		expectedB   int64
		expectedErr error
	}{
		{
			"Button A: X+94, Y+34",
			"Button B: X+22, Y+67",
			"Prize: X=8400, Y=5400",
			80, 40, nil,
		},
	}

	for _, tt := range testTable {
		valuesA, valuesB := ParseEquation(tt.eq1, tt.eq2, tt.sol)
		a, b, err := SolveEquation(valuesA, valuesB)
		if tt.expectedErr != err {
			t.Errorf("Expected error %v, got %v", tt.expectedErr, err)
		} else if a != tt.expectedA || b != tt.expectedB {
			t.Errorf("Expected %d, %d but got %d, %d", tt.expectedA, tt.expectedB, a, b)
		}
	}
}
