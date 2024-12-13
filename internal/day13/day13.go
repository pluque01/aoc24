package day13

import (
	"fmt"
	"log"
	"math"
	"pluque01/aoc24/pkg/reader"
	"regexp"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

func hasDecimal(f float64) bool {
	diff := math.Abs(math.Round(f) - f)
	return diff > 0.01
}

func SolveEquation(a []float64, b []float64) (int64, int64, error) {
	// Create the matrix A and vector b.
	A := mat.NewDense(2, 2, a)
	B := mat.NewVecDense(2, b)
	// log.Printf("A: %v", A)
	// log.Printf("B: %v", B)

	// Solve the equations using the Solve function.
	var x mat.VecDense
	if err := x.SolveVec(A, B); err != nil {
		return 0, 0, err
	}

	aFloat := x.At(0, 0)
	bFloat := x.At(1, 0)

	// if aFloat-math.Trunc(aFloat) > 0.001 || bFloat-math.Trunc(bFloat) > 0.001 {
	// 	log.Printf("Non decimal solutions: aFloat: %f, bFloat: %f", aFloat, bFloat)
	// 	return 0, 0, fmt.Errorf("No solution")
	// }
	if hasDecimal(aFloat) || hasDecimal(bFloat) {
		log.Printf("Non decimal solutions: aFloat: %f, bFloat: %f", aFloat, bFloat)
		return 0, 0, fmt.Errorf("No solution")
	}

	aInt := int64(x.At(0, 0))
	bInt := int64(x.At(1, 0))

	return aInt, bInt, nil
}

func ParseEquation(eq1 string, eq2 string, sol string) ([]float64, []float64) {
	rg := regexp.MustCompile(`\d+`)
	numbers := make([]string, 0)
	numbers = append(numbers, rg.FindAllString(eq1, -1)...)
	numbers = append(numbers, rg.FindAllString(eq2, -1)...)
	numbers = append(numbers, rg.FindAllString(sol, -1)...)
	if len(numbers) != 6 {
		log.Panicf("Could not parse equation:\n%s\n%s\n%s", eq1, eq2, sol)
	}
	numbersFloat := make([]float64, 6)
	for i, n := range numbers {
		converted, err := strconv.Atoi(n)
		if err != nil {
			log.Panicf("Could not convert to int: %s", n)
		}
		numbersFloat[i] = float64(converted)
	}

	aux := numbersFloat[1]
	numbersFloat[1] = numbersFloat[2]
	numbersFloat[2] = aux
	return numbersFloat[:4], numbersFloat[4:6]
}

func ParseEquation2(eq1 string, eq2 string, sol string) ([]float64, []float64) {
	rg := regexp.MustCompile(`\d+`)
	numbers := make([]string, 0)
	numbers = append(numbers, rg.FindAllString(eq1, -1)...)
	numbers = append(numbers, rg.FindAllString(eq2, -1)...)
	numbers = append(numbers, rg.FindAllString(sol, -1)...)
	if len(numbers) != 6 {
		log.Panicf("Could not parse equation:\n%s\n%s\n%s", eq1, eq2, sol)
	}
	numbersFloat := make([]float64, 6)
	for i, n := range numbers {
		converted, err := strconv.Atoi(n)
		if err != nil {
			log.Panicf("Could not convert to int: %s", n)
		}
		numbersFloat[i] = float64(converted)
	}

	aux := numbersFloat[1]
	numbersFloat[1] = numbersFloat[2]
	numbersFloat[2] = aux
	numbersFloat[4] += 10000000000000
	numbersFloat[5] += 10000000000000
	return numbersFloat[:4], numbersFloat[4:6]
}

func Solution1() int {
	data, err := reader.ReadStringByLineFile("/home/fallen/code/aoc24/inputs/day13.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	// fmt.Print(data)
	tokens := 0
	for i := 0; i < len(data); i += 3 {
		eq1, eq2 := ParseEquation(data[i], data[i+1], data[i+2])
		a, b, err := SolveEquation(eq1, eq2)
		if err == nil {
			tokens += 3 * int(a)
			tokens += int(b)
		}
	}
	return tokens
}

func Solution2() int {
	data, err := reader.ReadStringByLineFile("/home/fallen/code/aoc24/inputs/day13.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	// fmt.Print(data)
	tokens := 0
	equationsParsed := 0
	for i := 0; i < len(data); i += 3 {
		eq1, eq2 := ParseEquation2(data[i], data[i+1], data[i+2])
		equationsParsed++
		a, b, err := SolveEquation(eq1, eq2)
		if err == nil {
			tokens += 3 * int(a)
			tokens += int(b)
		}
	}
	log.Printf("Equations parsed: %d", equationsParsed)
	return tokens
}
