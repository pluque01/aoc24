package day01

import (
	"log"
	"pluque01/aoc24/pkg/array"
	pmath "pluque01/aoc24/pkg/math"
	"pluque01/aoc24/pkg/reader"
	"sort"
)

func CompareDistances() []int {
	content, err := reader.ReadColumnIntInput("./inputs/day01.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	c1 := content[0]
	c2 := content[1]

	if len(c1) != len(c2) {
		log.Fatalf("The columns have different lengths")
	}
	// sort the slices
	sort.Ints(c1)
	sort.Ints(c2)

	output := make([]int, len(c1))

	for i := 0; i < len(c1); i++ {
		output[i] = pmath.Abs(c1[i] - c2[i])
	}

	return output
}

func GetSimilarity() []int {
	content, err := reader.ReadColumnIntInput("./inputs/day01.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	c1 := content[0]
	c2 := content[1]

	if len(c1) != len(c2) {
		log.Fatalf("The columns have different lengths")
	}

	column2map := make(map[int]int)
	for _, value := range c2 {
		column2map[value] += 1
	}

	output := make([]int, len(c1))
	for i, value := range c1 {
		output[i] = value * column2map[value]
	}

	return output
}

func Solution1() int {
	distances := CompareDistances()
	totalSum := 0
	for i := 0; i < len(distances); i++ {
		totalSum += distances[i]
	}
	return totalSum
}

func Solution2() int {
	similarities := GetSimilarity()
	return array.Sum(similarities)
}
