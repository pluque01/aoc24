package main

import (
	"log"
	"pluque01/aoc24/internal/day05"
)

func main() {
	g, f := day05.Solution1And2()
	log.Printf("Solutions: \n GoodUpdatesSum: %d, FixedUpdatesSum: %d\n", g, f)
}
