package day03

import (
	"pluque01/aoc24/pkg/cmp"
	"testing"
)

func TestRemoveCorruptedMemory(t *testing.T) {
	testTable := []struct {
		value  string
		result []string
	}{
		{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			[]string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}},
	}
	for _, tt := range testTable {
		r := RemoveCorruptedMemory(tt.value)
		if !cmp.AreEquals(r, tt.result) {
			t.Errorf("ExecMul(%v) = %v; want %v", tt.value, r, tt.result)
		}
	}
}

func TestExecMul(t *testing.T) {
	testTable := []struct {
		value  string
		result int
	}{
		{"mul(3,2)", 6},
		{"mul(1,1)", 1},
		{"mul(2,0)", 0},
		// {"mul(2,-1)", -2},
	}

	for _, tt := range testTable {
		r, err := ExecMul(tt.value)
		if err != nil {
			t.Fatal(err)
		}
		if r != tt.result {
			t.Errorf("ExecMul(%v) = %v; want %v", tt.value, r, tt.result)
		}
	}
}

func TestRemoveCorruptedMemoryWithDo(t *testing.T) {
	testTable := []struct {
		value  string
		result []string
	}{
		{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			[]string{"mul(2,4)", "mul(8,5)"}},
		{"do():who(268,689)*mul(941,59))-{&&(mul(296,702):@)where()don't()why()when(),mul(405,639)~$),who(575,232),%mul(671,828)",
			[]string{"mul(941,59)", "mul(296,702)"}},
	}
	for _, tt := range testTable {
		r := RemoveCorruptedMemoryWithDo(tt.value)
		if !cmp.AreEquals(r, tt.result) {
			t.Errorf("ExecMul(%v) = %v; want %v", tt.value, r, tt.result)
		}
	}
}
