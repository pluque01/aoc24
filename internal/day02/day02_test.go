package day02

import "testing"

func TestIsReportSafeWithDampener(t *testing.T) {
	testTable := []struct {
		report []int
		isSafe bool
	}{
		{[]int{21, 18, 15, 12, 9, 7, 6}, true},
		{[]int{1, 3, 6, 9, 12, 12}, true},
		{[]int{1, 3, 6, 9, 9, 9}, false},
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, true},
		{[]int{1, 3, 6, 7, 9}, true},
		{[]int{17, 21, 23, 24, 31, 33, 34, 31}, false},
		{[]int{9, 13, 16, 14, 15, 16, 17, 20}, false},
		{[]int{69, 69, 64, 62, 58}, false},
		{[]int{65, 65, 64, 62, 60}, true},
		{[]int{1, 0, 3, 4, 5}, true},
		{[]int{1, 0, 1, 4, 5}, true},
		{[]int{1, 0, 0, 4, 5}, false},
		{[]int{1, 2, 3, 4, 2}, true},
	}

	for _, tt := range testTable {
		result := isReportSafeWithDampener(tt.report)
		if result != tt.isSafe {
			t.Errorf("isReportSafeWithDampener(%v, true) = %v; want %v", tt.report, result, tt.isSafe)
		}
	}
}
