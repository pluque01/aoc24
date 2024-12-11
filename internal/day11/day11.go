package day11

import (
	"log"
	"os"
	"pluque01/aoc24/pkg/reader"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

type MagicStones struct {
	stones []string
	cycles int
}

func NewMagicStones(stones []string) *MagicStones {
	return &MagicStones{stones: stones, cycles: 0}
}

func (ms *MagicStones) GetStones() []string {
	return ms.stones
}

func (ms *MagicStones) SplitStone(index int) {
	firstHalf := ms.stones[index][:len(ms.stones[index])/2]
	secondHalf := ms.stones[index][len(ms.stones[index])/2:]
	secondHalfInt, err := strconv.Atoi(secondHalf)
	if err != nil {
		log.Fatalf("Could not convert secondHalf to int: %s\nStone at Index = %v\nFirst Half = %s\nCycle: %d", secondHalf, ms.stones[index], firstHalf, ms.cycles)
	}
	ms.stones = append(ms.stones[:index], append([]string{firstHalf, strconv.Itoa(secondHalfInt)}, ms.stones[index+1:]...)...)
}

func (ms *MagicStones) GetStoneValue(index int) int {
	value, err := strconv.Atoi(ms.stones[index])
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func (ms *MagicStones) RunCycle() {
	for i := len(ms.stones) - 1; i >= 0; i-- {
		if ms.stones[i] == "0" {
			ms.stones[i] = "1"
		} else if len(ms.stones[i])%2 == 0 {
			ms.SplitStone(i)
		} else {
			ms.stones[i] = strconv.Itoa(ms.GetStoneValue(i) * 2024)
		}
	}
	ms.cycles++
	log.Printf("Completed cycle %d\n", ms.cycles)
}

func (ms *MagicStones) GetLength() int {
	return len(ms.stones)
}

func Solution1() int {
	data, err := reader.ReadStringByLineFile("/home/fallen/code/aoc24/inputs/day11.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	if len(data) != 1 {
		log.Fatalf("Expected 1 lines of input, got %d", len(data))
	}
	stones := strings.Split(data[0], " ")

	ms := NewMagicStones(stones)
	nCycles := 25
	for i := 0; i < nCycles; i++ {
		ms.RunCycle()
	}
	return ms.GetLength()
}

func lenLoop(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}
func SplitIntegerByMiddle(n int) (int, int) {
	// Calculate the number of digits manually
	digits := 0
	temp := n
	for temp > 0 {
		temp /= 10
		digits++
	}

	// Determine the divisor for splitting
	mid := digits / 2
	divisor := 1
	for i := 0; i < mid; i++ {
		divisor *= 10
	}

	// Split the number
	firstHalf := n / divisor
	secondHalf := n % divisor

	return firstHalf, secondHalf
}

func RunSimulationInteger(value int, iterations int) int {
	if v, ok := StoneGenerationMap[StoneGeneration{value, iterations}]; ok {
		return v
	}
	if iterations == 0 {
		return 0
	}
	stonesGenerated := 0
	var newValue int
	if value == 0 {
		newValue = 1
	} else if lenLoop(value)%2 == 0 {
		stonesGenerated++
		var newStone int
		newValue, newStone = SplitIntegerByMiddle(value)
		g := RunSimulationInteger(newStone, iterations-1)
		StoneGenerationMap[StoneGeneration{newStone, iterations - 1}] = g
		stonesGenerated += g
	} else {
		newValue = value * 2024
	}

	g := RunSimulationInteger(newValue, iterations-1)
	StoneGenerationMap[StoneGeneration{newValue, iterations - 1}] = g
	stonesGenerated += g
	return stonesGenerated
}

type StoneGeneration struct {
	stone     int
	iteration int
}

var StoneGenerationMap = make(map[StoneGeneration]int)

func Solution2() int {
	f, perr := os.Create("profile.pprof")
	if perr != nil {
		log.Fatal(perr)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	data, err := reader.ReadStringByLineFile("/home/fallen/code/aoc24/inputs/day11.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	if len(data) != 1 {
		log.Fatalf("Expected 1 lines of input, got %d", len(data))
	}
	stones := strings.Split(data[0], " ")
	stonesInt := make([]int, len(stones))
	for i, stone := range stones {
		stonesInt[i], _ = strconv.Atoi(stone)
	}
	nCycles := 75
	stonesGenerated := len(stones)
	totalTime := time.Now()
	for i, stone := range stonesInt {
		log.Print("Computing stone: ", i)
		start := time.Now()
		stonesGenerated += RunSimulationInteger(stone, nCycles)
		log.Printf("Stone: %s, Stones Generated: %d\n", stone, stonesGenerated)
		elapsed := time.Since(start)
		log.Printf("Iteration took %s", elapsed)
	}
	log.Printf("Total time: %s", time.Since(totalTime))

	return stonesGenerated
}
