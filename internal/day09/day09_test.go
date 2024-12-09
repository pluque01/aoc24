package day09

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	testTable := []struct {
		input    []rune
		expected []rune
	}{
		{
			[]rune{'2', '3', '3', '3', '1', '3', '3', '1', '2', '1', '4', '1', '4', '1', '3', '1', '4', '0', '2'},
			[]rune{'0', '0', '.', '.', '.', '1', '1', '1', '.', '.', '.', '2', '.', '.', '.', '3', '3', '3', '.', '4', '4', '.', '5', '5', '5', '5', '.', '6', '6', '6', '6', '.', '7', '7', '7', '.', '8', '8', '8', '8', '9', '9'},
		},
		{
			[]rune{'1', '2', '3', '4', '5'},
			[]rune{'0', '.', '.', '1', '1', '1', '.', '.', '.', '.', '2', '2', '2', '2', '2'},
		},
	}
	for _, test := range testTable {
		output := ParseDisk(&test.input)
		if !reflect.DeepEqual(*output, test.expected) {
			t.Errorf("\nExpected:\n%v\nGot:\n%v\n", string(test.expected), string(*output))
		}
	}
}

func TestCompactDisk(t *testing.T) {
	testTable := []struct {
		input    []rune
		expected []rune
	}{
		{
			[]rune{'0', '.', '.', '1', '1', '1', '.', '.', '.', '.', '2', '2', '2', '2', '2'},
			[]rune{'0', '2', '2', '1', '1', '1', '2', '2', '2', '.', '.', '.', '.', '.', '.'},
		},
		{
			[]rune{'0', '0', '.', '.', '.', '1', '1', '1', '.', '.', '.', '2', '.', '.', '.', '3', '3', '3', '.', '4', '4', '.', '5', '5', '5', '5', '.', '6', '6', '6', '6', '.', '7', '7', '7', '.', '8', '8', '8', '8', '9', '9'},
			[]rune{'0', '0', '9', '9', '8', '1', '1', '1', '8', '8', '8', '2', '7', '7', '7', '3', '3', '3', '6', '4', '4', '6', '5', '5', '5', '5', '6', '6', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		},
	}

	for _, test := range testTable {
		output := make([]rune, len(test.input))
		copy(output, test.input)
		CompactDisk(&output)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("\nExpected:\n%v\nGot:\n%v\n", string(test.expected), string(output))
		}
	}
}

func TestGetChecksum(t *testing.T) {
	testTable := []struct {
		input    []rune
		expected int
	}{
		{
			[]rune{'0', '0', '9', '9', '8', '1', '1', '1', '8', '8', '8', '2', '7', '7', '7', '3', '3', '3', '6', '4', '4', '6', '5', '5', '5', '5', '6', '6', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
			1928,
		},
	}

	for _, test := range testTable {
		output := GetChecksum(&test.input)
		if output != test.expected {
			t.Errorf("\nExpected:\n%v\nGot:\n%v\n", test.expected, output)
		}
	}
}

func TestCompactDiskWithoutFragmentation(t *testing.T) {
	testTable := []struct {
		input    []rune
		expected []rune
	}{
		{
			[]rune{'0', '0', '.', '.', '.', '1', '1', '1', '.', '.', '.', '2', '.', '.', '.', '3', '3', '3', '.', '4', '4', '.', '5', '5', '5', '5', '.', '6', '6', '6', '6', '.', '7', '7', '7', '.', '8', '8', '8', '8', '9', '9'},
			[]rune{'0', '0', '9', '9', '2', '1', '1', '1', '7', '7', '7', '.', '4', '4', '.', '3', '3', '3', '.', '.', '.', '.', '5', '5', '5', '5', '.', '6', '6', '6', '6', '.', '.', '.', '.', '.', '8', '8', '8', '8', '.', '.'},
		},
	}

	for _, test := range testTable {
		output := make([]rune, len(test.input))
		copy(output, test.input)
		CompactDiskWithoutFragmentation(&output)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("\nExpected:\n%v\nGot:\n%v\n", string(test.expected), string(output))
		}
	}
}
