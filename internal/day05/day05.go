package day05

import (
	"log"
	"pluque01/aoc24/internal/day05/printer"
	"pluque01/aoc24/pkg/reader"
)

func Solution1And2() (int, int) {
	content, err := reader.ReadStringByLineFile("./inputs/day05.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	p := printer.NewPrinter(content)
	goodUpdates, badUpdates := p.GetGoodAndBadUpdates()
	fixedUpdates := p.FixBadUpdates(badUpdates)
	return printer.SumMiddleElements(goodUpdates), printer.SumMiddleElements(fixedUpdates)
}
