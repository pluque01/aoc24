package day03

import (
	"fmt"
	"log"
	"pluque01/aoc24/pkg/reader"
	"regexp"
	"strconv"
)

func RemoveCorruptedMemory(s string) []string {
	rg := regexp.MustCompile(`mul\(\d+,\d+\)`)
	return rg.FindAllString(s, -1)
}

func ExecMul(mul string) (int, error) {
	rg := regexp.MustCompile(`\d+`)
	matches := rg.FindAllString(mul, -1)
	if len(matches) != 2 {
		return 0, fmt.Errorf("Numbers inside mul is different from 2: %v, %v, %v", len(matches), mul, matches)
	}
	mulNumbers := make([]int, 0)
	for _, value := range matches {
		v, err := strconv.Atoi(value)
		if err != nil {
			return 0, err
		}
		mulNumbers = append(mulNumbers, v)
	}
	return mulNumbers[0] * mulNumbers[1], nil

}

func Solution1() int {
	content, err := reader.ReadStringByLineFile("./inputs/day03.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	totalMul := 0
	for _, s := range content {
		cleanMuls := RemoveCorruptedMemory(s)
		for _, c := range cleanMuls {
			mul, err := ExecMul(c)
			if err != nil {
				log.Fatal(err)
			}
			totalMul += mul
		}
	}
	return totalMul
}

func RemoveCorruptedMemoryWithDo(s string) []string {
	rg := regexp.MustCompile(`(mul\(\d+,\d+\))|(do\(\))|(don't\(\))`)
	matches := rg.FindAllString(s, -1)
	log.Printf("Matches: %v", matches)

	mulActive := true
	var output []string
	for _, m := range matches {
		if m == "do()" {
			mulActive = true
		} else if m == "don't()" {
			mulActive = false
		} else if mulActive {
			output = append(output, m)
		}
	}
	return output
}

func Solution2() int {
	content, err := reader.ReadStringByLineFile("./inputs/day03.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	joinedContent := ""
	for _, s := range content {
		joinedContent += s
	}
	totalMul := 0
	cleanMul := RemoveCorruptedMemoryWithDo(joinedContent)
	for _, c := range cleanMul {
		mul, err := ExecMul(c)
		if err != nil {
			log.Fatal(err)
		}
		totalMul += mul
	}

	if err != nil {
		log.Fatal(err)
	}
	return totalMul

}
